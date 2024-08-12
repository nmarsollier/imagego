package main

import (
	"flag"

	"github.com/nmarsollier/imagego/rabbit"
	"github.com/nmarsollier/imagego/rest/routes"
)

//	@title			ImageGo
//	@version		1.0
//	@description	Microservicio de Imagenes.

//	@contact.name	Nestor Marsollier
//	@contact.email	nmarsollier@gmail.com

// @host		localhost:3001
// @BasePath	/v1
func main() {
	// For logging
	flag.Parse()
	flag.Set("logtostderr", "true")
	flag.Set("v", "2")

	rabbit.Init()
	routes.StartEngine()
}
