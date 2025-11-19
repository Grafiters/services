package monitoring

import (
	"riskmanagement/lib"
	models "riskmanagement/models/monitoring"
	repository "riskmanagement/repository/monitoring"

	"gitlab.com/golang-package-library/logger"
)

type MonitoringServicesDefinition interface {
	GetMonitoringTasklistUker(request models.MonitoringTasklistRequest) (response []models.MonitoringTasklistUnitKerjaResponse, pagination lib.Pagination, err error)
	GetMonitoringPekerja(request models.MonitoringTasklistRequest) (responses []models.MonitoringTasklistPekerjaResponse, pagination lib.Pagination, err error)
}

type MonitoringServices struct {
	logger     logger.Logger
	repository repository.MonitoringDefinition
}

func NewMontioringServices(
	logger logger.Logger,
	repository repository.MonitoringDefinition,
) MonitoringServicesDefinition {
	return MonitoringServices{
		logger:     logger,
		repository: repository,
	}
}

func (m MonitoringServices) GetMonitoringTasklistUker(request models.MonitoringTasklistRequest) (responses []models.MonitoringTasklistUnitKerjaResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalRow, totalData, err := m.repository.GetUnitKerjaTasklist(request)
	if err != nil {
		m.logger.Zap.Error(err)
		return responses, pagination, err
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRow, totalData)
	return data, pagination, nil
}

func (m MonitoringServices) GetMonitoringPekerja(request models.MonitoringTasklistRequest) (responses []models.MonitoringTasklistPekerjaResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalRow, totalData, err := m.repository.GetPekerjaTasklist(request)
	if err != nil {
		m.logger.Zap.Error(err)
		return responses, pagination, err
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRow, totalData)
	return data, pagination, nil
}
