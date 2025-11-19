package subincident

import (
	"riskmanagement/lib"
	models "riskmanagement/models/subincident"
	repository "riskmanagement/repository/subincident"

	"gitlab.com/golang-package-library/logger"
)

type SubIncidentDefinition interface {
	// GetAll() (responses []models.SubIncidentResponse, err error)
	GetAll() (responses []models.SubIncidentResponses, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.SubIncidentResponses, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.SubIncidentResponse, err error)
	GetSubIncidentByID(requests models.SubIncidentFilterRequest) (responses []models.SubIncidentResponses, err error)
	Store(request *models.SubIncidentRequest) (err error)
	Update(request *models.SubIncidentRequest) (err error)
	GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error)
	Delete(id int64) (err error)
}

type SubIncidentService struct {
	logger     logger.Logger
	repository repository.SubIncidentDefinition
}

func NewSubIncidentService(logger logger.Logger, repository repository.SubIncidentDefinition) SubIncidentDefinition {
	return SubIncidentService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements SubIncidentDefinition
func (subIncident SubIncidentService) Delete(id int64) (err error) {
	return subIncident.repository.Delete(id)
}

// GetAll implements SubIncidentDefinition
// func (subIncident SubIncidentService) GetAll() (responses []models.SubIncidentResponse, err error) {
// 	return subIncident.repository.GetAll()
// }

// GetSubIncidentByID implements SubIncidentDefinition
func (subIncident SubIncidentService) GetSubIncidentByID(request models.SubIncidentFilterRequest) (responses []models.SubIncidentResponses, err error) {
	dataSubIncident, err := subIncident.repository.GetSubIncidentByID(&request)
	if err != nil {
		subIncident.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataSubIncident {
		responses = append(responses, models.SubIncidentResponses{
			ID:                       response.ID.Int64,
			KodeKejadian:             response.KodeKejadian.String,
			PenyebabKejadian:         response.PenyebabKejadian.String,
			KodeSubKejadian:          response.KodeSubKejadian.String,
			KriteriaPenyebabKejadian: response.KriteriaPenyebabKejadian.String,
			CreatedAt:                &response.CreatedAt.String,
			UpdatedAt:                &response.UpdatedAt.String,
		})
	}

	return responses, err
}

func (subIncident SubIncidentService) GetAll() (responses []models.SubIncidentResponses, err error) {
	return subIncident.repository.GetAll()
}

// GetAllWithPaginate implements SubIncidentDefinition
func (pk2 SubIncidentService) GetAllWithPaginate(request models.Paginate) (responses []models.SubIncidentResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPK2, totalRows, totalData, err := pk2.repository.GetAllWithPaginate(&request)

	if err != nil {
		pk2.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		pk2.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPK2 {
		responses = append(responses, models.SubIncidentResponses{
			ID:                       response.ID,
			KodeKejadian:             response.KodeKejadian,
			PenyebabKejadian:         response.PenyebabKejadian,
			KodeSubKejadian:          response.KodeSubKejadian,
			KriteriaPenyebabKejadian: response.KriteriaPenyebabKejadian,
			Deskripsi:                response.Deskripsi,
			Status:                   response.Status,
			CreatedAt:                response.CreatedAt,
			UpdatedAt:                response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetOne implements SubIncidentDefinition
func (subIncident SubIncidentService) GetOne(id int64) (responses models.SubIncidentResponse, err error) {
	return subIncident.repository.GetOne(id)
}

// Store implements SubIncidentDefinition
func (subIncident SubIncidentService) Store(request *models.SubIncidentRequest) (err error) {
	status, err := subIncident.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements SubIncidentDefinition
func (subIncident SubIncidentService) Update(request *models.SubIncidentRequest) (err error) {
	status, err := subIncident.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetKodePenyebabKejadian implements SubIncidentDefinition
func (subIncident SubIncidentService) GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error) {
	dataSubIncident, err := subIncident.repository.GetKodePenyebabKejadian()
	if err != nil {
		subIncident.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataSubIncident {
		responses = append(responses, models.KodePenyebabKejadian{
			KodePenyebabKejadian: response.KodePenyebabKejadian,
		})
	}

	return responses, err
}
