package utils

import (
	"github.com/crossphoton/status-cron/src/config"
	"go.uber.org/zap"
)

var Logger *zap.Logger = nil

func init() {
	Logger, _ = zap.NewProduction()
	if config.DEV {
		Logger = Logger.WithOptions(zap.Development())
	}
	defer Logger.Sync() // flushes buffer, if any
}
