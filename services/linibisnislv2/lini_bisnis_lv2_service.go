package linibisnislv2

import (
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv2"
	repository "riskmanagement/repository/linibisnislv2"

	"gitlab.com/golang-package-library/logger"
)

type LiniBisnisLv2Definition interface {
	GetAll() (responses []models.LiniBisnisLv2Response, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.LiniBisnisLv2Response, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.LiniBisnisLv2Response, err error)
	Store(request *models.LiniBisnisLv2Request) (err error)
	Update(request *models.LiniBisnisLv2Request) (err error)
	Delete(id int64) (err error)
	GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error)
	GetLBByID(request models.KodeLB1) (responses []models.LiniBisnisLv2Response, err error)
}

type LiniBisnisLv2Service struct {
	logger     logger.Logger
	repository repository.LiniBisnisLv2Definition
}

func NewLiniBisnisLv2Service(
	logger logger.Logger,
	repository repository.LiniBisnisLv2Definition,
) LiniBisnisLv2Definition {
	return LiniBisnisLv2Service{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Service) Delete(id int64) (err error) {
	return lb2.repository.Delete(id)
}

// GetAllWithPaginate implements EventTypeLv1Definition
func (lb2 LiniBisnisLv2Service) GetAllWithPaginate(request models.Paginate) (responses []models.LiniBisnisLv2Response, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataLB2, totalRows, totalData, err := lb2.repository.GetAllWithPaginate(&request)

	if err != nil {
		lb2.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		lb2.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataLB2 {
		responses = append(responses, models.LiniBisnisLv2Response{
			ID:                response.ID,
			IDLiniBisnisLv1:   response.IDLiniBisnisLv1,
			KodeLiniBisnisLv2: response.KodeLiniBisnisLv2,
			LiniBisnisLv2:     response.LiniBisnisLv2,
			Deskripsi:         response.Deskripsi,
			Status:            response.Status,
			CreatedAt:         response.CreatedAt,
			UpdatedAt:         response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetAll implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Service) GetAll() (responses []models.LiniBisnisLv2Response, err error) {
	return lb2.repository.GetAll()
}

// GetKodeLiniBisnis implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Service) GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error) {
	dataLB2, err := lb2.repository.GetKodeLiniBisnis()

	if err != nil {
		lb2.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataLB2 {
		responses = append(responses, models.KodeLiniBisnis{
			KodeLiniBisnis: response.KodeLiniBisnis,
		})
	}

	return responses, err
}

// GetOne implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Service) GetOne(id int64) (responses models.LiniBisnisLv2Response, err error) {
	return lb2.repository.GetOne(id)
}

// Store implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Service) Store(request *models.LiniBisnisLv2Request) (err error) {
	status, err := lb2.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Service) Update(request *models.LiniBisnisLv2Request) (err error) {
	status, err := lb2.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetLBByID implements LiniBisnisLv2Definition
func (lb2 LiniBisnisLv2Service) GetLBByID(request models.KodeLB1) (responses []models.LiniBisnisLv2Response, err error) {
	dataLB2, err := lb2.repository.GetLBByID(&request)

	if err != nil {
		lb2.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataLB2 {
		responses = append(responses, models.LiniBisnisLv2Response{
			ID:                response.ID,
			IDLiniBisnisLv1:   response.IDLiniBisnisLv1,
			KodeLiniBisnisLv2: response.KodeLiniBisnisLv2,
			LiniBisnisLv2:     response.LiniBisnisLv2,
			Deskripsi:         response.Deskripsi,
			Status:            response.Status,
			CreatedAt:         response.CreatedAt,
			UpdatedAt:         response.UpdatedAt,
		})
	}

	return responses, err
}
