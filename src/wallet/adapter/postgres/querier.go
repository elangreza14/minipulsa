package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/wallet/entity"
)

type (
	Querier interface {
		ProductQuerier
	}

	// this is some method that contain some repository like in Product Core Folder
	// it create same interface so we can create successful docking between adapter and core
	ProductQuerier interface {
		GetWalletByUserID(ctx context.Context, userID int64) (*entity.DBWallet, error)
		UpdateWalletByUserID(ctx context.Context, amount int64, userID int64) error
		InsertWallet(ctx context.Context, amount int64, userID int64) error
		GetWalletHistories(ctx context.Context, userID int64) error
	}

	// we can also create another repository like logic here
	// we also can create transaction like with ease ini this postgres package
)

// we are casting querier that contain Product querier into queries
var _ Querier = (*Queries)(nil)
