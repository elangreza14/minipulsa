package httpserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
)

type AuthenticationController struct {
	authenticationService authentication.AuthenticationService
	log                   *logrus.Entry
}

func NewPublicAuthenticationController(
	router fiber.Router,
	authenticationService authentication.AuthenticationService,
	log *logrus.Entry,
) {
	handler := &AuthenticationController{
		authenticationService: authenticationService,
		log:                   log,
	}
	router.Post("/login-register", handler.CreateAuthentication)
}

func (wc *AuthenticationController) CreateAuthentication(ctx *fiber.Ctx) error {
	dataBody := new(entity.HTTPReqPostPutUser)
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

	token, _, err := wc.authenticationService.LoginRegister(ctx.Context(), *dataBody)

	if err != nil {
		wc.log.Error("authenticationController.CreateAuthentication ERROR: ", err)
		return ctx.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}

	res := entity.HTTPResPostPutUser{
		Token: token,
	}

	return ctx.Status(200).JSON(entity.ResponseHTTP{
		Code:    200,
		Message: []string{"ok"},
		Data:    res,
	})
}
