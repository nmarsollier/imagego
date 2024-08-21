package main

import (
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
	rabbit.Init()
	rest.StartEngine()
}
