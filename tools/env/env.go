package env

import (
	"os"
	"strconv"
)

// Configuration properties
type Configuration struct {
	Port              int    `json:"port"`
	GqlPort           int    `json:"gqlPort"`
	RabbitURL         string `json:"rabbitUrl"`
	RedisURL          string `json:"redisUrl"`
	SecurityServerURL string `json:"securityServerUrl"`
	FluentUrl         string `json:"fluentUrl"`
	Source            string `json:"fluentUrl"`
	AwsAccessKeyId    string `json:"AwsAccessKeyId"`
	AwsSecret         string `json:"AwsSecret"`
	AwsRegion         string `json:"AwsRegion"`
}

var config *Configuration

func new() *Configuration {
	return &Configuration{
		Port:              3001,
		GqlPort:           4001,
		RabbitURL:         "amqp://localhost",
		RedisURL:          "localhost:6379",
		SecurityServerURL: "http://localhost:3000",
		FluentUrl:         "localhost:24224",
		Source:            "redis",
		AwsAccessKeyId:    "placeholder_value",
		AwsSecret:         "placeholder_value",
		AwsRegion:         "placeholder_value",
	}
}

// Get Obtiene las variables de entorno del sistema
func Get() *Configuration {
	if config == nil {
		config = load()
	}

	return config
}

// Load file properties
func load() *Configuration {
	result := new()

	if value := os.Getenv("REDIS_URL"); len(value) > 0 {
		result.RedisURL = value
	}

	if value := os.Getenv("RABBIT_URL"); len(value) > 0 {
		result.RabbitURL = value
	}

	if value := os.Getenv("PORT"); len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.Port = intVal
		}
	}

	if value := os.Getenv("GQL_PORT"); len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.GqlPort = intVal
		}
	}

	if value := os.Getenv("FLUENT_URL"); len(value) > 0 {
		result.FluentUrl = value
	}

	if value := os.Getenv("AUTH_SERVICE_URL"); len(value) > 0 {
		result.SecurityServerURL = value
	}

	if value := os.Getenv("SOURCE"); len(value) > 0 {
		result.Source = value
	}
	if value := os.Getenv("AWS_ACCESS_KEY_ID"); len(value) > 0 {
		result.AwsAccessKeyId = value
	}
	if value := os.Getenv("AWS_SECRET_ACCESS_KEY"); len(value) > 0 {
		result.AwsSecret = value
	}
	if value := os.Getenv("AWS_REGION"); len(value) > 0 {
		result.AwsRegion = value
	}

	return result
}
