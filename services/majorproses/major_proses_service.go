package majorproses

import (
	"riskmanagement/lib"
	models "riskmanagement/models/majorproses"
	repository "riskmanagement/repository/majorproses"

	"gitlab.com/golang-package-library/logger"
)

type MajorProsesDefinition interface {
	GetAll() (responses []models.MajorProsesResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.MajorProsesResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.MajorProsesResponse, err error)
	Store(request *models.MajorProsesRequest) (err error)
	Update(request *models.MajorProsesRequest) (err error)
	Delete(id int64) (err error)
	GetKodeMajorProses(request models.KodeMegaProses) (responses []models.KodeMajorProses, err error)
	GetMajorByMegaProses(request models.KodeMegaProses) (responses []models.MajorProsesResponse, err error)
}

type MajorProsesService struct {
	logger     logger.Logger
	repository repository.MajorProsesDefinition
}

func NewMajorProsesService(
	logger logger.Logger,
	repository repository.MajorProsesDefinition,
) MajorProsesDefinition {
	return MajorProsesService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements MajorProsesDefinition
func (majorproses MajorProsesService) Delete(id int64) (err error) {
	return majorproses.repository.Delete(id)
}

// GetAll implements MajorProsesDefinition
func (majorproses MajorProsesService) GetAll() (responses []models.MajorProsesResponse, err error) {
	return majorproses.repository.GetAll()
}

// GetAllWithPaginate implements MajorProsesDefinition
func (major MajorProsesService) GetAllWithPaginate(request models.Paginate) (responses []models.MajorProsesResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	datamajor, totalRows, totalData, err := major.repository.GetAllWithPaginate(&request)

	if err != nil {
		major.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		major.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range datamajor {
		responses = append(responses, models.MajorProsesResponse{
			ID:              response.ID,
			IDMegaProses:    response.IDMegaProses,
			MegaProsesName:  response.MegaProsesName,
			KodeMajorProses: response.KodeMajorProses,
			MajorProses:     response.MajorProses,
			Deskripsi:       response.Deskripsi,
			Status:          response.Status,
			CreatedAt:       response.CreatedAt,
			UpdatedAt:       response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetKodeMajorProses implements MajorProsesDefinition
func (majorproses MajorProsesService) GetKodeMajorProses(request models.KodeMegaProses) (responses []models.KodeMajorProses, err error) {
	dataMAJORPROSES, err := majorproses.repository.GetKodeMajorProses(request)

	if err != nil {
		majorproses.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataMAJORPROSES {
		responses = append(responses, models.KodeMajorProses{
			KodeMegaProses:  request.KodeMegaProses,
			KodeMajorProses: response.KodeMajorProses,
		})
	}

	return responses, err
}

// GetOne implements MajorProsesDefinition
func (majorproses MajorProsesService) GetOne(id int64) (responses models.MajorProsesResponse, err error) {
	return majorproses.repository.GetOne(id)
}

// Store implements MajorProsesDefinition
func (majorproses MajorProsesService) Store(request *models.MajorProsesRequest) (err error) {
	status, err := majorproses.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements MajorProsesDefinition
func (majorproses MajorProsesService) Update(request *models.MajorProsesRequest) (err error) {
	status, err := majorproses.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetMajorByMegaProses implements MajorProsesDefinition
func (majorproses MajorProsesService) GetMajorByMegaProses(request models.KodeMegaProses) (responses []models.MajorProsesResponse, err error) {
	dataMajor, err := majorproses.repository.GetMajorByMegaProses(&request)
	if err != nil {
		majorproses.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataMajor {
		responses = append(responses, models.MajorProsesResponse{
			ID:              response.ID,
			IDMegaProses:    response.IDMegaProses,
			KodeMajorProses: response.KodeMajorProses,
			MajorProses:     response.MajorProses,
			Deskripsi:       response.Deskripsi,
			Status:          response.Status,
			CreatedAt:       response.CreatedAt,
			UpdatedAt:       response.UpdatedAt,
		})
	}

	return responses, err
}
