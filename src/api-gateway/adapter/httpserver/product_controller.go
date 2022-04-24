package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/elangreza14/minipulsa/api-gateway/core/product"
	"github.com/elangreza14/minipulsa/api-gateway/entity"
)

type ProductController struct {
	productService product.ProductService
	// productService product.ProductService
	log *logrus.Entry
}

func NewPublicProductController(
	router fiber.Router,
	productService product.ProductService,
	log *logrus.Entry,
) {
	handler := &ProductController{
		productService: productService,
		log:            log,
	}
	router.Get("/list", handler.GetProducts)
}

func (wc *ProductController) GetProducts(ctx *fiber.Ctx) error {
	res, err := wc.productService.GetProducts(ctx.Context())

	if err != nil {
		wc.log.Error("productController.GetProducts ERROR: ", err)
		return ctx.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(entity.ResponseHTTP{
		Code:    200,
		Message: []string{"ok"},
		Data:    res,
	})
}
