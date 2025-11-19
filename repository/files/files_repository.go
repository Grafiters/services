package files

import (
	"riskmanagement/lib"
	models "riskmanagement/models/files"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type FilesDefinition interface {
	GetAll() (responses []models.FilesResponse, err error)
	GetOne(id int64) (responses models.FilesResponse, err error)
	GetOneFiles(id int64) (responses []models.FilesResponses, err error)
	Store(request *models.Files, tx *gorm.DB) (responses *models.Files, err error)
	Update(request *models.FilesRequest) (responses bool, err error)
	Delete(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) FilesRepository
}

type FilesRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewFilesRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) FilesDefinition {
	return FilesRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements FilesDefinition
func (files FilesRepository) Delete(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.FilesResponse{}).Error
}

// GetAll implements FilesDefinition
func (files FilesRepository) GetAll() (responses []models.FilesResponse, err error) {
	return responses, files.db.DB.Find(&responses).Error
}

// GetOne implements FilesDefinition
func (files FilesRepository) GetOne(id int64) (responses models.FilesResponse, err error) {
	return responses, files.db.DB.Raw("SELECT mf.id id, f.filename filename, f.`path` path, f.extension extension, f.`size` size FROM materi_files mf JOIN files f on mf.files_id = f.id WHERE mf.id  = ? ", id).Find(&responses).Error
}

// GetOneFiles implements FilesDefinition
func (files FilesRepository) GetOneFiles(id int64) (responses []models.FilesResponses, err error) {
	rows, err := files.db.DB.Raw("SELECT mf.id id, f.filename filename, f.`path` path, f.extension extension, f.`size` size FROM materi_files mf JOIN files f on mf.files_id = f.id WHERE mf.id  = ? ", id).Find(&responses).Rows()

	defer rows.Close()
	var allFiles models.FilesResponses
	for rows.Next() {
		files.db.DB.ScanRows(rows, &allFiles)
		responses = append(responses, allFiles)
	}

	return responses, err
}

// Store implements FilesDefinition
func (files FilesRepository) Store(request *models.Files, tx *gorm.DB) (responses *models.Files, err error) {
	return request, tx.Save(&request).Error
}

// Update implements FilesDefinition
func (files FilesRepository) Update(request *models.FilesRequest) (responses bool, err error) {
	return true, files.db.DB.Save(&request).Error
}

// WithTrx implements FilesDefinition
func (files FilesRepository) WithTrx(trxHandle *gorm.DB) FilesRepository {
	if trxHandle == nil {
		files.logger.Zap.Error("transaction Database not found in gin context")
		return files
	}
	files.db.DB = trxHandle
	return files
}
