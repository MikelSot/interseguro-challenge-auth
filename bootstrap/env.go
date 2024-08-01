package bootstrap

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const _nameAppDefault = "interseguro-challenge-auth"

const _portDefault = ":3000"

const _expiresAtDefault = 1

const (
	_allowOriginsDefault = "*"
	_allowMethodsDefault = "GET,POST"
)

func getApplicationName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}

func getPort() string {
	port := os.Getenv("FIBER_PORT")
	if port == "" {
		return _portDefault
	}

	return port
}

func getAllowOrigins() string {
	allowedOrigins := os.Getenv("ALLOW_ORIGINS")
	if allowedOrigins == "" {
		return _allowOriginsDefault
	}

	return allowedOrigins
}

func getAllowMethods() string {
	allowedMethods := os.Getenv("ALLOW_METHODS")
	if allowedMethods == "" {
		return _allowMethodsDefault
	}

	return allowedMethods
}

func getSignKey() string {
	signKey := os.Getenv("SIGN_KEY")
	if signKey == "" {
		log.Warn("sign key not found")

		return ""
	}

	return signKey
}

func getExpiresAtHours() int {
	expiresAt := os.Getenv("EXPIRES_AT_HOURS")
	if expiresAt == "" {
		return _expiresAtDefault
	}

	num, err := strconv.Atoi(expiresAt)
	if err != nil {
		log.Warn("invalid expires at hours")

		return _expiresAtDefault
	}

	return num
}
