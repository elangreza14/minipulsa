#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

# Repository
mockgen -package=mock -destination=mock/mock_WalletRepository.go github.com/elangreza14/minipulsa/wallet/core/wallet WalletRepository

# Core
mockgen -package=mock -destination=mock/mock_WalletService.go github.com/elangreza14/minipulsa/wallet/core/wallet WalletService