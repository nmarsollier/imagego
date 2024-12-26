package security

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/imagego/tools/env"
	"github.com/nmarsollier/imagego/tools/errs"
	"github.com/nmarsollier/imagego/tools/httpx"
	"github.com/nmarsollier/imagego/tools/log"
	gocache "github.com/patrickmn/go-cache"
)

var cache = gocache.New(60*time.Minute, 10*time.Minute)

func getRemoteToken(token string, deps ...interface{}) (*User, error) {
	// Fetch the remote user
	req, err := http.NewRequest("GET", env.Get().SecurityServerURL+"/users/current", nil)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, errs.Unauthorized
	}
	req.Header.Add("Authorization", "Bearer "+token)
	if corrId, ok := log.Get(deps...).Data()[log.LOG_FIELD_CORRELATION_ID].(string); ok {
		req.Header.Add(log.LOG_FIELD_CORRELATION_ID, corrId)
	}

	resp, err := httpx.Get(deps...).Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Get(deps...).Error(err)
		return nil, errs.Unauthorized
	}
	defer resp.Body.Close()

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}
	return user, nil
}
