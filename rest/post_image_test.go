package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/image/image_dao"
	"github.com/nmarsollier/imagego/rest/server"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/httpx"
	"github.com/nmarsollier/imagego/tools/log"
	"github.com/stretchr/testify/assert"
)

func TestPostImageHappyPath(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// Redis
	redisMock := image_dao.NewMockImageDao(ctrl)
	redisMock.EXPECT().Set(gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 string, arg2 string) (string, error) {
			assert.NotEmpty(t, arg2)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(httpMock, redisMock, log.NewTestLogger(ctrl, 6, 0, 1, 1))
	InitRoutes()

	req, w := server.TestPostRequest("/v1/image", testImage, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostImageError(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// Redis
	redisMock := image_dao.NewMockImageDao(ctrl)
	redisMock.EXPECT().Set(gomock.Any(), gomock.Any()).Return(errs.NotFound).Times(1)

	// REQUEST
	r := server.TestRouter(httpMock, redisMock, log.NewTestLogger(ctrl, 6, 1, 1, 1))
	InitRoutes()

	req, w := server.TestPostRequest("/v1/image", testImage, user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}

func TestPostImageNotAuthorized(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock, log.NewTestLogger(ctrl, 5, 1, 1, 1))
	InitRoutes()

	req, w := server.TestPostRequest("/v1/image", testImage, user.ID)
	r.ServeHTTP(w, req)

	server.AssertUnauthorized(t, w)
}
