package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
	Addr string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	addr := host + ":" + port

	return Config{
		Port: port,
		Host: host,
		Addr: addr,
	}
}
