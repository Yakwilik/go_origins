// NO_LINT
package authmysql

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/structs"
	"reflect"
	"testing"
	"time"
)

// go: go test -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

func TestAuthMySQL_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	auth := AuthMySQL{
		SessionsDB: db,
		UsersDB:    db,
	}
	user := structs.User{
		ID:       0,
		Username: "username",
		Password: "password",
	}
	mock.ExpectExec("insert into `users`").
		WithArgs(user.Username, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := auth.CreateUser(user)
	if err != nil {
		t.Errorf("error %v", err)
		return
	}
	if id != 1 {
		t.Errorf("want id = 1, have %v", id)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectExec("insert into `users`").
		WithArgs(user.Username, user.Password).
		WillReturnError(errors.New("error"))

	_, err = auth.CreateUser(user)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestAuthMySQL_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	auth := AuthMySQL{
		SessionsDB: db,
		UsersDB:    db,
	}
	user := structs.User{
		ID:       0,
		Username: "username",
		Password: "password",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password_hash"})
	expect := []*structs.User{{ID: 1, Username: user.Username, Password: user.Password}}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Password)
	}
	mock.
		ExpectQuery("select * ").
		WithArgs(user.Username, user.Password).
		WillReturnRows(rows)

	item, err := auth.GetUser(user.Username, user.Password)
	if err != nil {
		t.Errorf("error %v", err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(&item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	mock.
		ExpectQuery("select * ").
		WithArgs(user.Username, user.Password).
		WillReturnError(errors.New("db_error"))

	_, err = auth.GetUser(user.Username, user.Password)

	if err == nil {
		t.Error("db_error expected")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestAuthMySQL_GenerateToken(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	auth := AuthMySQL{
		SessionsDB: db,
		UsersDB:    db,
	}
	user := structs.User{
		ID:       0,
		Username: "username",
		Password: "password",
	}
	mock.ExpectExec("insert into sessions").
		WithArgs(user.Username, user.ID, time.Now().Unix(), time.Now().Add(time.Hour*7*24).Unix()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = auth.GenerateToken(user.Username, user.ID, time.Now().Unix(), time.Now().Add(time.Hour*7*24).Unix())
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	mock.ExpectExec("insert into sessions").
		WithArgs(user.Username, user.ID, time.Now().Unix(), time.Now().Add(time.Hour*7*24).Unix()).
		WillReturnError(errors.New("db_error"))

	err = auth.GenerateToken(user.Username, user.ID, time.Now().Unix(), time.Now().Add(time.Hour*7*24).Unix())
	if err == nil {
		t.Errorf("expected error, got %v", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestAuthMySQL_CheckIdentity(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	auth := AuthMySQL{
		SessionsDB: db,
		UsersDB:    db,
	}
	user := structs.Session{
		ID:       0,
		Username: "username",
		UserID:   0,
	}

	rows := sqlmock.NewRows([]string{"id", "username", "user_id", "iat", "exp"})
	expect := []*structs.Session{{ID: 1, Username: user.Username, UserID: 0, IAT: time.Now().Unix(), EXP: time.Now().Add(time.Hour * 7 * 24).Unix()}}
	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.UserID, item.IAT, item.EXP)
	}
	mock.
		ExpectQuery("select * ").
		WithArgs(user.Username, user.UserID).
		WillReturnRows(rows)

	err = auth.CheckIdentity(user.Username, user.UserID)
	if err != nil {
		t.Errorf("error %v", err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	mock.
		ExpectQuery("select * ").
		WithArgs(user.Username, user.UserID).
		WillReturnError(errors.New("bd_error"))

	err = auth.CheckIdentity(user.Username, user.UserID)

	if err == nil {
		t.Errorf("error expected, but %v", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	empty := sqlmock.NewRows([]string{"id", "username", "user_id", "iat", "exp"})
	mock.
		ExpectQuery("select * ").
		WithArgs(user.Username, user.UserID).
		WillReturnRows(empty)

	err = auth.CheckIdentity(user.Username, user.UserID)
	if err == nil {
		t.Errorf("error expected")
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestNewAuthMySQL(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))

	if err != nil {
		t.Errorf("error did not expected:%v", err)
	}
	mock.ExpectPing()

	auth := NewAuthMySQL(db, db)
	err = auth.SessionsDB.Ping()
	if err != nil {
		t.Errorf("error did not expected:%v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	auth = NewAuthMySQL(nil, nil)
	if auth != nil {
		t.Errorf("expected nil")
	}
}
