package main

import (
	"github.com/nmarsollier/imagego/internal/engine/di"
	"github.com/nmarsollier/imagego/internal/engine/env"
	"github.com/nmarsollier/imagego/internal/engine/log"
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

	di.NewInjector(log.Get(env.Get().FluentUrl)).ConsumeLogoutService().Init()

	rest.StartEngine()
}
