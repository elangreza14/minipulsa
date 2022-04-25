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

	DBOrder struct {
		OrderID         int64  `json:"order_id" db:"order_id"`
		ProductID       int64  `json:"product_id" db:"product_id"`
		UserID          int64  `json:"user_id" db:"user_id"`
		WalletHistoryID int64  `json:"wallet_history_id" db:"wallet_history_id"`
		Status          string `json:"status" db:"status"`
		Date            string `json:"date" db:"date"`
	}
)
