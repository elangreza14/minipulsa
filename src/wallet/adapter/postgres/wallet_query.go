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
			"date"=NOW()
		WHERE
			user_id = $3;
	`
	insertWallet = `
		INSERT INTO 
			"wallets" 
		("user_id","amount","last_amount") 
		VALUES ($1,$2,$3);
	`

	getWalletHistories = `
		SELECT
			wh."amount",
			wh."date"
		FROM
			"wallets" AS w 
		LEFT JOIN
			"wallet_histories" as wh
		ON  
			w.wallet_id = wh.wallet_id
		WHERE
			user_id = $1;
	`
)

func (q *Queries) GetWalletByUserID(ctx context.Context, userID int64) (*entity.DBWallet, error) {
	i := &entity.DBWallet{}
	err := q.db.QueryRowContext(ctx, getWalletByID, userID).Scan(
		&i.WalletID,
		&i.UserID,
		&i.Amount,
		&i.Date,
	)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// ws.walletRepo.UpdateWalletByUserID(ctx, req.Amount, req.UserID)
func (q *Queries) UpdateWalletByUserID(ctx context.Context, amount int64, userID int64) error {
	_, err := q.db.ExecContext(ctx, updateWalletByID, amount, amount, userID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) InsertWallet(ctx context.Context, amount int64, userID int64) error {
	_, err := q.db.ExecContext(ctx, insertWallet, userID, amount, amount)
	if err != nil {
		return err
	}

	return nil
}

func (q *Queries) GetWalletHistories(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, getWalletHistories, userID)
	if err != nil {
		return err
	}

	return nil
}
