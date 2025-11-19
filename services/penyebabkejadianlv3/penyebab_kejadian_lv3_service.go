package penyebabkejadianlv3

import (
	"riskmanagement/lib"
	models "riskmanagement/models/penyebabkejadianlv3"
	repository "riskmanagement/repository/penyebabkejadianlv3"

	"gitlab.com/golang-package-library/logger"
)

type PenyebabKejadianLv3Definition interface {
	// GetAll() (responses []models.PenyebabKejadianLv3Response, err error)
	GetAll() (responses []models.PenyebabKejadianLv3Responses, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.PenyebabKejadianLv3Responses, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.PenyebabKejadianLv3Response, err error)
	Store(request *models.PenyebabKejadianLv3Request) (err error)
	Update(request *models.PenyebabKejadianLv3Request) (err error)
	GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error)
	Delete(id int64) (err error)
	GetKejadianByIDlv2(request models.KodeSubKejadianRequest) (responses []models.PenyebabKejadianLv3Response, err error)
	GetKejadianByIDlv1(request models.KodePenyebabKejadian) (responses []models.PenyebabKejadianLv3Response, err error)
	GetSubKejadian(id int64) (responses []models.PenyebabKejadianLv3Response, err error)
}

type PenyebabKejadianLv3Service struct {
	logger     logger.Logger
	repository repository.PenyebabKejadianLv3Definition
}

func NewPenyebabKejadianLv3Service(logger logger.Logger, repository repository.PenyebabKejadianLv3Definition) PenyebabKejadianLv3Definition {
	return PenyebabKejadianLv3Service{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements PenyebabKejadianLv3Definition
func (penyebabKejadianLv3 PenyebabKejadianLv3Service) Delete(id int64) (err error) {
	return penyebabKejadianLv3.repository.Delete(id)
}

// GetAll implements PenyebabKejadianLv3Definition
func (penyebabKejadianLv3 PenyebabKejadianLv3Service) GetAll() (responses []models.PenyebabKejadianLv3Responses, err error) {
	return penyebabKejadianLv3.repository.GetAll()
}

// GetAllWithPaginate implements PenyebabKejadianLv3Definition
func (pk3 PenyebabKejadianLv3Service) GetAllWithPaginate(request models.Paginate) (responses []models.PenyebabKejadianLv3Responses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPK3, totalRows, totalData, err := pk3.repository.GetAllWithPaginate(&request)

	if err != nil {
		pk3.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		pk3.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPK3 {
		responses = append(responses, models.PenyebabKejadianLv3Responses{
			ID:                      response.ID,
			KodeSubKejadian:         response.KodeSubKejadian,
			PenyebabKejadianLv2:     response.PenyebabKejadianLv2,
			KodePenyebabKejadianLv3: response.KodePenyebabKejadianLv3,
			PenyebabKejadianLv3:     response.PenyebabKejadianLv3,
			Deskripsi:               response.Deskripsi,
			Status:                  response.Status,
			CreatedAt:               response.CreatedAt,
			UpdatedAt:               response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements PenyebabKejadianLv3Definition
func (penyebabKejadianLv3 PenyebabKejadianLv3Service) GetOne(id int64) (responses models.PenyebabKejadianLv3Response, err error) {
	return penyebabKejadianLv3.repository.GetOne(id)
}

// Store implements PenyebabKejadianLv3Definition
func (penyebabKejadianLv3 PenyebabKejadianLv3Service) Store(request *models.PenyebabKejadianLv3Request) (err error) {
	status, err := penyebabKejadianLv3.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements PenyebabKejadianLv3Definition
func (penyebabKejadianLv3 PenyebabKejadianLv3Service) Update(request *models.PenyebabKejadianLv3Request) (err error) {
	status, err := penyebabKejadianLv3.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetKodePenyebabKejadian implements PenyebabKejadianLv3Definition
func (penyebabKejadianLv3 PenyebabKejadianLv3Service) GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error) {
	dataPenyebabKejadianLv3, err := penyebabKejadianLv3.repository.GetKodePenyebabKejadian()
	if err != nil {
		penyebabKejadianLv3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataPenyebabKejadianLv3 {
		responses = append(responses, models.KodePenyebabKejadian{
			KodePenyebabKejadian: response.KodePenyebabKejadian,
		})
	}

	return responses, err
}

// GetKejadianByIDlv2 implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Service) GetKejadianByIDlv2(request models.KodeSubKejadianRequest) (responses []models.PenyebabKejadianLv3Response, err error) {
	dataPK3, err := PK3.repository.GetKejadianByIDlv2(&request)
	if err != nil {
		PK3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataPK3 {
		responses = append(responses, models.PenyebabKejadianLv3Response{
			ID:                      response.ID,
			KodeSubKejadian:         response.KodeSubKejadian,
			KodePenyebabKejadianLv3: response.KodePenyebabKejadianLv3,
			PenyebabKejadianLv3:     response.PenyebabKejadianLv3,
			Deskripsi:               response.Deskripsi,
			Status:                  response.Status,
			CreatedAt:               response.CreatedAt,
			UpdatedAt:               response.UpdatedAt,
		})
	}

	return responses, err
}

// GetKejadianByIDlv1 implements PenyebabKejadianLv3Definition
func (PK3 PenyebabKejadianLv3Service) GetKejadianByIDlv1(request models.KodePenyebabKejadian) (responses []models.PenyebabKejadianLv3Response, err error) {
	dataPK3, err := PK3.repository.GetKejadianByIDlv1(&request)
	if err != nil {
		PK3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataPK3 {
		responses = append(responses, models.PenyebabKejadianLv3Response{
			ID:                      response.ID,
			KodeSubKejadian:         response.KodeSubKejadian,
			KodePenyebabKejadianLv3: response.KodePenyebabKejadianLv3,
			PenyebabKejadianLv3:     response.PenyebabKejadianLv3,
			Deskripsi:               response.Deskripsi,
			Status:                  response.Status,
			CreatedAt:               response.CreatedAt,
			UpdatedAt:               response.UpdatedAt,
		})
	}

	return responses, err
}

// GetSubKejadian implements PenyebabKejadianLv3Definition.
func (pk3 PenyebabKejadianLv3Service) GetSubKejadian(id int64) (responses []models.PenyebabKejadianLv3Response, err error) {
	dataPK3, err := pk3.repository.GetSubKejadian(id)
	if err != nil {
		pk3.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataPK3 {
		responses = append(responses, models.PenyebabKejadianLv3Response{
			ID:                      response.ID,
			KodeSubKejadian:         response.KodeSubKejadian,
			KodePenyebabKejadianLv3: response.KodePenyebabKejadianLv3,
			PenyebabKejadianLv3:     response.PenyebabKejadianLv3,
			Deskripsi:               response.Deskripsi,
			Status:                  response.Status,
			CreatedAt:               response.CreatedAt,
			UpdatedAt:               response.UpdatedAt,
		})
	}

	return responses, err
}
