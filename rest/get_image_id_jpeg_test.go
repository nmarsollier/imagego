package rest

import (
	"net/http"
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/imagego/rest/server"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/redisx"
	"github.com/nmarsollier/imagego/tools/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetImageIdJpegHappyPath(t *testing.T) {
	user := security.TestUser()
	testImage := tests.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID, arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdJpegInvalidImage(t *testing.T) {
	user := security.TestUser()
	testImage := tests.TestInvalidImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) *redis.StringCmd {
			assert.Equal(t, testImage.ID, arg1)
			return redis.NewStringResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertInternalServerError(t, w)
}

func TestGetImageIdJpegError(t *testing.T) {
	user := security.TestUser()
	testImage := tests.TestInvalidImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).Return(redis.NewStringResult("", errs.NotFound)).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/image/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}
