package mock

import (
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/commongo/httpx"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/redisx"
	"github.com/nmarsollier/commongo/security"
	"github.com/nmarsollier/commongo/test/mktools"
	"github.com/nmarsollier/imagego/internal/di"
	"github.com/nmarsollier/imagego/internal/image"
)

type TestInjector struct {
	di.Deps
}

func NewTestInjector(
	ctrl *gomock.Controller, withFieldCount int, errorCount int, infoCount int, dataCount int, warnCount int, fatalCount int,
) *TestInjector {

	result := &TestInjector{
		Deps: di.Deps{
			CurrLog: mktools.NewTestLogger(ctrl, withFieldCount, errorCount, infoCount, dataCount, warnCount, fatalCount),
		},
	}

	return result
}

func (t *TestInjector) SetLogger(log log.LogRusEntry) {
	t.CurrLog = log
}

func (t *TestInjector) SetRedisClient(redisClient redisx.RedisClient) {
	t.CurrRedisClient = redisClient
}

func (t *TestInjector) SetHttpClient(httpClient httpx.HTTPClient) {
	t.CurrHttpClient = httpClient
}

func (t *TestInjector) SetImageRepository(imageRepository image.ImageRepository) {
	t.CurrImageRepo = imageRepository
}

func (t *TestInjector) SetImageService(imageService image.ImageService) {
	t.CurrImageSvc = imageService
}

func (t *TestInjector) SetSecurityRepository(securityRepository security.SecurityRepository) {
	t.CurrSecRepo = securityRepository
}

func (t *TestInjector) SetSecurityService(securityService security.SecurityService) {
	t.CurrSecSvc = securityService
}
