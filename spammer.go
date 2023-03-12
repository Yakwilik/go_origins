package main

import (
	"fmt"
	"gopkg.in/typ.v4/slices"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	chans := make([]chan interface{}, 1)
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for i := range cmds {
		chans = append(chans, make(chan interface{}))
		wg.Add(1)
		go DoneMaker(wg, cmds[i], chans[i], chans[i+1])
	}

}

func DoneMaker(waiter *sync.WaitGroup, job cmd, in chan interface{}, out chan interface{}) {
	defer waiter.Done()
	job(in, out)
	close(out)
}

type userSyncTools struct {
	wg               *sync.WaitGroup
	uniqueUsers      map[string]bool
	uniqueUsersMutex *sync.RWMutex
}

func getUserIteration(out chan interface{}, rawData interface{}, userSync userSyncTools) {
	defer userSync.wg.Done()
	userEmail, ok := rawData.(string)
	if !ok {
		return
	}
	user := GetUser(userEmail)
	userSync.uniqueUsersMutex.Lock()
	defer userSync.uniqueUsersMutex.Unlock()
	if _, exists := userSync.uniqueUsers[user.Email]; !exists {
		userSync.uniqueUsers[user.Email] = true
		out <- user
	}
}

func SelectUsers(in, out chan interface{}) {
	userSyncTool := userSyncTools{
		wg:               &sync.WaitGroup{},
		uniqueUsers:      make(map[string]bool, 0),
		uniqueUsersMutex: &sync.RWMutex{},
	}
	defer userSyncTool.wg.Wait()

	for inChanData := range in {
		userSyncTool.wg.Add(1)
		go getUserIteration(out, inChanData, userSyncTool)
	}
}

func selectMessagesIteration(out chan interface{}, users []User, wg *sync.WaitGroup) {
	defer wg.Done()
	messages, err := GetMessages(users...)
	if err != nil {
		return
	}
	for _, msgID := range messages {
		out <- msgID
	}
}

func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	users := make([]User, 0, GetMessagesMaxUsersBatch)
	for inChanData := range in {
		users = append(users, inChanData.(User))
		if len(users) == GetMessagesMaxUsersBatch {
			wg.Add(1)
			go selectMessagesIteration(out, slices.Clone(users), wg)
			users = make([]User, 0, GetMessagesMaxUsersBatch)
		}
	}
	if len(users) > 0 {
		wg.Add(1)
		go selectMessagesIteration(out, users, wg)
	}
}

func hasSpamAsync(out chan interface{}, msgID MsgID, limiterChan chan interface{}, wg *sync.WaitGroup) {
	limiterChan <- true
	defer func() {
		<-limiterChan
		wg.Done()
	}()

	result, err := HasSpam(msgID)
	if err != nil {
		return
	}
	msgData := MsgData{
		ID:      msgID,
		HasSpam: result,
	}
	out <- msgData

}

func CheckSpam(in, out chan interface{}) {
	concurrentChan := make(chan interface{}, HasSpamMaxAsyncRequests)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	for inChanData := range in {
		wg.Add(1)
		go hasSpamAsync(out, inChanData.(MsgID), concurrentChan, wg)
	}
}

func CombineResults(in, out chan interface{}) {
	msgDataSlice := make([]MsgData, 0)

	for inChanData := range in {
		msgDataSlice = append(msgDataSlice, inChanData.(MsgData))
	}
	sort.Slice(msgDataSlice, func(i, j int) bool {
		if msgDataSlice[i].HasSpam == msgDataSlice[j].HasSpam {
			return msgDataSlice[i].ID < msgDataSlice[j].ID
		} else {
			return msgDataSlice[i].HasSpam
		}
	})
	for _, msgData := range msgDataSlice {
		out <- fmt.Sprintf("%t %d", msgData.HasSpam, msgData.ID)
	}
}
