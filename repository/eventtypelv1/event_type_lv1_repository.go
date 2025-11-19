package eventtypelv1

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/eventtypelv1"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type EventTypeLv1Definition interface {
	GetAll() (responses []models.EventTypeLv1Response, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.EventTypeLv1Response, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.EventTypeLv1Response, err error)
	Store(request *models.EventTypeLv1Request) (responses bool, err error)
	Update(request *models.EventTypeLv1Request) (responses bool, err error)
	Delete(id int64) (err error)
	GetKodeEventType() (responses []models.KodeEventType, err error)
	WithTrx(trxHandle *gorm.DB) EventTypeLv1Repository
}

type EventTypeLv1Repository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewEventTypeLv1Repository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) EventTypeLv1Definition {
	return EventTypeLv1Repository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// Delete implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) Delete(id int64) (err error) {
	return et1.db.DB.Where("id = ?", id).Delete(&models.EventTypeLv1Response{}).Error
}

// GetAll implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) GetAll() (responses []models.EventTypeLv1Response, err error) {
	return responses, et1.db.DB.Where("status = ?", 1).Find(&responses).Error
}

// GetAllWithPaginate implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) GetAllWithPaginate(request *models.Paginate) (responses []models.EventTypeLv1Response, totalRows int, totalData int, err error) {
	rows, err := et1.db.DB.Raw(`
		SELECT
			etl.id 'id',
			etl.kode_event_type 'kode_event_type',
			etl.event_type 'event_type',
			etl.deskripsi 'deskripsi',
			etl.status 'status',
			etl.created_at 'created_at',
			etl.updated_at 'updated_at'
		FROM event_type_lv1 etl  ORDER BY etl.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()

	var subMajorProses models.EventTypeLv1Response

	for rows.Next() {
		et1.db.DB.ScanRows(rows, &subMajorProses)
		responses = append(responses, subMajorProses)
	}

	paginateQuery := `SELECT COUNT(*) FROM event_type_lv1`
	err = et1.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

// GetKodeEventType implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) GetKodeEventType() (responses []models.KodeEventType, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.kode_event_type + 1)), 4)
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(etl.kode_event_type,'.', -1) as DECIMAL) kode_event_type
				FROM event_type_lv1 etl
				ORDER BY etl.id DESC LIMIT 1) 
			AS T`

	et1.logger.Zap.Info(query)
	rows, err := et1.dbRaw.DB.Query(query)
	defer rows.Close()

	et1.logger.Zap.Info("rows ", rows)
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

// GetOne implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) GetOne(id int64) (responses models.EventTypeLv1Response, err error) {
	return responses, et1.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) Store(request *models.EventTypeLv1Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = et1.db.DB.Save(&models.EventTypeLv1Request{
		KodeEventType: request.KodeEventType,
		EventType:     request.EventType,
		Deskripsi:     request.Deskripsi,
		Status:        request.Status,
		CreatedAt:     &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) Update(request *models.EventTypeLv1Request) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	err = et1.db.DB.Save(&models.EventTypeLv1Request{
		ID:            request.ID,
		KodeEventType: request.KodeEventType,
		EventType:     request.EventType,
		Deskripsi:     request.Deskripsi,
		Status:        request.Status,
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements EventTypeLv1Definition
func (et1 EventTypeLv1Repository) WithTrx(trxHandle *gorm.DB) EventTypeLv1Repository {
	if trxHandle == nil {
		et1.logger.Zap.Error("transaction Database not found in gin context")
		return et1
	}

	et1.db.DB = trxHandle
	return et1
}
