package core

import (
	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"github.com/elangreza14/minipulsa/api-gateway/core/order"
	"github.com/elangreza14/minipulsa/api-gateway/core/product"
	"github.com/elangreza14/minipulsa/api-gateway/core/wallet"
	"github.com/sirupsen/logrus"
)

// BaseApp struct handle all core logic with folder inside core
// if want to adding logic just add folder logic, repository and some test cases
type BaseApp struct {
	Log *logrus.Entry
	As  authentication.AuthenticationService
	Ps  product.ProductService
	Ws  wallet.WalletService
	Os  order.OrderService
}

func NewBaseApp(
	Log *logrus.Entry,
	As authentication.AuthenticationService,
	Ps product.ProductService,
	Ws wallet.WalletService,
	Os order.OrderService,

) BaseApp {
	return BaseApp{
		Log: Log,
		As:  As,
		Ps:  Ps,
		Ws:  Ws,
		Os:  Os,
	}
}
