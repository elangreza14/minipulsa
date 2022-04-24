package entity

import (
	"database/sql"
	"time"
)

type (

	//   int64 order_id = 1;
	//   int64 product_id = 2;
	//   int64 user_id = 3;
	//   int64 wallet_history_id = 4;
	//   string status = 5;
	//   string date = 6;

	DBOrder struct {
		OrderID         int64         `json:"order_id" db:"order_id"`
		ProductID       int64         `json:"product_id" db:"product_id"`
		UserID          int64         `json:"user_id" db:"user_id"`
		WalletHistoryID sql.NullInt64 `json:"wallet_history_id" db:"wallet_history_id"`
		Status          string        `json:"status" db:"status"`
		Date            time.Time     `json:"date" db:"date"`
	}

	ReqCreateOrder struct {
		ProductID int64 `json:"product_id" db:"product_id"`
		UserID    int64 `json:"user_id" db:"user_id"`
	}

	ReqUpdateOrder struct {
		OrderID         int64  `json:"order_id" db:"order_id"`
		Status          string `json:"status" db:"status"`
		WalletHistoryID int64  `json:"wallet_history_id" db:"wallet_history_id"`
	}

	OrderStatus string
)

const OrderStatusPending OrderStatus = "PENDING"
const OrderStatusSUCCESS OrderStatus = "SUCCESS"
const OrderStatusFAILED OrderStatus = "FAILED"
