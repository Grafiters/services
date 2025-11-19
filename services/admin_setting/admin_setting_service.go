package admin_setting

import (
	"errors"
	"riskmanagement/lib"
	models "riskmanagement/models/admin_setting"
	repository "riskmanagement/repository/admin_setting"
	"strings"

	"gitlab.com/golang-package-library/logger"
	"gitlab.com/golang-package-library/minio"
)

type AdminSettingDefinition interface {
	Show(request models.Paginate) (response []models.AdminSettingResponse, pagination lib.Pagination, err error)
	GetAll(request models.Paginate) (response []models.AdminSettingResponse, pagination lib.Pagination, err error)
	Store(request models.AdminSettingRequest) (status bool, err error)
	Update(request models.AdminSettingUpdateRequest) (status bool, err error)
	Delete(request *models.AdminSettingDelete) (status bool, err error)
	SearchTaskType(request models.KeywordRequest) (responses []models.TaskTypeResponse, pagination lib.Pagination, err error)
	SearchTaskTypeInput(request models.KeywordRequest) (responses []models.TaskTypeResponse, pagination lib.Pagination, err error)
	SearchTaskTypeInputByKegiatan(request models.KeywordRequest) (responses []models.TaskTypeResponse, pagination lib.Pagination, err error)
	GetOne(request models.TaskTypeRequestOne) (responses models.TaskTypeResponse, status bool, err error)
}

type AdminSettingService struct {
	db         lib.Database
	minio      minio.Minio
	logger     logger.Logger
	repository repository.AdminSettingDefinition
}

func NewAdminSettingService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	repository repository.AdminSettingDefinition,
) AdminSettingDefinition {
	return AdminSettingService{
		db:         db,
		minio:      minio,
		logger:     logger,
		repository: repository,
	}
}

func (setting AdminSettingService) Show(request models.Paginate) (response []models.AdminSettingResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	var dataSetting []models.AdminSetting
	var totalRows, totalData int

	dataSetting, totalRows, totalData, err = setting.repository.Show(&request)

	if err != nil {
		setting.logger.Zap.Error(err)
		return response, pagination, err
	}

	for _, set := range dataSetting {
		role, err := setting.repository.ShowRole(set.ID)

		if err != nil {
			setting.logger.Zap.Error(err)
			return response, pagination, err
		}

		response = append(response, models.AdminSettingResponse{
			ID:          set.ID,
			TaskType:    set.TaskType,
			Role:        role,
			Kegiatan:    set.Kegiatan,
			Range:       set.Range,
			Period:      set.Period,
			Upload:      set.Upload,
			TasklistMax: set.TasklistMax,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return response, pagination, err
}

func (setting AdminSettingService) GetAll(request models.Paginate) (response []models.AdminSettingResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	var dataSetting []models.TaskType
	var totalRows, totalData int

	dataSetting, totalRows, totalData, err = setting.repository.GetAll(&request)

	if err != nil {
		setting.logger.Zap.Error(err)
		return response, pagination, err
	}

	for _, set := range dataSetting {
		role, err := setting.repository.ShowRole(set.ID)

		if err != nil {
			setting.logger.Zap.Error(err)
			return response, pagination, err
		}

		response = append(response, models.AdminSettingResponse{
			ID:          set.ID,
			TaskType:    set.TaskType,
			Role:        role,
			Kegiatan:    set.Kegiatan,
			Range:       set.Range,
			Period:      set.Period,
			Upload:      set.Upload,
			TasklistMax: set.TasklistMax,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return response, pagination, err
}

func (setting AdminSettingService) SearchTaskTypeInput(request models.KeywordRequest) (responses []models.TaskTypeResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataTaskType, totalRows, totalData, err := setting.repository.SearchTaskTypeInput(&request)
	if err != nil {
		setting.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataTaskType {
		role, err := setting.repository.ShowRole(response.ID)

		if err != nil {
			setting.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.TaskTypeResponse{
			ID:          response.ID,
			TaskType:    response.TaskType,
			Kegiatan:    response.Kegiatan,
			Period:      response.Period,
			Range:       response.Range,
			Upload:      response.Upload,
			Role:        role,
			TasklistMax: response.TasklistMax,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (setting AdminSettingService) SearchTaskTypeInputByKegiatan(request models.KeywordRequest) (responses []models.TaskTypeResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataTaskType, totalRows, totalData, err := setting.repository.SearchTaskTypeInputByKegiatan(&request)
	if err != nil {
		setting.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataTaskType {
		role, err := setting.repository.ShowRole(response.ID)

		if err != nil {
			setting.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.TaskTypeResponse{
			ID:          response.ID,
			TaskType:    response.TaskType,
			Kegiatan:    response.Kegiatan,
			Period:      response.Period,
			Range:       response.Range,
			Upload:      response.Upload,
			Role:        role,
			TasklistMax: response.TasklistMax,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (setting AdminSettingService) SearchTaskType(request models.KeywordRequest) (responses []models.TaskTypeResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataTaskType, totalRows, totalData, err := setting.repository.SearchTaskType(&request)
	if err != nil {
		setting.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataTaskType {
		role, err := setting.repository.ShowRole(response.ID)

		if err != nil {
			setting.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.TaskTypeResponse{
			ID:          response.ID,
			TaskType:    response.TaskType,
			Kegiatan:    response.Kegiatan,
			Period:      response.Period,
			Range:       response.Range,
			Upload:      response.Upload,
			Role:        role,
			TasklistMax: response.TasklistMax,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (setting AdminSettingService) Store(request models.AdminSettingRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := setting.db.DB.Begin()

	reqCheckTaskType := &models.TaskTypeCheckRequest{
		TaskType: request.TaskType,
	}

	isAvailable, err := setting.repository.CheckAvailability(*reqCheckTaskType)

	if isAvailable.Total != 0 {
		err1 := errors.New("Nama jenis task sudah pernah digunakan!")
		setting.logger.Zap.Error("Duplicate data")
		return false, err1
	}

	reqSetting := &models.AdminSetting{
		TaskType:    request.TaskType,
		Kegiatan:    request.Kegiatan,
		Period:      request.Period,
		Range:       request.Range,
		Upload:      request.Upload,
		TasklistMax: request.TasklistMax,
		Status:      "Aktif",
		CreatedAt:   &timeNow,
		UpdatedAt:   &timeNow,
	}

	dataSetting, err := setting.repository.Store(reqSetting, tx)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	if len(request.Role) != 0 {
		for _, value := range request.Role {
			hilfms := strings.Split(value.HILFM, ",")
			stells := strings.Split(value.StellTX, ",")

			for _, hilfm := range hilfms {
				for _, stell := range stells {
					_, err := setting.repository.StoreRole(&models.AdminSettingRoleRequest{
						Orgeh:     value.Orgeh,
						KOSTL:     value.KOSTL,
						HILFM:     hilfm,
						TipeUker:  value.TipeUker,
						Role:      value.Role,
						Stell:     stell,
						JG:        value.JG,
						IDSetting: dataSetting.ID,
					}, tx)

					if err != nil {
						tx.Rollback()
						setting.logger.Zap.Error(err)
						return false, err
					}
				}
			}
		}
	}

	tx.Commit()
	return true, err
}

func (setting AdminSettingService) Update(request models.AdminSettingUpdateRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := setting.db.DB.Begin()

	updateSetting := &models.AdminSettingUpdate{
		ID:          request.ID,
		TaskType:    request.TaskType,
		Kegiatan:    request.Kegiatan,
		Period:      request.Period,
		Range:       request.Range,
		Upload:      request.Upload,
		TasklistMax: request.TasklistMax,
		UpdatedAt:   &timeNow,
	}
	_, err = setting.repository.Update(updateSetting, tx)

	if err != nil {
		tx.Rollback()
		setting.logger.Zap.Error(err)
		return false, err
	}

	deleteSetting := models.AdminSettingRole{
		IDSetting: request.ID,
	}
	err = setting.repository.DeleteRole(&deleteSetting, tx)

	if err != nil {
		tx.Rollback()
		setting.logger.Zap.Error(err)
		return false, err
	}

	if len(request.AdminSettingRole) != 0 {
		for _, value := range request.AdminSettingRole {
			_, err := setting.repository.StoreRole(&models.AdminSettingRoleRequest{
				Orgeh:     value.Orgeh,
				KOSTL:     value.KOSTL,
				HILFM:     value.HILFM,
				TipeUker:  value.TipeUker,
				Role:      value.Role,
				Stell:     value.StellTX,
				JG:        value.JG,
				IDSetting: request.ID,
			}, tx)

			if err != nil {
				tx.Rollback()
				setting.logger.Zap.Error(err)
				return false, err
			}
		}
	}

	tx.Commit()
	return true, err
}

func (setting AdminSettingService) Delete(request *models.AdminSettingDelete) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := setting.db.DB.Begin()

	deleteSetting := &models.AdminSettingDelete{
		ID:        request.ID,
		Status:    "Tidak Aktif",
		UpdatedAt: &timeNow,
	}

	err = setting.repository.Delete(deleteSetting, tx)

	if err != nil {
		tx.Rollback()
		setting.logger.Zap.Error(err)
		return false, err
	} else {
		tx.Commit()
		return true, err
	}
}

func (setting AdminSettingService) GetOne(request models.TaskTypeRequestOne) (responses models.TaskTypeResponse, status bool, err error) {
	dataTaskTypeRole, err := setting.repository.GetRole(request.ID)

	dataTaskType, err := setting.repository.GetOne(request.ID)

	if dataTaskType.ID != 0 {
		responses = models.TaskTypeResponse{
			ID:          dataTaskType.ID,
			TaskType:    dataTaskType.TaskType,
			Kegiatan:    dataTaskType.Kegiatan,
			Period:      dataTaskType.Period,
			Range:       dataTaskType.Range,
			Upload:      dataTaskType.Upload,
			TasklistMax: dataTaskType.TasklistMax,
			Role:        dataTaskTypeRole,
		}

		return responses, true, err
	}

	return responses, false, err
}
