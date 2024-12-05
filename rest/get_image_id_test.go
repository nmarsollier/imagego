package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/imagego/image"
	"github.com/nmarsollier/imagego/rest/server"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/log"
	"github.com/nmarsollier/imagego/tools/redisx"
	"github.com/stretchr/testify/assert"
)

func TestGetImageIdHappyPath(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock)
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizedHead(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID+"_800", arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 0, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "800")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizedParam(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID+"_640", arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 0, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID+"?Size=640", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizeInvalid(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 0, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "180")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdInvalidDocument(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return "", errs.NotFound
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 1, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}

func TestGetImageIdResizeNotNeed(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID+"_320", arg1)
			return "", errs.NotFound
		},
	).Times(1)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 string, arg2 string, arg3 interface{}) (string, error) {
			assert.Equal(t, testImage.ID+"_320", arg1)
			assert.Equal(t, testImage.Image, arg2)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 1, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "320")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResized(t *testing.T) {
	user := security.TestUser()
	testImage := image.TestResizeImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := redisx.NewMockImageDao(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID+"_160", arg1)
			return "", errs.NotFound
		},
	).Times(1)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 string, arg2 string, arg3 interface{}) (string, error) {
			assert.Equal(t, testImage.ID+"_160", arg1)
			assert.NotEmpty(t, arg2)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(redisMock, log.NewTestLogger(ctrl, 5, 1, 1, 0))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/image/"+testImage.ID, user.ID)
	req.Header.Add("Size", "160")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}
