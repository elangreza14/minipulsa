package order

import (
	"context"

	"github.com/elangreza14/minipulsa/order/entity"
)

type (
	OrderRepository interface {
		UpdateOrder(ctx context.Context, req entity.ReqUpdateOrder) (*entity.DBOrder, error)
		InsertOrder(ctx context.Context, req entity.ReqCreateOrder, status entity.OrderStatus) (*entity.DBOrder, error)
		GetOrder(ctx context.Context, orderID int64) (*entity.DBOrder, error)
	}
)
