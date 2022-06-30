package entities

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/crossphoton/status-cron/src/config"
	"github.com/crossphoton/status-cron/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	Id   int64       `mapstructure:"id" json:"id" gorm:"primaryKey"`
	Name string      `mapstructure:"name" json:"name"`
	Url  string      `mapstructure:"url" json:"url"`
	Type ServiceType `mapstructure:"type" json:"type"`
	Cron string      `mapstructure:"cron" json:"cron"`
}

type DB interface {
	Connect()
	GetServices() []Service
	SaveResult(Result)
	Close()
}

var DB_instance DB

type DB_TYPE string

const (
	Mongo_DB    DB_TYPE = "mongo"
	Postgres_DB DB_TYPE = "postgres"
	JSON        DB_TYPE = "json"
)

type PostgresDB struct {
	client *gorm.DB
}

func (db *PostgresDB) Connect() {
	var err error

	dsn := config.GetPostgresURI()
	db.client, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))

	if err != nil {
		panic(err)
	}

	db.client.AutoMigrate(&Service{}, &Result{})

	utils.Logger.Info("connected to database")
}

func (db *PostgresDB) GetServices() []Service {
	var services []Service
	res := db.client.Find(&services)
	if res.Error != nil {
		utils.Logger.Error(fmt.Sprint("couldn't connect to database", res.Error.Error()))
		return []Service{}
	}

	return services
}

func (db *PostgresDB) SaveResult(res Result) {
	db.client.Create(res)
}

func (db *PostgresDB) Close() {
}

type JsonFile struct {
	services []Service
	results  []Result
}

type JsonDB struct {
	file     JsonFile
	filepath string
	lock     sync.Mutex
}

func (db *JsonDB) Connect() {
	db.filepath = config.GetJSONFilePath()
	file, err := os.ReadFile(db.filepath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &db.file)
}

func (db *JsonDB) GetServices() []Service {
	return db.file.services
}

func (db *JsonDB) SaveResult(res Result) {
	db.lock.Lock()
	db.file.results = append(db.file.results, res)
	db.lock.Unlock()
}

func (db *JsonDB) Close() {
	data, err := json.Marshal(db.file)
	if err != nil {
		panic(err)
	}

	os.WriteFile(db.filepath, data, 0777)
}

var db_map = map[DB_TYPE]DB{
	Postgres_DB: &PostgresDB{},
	JSON:        &JsonDB{},
}

func init() {
	DB_instance = db_map[DB_TYPE(config.DB_TYPE)]
	DB_instance.Connect()
}
