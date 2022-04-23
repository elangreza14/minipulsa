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

	HTTPReqPostPutUser struct {
		Email    string `json:"email" db:"email" validate:"required"`
		Password string `json:"password" db:"password" validate:"required"`
	}

	HTTPResPostPutUser struct {
		Token string `json:"token" db:"token" validate:"required"`
	}
)
