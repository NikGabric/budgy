package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         int
	DBConnectionString string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, err
	}

	config := &Config{
		ServerPort:         serverPort,
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}

	return config, nil
}
