package main

import (
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/imagego/internal/di"
	"github.com/nmarsollier/imagego/internal/env"
	server "github.com/nmarsollier/imagego/internal/graph"
	"github.com/nmarsollier/imagego/internal/rest"
)

//	@title			ImageGo
//	@version		1.0
//	@description	Microservicio de Imagenes.

//	@contact.name	Nestor Marsollier
//	@contact.email	nmarsollier@gmail.com

// @host		localhost:3001
// @BasePath	/v1
func main() {
	go server.Start()

	di.NewInjector(log.Get(env.Get().FluentUrl, "imagego")).ConsumeLogoutService().Init()

	rest.StartEngine()
}
