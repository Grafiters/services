package pelaporan

import (
	"riskmanagement/lib"
	models "riskmanagement/models/pelaporan"
	service "riskmanagement/services/pelaporan"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type PelaporanController struct {
	logger  logger.Logger
	service service.PelaporanServicesDefinition
}

func NewPelaporanController(
	logger logger.Logger,
	PelaporanService service.PelaporanServicesDefinition,
) PelaporanController {
	return PelaporanController{
		logger:  logger,
		service: PelaporanService,
	}
}

func (p PelaporanController) GetDraftList(c *gin.Context) {
	data := models.DraftListRequest{}

	if err := c.Bind(&data); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, pagination, err := p.service.GetDraftList(data)
	if err != nil {
		lib.ReturnToJson(c, 200, "400", "Error :"+err.Error(), false)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, pagination)
}

func (p PelaporanController) GetApprovalList(c *gin.Context) {
	data := models.DraftListRequest{}

	if err := c.Bind(&data); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, pagination, err := p.service.GetApprovalList(data)
	if err != nil {
		lib.ReturnToJson(c, 200, "400", "Error :"+err.Error(), false)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, pagination)
}

func (p PelaporanController) GetDraftDetail(c *gin.Context) {
	requests := models.SuratRequestOne{}

	if err := c.Bind(&requests); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := p.service.GetDraftDetail(requests.ID)
	if err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Error: "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", result)
}

func (p PelaporanController) TambahDraftSurat(c *gin.Context) {
	data := models.DraftSuratRequest{}

	if err := c.Bind(&data); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := p.service.Generate(data)
	if err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", err.Error(), false)
		return
	}

	if !result {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Error apa: "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil menambahkan", result)
}

func (p PelaporanController) Approve(c *gin.Context) {
	data := models.ApprovalRequest{}
	if err := c.Bind(&data); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := p.service.Approve(data)
	if err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Error: "+err.Error(), false)
		return
	}

	if !result {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil Approve", result)
}

func (p PelaporanController) Reject(c *gin.Context) {
	data := models.PenolakanCatatan{}
	if err := c.Bind(&data); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := p.service.Reject(data)
	if err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Error : "+err.Error(), false)
		return
	}

	if !result {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Gagal : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil Ditolak", result)
}

func (p PelaporanController) Delete(c *gin.Context) {
	data := models.ApprovalRequest{}
	if err := c.Bind(&data); err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, err := p.service.Delete(data)
	if err != nil {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Error: "+err.Error(), false)
		return
	}

	if !result {
		p.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Berhasil Dihapus ", result)
}
