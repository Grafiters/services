package menu

import (
	"database/sql"
	"riskmanagement/lib"

	menuModels "riskmanagement/models/menu"
	service "riskmanagement/services/menu"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type MenuController struct {
	logger  logger.Logger
	service service.MenuServiceDefinition
}

func NewMenuController(logger logger.Logger, service service.MenuServiceDefinition) MenuController {
	return MenuController{
		logger:  logger,
		service: service,
	}
}

func (m MenuController) GetMenuTree(c *gin.Context) {
	request := menuModels.MenuRequest{}

	if err := c.Bind(&request); err != nil {
		m.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	menus, err := m.service.GetMenuTree(request)
	if err != nil {
		m.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", menus)
}

func (m MenuController) GetKuisioner(c *gin.Context) {
	request := menuModels.RequestKuisioner{}

	if err := c.Bind(&request); err != nil {
		m.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	menus, err := m.service.GetKuisioner(request)
	if err != nil {
		m.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", menus.Data, menus.Pagination)
}

func (menu MenuController) GetAll(c *gin.Context) {
	data, err := menu.service.GetAll()
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (menu MenuController) SubMenuCheck(c *gin.Context) {
	request := menuModels.MenuQnaRequest{}

	if err := c.Bind(&request); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	data, err := menu.service.SubMenuCheck(request)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (menu MenuController) GetAllMstMenu(c *gin.Context) {
	data, err := menu.service.GetAllMstMenu()
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (menu MenuController) DeleteMenuRRM(c *gin.Context) {
	requests := menuModels.MstMenuRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := menu.service.DeleteMenuRRM(requests)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (menu MenuController) StoreMstRRM(c *gin.Context) {
	requests := menuModels.MenuRoleRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := menu.service.StoreMstRRM(requests.MstMenu)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

func (menu MenuController) DeleteRole(c *gin.Context) {
	requests := menuModels.MstMenuRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	err := menu.service.DeleteRole(requests)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", true)
}

func (menu MenuController) StoreRoleRRM(c *gin.Context) {
	requests := menuModels.MenuRoleRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	err := menu.service.StoreRoleRRM(requests.Role)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", true)
}

func (menu MenuController) SetStatus(c *gin.Context) {
	requests := menuModels.MstMenu{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	response, err := menu.service.SetStatus(requests)
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", response)
}

func (menu MenuController) GetLastID(c *gin.Context) {
	id_menu, err := menu.service.GetLastID()
	if err != nil {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", id_menu)
}
