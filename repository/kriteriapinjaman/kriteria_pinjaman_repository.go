package kriteriapinjaman

import (
	"riskmanagement/lib"
	models "riskmanagement/models/kriteriapinjaman"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type KriteriaPinjamanDefinition interface {
	WithTrx(trxHandle *gorm.DB) KriteriaPinjamanRepository
	Store(request *models.KriteriaPinjamanRequest) (response bool, err error)
	Update(request *models.KriteriaPinjamanRequest) (response bool, err error)
	Delete(id int64) (err error)
	GetAll(request models.Paginate) (response []models.KriteriaPinjamanResponse, totalRows int, err error)
}
type KriteriaPinjamanRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewKriteriaPinjamanRepository(
	db lib.Database,
	logger logger.Logger,
) KriteriaPinjamanDefinition {
	return KriteriaPinjamanRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanRepository) WithTrx(trxHandle *gorm.DB) KriteriaPinjamanRepository {
	panic("unimplemented")
}

// Delete implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanRepository) Delete(id int64) (err error) {
	panic("unimplemented")
}

// GetAll implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanRepository) GetAll(request models.Paginate) (response []models.KriteriaPinjamanResponse, totalRows int, err error) {
	db := k.db.DB.Table("tbl_mst_kriteria_pinjaman")

	orderedBy := ""
	if request.Order == "DESC" {
		orderedBy = "id DESC"
	} else if request.Order == "ASC" {
		orderedBy = "id ASC"
	}

	db = db.Select("id, kriteria, status").Order(orderedBy)

	var count int64
	db.Count(&count)

	totalRows = int(count)

	if request.Limit != 0 && request.Offset != 0 {
		db = db.Limit(request.Limit).Offset(request.Offset)
	} else {
		db = db.Where("status != 0")
	}

	err = db.Scan(&response).Error

	return response, totalRows, err
}

// Store implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanRepository) Store(request *models.KriteriaPinjamanRequest) (response bool, err error) {
	panic("unimplemented")
}

// Update implements KriteriaPinjamanDefinition.
func (k KriteriaPinjamanRepository) Update(request *models.KriteriaPinjamanRequest) (response bool, err error) {
	panic("unimplemented")
}
