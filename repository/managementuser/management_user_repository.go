package managementuser

import (
	"database/sql"
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/managementuser"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type ManagementUserDefinition interface {
	GetAll() (responses []models.ManagementUserResponse, err error)
	GetOne(id int64) (responses models.ManagementUserResponse, err error)
	Store(request *models.ManagementUserRequest) (responses bool, err error)
	Update(request *models.ManagementUserRequest) (responese bool, err error)
	Delete(id int64) (err error)
	GetAllMenu() (responses []models.Menu, err error)
	WithTrx(trxHandle *gorm.DB) ManagementUserRepository
	DeleteMappingMenu(id int64, tx *gorm.DB) (err error)
	GetMenu(request string) (responses models.Menus, err error)
	GetMenuQuestionnaire(request string) (responses models.Menus, err error)

	GetChildMenu(menuID string, request string) (responses []models.ChildMenuResponse, err error)
	GetChildMenuQuest(menuID string, request string) (responses []models.ChildMenuResponse, err error)

	GetSubChildMenu(menuID string, request string) (responses []models.SubChildMenuResponse, err error)
	GetSubChildMenuQuest(menuID string, request string) (responses []models.SubChildMenuResponse, err error)

	GetAllWithPaginate(request *models.Paginate) (responses []models.ManagementUserFinResponse, totalRows int, totalData int, err error)
	CheckUserAccessibility(request *models.ManagementUserCheckAccessibilityRequest) (isExist bool, err error)
	GetUkerKelolaan(request *models.UkerKelolaanUserRequest) (responses []models.UkerKelolaanUserResponseNull, err error)
	GetTreeMenu() (responses models.Menus, err error)
	GetChildTreeMenu(menuID string) (responses []models.ChildMenuResponse, err error)
	GetSubChildTreeMenu(menuID string) (responses []models.SubChildMenuResponse, err error)
	GetLevelUker() (responses []models.LevelUkerResponse, err error)
	GetJabatanRole() (responses []models.JabatanRolesResponse, err error)
	GetAdditionalMenu() (response []models.AdditionalMenuResponse, err error)
	GetAdditionalMenuById(id string) (response []models.AdditionalMenuResponse, err error)
}

type ManagementUserRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewManagementUserRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) ManagementUserDefinition {
	return ManagementUserRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Delete implements ManagementUserDefinition
func (managementuser ManagementUserRepository) Delete(id int64) (err error) {
	return managementuser.db.DB.Where("id = ?", id).Delete(&models.ManagementUserResponse{}).Error
}

// GetAllWithPaginate implements ManagementUserDefinition
func (mu ManagementUserRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.ManagementUserFinResponse, totalRows int, totalData int, err error) {
	query := mu.db.DB.Table(`management_user_batch3 mu`).
		Select(`
			mu.id 'id',
			mr.role_name 'role',
			lu.deskripsi 'level_uker',
			mj.description 'stell_tx',
			mj.jgpg 'jgpg' 
		`).
		Joins(`LEFT JOIN mst_roles_batch3 mr ON mr.id = mu.role_id`).
		Joins(`LEFT JOIN level_uker lu ON lu.level_uker = mu.level_uker`).
		Joins(`LEFT JOIN mst_jabatan mj ON mj.id = mu.level_id`).
		Order("mu.id")

	if request.Role != 0 {
		query = query.Where("mu.role_id = ?", request.Role)
	}

	var count int64
	query.Count(&count)

	totalRows = int(count)

	query.Limit(request.Limit).Offset(request.Offset).Find(&responses)

	// calculate the total pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(request.Limit)))
	return responses, totalPages, totalRows, err
}

// GetAll implements ManagementUserDefinition
func (managementuser ManagementUserRepository) GetAll() (responses []models.ManagementUserResponse, err error) {
	return responses, managementuser.db.DB.Find(&responses).Error
}

// GetOne implements ManagementUserDefinition
func (managementuser ManagementUserRepository) GetOne(id int64) (responses models.ManagementUserResponse, err error) {
	return responses, managementuser.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements ManagementUserDefinition
func (managementuser ManagementUserRepository) Store(request *models.ManagementUserRequest) (responses bool, err error) {
	// rowsCheckCode, err := managementuser.dbRaw.DB.Query(`
	// SELECT
	// 	COUNT(*) 'count'
	// FROM management_user_batch3
	// WHERE
	// 	role_id = ? AND
	// 	level_uker = ? AND
	// 	level_id = ? AND
	// 	ORGEH = ?`, request.NamaJabatan, request.LevelUker, request.LevelID, request.ORGEH)
	rowsCheckCode, err := managementuser.dbRaw.DB.Query(`
	SELECT 
		COUNT(*) 'count'
	FROM management_user_batch3 
	WHERE 
		role_id = ? AND
		level_uker = ? AND 
		level_id = ?`, request.RoleID, request.LevelUker, request.LevelID)

	defer rowsCheckCode.Close()

	checkErr(err)

	if checkCount(rowsCheckCode) == 0 {
		timeNow := lib.GetTimeNow("timestime")
		err = managementuser.db.DB.Save(&models.ManagementUserRequest{
			RoleID:     request.RoleID,
			LevelUker:  request.LevelUker,
			LevelID:    request.LevelID,
			AddonPernr: request.AddonPernr,
			// ORGEH:       request.ORGEH,
			CreatedAt: &timeNow,
		}).Error

		return true, err
	}

	return false, err
}

// Update implements ManagementUserDefinition
func (managementuser ManagementUserRepository) Update(request *models.ManagementUserRequest) (responese bool, err error) {
	// rowsCheckCode, err := managementuser.dbRaw.DB.Query(`
	// SELECT
	// 	COUNT(*) 'count'
	// FROM management_user_batch3
	// WHERE
	// 	nama_jabatan = ? AND
	// 	level_uker = ? AND
	// 	level_id = ? AND
	// 	ORGEH = ?`, request.NamaJabatan, request.LevelUker, request.LevelID, request.ORGEH)

	// checkErr(err)

	// if checkCount(rowsCheckCode) == 0 {
	timeNow := lib.GetTimeNow("timestime")
	err = managementuser.db.DB.Save(&models.ManagementUserRequest{
		ID: request.ID,
		// NamaJabatan: request.NamaJabatan,
		RoleID:     request.RoleID,
		LevelUker:  request.LevelUker,
		LevelID:    request.LevelID,
		AddonPernr: request.AddonPernr,
		// ORGEH:       request.ORGEH,
		CreatedAt: request.CreatedAt,
		UpdatedAt: &timeNow,
	}).Error

	return true, err
	// }

	// return false, err
}

// WithTrx implements ManagementUserDefinition
func (managementuser ManagementUserRepository) WithTrx(trxHandle *gorm.DB) ManagementUserRepository {
	if trxHandle == nil {
		managementuser.logger.Zap.Error("transaction Database not found in gin context")
		return managementuser
	}

	managementuser.db.DB = trxHandle
	return managementuser
}

// GetAllMenu implements ManagementUserDefinition
func (managementuser ManagementUserRepository) GetAllMenu() (responses []models.Menu, err error) {
	return responses, managementuser.db.DB.Find(&responses).Error
	// db := managementuser.db.DB

	// query := db.Model(&responses).
	// 	Select(`
	// 		mst_menu.id_menu,
	// 		mst_menu.Title,
	// 		mst_menu.Url,
	// 		mst_menu.Deskripsi,
	// 		mst_menu.icon,
	// 		mst_menu.svg_icon,
	// 		mst_menu.font_icon,
	// 		mst_menu.id_parent,
	// 		mst_menu.child_status
	// 	`).
	// 	Order("mst_menu.id_parent").
	// 	Order("mst_menu.id_menu")

	// query.Find(&responses)

	// return responses, err
}

// DeleteMappingMenu implements ManagementUserDefinition
func (managementuser ManagementUserRepository) DeleteMappingMenu(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapMenu{}).Error
}

// GetMenu implements ManagementUserDefinition
func (mu ManagementUserRepository) GetMenu(request string) (responses models.Menus, err error) {
	splitMenu := strings.Split(request, ",")

	query := `SELECT 
				DISTINCT 
				m.id_menu, 
				m.Title, 
				m.Url, 
				m.Deskripsi, 
				m.icon, 
				m.svg_icon, 
				m.font_icon,
				m.Urutan
			FROM mst_menu m 
			WHERE m.role_access=1 AND m.Status=1 AND m.id_parent = 0
			AND m.id_menu IN (?)
			ORDER BY m.Urutan ASC`

	rows, err := mu.db.DB.Raw(query, splitMenu).Rows()

	defer rows.Close()

	var menu models.Menu
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

func (mu ManagementUserRepository) GetMenuQuestionnaire(request string) (responses models.Menus, err error) {
	splitMenu := strings.Split(request, ",")

	query := `SELECT 
				DISTINCT 
				m.id_menu, 
				m.Title, 
				m.Url, 
				m.Deskripsi, 
				m.icon, 
				m.svg_icon, 
				m.font_icon,
				m.Urutan
			FROM mst_menu_questionnaire m 
			WHERE m.role_access=1 AND m.Status=1 AND m.id_parent = 0
			AND m.id_menu IN (?)
			GROUP BY m.Url
			ORDER BY m.Urutan ASC`

	rows, err := mu.db.DB.Raw(query, splitMenu).Rows()

	defer rows.Close()

	var menu models.Menu
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

// GetChildMenu implements ManagementUserDefinition
func (mu ManagementUserRepository) GetChildMenu(menuID string, request string) (responses []models.ChildMenuResponse, err error) {
	splitMenu := strings.Split(request, ",")

	query := `SELECT 
					DISTINCT 
					m.id_menu, 
					m.Title, 
					m.Url, 
					m.Deskripsi, 
					m.icon, 
					m.svg_icon, 
					m.font_icon,
					m.Urutan
				FROM mst_menu m 
				WHERE m.role_access=1 AND m.Status=1 AND m.Status=1 AND
				m.id_parent = ? 
				AND m.id_menu IN (?)
				ORDER BY m.Urutan ASC`

	rows, err := mu.db.DB.Raw(query, menuID, splitMenu).Rows()

	defer rows.Close()

	var menu models.ChildMenuResponse
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

func (mu ManagementUserRepository) GetChildMenuQuest(menuID string, request string) (responses []models.ChildMenuResponse, err error) {
	splitMenu := strings.Split(request, ",")

	query := `SELECT 
					DISTINCT 
					m.id_menu, 
					m.Title, 
					m.Url, 
					m.Deskripsi, 
					m.icon, 
					m.svg_icon, 
					m.font_icon,
					m.Urutan
				FROM mst_menu_questionnaire m  
				WHERE m.role_access=1 AND m.Status=1 AND m.child_status=1 AND
				m.id_parent = ?
				AND m.id_menu IN (?)
				GROUP BY m.Url
				ORDER BY m.Urutan ASC`

	rows, err := mu.db.DB.Raw(query, menuID, splitMenu).Rows()

	defer rows.Close()

	var menu models.ChildMenuResponse
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

// GetSubChildMenu implements ManagementUserDefinition
func (mu ManagementUserRepository) GetSubChildMenu(menuID string, request string) (responses []models.SubChildMenuResponse, err error) {
	splitMenu := strings.Split(request, ",")
	query := `SELECT 
					DISTINCT 
					m.id_menu, 
					m.Title, 
					m.Url, 
					m.Deskripsi, 
					m.icon, 
					m.svg_icon, 
					m.font_icon,
					m.Urutan
				FROM mst_menu m 
				WHERE m.role_access=1 AND m.Status=1 AND m.Status=1 AND
				m.id_parent = ?
				AND m.id_menu IN ?
				ORDER BY m.Urutan ASC
				`
	rows, err := mu.db.DB.Raw(query, menuID, splitMenu).Rows()

	defer rows.Close()

	var menu models.SubChildMenuResponse
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

func (mu ManagementUserRepository) GetSubChildMenuQuest(menuID string, request string) (responses []models.SubChildMenuResponse, err error) {
	splitMenu := strings.Split(request, ",")

	query := `SELECT 
					DISTINCT 
					m.id_menu, 
					m.Title, 
					m.Url, 
					m.Deskripsi, 
					m.icon, 
					m.svg_icon, 
					m.font_icon,
					m.Urutan
				FROM mst_menu_questionnaire m 
				WHERE m.role_access=1 AND m.Status=1  AND m.child_status=2 AND
				m.id_parent = ?
				AND m.id_menu IN ?
				GROUP BY m.Url
				ORDER BY m.Urutan ASC
				`
	rows, err := mu.db.DB.Raw(query, menuID, splitMenu).Rows()

	defer rows.Close()

	var menu models.SubChildMenuResponse
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

func (mu ManagementUserRepository) CheckUserAccessibility(request *models.ManagementUserCheckAccessibilityRequest) (isExist bool, err error) {
	totalData := 0
	query := ""

	firstString := request.PERNR[0:1]

	if firstString == "0" {
		query = `SELECT 
					COUNT(id)
				FROM 
					management_user_batch3 
				WHERE
					level_uker = '` + request.LevelUker + `'
				AND level_id = '` + request.LevelID + `'
				AND ORGEH = '` + request.ORGEH + `'
			  `
	} else {
		query = `SELECT 
					COUNT(mu.id)
				FROM management_user_batch3 mu
				JOIN pgs_user pu  ON mu.nama_jabatan = pu.jabatan_pgs 
				WHERE pu.pn = '` + request.PERNR + `'
				`
	}

	//QUERY
	mu.logger.Zap.Info("management_user_batch3-query-activity-unknown", query)
	err = mu.dbRaw.DB.QueryRow(query).Scan(&totalData)

	fmt.Println("totalData", totalData)

	isExist = totalData > 0

	if err != nil {
		return isExist, err
	}

	return isExist, err
}

func (mu ManagementUserRepository) GetUkerKelolaan(request *models.UkerKelolaanUserRequest) (responses []models.UkerKelolaanUserResponseNull, err error) {
	timeNow := lib.GetTimeNow("timestime")

	query := `SELECT 
				  *
			  FROM 
			  	  uker_kelolaan_user uku
			  WHERE
				  uku.pn = ?
			  AND (ISNULL(uku.expired_at) OR uku.expired_at >= ?)
			  `
	//QUERY
	mu.logger.Zap.Info("uker_kelolaan_user-query-activity-unknown", query)
	rows, err := mu.dbRaw.DB.Query(query, request.PERNR, timeNow)
	defer rows.Close()

	mu.logger.Zap.Info("uker_kelolaan_user-rows-activity-unknown", rows)
	if err != nil {
		return responses, err
	}

	response := models.UkerKelolaanUserResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.Id,
			&response.CreatedAt,
			&response.UpdatedAt,
			&response.ExpiredAt,
			&response.IsTemp,
			&response.PN,
			&response.REGION,
			&response.RGDESC,
			&response.MAINBR,
			&response.MBDESC,
			&response.BRANCH,
			&response.BRDESC,
		)
		responses = append(responses, response)
	}

	fmt.Println("responses", responses)

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetTreeMenu implements ManagementUserDefinition
func (mu ManagementUserRepository) GetTreeMenu() (responses models.Menus, err error) {
	rows, err := mu.db.DB.Raw(`
		SELECT 
			DISTINCT 
			m.id_menu, 
			m.Title, 
			m.Url, 
			m.Deskripsi, 
			m.icon, 
			m.svg_icon, 
			m.font_icon,
			m.Urutan
		FROM mst_menu m 
		WHERE m.role_access=1 AND m.Status=1 AND m.id_parent = 0
		ORDER BY m.Urutan ASC
	`).Rows()

	// INNER JOIN mst_role_map_menu mapMenu ON mapMenu.id_menu = m.id_menu
	// INNER JOIN mst_roles_batch3 role ON mapMenu.id_role = role.id
	// INNER JOIN management_user_batch3 mu ON role.role_name = mu.nama_jabatan

	defer rows.Close()

	var menu models.Menu
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

// GetChildTreeMenu implements ManagementUserDefinition
func (mu ManagementUserRepository) GetChildTreeMenu(menuID string) (responses []models.ChildMenuResponse, err error) {
	rows, err := mu.db.DB.Raw(`
		SELECT 
			DISTINCT 
			m.id_menu, 
			m.Title, 
			m.Url, 
			m.Deskripsi, 
			m.icon, 
			m.svg_icon, 
			m.font_icon,
			m.Urutan
		FROM mst_menu m 
		WHERE m.role_access=1 AND m.Status=1 AND m.id_parent = ?
		ORDER BY m.urutan ASC
	`, menuID).Rows()

	// INNER JOIN mst_role_map_menu mapMenu ON mapMenu.id_menu = m.id_menu
	// INNER JOIN mst_roles_batch3 role ON mapMenu.id_role = role.id
	// INNER JOIN management_user_batch3 mu ON role.role_name = mu.nama_jabatan

	defer rows.Close()

	var menu models.ChildMenuResponse
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

// GetSubChildTreeMenu implements ManagementUserDefinition
func (mu ManagementUserRepository) GetSubChildTreeMenu(menuID string) (responses []models.SubChildMenuResponse, err error) {
	rows, err := mu.db.DB.Raw(`
		SELECT 
			DISTINCT 
			m.id_menu, 
			m.Title, 
			m.Url, 
			m.Deskripsi, 
			m.icon, 
			m.svg_icon, 
			m.font_icon,
			m.Urutan
		FROM mst_menu m 
		WHERE m.role_access=1 AND m.Status=1 AND m.id_parent = ?
		ORDER BY m.urutan ASC
	`, menuID).Rows()

	// INNER JOIN mst_role_map_menu mapMenu ON mapMenu.id_menu = m.id_menu
	// INNER JOIN mst_roles_batch3 role ON mapMenu.id_role = role.id
	// INNER JOIN management_user_batch3 mu ON role.role_name = mu.nama_jabatan

	defer rows.Close()

	var menu models.SubChildMenuResponse
	for rows.Next() {
		mu.db.DB.ScanRows(rows, &menu)
		responses = append(responses, menu)
	}

	return responses, err
}

// GetLevelUker implements ManagementUserDefinition
func (mu ManagementUserRepository) GetLevelUker() (responses []models.LevelUkerResponse, err error) {
	return responses, mu.db.DB.Find(&responses).Error
}

// Enhance Management User By Panji 02/02/2024
// GetJabatanRole implements ManagementUserDefinition.
func (mu ManagementUserRepository) GetJabatanRole() (responses []models.JabatanRolesResponse, err error) {
	return responses, mu.db.DB.Find(&responses).Error
}

// GetAdditionalMenu implements ManagementUserDefinition.
func (mu ManagementUserRepository) GetAdditionalMenu() (response []models.AdditionalMenuResponse, err error) {
	db := mu.db.DB.Table("mst_additional_menu").
		Select("id, nama, url, icon")

	err = db.Scan(&response).Error

	return response, err
}

// GetAdditionalMenuById implements ManagementUserDefinition.
func (mu ManagementUserRepository) GetAdditionalMenuById(id string) (response []models.AdditionalMenuResponse, err error) {
	db := mu.db.DB.Table("mst_additional_menu").
		Select("id, nama, url, icon")

	ids := strings.Split(id, ",")

	db = db.Where("id in (?)", ids)

	err = db.Scan(&response).Error

	return response, err
}
