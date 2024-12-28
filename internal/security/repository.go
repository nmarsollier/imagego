package security

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/commongo/cache"
	"github.com/nmarsollier/commongo/errs"
	"github.com/nmarsollier/commongo/httpx"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/imagego/internal/env"
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
		cache:  cache.NewCache[User](),
	}
}

type securityRepository struct {
	log    log.LogRusEntry
	cache  cache.Cache[User]
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
	r.cache.Add(token, user)

	return user, nil
}

func (r *securityRepository) CleanToken(token string) {
	r.cache.Remove(token)
}

func (r *securityRepository) GetToken(token string) (*User, bool) {
	user, err := r.cache.Get(token)
	if err != nil {
		return nil, false
	}
	return user, true
}
