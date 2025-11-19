package pgsuser

import (
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/pgsuser"
	"strconv"
	"time"

	"fmt"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type PgsUserDefinition interface {
	WithTrx(trxHandle *gorm.DB) PgsUserRepository
	GetAll(makerID string) (responses []models.PgsUserResponses, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.PgsUserResponses, totalData int, totalRows int, err error)
	GetOne(id int64) (responses models.PgsUserResponses, err error)
	Store(request *models.PgsUser, tx *gorm.DB) (responses *models.PgsUser, err error)
	Update(request *models.PgsUserRequestMaintainance, include []string, tx *gorm.DB) (responses bool, err error)
	GetPgsApproval(request *models.Paginate) (responses []models.PgsApprovalResponseNull, totalData int, totalRows int, err error)
	ApprovePgsUser(request *models.PgsUpdateRequest, include []string, tx *gorm.DB) (responses bool, err error)
	CekPGSAda(pernr string) (responses models.PgsUserResponses, err error)
	Delete(request *models.UpdateDelete, include []string, tx *gorm.DB) (response bool, err error)
	SearchPekerjaByPn(PERNR string) (responses models.UserResponseLocal, err error)
	CekPgsActive(PERNR string, BRANCH string) (count int, err error)

	// Batch 3
	GetUkerBinaan(request *models.UkerKelolaanRequest) (responses []models.UnitKerja, err error)
	GetRegion(BRANCH string) (REGION string, err error)
	CekUkerKelolaan(PERNR string) (count int, err error)
}

type PgsUserRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

// Delete implements PgsUserDefinition
func (PgsUserRepository) Delete(request *models.UpdateDelete, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(request).Error
}

func NewPgsUserRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) PgsUserDefinition {
	return PgsUserRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: 0,
	}
}

// GetAll implements PgsUserDefinition
func (pgsUser PgsUserRepository) GetAll(makerID string) (responses []models.PgsUserResponses, err error) {
	return responses, pgsUser.db.DB.Where("maker_id = ? AND delete_flag != 1", makerID).Find(&responses).Error
}

// GetAllWithPaginate implements PgsUserDefinition
func (pgs PgsUserRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.PgsUserResponses, totalData int, totalRows int, err error) {
	rows, err := pgs.db.DB.Raw(`
	SELECT 
		id,
		pn,
		nama_pekerja,
		CONCAT(BRANCH," - ", BRDESC) 'unit_kerja',
		jabatan_pgs,
		status,
		action
	FROM pgs_user WHERE maker_id = ? AND delete_flag != 1 ORDER BY id LIMIT ? OFFSET ?
	`, request.Penr, request.Limit, request.Offset).Rows()

	defer rows.Close()

	var roles models.PgsUserResponses

	for rows.Next() {
		pgs.db.DB.ScanRows(rows, &roles)
		responses = append(responses, roles)
	}

	paginateQuery := `SELECT COUNT(*) FROM pgs_user WHERE maker_id = ? AND delete_flag != 1`
	err = pgs.dbRaw.DB.QueryRow(paginateQuery, request.Penr).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))
	return responses, resultFinal, totalRows, err
}

// GetOne implements PgsUserDefinition
func (pgsUser PgsUserRepository) GetOne(id int64) (responses models.PgsUserResponses, err error) {
	err = pgsUser.db.DB.Raw(`SELECT * FROM pgs_user WHERE id = ?`, id).Find(&responses).Error
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// Store implements PgsUserDefinition
func (pgsUser PgsUserRepository) Store(request *models.PgsUser, tx *gorm.DB) (responses *models.PgsUser, err error) {
	return request, tx.Save(&request).Error
}

// Update implements PgsUserDefinition
func (pgsUser PgsUserRepository) Update(request *models.PgsUserRequestMaintainance, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements PgsUserDefinition
func (pgsUser PgsUserRepository) WithTrx(trxHandle *gorm.DB) PgsUserRepository {
	if trxHandle == nil {
		pgsUser.logger.Zap.Error("transaction Database not found in gin context")
		return pgsUser
	}

	pgsUser.db.DB = trxHandle
	return pgsUser
}

// GetPgsApproval implements PgsUserDefinition
func (pgsUser PgsUserRepository) GetPgsApproval(request *models.Paginate) (responses []models.PgsApprovalResponseNull, totalData int, totalRows int, err error) {
	query := `SELECT 
				pgs.id 'id',
				pgs.pn 'pn',
				pgs.nama_pekerja 'nama_pekerja',
				concat(pgs.BRANCH," - ",pgs.BRDESC) 'unit_kerja',
				pgs.jabatan_pgs 'jabatan_pgs',
				approval.id 'id_approval',
				approval.approval_id 'approval',
				approval.approval_date 'approval_date',
				approval.approval_status 'approval_status'
			FROM pgs_user pgs
			JOIN pgs_user_approval approval ON approval.id_pgs_user = pgs.id
			WHERE pgs.status like '%01%' AND approval.approval_id = ? AND approval.approval_status = "0" ORDER BY pgs.created_at DESC LIMIT ? OFFSET ?`

	pgsUser.logger.Zap.Info(query)
	rows, err := pgsUser.dbRaw.DB.Query(query, request.Penr, request.Limit, request.Offset)
	defer rows.Close()

	pgsUser.logger.Zap.Info("rows =>", rows)
	if err != nil {
		return responses, totalData, totalRows, err
	}

	response := models.PgsApprovalResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.PN,
			&response.NamaPekerja,
			&response.UnitKerja,
			&response.JabatanPgs,
			&response.IDApproval,
			&response.Approval,
			&response.ApprovalDate,
			&response.ApprovalStatus,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalData, totalRows, err
	}

	paginateQuery := `SELECT COUNT(*) FROM pgs_user pgs
				JOIN pgs_user_approval approval ON approval.id_pgs_user = pgs.id
				WHERE pgs.status like '%01%' AND approval.approval_id = ? AND approval.approval_status = "0"`

	err = pgsUser.dbRaw.DB.QueryRow(paginateQuery, request.Penr).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))
	return responses, resultFinal, totalRows, err
}

// ApprovePgsUser implements PgsUserDefinition
func (pgsUser PgsUserRepository) ApprovePgsUser(request *models.PgsUpdateRequest, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// CekPGSAda implements PgsUserDefinition
func (pgsUser PgsUserRepository) CekPGSAda(pernr string) (responses models.PgsUserResponses, err error) {
	err = pgsUser.db.DB.Raw(`SELECT * FROM pgs_user WHERE PN = ?`, pernr).Find(&responses).Error
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

// GetUkerBinaan implements PgsUserDefinition
func (pgs PgsUserRepository) GetUkerBinaan(request *models.UkerKelolaanRequest) (responses []models.UnitKerja, err error) {
	db := pgs.db.DB.Table("uker_kelolaan_user uku")

	db = db.Select(`
		uku.REGION 'REGION',
		uku.RGDESC 'RGDESC',
		uku.MAINBR 'MAINBR',
		uku.MBDESC 'MBDESC',
		uku.BRANCH 'BRANCH',
		uku.BRDESC 'BRDESC'
	`).Where("uku.status = 1")

	fmt.Println("tipe_uker =>>", request.TIPEUKER)
	if request.TIPEUKER != "KP" {
		if request.TIPEUKER == "KW" {
			REGION, err := pgs.GetRegion(request.BRANCH)
			if err != nil {
				return responses, err
			}

			ada, err := pgs.CekUkerKelolaan(request.PERNR)
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
		ada, err := pgs.CekUkerKelolaan(request.PERNR)
		if err != nil {
			return responses, err
		}

		if ada > 0 {
			db = db.Where("uku.pn = ?", request.PERNR)
		}

		db = db.Group("uku.BRANCH")
	}

	// if request.TIPEUKER == "KP" {
	// 	db = db.Group("uku.BRANCH")
	// }

	err = db.Find(&responses).Error
	if err != nil {
		return responses, err
	}

	ada, err := pgs.CekPgsActive(request.PERNR, "")
	if err != nil {
		return responses, err
	}

	if ada > 0 {
		pgsQuery := pgs.db.DB.Table("pgs_user").
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
			Where(`action = 'Active'`)

		var additionalResponses []models.UnitKerja // Replace YourResponseType with the actual type of the response

		pgsQuery.Find(&additionalResponses)

		// Append additionalResponses to the responses slice
		responses = append(responses, additionalResponses...)
	}

	return responses, err
}

func (repo PgsUserRepository) SearchPekerjaByPn(PERNR string) (responses models.UserResponseLocal, err error) {
	query := repo.db.DB

	query = query.Table("pa0001_eof pe").
		Select(`
				PERNR,
				WERKS,
				BTRTL,
				KOSTL,
				ORGEH,
				ORGEH_PGS 'ORGEHPGS',
				STELL,
				SNAME,
				WERKS_TX 'WERKSTX',
				BTRTL_TX 'BTRTLTX',
				KOSTL_TX 'KOSTLTX',
				ORGEH_TX 'ORGEHTX',
				ORGEH_PGS_TX 'ORGEHPGSTX',
				STELL_TX STELLTX,
				BRANCH,
				TIPE_UKER 'TIPEUKER',
				HILFM,
				HILFM_PGS 'HILFMPGS',
				HTEXT,
				HTEXT_PGS 'HTEXTPGS'
			`).Where("pe.PERNR = ?", PERNR).Find(&responses)

	return responses, err
}

// versioning 24/10/2023 by panji
// CekPgsActive implements PgsUserDefinition
func (pgs PgsUserRepository) CekPgsActive(PERNR string, BRANCH string) (count int, err error) {
	db := pgs.db.DB

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

func (pgs PgsUserRepository) GetRegion(BRANCH string) (REGION string, err error) {
	kodeBranch, err := strconv.Atoi(BRANCH)

	db := pgs.db.DB.Table(`dwh_branch`).Select(`REGION`).Where(`BRANCH = ?`, kodeBranch)

	err = db.Find(&REGION).Error

	return REGION, err
}

// CekUkerKelolaan implements PgsUserDefinition.
func (pgs PgsUserRepository) CekUkerKelolaan(PERNR string) (count int, err error) {
	// panic("unimplemented")
	db := pgs.db.DB

	db = db.Table("uker_kelolaan_user").
		Select(`COUNT(*) 'count'`).Where(`pn = ?`, PERNR).Where(`status = 1`)

	err = db.Scan(&count).Error

	if err != nil {
		return count, err
	}

	return count, err
}
