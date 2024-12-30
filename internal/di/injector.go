package di

import (
	"github.com/nmarsollier/commongo/httpx"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/redisx"
	"github.com/nmarsollier/commongo/security"
	"github.com/nmarsollier/imagego/internal/env"
	"github.com/nmarsollier/imagego/internal/image"
)

// Singletons
var redisClient redisx.RedisClient
var httpClient httpx.HTTPClient

type Injector interface {
	Logger() log.LogRusEntry
	RedisClient() redisx.RedisClient
	HttpClient() httpx.HTTPClient
	ImageRepository() image.ImageRepository
	ImageService() image.ImageService
	SecurityRepository() security.SecurityRepository
	SecurityService() security.SecurityService
}

type Deps struct {
	CurrLog         log.LogRusEntry
	CurrRedisClient redisx.RedisClient
	CurrHttpClient  httpx.HTTPClient
	CurrImageRepo   image.ImageRepository
	CurrImageSvc    image.ImageService
	CurrSecRepo     security.SecurityRepository
	CurrSecSvc      security.SecurityService
}

func NewInjector(log log.LogRusEntry) Injector {
	return &Deps{
		CurrLog: log,
	}
}

func (i *Deps) Logger() log.LogRusEntry {
	return i.CurrLog
}

func (i *Deps) RedisClient() redisx.RedisClient {
	if i.CurrRedisClient != nil {
		return i.CurrRedisClient
	}

	if redisClient != nil {
		return redisClient
	}

	redisClient = redisx.Get(env.Get().RedisURL)
	return redisClient
}

func (i *Deps) HttpClient() httpx.HTTPClient {
	if i.CurrHttpClient != nil {
		return i.CurrHttpClient
	}

	if httpClient != nil {
		return httpClient
	}

	httpClient = httpx.Get()
	return httpClient
}

func (i *Deps) ImageRepository() image.ImageRepository {
	if i.CurrImageRepo != nil {
		return i.CurrImageRepo
	}
	i.CurrImageRepo = image.NewImageRepository(i.Logger(), i.RedisClient())
	return i.CurrImageRepo
}

func (i *Deps) ImageService() image.ImageService {
	if i.CurrImageSvc != nil {
		return i.CurrImageSvc
	}
	i.CurrImageSvc = image.NewImageService(i.Logger(), i.ImageRepository())
	return i.CurrImageSvc
}

func (i *Deps) SecurityRepository() security.SecurityRepository {
	if i.CurrSecRepo != nil {
		return i.CurrSecRepo
	}
	i.CurrSecRepo = security.NewSecurityRepository(i.Logger(), i.HttpClient(), env.Get().SecurityServerURL)
	return i.CurrSecRepo
}

func (i *Deps) SecurityService() security.SecurityService {
	if i.CurrSecSvc != nil {
		return i.CurrSecSvc
	}
	i.CurrSecSvc = security.NewSecurityService(i.Logger(), i.SecurityRepository())
	return i.CurrSecSvc
}
