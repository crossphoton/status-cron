package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	DEV             bool   = false
	DB_TYPE         string = ""
	PRINT_RESULT    bool   = false
	SHUTDOWN_PERIOD int    = 30
	err             error
)

func init() {
	godotenv.Load()

	env := os.Getenv("env")
	if strings.ToLower(env) == "dev" {
		DEV = true
	}

	SHUTDOWN_PERIOD_env := os.Getenv("SHUTDOWN_PERIOD")
	if SHUTDOWN_PERIOD_env == "" {
		SHUTDOWN_PERIOD_env = "10"
	}

	SHUTDOWN_PERIOD, err = strconv.Atoi(SHUTDOWN_PERIOD_env)
	if err != nil {
		SHUTDOWN_PERIOD = 30
	}

	DB_TYPE = os.Getenv("DB_TYPE")
	if DB_TYPE == "" {
		DB_TYPE = "postgres"
	}

	PRINT_RESULT_env := os.Getenv("PRINT_RESULT")
	if PRINT_RESULT_env == "true" {
		PRINT_RESULT = true
	}
}
