package core

import (
	"github.com/elangreza14/minipulsa/api-gateway/core/authentication"
	"github.com/sirupsen/logrus"
)

// BaseApp struct handle all core logic with folder inside core
// if want to adding logic just add folder logic, repository and some test cases
type BaseApp struct {
	Log *logrus.Entry
	As  authentication.AuthenticationService
}