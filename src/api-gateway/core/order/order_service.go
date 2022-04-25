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
	CreateOrder(ctx context.Context, req entity.HTTPReqPostCreateOrder) (string, error)
	GetOrders(ctx context.Context, req entity.HTTPReqGetOrders) (*[]entity.DBOrder, error)
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

func (os orderService) CreateOrder(ctx context.Context, req entity.HTTPReqPostCreateOrder) (string, error) {
	grpcReq := &minipulsa.CreateOrderRequest{
		ProductId: req.ProductID,
		UserId:    req.UserID,
	}
	createOrderResp, err := os.OrderServiceClient.CreateOrder(ctx, grpcReq)
	if err != nil {
		os.log.Logger.Error("LoginRegister ERR: ", err)
		return "", err
	}

	grpcReqProduct := &minipulsa.GetProductRequest{
		ProductId: req.ProductID,
	}

	product, err := os.ProductServiceClient.GetProduct(ctx, grpcReqProduct)
	if err != nil {
		os.log.Logger.Error("LoginRegister ERR: ", err)
		return "", err
	}

	grpcReqWallet := &minipulsa.UseWalletRequest{
		UserId:  req.UserID,
		Amount:  -product.Product.Price,
		OrderId: createOrderResp.Order.OrderId,
	}

	resWallet, err := os.WalletServiceClient.UseWallet(ctx, grpcReqWallet)
	if err != nil {
		// os.log.Logger.Error("LoginRegister ERR: ", err)
		grpcReq := &minipulsa.UpdateOrderRequest{
			OrderId:         createOrderResp.Order.OrderId,
			Status:          "CANCELED",
			WalletHistoryId: 0,
		}
		_, err := os.OrderServiceClient.UpdateOrder(ctx, grpcReq)
		if err != nil {
			os.log.Logger.Error("LoginRegister ERR: ", err)
			return "", err
		}
		return "CANCELED", nil
	}

	grpcReqUpdate := &minipulsa.UpdateOrderRequest{
		OrderId:         createOrderResp.Order.OrderId,
		Status:          "SUCCESS",
		WalletHistoryId: resWallet.WalletHistoryId,
	}
	_, err = os.OrderServiceClient.UpdateOrder(ctx, grpcReqUpdate)
	if err != nil {
		os.log.Logger.Error("LoginRegister ERR: ", err)
		return "", err
	}

	return "SUCCESS", nil
}

func (os orderService) GetOrders(ctx context.Context, req entity.HTTPReqGetOrders) (*[]entity.DBOrder, error) {
	reqGrpc := &minipulsa.GetOrdersRequest{
		UserId: req.UserID,
	}

	createOrderResp, err := os.OrderServiceClient.GetOrders(ctx, reqGrpc)
	if err != nil {
		os.log.Logger.Error("LoginRegister ERR: ", err)
		return nil, err
	}
	baseRes := []entity.DBOrder{}
	for i := 0; i < len(createOrderResp.Orders); i++ {
		baseRes = append(baseRes, entity.DBOrder{
			OrderID:         createOrderResp.Orders[i].OrderId,
			ProductID:       createOrderResp.Orders[i].ProductId,
			UserID:          createOrderResp.Orders[i].UserId,
			WalletHistoryID: createOrderResp.Orders[i].WalletHistoryId,
			Status:          createOrderResp.Orders[i].Status,
			Date:            createOrderResp.Orders[i].Date,
		})
	}

	return &baseRes, nil
}
