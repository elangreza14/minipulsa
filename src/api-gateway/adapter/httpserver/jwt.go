package httpserver

import (
	"errors"
	"fmt"
	"strings"

	"github.com/elangreza14/minipulsa/api-gateway/entity"
	"github.com/gofiber/fiber/v2"
)

type JwtMiddleware interface {
	AuthorizeCurrentToken(c *fiber.Ctx) error
}

type jwtMiddleware struct {
}

func NewJwtMiddleware() JwtMiddleware {
	return &jwtMiddleware{}
}

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
)

func (ur *jwtMiddleware) AuthorizeCurrentToken(c *fiber.Ctx) error {
	baseAuthorizationHeader := c.GetReqHeaders()

	authorizationHeader := baseAuthorizationHeader[authorizationHeaderKey]

	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		return c.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}

	fields := strings.Fields(authorizationHeader)

	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return c.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		return c.Status(500).JSON(entity.ResponseHTTP{
			Code:    500,
			Message: []string{err.Error()},
			Data:    nil,
		})
	}

	accessToken := fields[1]

	c.Locals("jwt-token", accessToken)

	return c.Next()
}
