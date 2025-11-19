package repository

import (
	"riskmanagement/lib"
	models "riskmanagement/models/user"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	db     lib.Database
	logger logger.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db lib.Database, logger logger.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Zap.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.db.DB = trxHandle
	return r
}

// GetAll gets all users
func (r UserRepository) GetAll() (users []models.User, err error) {
	return users, r.db.DB.Find(&users).Error
}

// Save user
func (r UserRepository) Save(user models.User) (models.User, error) {
	return user, r.db.DB.Create(&user).Error
}

// Update updates user
func (r UserRepository) Update(user models.User) (models.User, error) {
	return user, r.db.DB.Save(&user).Error
}

// GetOne gets ont user
func (r UserRepository) GetOne(id uint) (user models.User, err error) {
	return user, r.db.DB.Where("id = ?", id).First(&user).Error
}

// GetOne gets user by email
func (r UserRepository) GetUserByEmail(email *string) (user models.User, err error) {
	return user, r.db.DB.Where("email = ?", email).First(&user).Error
}

// Delete deletes the row of data
func (r UserRepository) Delete(id uint) error {
	return r.db.DB.Where("id = ?", id).Delete(&models.User{}).Error
}

// GetOneAsset implements AssetAccessPlaceDefinition
func (u UserRepository) GetMenu(request models.MenuRequest) (responses models.Menus, err error) {
	rows, err := u.db.DB.Raw(`
	SELECT DISTINCT m.IDMenu, m.Title, m.Url, m.Deskripsi, m.Icon, m.svgIcon, m.fontIcon FROM mst_menu m INNER JOIN mst_access_menu n 
	ON m.IDMenu=n.IDMenu WHERE m.RoleAccess=1 AND m.Status=1 AND m.IDParent = 0 
	AND (
		(n.LevelUker='` + request.LevelUker + `' AND n.LevelID='` + request.LevelID + `') 
		OR (n.LevelUker='ALL' AND n.LevelID='ALL') 
		OR (n.LevelUker='ALL' AND n.LevelID='` + request.LevelID + `') 
		OR (n.LevelUker='` + request.LevelUker + `' AND n.LevelID='ALL')
		OR (n.KOSTL='` + request.Kostl + `' AND n.LevelID='` + request.LevelID + `')
		OR (n.ORGEH='` + request.Orgeh + `' AND n.LevelID='` + request.LevelID + `')
	)`).Rows()

	defer rows.Close()

	var menu models.Menu
	for rows.Next() {
		u.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}
	return responses, err
}

// GetOneAsset implements AssetAccessPlaceDefinition
func (u UserRepository) GetChildMenu(menuID int64, request models.MenuRequest) (responses []models.ChildMenuResponse, err error) {
	rows, err := u.db.DB.Raw(`
	SELECT m.IDMenu, m.Title, m.Url, m.Deskripsi, m.Icon, m.svgIcon, m.fontIcon 
	FROM mst_menu m INNER JOIN mst_access_menu n ON m.IDMenu=n.IDMenu 
	WHERE m.RoleAccess=1 AND m.Status=1 AND
	m.IDParent = ?
	AND (
		(n.LevelUker='`+request.LevelUker+`' AND n.LevelID='`+request.LevelID+`') 
		OR (n.LevelUker='ALL' AND n.LevelID='ALL') OR (n.LevelUker='ALL' AND n.LevelID='`+request.LevelID+`') 
		OR (n.LevelUker='`+request.LevelUker+`' AND n.LevelID='ALL')
		OR (n.KOSTL='`+request.Kostl+`' AND n.LevelID='`+request.LevelID+`')
		OR (n.ORGEH='`+request.Orgeh+`' AND n.LevelID='`+request.LevelID+`')
	)`, menuID).Rows()

	defer rows.Close()

	var menu models.ChildMenuResponse
	for rows.Next() {
		u.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}
	return responses, err
}

func (u UserRepository) GetSubChildMenu(menuID int64, request models.MenuRequest) (responses []models.SubChildMenuResponse, err error) {
	rows, err := u.db.DB.Raw(`
	SELECT m.IDMenu, m.Title, m.Url, m.Deskripsi, m.Icon, m.svgIcon, m.fontIcon 
	FROM mst_menu m INNER JOIN mst_access_menu n ON m.IDMenu=n.IDMenu 
	WHERE m.RoleAccess=1 AND m.Status=1 AND
	m.IDParent = ?
	AND (
		(n.LevelUker='`+request.LevelUker+`' AND n.LevelID='`+request.LevelID+`') 
		OR (n.LevelUker='ALL' AND n.LevelID='ALL') OR (n.LevelUker='ALL' AND n.LevelID='`+request.LevelID+`') 
		OR (n.LevelUker='`+request.LevelUker+`' AND n.LevelID='ALL')
		OR (n.KOSTL='`+request.Kostl+`' AND n.LevelID='`+request.LevelID+`')
		OR (n.ORGEH='`+request.Orgeh+`' AND n.LevelID='`+request.LevelID+`')
	)`, menuID).Rows()

	defer rows.Close()

	var menu models.SubChildMenuResponse
	for rows.Next() {
		u.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}
	return responses, err
}
