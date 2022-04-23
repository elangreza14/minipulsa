package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/authentication/entity"
)

type (
	Querier interface {
		UserQuerier
	}

	// this is some method that contain some repository like in User Core Folder
	// it create same interface so we can create successful docking between adapter and core
	UserQuerier interface {
		InsertUser(ctx context.Context, req entity.ReqPostPutUser) (*int64, error)
		UpdateUser(ctx context.Context, req entity.ReqPostPutUser, userID int64) error
		GetUser(ctx context.Context, userID int64) (*entity.DBUser, error)
		GetUserByEmail(ctx context.Context, email string) (*entity.DBUser, error)
	}

	// we can also create another repository like logic here
	// we also can create transaction like with ease ini this postgres package
)

// we are casting querier that contain User querier into queries
var _ Querier = (*Queries)(nil)
