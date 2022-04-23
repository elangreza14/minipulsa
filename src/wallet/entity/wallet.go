package entity

import "time"

type (
	DBWallet struct {
		WalletID int64     `json:"wallet_id" db:"wallet_id"`
		UserID   int64     `json:"user_id" db:"user_id"`
		Amount   int64     `json:"amount" db:"amount"`
		Date     time.Time `json:"date" db:"date"`
	}

	ReqUseWallet struct {
		UserID int64 `json:"user_id" db:"user_id"`
		Amount int64 `json:"amount" db:"amount"`
	}
)
