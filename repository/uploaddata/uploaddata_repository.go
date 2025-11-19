package uploaddata

import (
	"fmt"
	"riskmanagement/lib"
	model "riskmanagement/models/uploaddata"
	models "riskmanagement/models/uploaddata"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type UploadDataDefinition interface {
	WithTrx(trxHandle *gorm.DB) UploadDataRepository
	UploadRiskControl(request models.RiskControlRequest, tx *gorm.DB) (response bool, err error)
	GetKodeRiskControl() (numbered int, err error)
	CekRiskControl(riskcontrol string) (ada int, err error)
	UploadRisknIndicator(request models.RiskIndicatorRequest, tx *gorm.DB) (response bool, err error)
	CekRiskIndicator(kode string, riskindicator string) (ada int, err error)
	UploadRiskIssue(request *models.RiskEvent, tx *gorm.DB) (response *models.RiskEvent, err error)
	MapProses(request *models.MapProsesRequest, tx *gorm.DB) (response bool, err error)
	MapEvent(request *models.MapEventRequest, tx *gorm.DB) (response bool, err error)
	MapKejadian(request *models.MapKejadianRequest, tx *gorm.DB) (response bool, err error)
	MapProduct(request *models.MapProductRequest, tx *gorm.DB) (response bool, err error)
	MapLiniBisnis(request *models.MapLiniBisnisRequest, tx *gorm.DB) (response bool, err error)
	MapAktifitas(request *models.MapAktifitasRequest, tx *gorm.DB) (response bool, err error)
	CekRiskEvent(riskevent string) (ada int, err error)
	GetCounterEvent(kode string) (numbered int, err error)
	ValidasiMegaProses(request string) (response bool, err error)
	ValidasiMajorProses(request string) (response bool, err error)
	ValidasiSubMajorProses(request string) (response bool, err error)
	ValidasiEventLv1(request string) (response bool, err error)
	ValidasiEventLv2(request string) (response bool, err error)
	ValidasiEventLv3(request string) (response bool, err error)
	ValidasiLiniBisnisLv1(request string) (response bool, err error)
	ValidasiLiniBisnisLv2(request string) (response bool, err error)
	ValidasiLiniBisnisLv3(request string) (response bool, err error)
	ValidasiKejadianLv1(request string) (response bool, err error)
	ValidasiKejadianLv2(request string) (response bool, err error)
	ValidasiKejadianLv3(request string) (response bool, err error)
}

type UploadDataRepository struct {
	db     lib.Database
	dbRaw  lib.Databases
	logger logger.Logger
}

func NewUploadDataRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) UploadDataDefinition {
	return UploadDataRepository{
		db:     db,
		dbRaw:  dbRaw,
		logger: logger,
	}
}

// WithTrx implements RiskControlDefinition
func (upload UploadDataRepository) WithTrx(trxHandle *gorm.DB) UploadDataRepository {
	if trxHandle == nil {
		upload.logger.Zap.Error("transaction Database not found in gin context")
		return upload
	}

	upload.db.DB = trxHandle
	return upload
}

// UploadRiskControl implements UploadDataDefinition.
func (upload UploadDataRepository) UploadRiskControl(request models.RiskControlRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

// GetKodeRiskControl implements UploadDataDefinition.
func (upload UploadDataRepository) GetKodeRiskControl() (numbered int, err error) {
	err = upload.db.DB.Raw(`SELECT kode_risk_control 'numbered' FROM (SELECT 
		CASE
				WHEN LENGTH(kode) = 5 THEN CAST(RIGHT(kode,4) AS DECIMAL) 
				WHEN LENGTH(kode) = 4  THEN CAST(RIGHT(kode,3) AS DECIMAL)
				WHEN LENGTH(kode) = 3  THEN CAST(RIGHT(kode,2) AS DECIMAL)
				WHEN LENGTH(kode) = 2  THEN CAST(RIGHT(kode,1) AS DECIMAL)
		END + 1 'kode_risk_control' 
		FROM risk_control rc) as t
		ORDER BY t.kode_risk_control DESC LIMIT 1`).Scan(&numbered).Error

	if err != nil {
		return numbered, err
	}

	return numbered, err
}

// CekRiskControl implements UploadDataDefinition.
func (upload UploadDataRepository) CekRiskControl(riskcontrol string) (ada int, err error) {
	db := upload.db.DB.Table("risk_control").
		Select(`COUNT(*) 'ada'`).
		Where(`risk_control = ?`, riskcontrol)

	err = db.Scan(&ada).Error

	return ada, err
}

// UploadRisknIndicator implements UploadDataDefinition.
func (upload UploadDataRepository) UploadRisknIndicator(request models.RiskIndicatorRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

// CekRiskIndicator implements UploadDataDefinition.
func (upload UploadDataRepository) CekRiskIndicator(kode string, riskindicator string) (ada int, err error) {
	db := upload.db.DB.Table("risk_indicator").
		Select(`COUNT(*) 'ada'`).
		Where(`risk_indicator_code = ?`, kode).
		Where(`risk_indicator = ?`, riskindicator).
		Where(`delete_flag = 0`)

	err = db.Scan(&ada).Error

	return ada, err
}

// UploadRiskIssue implements UploadDataDefinition.
func (upload UploadDataRepository) UploadRiskIssue(request *models.RiskEvent, tx *gorm.DB) (response *models.RiskEvent, err error) {
	input := &model.RiskEvent{
		RiskTypeID:     request.RiskTypeID,
		RiskIssueCode:  request.RiskIssueCode,
		RiskIssue:      request.RiskIssue,
		Deskripsi:      request.Deskripsi,
		KategoriRisiko: request.KategoriRisiko,
		Status:         request.Status,
		Likelihood:     request.Likelihood,
		Impact:         request.Impact,
		CreatedAt:      request.CreatedAt,
		DeleteFlag:     false,
	}

	err = tx.Save(input).Error

	return input, err
}

// MapAktifitas implements UploadDataDefinition.
func (upload UploadDataRepository) MapAktifitas(request *models.MapAktifitasRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

// MapEvent implements UploadDataDefinition.
func (upload UploadDataRepository) MapEvent(request *models.MapEventRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

// MapKejadian implements UploadDataDefinition.
func (upload UploadDataRepository) MapKejadian(request *models.MapKejadianRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

// MapLiniBisnis implements UploadDataDefinition.
func (upload UploadDataRepository) MapLiniBisnis(request *models.MapLiniBisnisRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

// MapProduct implements UploadDataDefinition.
func (upload UploadDataRepository) MapProduct(request *models.MapProductRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

// MapProses implements UploadDataDefinition.
func (upload UploadDataRepository) MapProses(request *models.MapProsesRequest, tx *gorm.DB) (response bool, err error) {
	return response, tx.Save(&request).Error
}

func (upload UploadDataRepository) CekRiskEvent(riskevent string) (ada int, err error) {
	db := upload.db.DB.Table("risk_issue").
		Select(`COUNT(*) 'ada'`).
		Where(`delete_flag = 0`).
		// Where(`risk_issue_code = ?`, kode).
		Where(`risk_issue = ?`, riskevent)

	err = db.Scan(&ada).Error

	return ada, err
}

// GetCounterEvent implements UploadDataDefinition.
func (upload UploadDataRepository) GetCounterEvent(kode string) (numbered int, err error) {
	db := upload.db.DB.Table("risk_issue").
		Select(`COUNT(*) 'numbered'`).
		Where(`delete_flag = 0`).
		Where(`risk_issue_code LIKE ?`, fmt.Sprintf("%%%s%%", kode))

	err = db.Scan(&numbered).Error

	return numbered, err
}

// ValidasiEventLv1 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiEventLv1(request string) (response bool, err error) {
	db := upload.db.DB.Table("event_type_lv1").Where("kode_event_type = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiEventLv2 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiEventLv2(request string) (response bool, err error) {
	db := upload.db.DB.Table("event_type_lv2").Where("kode_event_type_lv2 = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiEventLv3 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiEventLv3(request string) (response bool, err error) {
	db := upload.db.DB.Table("event_type_lv3").Where("kode_event_type_lv3 = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiKejadianLv1 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiKejadianLv1(request string) (response bool, err error) {
	db := upload.db.DB.Table("penyebab_kejadian_lv1").Where("kode_kejadian = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiKejadianLv2 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiKejadianLv2(request string) (response bool, err error) {
	db := upload.db.DB.Table("penyebab_kejadian_lv2").Where("kode_sub_kejadian = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiKejadianLv3 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiKejadianLv3(request string) (response bool, err error) {
	db := upload.db.DB.Table("penyebab_kejadian_lv3").Where("kode_penyebab_kejadian_lv3 = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiLiniBisnisLv1 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiLiniBisnisLv1(request string) (response bool, err error) {
	db := upload.db.DB.Table("lini_bisnis_lv1").Where("kode_lini_bisnis = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiLiniBisnisLv2 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiLiniBisnisLv2(request string) (response bool, err error) {
	db := upload.db.DB.Table("lini_bisnis_lv2").Where("kode_lini_bisnis_lv2 = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiLiniBisnisLv3 implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiLiniBisnisLv3(request string) (response bool, err error) {
	db := upload.db.DB.Table("lini_bisnis_lv3").Where("kode_lini_bisnis_lv3 = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiMajorProses implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiMajorProses(request string) (response bool, err error) {
	db := upload.db.DB.Table("major_proses").Where("kode_major_proses = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiMegaProses implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiMegaProses(request string) (response bool, err error) {
	db := upload.db.DB.Table("mega_proses").Where("kode_mega_proses = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}

// ValidasiSubMajorProses implements UploadDataDefinition.
func (upload UploadDataRepository) ValidasiSubMajorProses(request string) (response bool, err error) {
	db := upload.db.DB.Table("sub_major_proses").Where("kode_sub_major_proses = ?", request)

	var totalData int64
	err = db.Count(&totalData).Error

	if totalData < 1 {
		return false, err
	}

	return true, nil
}
