package env

import (
	"cmp"
	"os"

	"github.com/nmarsollier/commongo/strs"
)

// Configuration properties
type Configuration struct {
	ServerName        string `json:"serverName"`
	Port              int    `json:"port"`
	GqlPort           int    `json:"gqlPort"`
	RabbitURL         string `json:"rabbitUrl"`
	RedisURL          string `json:"redisUrl"`
	SecurityServerURL string `json:"securityServerUrl"`
	FluentURL         string `json:"fluentUrl"`
}

var config *Configuration

func Get() *Configuration {
	if config == nil {
		config = load()
	}

	return config
}

// Load file properties
func load() *Configuration {
	return &Configuration{
		ServerName:        cmp.Or(os.Getenv("SERVER_NAME"), "imagego"),
		Port:              cmp.Or(strs.AtoiZero(os.Getenv("PORT")), 3001),
		GqlPort:           cmp.Or(strs.AtoiZero(os.Getenv("GQL_PORT")), 4001),
		RabbitURL:         cmp.Or(os.Getenv("RABBIT_URL"), "amqp://localhost"),
		RedisURL:          cmp.Or(os.Getenv("REDIS_URL"), "localhost:6379"),
		SecurityServerURL: cmp.Or(os.Getenv("AUTH_SERVICE_URL"), "http://localhost:3000"),
		FluentURL:         cmp.Or(os.Getenv("FLUENT_URL"), "localhost:24224"),
	}
}
