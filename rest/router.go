package rest

import (
	"fmt"

	_ "github.com/nmarsollier/imagego/docs"
	"github.com/nmarsollier/imagego/rest/engine"
	"github.com/nmarsollier/imagego/tools/env"
)

// StartEngine Runs gin server
func StartEngine() {
	engine.Router().Run(fmt.Sprintf(":%d", env.Get().Port))
}
