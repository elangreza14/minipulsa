package order

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/minipulsa/order/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/order/entity"
	"github.com/elangreza14/minipulsa/order/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_orderService_CreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("error", func(t *testing.T) {
		mockOrderRepo := mock.NewMockOrderRepository(ctrl)
		os := &orderService{
			log:       loggerserver.NewLogger(),
			orderRepo: mockOrderRepo,
		}
		base := entity.ReqCreateOrder{
			ProductID: 0,
			UserID:    0,
		}
		errBase := errors.New("err")
		mockOrderRepo.EXPECT().InsertOrder(context.Background(), base, entity.OrderStatusPending).Return(nil, errBase).AnyTimes()
		_, err := os.orderRepo.InsertOrder(context.Background(), base, entity.OrderStatusPending)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockOrderRepo := mock.NewMockOrderRepository(ctrl)
		os := &orderService{
			log:       loggerserver.NewLogger(),
			orderRepo: mockOrderRepo,
		}
		base := entity.ReqCreateOrder{
			ProductID: 0,
			UserID:    0,
		}
		resBase := &entity.DBOrder{}
		mockOrderRepo.EXPECT().InsertOrder(context.Background(), base, entity.OrderStatusPending).Return(resBase, nil).AnyTimes()
		res, err := os.orderRepo.InsertOrder(context.Background(), base, entity.OrderStatusPending)
		require.NoError(t, err)
		require.Equal(t, res, resBase)
	})
}

func Test_orderService_UpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("error GetOrder", func(t *testing.T) {
		mockRepo := mock.NewMockOrderRepository(ctrl)
		os := orderService{
			log:       loggerserver.NewLogger(),
			orderRepo: mockRepo,
		}
		base := entity.ReqUpdateOrder{
			OrderID:         1,
			Status:          "",
			WalletHistoryID: 1,
		}
		errBase := errors.New("err")
		mockRepo.EXPECT().GetOrder(context.Background(), base.OrderID).Return(nil, errBase).AnyTimes()
		_, err := os.UpdateOrder(context.Background(), base)
		require.Error(t, err)
	})

	t.Run("error UpdateOrder", func(t *testing.T) {
		mockRepo := mock.NewMockOrderRepository(ctrl)
		os := orderService{
			log:       loggerserver.NewLogger(),
			orderRepo: mockRepo,
		}
		base := entity.ReqUpdateOrder{
			OrderID:         1,
			Status:          "",
			WalletHistoryID: 1,
		}
		baseGetOrder := &entity.DBOrder{}
		errBase := errors.New("err")
		mockRepo.EXPECT().GetOrder(context.Background(), base.OrderID).Return(baseGetOrder, nil).AnyTimes()
		mockRepo.EXPECT().UpdateOrder(context.Background(), base).Return(nil, errBase).AnyTimes()
		_, err := os.UpdateOrder(context.Background(), base)
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo := mock.NewMockOrderRepository(ctrl)
		os := orderService{
			log:       loggerserver.NewLogger(),
			orderRepo: mockRepo,
		}
		base := entity.ReqUpdateOrder{
			OrderID:         1,
			Status:          "",
			WalletHistoryID: 1,
		}
		baseGetOrder := &entity.DBOrder{}
		mockRepo.EXPECT().GetOrder(context.Background(), base.OrderID).Return(baseGetOrder, nil).AnyTimes()
		mockRepo.EXPECT().UpdateOrder(context.Background(), base).Return(baseGetOrder, nil).AnyTimes()
		_, err := os.UpdateOrder(context.Background(), base)
		require.NoError(t, err)
	})
}
