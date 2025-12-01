package eventtypelv1

import (
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv1"
	repository "riskmanagement/repository/eventtypelv1"

	"gitlab.com/golang-package-library/logger"
)

type EventTypeLv1Definition interface {
	GetAll() (responses []models.EventTypeLv1Response, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.EventTypeLv1Response, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.EventTypeLv1Response, err error)
	Store(request *models.EventTypeLv1Request) (err error)
	Update(request *models.EventTypeLv1Request) (err error)
	Delete(id int64) (err error)
	GetKodeEventType() (responses []models.KodeEventType, err error)
}

type EventTypeLv1Service struct {
	logger     logger.Logger
	repository repository.EventTypeLv1Definition
}

func NewEventTypeLv1Service(
	logger logger.Logger,
	repository repository.EventTypeLv1Definition,
) EventTypeLv1Definition {
	return EventTypeLv1Service{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements EventTypeLv1Definition
func (et1 EventTypeLv1Service) Delete(id int64) (err error) {
	return et1.repository.Delete(id)
}

// GetAll implements EventTypeLv1Definition
func (et1 EventTypeLv1Service) GetAll() (responses []models.EventTypeLv1Response, err error) {
	return et1.repository.GetAll()
}

// GetAllWithPaginate implements EventTypeLv1Definition
func (et1 EventTypeLv1Service) GetAllWithPaginate(request models.Paginate) (responses []models.EventTypeLv1Response, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Limit = limit
	request.Page = page
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataEvent1, totalRows, totalData, err := et1.repository.GetAllWithPaginate(&request)

	if err != nil {
		et1.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		et1.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataEvent1 {
		responses = append(responses, models.EventTypeLv1Response{
			ID:            response.ID,
			KodeEventType: response.KodeEventType,
			EventType:     response.EventType,
			Deskripsi:     response.Deskripsi,
			Status:        response.Status,
			CreatedAt:     response.CreatedAt,
			UpdatedAt:     response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetKodeEventType implements EventTypeLv1Definition
func (et1 EventTypeLv1Service) GetKodeEventType() (responses []models.KodeEventType, err error) {
	dataET1, err := et1.repository.GetKodeEventType()

	if err != nil {
		et1.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataET1 {
		responses = append(responses, models.KodeEventType{
			KodeEventType: response.KodeEventType,
		})
	}

	return responses, err
}

// GetOne implements EventTypeLv1Definition
func (et1 EventTypeLv1Service) GetOne(id int64) (responses models.EventTypeLv1Response, err error) {
	return et1.repository.GetOne(id)
}

// Store implements EventTypeLv1Definition
func (et1 EventTypeLv1Service) Store(request *models.EventTypeLv1Request) (err error) {
	status, err := et1.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements EventTypeLv1Definition
func (et1 EventTypeLv1Service) Update(request *models.EventTypeLv1Request) (err error) {
	status, err := et1.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}
