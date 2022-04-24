package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/order/entity"
)

const (
	insertOrder = `
		INSERT INTO 
			"orders" 
		("product_id","user_id","status") 
		VALUES 
			($1,$2,$3) 
		RETURNING 
			"order_id",
			"product_id",
			"user_id",
			"wallet_history_id",
			"status",
			"date";
	`

	updateOrder = `
		UPDATE
			"orders" 
		SET
			"status"=$1,
			"wallet_history_id"=$2,
			"date"=NOW()
		WHERE
			"order_id" = $3
		RETURNING 
			"order_id",
			"product_id",
			"user_id",
			"wallet_history_id",
			"status",
			"date";
	`
	getOrder = `
		SELECT
			"order_id",
			"product_id",
			"user_id",
			"wallet_history_id",
			"status",
			"date"
		FROM
			"orders" 
		WHERE
			"order_id" = $1;
	`
)

func (q *Queries) InsertOrder(ctx context.Context, req entity.ReqCreateOrder, status entity.OrderStatus) (*entity.DBOrder, error) {
	i := &entity.DBOrder{}
	err := q.db.QueryRowContext(ctx,
		insertOrder,
		req.ProductID,
		req.UserID,
		status).Scan(
		&i.OrderID,
		&i.ProductID,
		&i.UserID,
		&i.WalletHistoryID,
		&i.Status,
		&i.Date,
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (q *Queries) UpdateOrder(ctx context.Context, req entity.ReqUpdateOrder) (*entity.DBOrder, error) {
	i := &entity.DBOrder{}
	err := q.db.QueryRowContext(ctx,
		updateOrder,
		req.Status,
		req.WalletHistoryID,
		req.OrderID).Scan(
		&i.OrderID,
		&i.ProductID,
		&i.UserID,
		&i.WalletHistoryID,
		&i.Status,
		&i.Date,
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (q *Queries) GetOrder(ctx context.Context, orderID int64) (*entity.DBOrder, error) {
	i := &entity.DBOrder{}
	err := q.db.QueryRowContext(ctx,
		getOrder,
		orderID).Scan(
		&i.OrderID,
		&i.ProductID,
		&i.UserID,
		&i.WalletHistoryID,
		&i.Status,
		&i.Date,
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}
