package app

import "os"

type Config struct {
	AppPort     string
	DataDirPath string
}

func GetConfig() Config {
	appEnv := os.Getenv("APP_ENV")
	envPort := os.Getenv("PORT")

	config := Config{}

	if appEnv == "" || appEnv != "production" {
		config.DataDirPath = "./data"
	} else {
		config.DataDirPath = "/data"
	}

	if envPort != "" {
		config.AppPort = envPort
	} else {
		config.AppPort = "5000"
	}

	return config
}
