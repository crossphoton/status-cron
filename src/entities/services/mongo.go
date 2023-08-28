package services

import (
	"context"
	"time"

	"github.com/crossphoton/status-cron/src/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO entities.ServiceType = "mongo"

type MongoService struct {
	Service entities.Service
	client  *mongo.Client
}

func (service *MongoService) Run() (res entities.Result) {
	res.ServiceId = service.Service.Id
	var err error

	data := service.Service.Data

	var mongoOptions = options.Client().ApplyURI(service.Service.Url)
	timeout, ok := data["timeout"].(float64)
	if ok && timeout != 0 {
		mongoOptions.SetConnectTimeout(time.Duration(timeout) * time.Second)
	}

	service.client, err = mongo.Connect(context.TODO(), mongoOptions)
	if err != nil {
		res.Reason = err.Error()
		return
	}

	err = service.client.Ping(context.TODO(), nil)
	if err != nil {
		res.Reason = err.Error()
		return
	}

	res.Success = true
	return
}
