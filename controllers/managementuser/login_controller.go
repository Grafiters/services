package controllers

import (
	"riskmanagement/lib"
	loginModel "riskmanagement/models/pgsuser"

	// services "riskmanagement/services/managementuser"

	"github.com/gin-gonic/gin"
)

// func NewLoginController(
// 	Service services.ManagementUserDefinition,
// 	logger logger.Logger,
// ) ManagementUserController {
// 	return ManagementUserController{
// 		logger:  logger,
// 		service: Service,
// 	}
// }

func (mu ManagementUserController) Login(c *gin.Context) {
	request := loginModel.LoginRequest{}
	if err := c.Bind(&request); err != nil {
		// singin.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}
	// fmt.Println(request)
	data, err := mu.service.Login(request)
	if err != nil {
		// singin.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// fmt.Println("catak ==>", data)

	// if login == nil {
	// 	lib.ReturnToJson(c, 200, "404", "User Cannot Access", nil)
	// 	return
	// }

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (mu ManagementUserController) LoginBrillianApps(c *gin.Context) {
	requests := loginModel.TokenRequest{}

	if err := c.Bind(&requests); err != nil {
		// singin.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	data, err := mu.service.LoginBrillianApps(requests)
	
	if err != nil {
		// singin.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal server error: "+err.Error(), "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}

func (mu ManagementUserController) ValidateToken (c *gin.Context) {
	requests := loginModel.LoginByToken{}

	if err := c.Bind(&requests); err != nil {
		// singin.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	data, err := mu.service.ValidateToken(requests)
	if err != nil {
		// singin.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquiry data berhasil", data)
}
