package services

import "github.com/crossphoton/status-cron/src/entities"

const EMPTY entities.ServiceType = ""

type EmptyService struct {
	service entities.Service
}

func (service *EmptyService) Run() entities.Result {
	return entities.Result{
		ServiceId: service.service.Id,
		Reason:    "Unknown service type",
		Success:   false,
	}
}
