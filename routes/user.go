package routes

import (
	controllers "riskmanagement/controllers/user"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

// UserRoutes struct
type UserRoutes struct {
	logger         logger.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/user")
	// .Use(s.authMiddleware.Handler())
	{
		api.POST("/login", s.userController.Login)
		// api.GET("/getAll", s.userController.GetUser)
		// api.GET("/getOne/:id", s.userController.GetOneUser)
		// api.POST("/store", s.userController.SaveUser)
		// api.POST("/store-no-trx", s.userController.SaveUserWOTrx)
		// api.POST("/update/:id", s.userController.UpdateUser)
		// api.DELETE("/delete/:id", s.userController.DeleteUser)
		api.POST("/getMenu", s.userController.GetMenu)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	userController controllers.UserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}
