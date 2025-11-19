package linkcage

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/linkcage"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type LinkcageDefinition interface {
	GetAll(request *models.LinkcageRequest) (responses []models.Linkcage, totalRows int, totalData int, err error)
	Store(request *models.Linkcage, tx *gorm.DB) (responses *models.Linkcage, err error)
	SetStatus(request *models.LinkcageRequest, tx *gorm.DB) (responses bool, err error)
	GetActive() (responses []models.ActiveLinkcage, err error)
	Delete(request *models.LinkcageRequest, tx *gorm.DB) (responses bool, err error)

	GetOne(request *models.LinkcageRequest) (response models.Linkcage, err error)
	StatusCheck(id int64) (status string, err error)
	Update(request *models.Linkcage, tx *gorm.DB) (responses *models.Linkcage, err error)
}

type LinkcageRepository struct {
	db      lib.Database
	timeout time.Duration
}

func NewLinkcageRepository(db lib.Database, dbRaw lib.Databases) LinkcageDefinition {
	return LinkcageRepository{
		db:      db,
		timeout: time.Second * 100,
	}
}

func (link LinkcageRepository) GetAll(request *models.LinkcageRequest) (responses []models.Linkcage, totalRows int, totalData int, err error) {
	// Linkcage Data
	query := link.db.DB.Select(`id, name, status`)
	query.Limit(10).Offset(request.Offset).Find(&responses)

	if query.Error != nil {
		return responses, 0, totalRows, query.Error
	}

	// Pagination
	totalRows64 := int64(totalRows)
	queryCount := link.db.DB.Table(`linkcage`).Count(&totalRows64)
	if queryCount.Error != nil {
		return responses, 0, totalRows, queryCount.Error
	}

	result := float64(int(totalRows64)) / float64(10)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, int(totalRows64), err
}

func (link LinkcageRepository) Store(request *models.Linkcage, tx *gorm.DB) (responses *models.Linkcage, err error) {
	return request, tx.Create(&request).Error
}

func (link LinkcageRepository) GetActive() (responses []models.ActiveLinkcage, err error) {
	rows, err := link.db.DB.Raw(`
		SELECT 
			IFNULL(fl.filename, 'NULL') 'filename',
			fl.path 'path',
			fl.extension 'ext',
			fl.size 'size',
			l.name 'name',
			l.url 'url'
	 		FROM linkcage l
			left join linkcage_image li on li.linkcage_id = l.id
			left join files fl ON fl.id = li.file_id
			WHERE l.status = 'Aktif'`).
		Rows()

	defer rows.Scan()

	var linkCage models.ActiveLinkcage

	for rows.Next() {
		link.db.DB.ScanRows(rows, &linkCage)
		responses = append(responses, linkCage)
	}

	return responses, err
}

func (link LinkcageRepository) Delete(request *models.LinkcageRequest, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Where("id = ?", request.ID).Delete(&models.Linkcage{}).Error
}

func (link LinkcageRepository) GetOne(request *models.LinkcageRequest) (response models.Linkcage, err error) {
	err = link.db.DB.Raw(`
		SELECT l.id, l.name, l.url, l.status FROM linkcage l
		WHERE l.id = ?`, request.ID).Find(&response).Error

	if err != nil {
		return response, err
	}

	return response, err
}

func (link LinkcageRepository) StatusCheck(id int64) (status string, err error) {
	query := link.db.DB.Raw(`SELECT status FROM linkcage l WHERE id = ?`, id).Scan(&status)

	return status, query.Error
}

func (link LinkcageRepository) SetStatus(request *models.LinkcageRequest, tx *gorm.DB) (responses bool, err error) {
	updateTableSQL := `UPDATE linkcage SET status = '` + request.Status + `', updated_at = '` + *request.UpdatedAt + `' WHERE id = ` + strconv.Itoa(int(request.ID)) + `;`

	return true, tx.Exec(updateTableSQL).Error
}

func (link LinkcageRepository) Update(request *models.Linkcage, tx *gorm.DB) (responses *models.Linkcage, err error) {
	// updateTableSQL := `UPDATE linkcage l SET name = '` + request.Name + `', url = '` + request.URL + `', status = '` + request.Status + `', created_at = '` + *request.CreatedAt + `', updated_at = '` + *request.UpdatedAt + `' WHERE id = ` + strconv.Itoa(int(request.ID)) + `;`

	// return responses, link.db.DB.Exec(updateTableSQL).Error
	return request, tx.Save(&request).Error
}
