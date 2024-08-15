package security

import (
	"bytes"
	"io"
	"net/http"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/imagego/tools/httpx"
	"github.com/nmarsollier/imagego/tools/strs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock Data
func TestUser() *User {
	return &User{
		ID:          primitive.NewObjectID().Hex(),
		Login:       "Login",
		Name:        "Name",
		Permissions: []string{"user"},
	}
}

// Http Mocks
func ExpectHttpToken(mock *httpx.MockHTTPClient, user *User) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(strs.ToJson(user))),
	}
	mock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(1)
}

func ExpectHttpUnauthorized(mock *httpx.MockHTTPClient) {
	response := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}
	mock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(1)
}
