package env

import (
	"encoding/json"
	"os"
)

// Configuration properties
type Configuration struct {
	Port              int    `json:"port"`
	RabbitURL         string `json:"rabbitUrl"`
	RedisURL          string `json:"redisUrl"`
	SecurityServerURL string `json:"securityServerUrl"`
	WWWWPath          string `json:"wwwPath"`
}

var config *Configuration

func new() *Configuration {
	return &Configuration{
		Port:              3001,
		RabbitURL:         "amqp://localhost",
		RedisURL:          "localhost:6379",
		SecurityServerURL: "http://localhost:3000",
		WWWWPath:          "www",
	}
}

// Get Obtiene las variables de entorno del sistema
func Get() *Configuration {
	if config == nil {
		if ok := Load("config.json"); !ok {
			config = new()
		}
	}

	return config
}

// Load file properties
func Load(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		return false
	}

	loaded := new()
	err = json.NewDecoder(file).Decode(&loaded)
	if err != nil {
		return false
	}

	config = loaded
	return true
}
