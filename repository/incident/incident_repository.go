package incident

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/incident"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type IncidentDefinition interface {
	GetAll() (responses []models.IncidentResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.IncidentResponses, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.IncidentResponse, err error)
	Store(request *models.IncidentRequest) (responses bool, err error)
	Update(request *models.IncidentRequest) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error)
	WithTrx(trxHandle *gorm.DB) IncidentRepository
}

type IncidentRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewIncidentRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) IncidentDefinition {
	return IncidentRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements IncidentDefinition
func (incident IncidentRepository) Delete(id int64) (err error) {
	return incident.db.DB.Where("id = ?", id).Delete(&models.IncidentResponse{}).Error
}

// GetAllWithPaginate implements IncidentDefinition
func (pk1 IncidentRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.IncidentResponses, totalRows int, totalData int, err error) {
	rows, err := pk1.db.DB.Raw(`
				SELECT 
					pkl.id 'id',
					pkl.kode_kejadian 'kode_kejadian',
					pkl.penyebab_kejadian 'penyebab_kejadian',
					pkl.deskripsi 'deskripsi',
					pkl.status  'status',
					pkl.created_at 'created_at',
					pkl.updated_at 'updated_at'
				FROM penyebab_kejadian_lv1 pkl ORDER BY pkl.id ASC LIMIT ? OFFSET ?`, request.Limit, request.Offset).Rows()

	defer rows.Scan()
	var kejadian models.IncidentResponses

	for rows.Next() {
		pk1.db.DB.ScanRows(rows, &kejadian)
		responses = append(responses, kejadian)
	}
	paginaiteQuery := `SELECT COUNT(*) FROM penyebab_kejadian_lv1`
	err = pk1.dbRaw.DB.QueryRow(paginaiteQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err

}

// GetAll implements IncidentDefinition
func (incident IncidentRepository) GetAll() (responses []models.IncidentResponse, err error) {
	return responses, incident.db.DB.Where("status = ?", 1).Find(&responses).Error
}

// GetOne implements IncidentDefinition
func (incident IncidentRepository) GetOne(id int64) (responses models.IncidentResponse, err error) {
	return responses, incident.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements IncidentDefinition
func (incident IncidentRepository) Store(request *models.IncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = incident.db.DB.Save(&models.IncidentRequest{
		KodeKejadian:     request.KodeKejadian,
		PenyebabKejadian: request.PenyebabKejadian,
		Deskripsi:        request.Deskripsi,
		Status:           request.Status,
		CreatedAt:        &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements IncidentDefinition
func (incident IncidentRepository) Update(request *models.IncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = incident.db.DB.Save(&models.IncidentRequest{
		ID:               request.ID,
		KodeKejadian:     request.KodeKejadian,
		PenyebabKejadian: request.PenyebabKejadian,
		Deskripsi:        request.Deskripsi,
		Status:           request.Status,
		CreatedAt:        request.CreatedAt,
		UpdatedAt:        &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements IncidentDefinition
func (incident IncidentRepository) WithTrx(trxHandle *gorm.DB) IncidentRepository {
	if trxHandle == nil {
		incident.logger.Zap.Error("transaction Database not found in gin context")
		return incident
	}

	incident.db.DB = trxHandle
	return incident
}

// GetKodePenyebabKejadian implements IncidentDefinition
func (incident IncidentRepository) GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error) {
	// query := `SELECT RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'kode_penyebab_kejadian' FROM penyebab_kejadian_lv1`
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_kejadian + 1)), 4)
				FROM(
					SELECT
						CAST(SUBSTRING_INDEX(pkl.kode_kejadian,'.', -1) as DECIMAL) kode_kejadian 
					FROM penyebab_kejadian_lv1 pkl
					ORDER BY pkl.id DESC LIMIT 1) 
				AS T`

	incident.logger.Zap.Info(query)
	rows, err := incident.dbRaw.DB.Query(query)
	defer rows.Scan()

	incident.logger.Zap.Info("rows", rows)
	for err != nil {
		return responses, err
	}

	response := models.KodePenyebabKejadian{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodePenyebabKejadian,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}
