package eventtypelv3

import (
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv3"
	repository "riskmanagement/repository/eventtypelv3"

	"gitlab.com/golang-package-library/logger"
)

type EventTypeLv3Definition interface {
	GetAll() (responses []models.EventTypeLv3Response, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.EventTypeLv3Response, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.EventTypeLv3Response, err error)
	Store(request *models.EventTypeLv3Request) (err error)
	Update(request *models.EventTypeLv3Request) (err error)
	Delete(id int64) (err error)
	GetKodeEventType() (responses []models.KodeEventType, err error)
	GetEventTypeById2(requests models.IDEventTypeLv2) (responses []models.EventTypeLv3Response, err error)
}

type EventTypeLv3Service struct {
	logger     logger.Logger
	repository repository.EventTypeLv3Definition
}

func NewEventTypeLv3Service(
	logger logger.Logger,
	repository repository.EventTypeLv3Definition,
) EventTypeLv3Definition {
	return EventTypeLv3Service{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements EventTypeLv3Definition
func (et3 EventTypeLv3Service) Delete(id int64) (err error) {
	return et3.repository.Delete(id)
}

// GetAllWithPaginate implements EventTypeLv1Definition
func (et3 EventTypeLv3Service) GetAllWithPaginate(request models.Paginate) (responses []models.EventTypeLv3Response, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataEvent1, totalRows, totalData, err := et3.repository.GetAllWithPaginate(&request)

	if err != nil {
		et3.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		et3.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataEvent1 {
		responses = append(responses, models.EventTypeLv3Response{
			ID:               response.ID,
			IDEventTypeLv2:   response.IDEventTypeLv2,
			KodeEventTypeLv3: response.KodeEventTypeLv3,
			EventTypeLv3:     response.EventTypeLv3,
			Deskripsi:        response.Deskripsi,
			Status:           response.Status,
			CreatedAt:        response.CreatedAt,
			UpdatedAt:        response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetAll implements EventTypeLv3Definition
func (et3 EventTypeLv3Service) GetAll() (responses []models.EventTypeLv3Response, err error) {
	return et3.repository.GetAll()
}

// GetKodeEventType implements EventTypeLv3Definition
func (et3 EventTypeLv3Service) GetKodeEventType() (responses []models.KodeEventType, err error) {
	dataET3, err := et3.repository.GetKodeEventType()

	if err != nil {
		et3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataET3 {
		responses = append(responses, models.KodeEventType{
			KodeEventType: response.KodeEventType,
		})
	}

	return responses, err
}

// GetOne implements EventTypeLv3Definition
func (et3 EventTypeLv3Service) GetOne(id int64) (responses models.EventTypeLv3Response, err error) {
	return et3.repository.GetOne(id)
}

// Store implements EventTypeLv3Definition
func (et3 EventTypeLv3Service) Store(request *models.EventTypeLv3Request) (err error) {
	status, err := et3.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements EventTypeLv3Definition
func (et3 EventTypeLv3Service) Update(request *models.EventTypeLv3Request) (err error) {
	status, err := et3.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetEventTypeById2 implements EventTypeLv3Definition
func (et3 EventTypeLv3Service) GetEventTypeById2(requests models.IDEventTypeLv2) (responses []models.EventTypeLv3Response, err error) {
	dataEt3, err := et3.repository.GetEventTypeById2(&requests)
	if err != nil {
		et3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataEt3 {
		responses = append(responses, models.EventTypeLv3Response{
			ID:               response.ID,
			IDEventTypeLv2:   response.IDEventTypeLv2,
			KodeEventTypeLv3: response.KodeEventTypeLv3,
			EventTypeLv3:     response.EventTypeLv3,
			Deskripsi:        response.Deskripsi,
			Status:           response.Status,
			CreatedAt:        response.CreatedAt,
			UpdatedAt:        response.UpdatedAt,
		})
	}

	return responses, err
}
