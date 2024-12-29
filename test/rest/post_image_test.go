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

func TestPostImageHappyPath(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := mockgen.NewMockHTTPClient(ctrl)
	mktools.ExpectHttpToken(httpMock, user)

	// Redis
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 string, arg2 string, arg3 interface{}) (string, error) {
			assert.NotEmpty(t, arg2)
			return testImage.Image, nil
		},
	).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 2, 0, 1, 1, 0, 0)
	deps.SetRedisClient(redisMock)
	deps.SetHttpClient(httpMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestPostRequest("/images/create", testImage, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostImageError(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := mockgen.NewMockHTTPClient(ctrl)
	mktools.ExpectHttpToken(httpMock, user)

	// Redis
	redisMock := mockgen.NewMockRedisClient(ctrl)
	redisMock.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return("", errs.NotFound).Times(1)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 2, 1, 1, 1, 0, 0)
	deps.SetRedisClient(redisMock)
	deps.SetHttpClient(httpMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestPostRequest("/images/create", testImage, user.ID)
	r.ServeHTTP(w, req)

	mktools.AssertDocumentNotFound(t, w)
}

func TestPostImageNotAuthorized(t *testing.T) {
	user := mktools.TestUser()
	testImage := mock.TestImage()

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := mockgen.NewMockHTTPClient(ctrl)
	mktools.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	deps := mock.NewTestInjector(ctrl, 1, 1, 1, 1, 0, 0)
	deps.SetHttpClient(httpMock)

	r := TestRouter(ctrl, deps)
	rest.InitRoutes(r)

	req, w := mktools.TestPostRequest("/images/create", testImage, user.ID)
	r.ServeHTTP(w, req)

	mktools.AssertUnauthorized(t, w)
}
