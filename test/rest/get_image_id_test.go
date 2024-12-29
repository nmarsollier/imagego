package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/commongo/errs"
	"github.com/nmarsollier/commongo/test/mktools"
	"github.com/nmarsollier/commongo/test/mockgen"
	"github.com/nmarsollier/imagego/internal/rest"
	"github.com/nmarsollier/imagego/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetImageIdHappyPath(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 1, 0, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizedHead(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID+"_800", arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 1, 0, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID, user.ID)
	req.Header.Add("Size", "800")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizedParam(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID+"_640", arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 1, 0, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID+"?Size=640", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResizeInvalid(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 1, 0, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID, user.ID)
	req.Header.Add("Size", "180")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdInvalidDocument(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).DoAndReturn(
		func(arg1 string) (string, error) {
			assert.Equal(t, testImage.ID, arg1)
			return "", errs.NotFound
		},
	).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 1, 1, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID, user.ID)
	r.ServeHTTP(w, req)

	mktools.AssertDocumentNotFound(t, w)
}

func TestGetImageIdResizeNotNeed(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
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
	deps := mock.NewTestInjector(ctrl, 1, 1, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID, user.ID)
	req.Header.Add("Size", "320")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdResized(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestResizeImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
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
	deps := mock.NewTestInjector(ctrl, 1, 1, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID, user.ID)
	req.Header.Add("Size", "160")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}
