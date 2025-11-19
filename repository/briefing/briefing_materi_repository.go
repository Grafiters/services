package briefing

import (
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type BriefingMateriDefinition interface {
	GetAll() (responses []models.BriefingMateriResponse, err error)
	GetOne(id int64) (responses models.BriefingMateriResponse, err error)
	GetOneBriefing(id int64) (responses []models.BriefingMateriResponses, err error)
	Store(request *models.BriefingMateri, tx *gorm.DB) (responses *models.BriefingMateri, err error)
	Update(request *models.BriefingMateriUpdate, tx *gorm.DB) (responses bool, err error)
	UpdatedIT(request *models.BriefingMateri, tx *gorm.DB) (responses bool, err error)
	// Update(request *models.BriefingMateriRequest) (responses bool, err error)
	Delete(id int64) (err error)
	DeleteBriefingID(id int64, tx *gorm.DB) (err error)
	WithTrx(trxHandle *gorm.DB) BriefingMateriRepository
	GetMateriReport(id int64) (responses []models.MateriReportResponse, err error)
}

type BriefingMateriRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewBriefingMateriRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) BriefingMateriDefinition {
	return BriefingMateriRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetOneBriefing implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) GetOneBriefing(id int64) (responses []models.BriefingMateriResponses, err error) {
	db := BriefingMateri.db.DB

	err = db.Table("briefing_materis bm").
		Select(`bm.id,
				bm.briefing_id, 
				bm.activity_id,
				a.name 'activity_text',
				bm.sub_activity_id,
				sa.name 'sub_activity_text',
				bm.product_id,
				p.product 'product_text',
				bm.title_materies,
				bm.judul_materi,
				bm.risk_issue_code,
				bm.rekomendasi_materi,
				bm.materi_tambahan`).
		Where(`bm.briefing_id = ?`, id).
		Joins(`LEFT JOIN activity a ON a.id = bm.activity_id `).
		Joins(`LEFT JOIN sub_activity sa ON sa.id = bm.sub_activity_id `).
		Joins(`LEFT JOIN product p ON p.id = bm.product_id`).
		Scan(&responses).Error

	// rows, err := BriefingMateri.db.DB.Raw(`
	// 	SELECT bm.*
	// 	FROM briefing_materis bm WHERE bm.briefing_id = ?`, id).Rows()

	// defer rows.Close()
	// var materi models.BriefingMateriResponses

	// for rows.Next() {
	// 	BriefingMateri.db.DB.ScanRows(rows, &materi)
	// 	responses = append(responses, materi)
	// }

	return responses, err
}

// Delete implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) Delete(id int64) (err error) {
	return BriefingMateri.db.DB.Where("id = ?", id).Delete(&models.BriefingMateriResponse{}).Error
}

// DeleteBriefingID implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) DeleteBriefingID(id int64, tx *gorm.DB) (err error) {
	return tx.Where("briefing_id = ?", id).Delete(&models.BriefingMateriResponse{}).Error
}

// GetAll implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) GetAll() (responses []models.BriefingMateriResponse, err error) {
	return responses, BriefingMateri.db.DB.Find(&responses).Error
}

// GetOne implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) GetOne(id int64) (responses models.BriefingMateriResponse, err error) {
	return responses, BriefingMateri.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) Store(request *models.BriefingMateri, tx *gorm.DB) (responses *models.BriefingMateri, err error) {
	return request, tx.Save(&request).Error
}

// Update implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) Update(request *models.BriefingMateriUpdate, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// UpdatedIT implements BriefingMateriDefinition
func (BriefingMateriRepository) UpdatedIT(request *models.BriefingMateri, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements BriefingMateriDefinition
func (BriefingMateri BriefingMateriRepository) WithTrx(trxHandle *gorm.DB) BriefingMateriRepository {
	if trxHandle == nil {
		BriefingMateri.logger.Zap.Error("transacton Database not found in gin context.")
		return BriefingMateri
	}
	BriefingMateri.db.DB = trxHandle
	return BriefingMateri
}

// GetMateriReport implements BriefingMateriDefinition
func (BM BriefingMateriRepository) GetMateriReport(id int64) (responses []models.MateriReportResponse, err error) {
	rows, err := BM.db.DB.Raw(`
			SELECT 
				bm.judul_materi 'judul_materi',
				bm.materi_tambahan 'rincian_materi',
				a.name 'aktifitas'
			FROM briefing_materis bm
			JOIN activity a ON a.id = bm.activity_id 
			WHERE bm.briefing_id = ?`, id).Rows()

	defer rows.Close()
	var materi models.MateriReportResponse

	for rows.Next() {
		BM.db.DB.ScanRows(rows, &materi)
		responses = append(responses, materi)
	}

	return responses, err
}
