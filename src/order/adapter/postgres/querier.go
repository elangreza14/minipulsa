package postgres

import (
	"context"

	"github.com/elangreza14/minipulsa/order/entity"
)

type (
	Querier interface {
		OrderQuerier
	}

	// this is some method that contain some repository like in Order Core Folder
	// it create same interface so we can create successful docking between adapter and core
	OrderQuerier interface {
		UpdateOrder(ctx context.Context, req entity.ReqUpdateOrder) (*entity.DBOrder, error)
		InsertOrder(ctx context.Context, req entity.ReqCreateOrder, status entity.OrderStatus) (*entity.DBOrder, error)
		GetOrder(ctx context.Context, orderID int64) (*entity.DBOrder, error)
	}

	// we can also create another repository like logic here
	// we also can create transaction like with ease ini this postgres package
)

// we are casting querier that contain Order querier into queries
var _ Querier = (*Queries)(nil)
