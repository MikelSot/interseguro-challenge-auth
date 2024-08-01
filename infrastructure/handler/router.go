package handler

import (
	"github.com/MikelSot/interseguro-challenge-auth/infrastructure/handler/login"
	"github.com/MikelSot/interseguro-challenge-auth/infrastructure/handler/register"
	"github.com/MikelSot/interseguro-challenge-auth/model"
)

func InitRoutes(spec model.RouterSpecification) {
	// L
	login.NewRouter(spec)

	// R
	register.NewRouter(spec)
}
