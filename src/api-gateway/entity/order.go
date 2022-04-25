package entity

type (

	//   int64 product_id = 1;
	//   int64 user_id = 2;

	HTTPReqCreateOrder struct {
		ProductID int64 `json:"product_id" db:"product_id" validate:"required,gt=0"`
	}

	HTTPReqPostCreateOrder struct {
		ProductID int64 `json:"product_id" db:"product_id" validate:"required,gt=0"`
		UserID    int64 `json:"user_id" db:"user_id" validate:"required,gt=0"`
	}

	HTTPReqGetOrders struct {
		UserID int64 `json:"user_id" db:"user_id" validate:"required,gt=0"`
	}
)
