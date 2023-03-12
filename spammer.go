package main

import (
	"gopkg.in/typ.v4/slices"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	chans := make([]chan interface{}, 0, 6)
	chans = append(chans, make(chan interface{}))
	wg := &sync.WaitGroup{}
	for i := range cmds {
		chans = append(chans, make(chan interface{}))
		wg.Add(1)
		go DoneMaker(wg, cmds[i], chans[i], chans[i+1])
	}
	wg.Wait()

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

	for inChanData := range in {
		userSyncTool.wg.Add(1)
		go getUserIteration(out, inChanData, userSyncTool)
	}
	userSyncTool.wg.Wait()
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
	users := make([]User, 0, 2)
	for inChanData := range in {
		users = append(users, inChanData.(User))
		if len(users) == 2 {
			wg.Add(1)
			go selectMessagesIteration(out, slices.Clone(users), wg)
			users = make([]User, 0, 2)
		}
	}
	if len(users) > 0 {
		wg.Add(1)
		go selectMessagesIteration(out, users, wg)
	}
	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
}
