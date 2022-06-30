package config

import "os"

func GetPostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}

func GetJSONFilePath() string {
	return os.Getenv("JSON_PATH")
}
