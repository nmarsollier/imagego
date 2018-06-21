package env

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	jsonErr, _ := json.Marshal(Get())
	assert.Equal(t, "{\"port\":3001,\"rabbitUrl\":\"amqp://localhost\",\"redisUrl\":\"localhost:6379\",\"securityServerUrl\":\"http://localhost:3000\",\"wwwPath\":\"www\"}", string(jsonErr))

}

func TestLoad(t *testing.T) {
	loaded := Load("./test/mocks/missing.json")
	assert.Equal(t, loaded, false)
	jsonErr, _ := json.Marshal(Get())
	assert.Equal(t, "{\"port\":3001,\"rabbitUrl\":\"amqp://localhost\",\"redisUrl\":\"localhost:6379\",\"securityServerUrl\":\"http://localhost:3000\",\"wwwPath\":\"www\"}", string(jsonErr))

	loaded = Load("missing.json")
	assert.Equal(t, loaded, false)
	jsonErr, _ = json.Marshal(Get())
	assert.Equal(t, "{\"port\":3001,\"rabbitUrl\":\"amqp://localhost\",\"redisUrl\":\"localhost:6379\",\"securityServerUrl\":\"http://localhost:3000\",\"wwwPath\":\"www\"}", string(jsonErr))

	loaded = Load("env_test_config.json")
	assert.Equal(t, loaded, true)
	jsonErr, _ = json.Marshal(Get())
	assert.Equal(t, "{\"port\":80,\"rabbitUrl\":\"otroUrl\",\"redisUrl\":\"localhost:6379\",\"securityServerUrl\":\"http://localhost:3000\",\"wwwPath\":\"www\"}", string(jsonErr))
}
