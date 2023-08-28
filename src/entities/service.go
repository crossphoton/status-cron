package entities

import (
	"time"
)

type Service struct {
	Id   int64       `mapstructure:"id" json:"id" gorm:"primaryKey"`
	Name string      `mapstructure:"name" json:"name"`
	Url  string      `mapstructure:"url" json:"url"`
	Type ServiceType `mapstructure:"type" json:"type"`
	Cron string      `mapstructure:"cron" json:"cron"`
	Data ServiceData `mapstructure:"data" json:"data"`
}

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
