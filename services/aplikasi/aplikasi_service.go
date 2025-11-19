package aplikasi

import (
	models "riskmanagement/models/aplikasi"
	repository "riskmanagement/repository/aplikasi"

	"gitlab.com/golang-package-library/logger"
)

type AplikasiDefinition interface {
	GetAll() (responses []models.AplikasiResponse, err error)
	GetOne(id int64) (responses models.AplikasiResponse, err error)
	Store(request *models.AplikasiRequest) (err error)
	Update(requests *models.AplikasiRequest) (err error)
	Delete(id int64) (err error)
}

type AplikasiService struct {
	logger     logger.Logger
	repository repository.AplikasiDefinition
}

func NewAplikasiService(
	logger logger.Logger,
	repository repository.AplikasiDefinition,
) AplikasiDefinition {
	return AplikasiService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements AplikasiDefinition
func (aplikasi AplikasiService) Delete(id int64) (err error) {
	return aplikasi.repository.Delete(id)
}

// GetAll implements AplikasiDefinition
func (aplikasi AplikasiService) GetAll() (responses []models.AplikasiResponse, err error) {
	return aplikasi.repository.GetAll()
}

// GetOne implements AplikasiDefinition
func (aplikasi AplikasiService) GetOne(id int64) (responses models.AplikasiResponse, err error) {
	return aplikasi.repository.GetOne(id)
}

// Store implements AplikasiDefinition
func (aplikasi AplikasiService) Store(request *models.AplikasiRequest) (err error) {
	status, err := aplikasi.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements AplikasiDefinition
func (aplikasi AplikasiService) Update(requests *models.AplikasiRequest) (err error) {
	status, err := aplikasi.repository.Update(requests)
	if !status || err != nil {
		return err
	}

	return nil
}
