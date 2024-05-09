package models

type User struct {
	//UserID   int64  `db:"user_id"`
	//Role     int64  `db:"role"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type ParamLogic struct {
	UserID int64 `db:"user_id"`
	//Role     int64  `db:"role"`
	Username string `db:"username"`
	Password string `db:"password"`
}
