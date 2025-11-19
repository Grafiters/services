package eventtypelv3

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv3"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type EventTypeLv3Definition interface {
	GetAll() (responses []models.EventTypeLv3Response, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.EventTypeLv3Response, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.EventTypeLv3Response, err error)
	Store(request *models.EventTypeLv3Request) (responses bool, err error)
	Update(request *models.EventTypeLv3Request) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeEventType() (responses []models.KodeEventType, err error)
	GetEventTypeById2(request *models.IDEventTypeLv2) (responses []models.EventTypeLv3Response, err error)
	WithTrx(trxHandle *gorm.DB) EventTypeLv3Repository
}

type EventTypeLv3Repository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewEventTypeLv3Repository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) EventTypeLv3Definition {
	return EventTypeLv3Repository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) Delete(id int64) (err error) {
	return et3.db.DB.Where("id = ?", id).Delete(&models.EventTypeLv3Response{}).Error
}

// GetAllWithPaginate implements EventTypeLv2Definition
func (et3 EventTypeLv3Repository) GetAllWithPaginate(request *models.Paginate) (responses []models.EventTypeLv3Response, totalRows int, totalData int, err error) {
	rows, err := et3.db.DB.Raw(`
	SELECT
		etl.id 'id',
		etl.id_event_type_lv2 'id_event_type_lv2',
		etl.kode_event_type_lv3 'kode_event_type_lv3',
		etl.event_type_lv3 'event_type_lv3',
		etl.deskripsi 'deskripsi',
		etl.status 'status',
		etl.created_at 'created_at',
		etl.updated_at 'updated_at'
	FROM event_type_lv3 etl  ORDER BY etl.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var subMajorProses models.EventTypeLv3Response

	for rows.Next() {
		et3.db.DB.ScanRows(rows, &subMajorProses)
		responses = append(responses, subMajorProses)
	}

	paginateQuery := `SELECT COUNT(*) FROM event_type_lv3`
	err = et3.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetAll implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) GetAll() (responses []models.EventTypeLv3Response, err error) {
	return responses, et3.db.DB.Find(&responses).Error
}

// GetKodeEventType implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) GetKodeEventType() (responses []models.KodeEventType, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_event_type_lv3 + 1)), 4)
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(etl.kode_event_type_lv3,'.', -1) as DECIMAL) kode_event_type_lv3
				FROM event_type_lv3 etl
				ORDER BY etl.id DESC LIMIT 1) 
			AS T`

	et3.logger.Zap.Info(query)
	rows, err := et3.dbRaw.DB.Query(query)

	defer rows.Close()

	et3.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, err
	}

	response := models.KodeEventType{}
	for rows.Next() {
		_ = rows.Scan(
			&response.KodeEventType,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetOne implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) GetOne(id int64) (responses models.EventTypeLv3Response, err error) {
	return responses, et3.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) Store(request *models.EventTypeLv3Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = et3.db.DB.Save(&models.EventTypeLv3Request{
		IDEventTypeLv2:   request.IDEventTypeLv2,
		KodeEventTypeLv3: request.KodeEventTypeLv3,
		EventTypeLv3:     request.EventTypeLv3,
		Deskripsi:        request.Deskripsi,
		Status:           request.Status,
		CreatedAt:        &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) Update(request *models.EventTypeLv3Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = et3.db.DB.Save(&models.EventTypeLv3Request{
		ID:               request.ID,
		IDEventTypeLv2:   request.IDEventTypeLv2,
		KodeEventTypeLv3: request.KodeEventTypeLv3,
		EventTypeLv3:     request.EventTypeLv3,
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

// WithTrx implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) WithTrx(trxHandle *gorm.DB) EventTypeLv3Repository {
	if trxHandle == nil {
		et3.logger.Zap.Error("transaction Database not found in gin context")
		return et3
	}

	et3.db.DB = trxHandle
	return et3
}

// GetEventTypeById2 implements EventTypeLv3Definition
func (et3 EventTypeLv3Repository) GetEventTypeById2(request *models.IDEventTypeLv2) (responses []models.EventTypeLv3Response, err error) {
	if request.IDEventTypeLv2 != "" {
		// where := "WHERE id_event_type_lv2 = '" + request.IDEventTypeLv2 + "'"
		where := "WHERE id_event_type_lv2 = ?"

		query := `SELECT * FROM event_type_lv3 ` + where + ` AND status != 0`
		et3.logger.Zap.Info(query)
		rows, err := et3.dbRaw.DB.Query(query, request.IDEventTypeLv2)
		defer rows.Close()

		et3.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.EventTypeLv3Response{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.IDEventTypeLv2,
				&response.KodeEventTypeLv3,
				&response.EventTypeLv3,
				&response.Deskripsi,
				&response.Status,
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
