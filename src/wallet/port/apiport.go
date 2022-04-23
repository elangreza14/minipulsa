package ports

import (
	"context"

	"github.com/elangreza14/minipulsa/wallet/entity"
)

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	UseWallet(ctx context.Context, req entity.ReqUseWallet) error
	GetWalletDetail(ctx context.Context, userId int64) (*entity.GetDetailWallet, error)
}
