package main

import (
	"github.com/nmarsollier/imagego/rabbit"
	"github.com/nmarsollier/imagego/rest/routes"
)

func main() {
	rabbit.Init()

	routes.StartEngine()
}
