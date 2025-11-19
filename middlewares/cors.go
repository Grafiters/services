package middlewares

import (
	"riskmanagement/lib"
	"riskmanagement/lib/env"

	cors "github.com/rs/cors/wrapper/gin"
	"gitlab.com/golang-package-library/logger"
)

// CorsMiddleware middleware for cors
type CorsMiddleware struct {
	handler lib.RequestHandler
	logger  logger.Logger
	env     env.Env
}

// NewCorsMiddleware creates new cors middleware
func NewCorsMiddleware(handler lib.RequestHandler, logger logger.Logger, env env.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		env:     env,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {
	// Actual response added headers:
	// map[Access-Control-Allow-Credentials:[true]
	// Access-Control-Allow-Origin:[http://localhost:7070] Vary:[Origin]]

	// Preflight response headers:
	// map[Access-Control-Allow-Credentials:[true]
	// Access-Control-Allow-Headers:[Content-Type]
	// Access-Control-Allow-Methods:[POST]
	// Access-Control-Allow-Origin:[http://localhost:7070]
	// Vary:[Origin Access-Control-Request-Method Access-Control-Request-Headers]]

	m.logger.Zap.Info("Setting up cors middleware")

	debug := m.env.Environment == "development"
	m.handler.Gin.Use(cors.New(cors.Options{
		AllowCredentials: false,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		Debug:            debug,
	}))
}
