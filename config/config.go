package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
	Addr string
}

func Load() Config {
	_ = godotenv.Load()

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
