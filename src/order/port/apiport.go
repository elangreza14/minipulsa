package port

import (
	"context"

	"github.com/elangreza14/minipulsa/order/entity"
)

type (
	ApiPort interface {
		CreateOrder(ctx context.Context, req entity.ReqCreateOrder) (*entity.DBOrder, error)
		UpdateOrder(ctx context.Context, req entity.ReqUpdateOrder) (*entity.DBOrder, error)
		GetOrders(ctx context.Context, userID int64) (*[]entity.DBOrder, error)
	}
)
