package linkcage

import (
	"riskmanagement/lib"
	models "riskmanagement/models/linkcage"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type LinkcageImageDefinition interface {
	Store(request *models.LinkcageImage, tx *gorm.DB) (responses *models.LinkcageImage, err error)

	GetLinkImage(request *models.LinkcageRequest) (responses []models.LinkcageImageResponse, err error)
	DeleteFilesByID(id int64, tx *gorm.DB) (err error)
}

type LinkcageImageRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewLinkcageImageRepository(
	db lib.Database,
	logger logger.Logger,
) LinkcageImageDefinition {
	return LinkcageImageRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (linkcageImage LinkcageImageRepository) Store(request *models.LinkcageImage, tx *gorm.DB) (responses *models.LinkcageImage, err error) {
	return request, tx.Create(&request).Error
}

func (linkcageImage LinkcageImageRepository) GetLinkImage(request *models.LinkcageRequest) (responses []models.LinkcageImageResponse, err error) {
	rows, err := linkcageImage.db.DB.Raw(`
		SELECT 
			li.id 'id_lampiran',
			li.linkcage_id 'linkcage_id',
			fl.filename 'filename',
			fl.path 'path',
			fl.extension 'ext',
			fl.size 'size'
		FROM linkcage_image li 
		JOIN files fl ON fl.id = li.file_id
		WHERE li.linkcage_id = ?`, request.ID).Rows()

	defer rows.Close()
	var img models.LinkcageImageResponse

	for rows.Next() {
		linkcageImage.db.DB.ScanRows(rows, &img)
		responses = append(responses, img)
	}

	return responses, err
}

func (linkcageImage LinkcageImageRepository) DeleteFilesByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("linkcage_id = ?", id).Delete(&models.LinkcageImage{}).Error
}
