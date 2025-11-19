package taskassignment

import (
	"net/http"
	"riskmanagement/lib"
	models "riskmanagement/models/taskassignment"
	services "riskmanagement/services/taskassignment"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

const (
	msgInputInvalid      = "Input Tidak Sesuai : "
	msgInternalError     = "Internal Error"
	msgDataEmpty         = "Data Kosong"
	msgInquerySuccessful = "Inquery Data Berhasil"
	msgInputSuccessful   = "Input data berhasil"
	msgUpdateFailed      = "Data Gagal Diupdate : "
	msgDeleteSuccessful  = "Delete data berhasil"
	msgApprovalRejected  = "Berhasil Ditolak"
	msgApprovalApproved  = "Berhasil Disetujui"
)

type TaskAssignmentsController struct {
	logger  logger.Logger
	service services.TaskAssignmentsDefinition
}

func NewTaskAssignmentsController(
	TaskAssignmentService services.TaskAssignmentsDefinition,
	logger logger.Logger,
) TaskAssignmentsController {
	return TaskAssignmentsController{
		service: TaskAssignmentService,
		logger:  logger,
	}
}

// @Summary Get Data Task
// @Tags Tasklist
// @Param taskFilter body models.TaskFilterRequest true "Task Filter"
// @Router /api/v1/task-assignment/getTaskData [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetDataTask(c *gin.Context) {
	request := models.TaskFilterRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	data, totalData, err := ts.service.GetDataTask(request)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", msgInquerySuccessful, data, totalData)
}

// @Summary Get Data Task Approval List
// @Tags Tasklist
// @Param taskFilter body models.TaskApprovalRequest true "Task Filter"
// @Router /api/v1/task-assignment/getTaskApprovalList [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetTaskApprovalList(c *gin.Context) {
	request := models.TaskApprovalRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	data, totalData, err := ts.service.GetTaskApprovalList(request)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", msgInquerySuccessful, data, totalData)
}

// @Summary Get Data Task Detail
// @Tags Tasklist
// @Param taskFilter body models.TaskDetailRequest true "Task Filter"
// @Router /api/v1/task-assignment/getTaskDataDatail [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetDataTaskDetail(c *gin.Context) {
	request := models.TaskDetailRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	data, err := ts.service.GetDataTaskDetail(request.ID)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	// if data.Approval != request.Pernr || data.Validation != request.Pernr {
	// 	lib.ReturnToJson(c, 200, "403", "Anda tidak memiliki akses untuk melihat data ini", nil)
	// 	return
	// }

	lib.ReturnToJson(c, 200, "200", msgInquerySuccessful, data)
}

// @Summary Get Detail Tematik
// @Tags Tasklist
// @Param taskFilter body models.DataTematikRequest true "Task Filter"
// @Router /api/v1/task-assignment/getDetailTematik [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetDetailTematik(c *gin.Context) {
	request := models.DataTematikRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	data, totalData, err := ts.service.GetDetailTematik(request)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", msgInquerySuccessful, data, totalData)
}

// @Summary Get Rejection Notes
// @Tags Tasklist
// @Param taskFilter body models.TaskDetailRequest true "Task Filter"
// @Router /api/v1/task-assignment/getRejectionNotes [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetRejectionNotes(c *gin.Context) {
	request := models.TaskDetailRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	data, err := ts.service.GetRejectionNotes(request.ID)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", msgInquerySuccessful, data)
}

// @Summary Delete Lampiran
// @Tags Tasklist
// @Param taskFilter body models.DataTematikRequest true "Task Filter"
// @Router /api/v1/task-assignment/deleteLampiran [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) DeleteLampiran(c *gin.Context) {
	request := models.DataTematikRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	status, err := ts.service.DeleteLampiran(request)

	if err != nil && !status {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, status)
		return
	}

	lib.ReturnToJson(c, 200, "200", msgDeleteSuccessful, status)
}

// func (ts TaskAssignmentsController) StoreData(c *gin.Context) {
// 	requests := models.CreateTaskDTO{}

// 	if err := c.Bind(&requests); err != nil {
// 		ts.logger.Zap.Error(err)
// 		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
// 		return
// 	}

// 	status, err := ts.service.StoreData(requests)

// 	if err != nil {
// 		ts.logger.Zap.Error(err)
// 		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
// 		return
// 	}

// 	lib.ReturnToJson(c, 200, "200", msgInputSuccessful, status)
// }

// @Summary Check If Table Exist
// @Tags Tasklist
// @Param taskFilter body models.CheckTableRequest true "Task Filter"
// @Router /api/v1/task-assignment/checkTableExist [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) CheckTableExist(c *gin.Context) {
	requests := models.CheckTableRequest{}

	if err := c.Bind(&requests); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	response, err := ts.service.CheckTableExist(requests)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", msgInquerySuccessful, response)
}

// add by panji 02-01-2025
// @Summary Generate No Task
// @Tags Tasklist
// @Param taskFilter body models.NoTaskRequest true "Task Filter"
// @Router /api/v1/task-assignment/generateNoTask [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) GenerateNoTask(c *gin.Context) {
	requests := models.NoTaskRequest{}

	today := lib.GetTimeNow("date2")

	if err := c.Bind(&requests); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}

	NoTask, err := ts.service.GenerateNoTask(requests.Orgeh)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", nil)
		return
	}

	NoPelaporan := "TSK-" + requests.Orgeh + "-" + today + "-" + NoTask

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", NoPelaporan)
}

// @Summary   Get All Tasklist
// @Tags      Tasklist
// @Router    /api/v1/task-assignment/ [get]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetAllTasklist(c *gin.Context) {
	datas, err := ts.service.GetAllTasklist()
	if err != nil {
		ts.logger.Zap.Error(err)
		return
	}
	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", datas)
}

// @Summary   Get One Tasklist By Id
// @Tags      Tasklist
// @Param     id   path  string  true  "Tasklist ID"
// @Router    /api/v1/task-assignment/{id} [get]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetOnebyIdTasklist(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ts.logger.Zap.Error("Id Kosong")
		lib.ReturnToJson(c, 400, "400", "Id tidak boleh kosong", "")
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ts.logger.Zap.Error("Gagal merubah Id")
		lib.ReturnToJson(c, 500, "500", "Gagal merubah Id"+err.Error(), "")
		return
	}

	data, err := ts.service.GetOnebyIdTasklist(idInt)
	if err != nil {
		lib.ReturnToJson(c, 500, "500", "Gagal Inquery Data : "+err.Error(), "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", data)
}

// StoreTasklist handles the creation of a new task.
// @Summary Create Tasklist
// @Tags Tasklist
// @Accept multipart/form-data
// @Param tasklist formData models.CreateTaskDTO true "Task payload"
// @Param file formData file true "File Uploader"
// @Router /api/v1/task-assignment/store [post]
func (ts TaskAssignmentsController) StoreTasklist(c *gin.Context) {
	var request models.CreateTaskDTO

	// Bind form data to struct
	if err := c.ShouldBind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", "")
		return
	}

	// fmt.Println("Controller ===>", tasklist)

	file, err := c.FormFile("file")
	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Gagal Menerima FIle : "+err.Error(), nil)
		return
	}
	response, err := ts.service.StoreTasklist(request, file)
	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Gagal Menyimpan Tasklist : "+err.Error(), response)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data berhasil disimpan dan sedang di proses", response)
}

// @Summary Delete Tasklist
// @Tags Tasklist
// @Param     id   body models.DeleteTasklistDTO  true  "Delete Payload"
// @Router /api/v1/task-assignment/deleteTasklist [delete]
func (ts TaskAssignmentsController) DeleteTasklist(c *gin.Context) {
	request := models.DeleteTasklistDTO{}
	if err := c.ShouldBind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Input tidak sesuai", "")
		return
	}

	err := ts.service.DeleteTasklist(request)
	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 500, "500", "Internal Server Error", "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data berhasil dihapus", "")
}

func (ts TaskAssignmentsController) ValidateData(c *gin.Context) {

	data := []models.ValidateLaporanRAPDTO{}

	if err := c.Bind(&data); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Format File tidak sesuai", nil)
		// fmt.Printf("%+v", data)
		return
	}

	responses, err := ts.service.ValidateData(data)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Data Uker Tidak Valid : "+err.Error(), responses)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data Valid", responses)

	// else {
	// 	fmt.Print(data)
	// 	responses, err := ts.service.ValidateData(data)
	// 	if err != nil {
	// 		ts.logger.Zap.Error(err)
	// 		lib.ReturnToJson(c, 400, "400", "Beberapa tidak Valid : "+err.Error(), responses)
	// 		return
	// 	} else {
	// 		lib.ReturnToJson(c, 200, "200", "Inquery Data Berhasil", responses)
	// 	}
	// }
}

// @Summary Approve Tasklist
// @Tags Tasklist
// @Param approvalRequest body models.ApprovalRequest true "Approval Request"
// @Router /api/v1/task-assignment/approval [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) Approval(c *gin.Context) {
	var approvalRequest models.ApprovalRequest
	if err := c.ShouldBind(&approvalRequest); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Input tidak sesuai", "")
		return
	}

	tasklist, err := ts.service.Approval(approvalRequest)
	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Error : "+err.Error(), "")
		return
	}

	if approvalRequest.Approval == "reject" {
		lib.ReturnToJson(c, 200, "200", "Data berhasil diReject", tasklist)
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data berhasil diApprove", tasklist)
}

// @Summary Update Tasklist
// @Tags Tasklist
// @Param id path int true "Tasklist ID"
// @Param tasklist formData models.UpdateTasklistDTO true "Update Task payload"
// @Param file formData file true "File Uploader"
// @Router /api/v1/task-assignment/update/{id} [put]
func (ts TaskAssignmentsController) UpdateTasklist(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ts.logger.Zap.Error("id is empty")
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", "")
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Input tidak sesuai", "")
		return
	}
	request := models.UpdateTasklistDTO{}
	if err := c.ShouldBind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Input tidak sesuai", "")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		if err != http.ErrMissingFile {
			ts.logger.Zap.Error(err)
			lib.ReturnToJson(c, 500, "500", "Gagal menerima file: "+err.Error(), nil)
			return
		}
		file = nil // If file not provided, pass nil
	}

	tasklist, err := ts.service.UpdateTasklist(idInt, request, file)
	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 400, "400", "Gagal memperbaharui Tasklist : "+err.Error(), "")
		return
	}

	lib.ReturnToJson(c, 200, "200", "Data Berhasil diperbaharui", tasklist)
}

// @Summary MyTasklist
// @Tags Tasklist
// @Param approvalRequest body models.TaskFilterRequest true "My Task Filter"
// @Router /api/v1/task-assignment/myTasklist [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) MyTasklist(c *gin.Context) {
	request := models.TaskFilterRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	data, totalData, err := ts.service.MyTasklist(request)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", msgInquerySuccessful, data, totalData)
}

// @Summary MyTasklistDetail
// @Tags Tasklist
// @Param approvalRequest body models.TaskDetailRequest true "My Task Detail"
// @Router /api/v1/task-assignment/myTasklistDetail [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) MyTasklistDetail(c *gin.Context) {
	request := models.TaskDetailRequest{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	data, err := ts.service.MyTasklistDetail(request.ID)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	lib.ReturnToJson(c, 200, "200", msgInquerySuccessful, data)
}

// @Summary GetMyTasklistTotal
// @Tags Tasklist
// @Param approvalRequest body models.RequestMyTasklist true "My Task Total"
// @Router /api/v1/task-assignment/getMyTasklistTotal [post]
// @Security BearerAuth
func (ts TaskAssignmentsController) GetMyTasklistTotal(c *gin.Context) {
	request := models.RequestMyTasklist{}

	if err := c.Bind(&request); err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", msgInputInvalid+err.Error(), "")
		return
	}

	total, err := ts.service.GetMyTasklistTotal(request)

	if err != nil {
		ts.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", msgInternalError, nil)
		return
	}

	response := models.MyTasklist{
		Total: total,
	}

	lib.ReturnToJson(c, 200, "200", msgInquerySuccessful, response)
}
