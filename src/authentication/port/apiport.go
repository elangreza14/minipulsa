package ports

import (
	"context"

	"github.com/elangreza14/minipulsa/authentication/entity"
)

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	LoginRegister(context.Context, entity.ReqPostPutUser) (string, *int64, error)
	ValidateToken(context.Context, string) (string, *int64, error)
}
