package authmysql

import (
	"errors"
	"github.com/blockloop/scan"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
)

func (a *AuthMySQL) GenerateToken(username string, id int, iat, exp int64) error {
	_, err := a.SessionsDB.Exec(`insert into sessions (username, user_id, iat, exp) VALUES (?, ?, ?, ?)`, username, id, iat, exp)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthMySQL) CheckIdentity(username string, id int) error {
	query, err := a.SessionsDB.Query(`select * from sessions where username = (?) and user_id = (?)`, username, id)
	if err != nil {
		return err
	}
	session := make([]structs.Session, 0)
	check := scan.Rows(&session, query)
	if check != nil {
		return check
	}
	if len(session) < 1 {
		return errors.New("no sessions with this user")
	}
	return nil

}
