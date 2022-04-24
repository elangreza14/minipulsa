package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/wallet/entity"
)

type (
	Querier interface {
		WalletQuerier
	}

	// this is some method that contain some repository like in Wallet Core Folder
	// it create same interface so we can create successful docking between adapter and core
	WalletQuerier interface {
		GetWalletByUserID(ctx context.Context, userID int64) (*entity.DBWallet, error)
		UpdateWalletByUserID(ctx context.Context, req entity.ReqUseWallet) error
		InsertWallet(ctx context.Context, req entity.ReqUseWallet) error
		GetWalletHistories(ctx context.Context, userID int64) (*[]entity.DBWalletHistories, error)
		GetWalletHistoryByReqUseWallet(ctx context.Context, req entity.ReqUseWallet) (*entity.DBWalletHistoriesDetail, error)
	}

	// we can also create another repository like logic here
	// we also can create transaction like with ease ini this postgres package
)

// we are casting querier that contain Wallet querier into queries
var _ Querier = (*Queries)(nil)
