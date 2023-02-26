package structs

type Session struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	UserID   int    `db:"user_id"`
	IAT      int64  `db:"iat"`
	EXP      int64  `db:"EXP"`
}
