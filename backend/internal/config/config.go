package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost   string
	ServerPort   string
	DbConnString string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("! Unable to read .env file !")
	}

	conf := Config{
		ServerHost:   os.Getenv("SERVER_HOST"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		DbConnString: os.Getenv("DB_CONN_STRING"),
	}

	return conf
}
