package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/crossphoton/status-cron/src"
	"github.com/crossphoton/status-cron/src/utils"
)

var (
	startTime    time.Time
	staticServer bool
	helpFlag     bool
)

func init() {
	startTime = time.Now().Round(time.Minute).Add(-time.Minute)
}

func main() {
	// Flags Parsing
	flag.BoolVar(&staticServer, "static", true, "whether to start a static server or one time job")
	flag.BoolVar(&helpFlag, "help", false, "print help")
	flag.Parse()

	if helpFlag {
		flag.PrintDefaults()
		return
	}

	if staticServer {
		utils.Logger.Info("starting static server")
		fmt.Println("starting static server")
		src.StaticServerRunner(startTime)
	} else {
		utils.Logger.Info("starting job")
		src.Runner(startTime)
	}

}
