package rest

import (
	"fmt"

	_ "github.com/nmarsollier/imagego/docs"
	"github.com/nmarsollier/imagego/rest/server"
	"github.com/nmarsollier/imagego/tools/env"
)

// StartEngine Runs gin server
func StartEngine() {
	InitRoutes()
	server.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}

func InitRoutes() {
	initGetImageId()
	initGetImageIdJpeg()
	initPostImage()
}
