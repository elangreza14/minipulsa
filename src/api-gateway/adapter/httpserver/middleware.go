package httpserver

import (
	"errors"
)

const (
	LocalsWeightID string = "Weight-id"
)

var (
	ErrorCannotLowerThanZero = errors.New("id must be positive")
	ErrorNoID                = errors.New("id is required")
)

// func (wc *WeightController) CheckIDWeightMiddleWare(ctx *fiber.Ctx) error {
// 	idParam, err := ctx.ParamsInt("id")
// 	if err != nil {
// 		wc.log.Error("error CheckIDWeightMiddleWare log: ", err)
// 		return ctx.Status(400).JSON(entity.ResponseHTTP{
// 			Code:    400,
// 			Message: []string{ErrorNoID.Error()},
// 			Data:    nil,
// 		})
// 	}

// 	if idParam < 1 {
// 		wc.log.Error("error CheckIDWeightMiddleWare log: ", err)
// 		return ctx.Status(400).JSON(entity.ResponseHTTP{
// 			Code:    400,
// 			Message: []string{ErrorCannotLowerThanZero.Error()},
// 			Data:    nil,
// 		})
// 	}

// 	ctx.Locals(LocalsWeightID, &idParam)

// 	return ctx.Next()
// }
