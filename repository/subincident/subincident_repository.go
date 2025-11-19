package subincident

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/subincident"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type SubIncidentDefinition interface {
	// GetAll() (responses []models.SubIncidentResponse, err error)
	GetAll() (responses []models.SubIncidentResponses, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.SubIncidentResponses, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.SubIncidentResponse, err error)
	GetSubIncidentByID(requests *models.SubIncidentFilterRequest) (responses []models.SubIncidentListFilter, err error)
	Store(request *models.SubIncidentRequest) (responses bool, err error)
	Update(request *models.SubIncidentRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) SubIncidentRepository
	GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error)
}

type SubIncidentRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewSubIncidentRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) SubIncidentDefinition {
	return SubIncidentRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetSubIncidentByID implements SubIncidentDefinition
func (subIncident SubIncidentRepository) GetSubIncidentByID(requests *models.SubIncidentFilterRequest) (responses []models.SubIncidentListFilter, err error) {
	if requests.KodeKejadian != "" {
		where := " WHERE pk2.kode_kejadian = ?"

		query := `SELECT
				pk2.id 'id',
				pk2.kode_kejadian 'kode_kejadian',
				pk1.penyebab_kejadian 'penyebab_kejadian',
				pk2.kode_sub_kejadian 'kode_sub_kejadian',
				pk2.kriteria_penyebab_kejadian 'kriteria_penyebab_kejadian',
				pk2.created_at 'created_at',
				pk2.updated_at 'updated_at'
			FROM penyebab_kejadian_lv2 pk2
			JOIN penyebab_kejadian_lv1 pk1 ON pk1.kode_kejadian = pk2.kode_kejadian` + where

		subIncident.logger.Zap.Info(query)
		rows, err := subIncident.dbRaw.DB.Query(query, requests.KodeKejadian)
		defer rows.Close()

		subIncident.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.SubIncidentListFilter{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.KodeKejadian,
				&response.PenyebabKejadian,
				&response.KodeSubKejadian,
				&response.KriteriaPenyebabKejadian,
				&response.CreatedAt,
				&response.UpdatedAt,
			)
			responses = append(responses, response)
		}

		if err = rows.Err(); err != nil {
			return responses, err
		}
	}

	return responses, err
}

// Delete implements SubIncidentDefinition
func (subIncident SubIncidentRepository) Delete(id int64) (err error) {
	return subIncident.db.DB.Where("id = ?", id).Delete(&models.SubIncidentResponse{}).Error
}

// GetAll implements SubIncidentDefinition
//
//	func (subIncident SubIncidentRepository) GetAll() (responses []models.SubIncidentResponse, err error) {
//		return responses, subIncident.db.DB.Find(&responses).Error
//	}
func (subIncident SubIncidentRepository) GetAll() (responses []models.SubIncidentResponses, err error) {
	rows, err := subIncident.db.DB.Raw(`
		SELECT
			pk2.id 'id',
			pk2.kode_kejadian 'kode_kejadian',
			pk1.penyebab_kejadian 'penyebab_kejadian',
			pk2.kode_sub_kejadian 'kode_sub_kejadian',
			pk2.kriteria_penyebab_kejadian 'kriteria_penyebab_kejadian',
			pk2.deskripsi 'deskripsi',
			pk2.status 'status',
			pk2.created_at 'created_at',
			pk2.updated_at 'updated_at'
		FROM penyebab_kejadian_lv2 pk2
		LEFT OUTER JOIN penyebab_kejadian_lv1 pk1 ON pk1.kode_kejadian = pk2.kode_kejadian 
		WHERE pk2.status = 1
	`).Rows()

	defer rows.Close()

	var subInci models.SubIncidentResponses

	for rows.Next() {
		subIncident.db.DB.ScanRows(rows, &subInci)
		responses = append(responses, subInci)
	}

	return responses, err
}

// GetAllWithPaginate implements SubIncidentDefinition
func (pk2 SubIncidentRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.SubIncidentResponses, totalRows int, totalData int, err error) {
	rows, err := pk2.db.DB.Raw(`
		SELECT
			pk2.id 'id',
			pk2.kode_kejadian 'kode_kejadian',
			pk1.penyebab_kejadian 'penyebab_kejadian',
			pk2.kode_sub_kejadian 'kode_sub_kejadian',
			pk2.kriteria_penyebab_kejadian 'kriteria_penyebab_kejadian',
			pk2.deskripsi 'deskripsi',
			pk2.status 'status',
			pk2.created_at 'created_at',
			pk2.updated_at 'updated_at'
		FROM penyebab_kejadian_lv2 pk2
		LEFT OUTER JOIN penyebab_kejadian_lv1 pk1 ON pk1.kode_kejadian = pk2.kode_kejadian LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	var subInci models.SubIncidentResponses

	for rows.Next() {
		pk2.db.DB.ScanRows(rows, &subInci)
		responses = append(responses, subInci)
	}

	paginateQuery := `SELECT COUNT(*) FROM penyebab_kejadian_lv2`
	err = pk2.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

// GetOne implements SubIncidentDefinition
func (subIncident SubIncidentRepository) GetOne(id int64) (responses models.SubIncidentResponse, err error) {
	return responses, subIncident.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements SubIncidentDefinition
func (subIncident SubIncidentRepository) Store(request *models.SubIncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = subIncident.db.DB.Save(&models.SubIncidentRequest{
		KodeKejadian:             request.KodeKejadian,
		KodeSubKejadian:          request.KodeSubKejadian,
		KriteriaPenyebabKejadian: request.KriteriaPenyebabKejadian,
		Deskripsi:                request.Deskripsi,
		Status:                   request.Status,
		CreatedAt:                &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements SubIncidentDefinition
func (subIncident SubIncidentRepository) Update(request *models.SubIncidentRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = subIncident.db.DB.Save(&models.SubIncidentRequest{
		ID:                       request.ID,
		KodeKejadian:             request.KodeKejadian,
		KodeSubKejadian:          request.KodeSubKejadian,
		KriteriaPenyebabKejadian: request.KriteriaPenyebabKejadian,
		Deskripsi:                request.Deskripsi,
		Status:                   request.Status,
		CreatedAt:                request.CreatedAt,
		UpdatedAt:                &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements SubIncidentDefinition
func (subIncident SubIncidentRepository) WithTrx(trxHandle *gorm.DB) SubIncidentRepository {
	if trxHandle == nil {
		subIncident.logger.Zap.Error("transaction Database not found in gin context")
		return subIncident
	}

	subIncident.db.DB = trxHandle
	return subIncident
}

func (subIncident SubIncidentRepository) GetKodePenyebabKejadian() (responses []models.KodePenyebabKejadian, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_sub_kejadian + 1)), 4)
				FROM(
					SELECT
						CAST(SUBSTRING_INDEX(pkl.kode_sub_kejadian,'.', -1) as DECIMAL) kode_sub_kejadian 
					FROM penyebab_kejadian_lv2 pkl
					ORDER BY pkl.id DESC LIMIT 1) 
				AS T`

	subIncident.logger.Zap.Info(query)
	rows, err := subIncident.dbRaw.DB.Query(query)
	defer rows.Close()

	subIncident.logger.Zap.Info("rows", rows)
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
