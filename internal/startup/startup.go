package startup

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server     string
	DBName     string
	DBPassword string
	DBUsername string
	CSRFSecret string
}

func LoadEnv(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	config := &Config{
		Server:     os.Getenv("SERVER_ADDRESS"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBUsername: os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		CSRFSecret: os.Getenv("CSRF_SECRET"),
	}
	if config.Server == "" {
		return nil, errors.New("server address cannot be blank")
	}
	if config.DBName == "" {
		return nil, errors.New("DBName cannot be blank")
	}
	if config.DBUsername == "" {
		return nil, errors.New("DBUsername cannot be blank")
	}
	if config.DBPassword == "" {
		return nil, errors.New("DBPassword cannot be blank")
	}
	return config, nil
}
