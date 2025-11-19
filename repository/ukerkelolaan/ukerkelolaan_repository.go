package ukerkelolaan

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/ukerkelolaan"
	"strconv"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type UkerKelolaanDefinition interface {
	WithTrx(trxHandle *gorm.DB) UkerKelolaanRepository
	GetOne(pn string) (responses models.UkerKelolaanResponse, err error)
	GetListUker(pn string) (responses []models.ListUker, err error)
	GetAllWithPaginate(request *models.KeywordRequest) (responses []models.UkerKelolaanResponse, totalData int, totalRows int, err error)
	Store(request *models.UkerKelolaan, tx *gorm.DB) (responses *models.UkerKelolaan, err error)
	Update(request *models.UkerKelolaan, include []string, tx *gorm.DB) (responses bool, err error)
	Delete(request *models.UkerKelolaanReqDelete, include []string, tx *gorm.DB) (response bool, err error)
	FilterUkerKelolaan(request *models.KeywordRequest) (responses []models.UkerKelolaanResponse, totalData int, totalRows int, err error)
	CekBRCKelolaan(PERNR string) (count int, err error)
	StoreMstUker(request *models.SaveMstRequest, tx *gorm.DB) (responses bool, err error)
	GetDetailData(id int64) (response models.UkerKelolaanResponse, err error)
	GetListUkerKelolaan(request *models.PencarianUker) (response []models.UkerList, err error)
}

type UkerKelolaanRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewUkerKelolaanRepository(
	db lib.Database,
	logger logger.Logger,
) UkerKelolaanDefinition {
	return UkerKelolaanRepository{
		db:      db,
		logger:  logger,
		timeout: 0,
	}
}

// WithTrx implements UkerKelolaanDefinition
func (uk UkerKelolaanRepository) WithTrx(trxHandle *gorm.DB) UkerKelolaanRepository {
	if trxHandle == nil {
		uk.logger.Zap.Error("transaction Database not found in gin context")
		return uk
	}

	uk.db.DB = trxHandle
	return uk
}

func (UkerKelolaanRepository) StoreMstUker(request *models.SaveMstRequest, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(request).Error
}

// GetDetailData implements UkerKelolaanDefinition.
func (uk UkerKelolaanRepository) GetDetailData(id int64) (response models.UkerKelolaanResponse, err error) {
	db := uk.db.DB
	err = db.Table("mst_uker_kelolaan muk").
		Select(`
				muk.id 'id',
				muk.pn 'pn',
				muk.sname 'SNAME',
				muk.aktif 'status'`).
		Where("id = ?", id).
		Scan(&response).Error

	return response, err
}

// Delete implements UkerKelolaanDefinition
func (UkerKelolaanRepository) Delete(request *models.UkerKelolaanReqDelete, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Where("id = ?", request.ID).Delete(&models.UkerKelolaan{}).Error
	// return true, tx.Save(request).Error

}

// GetAllWithPaginate implements UkerKelolaanDefinition
func (uk UkerKelolaanRepository) GetAllWithPaginate(request *models.KeywordRequest) (responses []models.UkerKelolaanResponse, totalData int, totalRows int, err error) {
	db := uk.db.DB

	// db = db.Table("uker_kelolaan_user").
	// 	Select(`
	// 		id,
	// 		created_at,
	// 		updated_at
	// 		expired_at,
	// 		is_temp,
	// 		pn,
	// 		SNAME,
	// 		REGION,
	// 		RGDESC,
	// 		MAINBR,
	// 		MBDESC,
	// 		BRANCH,
	// 		BRDESC,
	// 		status`).Group("pn").Order("created_at ASC")

	// if request.REGION != "" {
	// 	db = db.Where("REGION = ?", request.REGION)
	// }

	db = db.Table("mst_uker_kelolaan muk").
		Select(`
				muk.id 'id',
				muk.pn 'pn',
				muk.sname 'SNAME',
				muk.aktif 'status'`).
		Joins(`LEFT JOIN uker_kelolaan_user uku ON uku.pn = muk.pn`).
		Group(`muk.pn`)

	if request.REGION != "" {
		db = db.Where(`uku.REGION = ?`, request.REGION)
	}

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return responses, totalData, totalRows, err
	}

	totalData = int(count)

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error
	if err != nil {
		return responses, totalData, totalRows, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	return responses, totalData, totalRows, err
}

func totalRows(totalRows int) {
	panic("unimplemented")
}

// Store implements UkerKelolaanDefinition
func (UkerKelolaanRepository) Store(request *models.UkerKelolaan, tx *gorm.DB) (responses *models.UkerKelolaan, err error) {
	return request, tx.Save(&request).Error
}

// Update implements UkerKelolaanDefinition
func (UkerKelolaanRepository) Update(request *models.UkerKelolaan, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// FilterUkerKelolaan implements UkerKelolaanDefinition
func (uk UkerKelolaanRepository) FilterUkerKelolaan(request *models.KeywordRequest) (responses []models.UkerKelolaanResponse, totalData int, totalRows int, err error) {
	db := uk.db.DB

	// db = db.Table("uker_kelolaan_user uku").
	// 	Select(`
	// 		id,
	// 		created_at,
	// 		updated_at
	// 		expired_at,
	// 		is_temp,
	// 		pn,
	// 		SNAME,
	// 		REGION,
	// 		RGDESC,
	// 		MAINBR,
	// 		MBDESC,
	// 		BRANCH,
	// 		BRDESC,
	// 		status`).
	// 	Where("status = ?", request.Status).
	// 	Group("pn").
	// 	Order("created_at ASC")

	db = db.Table("mst_uker_kelolaan muk").
		Select(`
				muk.id,
				muk.pn,
				muk.sname 'SNAME',
				muk.aktif 'status'`).
		Joins(`LEFT JOIN uker_kelolaan_user uku ON uku.pn = muk.pn`).
		Where(`muk.aktif = ?`, request.Status).
		Group(`muk.pn`)

	if request.Pn != "Semua" && request.Pn != "" {
		db = db.Where("muk.pn = ?", request.Pn)
	}

	if request.REGION != "all" && request.REGION != "" {
		db = db.Where("uku.REGION = ?", request.REGION)
	}

	if request.MAINBR != "all" && request.MAINBR != "" {
		db = db.Where("uku.MAINBR = ?", request.MAINBR)
	}

	if request.BRANCH != "all" && request.BRANCH != "" {
		db = db.Where("uku.BRANCH = ?", request.BRANCH)
	}

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return responses, totalData, totalRows, err
	}

	totalData = int(count)
	fmt.Println("TotalRows =>", totalData)

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil((float64(totalData)) / float64(request.Limit)))
	}

	return responses, totalData, totalRows, err
}

// CekBRCKelolaan implements UkerKelolaanDefinition
func (uk UkerKelolaanRepository) CekBRCKelolaan(PERNR string) (count int, err error) {
	db := uk.db.DB

	// db = db.Table("uker_kelolaan_user").
	// 	Select(`COUNT(*) 'count'`).
	// 	Where(`pn = ?`, PERNR).
	// 	Where(`status = true`)

	db = db.Table("mst_uker_kelolaan").
		Select(`COUNT(*) 'count'`).
		Where(`pn = ?`, PERNR)

	err = db.Scan(&count).Error

	if err != nil {
		return count, err
	}

	return count, err
}

// GetOne implements UkerKelolaanDefinition
func (uk UkerKelolaanRepository) GetOne(pn string) (responses models.UkerKelolaanResponse, err error) {
	db := uk.db.DB
	fmt.Println("iniPN =>", pn)

	db = db.Table("uker_kelolaan_user").
		Select(`
			id,
			created_at,
			updated_at
			expired_at,
			is_temp,
			pn,
			SNAME,
			REGION,
			RGDESC,
			MAINBR,
			MBDESC,
			BRANCH,
			BRDESC,
			status`).
		Where("pn = ?", pn).
		Where("status = 1").
		Group("pn")

	err = db.Find(&responses).Error

	return responses, err
}

// GetListUker implements UkerKelolaanDefinition
func (uk UkerKelolaanRepository) GetListUker(pn string) (responses []models.ListUker, err error) {
	db := uk.db.DB

	db = db.Table("uker_kelolaan_user").
		Select(`
		id,
		REGION,
		RGDESC,
		MAINBR,
		MBDESC,
		BRANCH,
		BRDESC
	`).Where("pn = ?", pn)
	// .Where("status = 1")

	err = db.Find(&responses).Error

	return responses, err
}

// GetListUkerKelolaan implements UkerKelolaanDefinition.
func (uk UkerKelolaanRepository) GetListUkerKelolaan(request *models.PencarianUker) (responses []models.UkerList, err error) {
	db := uk.db.DB.Table("uker_kelolaan_user uku")

	db = db.Select(`
		REGION,
		RGDESC,
		MAINBR,
		MBDESC,
		BRANCH,
		BRDESC
	`).Where(`uku.status = 1`)

	if request.Keyword != "" {
		db = db.Where(`CONCAT(BRANCH, "-", BRDESC) like ?`, fmt.Sprintf("%%%s%%", request.Keyword))
	}

	if request.TIPEUKER != "KP" {
		if request.TIPEUKER == "KW" {
			REGION, err := uk.GetRegion(request.BRANCH)
			if err != nil {
				return responses, err
			}

			ada, err := uk.CekBRCKelolaan(request.PERNR)
			if err != nil {
				return responses, err
			}

			if ada > 0 {
				db = db.Where("uku.pn = ?", request.PERNR)
			}

			db = db.Where("uku.REGION = ?", REGION).Group("uku.BRANCH")
		} else {
			db = db.Where("uku.pn = ?", request.PERNR)
		}
	} else {
		ada, err := uk.CekUkerKelolaan(request.PERNR)
		if err != nil {
			return responses, err
		}

		if ada > 0 {
			db = db.Where("uku.pn = ?", request.PERNR)
		}

		db = db.Group("uku.BRANCH")
	}

	db.Limit(int(request.Limit)).Offset(int(request.Offset))

	err = db.Find(&responses).Error
	if err != nil {
		return responses, err
	}

	ada, err := uk.CekPgsActive(request.PERNR, "")
	if err != nil {
		return responses, err
	}

	if ada > 0 {
		fmt.Println("Ada => ", ada)
		pgsQuery := uk.db.DB.Table("pgs_user").
			Select(`
				REGION,
				RGDESC,
				MAINBR,
				MBDESC,
				BRANCH,
				BRDESC
			`).
			Where(`pn = ?`, request.PERNR).
			Where(`delete_flag = 0`).
			Where(`status IN ("02a")`).
			Where(`action = 'Active'`).
			Limit(int(request.Limit)).
			Offset(int(request.Offset))

		if request.Keyword != "" {
			pgsQuery = pgsQuery.Where(`CONCAT(BRANCH, "-", BRDESC) like ?`, fmt.Sprintf("%%%s%%", request.Keyword))
		}

		var additionalResponses []models.UkerList // Replace YourResponseType with the actual type of the response

		pgsQuery.Find(&additionalResponses)

		// Append additionalResponses to the responses slice
		responses = append(responses, additionalResponses...)
	}

	return responses, err
}

func (uk UkerKelolaanRepository) CekPgsActive(PERNR string, BRANCH string) (count int, err error) {
	db := uk.db.DB

	db = db.Table("pgs_user").
		Select(`COUNT(*) 'count'`).
		Where(`delete_flag = 0`).
		Where(`status != '03a'`)
		// Where(`status IN ("02a","01a","01b")`).
		// Where(`action IN ("Active", "New Request", "UpdateData")`)

	if BRANCH != "" {
		db = db.Where(`pn = ? OR BRANCH = ?`, PERNR, BRANCH)
	} else {
		db = db.Where(`pn = ?`, PERNR)
	}

	err = db.Scan(&count).Error

	if err != nil {
		return count, err
	}

	return count, err
}

func (uk UkerKelolaanRepository) GetRegion(BRANCH string) (REGION string, err error) {
	kodeBranch, err := strconv.Atoi(BRANCH)

	if err != nil {
		return "", err
	}

	db := uk.db.DB.Table(`dwh_branch`).Select(`REGION`).Where(`BRANCH = ?`, kodeBranch)

	err = db.Find(&REGION).Error

	return REGION, err
}

func (uk UkerKelolaanRepository) CekUkerKelolaan(PERNR string) (count int, err error) {
	// panic("unimplemented")
	db := uk.db.DB

	db = db.Table("uker_kelolaan_user").
		Select(`COUNT(*) 'count'`).Where(`pn = ?`, PERNR).Where(`status = 1`)

	err = db.Scan(&count).Error

	if err != nil {
		return count, err
	}

	return count, err
}
