package rest

import (
	"net/http"
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/imagego/rest/server"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/redis_client"
	"github.com/nmarsollier/imagego/tools/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetImageIdHappyPath(t *testing.T) {
	user := tests.TestUser()
	testImage := tests.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID, arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizedHead(t *testing.T) {
	user := tests.TestUser()
	testImage := tests.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID+"_160", arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "160")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizedParam(t *testing.T) {
	user := tests.TestUser()
	testImage := tests.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID+"_160", arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID+"?Size=160", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizeInvalid(t *testing.T) {
	user := tests.TestUser()
	testImage := tests.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID, arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "180")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdInvalidDocument(t *testing.T) {
	user := tests.TestUser()
	testImage := tests.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID, arg1)
			return redis.NewStringResult("", errs.NotFound)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}

func TestGetImageIdResizeNotNeed(t *testing.T) {
	user := tests.TestUser()
	testImage := tests.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID+"_160", arg1)
			return redis.NewStringResult("", errs.NotFound)
		},
	).Times(1)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID, arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 string, arg2 string, arg3 interface{}) *redis.StatusCmd {
			assert.Equal(t, testImage.ID+"_160", arg1)
			assert.Equal(t, testImage.Image, arg2)
			return redis.NewStatusResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "160")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResized(t *testing.T) {
	user := tests.TestUser()
	testImage := tests.TestResizeImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redis_client.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID+"_160", arg1)
			return redis.NewStringResult("", errs.NotFound)
		},
	).Times(1)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID, arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 string, arg2 string, arg3 interface{}) *redis.StatusCmd {
			assert.Equal(t, testImage.ID+"_160", arg1)
			assert.NotEmpty(t, arg2)
			return redis.NewStatusResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "160")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}
