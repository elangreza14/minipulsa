package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/product/entity"
)

type (
	Querier interface {
		ProductQuerier
	}

	// this is some method that contain some repository like in Product Core Folder
	// it create same interface so we can create successful docking between adapter and core
	ProductQuerier interface {
		GetProductByID(ctx context.Context, productID int64) (*entity.DBProduct, error)
		GetProducts(ctx context.Context) (*[]entity.DBProduct, error)
	}

	// we can also create another repository like logic here
	// we also can create transaction like with ease ini this postgres package
)

// we are casting querier that contain Product querier into queries
var _ Querier = (*Queries)(nil)
