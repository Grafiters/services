package megaproses

import (
	"riskmanagement/lib"
	models "riskmanagement/models/megaproses"
	repository "riskmanagement/repository/megaproses"

	"gitlab.com/golang-package-library/logger"
)

type MegaProsesDefinition interface {
	GetAll() (responses []models.MegaProsesResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.MegaProsesResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.MegaProsesResponse, err error)
	Store(request *models.MegaProsesRequest) (err error)
	Update(request *models.MegaProsesRequest) (err error)
	Delete(id int64) (err error)
	GetKodeMegaProses() (responses []models.KodeMegaProses, err error)
}

type MegaProsesService struct {
	logger     logger.Logger
	repository repository.MegaProsesDefinition
}

func NewMegaProsesService(
	logger logger.Logger,
	repository repository.MegaProsesDefinition,
) MegaProsesDefinition {
	return MegaProsesService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements MegaProsesDefinition
func (megaproses MegaProsesService) Delete(id int64) (err error) {
	return megaproses.repository.Delete(id)
}

// GetAll implements MegaProsesDefinition
func (megaproses MegaProsesService) GetAll() (responses []models.MegaProsesResponse, err error) {
	return megaproses.repository.GetAll()
}

// GetAllWithPaginate implements MegaProsesDefinition
func (mega MegaProsesService) GetAllWithPaginate(request models.Paginate) (responses []models.MegaProsesResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataMega, totalRows, totalData, err := mega.repository.GetAllWithPaginate(&request)

	if err != nil {
		mega.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		mega.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataMega {
		responses = append(responses, models.MegaProsesResponse{
			ID:             response.ID,
			KodeMegaProses: response.KodeMegaProses,
			MegaProses:     response.MegaProses,
			Deskripsi:      response.Deskripsi,
			Status:         response.Status,
			CreatedAt:      response.CreatedAt,
			UpdatedAt:      response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetKodeMegaProses implements MegaProsesDefinition
func (megaproses MegaProsesService) GetKodeMegaProses() (responses []models.KodeMegaProses, err error) {
	dataMEGAPROSES, err := megaproses.repository.GetKodeMegaProses()

	if err != nil {
		megaproses.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataMEGAPROSES {
		responses = append(responses, models.KodeMegaProses{
			KodeMegaProses: response.KodeMegaProses,
		})
	}

	return responses, err
}

// GetOne implements MegaProsesDefinition
func (megaproses MegaProsesService) GetOne(id int64) (responses models.MegaProsesResponse, err error) {
	return megaproses.repository.GetOne(id)
}

// Store implements MegaProsesDefinition
func (megaproses MegaProsesService) Store(request *models.MegaProsesRequest) (err error) {
	status, err := megaproses.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements MegaProsesDefinition
func (megaproses MegaProsesService) Update(request *models.MegaProsesRequest) (err error) {
	status, err := megaproses.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}
