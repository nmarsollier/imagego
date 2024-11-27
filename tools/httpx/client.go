package httpx

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func Get(deps ...interface{}) HTTPClient {
	for _, o := range deps {
		if client, ok := o.(HTTPClient); ok {
			return client
		}
	}

	return http.DefaultClient
}
