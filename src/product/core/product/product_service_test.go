package product

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/minipulsa/product/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/product/entity"
	"github.com/elangreza14/minipulsa/product/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_productService_GetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("success", func(t *testing.T) {
		mockProductRepo := mock.NewMockProductRepository(ctrl)

		ps := &productService{
			productRepo: mockProductRepo,
		}
		res := &entity.DBProduct{
			ProductID: 0,
			Price:     0,
		}
		var baseID int64 = 1
		mockProductRepo.EXPECT().GetProductByID(context.Background(), baseID).Return(res, nil).AnyTimes()
		resRepo, err := ps.GetProductByID(context.Background(), baseID)
		require.NoError(t, err)
		require.Equal(t, res, resRepo)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo := mock.NewMockProductRepository(ctrl)

		ps := &productService{
			productRepo: mockProductRepo,
			log:         loggerserver.NewLogger(),
		}
		var baseID int64 = 1
		var errBase error = errors.New("err")

		mockProductRepo.EXPECT().GetProductByID(context.Background(), baseID).Return(nil, errBase).AnyTimes()
		resRepo, err := ps.GetProductByID(context.Background(), baseID)
		require.Error(t, err)
		require.Empty(t, resRepo)
	})
}

func Test_productService_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("success", func(t *testing.T) {
		mockProductRepo := mock.NewMockProductRepository(ctrl)

		ps := &productService{
			productRepo: mockProductRepo,
			log:         loggerserver.NewLogger(),
		}
		res := &[]entity.DBProduct{}
		mockProductRepo.EXPECT().GetProducts(context.Background()).Return(res, nil).AnyTimes()
		resRepo, err := ps.GetProducts(context.Background())
		require.NoError(t, err)
		require.Equal(t, res, resRepo)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo := mock.NewMockProductRepository(ctrl)
		ps := &productService{
			productRepo: mockProductRepo,
			log:         loggerserver.NewLogger(),
		}
		var errBase error = errors.New("err")
		mockProductRepo.EXPECT().GetProducts(context.Background()).Return(nil, errBase).AnyTimes()
		res, err := ps.GetProducts(context.Background())
		require.Error(t, err)
		require.Empty(t, res)
	})

}
