package tasklists

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	services "riskmanagement/services/tasklists"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gitlab.com/golang-package-library/logger"
)

type TasklistsController struct {
	logger  logger.Logger
	service services.TasklistsDefinition
}

func NewTasklistsController(
	TasklistsService services.TasklistsDefinition,
	logger logger.Logger,
) TasklistsController {
	return TasklistsController{
		service: TasklistsService,
		logger:  logger,
	}
}

// @Summary Get All Tasklist
// @Tags Tasklist (Old)
// @Param paginate body models.Paginate true "Tasklist ID"
// @Router /api/v1/tasklists/getAll [post]
func (tasklists TasklistsController) GetAll(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, pagination, err := tasklists.service.GetAll(requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

// @Summary   Get One Tasklist By Id
// @Tags      Tasklist (Old)
// @Param     id   body models.GetTaskByID  true  "Tasklist ID"
// @Router    /api/v1/tasklists/getTaskByID [post]
// @Security BearerAuth
func (tasklists TasklistsController) GetTaskByID(c *gin.Context) {
	requests := models.GetTaskByID{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, pagination, err := tasklists.service.GetTaskByID(requests.ID)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

// @Summary   Get All Officer
// @Tags      Tasklist (Old)
// @Param     paginate   body  models.Paginate  true  "Tasklist ID"
// @Router    /api/v1/tasklists/getAllOfficer [post]
// @Security BearerAuth
func (tasklists TasklistsController) GetAllOfficer(c *gin.Context) {
	requests := models.Paginate{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		fmt.Println(&requests)
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := tasklists.service.GetAllOfficer(requests)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

// @Summary   Get Data Anomaly
// @Tags      Tasklist (Old)
// @Param     TasklistID   body  models.TasklistDataAnomaliRequest  true  "Tasklist ID"
// @Router    /api/v1/tasklists/getDataAnomali [post]
// @Security BearerAuth
func (tasklists TasklistsController) GetDataAnomali(c *gin.Context) {
	requests := models.TasklistDataAnomaliRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		fmt.Println(&requests)
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := tasklists.service.GetDataAnomali(requests)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Get Tasklist Filter
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsFilterRequest  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/filter [post]
// @Security BearerAuth
func (tasklists TasklistsController) Filter(c *gin.Context) {
	requests := models.TasklistsFilterRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		fmt.Println(&requests)
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := tasklists.service.Filter(requests)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

// @Summary   Get Tasklist Filter By ID
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsFilterRequest  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/filterByID [post]
// @Security BearerAuth
func (tasklists TasklistsController) FilterByID(c *gin.Context) {
	// requests := models.TasklistsFilterByIDRequest{}
	request := models.TasklistsFilterRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(&request)
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := tasklists.service.FilterByID(request)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

// @Summary   Get Tasklist Filter By Officer
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsFilterOfficerRequest  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/filterOfficer [post]
// @Security BearerAuth
func (tasklists TasklistsController) FilterOfficer(c *gin.Context) {
	requests := models.TasklistsFilterOfficerRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, pagination, err := tasklists.service.FilterOfficer(requests)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if pagination.Total == 0 {
		lib.ReturnToJson(c, 200, "404", "Data Kosong", data)
		return
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery Data Berhasil", data, pagination)
}

// @Summary   Check Availability
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsCheckRequest  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/checkAvailability [post]
func (tasklists TasklistsController) CheckAvailability(c *gin.Context) {
	requests := models.TasklistsCheckRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := tasklists.service.CheckAvailability(requests)

	if err != nil {
		tasklists.logger.Zap.Error(err)
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

// @Summary   Store Tasklist
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsStoreRequest  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/store [post]
// @Security BearerAuth
func (tasklists TasklistsController) Store(c *gin.Context) {
	data := models.TasklistsStoreRequest{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	status, err := tasklists.service.Store(data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", err.Error(), false)
		return
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

// @Summary   Update Tasklist
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsUpdateRequest  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/update [post]
// @Security BearerAuth
func (tasklists TasklistsController) Update(c *gin.Context) {
	data := models.TasklistsUpdateRequest{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := tasklists.service.Update(&data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate : ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

// @Summary   Update End Date Tasklist
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsUpdateEndDateRequest  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/updateEndDate [post]
func (tasklists TasklistsController) UpdateEndDate(c *gin.Context) {
	data := models.TasklistsUpdateEndDateRequest{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	status, err := tasklists.service.UpdateEndDate(&data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate : ", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

// @Summary   Delete Tasklist
// @Tags      Tasklist (Old)
// @Param     filter   body  models.TasklistsUpdateDelete  true  "Tasklist Filter"
// @Router    /api/v1/tasklists/delete [post]
// @Security BearerAuth
func (tasklists TasklistsController) Delete(c *gin.Context) {
	data := models.TasklistsUpdateDelete{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := tasklists.service.Delete(&data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal dihapus", false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Delete data berhasil", true)
}

// @Summary   Approval Tasklist
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistsAprrovalRequest true "Tasklist Approval"
// @Router    /api/v1/tasklists/approval [post]
// @Security BearerAuth
func (tasklists TasklistsController) Approval(c *gin.Context) {
	data := models.TasklistsAprrovalRequest{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := tasklists.service.Approval(&data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate", false)
		return
	}

	if data.ApprovalStatus == "Ditolak oleh Approver" || data.ApprovalStatus == "Ditolak oleh Validator" {
		lib.ReturnToJson(c, 200, "200", "Berhasil Ditolak", true)
	} else {
		lib.ReturnToJson(c, 200, "200", "Berhasil Disetujui", true)
	}
}

// @Summary   Validation Tasklist
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistsValidation true "Tasklist Validation"
// @Router    /api/v1/task-assignment/validation [post]
// @Security BearerAuth
func (tasklists TasklistsController) Validation(c *gin.Context) {
	data := models.TasklistsValidation{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	status, err := tasklists.service.Validation(&data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
		return
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Data Gagal Diupdate", false)
		return
	}

	if data.ValidationStatus == "Ditolak oleh Approver" || data.ValidationStatus == "Ditolak oleh Validator" {
		lib.ReturnToJson(c, 200, "200", "Berhasil Ditolak", true)
	} else {
		lib.ReturnToJson(c, 200, "200", "Berhasil Disetujui", true)
	}
}

// @Summary   Count Tasklist
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistCountRequest true "Tasklist Count"
// @Router    /api/v1/tasklists/count [post]
// @Security BearerAuth
func (tasklists TasklistsController) CountTask(c *gin.Context) {
	requests := models.TasklistCountRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, err := tasklists.service.CountTask(requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Store Done Tasklist
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistsDoneHistoryRequest true "Tasklist Done History"
// @Router    /api/v1/tasklists/storeDone [post]
func (tasklists TasklistsController) StoreDone(c *gin.Context) {
	data := models.TasklistsDoneHistoryRequest{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	status, err := tasklists.service.StoreDone(data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

// @Summary   Get Done Tasklist
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistsDoneHistoryCheckRequest true "Tasklist Done History"
// @Router    /api/v1/tasklists/getDone [post]
func (tasklists TasklistsController) GetDone(c *gin.Context) {
	requests := models.TasklistsDoneHistoryCheckRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, status, err := tasklists.service.GetDone(requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if status == false {
		lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", "Data kosong")
	} else {
		lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
	}
}

// @Summary   Lead Autocomplete Val
// @Tags      Tasklist (Old)
// @Param     tasklist body models.LeadAutocompleteRequest true "Lead Autocomplete"
// @Router    /api/v1/tasklists/leadAutocompleteVal [post]
func (tasklists TasklistsController) LeadAutocompleteVal(c *gin.Context) {
	requests := models.LeadAutocompleteRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, err := tasklists.service.LeadAutocompleteVal(requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Lead Autocomplete Apr
// @Tags      Tasklist (Old)
// @Param     tasklist body models.LeadAutocompleteRequest true "Lead Autocomplete"
// @Router    /api/v1/tasklists/leadAutocompleteApr [post]
func (tasklists TasklistsController) LeadAutocompleteApr(c *gin.Context) {
	requests := models.LeadAutocompleteRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, err := tasklists.service.LeadAutocompleteApr(requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   User Region
// @Tags      Tasklist (Old)
// @Param     tasklist body models.UserRegionRequest true "User Region"
// @Router    /api/v1/tasklists/userRegion [post]
func (tasklists TasklistsController) UserRegion(c *gin.Context) {
	requests := models.UserRegionRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, err := tasklists.service.UserRegion(requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Get One Tasklist By Id
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistRequestOne true "Tasklist ID"
// @Router    /api/v1/tasklists/getOne [post]
func (tasklists TasklistsController) GetOne(c *gin.Context) {
	requests := models.TasklistRequestOne{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, status, err := tasklists.service.GetOne(requests.ID)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Get Data Verifikasi
// @Tags      Tasklist (Old)
// @Param     tasklist body models.DataVerifikasiRequest true "Data Verifikasi"
// @Router    /api/v1/tasklists/getDataVerifikasi [post]
func (tasklists TasklistsController) GetDataVerifikasi(c *gin.Context) {
	requests := models.DataVerifikasiRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, err := tasklists.service.GetDataVerifikasi(requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Get Lampiran Indicator
// @Tags      Tasklist (Old)
// @Param     tasklist body models.LampiranIndikatorCheck true "Lampiran Indicator"
// @Router    /api/v1/tasklists/getLampiranIndicator [post]
func (tasklists TasklistsController) GetLampiranIndikator(c *gin.Context) {
	requests := &models.LampiranIndikatorCheck{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, status, err := tasklists.service.GetLampiranIndicator(requests)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	if !status {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error status", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Download Lampiran Indicator Template
// @Tags      Tasklist (Old)
// @Param     tasklist body models.LampiranIndikatorCheck true "Lampiran Indicator"
// @Router    /api/v1/tasklists/downloadLampiranIndicatorTemplate [post]
func (tasklists TasklistsController) DownloadLampiranIndikatorTemplate(c *gin.Context) {
	requests := &models.LampiranIndikatorCheck{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	responses, err := tasklists.service.DownloadLampiranIndikatorTemplate(requests)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	file := excelize.NewFile()

	sheet1Name := "Sheet One"
	file.SetSheetName(file.GetSheetName(1), sheet1Name)

	col := 0

	for _, response := range responses.Header {
		col += 1

		// colString := strconv.Itoa(col)
		// fmt.Println("isi intToChar(i): ", intToChar(col))
		file.SetCellValue("Sheet1", intToChar(col)+"1", response)
	}

	var b bytes.Buffer
	if err := file.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

func intToChar(n int) string {
	var result string
	for n > 0 {
		remainder := (n - 1) % 26
		result = string('A'+remainder) + result
		n = (n - 1) / 26
	}
	return result
}

// @Summary   Insert Tasklist Rejected
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistsRejected true "Tasklist Rejected"
// @Router    /api/v1/tasklists/insertTasklistRejected [post]
func (tasklists TasklistsController) InsertTasklistRejected(c *gin.Context) {
	requests := &models.TasklistsRejected{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	err := tasklists.service.InsertTasklistRejected(requests)

	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", true)
}

func (tasklists TasklistsController) GetTasklistData(c *gin.Context) {
	// requests := &models.TasklistCheckRequest{}

	// fmt.Println("isi request:", requests)

	// if err := c.ShouldBindJSON(&requests); err != nil {
	// 	tasklists.logger.Zap.Error()
	// 	lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
	// 	return
	// }
	// data, err := tasklists.service.GetTasklist(requests)

	// if err != nil {
	// 	tasklists.logger.Zap.Error()
	// }

	// lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Insert StoreDaily
// @Tags      Tasklist (Old)
// @Param     tasklist body models.TasklistDailyStore true "Tasklist Daily Store"
// @Router    /api/v1/tasklists/storeDaily [post]
func (tasklists TasklistsController) StoreDaily(c *gin.Context) {
	data := models.TasklistDailyStore{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	_, err := tasklists.service.StoreDaily(&data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", err.Error(), false)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Input data berhasil", true)
}

// @Summary   Update Daily
// @Tags      Tasklist (Old)
// @Param     tasklist body models.ProgresUpdateRequest true "Tasklist Daily Store"
// @Router    /api/v1/tasklists/updateDaily [post]
func (tasklists TasklistsController) UpdateDaily(c *gin.Context) {
	data := models.ProgresUpdateRequest{}

	if err := c.Bind(&data); err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), "")
		return
	}

	_, err := tasklists.service.UpdateDaily(&data)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error : ", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Update data berhasil", true)
}

// @Summary   Get Anomali Header
// @Tags      Tasklist (Old)
// @Param     tasklist body models.AnomaliHeader true "Tasklist Anomali Header"
// @Router    /api/v1/tasklists/anomaliHeader [post]
func (tasklists TasklistsController) GetAnomaliHeader(c *gin.Context) {
	requests := models.AnomaliHeader{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, err := tasklists.service.GetAnomaliHeader(&requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Get Anomali Value
// @Tags      Tasklist (Old)
// @Param     tasklist body models.AnomaliValue true "Tasklist Anomali Value"
// @Router    /api/v1/tasklists/anomaliValue [post]
func (tasklists TasklistsController) GetAnomaliValue(c *gin.Context) {
	requests := models.AnomaliValue{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := tasklists.service.GetAnomaliValue(&requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// @Summary   Get First Lampiran
// @Tags      Tasklist (Old)
// @Param     tasklist body models.GetFirstLampiranRequest true "Tasklist Get First Lampiran"
// @Router    /api/v1/tasklists/getFirstLampiran [post]
func (tasklists TasklistsController) GetFirstLampiran(c *gin.Context) {
	requests := models.GetFirstLampiranRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		tasklists.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	data, err := tasklists.service.GetFirstLampiran(&requests)
	if err != nil {
		tasklists.logger.Zap.Error()
	}

	if err == sql.ErrNoRows {
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}
