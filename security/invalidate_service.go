package security

import (
	"github.com/nmarsollier/imagego/tools/log"
)

// Invalidate invalida un token del cache
func Invalidate(token string, ctx ...interface{}) {
	if len(token) <= 7 {
		log.Get(ctx...).Info("Token no valido: ", token)
		return
	}

	cache.Delete(token)
	log.Get(ctx...).Info("Token invalidado: ", token)
}
