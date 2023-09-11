package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	PrivateKeys []string
	Port        uint
	RPC_URL     string
}

var ApplicationConfig Config

func InitConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Warning: Could not load .env file", err)
	}

	privateKeys := os.Getenv("PRIVATE_KEYS")

	ApplicationConfig = Config{}

	ApplicationConfig.PrivateKeys = strings.Split(privateKeys, ",")

	ApplicationConfig.RPC_URL = os.Getenv("RPC_URL")

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Warning: Could not load .env file", err)
	}

	ApplicationConfig.Port = uint(port)

	return ApplicationConfig
}
