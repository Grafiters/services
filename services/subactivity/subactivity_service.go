package subactivity

import (
	"riskmanagement/lib"
	models "riskmanagement/models/subactivity"
	repository "riskmanagement/repository/subactivity"

	"gitlab.com/golang-package-library/logger"
)

type SubActivityDefinition interface {
	GetAll() (responses []models.SubActivityResponses, err error)
	GetAllWithPagination(requests models.Paginate) (responses []models.SubActivityResponses, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.SubActivityResponse, err error)
	GetLastID(id int64) (responses []models.SubActivityResponse, err error)
	Store(request *models.SubActivityRequest) (err error)
	Update(request *models.SubActivityRequest) (err error)
	Delete(id int64) (err error)
	GetKodeSubActivity(request models.KodeSubActivityRequest) (responses []models.KodeSubActivityResponse, err error)
}

type SubActivityService struct {
	logger     logger.Logger
	repository repository.SubActivityDefinition
}

func (subactivity SubActivityService) GetAll() (responses []models.SubActivityResponses, err error) {
	return subactivity.repository.GetAll()
}

// GetLastID implements SubActivityDefinition
func (subactivity SubActivityService) GetLastID(id int64) (responses []models.SubActivityResponse, err error) {
	return subactivity.repository.GetLastID(id)
}

func NewSubActivityService(logger logger.Logger, repository repository.SubActivityDefinition) SubActivityDefinition {
	return SubActivityService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements SubAvtivityDefinition
func (subactivity SubActivityService) Delete(id int64) (err error) {
	return subactivity.repository.Delete(id)
}

// GetAll implements SubAvtivityDefinition
//
//	func (subactivity SubActivityService) GetAll() (responses []models.SubActivityResponse, err error) {
//		return subactivity.repository.GetAll()
//	}
func (subactivity SubActivityService) GetAllWithPagination(request models.Paginate) (responses []models.SubActivityResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataSubActivitas, totalRows, totalData, err := subactivity.repository.GetAllWithPagination(&request)
	if err != nil {
		subactivity.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		subactivity.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataSubActivitas {
		responses = append(responses, models.SubActivityResponses{
			ID:              response.ID,
			ActivityID:      response.ActivityID,
			ActivityName:    response.ActivityName,
			KodeSubActivity: response.KodeSubActivity,
			NameSubActivity: response.NameSubActivity,
			CreatedAt:       response.CreatedAt,
			UpdatedAt:       response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements SubAvtivityDefinition
func (subactivity SubActivityService) GetOne(id int64) (responses models.SubActivityResponse, err error) {
	return subactivity.repository.GetOne(id)
}

// Store implements SubAvtivityDefinition
func (subactivity SubActivityService) Store(request *models.SubActivityRequest) (err error) {
	status, err := subactivity.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements SubAvtivityDefinition
func (subactivity SubActivityService) Update(request *models.SubActivityRequest) (err error) {
	status, err := subactivity.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetKodeSubActivity implements SubActivityDefinition
func (subActivity SubActivityService) GetKodeSubActivity(request models.KodeSubActivityRequest) (responses []models.KodeSubActivityResponse, err error) {
	dataSubs, err := subActivity.repository.GetKodeSubActivity(request)

	if err != nil {
		subActivity.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataSubs {
		responses = append(responses, models.KodeSubActivityResponse{
			KodeActivity:    request.KodeActivity,
			KodeSubActivity: response.KodeSubActivity,
		})
	}

	return responses, err
}
