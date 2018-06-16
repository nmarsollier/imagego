package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	cors "github.com/itsjamie/gin-cors"
	"github.com/nmarsollier/imagego/rabbit"
	"github.com/nmarsollier/imagego/routes"
	"github.com/nmarsollier/imagego/tools/env"
)

func main() {
	if len(os.Args) > 1 {
		env.Load(os.Args[1])
	}

	rabbit.Init()

	// Hoy gin usa v8, para actualizar gin validator a v9.
	binding.Validator = new(defaultValidator)

	server := gin.Default()

	server.Use(gzip.Gzip(gzip.DefaultCompression))

	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Size",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	server.Use(static.Serve("/", static.LocalFile(env.Get().WWWWPath, true)))

	server.POST("/v1/image", routes.NewImage)
	server.GET("/v1/image/:imageID", routes.GetImage)
	server.GET("/v1/image/:imageID/jpeg", routes.GetImageJpeg)

	server.Run(fmt.Sprintf(":%d", env.Get().Port))
}
