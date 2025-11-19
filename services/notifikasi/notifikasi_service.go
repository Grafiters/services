package notifikasi

import (
	"riskmanagement/lib"
	models "riskmanagement/models/notifikasi"
	repository "riskmanagement/repository/notifikasi"

	"gitlab.com/golang-package-library/logger"
)

type NotifikasiServicesDefinition interface {
	GetNotifikasi(request models.NotifikasiRequest) (responses []models.NotifikasiResponse, pagination lib.Pagination, err error)
	GetTotalNotifikasi(request models.NotifikasiTotalRequest) (responses []models.NotifikasiSimpleResponse, totalRow int, err error)
	UpdateStatus(request models.NotifikasiUpdateStatus) (response bool, err error)
	DeleteStatus(id int) (response bool, err error)
	Store(request models.TasklistNotifikasiRequest) (response models.TasklistNotifikasi, err error)
}

type NotifikasiServices struct {
	logger     logger.Logger
	repository repository.NotifikasiDefinition
}

func (n NotifikasiServices) Store(request models.TasklistNotifikasiRequest) (response models.TasklistNotifikasi, err error) {
	data, err := n.repository.CreateNotification(request)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (n NotifikasiServices) UpdateStatus(request models.NotifikasiUpdateStatus) (response bool, err error) {
	data, err := n.repository.UpdateStatus(request)

	return data, err
}

func (n NotifikasiServices) DeleteStatus(id int) (response bool, err error) {
	data, err := n.repository.DeleteStatus(id)

	return data, err
}

func (n NotifikasiServices) GetNotifikasi(request models.NotifikasiRequest) (responses []models.NotifikasiResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalRows, totalData, err := n.repository.GetNotifikasi(request)

	if err != nil {
		n.logger.Zap.Error(err)
		return responses, pagination, err
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return data, pagination, err
}

func (n NotifikasiServices) GetTotalNotifikasi(request models.NotifikasiTotalRequest) (responses []models.NotifikasiSimpleResponse, totalRow int, err error) {
	responses, totalRow, err = n.repository.GetTotalNotifikasi(request)

	if err != nil {
		n.logger.Zap.Error(err)
		return nil, 0, err
	}

	return responses, totalRow, err
}

func NewNotifikasiServices(
	logger logger.Logger,
	repository repository.NotifikasiDefinition,
) NotifikasiServicesDefinition {
	return NotifikasiServices{
		logger:     logger,
		repository: repository,
	}
}
