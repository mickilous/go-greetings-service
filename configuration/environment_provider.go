package configuration

import (
	"fmt"
	"os"
	"strconv"
)

type EnvironmentProvider struct{}

func NewEnvironmentProvider() *EnvironmentProvider {
	return &EnvironmentProvider{}
}

func (e *EnvironmentProvider) GetString(key string) string {
	return e.GetStringOr(key, "")
}

func (e *EnvironmentProvider) GetStringOr(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func (e *EnvironmentProvider) GetInt(key string) (int, error) {
	return e.GetIntOr(key, 0)
}

func (e *EnvironmentProvider) GetIntOr(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback, nil
	}
	ret, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("unsupported int env var - %v - %v : %v", err, key, value)
	}
	return int(ret), nil
}
