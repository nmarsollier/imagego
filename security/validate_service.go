package security

import (
	"github.com/nmarsollier/imagego/tools/errs"
	gocache "github.com/patrickmn/go-cache"
)

// Validate checks if the token is valid
func Validate(token string, ctx ...interface{}) (*User, error) {
	// If it is in cache, return the cache
	if found, ok := cache.Get(token); ok {
		if user, ok := found.(*User); ok {
			return user, nil
		}
	}

	user, err := getRemoteToken(token, ctx...)
	if err != nil {
		return nil, errs.Unauthorized
	}

	// Add to cache and return
	cache.Set(token, user, gocache.DefaultExpiration)

	return user, nil
}
