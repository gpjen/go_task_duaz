package config

import (
	"fmt"
	"os"
)

type environmentConfig struct {
	NodeEnv    string
	DBUsername string
	DBPassword string
	DBHost     string
	DBDatabase string
	DBPort     string
	JWTSecret  string
}

func LoadEnvConfig() environmentConfig {

	nodeENV := os.Getenv("NODE_ENV")

	fmt.Println("> NODE_ENV: ", nodeENV)

	switch nodeENV {
	case "development":
		return environmentConfig{
			NodeEnv:    nodeENV,
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBDatabase: os.Getenv("DB_DATABASE_DEVELOPMENT"),
			DBPort:     os.Getenv("DB_PORT"),
			JWTSecret:  os.Getenv("JWT_SECRET"),
		}
	case "test":
		return environmentConfig{
			NodeEnv:    nodeENV,
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBDatabase: os.Getenv("DB_DATABASE_TEST"),
			DBPort:     os.Getenv("DB_PORT"),
			JWTSecret:  os.Getenv("JWT_SECRET"),
		}
	case "production":
		return environmentConfig{
			NodeEnv:    nodeENV,
			DBUsername: os.Getenv("DB_USERNAME"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBDatabase: os.Getenv("DB_DATABASE_PRODUCTION"),
			DBPort:     os.Getenv("DB_PORT"),
			JWTSecret:  os.Getenv("JWT_SECRET"),
		}
	default:
		panic("Unknown or unsupported NODE_ENV value")
	}
}
