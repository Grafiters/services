package linibisnislv1

import (
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv1"
	repository "riskmanagement/repository/linibisnislv1"

	"gitlab.com/golang-package-library/logger"
)

type LiniBisnisLv1Definition interface {
	GetAll() (responses []models.LiniBisnisLv1Response, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.LiniBisnisLv1Response, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.LiniBisnisLv1Response, err error)
	Store(request *models.LiniBisnisLv1Request) (err error)
	Update(request *models.LiniBisnisLv1Request) (err error)
	Delete(id int64) (err error)
	GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error)
}

type LiniBisnisLv1Service struct {
	logger     logger.Logger
	repository repository.LiniBisnisLv1Definition
}

func NewLiniBisnisLv1Service(
	logger logger.Logger,
	repository repository.LiniBisnisLv1Definition,
) LiniBisnisLv1Definition {
	return LiniBisnisLv1Service{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Service) Delete(id int64) (err error) {
	return lb1.repository.Delete(id)
}

// GetAllWithPaginate implements EventTypeLv1Definition
func (lb1 LiniBisnisLv1Service) GetAllWithPaginate(request models.Paginate) (responses []models.LiniBisnisLv1Response, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataLB1, totalRows, totalData, err := lb1.repository.GetAllWithPaginate(&request)

	if err != nil {
		lb1.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		lb1.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataLB1 {
		responses = append(responses, models.LiniBisnisLv1Response{
			ID:             response.ID,
			KodeLiniBisnis: response.KodeLiniBisnis,
			LiniBisnis1:    response.LiniBisnis1,
			Deskripsi:      response.Deskripsi,
			Status:         response.Status,
			CreatedAt:      response.CreatedAt,
			UpdatedAt:      response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetAll implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Service) GetAll() (responses []models.LiniBisnisLv1Response, err error) {
	return lb1.repository.GetAll()
}

// GetKodeLiniBisnis implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Service) GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error) {
	dataLB1, err := lb1.repository.GetKodeLiniBisnis()

	if err != nil {
		lb1.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataLB1 {
		responses = append(responses, models.KodeLiniBisnis{
			KodeLiniBisnis: response.KodeLiniBisnis,
		})
	}

	return responses, err
}

// GetOne implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Service) GetOne(id int64) (responses models.LiniBisnisLv1Response, err error) {
	return lb1.repository.GetOne(id)
}

// Store implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Service) Store(request *models.LiniBisnisLv1Request) (err error) {
	status, err := lb1.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements LiniBisnisLv1Definition
func (lb1 LiniBisnisLv1Service) Update(request *models.LiniBisnisLv1Request) (err error) {
	status, err := lb1.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}
