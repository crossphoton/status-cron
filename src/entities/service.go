package entities

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceI interface {
	Run() Result
}

type Result struct {
	ServiceId int64     `mapstructure:"id" json:"id" gorm:"primaryKey"`
	Success   bool      `mapstructure:"success" json:"success"`
	Reason    string    `mapstructure:"reason" json:"reason"`
	CronTime  time.Time `mapstructure:"cron_time" json:"cron_time"`
}

type ServiceType string

const (
	HTTP  ServiceType = "http"
	Redis ServiceType = "redis"
	SQL   ServiceType = "sql"
	Mongo ServiceType = "mongo"
	Empty ServiceType = ""
)

type HTTPService struct {
	service Service
}

func (service *HTTPService) Run() (res Result) {
	res.ServiceId = service.service.Id

	data := service.service.Data
	status, ok := data["status"].(float64)
	if !ok || data["status"] == 0 {
		status = 200
	}
	method, ok := data["method"].(string)
	if !ok {
		method = "GET"
	}

	var body bytes.Buffer
	enc := gob.NewEncoder(&body)
	enc.Encode(data["body"])

	// Form http request
	req, err := http.NewRequest(method, service.service.Url, &body)
	if err != nil {
		res.Reason = err.Error()
		return
	}

	// Add headers
	d, ok := data["headers"].(map[string]interface{})
	if ok {
		for k, v := range d {
			req.Header.Add(k, v.(string))
		}
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Result{
			ServiceId: service.service.Id,
			Success:   false,
			Reason:    err.Error(),
		}
	}

	if resp.StatusCode != int(status) {
		res.Reason = fmt.Sprintf("Got status %d instead.", resp.StatusCode)
		return
	}

	res.Success = true
	return
}

type EmptyService struct {
	service Service
}

func (service *EmptyService) Run() Result {
	return Result{
		ServiceId: service.service.Id,
		Reason:    "Unknown service type",
		Success:   false,
	}
}

type RedisService struct {
	service Service
	client  *redis.Client
}

func (service *RedisService) Run() (res Result) {
	res.ServiceId = service.service.Id

	data := service.service.Data

	options, err := redis.ParseURL(service.service.Url)
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

type SQLService struct {
	service Service
	client  *sql.DB
}

func (service *SQLService) Run() (res Result) {
	res.ServiceId = service.service.Id
	var err error

	data := service.service.Data

	driver, ok := data["driver"].(string)
	if !ok {
		return Result{
			ServiceId: res.ServiceId,
			Success:   false,
			Reason:    "invalid driver",
		}
	}
	service.client, err = sql.Open(driver, service.service.Url)
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

type MongoService struct {
	service Service
	client  *mongo.Client
}

func (service *MongoService) Run() (res Result) {
	res.ServiceId = service.service.Id
	var err error

	data := service.service.Data

	var mongoOptions = options.Client().ApplyURI(service.service.Url)
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

func MapService(s Service) ServiceI {
	switch s.Type {
	case HTTP:
		return &HTTPService{
			service: s,
		}
	case Redis:
		return &RedisService{
			service: s,
		}
	case SQL:
		return &SQLService{
			service: s,
		}
	case Mongo:
		return &MongoService{
			service: s,
		}
	default:
		return &EmptyService{}
	}
}
