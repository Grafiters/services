package controllers

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/managementuser"
	services "riskmanagement/services/managementuser"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type ManagementUserController struct {
	logger  logger.Logger
	service services.ManagementUserService
}

func NewManagementUserController(
	MUService services.ManagementUserService,
	logger logger.Logger,
) ManagementUserController {
	return ManagementUserController{
		logger:  logger,
		service: MUService,
	}
}

func (mu ManagementUserController) GetAll(c *gin.Context) {
	datas, err := mu.service.GetAll()
	if err != nil {
		mu.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (mu ManagementUserController) GetAllWithPaginate(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.Bind(&requests); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	datas, pagination, err := mu.service.GetAllWithPaginate(requests)
	if err != nil {
		mu.logger.Zap.Error(err)
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", nil)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", datas, pagination)
}

func (mu ManagementUserController) GetOne(c *gin.Context) {
	request := models.ManagementUserRequest{}

	if err := c.Bind(&request); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := mu.service.GetOne(request.ID)
	if err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (mu ManagementUserController) Store(c *gin.Context) {
	data := models.ManagementUserRequest{}

	if err := c.Bind(&data); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	response, err := mu.service.Store(&data)
	if err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !response {
		lib.ReturnToJson(c, 200, "400", "User sudah terdaftar", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Input Data Berhasil", true)
}

func (mu ManagementUserController) Update(c *gin.Context) {
	data := models.ManagementUserRequest{}

	if err := c.Bind(&data); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}
	fmt.Println(data)
	response, err := mu.service.Update(&data)
	if err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	if !response {
		lib.ReturnToJson(c, 200, "400", "User sudah terdaftar", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update Data Berhasil", true)
}

func (mu ManagementUserController) Delete(c *gin.Context) {
	request := models.ManagementUserRequest{}

	if err := c.Bind(&request); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	if err := mu.service.Delete(request.ID); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", true)
}

func (mu ManagementUserController) GetMappingMenu(c *gin.Context) {
	request := models.ManagementUserRequest{}

	if err := c.Bind(&request); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	data, err := mu.service.GetMappingMenu(request.ID)
	if err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (mu ManagementUserController) MappingMenu(c *gin.Context) {
	data := models.MappingMenuRequest{}

	if err := c.Bind(&data); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := mu.service.MappingMenu(data)
	if err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Mapping data berhasil", true)
}

func (mu ManagementUserController) GetAllMenu(c *gin.Context) {
	datas, err := mu.service.GetAllMenu()
	if err != nil {
		mu.logger.Zap.Error(err)
	}

	fmt.Print("data => ", datas)
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (mu ManagementUserController) DeleteMapControl(c *gin.Context) {
	data := models.MapMenu{}

	if err := c.Bind(&data); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := mu.service.DeleteMappingMenu(&data)
	if err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

// func (mu ManagementUserController) GetMenu(c *gin.Context) {
// 	request := models.MenuRequest{}

// 	if err := c.Bind(&request); err != nil {
// 		mu.logger.Zap.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error bind JSON": err.Error(),
// 		})
// 		return
// 	}
// 	menus, err := mu.service.GetMenu(request)
// 	if err != nil {
// 		mu.logger.Zap.Error(err)
// 		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
// 		return
// 	}

// 	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", menus)
// }

func (mu ManagementUserController) GetTreeMenu(c *gin.Context) {
	menus, err := mu.service.GetTreeMenu()
	if err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", menus)
}

func (mu ManagementUserController) GetUkerKelolaan(c *gin.Context) {
	request := models.UkerKelolaanUserRequest{}

	if err := c.Bind(&request); err != nil {
		mu.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	datas, err := mu.service.GetUkerKelolaan(request)
	if err != nil {
		mu.logger.Zap.Error(err)
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (mu ManagementUserController) GetLevelUker(c *gin.Context) {
	datas, err := mu.service.GetLevelUker()
	if err != nil {
		mu.logger.Zap.Error(err)
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

func (mu ManagementUserController) GetJabatanRole(c *gin.Context) {
	datas, err := mu.service.GetJabatanRole()
	if err != nil {
		mu.logger.Zap.Error(err)
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", datas)
}

func (mu ManagementUserController) GetAdditionalMenu(c *gin.Context) {
	data, err := mu.service.GetAdditionalMenu()

	if err != nil {
		mu.logger.Zap.Error(err)
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}
