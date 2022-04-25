#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

# Repository
mockgen -package=mock -destination=mock/mock_ProductRepository.go github.com/elangreza14/minipulsa/product/core/product ProductRepository

# Core
mockgen -package=mock -destination=mock/mock_ProductService.go github.com/elangreza14/minipulsa/product/core/product ProductService