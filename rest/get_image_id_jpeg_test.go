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
	"github.com/nmarsollier/imagego/tools/log"
	"github.com/stretchr/testify/assert"
)

func TestGetImageIdJpegHappyPath(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := image_dao.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 0, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdJpegInvalidImage(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestInvalidImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := image_dao.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 0, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	server.AssertInternalServerError(t, w)
}

func TestGetImageIdJpegError(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestInvalidImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := image_dao.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).Return("", errs.NotFound).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 1, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}
