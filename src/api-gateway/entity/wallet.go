package entity

type (
	HTTPReqPostUseWallet struct {
		Amount int64 `json:"amount" db:"amount" validate:"required,gt=0"`
	}

	DBWallet struct {
		WalletID int64  `json:"wallet_id" db:"wallet_id"`
		UserID   int64  `json:"user_id" db:"user_id"`
		OrderID  int64  `json:"order_id" db:"order_id"`
		Amount   int64  `json:"amount" db:"amount"`
		Date     string `json:"date" db:"date"`
	}

	DBHistoryWallet struct {
		LastAmount int64  `json:"last_amount" db:"last_amount"`
		OrderID    int64  `json:"order_id" db:"order_id"`
		Date       string `json:"date" db:"date"`
	}

	GetDBWallet struct {
		Detail  DBWallet          `json:"detail" db:"detail"`
		History []DBHistoryWallet `json:"history" db:"history"`
	}
)
