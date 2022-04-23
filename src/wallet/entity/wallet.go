package entity

import "time"

type (
	DBWallet struct {
		WalletID int64     `json:"wallet_id" db:"wallet_id"`
		UserID   int64     `json:"user_id" db:"user_id"`
		Amount   int64     `json:"amount" db:"amount"`
		Date     time.Time `json:"date" db:"date"`
	}

	DBWalletHistories struct {
		LastAmount int64     `json:"last_amount" db:"last_amount"`
		Date       time.Time `json:"date" db:"date"`
	}

	GetDetailWallet struct {
		WalletID      int64               `json:"wallet_id" db:"wallet_id"`
		UserID        int64               `json:"user_id" db:"user_id"`
		Amount        int64               `json:"amount" db:"amount"`
		Date          time.Time           `json:"date" db:"date"`
		HistoryWallet []DBWalletHistories `json:"history_wallet" db:"history_wallet"`
	}

	ReqUseWallet struct {
		UserID int64 `json:"user_id" db:"user_id"`
		Amount int64 `json:"amount" db:"amount"`
	}
)
