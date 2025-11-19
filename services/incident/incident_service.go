package incident

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/incident"
	repository "riskmanagement/repository/incident"

	"gitlab.com/golang-package-library/logger"
)

type IncidentDefinition interface {
	GetAll() (responses []models.IncidentResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.IncidentResponses, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.IncidentResponse, err error)
	Store(request *models.IncidentRequest) (err error)
	Update(request *models.IncidentRequest) (err error)
	Delete(id int64) (err error)
	GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error)
}

type IncidentService struct {
	logger     logger.Logger
	repository repository.IncidentDefinition
}

func NewIncidentService(
	logger logger.Logger,
	repository repository.IncidentDefinition,
) IncidentDefinition {
	return IncidentService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements IncidentDefinition
func (incident IncidentService) Delete(id int64) (err error) {
	return incident.repository.Delete(id)
}

// GetAll implements IncidentDefinition
func (incident IncidentService) GetAll() (responses []models.IncidentResponse, err error) {
	return incident.repository.GetAll()
}

// GetAllWithPaginate implements IncidentDefinition
func (pk1 IncidentService) GetAllWithPaginate(request models.Paginate) (responses []models.IncidentResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPK1, totalRows, totalData, err := pk1.repository.GetAllWithPaginate(&request)

	fmt.Println("dataPK1 ===>", dataPK1)
	fmt.Println("totalRows =>", totalRows)
	fmt.Println("totalData =>", totalData)
	fmt.Println("Error ==>", err)

	if err != nil {
		pk1.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPK1 {
		responses = append(responses, models.IncidentResponses{
			ID:               response.ID,
			KodeKejadian:     response.KodeKejadian,
			PenyebabKejadian: response.PenyebabKejadian,
			Deskripsi:        response.Deskripsi,
			Status:           response.Status,
			CreatedAt:        response.CreatedAt,
			UpdatedAt:        response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements IncidentDefinition
func (incident IncidentService) GetOne(id int64) (responses models.IncidentResponse, err error) {
	return incident.repository.GetOne(id)
}

// Store implements IncidentDefinition
func (incident IncidentService) Store(request *models.IncidentRequest) (err error) {
	fmt.Println("service = ", request)
	status, err := incident.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements IncidentDefinition
func (incident IncidentService) Update(request *models.IncidentRequest) (err error) {
	status, err := incident.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetKodePenyebabKejadian implements IncidentDefinition
func (incident IncidentService) GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error) {
	dataIncident, err := incident.repository.GetKodePenyebabKejadian()

	if err != nil {
		incident.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataIncident {
		responses = append(responses, models.KodePenyebabKejadian{
			KodePenyebabKejadian: response.KodePenyebabKejadian,
		})
	}

	return responses, err
}
