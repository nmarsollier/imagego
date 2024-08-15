package security

import (
	"time"

	"github.com/golang/glog"
	"github.com/nmarsollier/imagego/tools/errs"
	gocache "github.com/patrickmn/go-cache"
)

var cache = gocache.New(60*time.Minute, 10*time.Minute)

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

// Invalidate invalida un token del cache
func Invalidate(token string) {
	cache.Delete(token)
	glog.Info("Token invalidado: ", token)
}
