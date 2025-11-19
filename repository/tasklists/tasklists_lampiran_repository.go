package tasklists

import (
	"riskmanagement/lib"
	models "riskmanagement/models/tasklists"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TasklistsLampiranDefinition interface {
	GetOneFileByID(id int64) (responses []models.TasklistFilesResponses, err error)
	Store(request *models.TasklistsFiles, tx *gorm.DB) (responses *models.TasklistsFiles, err error)
	DeleteFilesByID(id int64, tx *gorm.DB) (err error)
}

type TasklistsLampiranRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewTasklistsLampiranRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) TasklistsLampiranDefinition {
	return TasklistsLampiranRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (lampiranFiles TasklistsLampiranRepository) GetOneFileByID(id int64) (responses []models.TasklistFilesResponses, err error) {
	rows, err := lampiranFiles.db.DB.Raw(`
		SELECT 
			lam.id 'id_lampiran',
			lam.tasklists_id 'tasklists_id',
			fl.filename 'filename',
			fl.path 'path',
			fl.extension 'ext',
			fl.size 'size'
		FROM tasklists_lampiran lam 
		JOIN files fl ON fl.id = lam.files_id
		WHERE lam.tasklists_id = ?`, id).Rows()

	defer rows.Close()
	var lampiran models.TasklistFilesResponses

	for rows.Next() {
		lampiranFiles.db.DB.ScanRows(rows, &lampiran)
		responses = append(responses, lampiran)
	}

	return responses, err
}

// Store implements VerifikasiFilesDefinition
func (lampiranFiles TasklistsLampiranRepository) Store(request *models.TasklistsFiles, tx *gorm.DB) (responses *models.TasklistsFiles, err error) {
	// return request, tx.Save(&request).Error
	return request, tx.Create(&request).Error
}

// DeleteFilesByID implements VerifikasiFilesDefinition
func (lampiranFiles TasklistsLampiranRepository) DeleteFilesByID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("tasklists_id = ?", id).Delete(&models.TasklistsFiles{}).Error
}
