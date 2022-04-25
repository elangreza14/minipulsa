package product

import (
	"context"
	"errors"
	"testing"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/adapter/loggerserver"
	"github.com/elangreza14/minipulsa/api-gateway/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_productService_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("err GetProducts", func(t *testing.T) {
		productRepoMock := mock.NewMockProductRepo(ctrl)
		ps := productService{
			ProductServiceClient: productRepoMock,
			log:                  loggerserver.NewLogger(),
		}
		errBase := errors.New("err")
		productRepoMock.EXPECT().GetProducts(context.Background(), &minipulsa.Empty{}).Return(nil, errBase).AnyTimes()

		_, err := ps.GetProducts(context.Background())
		require.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		productRepoMock := mock.NewMockProductRepo(ctrl)
		ps := productService{
			ProductServiceClient: productRepoMock,
			log:                  loggerserver.NewLogger(),
		}
		res := &minipulsa.GetProductsResponse{}
		productRepoMock.EXPECT().GetProducts(context.Background(), &minipulsa.Empty{}).Return(res, nil).AnyTimes()

		_, err := ps.GetProducts(context.Background())
		require.NoError(t, err)
	})
}
