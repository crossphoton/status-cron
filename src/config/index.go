package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	DEV          bool   = false
	DB_TYPE      string = ""
	PRINT_RESULT bool   = false
)

func init() {
	godotenv.Load()

	env := os.Getenv("env")
	if strings.ToLower(env) == "dev" {
		DEV = true
	}

	DB_TYPE = os.Getenv("DB_TYPE")
	if DB_TYPE == "" {
		DB_TYPE = "postgres"
	}

	PRINT_RESULT_env := os.Getenv("DB_TYPE")
	if PRINT_RESULT_env == "true" {
		PRINT_RESULT = true
	}
}
