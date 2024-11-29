package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string
	DbConnString string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("! Unable to read .env file !")
	}

	conf := Config{
		ServerPort:   os.Getenv("SERVER_PORT"),
		DbConnString: os.Getenv("DB_CONN_STRING"),
	}

	return conf
}
