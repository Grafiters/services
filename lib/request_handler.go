package lib

import (
	_ "riskmanagement/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler() RequestHandler {
	engine := gin.New()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return RequestHandler{Gin: engine}
}
