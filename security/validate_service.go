package security

import (
	"github.com/nmarsollier/imagego/tools/errs"
	gocache "github.com/patrickmn/go-cache"
)

// Validate valida si el token es valido
func Validate(token string, ctx ...interface{}) (*User, error) {
	// Si esta en cache, retornamos el cache
	if found, ok := cache.Get(token); ok {
		if user, ok := found.(*User); ok {
			return user, nil
		}
	}

	user, err := getRemoteToken(token, ctx...)
	if err != nil {
		return nil, errs.Unauthorized
	}

	// Todo bien, se agrega al cache y se retorna
	cache.Set(token, user, gocache.DefaultExpiration)

	return user, nil
}
