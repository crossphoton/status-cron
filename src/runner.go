package src

import (
	"fmt"
	"sync"
	"time"

	"github.com/crossphoton/status-cron/src/config"
	"github.com/crossphoton/status-cron/src/entities"
	"github.com/crossphoton/status-cron/src/utils"
	"github.com/gorhill/cronexpr"
)

// runService runs the services if they pass the cron
func runService(ct time.Time, service entities.Service, wg *sync.WaitGroup) {
	if ct == cronexpr.MustParse(service.Cron).Next(ct.Add(-time.Minute)) {
		a := entities.MapService(service)
		res := a.Run()
		res.CronTime = ct

		if config.PRINT_RESULT {
			utils.Logger.Info(fmt.Sprintf("%+v", res))
		}

		entities.DB_instance.SaveResult(res)
	}

	wg.Done()
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
