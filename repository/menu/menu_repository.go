package menu

import (
	"fmt"
	"riskmanagement/lib"
	menuModels "riskmanagement/models/menu"
	types "riskmanagement/models/type"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MenuDefinition interface {
	GetMenuTree(idMenu string, parentID int) (responses []menuModels.Menu, err error)
	LoadRole(request menuModels.RequestKuisioner) (response menuModels.RoleDesc, err error)
	GetListKuisioner(keyword string, id_menu string, limit int64, offset int64) (responses []menuModels.MenuKuisioner, totalData int64, err error)

	// From MQ
	GetAll() (responses []menuModels.MenuQna, err error)
	SubMenuCheck(request menuModels.MenuQnaRequest) (response menuModels.MenuQna, err error)
	GetAllMstMenu() (responses []menuModels.MstMenu, err error)
	DeleteMenuRRM(request menuModels.MstMenuRequest) (err error)
	StoreMstRRM(request menuModels.MstMenu, tx *gorm.DB) (id string, err error)
	SetStatus(request *menuModels.MstMenu) (responses bool, err error)

	StoreRoleRRM(request *types.Role) (err error)
	DeleteRole(request *types.Role) (err error)

	GetLastID() (id_menu int64, err error)

	GetMstMenuRRM(id string) (response menuModels.MstMenuResponse, err error)
	CheckMenuParentExist(title string) (status bool, err error)
}

type MenuRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewMenuRepository(
	db lib.Database,
	logger logger.Logger,
) MenuDefinition {
	return MenuRepository{
		db:      db,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetMenu implements MenuDefinition.
func (m MenuRepository) GetMenuTree(idMenu string, parentID int) (responses []menuModels.Menu, err error) {
	listId := strings.Split(idMenu, ",")

	db := m.db.DB.Table(`mst_menu_mfe`).
		Select(`
			id,
			title,
			icon,
			path,
			is_section,
			parent_id
		`).
		Where(`parent_id = ?`, parentID).
		Where(`id IN (?)`, listId).
		Order(`ordered ASC`)

	err = db.Scan(&responses).Error

	if err != nil {
		m.logger.Zap.Error(err)
		return nil, err
	}

	return responses, nil
}

func (menu MenuRepository) GetAll() (responses []menuModels.MenuQna, err error) {
	return responses, menu.db.DB.Find(&responses).Error
}

func (menu MenuRepository) SubMenuCheck(request menuModels.MenuQnaRequest) (response menuModels.MenuQna, err error) {
	return response, menu.db.DB.Where("id = ?", request.ID).Where("submenu = ?", "y").Find(&response).Error
}

func (menu MenuRepository) GetAllMstMenu() (responses []menuModels.MstMenu, err error) {
	getMenu := menu.db.DB.Table("mst_menu").
		Select("id_menu, Title, Urutan, child_status, id_parent").
		Where("Url = ''").
		Order("CASE WHEN child_status = 1 THEN id_parent ELSE id_menu END").
		Order("CASE WHEN child_status = 1 THEN urutan ELSE NULL END").
		Order("CASE WHEN child_status = 2 THEN id_parent ELSE id_menu END").
		Order("CASE WHEN child_status = 2 THEN urutan ELSE NULL END").
		Order("urutan").
		Find(&responses)

	return responses, getMenu.Error
}

func (menu MenuRepository) DeleteMenuRRM(request menuModels.MstMenuRequest) (err error) {
	return menu.db.DB.Where("id_menu", request.IDMenu).Delete(&request).Error
}

func (menu MenuRepository) StoreMstRRM(request menuModels.MstMenu, tx *gorm.DB) (id string, err error) {
	return request.IDMenu, tx.Create(&request).Error
}

func (menu MenuRepository) SetStatus(request *menuModels.MstMenu) (responses bool, err error) {
	updateTableSQL := `UPDATE mst_menu_questionnaire SET status = ? WHERE id_menu = ?`

	return true, menu.db.DB.Exec(updateTableSQL, request.Status, request.IDMenu).Error
}

func (menu MenuRepository) StoreRoleRRM(request *types.Role) (err error) {
	return menu.db.DB.Create(&request).Error
}

func (menu MenuRepository) DeleteRole(request *types.Role) (err error) {
	return menu.db.DB.Where("menu_id", request.MenuID).Delete(&request).Error
}

func (menu MenuRepository) GetLastID() (id_menu int64, err error) {
	err = menu.db.DB.Table("mst_menu").Select("MAX(id_menu)").Scan(&id_menu).Error

	if err != nil {
		return id_menu, err
	}

	return id_menu, err
}

func (menu MenuRepository) GetMstMenuRRM(id string) (responses menuModels.MstMenuResponse, err error) {
	db := menu.db.DB

	query := db.Table(`mst_menu`).
		Select(`Title "title", id_parent "id_parent"`).
		Where(`id_menu = ?`, id)

	err = query.Scan(&responses).Error
	if err != nil {
		return responses, err
	}

	return responses, err
}

func (menu MenuRepository) CheckMenuParentExist(title string) (status bool, err error) {
	db := menu.db.DB

	query := db.Table(`mst_menu`).
		Where(`title like ?`, "%"+title+"%")

	var count int64
	query.Count(&count)

	if count > 0 {
		return false, err
	}

	err = query.Error
	if err != nil {
		return status, err
	}

	return true, err
}

func (m MenuRepository) LoadRole(request menuModels.RequestKuisioner) (response menuModels.RoleDesc, err error) {
	db := m.db.DB.Table(`mst_roles_batch3 role`)

	query := db.Select(`
		role.id ,
    	role.role_name 'role_desc'
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
			request.TipeUker, "ALL", "ALL",
			"ALL", request.Hilfm, request.Jgpg,
			request.TipeUker, request.Hilfm, request.Jgpg,
			request.Kostl, request.Hilfm, request.Jgpg,
			request.Orgeh, request.Hilfm, request.Jgpg,
			request.TipeUker, fmt.Sprintf("%%%s%%", request.Pernr),
			request.Kostl, fmt.Sprintf("%%%s%%", request.Pernr),
			request.Orgeh, fmt.Sprintf("%%%s%%", request.Pernr),
		)

	err = query.Find(&response).Error

	return response, err
}

// GetListKuisioner implements MenuDefinition.
func (m MenuRepository) GetListKuisioner(keyword string, id_menu string, limit int64, offset int64) (responses []menuModels.MenuKuisioner, totalData int64, err error) {
	splitMenu := strings.Split(id_menu, ",")

	db := m.db.DB.Table(`mst_menu_questionnaire`)

	query := db.Select(`id_menu 'id',title,icon, url 'path'`).
		Where(`id_menu in (?)`, splitMenu).
		Where("Status = 1").
		Order(`title ASC`)

	if keyword != "" {
		query = query.Where("title LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}

	query.Count(&totalData)

	err = query.Limit(int(limit)).Offset(int(offset)).Find(&responses).Error

	return responses, totalData, err
}
