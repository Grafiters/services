package eventtypelv2

import (
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv2"
	repository "riskmanagement/repository/eventtypelv2"

	"gitlab.com/golang-package-library/logger"
)

type EventTypeLv2Definition interface {
	GetAll() (responses []models.EventTypeLv2Response, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.EventTypeLv2Response, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.EventTypeLv2Response, err error)
	Store(request *models.EventTypeLv2Request) (err error)
	Update(request *models.EventTypeLv2Request) (err error)
	Delete(id int64) (err error)
	GetKodeEventType() (responses []models.KodeEventType, err error)
	GetEventTypeById1(request models.IDEventTypeLv1) (responses []models.EventTypeLv2Response, err error)
}

type EventTypeLv2Service struct {
	logger     logger.Logger
	repository repository.EventTypeLv2Definition
}

func NewEventTypeLv2Service(
	logger logger.Logger,
	repository repository.EventTypeLv2Definition,
) EventTypeLv2Definition {
	return EventTypeLv2Service{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements EventTypeLv2Definition
func (et2 EventTypeLv2Service) Delete(id int64) (err error) {
	return et2.repository.Delete(id)
}

// GetAll implements EventTypeLv2Definition
func (et2 EventTypeLv2Service) GetAll() (responses []models.EventTypeLv2Response, err error) {
	return et2.repository.GetAll()
}

// GetAllWithPaginate implements EventTypeLv1Definition
func (et2 EventTypeLv2Service) GetAllWithPaginate(request models.Paginate) (responses []models.EventTypeLv2Response, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataEvent1, totalRows, totalData, err := et2.repository.GetAllWithPaginate(&request)

	if err != nil {
		et2.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		et2.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataEvent1 {
		responses = append(responses, models.EventTypeLv2Response{
			ID:               response.ID,
			IDEventTypeLv1:   response.IDEventTypeLv1,
			KodeEventTypeLv2: response.KodeEventTypeLv2,
			EventTypeLv2:     response.EventTypeLv2,
			Deskripsi:        response.Deskripsi,
			Status:           response.Status,
			CreatedAt:        response.CreatedAt,
			UpdatedAt:        response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetKodeEventType implements EventTypeLv2Definition
func (et2 EventTypeLv2Service) GetKodeEventType() (responses []models.KodeEventType, err error) {
	dataET2, err := et2.repository.GetKodeEventType()

	if err != nil {
		et2.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataET2 {
		responses = append(responses, models.KodeEventType{
			KodeEventType: response.KodeEventType,
		})
	}

	return responses, err
}

// GetOne implements EventTypeLv2Definition
func (et2 EventTypeLv2Service) GetOne(id int64) (responses models.EventTypeLv2Response, err error) {
	return et2.repository.GetOne(id)
}

// Store implements EventTypeLv2Definition
func (et2 EventTypeLv2Service) Store(request *models.EventTypeLv2Request) (err error) {
	status, err := et2.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements EventTypeLv2Definition
func (et2 EventTypeLv2Service) Update(request *models.EventTypeLv2Request) (err error) {
	status, err := et2.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetEventTypeById1 implements EventTypeLv2Definition
func (et2 EventTypeLv2Service) GetEventTypeById1(request models.IDEventTypeLv1) (responses []models.EventTypeLv2Response, err error) {
	dataEt2, err := et2.repository.GetEventTypeById1(&request)
	if err != nil {
		et2.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataEt2 {
		responses = append(responses, models.EventTypeLv2Response{
			ID:               response.ID,
			IDEventTypeLv1:   response.IDEventTypeLv1,
			KodeEventTypeLv2: response.KodeEventTypeLv2,
			EventTypeLv2:     response.EventTypeLv2,
			Deskripsi:        response.Deskripsi,
			Status:           response.Status,
			CreatedAt:        response.CreatedAt,
			UpdatedAt:        response.UpdatedAt,
		})
	}

	return responses, err
}
