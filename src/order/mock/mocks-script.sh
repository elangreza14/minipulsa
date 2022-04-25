#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

# Repository
mockgen -package=mock -destination=mock/mock_OrderRepository.go github.com/elangreza14/minipulsa/order/core/order OrderRepository

# Core
mockgen -package=mock -destination=mock/mock_OrderService.go github.com/elangreza14/minipulsa/order/core/order OrderService