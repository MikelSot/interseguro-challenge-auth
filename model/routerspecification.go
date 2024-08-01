package model

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type RouterSpecification struct {
	App       *fiber.App
	Logger    Logger
	DB        *sql.DB
	ExpiresAt int
	SignKey   string
}
