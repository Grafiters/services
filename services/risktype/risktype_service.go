package risktype

import (
	"riskmanagement/lib"
	models "riskmanagement/models/risktype"
	repository "riskmanagement/repository/risktype"

	"gitlab.com/golang-package-library/logger"
)

type RiskTypeDefinition interface {
	GetAll() (responses []models.RiskTypeResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.RiskTypeResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.RiskTypeResponse, err error)
	Store(request *models.RiskTypeRequest) (err error)
	Update(request *models.RiskTypeRequest) (err error)
	Delete(id int64) (err error)
}

type RiskTypeService struct {
	logger     logger.Logger
	repository repository.RiskTypeDefinition
}

// GetAllWithPaginate implements RiskTypeDefinition
func (rt RiskTypeService) GetAllWithPaginate(request models.Paginate) (responses []models.RiskTypeResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	datart, totalRows, totalData, err := rt.repository.GetAllWithPaginate(&request)

	if err != nil {
		rt.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		rt.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range datart {
		responses = append(responses, models.RiskTypeResponse{
			ID:           response.ID,
			RiskTypeCode: response.RiskTypeCode,
			RiskType:     response.RiskType,
			Deskripsi:    response.Deskripsi,
			Status:       response.Status,
			CreatedAt:    response.CreatedAt,
			UpdatedAt:    response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// Delete implements RiskIssueDefinition
func (riskType RiskTypeService) Delete(id int64) (err error) {
	return riskType.repository.Delete(id)
}

// GetAll implements RiskIssueDefinition
func (riskType RiskTypeService) GetAll() (responses []models.RiskTypeResponse, err error) {
	return riskType.repository.GetAll()
}

// GetOne implements RiskIssueDefinition
func (riskType RiskTypeService) GetOne(id int64) (responses models.RiskTypeResponse, err error) {
	return riskType.repository.GetOne(id)
}

// Store implements RiskIssueDefinition
func (riskType RiskTypeService) Store(request *models.RiskTypeRequest) (err error) {
	// fmt.Println("service = ", request)
	status, err := riskType.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements RiskIssueDefinition
func (riskType RiskTypeService) Update(request *models.RiskTypeRequest) (err error) {
	status, err := riskType.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

func NewRiskTypeService(
	logger logger.Logger,
	repository repository.RiskTypeDefinition,
) RiskTypeDefinition {
	return RiskTypeService{
		logger:     logger,
		repository: repository,
	}
}
