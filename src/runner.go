package src

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/crossphoton/status-cron/src/config"
	"github.com/crossphoton/status-cron/src/entities"
	"github.com/crossphoton/status-cron/src/utils"
	"github.com/crossphoton/status-cron/src/utils/services"
	"github.com/gorhill/cronexpr"
)

// runService runs the services if they pass the cron
func runService(ct time.Time, service entities.Service, wg *sync.WaitGroup) {
	if ct == cronexpr.MustParse(service.Cron).Next(ct.Add(-time.Minute)) {
		a := services.MapService(service)
		res := a.Run()
		res.CronTime = ct

		if config.PRINT_RESULT {
			utils.Logger.Info(fmt.Sprintf("%+v", res))
		}

		entities.DB_instance.SaveResult(res)
	}

	wg.Done()
}

func gracefulShutdown(wg *sync.WaitGroup) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	ctx := context.Background()
	go func() {
		wg.Wait()
		ctx.Done()
	}()

	// Wait for pending jobs to complete
	utils.Logger.Info("shutting down gracefully")
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(config.SHUTDOWN_PERIOD))
	defer cancel()

	<-ctx.Done()

	utils.Logger.Info("all jobs completed")

	// Close the database connection
	utils.Logger.Info("closing database connection")
	entities.DB_instance.Close()
	utils.Logger.Info("database connection closed")

	// Shutdown with exit code 1
	os.Exit(1)
}

func StaticServerRunner(ct time.Time) {
	wg := new(sync.WaitGroup)
	var next time.Time = ct

	// Graceful Shutdown
	go gracefulShutdown(wg)

	for {
		ct = next
		next = ct.Add(time.Minute)
		wg.Add(1)
		go func() {
			Runner(ct)
			wg.Done()
		}()
	}
}

// Runner runs the job
func Runner(ct time.Time) {
	services := entities.DB_instance.GetServices()
	wg := sync.WaitGroup{}
	wg.Add(len(services))

	for _, s := range services {
		go runService(ct, s, &wg)
	}

	wg.Wait()

	entities.DB_instance.Close()
}
