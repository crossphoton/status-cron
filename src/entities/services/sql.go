package services

import (
	"database/sql"

	"github.com/crossphoton/status-cron/src/entities"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora"
)

const SQL entities.ServiceType = "sql"

type SQLService struct {
	Service entities.Service
	client  *sql.DB
}

func (service *SQLService) Run() (res entities.Result) {
	res.ServiceId = service.Service.Id
	var err error

	data := service.Service.Data

	driver, ok := data["driver"].(string)
	if !ok {
		return entities.Result{
			ServiceId: res.ServiceId,
			Success:   false,
			Reason:    "invalid driver",
		}
	}
	service.client, err = sql.Open(driver, service.Service.Url)
	if err != nil {
		res.Reason = err.Error()
		return
	}

	err = service.client.Ping()
	if err != nil {
		res.Reason = err.Error()
		return
	}

	res.Success = true
	return
}
