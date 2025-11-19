package routes

import (
	controllers "riskmanagement/controllers/filemanager"
	"riskmanagement/lib"
	"riskmanagement/middlewares"

	"gitlab.com/golang-package-library/logger"
)

type FileManagerRoutes struct {
	logger                logger.Logger
	handler               lib.RequestHandler
	FileManagerController controllers.FileManagerController
	authMiddleware        middlewares.JWTAuthMiddleware
}

func (s FileManagerRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api/v1/fileManager").Use(s.authMiddleware.Handler())
	{
		api.POST("/uploadFile", s.FileManagerController.MakeUpload)
		api.POST("/getFile", s.FileManagerController.GetFile)
		api.POST("/removeFile", s.FileManagerController.RemoveObject)
	}
}

func NewFileManagerRoutes(
	logger logger.Logger,
	handler lib.RequestHandler,
	FileManagerController controllers.FileManagerController,
	authMiddleware middlewares.JWTAuthMiddleware,
) FileManagerRoutes {
	return FileManagerRoutes{
		logger:                logger,
		handler:               handler,
		FileManagerController: FileManagerController,
		authMiddleware:        authMiddleware,
	}
}
