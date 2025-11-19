package mstrole

import (
	"riskmanagement/lib"
	models "riskmanagement/models/mstrole"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MstRoleMenuDefinition interface {
	WithTrx(trxHandle *gorm.DB) MstRoleMenuRepository
	Store(request *models.MstRoleMapMenu, tx *gorm.DB) (responses *models.MstRoleMapMenu, err error)
	GetByIDRole(id int64) (responses []models.MstRoleMapMenuResponseOne, err error)
	DeleteByID(id int64, tx *gorm.DB) (err error)
}

type MstRoleMenuRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMstRoleMenuRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MstRoleMenuDefinition {
	return MstRoleMenuRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Store implements MstRoleMenuDefinition
func (MstRoleMenuRepository) Store(request *models.MstRoleMapMenu, tx *gorm.DB) (responses *models.MstRoleMapMenu, err error) {
	return request, tx.Save(&request).Error
}

// DeleteByID implements MstRoleMenuDefinition
func (mstRoleMenu MstRoleMenuRepository) DeleteByID(id int64, tx *gorm.DB) (err error) {
	panic("unimplemented")
}

// GetByIDRole implements MstRoleMenuDefinition
func (mstRoleMenu MstRoleMenuRepository) GetByIDRole(id int64) (responses []models.MstRoleMapMenuResponseOne, err error) {
	rows, err := mstRoleMenu.db.DB.Raw(`SELECT * FROM mst_role_map_menu WHERE id_role = ?`, id).Rows()
	// rows, err := mstRoleMenu.db.DB.Raw(`SELECT DISTINCT
	// 										mapmenu.id 'id',
	// 										mapmenu.id_role 'id_role',
	// 										mapmenu.id_menu 'id_menu',
	// 										menu.Title 'title'
	// 									FROM
	// 									mst_role_map_menu mapmenu
	// 									INNER JOIN mst_menu menu ON menu.id_menu = mapmenu.id_menu WHERE mapmenu.id_role = ?`, id).Rows()

	defer rows.Close()
	var menu models.MstRoleMapMenuResponseOne

	for rows.Next() {
		mstRoleMenu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

// WithTrx implements MstRoleMenuDefinition
func (mstRoleMenu MstRoleMenuRepository) WithTrx(trxHandle *gorm.DB) MstRoleMenuRepository {
	if trxHandle == nil {
		mstRoleMenu.logger.Zap.Error("transaction Database not found in gin context")
		return mstRoleMenu
	}

	mstRoleMenu.db.DB = trxHandle
	return mstRoleMenu
}
