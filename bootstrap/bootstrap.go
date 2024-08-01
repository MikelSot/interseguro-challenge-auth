package bootstrap

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-auth/infrastructure/handler"
	"github.com/MikelSot/interseguro-challenge-auth/infrastructure/handler/response"
	"github.com/MikelSot/interseguro-challenge-auth/model"
)

func Run() {
	_ = godotenv.Load()

	app := newFiber(response.ErrorHandler)
	logger := newLogger(false)
	db := getConnectionDB()

	handler.InitRoutes(model.RouterSpecification{
		Api:       app,
		Logger:    logger,
		DB:        db,
		ExpiresAt: getExpiresAtHours(),
		SignKey:   signKey,
	})

	log.Fatal(app.Listen(getPort()))
}
