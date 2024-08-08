package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	_ "github.com/nmarsollier/imagego/docs"
	"github.com/nmarsollier/imagego/rest/middlewares"
	"github.com/nmarsollier/imagego/tools/env"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// StartEngine Runs gin server
func StartEngine() {
	router().Run(fmt.Sprintf(":%d", env.Get().Port))
}

var engine *gin.Engine = nil

func router() *gin.Engine {
	if engine != nil {
		return engine
	}

	// Hoy gin usa v8, para actualizar gin validator a v9.
	// binding.Validator = new(defaultValidator)

	engine = gin.Default()
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Use(middlewares.ErrorHandler)

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
