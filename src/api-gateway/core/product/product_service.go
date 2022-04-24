package product

import (
	"context"

	"github.com/elangreza14/minipulsa/api-gateway/adapter/grpc/minipulsa"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/elangreza14/minipulsa/api-gateway/port"
	"github.com/sirupsen/logrus"
)

// ProductService is service layer that handle interaction between core and adapter
type ProductService interface {
	GetProducts(context.Context) (*entity.HTTPResGetProducts, error)
}

type productService struct {
	ProductServiceClient port.ProductRepo
	log                  *logrus.Entry
}

// NewProductService is generating new instance
func NewProductService(
	log *logrus.Entry,
	ProductServiceClient port.ProductRepo,
) ProductService {
	return &productService{
		ProductServiceClient: ProductServiceClient,
		log:                  log,
	}
}

func (ps productService) GetProducts(ctx context.Context) (*entity.HTTPResGetProducts, error) {
	res, err := ps.ProductServiceClient.GetProducts(ctx, &minipulsa.Empty{})
	if err != nil {
		ps.log.Logger.Error("LoginRegister ERR: ", err)
		return nil, err
	}
	products := []entity.Product{}
	for i := 0; i < len(res.Products); i++ {
		products = append(products, entity.Product{
			ProductId: res.Products[i].ProductId,
			Price:     res.Products[i].Price,
			Name:      res.Products[i].Name,
		})
	}

	resp := &entity.HTTPResGetProducts{
		Products: products,
	}

	return resp, nil
}
