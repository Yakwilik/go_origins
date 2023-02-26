package authmysql

import (
	"github.com/blockloop/scan"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
)

func (a *AuthMySQL) CreateUser(user structs.User) (int, error) {

	exec, err := a.UsersDB.Exec("insert into `users` (username, password_hash) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return -1, err
	}
	id, err := exec.LastInsertId()

	return int(id), err
}

func (a *AuthMySQL) GetUser(username, password string) (structs.User, error) {
	row, err := a.UsersDB.Query("select * from users where username = (?) and password_hash = (?)", username, password)

	if err != nil {
		return structs.User{}, err
	}
	if row.Err() != nil {
		return structs.User{}, row.Err()
	}
	var user structs.User
	err = scan.Row(&user, row)
	if err != nil {
		return structs.User{}, err
	}

	return user, nil
}
