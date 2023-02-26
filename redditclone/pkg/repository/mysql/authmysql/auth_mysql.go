package authmysql

import (
	"database/sql"
)

type AuthMySQL struct {
	SessionsDB *sql.DB
	UsersDB    *sql.DB
}

func NewAuthMySQL(sessionsDB, usersDB *sql.DB) *AuthMySQL {
	if sessionsDB == nil || usersDB == nil {
		return nil
	}
	return &AuthMySQL{
		SessionsDB: sessionsDB,
		UsersDB:    usersDB,
	}
}
