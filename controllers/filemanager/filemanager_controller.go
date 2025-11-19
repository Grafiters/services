package filemanager

import (
	"riskmanagement/lib"
	models "riskmanagement/models/filemanager"
	services "riskmanagement/services/filemanager"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"

	minio "gitlab.com/golang-package-library/minio"
)

type FileManagerController struct {
	minio   minio.Minio
	logger  logger.Logger
	service services.FileManagerDefinition
}

func NewFileManagerController(
	FileManagerService services.FileManagerDefinition,
	logger logger.Logger,
	minio minio.Minio,
) FileManagerController {
	return FileManagerController{
		minio:   minio,
		logger:  logger,
		service: FileManagerService,
	}
}

func (filemanager FileManagerController) MakeUpload(c *gin.Context) {
	request := models.FileManagerRequest{}
	file, err := c.FormFile("file")
	subdir := c.PostForm("subdir")
	request.File = file
	request.Subdir = subdir

	if err != nil {
		filemanager.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), nil)
		return
	}

	datas, err := filemanager.service.MakeUpload(request)
	if err != nil {
		filemanager.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), datas)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Upload Berhasil", datas)
}

func (filemanager FileManagerController) GetFile(c *gin.Context) {
	request := models.FileManagerRequest{}
	if err := c.Bind(&request); err != nil {
		filemanager.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
		return
	}

	datas, err := filemanager.service.GetFile(request)
	if err != nil {
		filemanager.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : "+err.Error(), datas)
		// return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (filemanager FileManagerController) RemoveObject(c *gin.Context) {
	request := models.FileManagerRequest{}

	if err := c.Bind(&request); err != nil {
		filemanager.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), nil)
	}

	datas, err := filemanager.service.RemoveObject(request)
	if err != nil {
		filemanager.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error :"+err.Error(), datas)
	}

	if !datas {
		lib.ReturnToJson(c, 200, "200", "Remove Gagal :"+err.Error(), datas)
	}

	lib.ReturnToJson(c, 200, "200", "Remove Success", datas)
}
