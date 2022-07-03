package main

import (
	"time"

	"github.com/crossphoton/status-cron/src"
	"github.com/crossphoton/status-cron/src/utils"
)

var startTime time.Time

func init() {
	utils.Logger.Info("starting job")
	startTime = time.Now().Round(time.Minute).Add(-time.Minute)
}

func main() {
	src.Runner(startTime)
}
