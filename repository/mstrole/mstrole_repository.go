package mstrole

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/mstrole"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MstRoleDefinition interface {
	GetAll() (responses []models.MstRoleResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.MstRoleResponse, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.MstRoleResponse, err error)
	Store(request *models.MstRole, tx *gorm.DB) (responses *models.MstRole, err error)
	Update(request *models.MstRoleRequestUpdate, include []string, tx *gorm.DB) (responses bool, err error)
	Delete(request *models.MstRoleRequestDelete, include []string, tx *gorm.DB) (responses bool, err error)
	GetMenuList(request models.MenuListRequest) (responses models.MstRoleResponseOne, err error)
	GetMenuListQuestionnaire(request models.MenuListRequest) (responses models.MstRoleQuestionnaireResponseOne, err error)

	WithTrx(trxHandle *gorm.DB) MstRoleRepository
}

type MstRoleRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMstRoleRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MstRoleDefinition {
	return MstRoleRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (repo MstRoleRepository) GetMenuList(request models.MenuListRequest) (responses models.MstRoleResponseOne, err error) {
	db := repo.db.DB

	query := db.Table("mst_roles_batch3 role").
		Select(`
			role.id,
			role.role_name,
			role.menu,
			role.additional_menu,
			role.mfe_menu
		`).
		Joins("LEFT JOIN management_user_batch3 mu ON mu.role_id = role.id").
		Joins("LEFT JOIN mst_jabatan mj ON mj.id = mu.level_id").
		Where(`(mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker =  ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker =  ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND mu.addon_pernr LIKE ?)`+
			`OR (mu.level_uker = ? AND mu.addon_pernr LIKE ?)`+
			`OR (mu.level_uker = ? AND mu.addon_pernr LIKE ?)`,
			"ALL", "ALL", "ALL",
			request.TIPEUKER, "ALL", "ALL",
			"ALL", request.HILFM, request.Jgpg,
			request.TIPEUKER, request.HILFM, request.Jgpg,
			request.KOSTL, request.HILFM, request.Jgpg,
			request.ORGEH, request.HILFM, request.Jgpg,
			request.TIPEUKER, fmt.Sprintf("%%%s%%", request.PERNR),
			request.KOSTL, fmt.Sprintf("%%%s%%", request.PERNR),
			request.ORGEH, fmt.Sprintf("%%%s%%", request.PERNR),
		)
	query.Find(&responses)

	return responses, err
}

func (repo MstRoleRepository) GetMenuListQuestionnaire(request models.MenuListRequest) (responses models.MstRoleQuestionnaireResponseOne, err error) {
	db := repo.db.DB

	query := db.Table("mst_roles_questionnaire").
		Select(`
			mst_roles_batch3.role_name,
			GROUP_CONCAT(TRIM(mst_roles_questionnaire.menu_id) SEPARATOR ',') AS 'menu'
		`).
		Joins("LEFT JOIN mst_roles_batch3 ON mst_roles_batch3.id = mst_roles_questionnaire.role_id").
		Joins("LEFT JOIN management_user_batch3 mu ON mu.role_id = mst_roles_batch3.id").
		Joins("LEFT JOIN mst_jabatan mj ON mj.id = mu.level_id").
		Where(`(mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker =  ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker =  ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND (mj.hilfm = ? AND mj.jgpg = ?))`+
			`OR (mu.level_uker = ? AND mu.addon_pernr LIKE ?)`+
			`OR (mu.level_uker = ? AND mu.addon_pernr LIKE ?)`+
			`OR (mu.level_uker = ? AND mu.addon_pernr LIKE ?)`,
			"ALL", "ALL", "ALL",
			request.TIPEUKER, "ALL", "ALL",
			"ALL", request.HILFM, request.Jgpg,
			request.TIPEUKER, request.HILFM, request.Jgpg,
			request.KOSTL, request.HILFM, request.Jgpg,
			request.ORGEH, request.HILFM, request.Jgpg,
			request.TIPEUKER, fmt.Sprintf("%%%s%%", request.PERNR),
			request.KOSTL, fmt.Sprintf("%%%s%%", request.PERNR),
			request.ORGEH, fmt.Sprintf("%%%s%%", request.PERNR),
		)

	query.Group("mst_roles_questionnaire.role_id")

	query.Find(&responses)

	return responses, err
}

// Delete implements MstRoleDefinition
func (mstRole MstRoleRepository) Delete(request *models.MstRoleRequestDelete, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// GetAllWithPaginate implements MstRoleDefinition
func (msRole MstRoleRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.MstRoleResponse, totalRows int, totalData int, err error) {
	rows, err := msRole.db.DB.Raw(`
	SELECT 
		mr.id 'id',
		mr.role_name 'role_name',
		mr.delete_flag 'delete_flag',
		mr.created_at 'created_at',
		mr.updated_at 'updated_at'
	FROM mst_roles_batch3 mr 
	WHERE mr.delete_flag != 1 ORDER BY mr.id LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()

	defer rows.Close()

	var roles models.MstRoleResponse

	for rows.Next() {
		msRole.db.DB.ScanRows(rows, &roles)
		responses = append(responses, roles)
	}

	paginateQuery := `SELECT COUNT(*) FROM mst_roles_batch3`
	err = msRole.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))
	return responses, resultFinal, totalRows, err
}

// GetAll implements MstRoleDefinition
func (mstRole MstRoleRepository) GetAll() (responses []models.MstRoleResponse, err error) {
	return responses, mstRole.db.DB.Raw(`SELECT * FROM mst_roles_batch3 WHERE delete_flag != '1'`).Find(&responses).Error
}

// GetOne implements MstRoleDefinition
func (mstRole MstRoleRepository) GetOne(id int64) (responses models.MstRoleResponse, err error) {
	db := mstRole.db.DB

	query := db.Table("mst_roles_batch3").
		Select(`
				id, 
				menu,
				additional_menu,
				role_name,
				addon_pernr,
				delete_flag,
				created_at
			`).Where("id = ?", id)

	err = query.Find(&responses).Error

	// err = mstRole.db.DB.Raw(`SELECT * from mst_roles_batch3 WHERE id = ?`, id).Find(&responses).Error
	if err != nil {
		mstRole.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

// Store implements MstRoleDefinition
func (mstRole MstRoleRepository) Store(request *models.MstRole, tx *gorm.DB) (responses *models.MstRole, err error) {
	return request, tx.Save(&request).Error
}

// Update implements MstRoleDefinition
func (mstRole MstRoleRepository) Update(request *models.MstRoleRequestUpdate, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// WithTrx implements MstRoleDefinition
func (mstRole MstRoleRepository) WithTrx(trxHandle *gorm.DB) MstRoleRepository {
	if trxHandle == nil {
		mstRole.logger.Zap.Error("transaction Database not found in gin Context")
		return mstRole
	}

	mstRole.db.DB = trxHandle
	return mstRole
}
