package wallet

import (
	"context"

	"github.com/elangreza14/minipulsa/wallet/entity"
)

type (
	// WalletRepository handle all repository with this 4 method availabel
	// This Repository is not concern of any adapters. Can be postgres / redis/ mongo/ maria
	// This Repository also concern about handling all method availabel

	WalletRepository interface {
		GetWalletByUserID(ctx context.Context, userID int64) (*entity.DBWallet, error)
		UpdateWalletByUserID(ctx context.Context, amount int64, userID int64) error
		InsertWallet(ctx context.Context, amount int64, userID int64) error
		GetWalletHistories(ctx context.Context, userID int64) (*[]entity.DBWalletHistories, error)
	}
)
