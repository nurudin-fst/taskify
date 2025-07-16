package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Overload(".env")
	if err == nil {
		slog.Info(`file .env loaded (development environment)`)
		return
	}
}
