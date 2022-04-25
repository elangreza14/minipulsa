#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

# Repository
mockgen -package=mock -destination=mock/mock_AuthenticationRepo.go github.com/elangreza14/minipulsa/api-gateway/port AuthenticationRepo
mockgen -package=mock -destination=mock/mock_ProductRepo.go github.com/elangreza14/minipulsa/api-gateway/port ProductRepo
mockgen -package=mock -destination=mock/mock_WalletRepo.go github.com/elangreza14/minipulsa/api-gateway/port WalletRepo
mockgen -package=mock -destination=mock/mock_OrderRepo.go github.com/elangreza14/minipulsa/api-gateway/port OrderRepo

# Service
mockgen -package=mock -destination=mock/mock_AuthenticationService.go github.com/elangreza14/minipulsa/api-gateway/core/authentication AuthenticationService
mockgen -package=mock -destination=mock/mock_ProductService.go github.com/elangreza14/minipulsa/api-gateway/core/product ProductService
mockgen -package=mock -destination=mock/mock_WalletService.go github.com/elangreza14/minipulsa/api-gateway/core/wallet WalletService
mockgen -package=mock -destination=mock/mock_OrderService.go github.com/elangreza14/minipulsa/api-gateway/core/order OrderService
