package httpx

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func Get() HTTPClient {
	return http.DefaultClient
}
