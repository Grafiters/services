package eventtypelv2

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv2"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type EventTypeLv2Definition interface {
	GetAll() (responses []models.EventTypeLv2Response, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.EventTypeLv2Response, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.EventTypeLv2Response, err error)
	Store(request *models.EventTypeLv2Request) (responses bool, err error)
	Update(request *models.EventTypeLv2Request) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeEventType() (responses []models.KodeEventType, err error)
	GetEventTypeById1(request *models.IDEventTypeLv1) (responses []models.EventTypeLv2Response, err error)
	WithTrx(trxHandle *gorm.DB) EventTypeLv2Repository
}

type EventTypeLv2Repository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewEventTypeLv2Repository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) EventTypeLv2Definition {
	return EventTypeLv2Repository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) Delete(id int64) (err error) {
	return et2.db.DB.Where("id = ?", id).Delete(&models.EventTypeLv2Response{}).Error
}

// GetAllWithPaginate implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) GetAllWithPaginate(request *models.Paginate) (responses []models.EventTypeLv2Response, totalRows int, totalData int, err error) {
	rows, err := et2.db.DB.Raw(`
	SELECT
		etl.id 'id',
		etl.id_event_type_lv1 'id_event_type_lv1',
		etl.kode_event_type_lv2 'kode_event_type_lv2',
		etl.event_type_lv2 'event_type_lv2',
		etl.deskripsi 'deskripsi',
		etl.status 'status',
		etl.created_at 'created_at',
		etl.updated_at 'updated_at'
	FROM event_type_lv2 etl  ORDER BY etl.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var subMajorProses models.EventTypeLv2Response

	for rows.Next() {
		et2.db.DB.ScanRows(rows, &subMajorProses)
		responses = append(responses, subMajorProses)
	}

	paginateQuery := `SELECT COUNT(*) FROM event_type_lv2`
	err = et2.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

// GetAll implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) GetAll() (responses []models.EventTypeLv2Response, err error) {
	return responses, et2.db.DB.Where("status = ?", 1).Find(&responses).Error
}

// GetKodeEventType implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) GetKodeEventType() (responses []models.KodeEventType, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_event_type_lv2 + 1)), 4)
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(etl.kode_event_type_lv2,'.', -1) as DECIMAL) kode_event_type_lv2
				FROM event_type_lv2 etl
				ORDER BY etl.id DESC LIMIT 1) 
			AS T`

	et2.logger.Zap.Info(query)
	rows, err := et2.dbRaw.DB.Query(query)
	defer rows.Close()

	et2.logger.Zap.Info("rows ", rows)
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

// GetOne implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) GetOne(id int64) (responses models.EventTypeLv2Response, err error) {
	return responses, et2.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) Store(request *models.EventTypeLv2Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = et2.db.DB.Save(&models.EventTypeLv2Request{
		IDEventTypeLv1:   request.IDEventTypeLv1,
		KodeEventTypeLv2: request.KodeEventTypeLv2,
		EventTypeLv2:     request.EventTypeLv2,
		Deskripsi:        request.Deskripsi,
		Status:           request.Status,
		CreatedAt:        &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) Update(request *models.EventTypeLv2Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = et2.db.DB.Save(&models.EventTypeLv2Request{
		ID:               request.ID,
		IDEventTypeLv1:   request.IDEventTypeLv1,
		KodeEventTypeLv2: request.KodeEventTypeLv2,
		EventTypeLv2:     request.EventTypeLv2,
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

// WithTrx implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) WithTrx(trxHandle *gorm.DB) EventTypeLv2Repository {
	if trxHandle == nil {
		et2.logger.Zap.Error("transaction Database not found in gin context")
		return et2
	}

	et2.db.DB = trxHandle
	return et2
}

// GetEventTypeById1 implements EventTypeLv2Definition
func (et2 EventTypeLv2Repository) GetEventTypeById1(request *models.IDEventTypeLv1) (responses []models.EventTypeLv2Response, err error) {
	if request.IDEventTypeLv1 != "" {
		// where := "WHERE id_event_type_lv1 = '" + request.IDEventTypeLv1 + "'"
		where := "WHERE id_event_type_lv1 = ?"

		query := `SELECT * FROM event_type_lv2 ` + where + ` AND status != 0`

		et2.logger.Zap.Info(query)
		rows, err := et2.dbRaw.DB.Query(query, request.IDEventTypeLv1)
		defer rows.Close()

		et2.logger.Zap.Info("rows =>", rows)
		if err != nil {
			return responses, err
		}

		response := models.EventTypeLv2Response{}
		for rows.Next() {
			_ = rows.Scan(
				&response.ID,
				&response.IDEventTypeLv1,
				&response.KodeEventTypeLv2,
				&response.EventTypeLv2,
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
