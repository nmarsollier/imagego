package security

import (
	"github.com/nmarsollier/imagego/tools/log"
)

// Invalidate invalidates a token from the cache
func Invalidate(token string, deps ...interface{}) {
	if len(token) <= 7 {
		log.Get(deps...).Info("Token no valido: ", token)
		return
	}

	cache.Delete(token)
	log.Get(deps...).Info("Token invalidado: ", token)
}
