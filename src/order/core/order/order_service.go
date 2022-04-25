package order

import (
	"context"
	"database/sql"

	"github.com/elangreza14/minipulsa/order/entity"
	"github.com/sirupsen/logrus"
)

type (
	OrderService interface {
		CreateOrder(ctx context.Context, req entity.ReqCreateOrder) (*entity.DBOrder, error)
		UpdateOrder(ctx context.Context, req entity.ReqUpdateOrder) (*entity.DBOrder, error)
		GetOrders(ctx context.Context, userID int64) (*[]entity.DBOrder, error)
	}

	orderService struct {
		log       *logrus.Entry
		orderRepo OrderRepository
	}
)

func NewOrderService(
	orderRepo OrderRepository,
	log *logrus.Entry,
) OrderService {
	return &orderService{
		log:       log,
		orderRepo: orderRepo,
	}
}

func (os *orderService) CreateOrder(ctx context.Context, req entity.ReqCreateOrder) (*entity.DBOrder, error) {
	order, err := os.orderRepo.InsertOrder(ctx, req, entity.OrderStatusPending)
	if err != nil {
		os.log.Logger.Error("CreateOrder ERR: ", err)
		return nil, entity.ErrorGRPCInternalServer
	}

	return order, nil
}

func (os *orderService) UpdateOrder(ctx context.Context, req entity.ReqUpdateOrder) (*entity.DBOrder, error) {
	_, err := os.orderRepo.GetOrder(ctx, req.OrderID)
	if err != nil {
		os.log.Logger.Error("CreateOrder ERR: ", err)
		if err == sql.ErrNoRows {
			return nil, entity.ErrorGRPCNotFound
		}
		return nil, entity.ErrorGRPCInternalServer
	}

	order, err := os.orderRepo.UpdateOrder(ctx, req)
	if err != nil {
		os.log.Logger.Error("CreateOrder ERR: ", err)
		return nil, entity.ErrorGRPCInternalServer
	}

	return order, nil
}

func (os *orderService) GetOrders(ctx context.Context, userID int64) (*[]entity.DBOrder, error) {
	orders, err := os.orderRepo.GetOrders(ctx, userID)
	if err != nil {
		os.log.Logger.Error("CreateOrder ERR: ", err)
		if err == sql.ErrNoRows {
			return nil, entity.ErrorGRPCNotFound
		}
		return nil, entity.ErrorGRPCInternalServer
	}

	return orders, nil
}
