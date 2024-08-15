package rest

import (
	"net/http"
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/imagego/rest/server"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/httpx"
	"github.com/nmarsollier/imagego/tools/redisx"
	"github.com/nmarsollier/imagego/tools/tests"
	"github.com/stretchr/testify/assert"
)

func TestPostImageHappyPath(t *testing.T) {
	user := security.TestUser()
	testImage := tests.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// Redis
	redisMock := redisx.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 string, arg2 string, arg3 interface{}) *redis.StatusCmd {
			assert.NotEmpty(t, arg2)
			return redis.NewStatusResult(testImage.Image, nil)
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(httpMock, redisMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/image", testImage, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostImageError(t *testing.T) {
	user := security.TestUser()
	testImage := tests.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// Redis
	redisMock := redisx.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(redis.NewStatusResult("", errs.NotFound)).Times(1)

	// REQUEST
	r := server.TestRouter(httpMock, redisMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/image", testImage, user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}

func TestPostImageNotAuthorized(t *testing.T) {
	user := security.TestUser()
	testImage := tests.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/image", testImage, user.ID)
	r.ServeHTTP(w, req)

	tests.AssertUnauthorized(t, w)
}
