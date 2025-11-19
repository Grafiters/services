package routes

import (
	controllers "riskmanagement/controllers/product"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type ProductRoutes struct {
	logger            logger.Logger
	handler           lib.RequestHandler
	ProductController controllers.ProductController
	authMiddleware    middlewares.JWTAuthMiddleware
}

func (s ProductRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/product").Use(s.authMiddleware.Handler())
	{
		api.POST("/getAll", s.ProductController.GetAll)
		api.POST("/getAllWithPage", s.ProductController.GetAllWithPage)
		api.POST("/getOne", s.ProductController.GetOne)
		api.POST("/store", s.ProductController.Store)
		api.POST("/update", s.ProductController.Update)
		api.POST("/delete", s.ProductController.Delete)
		api.POST("/getKodeProduct", s.ProductController.GetKodeProduct)
		api.POST("/getProductByActivity", s.ProductController.GetProductByActivity)
		api.POST("/searchProduct", s.ProductController.SearchProduct)
		api.POST("/getProductBySegment", s.ProductController.GetProductBySegment)
	}
}

func NewProductRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	ProductController controllers.ProductController,
	authMiddleware middlewares.JWTAuthMiddleware,
) ProductRoutes {
	return ProductRoutes{
		handler:           handler,
		logger:            logger,
		ProductController: ProductController,
		authMiddleware:    authMiddleware,
	}
}
