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

func TestGetImageIdJpegHappyPath(t *testing.T) {
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

	req, w := mktools.TestGetRequest("/images/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetImageIdJpegInvalidImage(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestInvalidImage()

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

	req, w := mktools.TestGetRequest("/images/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	mktools.AssertInternalServerError(t, w)
}

func TestGetImageIdJpegError(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestInvalidImage()

	// Mocks Redis
	ctrl := gomock.NewController(t)
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Get(gomock.Any()).Return("", errs.NotFound).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 1, 1, 1, 0, 0, 0)
	deps.SetRedisClient(redisMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestGetRequest("/images/"+testImage.ID+"/jpeg", user.ID)
	r.ServeHTTP(w, req)

	mktools.AssertDocumentNotFound(t, w)
}
