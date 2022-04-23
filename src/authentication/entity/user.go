package entity

type (
	DBUser struct {
		UserID   int64  `db:"user_id"`
		Email    string `db:"email"`
		Password string `db:"password"`
		Token    string `db:"token"`
	}

	ReqPostPutUser struct {
		Email    string `db:"email"`
		Password string `db:"password"`
		Token    string `db:"token"`
	}
)
