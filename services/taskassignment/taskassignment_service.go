package taskassignment

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"riskmanagement/consumers/publisher"
	"riskmanagement/lib"
	fileModels "riskmanagement/models/filemanager"
	publisherModels "riskmanagement/models/publisher"
	models "riskmanagement/models/taskassignment"
	tasklistModels "riskmanagement/models/tasklists"
	fileRepo "riskmanagement/repository/files"
	repository "riskmanagement/repository/taskassignment"
	tasklistRepo "riskmanagement/repository/tasklists"
	filemanager "riskmanagement/services/filemanager"

	"strings"

	"gitlab.com/golang-package-library/logger"
	"gitlab.com/golang-package-library/minio"
)

type TaskAssignmentsDefinition interface {
	// StoreData(request models.CreateTaskDTO) (status bool, err error)
	CheckTableExist(request models.CheckTableRequest) (response models.TemplateResponse, err error)
	GetDataTask(request models.TaskFilterRequest) (responses []models.TaskResponses, totalData int64, err error)
	GetDataTaskDetail(id int64) (responses models.TaskResponses, err error)
	GetDetailTematik(request models.DataTematikRequest) (responses models.DataTematikResponse, totalData int64, err error)
	GetTaskApprovalList(request models.TaskApprovalRequest) (responses []models.TaskResponses, totalData int64, err error)
	GetRejectionNotes(id int64) (response models.TaskRejectionNotes, err error)
	DeleteLampiran(request models.DataTematikRequest) (responses bool, err error)
	MyTasklist(request models.TaskFilterRequest) (responses []models.MyTasklistResponse, totalData int64, err error)
	MyTasklistDetail(id int64) (responses models.MyTasklistResponse, err error)

	// Add by Panji 02-01-2025
	GenerateNoTask(orgeh string) (response string, err error)
	GetAllTasklist() (responses []models.Task, err error)
	GetOnebyIdTasklist(id int64) (responses models.Task, err error)
	StoreTasklist(request models.CreateTaskDTO, file *multipart.FileHeader) (response models.Task, err error)
	UpdateTasklist(id int64, request models.UpdateTasklistDTO, file *multipart.FileHeader) (response models.Task, err error)
	DeleteTasklist(request models.DeleteTasklistDTO) (err error)
	ValidateData(request []models.ValidateLaporanRAPDTO) (responses []models.ValidateLaporanRAPDTO, err error)
	// InsertFromFile(payload map[string]interface{}) error
	Approval(request models.ApprovalRequest) (response models.Task, err error)
	GetMyTasklistTotal(request models.RequestMyTasklist) (total int64, err error)
}

type TaskAssignmentsService struct {
	db           lib.Database
	minio        minio.Minio
	logger       logger.Logger
	repository   repository.TaskAssignmentsDefinition
	fileRepo     fileRepo.FilesDefinition
	tasklistRepo tasklistRepo.TasklistsDefinition
	filemanager  filemanager.FileManagerDefinition
	publisher    publisher.PublisherInterface
}

func NewTaskAssignmentService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	repository repository.TaskAssignmentsDefinition,
	fileRepo fileRepo.FilesDefinition,
	tasklistRepo tasklistRepo.TasklistsDefinition,
	filemanager filemanager.FileManagerDefinition,
	publisher publisher.PublisherInterface,
) TaskAssignmentsDefinition {
	return TaskAssignmentsService{
		db:           db,
		minio:        minio,
		logger:       logger,
		repository:   repository,
		fileRepo:     fileRepo,
		tasklistRepo: tasklistRepo,
		filemanager:  filemanager,
		publisher:    publisher,
	}
}

// StoreData implements TaskAssignmentsDefinition.
// func (taskassignments TaskAssignmentsService) StoreData(request models.CreateTaskDTO) (status bool, err error) {
// 	timeNow := lib.GetTimeNow("timestime")

// 	tx := taskassignments.db.DB.Begin()

// 	storedData, err := taskassignments.repository.StoreData(models.Task{
// 		NoTasklist:      request.NoTasklist,
// 		NamaTasklist:    request.NamaTasklist,
// 		RiskIndicator:   request.RiskIndicator,
// 		ActivityID:      request.ActivityID,
// 		ProductID:       request.ProductID,
// 		ProductName:     request.ProductName,
// 		RiskIssueID:     request.RiskIssueID,
// 		RiskIssue:       request.RiskIssue,
// 		RiskIndicatorID: request.RiskIndicatorID,
// 		TaskType:        request.TaskType,
// 		TaskTypeName:    request.TaskTypeName,
// 		Kegiatan:        request.Kegiatan,
// 		Period:          request.Period,
// 		SumberData:      request.SumberData,
// 		StartDate:       request.StartDate,
// 		EndDate:         request.EndDate,
// 		RAP:             request.RAP,
// 		Validation:      request.Validation,
// 		ValidationName:  request.ValidationName,
// 		Approval:        request.Approval,
// 		ApprovalName:    request.ApprovalName,
// 		ApprovalStatus:  "Minta Persetujuan Validasi",
// 		Status:          "Aktif",
// 		MakerID:         request.MakerID,
// 		CreatedAt:       timeNow,
// 		UpdatedAt:       timeNow,
// 	}, tx)

// 	if err != nil {
// 		taskassignments.logger.Zap.Error("error when store data: ", err)
// 		tx.Rollback()
// 		return false, fmt.Errorf("error when store data: %v", err)
// 	}

// 	if request.File.Filename != "" {
// 		files, err := taskassignments.fileRepo.Store(&fileModels.Files{
// 			Filename:  request.File.Filename,
// 			Path:      request.File.Path,
// 			Extension: request.File.Extension,
// 			Size:      request.File.Size,
// 			CreatedAt: &timeNow,
// 		}, tx)

// 		if err != nil {
// 			taskassignments.logger.Zap.Error("Error when store file: ", err)
// 			tx.Rollback()
// 			return false, fmt.Errorf("error when store file: %v", err)
// 		}

// 		_, err = taskassignments.repository.StoreFile(models.TaskFile{
// 			TasklistsID: storedData.ID,
// 			FilesID:     files.ID,
// 			CreatedAt:   timeNow,
// 		}, tx)

// 		if err != nil {
// 			taskassignments.logger.Zap.Error("Error when store file: ", err)
// 			tx.Rollback()
// 			return false, fmt.Errorf("error when store file: %v", err)
// 		}
// 	}

// 	_, err = taskassignments.tasklistRepo.StoreNotif(
// 		&tasklistModels.TasklistNotif{
// 			TaskID:     storedData.ID,
// 			Tanggal:    &timeNow,
// 			Keterangan: "Ada 1 tasklist baru yang perlu divalidasi",
// 			Status:     0,
// 			Jenis:      "Validasi Task",
// 			Receiver:   storedData.Validation,
// 			Uker:       "",
// 		}, tx)

// 	if err != nil {
// 		taskassignments.logger.Zap.Error("Error when store notif: ", err)
// 		tx.Rollback()
// 		return false, fmt.Errorf("error when store notif: %v", err)
// 	}

// 	tx.Commit()

// 	return true, err
// }

// CheckTableExist implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) CheckTableExist(request models.CheckTableRequest) (response models.TemplateResponse, err error) {
	response, err = taskassignments.repository.CheckTableExist(request)

	return response, err
}

//Add By Panji 02-01-2025

// GenerateNoTask implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) GenerateNoTask(orgeh string) (response string, err error) {
	response, err = taskassignments.repository.GenerateNoTask(orgeh)

	return response, err
}

// GetDataTask implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) GetDataTask(request models.TaskFilterRequest) (responses []models.TaskResponses, totalData int64, err error) {
	dataTask, totalData, err := taskassignments.repository.GetDataTask(request)

	if err != nil {
		taskassignments.logger.Zap.Error(err.Error())
		return nil, totalData, err
	}

	for _, value := range dataTask {
		responses = append(responses, models.TaskResponses{
			No:              value.No,
			ID:              value.ID,
			NoTasklist:      value.NoTasklist,
			NamaTasklist:    value.NamaTasklist,
			RiskIndicator:   value.RiskIndicator,
			ActivityID:      value.ActivityID,
			ActivityName:    value.ActivityName,
			ProductID:       value.ProductID,
			ProductName:     value.ProductName,
			RiskIssueID:     value.RiskIssueID,
			RiskIssue:       value.RiskIssue,
			RiskIndicatorID: value.RiskIndicatorID,
			TaskType:        value.TaskType,
			TaskTypeName:    value.TaskTypeName,
			Kegiatan:        value.Kegiatan,
			Period:          value.Period,
			SumberData:      value.SumberData,
			StartDate:       value.StartDate,
			EndDate:         value.EndDate,
			RAP:             value.RAP,
			Validation:      value.Validation,
			ValidationName:  value.ValidationName,
			Approval:        value.Approval,
			ApprovalName:    value.ApprovalName,
			MakerID:         value.MakerID,
			Status:          value.Status,
			ApprovalStatus:  value.ApprovalStatus,
			StatusFile:      value.StatusFile,
		})
	}

	return responses, totalData, err
}

// GetTaskApprovalList implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) GetTaskApprovalList(request models.TaskApprovalRequest) (responses []models.TaskResponses, totalData int64, err error) {
	dataTask, totalData, err := taskassignments.repository.GetTaskApprovalList(request)

	if err != nil {
		taskassignments.logger.Zap.Error(err.Error())
		return nil, totalData, err
	}

	for _, value := range dataTask {
		// branches, _ := taskassignments.repository.GetBranchList(value.ID)

		responses = append(responses, models.TaskResponses{
			No:              value.No,
			ID:              value.ID,
			NoTasklist:      value.NoTasklist,
			NamaTasklist:    value.NamaTasklist,
			RiskIndicator:   value.RiskIndicator,
			ActivityID:      value.ActivityID,
			ActivityName:    value.ActivityName,
			ProductID:       value.ProductID,
			ProductName:     value.ProductName,
			RiskIssueID:     value.RiskIssueID,
			RiskIssue:       value.RiskIssue,
			RiskIndicatorID: value.RiskIndicatorID,
			TaskType:        value.TaskType,
			TaskTypeName:    value.TaskTypeName,
			Kegiatan:        value.Kegiatan,
			Period:          value.Period,
			SumberData:      value.SumberData,
			StartDate:       value.StartDate,
			EndDate:         value.EndDate,
			RAP:             value.RAP,
			Validation:      value.Validation,
			ValidationName:  value.ValidationName,
			Approval:        value.Approval,
			ApprovalName:    value.ApprovalName,
			MakerID:         value.MakerID,
			Status:          value.Status,
			ApprovalStatus:  value.ApprovalStatus,
			StatusFile:      value.StatusFile,
			// Branch:          branches,
		})
	}

	return responses, totalData, err
}

// GetDataTaskDetail implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) GetDataTaskDetail(id int64) (responses models.TaskResponses, err error) {
	responses, err = taskassignments.repository.GetDataTaskDetail(id)
	branches, _ := taskassignments.repository.GetBranchList(responses.ID)
	responses.Branch = branches

	return responses, err
}

// GetRejectionNotes implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) GetRejectionNotes(id int64) (response models.TaskRejectionNotes, err error) {
	response, err = taskassignments.repository.GetRejectionNotes(id)

	return response, err
}

// GetDetailTematik implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) GetDetailTematik(request models.DataTematikRequest) (responses models.DataTematikResponse, totalData int64, err error) {
	responses, totalData, err = taskassignments.repository.GetDetailTematik(request)

	return responses, totalData, err
}

// DeleteLampiran implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) DeleteLampiran(request models.DataTematikRequest) (responses bool, err error) {
	responses, err = taskassignments.repository.DeleteLampiran(request)

	if err != nil {
		return false, fmt.Errorf("gagal Menghapus Lampiran RAP")
	}

	status, err := taskassignments.repository.DeleteTaskUker(request.Id)

	if err != nil {
		return status, fmt.Errorf("gagal menghapus tasklist uker")
	}

	return responses, err
}

func (taskassignments TaskAssignmentsService) GetAllTasklist() (responses []models.Task, err error) {
	fmt.Println("aa")
	return taskassignments.repository.GetAll()
}

func (taskassignments TaskAssignmentsService) GetOnebyIdTasklist(id int64) (responses models.Task, err error) {
	return taskassignments.repository.GetOneById(id)
}

func (taskassignments TaskAssignmentsService) StoreTasklist(request models.CreateTaskDTO, file *multipart.FileHeader) (response models.Task, err error) {
	timeNow := lib.GetTimeNow("timestime")
	allowedExtensions := map[string]struct{}{
		".xlsx": {},
		".xls":  {},
		".csv":  {},
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if _, valid := allowedExtensions[ext]; !valid {
		return models.Task{}, fmt.Errorf("hanya file dengan tipe .XLS, .XLSX, and .CSV yang diizinkan")
	}
	filename := fmt.Sprintf("LampiranRAP%s%s", strings.ReplaceAll(strings.ReplaceAll(lib.GetTimeNow("timestime"), " ", "_"), ":", "-"), ext)
	fmt.Println("filename =>", filename)
	// filename := "testUpload"
	savePath, err := taskassignments.filemanager.MakeUpload(fileModels.FileManagerRequest{
		File:     file,
		Filename: filename,
		Subdir:   "LampiranRAP",
	})

	if err != nil {
		return models.Task{}, fmt.Errorf("gagal menyimpan file : %w", err)
	}

	tx := taskassignments.db.DB.Begin()
	tasklist := models.Task{
		NoTasklist:      request.NoTasklist,
		NamaTasklist:    request.NamaTasklist,
		RiskIndicator:   request.RiskIndicator,
		ActivityID:      request.ActivityID,
		ProductID:       request.ProductID,
		ProductName:     request.ProductName,
		RiskIssueID:     request.RiskIssueID,
		RiskIssue:       request.RiskIssue,
		RiskIndicatorID: request.RiskIndicatorID,
		TaskType:        request.TaskType,
		TaskTypeName:    request.TaskTypeName,
		Kegiatan:        request.Kegiatan,
		Period:          request.Period,
		EnableRange:     request.EnableRange,
		SumberData:      request.SumberData,
		RangeDate:       request.RangeDate,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
		RAP:             request.RAP,
		Sample:          request.Sample,
		Validation:      request.Validation,
		ValidationName:  request.ValidationName,
		Approval:        request.Approval,
		ApprovalName:    request.ApprovalName,
		ApprovalStatus:  models.M1,
		Status:          "Aktif",
		MakerID:         request.MakerID,
		StatusFile:      "Sedang Diproses",
		CreatedAt:       timeNow,
		UpdatedAt:       timeNow,
	}

	response, err = taskassignments.repository.StoreTasklist(tasklist)
	if err != nil {
		taskassignments.logger.Zap.Error("Error Push Notif : ", err)
		tx.Rollback()
		return models.Task{}, fmt.Errorf("gagal menyimpan data Task : %w", err)
	}

	_, err = taskassignments.tasklistRepo.StoreNotif(&tasklistModels.TasklistNotif{
		TaskID:     response.ID,
		Tanggal:    &timeNow,
		Keterangan: "Ada 1 task baru untuk divalidasi",
		Status:     0,
		Jenis:      "Validasi Tasklist",
		Receiver:   request.Validation,
		Uker:       "-",
	}, tx)

	if err != nil {
		taskassignments.logger.Zap.Error("Error Push Notif : ", err)
		tx.Rollback()
		return models.Task{}, fmt.Errorf("error push notif : ", err)
	}

	tx.Commit()

	// savePath, err := taskassignments.SaveFile(file)
	// if err != nil {
	// 	return models.Task{}, err
	// }
	data := map[string]interface{}{
		"minioPath":   savePath.Path,
		"filename":    filename,
		"tasklist_id": response.ID, // Store the struct directly; we'll serialize everything in one step
	}

	fmt.Println("Data to send =>", data)

	payload := publisherModels.PublishMessageDTO{
		QueueName: "taskassignment-read-file-minio",
		Pattern:   "insert-tasklist",
		Body:      data,
	}
	err = taskassignments.publisher.PublishMessage(payload)
	if err != nil {
		taskassignments.logger.Zap.Error(err)
		return models.Task{}, fmt.Errorf("gagal mengirim message ke RabbitMQ : %w", err)
	}
	return response, nil
}

func (taskassignments TaskAssignmentsService) UpdateTasklist(id int64, request models.UpdateTasklistDTO, file *multipart.FileHeader) (response models.Task, err error) {
	timeNow := lib.GetTimeNow("timestime")
	existingTasklist, err := taskassignments.repository.GetOneById(id)
	if err != nil {
		return models.Task{}, fmt.Errorf("gagal mendapatkan tasklist : %w", err)
	}

	if existingTasklist.ApprovalStatus == models.C0 || existingTasklist.ApprovalStatus == models.S0 {
		existingTasklist.NoTasklist = request.NoTasklist
		existingTasklist.NamaTasklist = request.NamaTasklist
		existingTasklist.RiskIndicator = request.RiskIndicator
		existingTasklist.ActivityID = request.ActivityID
		existingTasklist.ProductID = request.ProductID
		existingTasklist.ProductName = request.ProductName
		existingTasklist.RiskIssueID = request.RiskIssueID
		existingTasklist.RiskIssue = request.RiskIssue
		existingTasklist.RiskIndicatorID = request.RiskIndicatorID
		existingTasklist.TaskType = request.TaskType
		existingTasklist.TaskTypeName = request.TaskTypeName
		existingTasklist.Kegiatan = request.Kegiatan
		existingTasklist.EnableRange = request.EnableRange
		existingTasklist.Period = request.Period
		existingTasklist.SumberData = request.SumberData
		existingTasklist.RangeDate = request.RangeDate
		existingTasklist.StartDate = request.StartDate
		existingTasklist.EndDate = request.EndDate
		existingTasklist.RAP = request.RAP
		existingTasklist.Sample = request.Sample
		existingTasklist.Validation = request.Validation
		existingTasklist.ValidationName = request.ValidationName
		existingTasklist.Approval = request.Approval
		existingTasklist.ApprovalName = request.ApprovalName
		existingTasklist.ApprovalStatus = models.M1
		existingTasklist.MakerID = request.MakerID

		existingTasklist.UpdatedAt = lib.GetTimeNow("timestime")
		tx := taskassignments.db.DB.Begin()

		response, err = taskassignments.repository.UpdateTasklist(existingTasklist)
		if err != nil {
			taskassignments.logger.Zap.Error("Error Update data :", err)
			tx.Rollback()
			return models.Task{}, err
		}

		_, err = taskassignments.tasklistRepo.StoreNotif(&tasklistModels.TasklistNotif{
			TaskID:     response.ID,
			Tanggal:    &timeNow,
			Keterangan: "Ada 1 task baru untuk divalidasi",
			Status:     0,
			Jenis:      "Validasi Tasklist",
			Receiver:   request.Validation,
			Uker:       "-",
		}, tx)

		if err != nil {
			taskassignments.logger.Zap.Error("Error Push Notif : ", err)
			tx.Rollback()
			return models.Task{}, fmt.Errorf("error push notif : ", err)
		}

		tx.Commit()

		if file != nil {
			allowedExtensions := map[string]struct{}{
				".xlsx": {},
				".xls":  {},
				".csv":  {},
			}
			ext := strings.ToLower(filepath.Ext(file.Filename))
			if _, valid := allowedExtensions[ext]; !valid {
				return models.Task{}, fmt.Errorf("hanya file dengan tipe .XLS, .XLSX, and .CSV yang diizinkan")
			}

			filename := fmt.Sprintf("LampiranRAP%s%s", strings.ReplaceAll(strings.ReplaceAll(lib.GetTimeNow("timestime"), " ", "_"), ":", "-"), ext)
			fmt.Println("filename =>", filename)
			// filename := "testUpload"
			savePath, err := taskassignments.filemanager.MakeUpload(fileModels.FileManagerRequest{
				File:     file,
				Filename: filename,
				Subdir:   "LampiranRAP",
			})

			if err != nil {
				return models.Task{}, fmt.Errorf("gagal menyimpan file : %w", err)
			}

			fmt.Println("savePath =>", savePath.Path)

			data := map[string]interface{}{
				"minioPath":   savePath.Path,
				"filename":    filename,
				"tasklist_id": response.ID, // Store the struct directly; we'll serialize everything in one step
			}

			payload := publisherModels.PublishMessageDTO{
				QueueName: "taskassignment-read-file-minio",
				Pattern:   "insert-tasklist",
				Body:      data,
			}
			err = taskassignments.publisher.PublishMessage(payload)
			if err != nil {
				taskassignments.logger.Zap.Error(err)
				return models.Task{}, fmt.Errorf("gagal mengirim message ke RabbitMQ")
			}
		}

	} else {
		return models.Task{}, fmt.Errorf("tasklist tidak bisa diubah karena status approval %v", existingTasklist.ApprovalStatus)
	}

	return response, nil
}

func (taskassignments TaskAssignmentsService) DeleteTasklist(request models.DeleteTasklistDTO) (error error) {
	tasklist, err := taskassignments.repository.GetOneById(request.ID)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan tasklist : %w", err)
	}

	if tasklist.MakerID != request.Perner {
		return fmt.Errorf("anda tidak memiliki akses untuk menghapus Tasklist")
	}

	deleteLampiranDTO := models.DataTematikRequest{
		Id:            tasklist.ID,
		RiskEvent:     tasklist.RiskIssueID,
		RiskIndicator: tasklist.RiskIndicatorID,
	}

	tableColumn, err := taskassignments.repository.CheckTableExist(models.CheckTableRequest{
		RiskIssue:     tasklist.RiskIssueID,
		RiskIndicator: tasklist.RiskIndicatorID,
	})

	if err != nil {
		return fmt.Errorf("gagal mendapatkan data table : %w", err)
	}

	if tableColumn.Status == "ada" {
		response, err := taskassignments.repository.DeleteLampiran(deleteLampiranDTO)
		if err != nil {
			return err
		}

		if response {
			err = taskassignments.repository.DeleteTasklistUker(tasklist.ID)
			if err != nil {
				return fmt.Errorf("gagal menghapus tasklist uker: %w", err)
			}
		} else {
			return fmt.Errorf("gagal menghapus lampiran RAP")
		}
	}

	if tasklist.ApprovalStatus == models.C0 || tasklist.ApprovalStatus == models.S0 {
		return taskassignments.repository.DeleteTasklist(request.ID)
	} else {
		return fmt.Errorf("tasklist tidak bisa dihapus karena status approval %v", tasklist.ApprovalStatus)
	}
}

func (taskassignments TaskAssignmentsService) ValidateData(data []models.ValidateLaporanRAPDTO) (responses []models.ValidateLaporanRAPDTO, err error) {
	var invalidDwh []models.ValidateLaporanRAPDTO
	for _, item := range data {
		valid := taskassignments.repository.ValidateData(item)
		if !valid {
			invalidDwh = append(invalidDwh, item)
		}
	}
	if len(invalidDwh) > 0 {
		return invalidDwh, fmt.Errorf("jumlah data invalid: %d", len(invalidDwh))
	}

	return data, nil
}

// func (taskassignments TaskAssignmentsService) InsertFromFile(payload map[string]interface{}) error {
// 	tasklist, err := taskassignments.repository.GetOneById(int64(payload["tasklist_id"].(float64)))
// 	if err != nil {
// 		os.Remove(payload["filepath"].(string))
// 		return fmt.Errorf("gagal mendapatkan tasklist: %w", err)
// 	}

// 	file, err := excelize.OpenFile(payload["filepath"].(string))
// 	if err != nil {
// 		tasklist.StatusFile = "Gagal Membuka File : " + err.Error()
// 		taskassignments.repository.UpdateTasklist(tasklist)
// 		os.Remove(payload["filepath"].(string))
// 		return fmt.Errorf("sistem gagal membuka file: %w", err)
// 	}

// 	rows := file.GetRows("Sheet1")
// 	if len(rows) < 1 {
// 		tasklist.StatusFile = "Gagal : sheet harus bernama Sheet1 (Default)"
// 		taskassignments.repository.UpdateTasklist(tasklist)
// 		os.Remove(payload["filepath"].(string))
// 		return fmt.Errorf("sheet harus bernama Sheet1 (Default)")
// 	}

// 	re := regexp.MustCompile(`[^a-zA-Z0-9_]`)
// 	for i := range rows {
// 		for j := range rows[i] {
// 			// Sanitize each string using the regex
// 			rows[i][j] = re.ReplaceAllString(rows[i][j], "")
// 		}
// 	}
// 	headers := rows[0]
// 	rows = rows[1:]
// 	duplicateCounts := make(map[string]int)
// 	distinctData := make([]models.TasklistUker, 0, len(rows))

// 	for _, row := range rows {
// 		key := fmt.Sprintf("%s|%s|%s|%s|%s|%s", row[0], row[1], row[2], row[3], row[4], row[5])
// 		duplicateCounts[key]++

// 		if duplicateCounts[key] == 1 {
// 			distinctData = append(distinctData, models.TasklistUker{
// 				TasklistId: tasklist.ID,
// 				REGION:     row[0], RGDESC: row[1], MAINBR: row[2],
// 				MBDESC: row[3], BRANCH: row[4], BRDESC: row[5],
// 				JumlahNominatif: duplicateCounts[key],
// 			})
// 		} else {
// 			for i := range distinctData {
// 				if distinctData[i].REGION == row[0] && distinctData[i].RGDESC == row[1] &&
// 					distinctData[i].MAINBR == row[2] && distinctData[i].MBDESC == row[3] &&
// 					distinctData[i].BRANCH == row[4] && distinctData[i].BRDESC == row[5] {

// 					distinctData[i].JumlahNominatif = duplicateCounts[key]
// 					break
// 				}
// 			}
// 		}
// 	}

// 	err = taskassignments.repository.BatchStoreTasklistUker(distinctData)
// 	if err != nil {
// 		tasklist.StatusFile = "Gagal input file ke DB : " + err.Error()
// 		taskassignments.repository.UpdateTasklist(tasklist)
// 		os.Remove(payload["filepath"].(string))
// 		return err
// 	}

// 	tableName := "lampiran_rap_" + strconv.FormatInt(tasklist.RiskIssueID, 10) + "_" + strconv.FormatInt(tasklist.RiskIndicatorID, 10)
// 	isExist, err := taskassignments.repository.ValidatFileToDB(tableName, headers)
// 	if err != nil {
// 		tasklist.StatusFile = "Gagal validasi file : " + err.Error()
// 		taskassignments.repository.UpdateTasklist(tasklist)
// 		os.Remove(payload["filepath"].(string))
// 		return err
// 	} else {
// 		if isExist {
// 			err := taskassignments.repository.BatchStoreLampiranRAP(tableName, rows, tasklist.ID)
// 			if err != nil {
// 				tasklist.StatusFile = "Gagal input file ke DB : " + err.Error()
// 				taskassignments.repository.UpdateTasklist(tasklist)
// 				os.Remove(payload["filepath"].(string))
// 				return err
// 			}
// 		} else if !isExist {
// 			err := taskassignments.repository.CreateTableLampiranRAP(tableName, headers)
// 			if err != nil {
// 				tasklist.StatusFile = "Gagal Membuat Table file : " + err.Error()
// 				taskassignments.repository.UpdateTasklist(tasklist)
// 				os.Remove(payload["filepath"].(string))
// 				return err
// 			}
// 			err = taskassignments.repository.BatchStoreLampiranRAP(tableName, rows, tasklist.ID)
// 			if err != nil {
// 				tasklist.StatusFile = "Gagal input file ke DB : " + err.Error()
// 				taskassignments.repository.UpdateTasklist(tasklist)
// 				os.Remove(payload["filepath"].(string))
// 				return err
// 			}
// 		}
// 	}
// 	tasklist.StatusFile = "Selesai"
// 	taskassignments.repository.UpdateTasklist(tasklist)
// 	os.Remove(payload["filepath"].(string))
// 	return nil
// }

func (taskassignments TaskAssignmentsService) PublisherApproval(request models.ApprovalRequest) (err error) {
	// if request.Approval == "Ditolak oleh Approver" || request.Approval == "Ditolak oleh Validator" {
	// 	err := taskassignments.repository.StoreTasklistRejected(models.TasklistRejected{TasklsitID: request.TasklistID, Notes: request.Notes, Status: "Not Done"})
	// 	if err != nil {
	// 		return fmt.Errorf("gagal menyimpan data Tasklists yang di tolak: %v", err)
	// 	}
	// } else if request.Approval == "Disetujui" {
	// 	tasklist, err := taskassignments.repository.GetOneById(request.TasklistID)
	// 	if err != nil {
	// 		return fmt.Errorf("gagal mendapatkan data Tasklist: %v", err)
	// 	}
	// 	if tasklist.Period != "Daily" {
	// 		_, err := taskassignments.repository.GetDataBRC(request.Branch)
	// 		if err != nil {
	// 			return fmt.Errorf("gagal mendapatkan data BRC: %v", err)
	// 		}
	// 	}
	// }
	return nil
}

func (taskassignments TaskAssignmentsService) Approval(request models.ApprovalRequest) (models.Task, error) {
	timeNow := lib.GetTimeNow("timestime")
	tasklist, err := taskassignments.repository.GetOneById(request.ID)
	if err != nil {
		return models.Task{}, fmt.Errorf("gagal mendapatkan data Tasklist: %v", err)
	}

	if (request.Perner != tasklist.Validation && tasklist.ApprovalStatus == models.M1) ||
		(request.Perner != tasklist.Approval && tasklist.ApprovalStatus == models.C1) {
		return models.Task{}, fmt.Errorf("anda tidak memiliki akses untuk approval")
	}

	newStatus := models.ApprovalEnum("")
	switch request.Approval {
	case "approve":
		if tasklist.ApprovalStatus == models.M1 {
			newStatus = models.C1
		} else if tasklist.ApprovalStatus == models.C1 {
			newStatus = models.S1
		} else {
			return models.Task{}, fmt.Errorf("task sudah disetujui Approver")
		}
	case "reject":
		err = taskassignments.repository.StoreTasklistRejected(models.TasklistRejected{TasklistID: request.ID, Notes: request.Notes})
		if err != nil {
			return models.Task{}, fmt.Errorf("gagal menyimpan ke daftar tasklist ditolak: %v", err)
		}
		if tasklist.ApprovalStatus == models.M1 {
			newStatus = models.C0
		} else if tasklist.ApprovalStatus == models.C1 {
			newStatus = models.S0
		} else {
			return models.Task{}, fmt.Errorf("task sudah ditolak")
		}
	default:
		return models.Task{}, fmt.Errorf("status approval tidak sesuai")
	}

	tasklist.ApprovalStatus = newStatus
	tx := taskassignments.db.DB.Begin()

	updatedTask, err := taskassignments.repository.UpdateTasklist(tasklist)
	if err != nil {
		tx.Rollback()
		return models.Task{}, fmt.Errorf("gagal update data Tasklist: %v", err)
	}

	if newStatus == models.C1 {
		_, err = taskassignments.tasklistRepo.StoreNotif(&tasklistModels.TasklistNotif{
			TaskID:     updatedTask.ID,
			Tanggal:    &timeNow,
			Keterangan: "Ada 1 task baru untuk diapprove",
			Status:     0,
			Jenis:      "Validasi Tasklist",
			Receiver:   updatedTask.Approval,
			Uker:       "-",
		}, tx)

		if err != nil {
			taskassignments.logger.Zap.Error("Error Push Notif : ", err)
			tx.Rollback()
			return models.Task{}, fmt.Errorf("error push notif : ", err)
		}
	}

	tx.Commit()

	if updatedTask.ApprovalStatus == models.S1 {
		if updatedTask.Period != "Daily" {
			brcList, err := taskassignments.repository.GetDataBRC(request.Branch, tasklist.ID)
			if err != nil {
				return models.Task{}, fmt.Errorf("gagal mendapatkan data BRC: %v", err)
			}

			data := map[string]interface{}{
				"brcList":  brcList,
				"tasklist": updatedTask, // Store the struct directly; we'll serialize everything in one step
			}

			// fmt.Println("DATA NYA nih =>", data)

			payload := publisherModels.PublishMessageDTO{
				QueueName: "taskassignment-approval",
				Pattern:   "tasklist-approval",
				Body:      data,
			}
			err = taskassignments.publisher.PublishMessage(payload)
			if err != nil {
				taskassignments.logger.Zap.Error(err)
				return models.Task{}, fmt.Errorf("gagal mengirim message ke RabbitMQ")
			}
		}

	}

	return updatedTask, nil
}

// MyTasklist implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) MyTasklist(request models.TaskFilterRequest) (responses []models.MyTasklistResponse, totalData int64, err error) {
	responses, totalData, err = taskassignments.repository.MyTasklist(request)

	return responses, totalData, err
}

// MyTasklistDetail implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) MyTasklistDetail(id int64) (responses models.MyTasklistResponse, err error) {
	responses, err = taskassignments.repository.MyTasklistDetail(id)

	return responses, err
}

// GetMyTasklistTotal implements TaskAssignmentsDefinition.
func (taskassignments TaskAssignmentsService) GetMyTasklistTotal(request models.RequestMyTasklist) (total int64, err error) {
	total, err = taskassignments.repository.GetMyTasklistTotal(request)

	if err != nil {
		return 0, err
	}

	return total, err
}
