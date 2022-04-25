package order

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/elangreza14/minipulsa/api-gateway/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_orderService_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("err CreateOrder", func(t *testing.T) {
		orderRepoMock := mock.NewMockOrderRepo(ctrl)
		walletRepoMock := mock.NewMockWalletRepo(ctrl)
		productRepoMock := mock.NewMockProductRepo(ctrl)
		ps := orderService{
			OrderServiceClient:   orderRepoMock,
			log:                  loggerserver.NewLogger(),
			WalletServiceClient:  walletRepoMock,
			ProductServiceClient: productRepoMock,
		}

		req := entity.HTTPReqPostCreateOrder{
			ProductID: 1,
			UserID:    1,
		}
		grpcReq := &minipulsa.CreateOrderRequest{
			ProductId: req.ProductID,
			UserId:    req.UserID,
		}
		errBase := errors.New("err")
		orderRepoMock.EXPECT().CreateOrder(context.Background(), grpcReq).Return(nil, errBase).AnyTimes()
		_, err := ps.CreateOrder(context.Background(), req)
		require.Error(t, err)
	})

	t.Run("err GetProduct", func(t *testing.T) {
		orderRepoMock := mock.NewMockOrderRepo(ctrl)
		walletRepoMock := mock.NewMockWalletRepo(ctrl)
		productRepoMock := mock.NewMockProductRepo(ctrl)
		ps := orderService{
			OrderServiceClient:   orderRepoMock,
			log:                  loggerserver.NewLogger(),
			WalletServiceClient:  walletRepoMock,
			ProductServiceClient: productRepoMock,
		}

		req := entity.HTTPReqPostCreateOrder{
			ProductID: 1,
			UserID:    1,
		}
		grpcReq := &minipulsa.CreateOrderRequest{
			ProductId: req.ProductID,
			UserId:    req.UserID,
		}
		errBase := errors.New("err")
		createOrderRes := &minipulsa.CreateORUpdateOrderResponse{
			Order: &minipulsa.Order{},
		}

		orderRepoMock.EXPECT().CreateOrder(context.Background(), grpcReq).Return(createOrderRes, nil).AnyTimes()
		grpcReqProduct := &minipulsa.GetProductRequest{
			ProductId: req.ProductID,
		}

		productRepoMock.EXPECT().GetProduct(context.Background(), grpcReqProduct).Return(nil, errBase).AnyTimes()
		_, err := ps.CreateOrder(context.Background(), req)
		require.Error(t, err)
	})

	t.Run("err UseWallet 1", func(t *testing.T) {
		orderRepoMock := mock.NewMockOrderRepo(ctrl)
		walletRepoMock := mock.NewMockWalletRepo(ctrl)
		productRepoMock := mock.NewMockProductRepo(ctrl)
		ps := orderService{
			OrderServiceClient:   orderRepoMock,
			log:                  loggerserver.NewLogger(),
			WalletServiceClient:  walletRepoMock,
			ProductServiceClient: productRepoMock,
		}

		req := entity.HTTPReqPostCreateOrder{
			ProductID: 1,
			UserID:    1,
		}
		grpcReq := &minipulsa.CreateOrderRequest{
			ProductId: req.ProductID,
			UserId:    req.UserID,
		}
		errBase := errors.New("err")
		createOrderRes := &minipulsa.CreateORUpdateOrderResponse{
			Order: &minipulsa.Order{},
		}

		orderRepoMock.EXPECT().CreateOrder(context.Background(), grpcReq).Return(createOrderRes, nil).AnyTimes()
		grpcReqProduct := &minipulsa.GetProductRequest{
			ProductId: req.ProductID,
		}
		getProductRes := &minipulsa.GetProductResponse{
			Product: &minipulsa.Product{
				ProductId: 1,
				Price:     1,
				Name:      "1",
			},
		}
		productRepoMock.EXPECT().GetProduct(context.Background(), grpcReqProduct).Return(getProductRes, nil).AnyTimes()

		grpcReqWallet := &minipulsa.UseWalletRequest{
			UserId:  req.UserID,
			Amount:  -getProductRes.Product.Price,
			OrderId: createOrderRes.Order.OrderId,
		}
		walletRepoMock.EXPECT().UseWallet(context.Background(), grpcReqWallet).Return(nil, errBase).AnyTimes()

		grpcReqUpdate := &minipulsa.UpdateOrderRequest{
			OrderId: createOrderRes.Order.OrderId,
			Status:  "CANCELED",
		}
		orderRepoMock.EXPECT().UpdateOrder(context.Background(), grpcReqUpdate).Return(nil, errBase).AnyTimes()

		_, err := ps.CreateOrder(context.Background(), req)
		require.Error(t, err)
	})

	t.Run("err UseWallet 2", func(t *testing.T) {
		orderRepoMock := mock.NewMockOrderRepo(ctrl)
		walletRepoMock := mock.NewMockWalletRepo(ctrl)
		productRepoMock := mock.NewMockProductRepo(ctrl)
		ps := orderService{
			OrderServiceClient:   orderRepoMock,
			log:                  loggerserver.NewLogger(),
			WalletServiceClient:  walletRepoMock,
			ProductServiceClient: productRepoMock,
		}

		req := entity.HTTPReqPostCreateOrder{
			ProductID: 1,
			UserID:    1,
		}
		grpcReq := &minipulsa.CreateOrderRequest{
			ProductId: req.ProductID,
			UserId:    req.UserID,
		}
		errBase := errors.New("err")
		createOrderRes := &minipulsa.CreateORUpdateOrderResponse{
			Order: &minipulsa.Order{},
		}

		orderRepoMock.EXPECT().CreateOrder(context.Background(), grpcReq).Return(createOrderRes, nil).AnyTimes()
		grpcReqProduct := &minipulsa.GetProductRequest{
			ProductId: req.ProductID,
		}
		getProductRes := &minipulsa.GetProductResponse{
			Product: &minipulsa.Product{
				ProductId: 1,
				Price:     1,
				Name:      "1",
			},
		}
		productRepoMock.EXPECT().GetProduct(context.Background(), grpcReqProduct).Return(getProductRes, nil).AnyTimes()

		grpcReqWallet := &minipulsa.UseWalletRequest{
			UserId:  req.UserID,
			Amount:  -getProductRes.Product.Price,
			OrderId: createOrderRes.Order.OrderId,
		}

		grpcResWallet := &minipulsa.UseWalletResponse{
			WalletHistoryId: 1,
		}
		walletRepoMock.EXPECT().UseWallet(context.Background(), grpcReqWallet).Return(grpcResWallet, nil).AnyTimes()

		grpcReqUpdate := &minipulsa.UpdateOrderRequest{
			OrderId:         createOrderRes.Order.OrderId,
			Status:          "SUCCESS",
			WalletHistoryId: grpcResWallet.WalletHistoryId,
		}
		orderRepoMock.EXPECT().UpdateOrder(context.Background(), grpcReqUpdate).Return(nil, errBase).AnyTimes()

		_, err := ps.CreateOrder(context.Background(), req)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		orderRepoMock := mock.NewMockOrderRepo(ctrl)
		walletRepoMock := mock.NewMockWalletRepo(ctrl)
		productRepoMock := mock.NewMockProductRepo(ctrl)
		ps := orderService{
			OrderServiceClient:   orderRepoMock,
			log:                  loggerserver.NewLogger(),
			WalletServiceClient:  walletRepoMock,
			ProductServiceClient: productRepoMock,
		}

		req := entity.HTTPReqPostCreateOrder{
			ProductID: 1,
			UserID:    1,
		}
		grpcReq := &minipulsa.CreateOrderRequest{
			ProductId: req.ProductID,
			UserId:    req.UserID,
		}
		// errBase := errors.New("err")
		createOrderRes := &minipulsa.CreateORUpdateOrderResponse{
			Order: &minipulsa.Order{},
		}

		orderRepoMock.EXPECT().CreateOrder(context.Background(), grpcReq).Return(createOrderRes, nil).AnyTimes()
		grpcReqProduct := &minipulsa.GetProductRequest{
			ProductId: req.ProductID,
		}
		getProductRes := &minipulsa.GetProductResponse{
			Product: &minipulsa.Product{
				ProductId: 1,
				Price:     1,
				Name:      "1",
			},
		}
		productRepoMock.EXPECT().GetProduct(context.Background(), grpcReqProduct).Return(getProductRes, nil).AnyTimes()

		grpcReqWallet := &minipulsa.UseWalletRequest{
			UserId:  req.UserID,
			Amount:  -getProductRes.Product.Price,
			OrderId: createOrderRes.Order.OrderId,
		}

		grpcResWallet := &minipulsa.UseWalletResponse{
			WalletHistoryId: 1,
		}
		walletRepoMock.EXPECT().UseWallet(context.Background(), grpcReqWallet).Return(grpcResWallet, nil).AnyTimes()

		grpcReqUpdate := &minipulsa.UpdateOrderRequest{
			OrderId:         createOrderRes.Order.OrderId,
			Status:          "SUCCESS",
			WalletHistoryId: grpcResWallet.WalletHistoryId,
		}
		grpcResUpdate := &minipulsa.CreateORUpdateOrderResponse{}
		orderRepoMock.EXPECT().UpdateOrder(context.Background(), grpcReqUpdate).Return(grpcResUpdate, nil).AnyTimes()

		_, err := ps.CreateOrder(context.Background(), req)
		require.NoError(t, err)
	})

}
