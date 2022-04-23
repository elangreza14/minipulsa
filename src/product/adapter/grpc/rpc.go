package grpc

import (
	"context"

	minipulsa "github.com/elangreza14/minipulsa/product/adapter/grpc/minipulsa"
)

func (a adapter) GetProducts(ctx context.Context, req *minipulsa.Empty) (*minipulsa.GetProductsResponse, error) {
	resProducts, err := a.api.GetProducts(ctx)

	if err != nil {
		return nil, err
	}

	baseRes := []*minipulsa.Product{}
	rawResProducts := *resProducts
	for i := 0; i < len(rawResProducts); i++ {
		input := &minipulsa.Product{
			ProductId: rawResProducts[i].ProductID,
			Price:     rawResProducts[i].Price,
			Name:      rawResProducts[i].Name,
		}
		baseRes = append(baseRes, input)
	}

	finalRes := &minipulsa.GetProductsResponse{
		Products: baseRes,
	}

	return finalRes, nil
}

func (a adapter) GetProduct(ctx context.Context, req *minipulsa.GetProductRequest) (*minipulsa.GetProductResponse, error) {

	resProduct, err := a.api.GetProductByID(ctx, req.ProductId)

	if err != nil {
		return nil, err
	}

	finalRes := &minipulsa.GetProductResponse{
		Product: &minipulsa.Product{
			ProductId: resProduct.ProductID,
			Price:     resProduct.Price,
			Name:      resProduct.Name,
		},
	}

	return finalRes, nil
}
