package main

import (
	"github.com/nmarsollier/imagego/graph/server"
	"github.com/nmarsollier/imagego/rabbit"
	"github.com/nmarsollier/imagego/rest"
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
	rabbit.Init()
	rest.StartEngine()
}
