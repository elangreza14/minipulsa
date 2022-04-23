package ports

import (
	"context"

	"github.com/elangreza14/minipulsa/product/entity"
)

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	GetProductByID(ctx context.Context, productID int64) (*entity.DBProduct, error)
	GetProducts(context.Context) (*[]entity.DBProduct, error)
}
