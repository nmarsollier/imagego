package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	// Test case 1: Default values
	config := load()
	assert.Equal(t, 3001, config.Port)
	assert.Equal(t, "amqp://localhost", config.RabbitURL)
	assert.Equal(t, "postgres://postgres@localhost:5432/postgres", config.PostgresURL)
	assert.Equal(t, "http://localhost:3000", config.SecurityServerURL)
	assert.Equal(t, "localhost:24224", config.FluentUrl)
}

func TestLoad(t *testing.T) {
	// Test case 2: Custom values from environment variables
	os.Setenv("POSTGRES_URL", "custom_postgres_url")
	os.Setenv("RABBIT_URL", "custom_rabbit_url")
	os.Setenv("PORT", "8080")
	os.Setenv("FLUENT_URL", "custom_fluent_url")
	os.Setenv("AUTH_SERVICE_URL", "custom_auth_service_url")

	config = load()
	assert.Equal(t, "custom_postgres_url", config.PostgresURL)
	assert.Equal(t, "custom_rabbit_url", config.RabbitURL)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, "custom_fluent_url", config.FluentUrl)
	assert.Equal(t, "custom_auth_service_url", config.SecurityServerURL)

	// Clean up environment variables
	os.Unsetenv("REDIS_URL")
	os.Unsetenv("RABBIT_URL")
	os.Unsetenv("PORT")
	os.Unsetenv("FLUENT_URL")
	os.Unsetenv("AUTH_SERVICE_URL")
}
