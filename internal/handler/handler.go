package handler

import (
	"mvp/utils/log"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserHandler,
)

func NewUserHandler(userRepo UserRepo, log log.Logger) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
		log:      log,
	}
}
