package routes

import (
	"fmt"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/nmarsollier/imagego/middlewares"
	"github.com/nmarsollier/imagego/tools/env"
)

var router *gin.Engine = nil

// StartEngine Runs gin server
func StartEngine() {
	getRouter().Run(fmt.Sprintf(":%d", env.Get().Port))
}

func getRouter() *gin.Engine {
	if router != nil {
		return router
	}

	// Hoy gin usa v8, para actualizar gin validator a v9.
	// binding.Validator = new(defaultValidator)

	router = gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middlewares.ErrorHandler)

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Size",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Use(static.Serve("/", static.LocalFile(env.Get().WWWWPath, true)))

	return router
}
