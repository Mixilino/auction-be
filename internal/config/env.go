package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"sync"
)

var appConfig AppConfig
var configOnce sync.Once

func loadConfigFromEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := env.Parse(&appConfig); err != nil {
		panic(err)
	}
}

// Load loads the environment variables into the AppConfig struct if it's not already initialized
func Load() {
	configOnce.Do(loadConfigFromEnv)
}

func GetDBHost() string {
	return appConfig.DBHost
}

func GetDBPort() int {
	return appConfig.DBPort
}

func GetDBUser() string {
	return appConfig.DBUser
}

func GetDBPassword() string {
	return appConfig.DBPassword
}

func GetDBDriver() string {
	return appConfig.DBDriver
}

func GetDBName() string {
	return appConfig.DBName
}

func GetSSLMode() string {
	return appConfig.SSLMode
}

func GetJWTSecret() string {
	return appConfig.JWTSecret
}
