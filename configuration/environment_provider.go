package configuration

import (
	"log"
	"os"
	"strconv"
)

type EnvironmentProvider struct {
	*log.Logger
}

func NewEnvironmentProvider(logger *log.Logger) *EnvironmentProvider {
	return &EnvironmentProvider{Logger: logger}
}

func (e *EnvironmentProvider) GetString(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value

}

func (e *EnvironmentProvider) GetInt(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	ret, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		e.Logger.Fatalf("Unsupported Int ENV Var - %v : %v", key, value)
	}
	return int(ret)

}
