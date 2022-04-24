package entity

import (
	"database/sql"
	"time"
)

type (
	DBWallet struct {
		WalletID int64         `json:"wallet_id" db:"wallet_id"`
		OrderID  sql.NullInt64 `json:"order_id" db:"order_id"`
		UserID   int64         `json:"user_id" db:"user_id"`
		Amount   int64         `json:"amount" db:"amount"`
		Date     time.Time     `json:"date" db:"date"`
	}

	DBWalletHistories struct {
		OrderID    sql.NullInt64 `json:"order_id" db:"order_id"`
		LastAmount int64         `json:"last_amount" db:"last_amount"`
		Date       time.Time     `json:"date" db:"date"`
	}

	DBWalletHistoriesDetail struct {
		WalletHistoryID int64         `json:"wallet_history_id" db:"wallet_history_id"`
		OrderID         sql.NullInt64 `json:"order_id" db:"order_id"`
		LastAmount      int64         `json:"last_amount" db:"last_amount"`
		Date            time.Time     `json:"date" db:"date"`
	}

	GetDetailWallet struct {
		WalletID      int64               `json:"wallet_id" db:"wallet_id"`
		UserID        int64               `json:"user_id" db:"user_id"`
		Amount        int64               `json:"amount" db:"amount"`
		OrderID       int64               `json:"order_id" db:"order_id"`
		Date          time.Time           `json:"date" db:"date"`
		HistoryWallet []DBWalletHistories `json:"history_wallet" db:"history_wallet"`
	}

	ReqUseWallet struct {
		UserID  int64 `json:"user_id" db:"user_id"`
		Amount  int64 `json:"amount" db:"amount"`
		OrderID int64 `json:"order_id" db:"order_id"`
	}
)
