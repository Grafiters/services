package riskcontrol

import (
	"riskmanagement/lib"
	models "riskmanagement/models/riskcontrol"
	repository "riskmanagement/repository/riskcontrol"

	"gitlab.com/golang-package-library/logger"
)

type RiskControlDefinition interface {
	GetAll() (responses []models.RiskControlResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.RiskControlResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.RiskControlResponse, err error)
	Store(request *models.RiskControlRequest) (err error)
	Update(request *models.RiskControlRequest) (err error)
	Delete(id int64) (err error)
	GetKodeRiskControl() (responses []models.KodeRiskControl, err error)
	SearchRiskControlByIssue(request models.KeywordRequest) (response []models.RiskControlResponses, pagination lib.Pagination, err error)
}

type RiskControlService struct {
	logger     logger.Logger
	repository repository.RiskControlDefinition
}

func NewRiskControService(
	logger logger.Logger,
	repository repository.RiskControlDefinition,
) RiskControlDefinition {
	return RiskControlService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements RiskControlDefinition
func (riskControl RiskControlService) Delete(id int64) (err error) {
	return riskControl.repository.Delete(id)
}

// GetAllWithPaginate implements RiskControlDefinition
func (rc RiskControlService) GetAllWithPaginate(request models.Paginate) (responses []models.RiskControlResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPgs, totalRows, totalData, err := rc.repository.GetAllWithPaginate(&request)
	if err != nil {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPgs {
		responses = append(responses, models.RiskControlResponse{
			ID:          response.ID,
			Kode:        response.Kode,
			RiskControl: response.RiskControl,
			Deskripsi:   response.Deskripsi,
			Status:      response.Status,
			CreatedAt:   response.CreatedAt,
			UpdatedAt:   response.UpdatedAt,
		})
	}

	// pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	pagination = lib.SetPaginationResponse(page, limit, int(totalRows), int(totalData))
	return responses, pagination, err
}

// GetAll implements RiskControlDefinition
func (riskControl RiskControlService) GetAll() (responses []models.RiskControlResponse, err error) {
	return riskControl.repository.GetAll()
}

// GetOne implements RiskControlDefinition
func (riskControl RiskControlService) GetOne(id int64) (responses models.RiskControlResponse, err error) {
	// return riskControl.repository.GetOne(id)
	data, err := riskControl.repository.GetOne(id)

	return data, err
}

// Store implements RiskControlDefinition
func (riskControl RiskControlService) Store(request *models.RiskControlRequest) (err error) {
	status, err := riskControl.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements RiskControlDefinition
func (riskControl RiskControlService) Update(request *models.RiskControlRequest) (err error) {
	status, err := riskControl.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetKodeRiskControl implements RiskControlDefinition
func (riskControl RiskControlService) GetKodeRiskControl() (responses []models.KodeRiskControl, err error) {
	dataRC, err := riskControl.repository.GetKodeRiskControl()

	if err != nil {
		riskControl.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataRC {
		responses = append(responses, models.KodeRiskControl{
			KodeRiskControl: response.KodeRiskControl,
		})
	}

	return responses, err
}

// SearchRiskControlByIssue implements RiskControlDefinition
func (rc RiskControlService) SearchRiskControlByIssue(request models.KeywordRequest) (responses []models.RiskControlResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataControl, totalRows, totalData, err := rc.repository.SearchRiskControlByIssue(&request)
	if err != nil {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataControl {
		responses = append(responses, models.RiskControlResponses{
			ID:          response.ID,
			Kode:        response.Kode,
			RiskControl: response.RiskControl,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}
