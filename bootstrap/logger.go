package bootstrap

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

func newLogger(pretty bool) model.Logger {
	formatter := &log.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     pretty,
	}

	log.SetReportCaller(true)
	log.SetFormatter(formatter)

	return log.StandardLogger()
}
