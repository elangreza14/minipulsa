package order

import (
	"context"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/elangreza14/minipulsa/api-gateway/port"
	"github.com/sirupsen/logrus"
)

// OrderService is service layer that handle interaction between core and adapter
type OrderService interface {
	CreateOrder(ctx context.Context, req entity.HTTPReqPostCreateOrder) error
}

type orderService struct {
	OrderServiceClient   port.OrderRepo
	log                  *logrus.Entry
	WalletServiceClient  port.WalletRepo
	ProductServiceClient port.ProductRepo
}

func NewOrderService(
	log *logrus.Entry,
	OrderServiceClient port.OrderRepo,
	WalletServiceClient port.WalletRepo,
	ProductServiceClient port.ProductRepo,

) OrderService {
	return &orderService{
		OrderServiceClient:   OrderServiceClient,
		log:                  log,
		WalletServiceClient:  WalletServiceClient,
		ProductServiceClient: ProductServiceClient,
	}
}

func (ws orderService) CreateOrder(ctx context.Context, req entity.HTTPReqPostCreateOrder) error {
	grpcReq := &minipulsa.CreateOrderRequest{
		ProductId: req.ProductID,
		UserId:    req.UserID,
	}
	createOrderResp, err := ws.OrderServiceClient.CreateOrder(ctx, grpcReq)
	if err != nil {
		ws.log.Logger.Error("LoginRegister ERR: ", err)
		return err
	}

	grpcReqProduct := &minipulsa.GetProductRequest{
		ProductId: req.ProductID,
	}

	product, err := ws.ProductServiceClient.GetProduct(ctx, grpcReqProduct)
	if err != nil {
		ws.log.Logger.Error("LoginRegister ERR: ", err)
		return err
	}

	grpcReqWallet := &minipulsa.UseWalletRequest{
		UserId:  req.UserID,
		Amount:  -product.Product.Price,
		OrderId: createOrderResp.Order.OrderId,
	}

	resWallet, err := ws.WalletServiceClient.UseWallet(ctx, grpcReqWallet)
	if err != nil {
		// ws.log.Logger.Error("LoginRegister ERR: ", err)
		grpcReq := &minipulsa.UpdateOrderRequest{
			OrderId:         createOrderResp.Order.OrderId,
			Status:          "CANCELED",
			WalletHistoryId: 0,
		}
		_, err := ws.OrderServiceClient.UpdateOrder(ctx, grpcReq)
		if err != nil {
			ws.log.Logger.Error("LoginRegister ERR: ", err)
			return err
		}
		return nil
	}

	grpcReqUpdate := &minipulsa.UpdateOrderRequest{
		OrderId:         createOrderResp.Order.OrderId,
		Status:          "SUCCESS",
		WalletHistoryId: resWallet.WalletHistoryId,
	}
	_, err = ws.OrderServiceClient.UpdateOrder(ctx, grpcReqUpdate)
	if err != nil {
		ws.log.Logger.Error("LoginRegister ERR: ", err)
		return err
	}

	return nil
}
