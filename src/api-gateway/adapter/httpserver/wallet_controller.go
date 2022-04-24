package httpserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"github.com/elangreza14/minipulsa/api-gateway/core/wallet"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
)

type WalletController struct {
	walletService         wallet.WalletService
	AuthenticationService authentication.AuthenticationService
	// walletService wallet.WalletService
	log *logrus.Entry
	jwt JwtMiddleware
}

func NewPublicWalletController(
	router fiber.Router,
	walletService wallet.WalletService,
	AuthenticationService authentication.AuthenticationService,
	log *logrus.Entry,
	jwt JwtMiddleware,
) {
	handler := &WalletController{
		walletService:         walletService,
		AuthenticationService: AuthenticationService,
		log:                   log,
		jwt:                   jwt,
	}
	router.Post("/use", jwt.AuthorizeCurrentToken, handler.UseWallet)
}

func (wc *WalletController) UseWallet(ctx *fiber.Ctx) error {
	baseUserLocal := ctx.Locals("jwt-token").(string)

	dataBody := new(entity.HTTPReqPostUseWallet)
	if err := ctx.BodyParser(dataBody); err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(*dataBody); err != nil {
		wc.log.Error("authenticationController.CreateAuthentication ERROR: ", err)
		return ctx.Status(400).JSON(entity.ResponseHTTP{
			Code:    400,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}

	res, err := wc.AuthenticationService.ValidateToken(ctx.Context(), baseUserLocal)

	if err != nil {
		wc.log.Error("walletController.UseWallet ERROR: ", err)
		return ctx.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}
	reqUseWallet := &entity.HTTPReqPostUseWallet{
		Amount: dataBody.Amount,
	}
	err = wc.walletService.UseWallet(ctx.Context(), *reqUseWallet, *res)

	if err != nil {
		wc.log.Error("walletController.UseWallet ERROR: ", err)
		return ctx.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(entity.ResponseHTTP{
		Code:    200,
		Message: []string{"ok"},
		Data:    nil,
	})
}

// wc.log.Info("walletController.UseWallet INFO: ", baseUserLocal)
