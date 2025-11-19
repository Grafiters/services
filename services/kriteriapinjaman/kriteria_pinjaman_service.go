package kriteriapinjaman

import (
	models "riskmanagement/models/kriteriapinjaman"
	repo "riskmanagement/repository/kriteriapinjaman"

	"gitlab.com/golang-package-library/logger"
)

type KriteriaPinjamanDefinition interface {
	Store(request models.KriteriaPinjamanRequest) (response bool, err error)
	Update(request models.KriteriaPinjamanRequest) (response bool, err error)
	Delete(id int64) (err error)
	GetAll(request models.Paginate) (response []models.KriteriaPinjamanResponse, totalRows int, err error)
}

type KriteriaPinjamanService struct {
	logger logger.Logger
	repo   repo.KriteriaPinjamanDefinition
}

func NewKriteriaPinjamanService(
	logger logger.Logger,
	repo repo.KriteriaPinjamanDefinition,
) KriteriaPinjamanDefinition {
	return KriteriaPinjamanService{
		logger: logger,
		repo:   repo,
	}
}

// Delete implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanService) Delete(id int64) (err error) {
	panic("unimplemented")
}

// GetAll implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanService) GetAll(request models.Paginate) (response []models.KriteriaPinjamanResponse, totalRows int, err error) {
	kriteria, totalRows, err := k.repo.GetAll(request)

	if err != nil {
		k.logger.Zap.Error(err)
		return response, totalRows, err
	}

	for _, value := range kriteria {
		response = append(response, models.KriteriaPinjamanResponse{
			ID:       value.ID,
			Kriteria: value.Kriteria,
			Status:   value.Status,
		})
	}

	return response, totalRows, err
}

// Store implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanService) Store(request models.KriteriaPinjamanRequest) (response bool, err error) {
	panic("unimplemented")
}

// Update implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanService) Update(request models.KriteriaPinjamanRequest) (response bool, err error) {
	panic("unimplemented")
}
