package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/nmarsollier/imagego/docs"
	"github.com/nmarsollier/imagego/internal/env"
	"github.com/nmarsollier/imagego/internal/rest/server"
)

// StartEngine Runs gin server
func StartEngine() {
	engine := server.Router()
	InitRoutes(engine)
	engine.Run(fmt.Sprintf(":%d", env.Get().Port))
}

func InitRoutes(engine *gin.Engine) {
	initGetImageId(engine)
	initGetImageIdJpeg(engine)
	initPostImage(engine)
}
