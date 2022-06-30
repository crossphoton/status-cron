package entities

import "net/http"

type ServiceI interface {
	Run() Result
}

type Result struct {
	ServiceId int64 `mapstructure:"id" json:"id" gorm:"primaryKey"`
	Success   bool  `mapstructure:"success" json:"success"`
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

func (service HTTPService) Run() Result {
	resp, err := http.Get(service.service.Url)
	if err != nil {
		return Result{
			ServiceId: service.service.Id,
			Success:   false,
		}
	}

	return Result{
		ServiceId: service.service.Id,
		Success:   resp.StatusCode == 200,
	}
}

type EmptyService struct {
}

func (service EmptyService) Run() Result {
	return Result{}
}

func MapService(s Service) ServiceI {
	switch s.Type {
	case HTTP:
		return HTTPService{
			service: s,
		}
	// case Redis:
	// 	return HTTPService{
	// 		service: s,
	// 	}
	// case SQL:
	// 	return HTTPService{
	// 		service: s,
	// 	}
	// case Mongo:
	// 	return HTTPService{
	// 		service: s,
	// 	}
	default:
		return EmptyService{}
	}
}
