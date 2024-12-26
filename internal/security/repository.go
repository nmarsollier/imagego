package security

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/imagego/internal/engine/env"
	"github.com/nmarsollier/imagego/internal/engine/errs"
	"github.com/nmarsollier/imagego/internal/engine/httpx"
	"github.com/nmarsollier/imagego/internal/engine/log"
	gocache "github.com/patrickmn/go-cache"
)

type SecurityRepository interface {
	GetRemoteToken(token string) (*User, error)
	CleanToken(token string)
	GetToken(token string) (*User, bool)
}

func NewSecurityRepository(
	log log.LogRusEntry,
	client httpx.HTTPClient,
) SecurityRepository {
	return &securityRepository{
		log:    log,
		client: client,
		cache:  gocache.New(60*time.Minute, 10*time.Minute),
	}
}

type securityRepository struct {
	log    log.LogRusEntry
	cache  *gocache.Cache
	client httpx.HTTPClient
}

func (r *securityRepository) GetRemoteToken(token string) (*User, error) {
	// Fetch the remote user
	req, err := http.NewRequest("GET", env.Get().SecurityServerURL+"/users/current", nil)
	if err != nil {
		r.log.Error(err)
		return nil, errs.Unauthorized
	}

	req.Header.Add("Authorization", "Bearer "+token)
	if corrId, ok := r.log.Data()[log.LOG_FIELD_CORRELATION_ID].(string); ok {
		req.Header.Add(log.LOG_FIELD_CORRELATION_ID, corrId)
	}

	resp, err := r.client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		r.log.Error(err)
		return nil, errs.Unauthorized
	}
	defer resp.Body.Close()

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		r.log.Error(err)
		return nil, err
	}

	// Add to cache and return
	r.cache.Set(token, user, gocache.DefaultExpiration)

	return user, nil
}

func (r *securityRepository) CleanToken(token string) {
	r.cache.Delete(token)
}

func (r *securityRepository) GetToken(token string) (*User, bool) {
	if found, ok := r.cache.Get(token); ok {
		if user, ok := found.(*User); ok {
			return user, true
		}
	}
	return nil, false
}
