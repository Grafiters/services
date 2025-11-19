package routes

import (
	controllers "riskmanagement/controllers/auth"
	login "riskmanagement/controllers/managementuser"

	"riskmanagement/lib"

	"gitlab.com/golang-package-library/logger"
)

// AuthRoutes struct
type AuthRoutes struct {
	logger          logger.Logger
	handler         lib.RequestHandler
	authController  controllers.JWTAuthController
	LoginController login.ManagementUserController
}

// Setup user routes
func (s AuthRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	auth := s.handler.Gin.Group("/api/v1/auth")
	{
		auth.POST("/generateToken", s.authController.GenerateToken)
		auth.POST("/login", s.LoginController.Login)
		auth.POST("/loginBrillianApps", s.LoginController.LoginBrillianApps)
		auth.POST("/validateToken", s.LoginController.ValidateToken)
	}
}

// NewAuthRoutes creates new user controller
func NewAuthRoutes(
	handler lib.RequestHandler,
	authController controllers.JWTAuthController,
	logger logger.Logger,
	LoginController login.ManagementUserController,
) AuthRoutes {
	return AuthRoutes{
		handler:         handler,
		logger:          logger,
		authController:  authController,
		LoginController: LoginController,
	}
}
