package loadenv

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	envVar map[string]string
}

func LoadEnvFile() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error in loading env")
	}

	envVars, err := godotenv.Read()
	if err != nil {
		log.Print("Error reading env")
		return nil, err
	}

	return &Config{envVars}, nil
}

func (c *Config) GetEnv(key string) string {
	return c.envVar[key]
}
