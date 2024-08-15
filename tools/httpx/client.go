package httpx

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func Get(ctx ...interface{}) HTTPClient {
	for _, o := range ctx {
		if client, ok := o.(HTTPClient); ok {
			return client
		}
	}

	return http.DefaultClient
}
