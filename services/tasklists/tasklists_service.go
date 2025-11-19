package tasklists

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	settingRepository "riskmanagement/repository/admin_setting"
	repository "riskmanagement/repository/tasklists"
	"strconv"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gitlab.com/golang-package-library/minio"

	"github.com/google/uuid"

	requestFile "riskmanagement/models/files"
	fileRepo "riskmanagement/repository/files"

	fileModel "riskmanagement/models/filemanager"
	filemanager "riskmanagement/services/filemanager"
)

type TasklistsDefinition interface {
	GetAll(request models.Paginate) (response []models.TasklistsResponse, pagination lib.Pagination, err error)
	GetTaskByID(id int64) (response []models.TasklistsResponse, pagination lib.Pagination, err error)
	GetAllOfficer(request models.Paginate) (responses []models.TasklistsFilterResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (response models.TasklistResponse, status bool, err error)
	Filter(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterResponse, pagination lib.Pagination, err error)
	FilterByID(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterResponse, pagination lib.Pagination, err error)
	FilterOfficer(request models.TasklistsFilterOfficerRequest) (responses []models.TasklistsFilterOfficerResponse, pagination lib.Pagination, err error)
	CheckAvailability(request models.TasklistsCheckRequest) (response models.TasklistsCheckResponse, err error)
	Store(request models.TasklistsStoreRequest) (status bool, err error)
	Update(request *models.TasklistsUpdateRequest) (status bool, err error)
	UpdateEndDate(request *models.TasklistsUpdateEndDateRequest) (status bool, err error)
	Delete(request *models.TasklistsUpdateDelete) (response bool, err error)
	Approval(request *models.TasklistsAprrovalRequest) (response bool, err error)
	Validation(request *models.TasklistsValidation) (response bool, err error)
	CountTask(request models.TasklistCountRequest) (response models.TasklistsCountResponse, err error)
	StoreDone(request models.TasklistsDoneHistoryRequest) (status bool, err error)
	GetDone(request models.TasklistsDoneHistoryCheckRequest) (responses models.TasklistsDoneHistoryResponse, status bool, err error)
	LeadAutocompleteVal(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error)
	LeadAutocompleteApr(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error)
	UserRegion(request models.UserRegionRequest) (responses models.UserRegionResponse, err error)
	GetDataAnomali(request models.TasklistDataAnomaliRequest) (responses []models.TasklistDataAnomaliResponse, err error)
	GetDataVerifikasi(request models.DataVerifikasiRequest) (responses []models.DataVerifikasiResponse, err error)
	GetLampiranIndicator(request *models.LampiranIndikatorCheck) (response models.LampiranIndikatorResponse, status bool, err error)
	DownloadLampiranIndikatorTemplate(request *models.LampiranIndikatorCheck) (response models.AnomaliHeaderResponse, err error)

	InsertTasklistRejected(request *models.TasklistsRejected) (err error)

	GetTasklist(request *models.TasklistCheckRequest) (response models.TasklistCheckResponse, err error)
	StoreDaily(request *models.TasklistDailyStore) (response *models.TasklistDailyStore, err error)
	UpdateDaily(request *models.ProgresUpdateRequest) (response *models.ProgresUpdateRequest, err error)

	GetAnomaliHeader(request *models.AnomaliHeader) (response models.AnomaliHeaderResponse, err error)
	GetAnomaliValue(request *models.AnomaliValue) (response []models.AnomaliValueResponse, err error)
	GetFirstLampiran(request *models.GetFirstLampiranRequest) (response models.GetFirstLampiranResponse, err error)
}

type TasklistsService struct {
	db                   lib.Database
	minio                minio.Minio
	logger               logger.Logger
	repository           repository.TasklistsDefinition
	settingRepository    settingRepository.AdminSettingDefinition
	uker                 repository.TasklistsUkerDefinition
	activity             repository.TasklistsActivityDefinition
	lampiran             repository.TasklistsLampiranDefinition
	tasklistsAnomaliKRID repository.TasklistsDataAnomaliKRIDDefinition
	fileRepo             fileRepo.FilesDefinition
	doneHistory          repository.TasklistsDoneHistoryDefinition
	filemanager          filemanager.FileManagerDefinition
	daily                repository.TasklistsDailyDefinition
}

func NewTasklistsService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	repository repository.TasklistsDefinition,
	settingRepository settingRepository.AdminSettingDefinition,
	uker repository.TasklistsUkerDefinition,
	activity repository.TasklistsActivityDefinition,
	lampiran repository.TasklistsLampiranDefinition,
	tasklistsAnomaliKRID repository.TasklistsDataAnomaliKRIDDefinition,
	fileRepo fileRepo.FilesDefinition,
	doneHistory repository.TasklistsDoneHistoryDefinition,
	filemanager filemanager.FileManagerDefinition,
) TasklistsDefinition {
	return TasklistsService{
		db:                   db,
		minio:                minio,
		logger:               logger,
		repository:           repository,
		settingRepository:    settingRepository,
		uker:                 uker,
		activity:             activity,
		lampiran:             lampiran,
		tasklistsAnomaliKRID: tasklistsAnomaliKRID,
		fileRepo:             fileRepo,
		doneHistory:          doneHistory,
		filemanager:          filemanager,
	}
}

func (tasklists TasklistsService) GetAll(request models.Paginate) (response []models.TasklistsResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	var dataTasklists []models.TasklistsResponse
	var totalRows, totalData int

	if request.Period == "All" {
		dataTasklists, totalRows, totalData, err = tasklists.repository.GetAll(&request)
		if err != nil {
			tasklists.logger.Zap.Error(err)
			return response, pagination, err
		}
	} else {
		dataTasklists, totalRows, totalData, err = tasklists.repository.GetByTaskType(&request)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			return response, pagination, err
		}
	}

	for _, tasklist := range dataTasklists {
		// subtask, err := tasklists.repository.GetSubTask(tasklist.TasklistID)
		// if err != nil {
		// 	tasklists.logger.Zap.Error(err)
		// 	return response, pagination, err
		// }

		// countDoneSampleRequest := &models.CountDoneSampleRequest{
		// 	Kegiatan:      tasklist.Kegiatan,
		// 	ActivityID:    tasklist.ActivityID,
		// 	ProductID:     tasklist.ProductID,
		// 	RiskIssueID:   tasklist.RiskIssueID,
		// 	RiskIssueCode: tasklist.RiskIssueCode,
		// 	Branch:        tasklist.Branch,
		// }

		// countDoneSample, err := tasklists.repository.DoneSampleCount(countDoneSampleRequest)

		// if err != nil {
		// 	tasklists.logger.Zap.Error(err)
		// 	return response, pagination, err
		// }

		response = append(response, models.TasklistsResponse{
			BRDESC:          tasklist.BRDESC,
			Kegiatan:        tasklist.Kegiatan,
			TasklistID:      tasklist.TasklistID,
			Activity:        tasklist.Activity,
			Product:         tasklist.Product,
			RiskIssue:       tasklist.RiskIssue,
			RiskIssueID:     tasklist.RiskIssueID,
			RiskIndicator:   tasklist.RiskIndicator,
			RiskIndicatorID: tasklist.RiskIndicatorID,
			StartDate:       tasklist.StartDate,
			EndDate:         tasklist.EndDate,
			DoneSample:      tasklist.DoneSample,
			Sample:          tasklist.Sample,
			Status:          tasklist.Status,
			JenisTask:       tasklist.JenisTask,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return response, pagination, err
}

func (tasklists TasklistsService) GetTaskByID(id int64) (response []models.TasklistsResponse, pagination lib.Pagination, err error) {
	var dataTasklists []models.TasklistsResponse
	var totalRows, totalData int

	dataTasklists, totalRows, totalData, err = tasklists.repository.GetTaskByID(id)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		return response, pagination, err
	}

	for _, tasklist := range dataTasklists {
		// subtask, err := tasklists.repository.GetSubTask(tasklist.TasklistID)
		// if err != nil {
		// 	tasklists.logger.Zap.Error(err)
		// 	return response, pagination, err
		// }

		// countDoneSampleRequest := &models.CountDoneSampleRequest{
		// 	Kegiatan:      tasklist.Kegiatan,
		// 	ActivityID:    tasklist.ActivityID,
		// 	ProductID:     tasklist.ProductID,
		// 	RiskIssueID:   tasklist.RiskIssueID,
		// 	RiskIssueCode: tasklist.RiskIssueCode,
		// 	Branch:        tasklist.Branch,
		// }

		// countDoneSample, err := tasklists.repository.DoneSampleCount(countDoneSampleRequest)

		// if err != nil {
		// 	tasklists.logger.Zap.Error(err)
		// 	return response, pagination, err
		// }

		response = append(response, models.TasklistsResponse{
			TasklistID:      tasklist.TasklistID,
			BRDESC:          tasklist.BRDESC,
			Kegiatan:        tasklist.Kegiatan,
			Activity:        tasklist.Activity,
			Product:         tasklist.Product,
			RiskIssue:       tasklist.RiskIssue,
			RiskIssueID:     tasklist.RiskIssueID,
			RiskIndicator:   tasklist.RiskIndicator,
			RiskIndicatorID: tasklist.RiskIndicatorID,
			StartDate:       tasklist.StartDate,
			EndDate:         tasklist.EndDate,
			DoneSample:      tasklist.DoneSample,
			Sample:          tasklist.Sample,
			Status:          tasklist.Status,
			JenisTask:       tasklist.JenisTask,
		})
	}

	pagination = lib.SetPaginationResponse(1, 1, totalRows, totalData)

	return response, pagination, err
}

func (tasklists TasklistsService) GetOne(id int64) (response models.TasklistResponse, status bool, err error) {
	dataTasklistsUker, err := tasklists.repository.ShowUker(id)

	if err != nil {
		fmt.Println(err)
	}

	dataTasklistsNote, err := tasklists.repository.GetNotes(id)

	if err != nil {
		fmt.Println(err)
	}

	files, err := tasklists.lampiran.GetOneFileByID(id)

	if err != nil {
		fmt.Println(err)
	}

	dataTasklists, err := tasklists.repository.GetOne(id)

	if dataTasklists.ID != 0 {
		fmt.Println("bukan 0")

		response = models.TasklistResponse{
			ID:              dataTasklists.ID,
			NoTasklist:      dataTasklists.NoTasklist,
			NamaTasklist:    dataTasklists.NamaTasklist,
			Uker:            dataTasklistsUker,
			SumberData:      dataTasklists.SumberData,
			RAP:             dataTasklists.RAP,
			ActivityID:      dataTasklists.ActivityID,
			Activity:        dataTasklists.Activity,
			ProductID:       dataTasklists.ProductID,
			ProductName:     dataTasklists.ProductName,
			RiskIssueID:     dataTasklists.RiskIssueID,
			RiskIssue:       dataTasklists.RiskIssue,
			RiskIndicatorID: dataTasklists.RiskIndicatorID,
			RiskIndicator:   dataTasklists.RiskIndicator,
			StartDate:       dataTasklists.StartDate,
			EndDate:         dataTasklists.EndDate,
			Status:          dataTasklists.Status,
			IsiLampiran:     dataTasklists.IsiLampiran,
			Notes:           dataTasklistsNote,
			TaskType:        dataTasklists.TaskType,
			TaskTypeName:    dataTasklists.TaskTypeName,
			TaskTypePeriod:  dataTasklists.TaskTypePeriod,
			Validation:      dataTasklists.Validation,
			ValidationName:  dataTasklists.ValidationName,
			Approval:        dataTasklists.Approval,
			ApprovalName:    dataTasklists.ApprovalName,
			ApprovalStatus:  dataTasklists.ApprovalStatus,
			Range:           dataTasklists.Range,
			Upload:          dataTasklists.Upload,
			Kegiatan:        dataTasklists.Kegiatan,
			Period:          dataTasklists.Period,
			Sample:          dataTasklists.Sample,
			Files:           files,
			MakerID:         dataTasklists.MakerID,
			CreatedAt:       dataTasklists.CreatedAt,
		}

		return response, true, err
	}

	return response, false, err
}

func (tasklists TasklistsService) GetAllOfficer(request models.Paginate) (responses []models.TasklistsFilterResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	tasklistsData, totalRows, totalData, err := tasklists.repository.GetAllOfficer(request)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range tasklistsData {
		ukerListReq := &models.UkerListReq{
			ID:       response.ID,
			PERNR:    request.PERNR,
			TipeUker: request.TipeUker,
			HILFM:    request.HILFM,
			StellTX:  request.StellTX,
			REGION:   request.REGION,
		}

		dataPegawai, err := tasklists.repository.GetTipeUkerMaker(response.MakerID)
		if err != nil {
			tasklists.logger.Zap.Error(err)
			return responses, pagination, err
		}

		uker, err := tasklists.repository.GetUkerList(*ukerListReq)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			return responses, pagination, err
		}
		responses = append(responses, models.TasklistsFilterResponse{
			ID: response.ID,
			// RGDESC:         response.RGDESC,
			// MBDESC:         response.MBDESC,
			// BRDESC:         response.BRDESC,
			Activity:       response.Activity,
			Product:        response.Product,
			Period:         response.Period,
			StartDate:      response.StartDate,
			EndDate:        response.EndDate,
			RiskIssue:      response.RiskIssue,
			RiskIndicator:  response.RiskIndicator,
			StatusApproval: response.StatusApproval,
			JenisTask:      response.JenisTask,
			Status:         response.Status,
			MakerID:        response.MakerID,
			Uker:           uker,
			DataPegawai:    dataPegawai,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

func (tasklists TasklistsService) UserRegion(request models.UserRegionRequest) (responses models.UserRegionResponse, err error) {
	return tasklists.repository.UserRegion(request)
}

func (tasklists TasklistsService) Filter(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	tasklistsData, totalRows, totalData, err := tasklists.repository.Filter(request)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range tasklistsData {
		ukerListReq := &models.UkerListReq{
			ID:       response.ID,
			PERNR:    request.PERNR,
			TipeUker: request.TipeUker,
		}

		uker, err := tasklists.repository.GetUkerList(*ukerListReq)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			return responses, pagination, err
		}

		var mini_link fileModel.FileManagerResponseUrl

		if response.Filename != "" {
			mini_link, err = tasklists.filemanager.GetFile(fileModel.FileManagerRequest{
				Subdir:   response.Path,
				Filename: response.Filename,
			})

			if err != nil {
				tasklists.logger.Zap.Error(err)
			}
		}

		responses = append(responses, models.TasklistsFilterResponse{
			ID: response.ID,
			// RGDESC:         response.RGDESC,
			// MBDESC:         response.MBDESC,
			// BRANCH:         response.BRANCH,
			// BRDESC:         response.BRDESC,
			StartDate:      response.StartDate,
			EndDate:        response.EndDate,
			Activity:       response.Activity,
			Product:        response.Product,
			RiskIssue:      response.RiskIssue,
			RiskIndicator:  response.RiskIndicator,
			StatusApproval: response.StatusApproval,
			MakerID:        response.MakerID,
			JenisTask:      response.JenisTask,
			Period:         response.Period,
			IDLampiran:     response.IDLampiran,
			Filename:       response.Filename,
			Path:           mini_link.MinioPath,
			Ext:            response.Ext,
			Size:           response.Size,
			Validation:     response.Validation,
			Approval:       response.Approval,
			Status:         response.Status,
			Uker:           uker,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

func (tasklists TasklistsService) FilterByID(request models.TasklistsFilterRequest) (responses []models.TasklistsFilterResponse, pagination lib.Pagination, err error) {
	tasklistsData, totalRows, totalData, err := tasklists.repository.FilterByID(request)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range tasklistsData {
		ukerListReq := &models.UkerListReq{
			ID:       response.ID,
			PERNR:    request.PERNR,
			TipeUker: request.TipeUker,
		}

		uker, err := tasklists.repository.GetUkerList(*ukerListReq)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			return responses, pagination, err
		}

		var mini_link fileModel.FileManagerResponseUrl

		if response.Filename != "" {
			mini_link, err = tasklists.filemanager.GetFile(fileModel.FileManagerRequest{
				Subdir:   response.Path,
				Filename: response.Filename,
			})

			if err != nil {
				tasklists.logger.Zap.Error(err)
			}
		}

		responses = append(responses, models.TasklistsFilterResponse{
			ID: response.ID,
			// RGDESC:         response.RGDESC,
			// MBDESC:         response.MBDESC,
			// BRDESC:         response.BRDESC,
			StartDate:      response.StartDate,
			EndDate:        response.EndDate,
			Activity:       response.Activity,
			Product:        response.Product,
			RiskIssue:      response.RiskIssue,
			RiskIndicator:  response.RiskIndicator,
			StatusApproval: response.StatusApproval,
			MakerID:        response.MakerID,
			JenisTask:      response.JenisTask,
			Period:         response.Period,
			IDLampiran:     response.IDLampiran,
			Filename:       response.Filename,
			Path:           mini_link.MinioPath,
			Ext:            response.Ext,
			Size:           response.Size,
			Validation:     response.Validation,
			Approval:       response.Approval,
			Status:         response.Status,
			Uker:           uker,
		})
	}

	pagination = lib.SetPaginationResponse(1, 1, totalRows, totalData)

	return responses, pagination, err
}

func (tasklists TasklistsService) FilterOfficer(request models.TasklistsFilterOfficerRequest) (responses []models.TasklistsFilterOfficerResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	tasklistsData, totalRows, totalData, err := tasklists.repository.FilterOfficer(request)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range tasklistsData {
		ukerListReq := &models.UkerListReq{
			ID:       response.ID,
			PERNR:    request.PERNR,
			TipeUker: request.TipeUker,
			HILFM:    request.HILFM,
			StellTX:  request.StellTX,
			REGION:   request.REGION,
		}

		dataPegawai, err := tasklists.repository.GetTipeUkerMaker(response.MakerID)
		if err != nil {
			tasklists.logger.Zap.Error(err)
			return responses, pagination, err
		}

		uker, err := tasklists.repository.GetUkerList(*ukerListReq)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.TasklistsFilterOfficerResponse{
			ID: response.ID,
			// RGDESC:         response.RGDESC,
			// MBDESC:         response.MBDESC,
			// BRDESC:         response.BRDESC,
			StartDate:      response.StartDate,
			EndDate:        response.EndDate,
			Activity:       response.Activity,
			Product:        response.Product,
			RiskIssue:      response.RiskIssue,
			RiskIndicator:  response.RiskIndicator,
			StatusApproval: response.StatusApproval,
			MakerID:        response.MakerID,
			JenisTask:      response.JenisTask,
			Period:         response.Period,
			Validation:     response.Validation,
			Approval:       response.Approval,
			Status:         response.Status,
			Uker:           uker,
			DataPegawai:    dataPegawai,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

func (tasklists TasklistsService) CheckAvailability(request models.TasklistsCheckRequest) (response models.TasklistsCheckResponse, err error) {
	return tasklists.repository.CheckAvailability(request)
}

func (tasklists TasklistsService) Store(request models.TasklistsStoreRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := tasklists.db.DB.Begin()

	TaskID := request.TaskType

	limit, err := tasklists.repository.LimitTask(TaskID)

	reqTasklists := &models.Tasklists{}

	//input data tasklists
	if *request.StartDate == "" && *request.EndDate == "" {
		reqTasklists = &models.Tasklists{
			NoTasklist:      request.NoTasklist,
			NamaTasklist:    request.NamaTasklist,
			ActivityID:      request.ActivityID,
			ProductID:       request.ProductID,
			ProductName:     request.ProductName,
			RiskIssueID:     request.RiskIssueID,
			RiskIssue:       request.RiskIssue,
			RiskIndicatorID: request.RiskIndicatorID,
			RiskIndicator:   request.RiskIndicator,
			TaskType:        request.TaskType,
			TaskTypeName:    request.TaskTypeName,
			Kegiatan:        request.Kegiatan,
			Period:          request.Period,
			SumberData:      request.SumberData,
			RAP:             request.RAP,
			Sample:          request.Sample,
			Approval:        request.Approval,
			ApprovalName:    request.ApprovalName,
			Validation:      request.Validation,
			ValidationName:  request.ValidationName,
			ApprovalStatus:  "Minta Persetujuan Validasi",
			Status:          "Aktif",
			MakerID:         request.MakerID,
			CreatedAt:       &timeNow,
			UpdatedAt:       &timeNow,
		}
	} else {
		reqTasklists = &models.Tasklists{
			NoTasklist:      request.NoTasklist,
			NamaTasklist:    request.NamaTasklist,
			ActivityID:      request.ActivityID,
			ProductID:       request.ProductID,
			ProductName:     request.ProductName,
			RiskIssueID:     request.RiskIssueID,
			RiskIssue:       request.RiskIssue,
			RiskIndicatorID: request.RiskIndicatorID,
			RiskIndicator:   request.RiskIndicator,
			TaskType:        request.TaskType,
			TaskTypeName:    request.TaskTypeName,
			Kegiatan:        request.Kegiatan,
			Period:          request.Period,
			SumberData:      request.SumberData,
			RAP:             request.RAP,
			Sample:          request.Sample,
			StartDate:       request.StartDate,
			EndDate:         request.EndDate,
			Approval:        request.Approval,
			ApprovalName:    request.ApprovalName,
			Validation:      request.Validation,
			ValidationName:  request.ValidationName,
			ApprovalStatus:  "Minta Persetujuan Validasi",
			Status:          "Aktif",
			MakerID:         request.MakerID,
			CreatedAt:       &timeNow,
			UpdatedAt:       &timeNow,
		}
	}

	dataTasklists, err := tasklists.repository.Store(reqTasklists, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}
	//end data tasklists

	risString := strconv.FormatInt(request.RiskIssueID, 10)
	rinString := strconv.FormatInt(request.RiskIndicatorID, 10)

	if len(request.IsiLampiran) > 0 {
		tasklistLampiranIndikator := &models.LampiranIndikatorCheck{
			RiskIssueID:     risString,
			RiskIndicatorID: rinString,
		}

		dataLampiranIndikator, _ := tasklists.repository.GetLampiranIndicator(tasklistLampiranIndikator)

		var jsonString []string

		if request.HeaderLampiran == "" {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, errors.New("Isi file tidak sesuai format yang seharusnya!")
		}

		replacerHeader := strings.NewReplacer("\\", "", "[", "", "]", "", `"`, "`")
		removedHeader := replacerHeader.Replace(request.HeaderLampiran)

		if dataLampiranIndikator.ID == 0 {
			tasklistHeader := &models.TasklistHeaderStore{
				TasklistID:      dataTasklists.ID,
				RiskIssueID:     risString,
				RiskIndicatorID: rinString,
				HeaderLampiran:  removedHeader,
			}

			err = tasklists.repository.CreateTableLampiran(tasklistHeader)

			if err != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(err)
				return false, err
			}

			lampiranIndicator := &models.LampiranIndikatorStore{
				RiskIssueID:       request.RiskIssueID,
				RiskIndicatorID:   request.RiskIndicatorID,
				NamaTable:         "lampiran_rap_" + risString + "_" + rinString,
				JumlahKolom:       request.JumlahKolom,
				RiskIndicatorDesc: request.RiskIndicator,
			}

			err = tasklists.repository.InsertTableLampiranIndikator(lampiranIndicator)
		}

		for _, item := range request.IsiLampiran {
			jsonData, _ := json.Marshal(item)
			jsonString = append(jsonString, string(jsonData))
		}

		columnEntry := ""
		for i, data := range jsonString {
			if i > 0 {
				// removed := data[7 : len(data)-4]

				replacer := strings.NewReplacer("\\", "", "[", "", "]", "")
				removed := replacer.Replace(data)
				idTaskConv := strconv.Itoa(int(dataTasklists.ID))
				if i != len(jsonString)-1 {
					// columnEntry += "(" + removed[6:len(removed)-1] + "), "
					columnEntry += "(" + idTaskConv + ", '', " + removed[6:len(removed)-1] + "), "
					// columnEntry += "(" + removed[6:len(removed)-1] + "), "
				} else {
					// columnEntry += "(" + removed[6:len(removed)-1] + ")"
					columnEntry += "(" + idTaskConv + ", '', " + removed[6:len(removed)-1] + ")"
					// columnEntry += "(" + removed[6:len(removed)-1] + ")"
				}
			}
		}

		taskIDInt := strconv.Itoa(int(dataTasklists.ID))

		tasklistColumn := &models.TasklistColumnStore{
			TasklistID:      taskIDInt,
			RiskIssueID:     risString,
			RiskIndicatorID: rinString,
			HeaderLampiran:  removedHeader,
			IsiLampiran:     columnEntry,
		}

		err = tasklists.repository.InsertTableLampiran(tasklistColumn)

		if err != nil {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, err
		}
	}

	ukerList := ""
	//Begin Input data uker
	if len(request.Uker) != 0 {
		count := 0
		for _, value := range request.Uker {
			// objectString := []byte(value.Object)
			tasklistRiskIssueAvailReq := &models.TasklistRiskIssueAvailReq{
				RiskIssueID: request.RiskIssueID,
				Branch:      value.BRANCH,
				ActivityID:  request.ActivityID,
				ProductID:   request.ProductID,
			}
			totalRiskIssue, _ := tasklists.repository.CekRiskIssueAvail(tasklistRiskIssueAvailReq)

			total, _ := tasklists.repository.CheckMaxLimit(TaskID, value.BRANCH)
			if limit != 0 && limit <= total {
				fmt.Println("isi limit:", limit)
				fmt.Println("isi total:", total)
				tx.Rollback()
				err1 := errors.New("Tasklist pada unit kerja ini sudah mencapai batas limit")
				tasklists.logger.Zap.Error("Limit tasklist")
				return false, err1
			} else if totalRiskIssue.Total > 0 {
				tx.Rollback()
				err2 := errors.New("Risk issue sudah terpakai")
				tasklists.logger.Zap.Error("Risk issue terpakai")
				return false, err2
			} else {
				if count == 0 {
					ukerList += value.BRDESC
				} else {
					ukerList += ", " + value.BRDESC
				}

				_, err := tasklists.uker.Store(&models.TasklistsUker{
					REGION:          value.REGION,
					RGDESC:          value.RGDESC,
					MAINBR:          value.MAINBR,
					MBDESC:          value.MBDESC,
					BRANCH:          value.BRANCH,
					BRDESC:          value.BRDESC,
					JumlahNominatif: request.Sample,
					TasklistID:      dataTasklists.ID,
				}, tx)

				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}
			}
			count++
		}
	}
	//End Input data uker

	//Begin Input Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	if request.Files[0].Filename != "" {
		for _, value := range request.Files {
			UUID := uuid.New()
			var destinationPath string
			bucketExist := tasklists.minio.BucketExist(tasklists.minio.Client(), bucket)

			pathSplit := strings.Split(value.Path, "/")
			sourcePath := fmt.Sprint(value.Path)
			destinationPath = pathSplit[1] + "/" +
				lib.GetTimeNow("year") + "/" +
				lib.GetTimeNow("month") + "/" +
				lib.GetTimeNow("day") + "/" +
				UUID.String() + "/" +
				pathSplit[2] + "/" + value.Filename

			if bucketExist {
				fmt.Println("Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				uploaded := tasklists.minio.CopyObject(tasklists.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			} else {
				fmt.Println("Not Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				tasklists.minio.MakeBucket(tasklists.minio.Client(), bucket, "")
				uploaded := tasklists.minio.CopyObject(tasklists.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			}

			files, err := tasklists.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(err)
				return false, err
			}

			_, err = tasklists.lampiran.Store(&models.TasklistsFiles{
				TasklistsID: dataTasklists.ID,
				FilesID:     files.ID,
				CreatedAt:   &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(err)
				return false, err
			}
		}
	}

	//End Input Lampiran

	reqNotif := &models.TasklistNotif{
		TaskID:     dataTasklists.ID,
		Tanggal:    &timeNow,
		Keterangan: "Ada 1 tasklist baru yang perlu divalidasi",
		Status:     0,
		Jenis:      "Validasi Tasklist",
		Receiver:   dataTasklists.Validation,
		Uker:       ukerList,
	}

	_, err = tasklists.repository.StoreNotif(reqNotif, tx)
	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

func (tasklists TasklistsService) Update(request *models.TasklistsUpdateRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := tasklists.db.DB.Begin()

	updateTasklists := &models.TasklistsUpdate{
		ID:                 request.ID,
		NoTasklist:         request.NoTasklist,
		NamaTasklist:       request.NamaTasklist,
		ActivityID:         request.ActivityID,
		ProductID:          request.ProductID,
		ProductName:        request.ProductName,
		RiskIssueID:        request.RiskIssueID,
		RiskIssue:          request.RiskIssue,
		RiskIndicatorID:    request.RiskIndicatorID,
		RiskIndicator:      request.RiskIndicator,
		TaskType:           request.TaskType,
		TaskTypeName:       request.TaskTypeName,
		Kegiatan:           request.Kegiatan,
		Period:             request.Period,
		SumberData:         request.SumberData,
		RAP:                request.RAP,
		Sample:             request.Sample,
		StartDate:          request.StartDate,
		EndDate:            request.EndDate,
		Approval:           request.Approval,
		ApprovalName:       request.ApprovalName,
		Validation:         request.Validation,
		ValidationName:     request.ValidationName,
		ApprovalStatus:     "Minta Persetujuan Validasi",
		VerificationStatus: "Not Done",
		Status:             "Aktif",
		MakerID:            request.MakerID,
		CreatedAt:          request.CreatedAt,
		UpdatedAt:          &timeNow,
	}

	dataTasklists, err := tasklists.repository.Update(updateTasklists, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	//Update & add Uker
	deleteUker := models.TasklistsUker{
		TasklistID: request.ID,
	}
	err = tasklists.uker.Delete(&deleteUker, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	ukerList := ""

	if len(request.Uker) != 0 {
		count := 0
		for _, value := range request.Uker {
			// objectString := []byte(value.Object)

			fmt.Println("object Stringg =======>s", value)
			sampleInt64, errConv := strconv.ParseInt(request.Sample, 10, 64)
			if errConv != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(errConv)
				return false, errConv
			}
			_, err := tasklists.uker.Store(&models.TasklistsUker{
				REGION:          value.REGION,
				RGDESC:          value.RGDESC,
				MAINBR:          value.MAINBR,
				MBDESC:          value.MBDESC,
				BRANCH:          value.BRANCH,
				BRDESC:          value.BRDESC,
				JumlahNominatif: sampleInt64,
				TasklistID:      dataTasklists.ID,
			}, tx)

			if err != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(err)
				return false, err
			}

			if count == 0 {
				ukerList += value.BRDESC
			} else {
				ukerList += ", " + value.BRDESC
			}

			count++
		}
	}

	//End Update & add Uker

	risStringCek := strconv.FormatInt(request.RiskIssueID, 10)
	rinStringCek := strconv.FormatInt(request.RiskIndicatorID, 10)

	tasklistLampiranIndikatorCek := &models.LampiranIndikatorCheck{
		RiskIssueID:     risStringCek,
		RiskIndicatorID: rinStringCek,
	}

	dataLampiranIndikatorCek, _ := tasklists.repository.GetLampiranIndicator(tasklistLampiranIndikatorCek)

	if dataLampiranIndikatorCek.ID != 0 {
		fmt.Println("table ditemukan")
		deleteIsiLampiran := models.TasklistLampiranDelete{
			TasklistID:    request.ID,
			RiskIssue:     request.RiskIssueID,
			RiskIndicator: request.RiskIndicatorID,
		}
		err = tasklists.repository.DeleteIsiLampiran(&deleteIsiLampiran, tx)

		if err != nil {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, err
		}
		fmt.Println("delete isi lampiran berhasil")
	}

	//Update & add Lampiran
	if len(request.IsiLampiran) > 0 {
		risString := strconv.FormatInt(request.RiskIssueID, 10)
		rinString := strconv.FormatInt(request.RiskIndicatorID, 10)

		tasklistLampiranIndikator := &models.LampiranIndikatorCheck{
			RiskIssueID:     risString,
			RiskIndicatorID: rinString,
		}

		dataLampiranIndikator, _ := tasklists.repository.GetLampiranIndicator(tasklistLampiranIndikator)

		var jsonString []string

		if request.HeaderLampiran == "" {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, errors.New("Isi file tidak sesuai format yang seharusnya!")
		}

		replacerHeader := strings.NewReplacer("\\", "", "[", "", "]", "", `"`, "`")
		removedHeader := replacerHeader.Replace(request.HeaderLampiran)

		if dataLampiranIndikator.ID == 0 {
			tasklistHeader := &models.TasklistHeaderStore{
				TasklistID:      dataTasklists.ID,
				RiskIssueID:     risString,
				RiskIndicatorID: rinString,
				HeaderLampiran:  removedHeader,
			}

			err = tasklists.repository.CreateTableLampiran(tasklistHeader)

			if err != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(err)
				return false, err
			}

			lampiranIndicator := &models.LampiranIndikatorStore{
				RiskIssueID:       request.RiskIssueID,
				RiskIndicatorID:   request.RiskIndicatorID,
				NamaTable:         "lampiran_rap_" + risString + "_" + rinString,
				JumlahKolom:       request.JumlahKolom,
				RiskIndicatorDesc: request.RiskIndicator,
			}

			err = tasklists.repository.InsertTableLampiranIndikator(lampiranIndicator)
		}

		deleteIsiLampiran := models.TasklistLampiranDelete{
			TasklistID:    request.ID,
			RiskIssue:     request.RiskIssueID,
			RiskIndicator: request.RiskIndicatorID,
		}
		err = tasklists.repository.DeleteIsiLampiran(&deleteIsiLampiran, tx)

		if err != nil {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, err
		}

		for _, item := range request.IsiLampiran {
			jsonData, _ := json.Marshal(item)
			jsonString = append(jsonString, string(jsonData))
		}

		columnEntry := ""
		for i, data := range jsonString {
			if i > 0 {
				// removed := data[7 : len(data)-4]

				replacer := strings.NewReplacer("\\", "", "[", "", "]", "")
				removed := replacer.Replace(data)
				idTaskConv := strconv.Itoa(int(dataTasklists.ID))
				if i != len(jsonString)-1 {
					// columnEntry += "(" + removed[6:len(removed)-1] + "), "
					columnEntry += "(" + idTaskConv + ", '', " + removed[6:len(removed)-1] + "), "
					// columnEntry += "(" + removed[6:len(removed)-1] + "), "
				} else {
					// columnEntry += "(" + removed[6:len(removed)-1] + ")"
					columnEntry += "(" + idTaskConv + ", '', " + removed[6:len(removed)-1] + ")"
					// columnEntry += "(" + removed[6:len(removed)-1] + ")"
				}
			}
		}

		taskIDInt := strconv.Itoa(int(dataTasklists.ID))

		tasklistColumn := &models.TasklistColumnStore{
			TasklistID:      taskIDInt,
			RiskIssueID:     risString,
			RiskIndicatorID: rinString,
			HeaderLampiran:  removedHeader,
			IsiLampiran:     columnEntry,
		}

		err = tasklists.repository.InsertTableLampiran(tasklistColumn)

		if err != nil {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, err
		}
	}

	//End Update & add isi Lampiran

	//Update & add Activity
	deleteIndicator := models.TasklistsRiskIndicator{
		TasklistsID: request.ID,
	}
	err = tasklists.activity.Delete(&deleteIndicator, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	// if len(request.RiskIndicator) != 0 {
	// 	for _, value := range request.RiskIndicator {
	// 		_, err := tasklists.activity.Store(&models.TasklistsRiskIndicator{
	// 			RiskIndicatorID: value.RiskIndicatorID,
	// 			TasklistsID:     dataTasklists.ID,
	// 		}, tx)

	// 		if err != nil {
	// 			tx.Rollback()
	// 			tasklists.logger.Zap.Error(err)
	// 			return false, err
	// 		}
	// 	}
	// } else {
	// 	if err != nil {
	// 		tx.Rollback()
	// 		tasklists.logger.Zap.Error(err)
	// 		return false, err
	// 	}
	// }
	//End Update & add Activity

	//Begin Input data anomali
	deleteAnomali := models.TasklistsAnomaliDataKRIDDelete{
		TasklistID: request.ID,
	}
	err = tasklists.tasklistsAnomaliKRID.Delete(&deleteAnomali, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	// if request.SumberData != "KRID" {
	// 	if len(request.DataAnomaliKRID) != 0 {
	// 		for _, value := range request.DataAnomaliKRID {
	// 			fmt.Println("object Stringg =======>s", value)
	// 			if value.Object != "" {
	// 				_, err = tasklists.tasklistsAnomaliKRID.Store(&models.TasklistsAnomaliDataKRIDRequest{
	// 					TasklistID: dataTasklists.ID,
	// 					Object:     value.Object,
	// 				}, tx)

	// 				if err != nil {
	// 					tx.Rollback()
	// 					tasklists.logger.Zap.Error(err)
	// 					return false, err
	// 				}
	// 			}

	// 		}
	// 	} else {
	// 		tx.Rollback()
	// 		tasklists.logger.Zap.Error(err)
	// 		return false, err
	// 	}
	// }
	//End Input data anomali

	//Begin Update Lampiran
	bucket := os.Getenv("BUCKET_NAME")

	// if len(request.Files) > 1 {
	err = tasklists.lampiran.DeleteFilesByID(request.ID, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	if request.Files[0].Filename != "" {
		for _, value := range request.Files {
			UUID := uuid.New()

			var destinationPath string
			bucketExist := tasklists.minio.BucketExist(tasklists.minio.Client(), bucket)

			pathSplit := strings.Split(value.Path, "/")
			sourcePath := fmt.Sprint(value.Path)
			destinationPath = pathSplit[1] + "/" +
				lib.GetTimeNow("year") + "/" +
				lib.GetTimeNow("month") + "/" +
				lib.GetTimeNow("day") + "/" +
				UUID.String() + "/" +
				pathSplit[2] + "/" + value.Filename

			if bucketExist {
				fmt.Println("Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				uploaded := tasklists.minio.CopyObject(tasklists.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			} else {
				fmt.Println("Not Exist")
				fmt.Println(bucket)
				fmt.Println(sourcePath)
				fmt.Println(destinationPath)
				tasklists.minio.MakeBucket(tasklists.minio.Client(), bucket, "")
				uploaded := tasklists.minio.CopyObject(tasklists.minio.Client(), bucket, sourcePath, bucket, destinationPath)

				fmt.Println(uploaded)
			}

			files, err := tasklists.fileRepo.Store(&requestFile.Files{
				Filename:  value.Filename,
				Path:      destinationPath,
				Extension: value.Extension,
				Size:      value.Size,
				CreatedAt: &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(err)
				return false, err
			}

			_, err = tasklists.lampiran.Store(&models.TasklistsFiles{
				TasklistsID: dataTasklists.ID,
				FilesID:     files.ID,
				CreatedAt:   &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				tasklists.logger.Zap.Error(err)
				return false, err
			}
		}
	}

	reqNotif := &models.TasklistNotif{
		TaskID:     dataTasklists.ID,
		Tanggal:    &timeNow,
		Keterangan: "Ada 1 tasklist baru yang perlu divalidasi",
		Status:     0,
		Jenis:      "Validasi Tasklist",
		Receiver:   dataTasklists.Validation,
		Uker:       ukerList,
	}

	_, err = tasklists.repository.StoreNotif(reqNotif, tx)
	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	//End Update Lampiran

	tx.Commit()
	return true, err
}

func (tasklists TasklistsService) UpdateEndDate(request *models.TasklistsUpdateEndDateRequest) (status bool, err error) {
	tx := tasklists.db.DB.Begin()

	updateTasklists := &models.TasklistsUpdateEndDate{
		ID:      request.ID,
		EndDate: request.EndDate,
	}
	_, err = tasklists.repository.UpdateEndDate(updateTasklists, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

func (tasklists TasklistsService) Delete(request *models.TasklistsUpdateDelete) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := tasklists.db.DB.Begin()

	getOneTasklists, exist, err := tasklists.GetOne(request.ID)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	UpdateDataTasklists := &models.TasklistsUpdateDelete{
		ID:        request.ID,
		Status:    "Non Aktif",
		UpdatedAt: &timeNow,
	}

	_, err = tasklists.repository.Delete(UpdateDataTasklists, tx)
	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneTasklists", getOneTasklists)
		tx.Commit()
		return true, err
	}

	return false, err
}

func (tasklists TasklistsService) Approval(request *models.TasklistsAprrovalRequest) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := tasklists.db.DB.Begin()

	UpdateApproval := &models.TasklistsAprroval{
		ID:             request.ID,
		ApprovalStatus: request.ApprovalStatus,
		UpdatedAt:      &timeNow,
	}

	_, err = tasklists.repository.Approval(UpdateApproval, tx)
	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	if request.ApprovalStatus == "Ditolak oleh Approver" || request.ApprovalStatus == "Ditolak oleh Validator" {
		// if request.ApprovalStatus == "Ditolak" {
		reqTasklistRejected := &models.TasklistsRejected{
			TasklistID: request.ID,
			Notes:      request.Notes,
			Status:     "Not Done",
		}

		err = tasklists.repository.InsertTasklistRejected(reqTasklistRejected, tx)

		if err != nil {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, err
		}
	} else if request.ApprovalStatus == "Disetujui" {
		today := time.Now()

		dataTasklist, _, err := tasklists.GetOne(request.ID)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			tx.Rollback()
			return false, err
		}

		dataTaskType, err := tasklists.settingRepository.GetOne(dataTasklist.TaskType)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			tx.Rollback()
			return false, err
		}

		if dataTasklist.Period == "Daily" {
			tx.Commit()
			return true, err
		}

		// Parse string ke objek time.Time
		startDay, err := time.Parse(time.RFC3339, *dataTasklist.StartDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
		}

		endDay, err := time.Parse(time.RFC3339, *dataTasklist.EndDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
		}

		// Mendapatkan bagian tanggal (hari) dari objek time.Time
		day := startDay.Day()

		if dataTasklist.Period == "Monthly" && day == today.Day() {
			brcRequest := &models.GetBRCRequest{
				// Branch: request.Branch,
				Uker: request.Uker,
			}

			brcList, err := tasklists.repository.GetDataBRC(brcRequest)
			if err != nil {
				tasklists.logger.Zap.Error(err)
				return response, err
			}

			for _, brc := range brcList {
				reqTaskDaily := &models.TasklistsToday{
					TasklistID:      dataTasklist.ID,
					ActivityID:      dataTasklist.ActivityID,
					ProductID:       dataTasklist.ProductID,
					Product:         dataTasklist.ProductName,
					RiskIssueID:     dataTasklist.RiskIssueID,
					RiskIssue:       dataTasklist.RiskIssue,
					RiskIndicatorID: dataTasklist.RiskIndicatorID,
					RiskIndicator:   dataTasklist.RiskIndicator,
					StartDate:       dataTasklist.StartDate,
					EndDate:         dataTasklist.EndDate,
					Status:          dataTasklist.Status,
					TaskType:        dataTasklist.TaskType,
					TaskTypeName:    dataTasklist.TaskTypeName,
					Kegiatan:        dataTasklist.Kegiatan,
					Period:          dataTasklist.Period,
					Sample:          dataTasklist.Sample,
					Progres:         0,
					Persentase:      0,
					MakerID:         dataTasklist.MakerID,
					PERNR:           brc.PERNR,
					Assigned:        brc.SNAME,
					REGION:          brc.REGION,
					RGDESC:          brc.RGDESC,
					MAINBR:          brc.MAINBR,
					MBDESC:          brc.MBDESC,
					BRANCH:          brc.BRANCH,
					BRDESC:          brc.BRDESC,
					CreatedAt:       &timeNow,
					UpdatedAt:       &timeNow,
				}

				taskToday, err := tasklists.repository.StoreTasklistDaily(reqTaskDaily, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}

				reqNotif := &models.TasklistNotif{
					// TaskID:     request.ID,
					TaskID:     int64(taskToday.ID),
					Tanggal:    &timeNow,
					Keterangan: "Ada 1 tasklist baru yang harus dikerjakan",
					Status:     0,
					Jenis:      "Tasklist",
					Receiver:   brc.PERNR,
					// Uker:       brc.Uker,
					Uker: brc.BRDESC,
					// Uker: request.Uker,
				}

				_, err = tasklists.repository.StoreNotif(reqNotif, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}
			}

			tx.Commit()
			return true, err
		}

		if dataTasklist.Period == "Custom" && dataTaskType.Range == "Quarter" && today.After(startDay) && today.Before(endDay) {
			brcRequest := &models.GetBRCRequest{
				// Branch: request.Branch,
				Uker: request.Uker,
			}

			brcList, err := tasklists.repository.GetDataBRC(brcRequest)
			if err != nil {
				tasklists.logger.Zap.Error(err)
				return response, err
			}

			for _, brc := range brcList {
				reqTaskDaily := &models.TasklistsToday{
					TasklistID:      dataTasklist.ID,
					ActivityID:      dataTasklist.ActivityID,
					ProductID:       dataTasklist.ProductID,
					Product:         dataTasklist.ProductName,
					RiskIssueID:     dataTasklist.RiskIssueID,
					RiskIssue:       dataTasklist.RiskIssue,
					RiskIndicatorID: dataTasklist.RiskIndicatorID,
					RiskIndicator:   dataTasklist.RiskIndicator,
					StartDate:       dataTasklist.StartDate,
					EndDate:         dataTasklist.EndDate,
					Status:          dataTasklist.Status,
					TaskType:        dataTasklist.TaskType,
					TaskTypeName:    dataTasklist.TaskTypeName,
					Kegiatan:        dataTasklist.Kegiatan,
					Period:          dataTasklist.Period,
					Sample:          dataTasklist.Sample,
					Progres:         0,
					Persentase:      0,
					MakerID:         dataTasklist.MakerID,
					PERNR:           brc.PERNR,
					Assigned:        brc.SNAME,
					REGION:          brc.REGION,
					RGDESC:          brc.RGDESC,
					MAINBR:          brc.MAINBR,
					MBDESC:          brc.MBDESC,
					BRANCH:          brc.BRANCH,
					BRDESC:          brc.BRDESC,
					CreatedAt:       &timeNow,
					UpdatedAt:       &timeNow,
				}

				taskToday, err := tasklists.repository.StoreTasklistDaily(reqTaskDaily, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}

				reqNotif := &models.TasklistNotif{
					// TaskID:     request.ID,
					TaskID:     int64(taskToday.ID),
					Tanggal:    &timeNow,
					Keterangan: "Ada 1 tasklist baru yang harus dikerjakan",
					Status:     0,
					Jenis:      "Tasklist",
					Receiver:   brc.PERNR,
					Uker:       brc.BRDESC,
					// Uker: request.Branch,
					// Uker: request.Uker,
				}

				_, err = tasklists.repository.StoreNotif(reqNotif, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}
			}

			tx.Commit()
			return true, err
		}

		if dataTasklist.Period == "Custom" && dataTaskType.Range == "Semester" && today.After(startDay) && today.Before(endDay) {
			brcRequest := &models.GetBRCRequest{
				// Branch: request.Branch,
				Uker: request.Uker,
			}

			brcList, err := tasklists.repository.GetDataBRC(brcRequest)
			if err != nil {
				tasklists.logger.Zap.Error(err)
				return response, err
			}

			for _, brc := range brcList {
				reqTaskDaily := &models.TasklistsToday{
					TasklistID:      dataTasklist.ID,
					ActivityID:      dataTasklist.ActivityID,
					ProductID:       dataTasklist.ProductID,
					Product:         dataTasklist.ProductName,
					RiskIssueID:     dataTasklist.RiskIssueID,
					RiskIssue:       dataTasklist.RiskIssue,
					RiskIndicatorID: dataTasklist.RiskIndicatorID,
					RiskIndicator:   dataTasklist.RiskIndicator,
					StartDate:       dataTasklist.StartDate,
					EndDate:         dataTasklist.EndDate,
					Status:          dataTasklist.Status,
					TaskType:        dataTasklist.TaskType,
					TaskTypeName:    dataTasklist.TaskTypeName,
					Kegiatan:        dataTasklist.Kegiatan,
					Period:          dataTasklist.Period,
					Sample:          dataTasklist.Sample,
					Progres:         0,
					Persentase:      0,
					MakerID:         dataTasklist.MakerID,
					PERNR:           brc.PERNR,
					Assigned:        brc.SNAME,
					REGION:          brc.REGION,
					RGDESC:          brc.RGDESC,
					MAINBR:          brc.MAINBR,
					MBDESC:          brc.MBDESC,
					BRANCH:          brc.BRANCH,
					BRDESC:          brc.BRDESC,
					CreatedAt:       &timeNow,
					UpdatedAt:       &timeNow,
				}

				taskToday, err := tasklists.repository.StoreTasklistDaily(reqTaskDaily, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}

				reqNotif := &models.TasklistNotif{
					// TaskID:     request.ID,
					TaskID:     int64(taskToday.ID),
					Tanggal:    &timeNow,
					Keterangan: "Ada 1 tasklist baru yang harus dikerjakan",
					Status:     0,
					Jenis:      "Tasklist",
					Receiver:   brc.PERNR,
					// Uker:       brc.Uker,
					// Uker: request.Branch,
					Uker: brc.BRDESC,
					// Uker: request.Uker,
				}

				_, err = tasklists.repository.StoreNotif(reqNotif, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}
			}

			tx.Commit()
			return true, err
		}

		if dataTasklist.Period == "Custom" && startDay.Format("2006-01-02") == today.Format("2006-01-02") {

			brcRequest := &models.GetBRCRequest{
				// Branch: request.Branch,
				Uker: request.Uker,
			}

			brcList, err := tasklists.repository.GetDataBRC(brcRequest)
			if err != nil {
				tasklists.logger.Zap.Error(err)
				return response, err
			}

			for _, brc := range brcList {
				reqTaskDaily := &models.TasklistsToday{
					TasklistID:      dataTasklist.ID,
					ActivityID:      dataTasklist.ActivityID,
					ProductID:       dataTasklist.ProductID,
					Product:         dataTasklist.ProductName,
					RiskIssueID:     dataTasklist.RiskIssueID,
					RiskIssue:       dataTasklist.RiskIssue,
					RiskIndicatorID: dataTasklist.RiskIndicatorID,
					RiskIndicator:   dataTasklist.RiskIndicator,
					StartDate:       dataTasklist.StartDate,
					EndDate:         dataTasklist.EndDate,
					Status:          dataTasklist.Status,
					TaskType:        dataTasklist.TaskType,
					TaskTypeName:    dataTasklist.TaskTypeName,
					Kegiatan:        dataTasklist.Kegiatan,
					Period:          dataTasklist.Period,
					Sample:          dataTasklist.Sample,
					Progres:         0,
					Persentase:      0,
					MakerID:         dataTasklist.MakerID,
					PERNR:           brc.PERNR,
					Assigned:        brc.SNAME,
					REGION:          brc.REGION,
					RGDESC:          brc.RGDESC,
					MAINBR:          brc.MAINBR,
					MBDESC:          brc.MBDESC,
					BRANCH:          brc.BRANCH,
					BRDESC:          brc.BRDESC,
					CreatedAt:       &timeNow,
					UpdatedAt:       &timeNow,
				}

				taskToday, err := tasklists.repository.StoreTasklistDaily(reqTaskDaily, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}

				reqNotif := &models.TasklistNotif{
					// TaskID:     request.ID,
					TaskID:     int64(taskToday.ID),
					Tanggal:    &timeNow,
					Keterangan: "Ada 1 tasklist baru yang harus dikerjakan",
					Status:     0,
					Jenis:      "Tasklist",
					Receiver:   brc.PERNR,
					// Uker:       brc.Uker,
					// Uker: request.Uker,
					// Uker: request.Branch,
					Uker: brc.BRDESC,
				}

				_, err = tasklists.repository.StoreNotif(reqNotif, tx)
				if err != nil {
					tx.Rollback()
					tasklists.logger.Zap.Error(err)
					return false, err
				}
			}

			tx.Commit()
			return true, err
		}
	} else if request.ApprovalStatus == "Minta Persetujuan Approval" {
		dataTasklistReceiver, _, err := tasklists.GetOne(request.ID)

		if err != nil {
			tasklists.logger.Zap.Error(err)
			tx.Rollback()
			return false, err
		}

		reqNotif := &models.TasklistNotif{
			TaskID:     request.ID,
			Tanggal:    &timeNow,
			Keterangan: "Ada 1 tasklist baru yang harus diapprove",
			Status:     0,
			Jenis:      "Validasi Tasklist",
			Receiver:   dataTasklistReceiver.Approval,
			// Uker:       request.Uker,
			Uker: request.Branch,
		}

		_, err = tasklists.repository.StoreNotif(reqNotif, tx)

		if err != nil {
			tx.Rollback()
			tasklists.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

func (tasklists TasklistsService) Validation(request *models.TasklistsValidation) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := tasklists.db.DB.Begin()

	UpdateValidation := &models.TasklistsValidation{
		ID:               request.ID,
		ValidationStatus: request.ValidationStatus,
		UpdatedAt:        &timeNow,
	}

	_, err = tasklists.repository.Validation(UpdateValidation, tx)
	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	reqNotif := &models.TasklistNotif{
		TaskID:     request.ID,
		Tanggal:    &timeNow,
		Keterangan: "Ada 1 tasklist baru yang perlu diapprove",
		Status:     0,
		Jenis:      "Validasi Tasklist",
		Receiver:   request.Receiver,
		Uker:       request.Branch,
		// Uker:       request.Uker,
	}

	_, err = tasklists.repository.StoreNotif(reqNotif, tx)
	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

func (tasklists TasklistsService) CountTask(request models.TasklistCountRequest) (response models.TasklistsCountResponse, err error) {
	countTaskDone, err := tasklists.repository.CountTaskDone(request)

	if err != nil {
		tasklists.logger.Zap.Error(err)
		return response, err
	}

	countTask, err := tasklists.repository.CountTask(request)

	if err != nil {
		tasklists.logger.Zap.Error(err)
		return response, err
	}

	responses := models.TasklistsCountResponse{
		Total: countTask.Total - countTaskDone.Total,
	}

	return responses, err
}

func (tasklists TasklistsService) StoreDone(request models.TasklistsDoneHistoryRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := tasklists.db.DB.Begin()

	reqTasklistsDone := &models.TasklistsDoneHistory{
		TasklistID: request.TasklistID,
		PERNR:      request.PERNR,
		CreatedAt:  &timeNow,
	}

	_, err = tasklists.doneHistory.Store(reqTasklistsDone, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

func (tasklists TasklistsService) GetDone(request models.TasklistsDoneHistoryCheckRequest) (responses models.TasklistsDoneHistoryResponse, status bool, err error) {
	dataTasklistsDone, err := tasklists.doneHistory.Get(request)

	if dataTasklistsDone.TasklistID != 0 {
		fmt.Println("bukan 0")

		responses = models.TasklistsDoneHistoryResponse{
			PERNR:      dataTasklistsDone.PERNR,
			TasklistID: dataTasklistsDone.TasklistID,
			CreatedAt:  dataTasklistsDone.CreatedAt,
		}

		return responses, true, err
	}

	return responses, false, err
}

func (tasklists TasklistsService) LeadAutocompleteVal(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error) {
	return tasklists.repository.LeadAutocompleteVal(request)
}

func (tasklists TasklistsService) LeadAutocompleteApr(request models.LeadAutocompleteRequest) (responses []models.LeadAutocomplete, err error) {
	return tasklists.repository.LeadAutocompleteApr(request)
}

func (tasklists TasklistsService) GetDataAnomali(request models.TasklistDataAnomaliRequest) (responses []models.TasklistDataAnomaliResponse, err error) {
	tasklistsDataAnomali, err := tasklists.tasklistsAnomaliKRID.GetDataAnomali(request)
	if err != nil {
		tasklists.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range tasklistsDataAnomali {
		responses = append(responses, models.TasklistDataAnomaliResponse{
			Object: response.Object,
		})
	}

	return responses, err
}

func (tasklists TasklistsService) GetDataVerifikasi(request models.DataVerifikasiRequest) (responses []models.DataVerifikasiResponse, err error) {
	return tasklists.repository.GetDataVerifikasi(request)
}

func (tasklists TasklistsService) GetLampiranIndicator(request *models.LampiranIndikatorCheck) (response models.LampiranIndikatorResponse, status bool, err error) {
	dataLampiran, err := tasklists.repository.GetLampiranIndicator(request)

	// if dataLampiran.ID != 0 {
	if err == nil {
		fmt.Println("bukan 0")

		response = models.LampiranIndikatorResponse{
			ID:              dataLampiran.ID,
			RiskIssueID:     dataLampiran.RiskIssueID,
			RiskIndicatorID: dataLampiran.RiskIndicatorID,
			NamaTable:       dataLampiran.NamaTable,
			JumlahKolom:     dataLampiran.JumlahKolom,
		}

		return response, true, err
	}

	return response, false, err
}

func (tasklists TasklistsService) DownloadLampiranIndikatorTemplate(request *models.LampiranIndikatorCheck) (response models.AnomaliHeaderResponse, err error) {
	return tasklists.repository.DownloadLampiranIndikatorTemplate(request)
}

func (tasklists TasklistsService) InsertTasklistRejected(request *models.TasklistsRejected) (err error) {
	tx := tasklists.db.DB.Begin()

	reqTasklistRejected := &models.TasklistsRejected{
		TasklistID: request.TasklistID,
		Notes:      request.Notes,
		Status:     "Not Done",
	}

	err = tasklists.repository.InsertTasklistRejected(reqTasklistRejected, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return err
	}

	tx.Commit()
	return err
}

func (tasklists TasklistsService) GetTasklist(request *models.TasklistCheckRequest) (response models.TasklistCheckResponse, err error) {
	dataTasklists, err := tasklists.daily.GetTasklist(request)

	if err != nil {
		fmt.Println(err)
	}

	response = models.TasklistCheckResponse{
		ID:              dataTasklists.ID,
		ActivityID:      dataTasklists.ActivityID,
		ProductID:       dataTasklists.ProductID,
		ProductName:     dataTasklists.ProductName,
		RiskIssueID:     dataTasklists.RiskIssueID,
		RiskIssue:       dataTasklists.RiskIssue,
		RiskIndicatorID: dataTasklists.RiskIndicatorID,
		RiskIndicator:   dataTasklists.RiskIndicator,
		StartDate:       dataTasklists.StartDate,
		EndDate:         dataTasklists.EndDate,
		TaskType:        dataTasklists.TaskType,
		TaskTypeName:    dataTasklists.TaskTypeName,
		SumberData:      dataTasklists.SumberData,
		RAP:             dataTasklists.RAP,
		Validation:      dataTasklists.Validation,
		ValidationName:  dataTasklists.ValidationName,
		Approval:        dataTasklists.Approval,
		ApprovalName:    dataTasklists.ApprovalName,
		Sample:          dataTasklists.Sample,
		REGION:          dataTasklists.REGION,
		RGDESC:          dataTasklists.RGDESC,
		MAINBR:          dataTasklists.MAINBR,
		MBDESC:          dataTasklists.MBDESC,
		BRANCH:          dataTasklists.BRANCH,
		BRDESC:          dataTasklists.BRDESC,
		PERNR:           dataTasklists.PERNR,
		ApprovalStatus:  dataTasklists.ApprovalStatus,
		Status:          dataTasklists.Status,
		MakerID:         dataTasklists.MakerID,
	}

	return response, err
}

func (tasklists TasklistsService) StoreDaily(request *models.TasklistDailyStore) (response *models.TasklistDailyStore, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := tasklists.db.DB.Begin()

	reqTasklists := &models.TasklistDailyStore{
		ActivityID:      request.ActivityID,
		ProductID:       request.ProductID,
		ProductName:     request.ProductName,
		RiskIssueID:     request.RiskIssueID,
		RiskIssue:       request.RiskIssue,
		RiskIndicatorID: request.RiskIndicatorID,
		RiskIndicator:   request.RiskIndicator,
		TaskType:        request.TaskType,
		TaskTypeName:    request.TaskTypeName,
		SumberData:      request.SumberData,
		RAP:             request.RAP,
		Progres:         0,
		Sample:          request.Sample,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
		Approval:        request.Approval,
		ApprovalName:    request.ApprovalName,
		Validation:      request.Validation,
		ValidationName:  request.ValidationName,
		ApprovalStatus:  request.ApprovalStatus,
		REGION:          request.REGION,
		RGDESC:          request.RGDESC,
		MBDESC:          request.MBDESC,
		MAINBR:          request.MAINBR,
		BRANCH:          request.BRANCH,
		BRDESC:          request.BRDESC,
		Status:          request.Status,
		MakerID:         request.MakerID,
		CreatedAt:       &timeNow,
		UpdatedAt:       &timeNow,
	}

	response, err = tasklists.daily.StoreDaily(reqTasklists, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return response, err
	}

	tx.Commit()
	return response, err
}

func (tasklists TasklistsService) UpdateDaily(request *models.ProgresUpdateRequest) (response *models.ProgresUpdateRequest, err error) {
	tx := tasklists.db.DB.Begin()

	updateTasklists := &models.ProgresUpdateRequest{
		ID:      request.ID,
		Progres: request.Progres,
	}
	response, err = tasklists.daily.UpdateDaily(updateTasklists, tx)

	if err != nil {
		tx.Rollback()
		tasklists.logger.Zap.Error(err)
		return response, err
	}

	tx.Commit()
	return response, err
}

func (tasklists TasklistsService) GetAnomaliHeader(request *models.AnomaliHeader) (response models.AnomaliHeaderResponse, err error) {
	return tasklists.repository.GetAnomaliHeader(request)
}

func (tasklists TasklistsService) GetAnomaliValue(request *models.AnomaliValue) (response []models.AnomaliValueResponse, err error) {
	return tasklists.repository.GetAnomaliValue(request)
}

func (tasklists TasklistsService) GetFirstLampiran(request *models.GetFirstLampiranRequest) (response models.GetFirstLampiranResponse, err error) {
	response, err = tasklists.repository.GetFirstLampiran(request)

	if err != nil {
		tasklists.logger.Zap.Error(err)
		return response, err
	}

	var mini_link fileModel.FileManagerResponseUrl

	mini_link, err = tasklists.filemanager.GetFile(fileModel.FileManagerRequest{
		Subdir:   response.Path,
		Filename: response.Filename,
	})

	if err != nil {
		tasklists.logger.Zap.Error(err)
	}

	response = models.GetFirstLampiranResponse{
		Filename: response.Filename,
		Path:     mini_link.MinioPath,
	}

	return response, err
}
