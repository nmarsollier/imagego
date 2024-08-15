package tests

import (
	"bytes"
	"io"
	"net/http"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/imagego/security"
	"github.com/nmarsollier/imagego/tools/http_client"
	"github.com/nmarsollier/imagego/tools/str_tools"
)

func ExpectHttpToken(mock *http_client.MockHTTPClient, user *security.User) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(str_tools.ToJson(user))),
	}
	mock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(1)
}

func ExpectHttpUnauthorized(mock *http_client.MockHTTPClient) {
	response := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}
	mock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(1)
}
