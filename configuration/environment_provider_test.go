package configuration_test

import (
	"github.com/stretchr/testify/assert"
	"greetings-service/configuration"
	"os"
	"strconv"
	"testing"
)

var (
	provider = configuration.NewEnvironmentProvider()
)

func TestGetString(t *testing.T) {
	key := "key"
	value := "value"

	os.Setenv(key, "value")

	result := provider.GetString(key)
	assert.Equal(t, value, result, "Invalid value returned")
}

func TestGetStringOr(t *testing.T) {
	key := "key"
	fallback := "fallback"

	os.Clearenv()

	result := provider.GetStringOr(key, fallback)
	assert.Equal(t, fallback, result, "Fallback should be returned")
}

func TestGetInt(t *testing.T) {
	key := "key"
	value := 42

	os.Setenv(key, strconv.Itoa(value))

	result, _ := provider.GetInt(key)
	assert.Equal(t, value, result, "Invalid value returned")
}

func TestGetIntOr(t *testing.T) {
	key := "key"
	fallback := 22

	os.Clearenv()

	result, _ := provider.GetIntOr(key, fallback)
	assert.Equal(t, fallback, result, "Fallback should be returned")
}

func TestGetInt_ErrorConversion(t *testing.T) {

}
