package product

import (
	"context"
	"database/sql"

	"github.com/elangreza14/minipulsa/product/entity"
	"github.com/sirupsen/logrus"
)

// ProductService is service layer that handle interaction between core and adapter
type ProductService interface {
	GetProductByID(ctx context.Context, productID int64) (*entity.DBProduct, error)
	GetProducts(context.Context) (*[]entity.DBProduct, error)
}

type productService struct {
	productRepo ProductRepository
	log         *logrus.Entry
}

// NewProductService is generating new instance
func NewProductService(ProductRepo ProductRepository,
	log *logrus.Entry,
) ProductService {
	return &productService{
		productRepo: ProductRepo,
		log:         log,
	}
}

func (ps *productService) GetProductByID(ctx context.Context, productID int64) (*entity.DBProduct, error) {
	res, err := ps.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		ps.log.Error("GetProductByID ERROR: ", err)
		if err == sql.ErrNoRows {
			return nil, entity.ErrorGRPCNotFound
		}
		return nil, entity.ErrorGRPCInternalServer
	}
	return res, nil
}

func (ps *productService) GetProducts(ctx context.Context) (*[]entity.DBProduct, error) {
	res, err := ps.productRepo.GetProducts(ctx)
	if err != nil {
		ps.log.Error("GetProducts ERROR: ", err)
		if err == sql.ErrNoRows {
			return nil, entity.ErrorGRPCNotFound
		}
		return nil, entity.ErrorGRPCInternalServer
	}
	return res, nil
}
