package materi

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/materi"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MateriDefinition interface {
	// GetAll() (responses []models.Materi, err error)
	GetAll() (responses []models.MateriAllResponse, err error)
	GetAllMateriFiles(materiID int64) (responses []models.MateriFilesResponse, err error)
	Store(request *models.Materi, tx *gorm.DB) (responses *models.Materi, err error)
	StoreMateriFiles(request *models.MateriRequest, tx *gorm.DB) (responses bool, err error)
	Delete(request models.MateriFilesRequest) (status bool, err error)
	WithTrx(trxHandle *gorm.DB) MateriRepository
	GetMateriByActivityAndProduct(request *models.GetMateriByActivityAndProductRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error)
	GetVerifikasiMateri(request *models.GetMateriVerifikasiRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error)
}

type MateriRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMateriRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MateriDefinition {
	return MateriRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements MaterDefinition
func (materi MateriRepository) Delete(request models.MateriFilesRequest) (status bool, err error) {
	err = materi.db.DB.Where("files_id = ?", request.FilesID).Delete(&models.MateriFiles{}).Error
	if err != nil {
		materi.logger.Zap.Error(err)
		return false, err
	}
	return true, err
}

// GetAll implements MaterDefinition
// func (materi MateriRepository) GetAll() (responses []models.Materi, err error) {
// 	return responses, materi.db.DB.Find(&responses).Error
// }

func (materi MateriRepository) GetAll() (responses []models.MateriAllResponse, err error) {
	return responses, materi.db.DB.Find(&responses).Error
}

// GetAllMateriFiles implements MaterDefinition
func (materi MateriRepository) GetAllMateriFiles(materiID int64) (responses []models.MateriFilesResponse, err error) {
	rows, err := materi.db.DB.Raw(`
				SELECT 
					mf.id materi_files_id,
					mf.materi_id, 
					mf.files_id, f.filename, 
					f.path ,f.extension,f.size  
				FROM materi_files mf
				LEFT JOIN files f on mf.files_id = f.id 
				LEFT JOIN materi m on mf.materi_id = m.id
				WHERE m.id = ? `, materiID).Rows()

	defer rows.Close()
	var materiFiles models.MateriFilesResponse
	for rows.Next() {
		materi.db.DB.ScanRows(rows, &materiFiles)
		responses = append(responses, materiFiles)
	}

	return responses, err
}

// Store implements MaterDefinition
func (materi MateriRepository) Store(request *models.Materi, tx *gorm.DB) (responses *models.Materi, err error) {
	return request, tx.Save(&request).Error
}

// StoreMateriFiles implements MaterDefinition
func (materi MateriRepository) StoreMateriFiles(request *models.MateriRequest, tx *gorm.DB) (responses bool, err error) {
	err = tx.Save(&models.MateriFiles{
		MateriID: request.MateriID,
		FilesID:  request.FilesID,
	}).Error
	fmt.Println(err)
	return true, err
}

// WithTrx implements MaterDefinition
func (materi MateriRepository) WithTrx(trxHandle *gorm.DB) MateriRepository {
	if trxHandle == nil {
		materi.logger.Zap.Error("transaction Database not found in gin context")
		return materi
	}

	materi.db.DB = trxHandle
	return materi
}

// briefing materis
func (materi MateriRepository) GetMateriByActivityAndProduct(request *models.GetMateriByActivityAndProductRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error) {
	query := `
				SELECT  DISTINCT activity_id  as id, product_id as code, judul_materi as name
				FROM briefing_materis
				WHERE activity_id = ?
				AND product_id = ?
			`

	materi.logger.Zap.Info("materi-query-activity-unknown", query)
	rows, err := materi.dbRaw.DB.Query(query, request.ActivityID, request.ProductID)
	defer rows.Close()

	fmt.Println(rows)

	materi.logger.Zap.Info("materi-rows-activity-unknown", rows)
	if err != nil {
		return responses, err
	}

	response := models.GetMateriByActivityAndProductResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Code,
			&response.Name,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// materi by activity, product, risk_issue on verifikasi
func (materi MateriRepository) GetVerifikasiMateri(request *models.GetMateriVerifikasiRequest) (responses []models.GetMateriByActivityAndProductResponseNull, err error) {
	query := `
				SELECT DISTINCT 
					v.risk_indicator_id as id,
					v.product_id as code,
					v.risk_indicator as name
				FROM verifikasi v
				WHERE v.activity_id = ?
				AND v.product_id = ?
				AND v.risk_issue_id = ?
			`

	materi.logger.Zap.Info("materi-query-activity-unknown", query)
	rows, err := materi.dbRaw.DB.Query(query, request.ActivityID, request.ProductID, request.RiskIssueID)
	defer rows.Close()

	// fmt.Println(rows)

	materi.logger.Zap.Info("materi-rows-activity-unknown", rows)
	if err != nil {
		return responses, err
	}

	response := models.GetMateriByActivityAndProductResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.Code,
			&response.Name,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}
