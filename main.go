package main

import (
	"time"

	"github.com/crossphoton/status-cron/src"
	"github.com/crossphoton/status-cron/src/utils"
)

func init() {
	utils.Logger.Info("starting job")
}

func main() {
	src.Runner(time.Now())
}
