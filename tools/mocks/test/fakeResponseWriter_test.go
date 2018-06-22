package test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/imagego/tools/mocks"
)

func TestResponseWriter(t *testing.T) {
	response := mocks.NewFakeResponseWriter(t)
	context, _ := gin.CreateTestContext(response)
	context.JSON(500, gin.H{"error": "Internal server error"})
	response.Assert(500, "{\"error\":\"Internal server error\"}")
}
