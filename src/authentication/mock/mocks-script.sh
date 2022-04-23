#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen

# Repository
mockgen -package=mock -destination=mock/mock_UserRepository.go github.com/elangreza14/minipulsa/authentication/core/user UserRepository

# Core
mockgen -package=mock -destination=mock/mock_UserService.go github.com/elangreza14/minipulsa/authentication/core/user UserService