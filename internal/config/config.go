package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PrivateKey string
	Port       uint
}

var ApplicationConfig Config

func InitConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Warning: Could not load .env file", err)
	}

	ApplicationConfig.PrivateKey = os.Getenv("PRIVATE_KEY")

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Warning: Could not load .env file", err)
	}

	ApplicationConfig.Port = uint(port)

	return ApplicationConfig
}
