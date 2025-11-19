package activity

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/activity"
	repository "riskmanagement/repository/activity"

	"gitlab.com/golang-package-library/logger"
)

type ActivityDefinition interface {
	GetAll() (responses []models.ActivityResponse, err error)
	GetAllWithPagination(request models.Paginate) (response []models.ActivityResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.ActivityResponse, err error)
	Store(request *models.ActivityRequest) (err error)
	Update(request *models.ActivityRequest) (err error)
	Delete(id int64) (err error)
	GetKodeActivity() (responses []models.KodeActivity, err error)
}

type ActivityService struct {
	logger     logger.Logger
	repository repository.ActivityDefinition
}

func NewActivityService(
	logger logger.Logger,
	repository repository.ActivityDefinition,
) ActivityDefinition {
	return ActivityService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements ActivityDefinition
func (activity ActivityService) Delete(id int64) (err error) {
	return activity.repository.Delete(id)
}

// GetAll implements ActivityDefinition
func (activity ActivityService) GetAll() (responses []models.ActivityResponse, err error) {
	return activity.repository.GetAll()
}

// GetAllWithPagination implements ActivityDefinition
func (activity ActivityService) GetAllWithPagination(request models.Paginate) (response []models.ActivityResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataActivitas, totalRows, totalData, err := activity.repository.GetAllWithPagination(&request)
	if err != nil {
		activity.logger.Zap.Error(err)
		return response, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		activity.logger.Zap.Error(err)
		return response, pagination, err
	}

	for _, aktifitas := range dataActivitas {
		response = append(response, models.ActivityResponse{
			ID:           aktifitas.ID,
			KodeActivity: aktifitas.KodeActivity,
			Name:         aktifitas.Name,
			CreateAt:     aktifitas.CreateAt,
			UpdateAt:     aktifitas.UpdateAt,
		})
	}
	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return response, pagination, err
}

// GetOne implements ActivityDefinition
func (activity ActivityService) GetOne(id int64) (responses models.ActivityResponse, err error) {
	return activity.repository.GetOne(id)
}

// Store implements ActivityDefinition
func (activity ActivityService) Store(request *models.ActivityRequest) (err error) {
	fmt.Println("service =", request)
	status, err := activity.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements ActivityDefinition
func (activity ActivityService) Update(request *models.ActivityRequest) (err error) {
	status, err := activity.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// getKodeActivity implements ActivityDefinition
func (activity ActivityService) GetKodeActivity() (responses []models.KodeActivity, err error) {
	dataActivity, err := activity.repository.GetKodeActivity()

	if err != nil {
		activity.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataActivity {
		responses = append(responses, models.KodeActivity{
			KodeActivity: response.KodeActivity,
		})
	}

	return responses, err
}
