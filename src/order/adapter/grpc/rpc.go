package grpc

import (
	"context"

	minipulsa "github.com/elangreza14/minipulsa/order/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/order/entity"
)

func (a adapter) CreateOrder(ctx context.Context, req *minipulsa.CreateOrderRequest) (*minipulsa.CreateORUpdateOrderResponse, error) {
	reqApi := &entity.ReqCreateOrder{
		ProductID: req.ProductId,
		UserID:    req.UserId,
	}
	res, err := a.api.CreateOrder(ctx, *reqApi)
	if err != nil {
		return nil, err
	}

	return &minipulsa.CreateORUpdateOrderResponse{
		Order: &minipulsa.Order{
			OrderId:         res.OrderID,
			ProductId:       res.ProductID,
			UserId:          res.UserID,
			WalletHistoryId: res.WalletHistoryID.Int64,
			Status:          res.Status,
			Date:            res.Date.Format(entity.YMDT),
		},
	}, nil

}
func (a adapter) UpdateOrder(ctx context.Context, req *minipulsa.UpdateOrderRequest) (*minipulsa.CreateORUpdateOrderResponse, error) {
	reqApi := &entity.ReqUpdateOrder{
		OrderID:         req.OrderId,
		Status:          req.Status,
		WalletHistoryID: req.WalletHistoryId,
	}
	res, err := a.api.UpdateOrder(ctx, *reqApi)
	if err != nil {
		return nil, err
	}

	return &minipulsa.CreateORUpdateOrderResponse{
		Order: &minipulsa.Order{
			OrderId:         res.OrderID,
			ProductId:       res.ProductID,
			UserId:          res.UserID,
			WalletHistoryId: res.WalletHistoryID.Int64,
			Status:          res.Status,
			Date:            res.Date.Format(entity.YMDT),
		},
	}, nil
}
