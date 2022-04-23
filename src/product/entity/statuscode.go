package entity

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrPassWrong  error = errors.New("password not valid")
	ErrTokenWrong error = errors.New("token not valid")

	ErrorGRPCInternalServer error = status.Errorf(codes.Internal, "internal server error")
	ErrorGRPCPasswordWrong  error = status.Errorf(codes.Unauthenticated, ErrPassWrong.Error())
	ErrorGRPCTokenWrong     error = status.Errorf(codes.Unauthenticated, ErrTokenWrong.Error())
	ErrorGRPCUnauthorized   error = status.Errorf(codes.Unauthenticated, "request unauthenticated")
	ErrorGRPCBadRequest     error = status.Errorf(codes.InvalidArgument, "bad request")
	ErrorGRPCNotFound       error = status.Errorf(codes.NotFound, "response no found")
)
