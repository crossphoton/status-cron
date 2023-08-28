package services

import (
	"github.com/crossphoton/status-cron/src/entities"
	"github.com/crossphoton/status-cron/src/entities/services"
)

func MapService(s entities.Service) entities.ServiceI {
	switch s.Type {
	case services.HTTP:
		return &services.HTTPService{
			Service: s,
		}
	case services.REDIS:
		return &services.RedisService{
			Service: s,
		}
	case services.SQL:
		return &services.SQLService{
			Service: s,
		}
	case services.MONGO:
		return &services.MongoService{
			Service: s,
		}
	default:
		return &services.EmptyService{}
	}
}
