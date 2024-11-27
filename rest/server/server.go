package server

import (
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	_ "github.com/nmarsollier/imagego/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var engine *gin.Engine = nil

func Router(deps ...interface{}) *gin.Engine {
	if engine != nil {
		return engine
	}

	engine = gin.Default()
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Use(GinLoggerMiddleware(deps...))
	engine.Use(ErrorHandler)

	engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Size",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return engine
}

// Gets the context for external services
func GinDeps(c *gin.Context) []interface{} {
	var deps []interface{}
	// mock_deps solo es para mocks en testing
	if mocks, ok := c.Get("mock_deps"); ok {
		deps = mocks.([]interface{})
	}

	deps = append(deps, ginLogger(c))

	return deps
}
