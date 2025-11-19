package controllers

import (
	"riskmanagement/lib"

	services "riskmanagement/services/auth"
	user "riskmanagement/services/user"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      logger.Logger
	service     services.JWTAuthService
	userService user.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger logger.Logger,
	service services.JWTAuthService,
	userService user.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) GenerateToken(c *gin.Context) {
	token := jwt.service.CreateTokenGlobal()
	lib.ReturnToJson(c, 200, "200", "Generate Token Successfully", token)
}
