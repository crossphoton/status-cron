package services

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"

	"github.com/crossphoton/status-cron/src/entities"
)

const REDIS entities.ServiceType = "redis"

type RedisService struct {
	Service entities.Service
	client  *redis.Client
}

func (service *RedisService) Run() (res entities.Result) {
	res.ServiceId = service.Service.Id

	data := service.Service.Data

	options, err := redis.ParseURL(service.Service.Url)
	if err != nil {
		res.Reason = err.Error()
		return res
	}

	password, ok := data["password"].(string)
	if ok && password != "" {
		options.Password = data["password"].(string)
	}

	timeout, ok := data["timeout"].(float64)
	if ok && timeout != 0 {
		options.DialTimeout = time.Duration(timeout) * time.Second
	}

	service.client = redis.NewClient(options)
	response := service.client.Ping(context.Background())

	if response.Err() != nil {
		res.Reason = response.Err().Error()
		res.Success = false
	} else {
		res.Success = true
	}

	return
}
