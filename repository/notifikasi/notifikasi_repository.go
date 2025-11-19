package notifikasi

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/notifikasi"
	"time"

	"gitlab.com/golang-package-library/logger"
)

var (
	timeNow = lib.GetTimeNow("timestime")
)

type NotifikasiDefinition interface {
	GetNotifikasi(request models.NotifikasiRequest) (responses []models.NotifikasiResponse, totalRows int, totalData int, err error)
	GetTotalNotifikasi(request models.NotifikasiTotalRequest) (responses []models.NotifikasiSimpleResponse, Row int, err error)
	UpdateStatus(request models.NotifikasiUpdateStatus) (responses bool, err error)
	DeleteStatus(id int) (responses bool, err error)
	CreateNotification(request models.TasklistNotifikasiRequest) (responses models.TasklistNotifikasi, err error)
}

type NotifikasiRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func (n NotifikasiRepository) UpdateStatus(request models.NotifikasiUpdateStatus) (responses bool, err error) {
	err = n.db.DB.Exec("UPDATE tasklist_notifikasis SET status = ? WHERE id = ?", request.Status, request.ID).Error
	if err != nil {
		n.logger.Zap.Error(err)
		return false, err
	}

	return true, nil
}

func (n NotifikasiRepository) DeleteStatus(id int) (responses bool, err error) {
	err = n.db.DB.Exec("DELETE FROM tasklist_notifikasis WHERE id = ?", id).Error
	if err != nil {
		n.logger.Zap.Error(err)
		return false, err
	}

	return true, nil
}

func (n NotifikasiRepository) GetTotalNotifikasi(request models.NotifikasiTotalRequest) (responses []models.NotifikasiSimpleResponse, totalRow int, err error) {
	db := n.db.DB.Table(`tasklist_notifikasis`).
		Select(`
			DISTINCT jenis,
			count(jenis) as total
		`).
		// Joins(`LEFT JOIN tasklists on tasklists.id = tasklist_notifikasis.task_id`).
		// Joins(`LEFT JOIN tasklists_uker on tasklists_uker.tasklist_id = tasklists.id`).
		Where(`tasklist_notifikasis.status = 0`).
		Where(`tasklist_notifikasis.receiver = ?`, request.PERNR)
		// Where(`(tasklists.validation = ? OR tasklists.approval = ?)`, request.PERNR, request.PERNR)

	// var count int64
	err = db.Find(&responses).Error

	queryCount := `select count(*) as total from tasklist_notifikasis 
					WHERE status = 0 AND receiver = '` + request.PERNR + `'`

	err = n.db.DB.Raw(queryCount).Scan(&totalRow).Error
	if err != nil {
		n.logger.Zap.Error(err)
		return responses, 0, err
	}
	n.logger.Zap.Info(responses)

	return responses, totalRow, err
}

func (n NotifikasiRepository) GetNotifikasi(request models.NotifikasiRequest) (responses []models.NotifikasiResponse, totalRows int, totalData int, err error) {
	baseQuery := ``

	// if request.Branch != "" {
	// 	baseQuery = `Select tasklist_notifikasis.id, DATE_FORMAT(tanggal,"%d-%b-%Y") as tanggal, task_id, keterangan, tasklist_notifikasis.status, jenis from tasklist_notifikasis JOIN tasklists_uker ON tasklists_uker.tasklist_id = tasklist_notifikasis.task_id where tasklist_notifikasis.status = 0 AND tasklists_uker.branch =  ` + request.Branch + ` `
	// } else {
	baseQuery = `Select tasklist_notifikasis.id, DATE_FORMAT(tanggal,"%d-%b-%Y") as tanggal, task_id, keterangan, tasklist_notifikasis.status, jenis, uker from tasklist_notifikasis where status = 0 AND receiver = '` + request.PERNR + `'`
	// }

	// query := baseQuery + `order by tanggal desc LIMIT ? OFFSET ?`
	query := baseQuery + `order by id desc LIMIT ? OFFSET ?`
	queryCount := `select count(*) from (` + baseQuery + `) as totalData`

	err = n.db.DB.Raw(query,
		request.Limit,
		request.Offset).Scan(&responses).Error

	err = n.db.DB.Raw(queryCount).Scan(&totalRows).Error

	if err != nil {
		n.logger.Zap.Error(err)
		return responses, 0, 0, err
	}

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (n NotifikasiRepository) CreateNotification(request models.TasklistNotifikasiRequest) (responses models.TasklistNotifikasi, err error) {
	//TODO implement me
	timeNow := lib.GetTimeNow("timestime")
	createNotifikasi := models.TasklistNotifikasi{
		Keterangan: request.Keterangan,
		Jenis:      request.Jenis,
		Tanggal:    timeNow,
		CreatedBy:  request.Pernr,
		DateShow:   request.DateShow,
	}
	return createNotifikasi, n.db.DB.Save(&createNotifikasi).Error
}

func NewNotifiksaiRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) NotifikasiDefinition {
	return NotifikasiRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}
