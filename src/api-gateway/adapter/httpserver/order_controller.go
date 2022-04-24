package httpserver

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"github.com/elangreza14/minipulsa/api-gateway/core/order"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
)

type OrderController struct {
	orderService          order.OrderService
	AuthenticationService authentication.AuthenticationService
	// orderService order.OrderService
	log *logrus.Entry
	jwt JwtMiddleware
}

func NewPublicOrderController(
	router fiber.Router,
	orderService order.OrderService,
	AuthenticationService authentication.AuthenticationService,
	log *logrus.Entry,
	jwt JwtMiddleware,
) {
	handler := &OrderController{
		orderService:          orderService,
		AuthenticationService: AuthenticationService,
		log:                   log,
		jwt:                   jwt,
	}
	router.Post("/process", jwt.AuthorizeCurrentToken, handler.ProcessOrder)
}

func (wc *OrderController) ProcessOrder(ctx *fiber.Ctx) error {
	baseUserLocal := ctx.Locals("jwt-token").(string)

	dataBody := new(entity.HTTPReqCreateOrder)
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
		wc.log.Error("orderController.ProcessOrder ERROR: ", err)
		return ctx.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}
	ProcessOrder := &entity.HTTPReqPostCreateOrder{
		ProductID: dataBody.ProductID,
		UserID:    *res,
	}
	err = wc.orderService.CreateOrder(ctx.Context(), *ProcessOrder)

	if err != nil {
		wc.log.Error("orderController.ProcessOrder ERROR: ", err)
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

// wc.log.Info("orderController.ProcessOrder INFO: ", baseUserLocal)
