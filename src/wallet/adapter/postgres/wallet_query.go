package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/wallet/entity"
)

const (
	getWalletByID = `
		SELECT
			"wallet_id",
			"user_id",
			"order_id",
			"amount",
			"date"
		FROM
			"wallets" 
		WHERE
			user_id = $1;
	`
	updateWalletByID = `
		UPDATE
			"wallets" 
		SET
			"amount"="amount"+$1,
			"last_amount"=$2,
			"order_id"=$3,
			"date"=NOW()
		WHERE
			user_id = $4;
	`
	insertWallet = `
		INSERT INTO 
			"wallets" 
		("user_id","amount","last_amount","order_id") 
		VALUES ($1,$2,$3,$4);
	`

	getWalletHistories = `
		SELECT
			wh."last_amount",
			wh."order_id",
			wh."date"
		FROM
			"wallets" AS w 
		LEFT JOIN
			"wallet_histories" as wh
		ON  
			w.wallet_id = wh.wallet_id
		WHERE
			user_id = $1 
		ORDER BY "date" DESC;
	`

	getWalletHistoryByReqUseWallet = `
		SELECT
			wh."wallet_history_id",
			wh."last_amount",
			wh."order_id",
			wh."date"
		FROM
			"wallets" AS w 
		LEFT JOIN
			"wallet_histories" as wh
		ON  
			w.wallet_id = wh.wallet_id
		WHERE
			w.user_id = $1 
		AND
			wh.last_amount = $2
		AND
			w.order_id = $3 
		ORDER BY "date" DESC LIMIT 1;
	`
)

func (q *Queries) GetWalletByUserID(ctx context.Context, userID int64) (*entity.DBWallet, error) {
	i := &entity.DBWallet{}
	err := q.db.QueryRowContext(ctx, getWalletByID, userID).Scan(
		&i.WalletID,
		&i.UserID,
		&i.OrderID,
		&i.Amount,
		&i.Date,
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// ws.walletRepo.UpdateWalletByUserID(ctx, req.Amount, req.UserID)
func (q *Queries) UpdateWalletByUserID(ctx context.Context, req entity.ReqUseWallet) error {
	_, err := q.db.ExecContext(ctx, updateWalletByID, req.Amount, req.Amount, req.OrderID, req.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) InsertWallet(ctx context.Context, req entity.ReqUseWallet) error {
	_, err := q.db.ExecContext(ctx, insertWallet, req.UserID, req.Amount, req.Amount, req.OrderID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) GetWalletHistories(ctx context.Context, userID int64) (*[]entity.DBWalletHistories, error) {
	rows, err := q.db.QueryContext(ctx, getWalletHistories, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []entity.DBWalletHistories{}
	for rows.Next() {
		var i entity.DBWalletHistories
		if err := rows.Scan(
			&i.LastAmount,
			&i.OrderID,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return &items, nil
}

func (q *Queries) GetWalletHistoryByReqUseWallet(ctx context.Context, req entity.ReqUseWallet) (*entity.DBWalletHistoriesDetail, error) {
	i := &entity.DBWalletHistoriesDetail{}
	err := q.db.QueryRowContext(ctx,
		getWalletHistoryByReqUseWallet,
		req.UserID,
		req.Amount,
		req.OrderID).Scan(
		&i.WalletHistoryID,
		&i.LastAmount,
		&i.OrderID,
		&i.Date,
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}
