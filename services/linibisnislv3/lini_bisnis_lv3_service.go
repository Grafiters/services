package linibisnislv3

import (
	"riskmanagement/lib"
	models "riskmanagement/models/linibisnislv3"
	repository "riskmanagement/repository/linibisnislv3"

	"gitlab.com/golang-package-library/logger"
)

type LiniBisnisLv3Definition interface {
	GetAll() (responses []models.LiniBisnisLv3Response, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.LiniBisnisLv3Response, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.LiniBisnisLv3Response, err error)
	Store(request *models.LiniBisnisLv3Request) (err error)
	Update(request *models.LiniBisnisLv3Request) (err error)
	Delete(id int64) (err error)
	GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error)
	GetLBByID(request models.KodeLB2) (responses []models.LiniBisnisLv3Response, err error)
}

type LiniBisnisLv3Service struct {
	logger     logger.Logger
	repository repository.LiniBisnisLv3Definition
}

func NewLiniBisnisLv3Service(
	logger logger.Logger,
	repository repository.LiniBisnisLv3Definition,
) LiniBisnisLv3Definition {
	return LiniBisnisLv3Service{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) Delete(id int64) (err error) {
	return lb3.repository.Delete(id)
}

// GetAllWithPaginate implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) GetAllWithPaginate(request models.Paginate) (responses []models.LiniBisnisLv3Response, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataLB3, totalRows, totalData, err := lb3.repository.GetAllWithPaginate(&request)

	if err != nil {
		lb3.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		lb3.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataLB3 {
		responses = append(responses, models.LiniBisnisLv3Response{
			ID:                response.ID,
			IDLiniBisnisLv2:   response.IDLiniBisnisLv2,
			KodeLiniBisnisLv3: response.KodeLiniBisnisLv3,
			LiniBisnisLv3:     response.LiniBisnisLv3,
			Deskripsi:         response.Deskripsi,
			Status:            response.Status,
			CreatedAt:         response.CreatedAt,
			UpdatedAt:         response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetAll implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) GetAll() (responses []models.LiniBisnisLv3Response, err error) {
	return lb3.repository.GetAll()
}

// GetKodeLiniBisnis implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) GetKodeLiniBisnis() (responses []models.KodeLiniBisnis, err error) {
	dataLB3, err := lb3.repository.GetKodeLiniBisnis()

	if err != nil {
		lb3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataLB3 {
		responses = append(responses, models.KodeLiniBisnis{
			KodeLiniBisnis: response.KodeLiniBisnis,
		})
	}

	return responses, err
}

// GetOne implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) GetOne(id int64) (responses models.LiniBisnisLv3Response, err error) {
	return lb3.repository.GetOne(id)
}

// Store implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) Store(request *models.LiniBisnisLv3Request) (err error) {
	status, err := lb3.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) Update(request *models.LiniBisnisLv3Request) (err error) {
	status, err := lb3.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetLBByID implements LiniBisnisLv3Definition
func (lb3 LiniBisnisLv3Service) GetLBByID(request models.KodeLB2) (responses []models.LiniBisnisLv3Response, err error) {
	dataLB3, err := lb3.repository.GetLBByID(&request)

	if err != nil {
		lb3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataLB3 {
		responses = append(responses, models.LiniBisnisLv3Response{
			ID:                response.ID,
			IDLiniBisnisLv2:   response.IDLiniBisnisLv2,
			KodeLiniBisnisLv3: response.KodeLiniBisnisLv3,
			LiniBisnisLv3:     response.LiniBisnisLv3,
			Deskripsi:         response.Deskripsi,
			Status:            response.Status,
			CreatedAt:         response.CreatedAt,
			UpdatedAt:         response.UpdatedAt,
		})
	}

	return responses, err
}
