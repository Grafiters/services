package routes

import (
	controllers "riskmanagement/controllers/linkcage"
	"riskmanagement/lib"
	"riskmanagement/middlewares"
)

type LinkcageRoutes struct {
	handler            lib.RequestHandler
	LinkcageController controllers.LinkcageController
	authMiddleware     middlewares.JWTAuthMiddleware
}

func (s LinkcageRoutes) Setup() {
	api := s.handler.Gin.Group("/api/v1/linkcage").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.LinkcageController.GetAll)
		api.POST("/store", s.LinkcageController.Store)
		api.POST("/setStatus", s.LinkcageController.SetStatus)
		api.POST("/getActive", s.LinkcageController.GetActive)
		api.POST("/delete", s.LinkcageController.Delete)
		api.POST("/getOne", s.LinkcageController.GetOne)
		api.POST("/update", s.LinkcageController.Update)
	}
}

func NewLinkcageRouter(
	handler lib.RequestHandler,
	LinkcageController controllers.LinkcageController,
	authMiddleware middlewares.JWTAuthMiddleware,
) LinkcageRoutes {
	return LinkcageRoutes{
		handler:            handler,
		LinkcageController: LinkcageController,
		authMiddleware:     authMiddleware,
	}
}
