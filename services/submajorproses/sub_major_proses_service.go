package submajorproses

import (
	"riskmanagement/lib"
	models "riskmanagement/models/submajorproses"
	repository "riskmanagement/repository/submajorproses"

	"gitlab.com/golang-package-library/logger"
)

type SubMajorProsesDefinition interface {
	GetAll() (responses []models.SubMajorProsesResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.SubMajorProsesResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.SubMajorProsesResponse, err error)
	Store(request *models.SubMajorProsesRequest) (responses bool, err error)
	Update(request *models.SubMajorProsesRequest) (err error)
	Delete(id int64) (err error)
	GetDataByID(request models.KodeMajor) (responses []models.SubMajorProsesResponse, err error)
}

type SubMajorProsesService struct {
	logger     logger.Logger
	repository repository.SubMajorProsesDefinition
}

func NewSubMajorProsesService(
	logger logger.Logger,
	repository repository.SubMajorProsesDefinition,
) SubMajorProsesDefinition {
	return SubMajorProsesService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesService) Delete(id int64) (err error) {
	return submajorproses.repository.Delete(id)
}

// GetAll implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesService) GetAll() (responses []models.SubMajorProsesResponse, err error) {
	return submajorproses.repository.GetAll()
}

// GetAllWithPaginate implements SubMajorProsesDefinition
func (subMajor SubMajorProsesService) GetAllWithPaginate(request models.Paginate) (responses []models.SubMajorProsesResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	datasubMajor, totalRows, totalData, err := subMajor.repository.GetAllWithPaginate(&request)

	if err != nil {
		subMajor.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		subMajor.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range datasubMajor {
		responses = append(responses, models.SubMajorProsesResponse{
			ID:                 response.ID,
			IDMajorProses:      response.IDMajorProses,
			MajorProses:        response.MajorProses,
			KodeSubMajorProses: response.KodeSubMajorProses,
			SubMajorProses:     response.SubMajorProses,
			Deskripsi:          response.Deskripsi,
			Status:             response.Status,
			CreatedAt:          response.CreatedAt,
			UpdatedAt:          response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetOne implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesService) GetOne(id int64) (responses models.SubMajorProsesResponse, err error) {
	return submajorproses.repository.GetOne(id)
}

// Store implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesService) Store(request *models.SubMajorProsesRequest) (responses bool, err error) {
	status, err := submajorproses.repository.Store(request)
	return status, err
}

// Update implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesService) Update(request *models.SubMajorProsesRequest) (err error) {
	status, err := submajorproses.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetDataByID implements SubMajorProsesDefinition
func (submajorproses SubMajorProsesService) GetDataByID(request models.KodeMajor) (responses []models.SubMajorProsesResponse, err error) {
	dataSubMajor, err := submajorproses.repository.GetDataByID(&request)

	if err != nil {
		submajorproses.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataSubMajor {
		responses = append(responses, models.SubMajorProsesResponse{
			ID:                 response.ID,
			IDMajorProses:      response.IDMajorProses,
			KodeSubMajorProses: response.KodeSubMajorProses,
			SubMajorProses:     response.SubMajorProses,
			Deskripsi:          response.Deskripsi,
			Status:             response.Status,
			CreatedAt:          response.CreatedAt,
			UpdatedAt:          response.UpdatedAt,
		})
	}

	return responses, err
}
