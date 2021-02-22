package main

import (
	"github.com/nmarsollier/imagego/rabbit"
	"github.com/nmarsollier/imagego/routes"
)

func main() {
	rabbit.Init()

	routes.StartEngine()
}
