package admin_setting

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/admin_setting"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type AdminSettingDefinition interface {
	GetAll(request *models.Paginate) (responses []models.TaskType, totalRows int, totalData int, err error)
	Show(request *models.Paginate) (responses []models.AdminSetting, totalRows int, totalData int, err error)
	ShowRole(id int64) (responses []models.TaskTypeRoleResponse, err error)
	Store(request *models.AdminSetting, tx *gorm.DB) (responses *models.AdminSetting, err error)
	StoreRole(request *models.AdminSettingRoleRequest, tx *gorm.DB) (responses *models.AdminSettingRoleRequest, err error)
	Update(request *models.AdminSettingUpdate, tx *gorm.DB) (responses *models.AdminSettingUpdate, err error)
	Delete(request *models.AdminSettingDelete, tx *gorm.DB) (err error)
	DeleteRole(request *models.AdminSettingRole, tx *gorm.DB) (err error)
	SearchTaskType(request *models.KeywordRequest) (responses []models.TaskType, totalRows int, totalData int, err error)
	SearchTaskTypeInput(request *models.KeywordRequest) (responses []models.TaskType, totalRows int, totalData int, err error)
	SearchTaskTypeInputByKegiatan(request *models.KeywordRequest) (responses []models.TaskType, totalRows int, totalData int, err error)
	GetOne(id int64) (responses models.TaskType, err error)
	GetRole(id int64) (responses []models.TaskTypeRoleResponse, err error)
	CheckAvailability(request models.TaskTypeCheckRequest) (response models.TaskTypeCheckResponse, err error)
}

type AdminSettingRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewAdminSettingRepository(db lib.Database, dbRaw lib.Databases, logger logger.Logger) AdminSettingDefinition {
	return AdminSettingRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (setting AdminSettingRepository) GetAll(request *models.Paginate) (responses []models.TaskType, totalRows int, totalData int, err error) {
	whereRole := ``

	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		whereRole = " AND asr.role = 4 "
		// } else if request.StellTX == "MANAGER - RISK MANAGEMENT & COMPLIANCE" || request.StellTX == "SENIOR MANAGER - RISK MANAGEMENT & COMPLIANCE" || request.StellTX == "ASSISTANT MANAGER - RISK MANAGEMENT & COMPLIANCE" || request.StellTX == "OFFICER - RISK MANAGEMENT & COMPLIANCE" {
		// 	whereRole = " AND asr.role = 2 "
	} else if (request.Stell == "70511843" || request.Stell == "70511844" || request.Stell == "70511845") && (request.JG >= "JG03" && request.JG <= "JG06") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511843|70511844|70511845%' OR asr.JG = 'JG03|JG04|JG05, JG06') "
	} else if (request.Stell == "70511846") && (request.JG == "JG07") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511846%' OR asr.JG = 'JG07') "
	} else if (request.Stell == "70511848" || request.Stell == "70511849") && (request.JG == "JG08" || request.JG == "JG09") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511848|70511849%' OR asr.JG = 'JG08|JG09') "
	} else if (request.Stell == "70511850" || request.Stell == "70348729") && (request.JG == "JG10" || request.JG == "JG11") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511850|70348729%' OR asr.JG = 'JG10|JG11') "
		// } else if (request.Orgeh == "70355061" || request.Orgeh == "70355063") && request.HILFM == "158" {
		// 	whereRole = " AND asr.role = 1 "
		// } else if (request.Orgeh == "70355061" || request.Orgeh == "70355063") && request.HILFM == "159" {
		// 	whereRole = " AND asr.role = 1 "
		// } else if (request.Orgeh == "70355061" || request.Orgeh == "70355063") && request.HILFM == "160" {
		// 	whereRole = " AND asr.role = 1 "
		// } else if (request.Orgeh == "70355061" || request.Orgeh == "70355063") && request.HILFM == "162" {
		// 	whereRole = " AND asr.role = 1 "
	} else if request.KOSTL == "PS21014" && request.HILFM == "158" {
		whereRole = " AND asr.role = 1 AND asr.hilfm LIKE '%158%'"
	} else if request.KOSTL == "PS21014" && request.HILFM == "159" {
		whereRole = " AND asr.role = 1 AND asr.hilfm LIKE '%159%'"
	} else if request.KOSTL == "PS21014" && request.HILFM == "160" {
		whereRole = " AND asr.role = 1 AND asr.hilfm LIKE '%160%'"
	} else if request.KOSTL == "PS21014" && request.HILFM == "162" {
		whereRole = " AND asr.role = 1 AND asr.hilfm LIKE '%162%'"
	} else if request.TipeUker == "KC" {
		whereRole = " AND asr.role = 3 "
	}

	query := ""

	if whereRole != "" {
		query = `SELECT DISTINCT admin_setting.id 'id', admin_setting.task_type 'task_type', admin_setting.kegiatan 'kegiatan', admin_setting.period 'period', admin_setting.range 'range', admin_setting.upload 'upload', admin_setting.tasklist_max 'tasklist_max' FROM admin_setting 
			LEFT JOIN admin_setting_role asr ON asr.id_setting = admin_setting.id 
			WHERE admin_setting.status = 'Aktif' ` + whereRole + ``
	}

	setting.logger.Zap.Info(query)
	rows, err := setting.dbRaw.DB.Query(query)

	setting.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.TaskType{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.TaskType,
			&response.Kegiatan,
			&response.Period,
			&response.Range,
			&response.Upload,
			&response.TasklistMax,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginateQuery := fmt.Sprintf(`SELECT count(*) FROM admin_setting WHERE status = "Aktif"`)
	err = setting.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (setting AdminSettingRepository) Show(request *models.Paginate) (responses []models.AdminSetting, totalRows int, totalData int, err error) {
	rows, err := setting.db.DB.Raw(
		`SELECT * FROM admin_setting WHERE status = "Aktif" 
		ORDER BY id ASC LIMIT ? OFFSET ?`, request.Limit, request.Offset).Rows()

	defer rows.Close()

	defer rows.Scan()

	var tasklists models.AdminSetting

	for rows.Next() {
		setting.db.DB.ScanRows(rows, &tasklists)
		responses = append(responses, tasklists)
	}

	paginateQuery := fmt.Sprintf(`SELECT count(*) FROM admin_setting WHERE status = "Aktif"`)
	err = setting.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (setting AdminSettingRepository) ShowRole(id int64) (responses []models.TaskTypeRoleResponse, err error) {
	// err = setting.db.DB.Raw(`
	// 	SELECT mst_roles.role_name FROM admin_setting_role JOIN mst_roles ON mst_roles.id = admin_setting_role.id_role WHERE admin_setting_role.id_setting = ?`, id).Find(&responses).Error
	err = setting.db.DB.Raw(`
		SELECT id_setting, orgeh, kostl, hilfm, tipe_uker, role, stell_tx, stell, jg FROM admin_setting_role WHERE admin_setting_role.id_setting = ?`, id).Find(&responses).Error

	if err != nil {
		setting.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (setting AdminSettingRepository) Store(request *models.AdminSetting, tx *gorm.DB) (responses *models.AdminSetting, err error) {
	return request, tx.Create(&request).Error
}

func (setting AdminSettingRepository) StoreRole(request *models.AdminSettingRoleRequest, tx *gorm.DB) (responses *models.AdminSettingRoleRequest, err error) {
	return request, tx.Create(&request).Error
}

func (setting AdminSettingRepository) Update(request *models.AdminSettingUpdate, tx *gorm.DB) (responses *models.AdminSettingUpdate, err error) {
	return request, tx.Save(&request).Error
}

func (setting AdminSettingRepository) Delete(request *models.AdminSettingDelete, tx *gorm.DB) (err error) {
	return tx.Save(&request).Error
	// return tx.Delete(&request).Error
}

func (setting AdminSettingRepository) DeleteRole(request *models.AdminSettingRole, tx *gorm.DB) (err error) {
	return tx.Where("id_setting", request.IDSetting).Delete(&request).Error
}

func (setting AdminSettingRepository) SearchTaskType(request *models.KeywordRequest) (responses []models.TaskType, totalRows int, totalData int, err error) {
	keyword := ``
	kegiatan := ``

	if request.Keyword != "" {
		keyword += `and admin_setting.task_type LIKE '%` + request.Keyword + `%'`
	}

	if request.Kegiatan != "" {
		kegiatan += `and admin_setting.kegiatan = '` + request.Kegiatan + `'`
	}

	query := `SELECT admin_setting.id 'id', admin_setting.task_type 'task_type', admin_setting.kegiatan 'kegiatan', admin_setting.period 'period', admin_setting.range 'range', admin_setting.upload 'upload', admin_setting.tasklist_max 'tasklist_max' FROM admin_setting 
		WHERE admin_setting.status = 'Aktif' ` + keyword + ` ` + kegiatan + ` LIMIT ? OFFSET ?`

	setting.logger.Zap.Info(query)
	rows, err := setting.dbRaw.DB.Query(query, request.Limit, request.Offset)

	setting.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.TaskType{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.TaskType,
			&response.Kegiatan,
			&response.Period,
			&response.Range,
			&response.Upload,
			&response.TasklistMax,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT COUNT(*) FROM admin_setting WHERE status = 'Aktif' ` + keyword
	err = setting.dbRaw.DB.QueryRow(paginationQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (setting AdminSettingRepository) SearchTaskTypeInput(request *models.KeywordRequest) (responses []models.TaskType, totalRows int, totalData int, err error) {
	keyword := ``

	if request.Keyword != "" {
		keyword += `and admin_setting.task_type LIKE '%` + request.Keyword + `%'`
	}

	whereRole := ``

	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		whereRole = " AND asr.role = 4 "
		// } else if request.StellTX == "MANAGER - RISK MANAGEMENT & COMPLIANCE" || request.StellTX == "SENIOR MANAGER - RISK MANAGEMENT & COMPLIANCE" || request.StellTX == "ASSISTANT MANAGER - RISK MANAGEMENT & COMPLIANCE" || request.StellTX == "OFFICER - RISK MANAGEMENT & COMPLIANCE" {
		// 	whereRole = " AND asr.role = 2 "
	} else if (request.Stell == "70511843" || request.Stell == "70511844" || request.Stell == "70511845") && (request.JG >= "JG03" && request.JG <= "JG06") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511843|70511844|70511845%' OR asr.JG = 'JG03|JG04|JG05, JG06') "
	} else if (request.Stell == "70511846") && (request.JG == "JG07") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511846%' OR asr.JG = 'JG07') "
	} else if (request.Stell == "70511848" || request.Stell == "70511849") && (request.JG == "JG08" || request.JG == "JG09") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511848|70511849%' OR asr.JG = 'JG08|JG09') "
	} else if (request.Stell == "70511850" || request.Stell == "70348729") && (request.JG == "JG10" || request.JG == "JG11") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511850|70348729%' OR asr.JG = 'JG10|JG11') "
	} else if request.KOSTL == "PS21014" && request.HILFM == "158" {
		whereRole = " AND asr.role = 1 "
	} else if request.KOSTL == "PS21014" && request.HILFM == "159" {
		whereRole = " AND asr.role = 1 "
	} else if request.KOSTL == "PS21014" && request.HILFM == "160" {
		whereRole = " AND asr.role = 1 "
	} else if request.KOSTL == "PS21014" && request.HILFM == "162" {
		whereRole = " AND asr.role = 1 "
	} else if request.TipeUker == "KC" {
		whereRole = " AND asr.role = 3 "
	}

	query := ""

	if whereRole != "" {
		query = `SELECT DISTINCT admin_setting.id 'id', admin_setting.task_type 'task_type', admin_setting.kegiatan 'kegiatan', admin_setting.period 'period', admin_setting.range 'range', admin_setting.upload 'upload', admin_setting.tasklist_max 'tasklist_max' FROM admin_setting 
			LEFT JOIN admin_setting_role asr ON asr.id_setting = admin_setting.id 
			WHERE admin_setting.status = 'Aktif' ` + keyword + ` ` + whereRole + ` LIMIT ? OFFSET ?`
	}

	setting.logger.Zap.Info(query)
	rows, err := setting.dbRaw.DB.Query(query, request.Limit, request.Offset)

	setting.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.TaskType{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.TaskType,
			&response.Kegiatan,
			&response.Period,
			&response.Range,
			&response.Upload,
			&response.TasklistMax,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT COUNT(*) FROM admin_setting WHERE status = 'Aktif' ` + keyword
	err = setting.dbRaw.DB.QueryRow(paginationQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (setting AdminSettingRepository) SearchTaskTypeInputByKegiatan(request *models.KeywordRequest) (responses []models.TaskType, totalRows int, totalData int, err error) {
	keyword := ``

	if request.Keyword != "" {
		keyword += `and admin_setting.task_type LIKE '%` + request.Keyword + `%'`
	}

	whereRole := ``

	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		whereRole = " AND asr.role = 4 "
	} else if (request.Stell == "70511843" || request.Stell == "70511844" || request.Stell == "70511845") && (request.JG >= "JG03" && request.JG <= "JG06") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511843|70511844|70511845%' OR asr.JG = 'JG03|JG04|JG05, JG06') "
	} else if (request.Stell == "70511846") && (request.JG == "JG07") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511846%' OR asr.JG = 'JG07') "
	} else if (request.Stell == "70511848" || request.Stell == "70511849") && (request.JG == "JG08" || request.JG == "JG09") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511848|70511849%' OR asr.JG = 'JG08|JG09') "
	} else if (request.Stell == "70511850" || request.Stell == "70348729") && (request.JG == "JG10" || request.JG == "JG11") {
		whereRole = " AND asr.role = 2 AND (asr.stell LIKE '%70511850|70348729%' OR asr.JG = 'JG10|JG11') "
	} else if request.KOSTL == "PS21014" && request.HILFM == "158" {
		whereRole = " AND asr.role = 1 "
	} else if request.KOSTL == "PS21014" && request.HILFM == "159" {
		whereRole = " AND asr.role = 1 "
	} else if request.KOSTL == "PS21014" && request.HILFM == "160" {
		whereRole = " AND asr.role = 1 "
	} else if request.KOSTL == "PS21014" && request.HILFM == "162" {
		whereRole = " AND asr.role = 1 "
	} else if request.TipeUker == "KC" {
		whereRole = " AND asr.role = 3 "
	}

	query := ""

	kegiatan := " AND admin_setting.kegiatan = '" + request.Kegiatan + "'"

	if whereRole != "" {
		query = `SELECT DISTINCT admin_setting.id 'id', admin_setting.task_type 'task_type', admin_setting.kegiatan 'kegiatan', admin_setting.period 'period', admin_setting.range 'range', admin_setting.upload 'upload', admin_setting.tasklist_max 'tasklist_max' FROM admin_setting 
			LEFT JOIN admin_setting_role asr ON asr.id_setting = admin_setting.id 
			WHERE admin_setting.status = 'Aktif'` + kegiatan + ` ` + keyword + ` ` + whereRole + ` LIMIT ? OFFSET ?`
	}

	setting.logger.Zap.Info(query)
	rows, err := setting.dbRaw.DB.Query(query, request.Limit, request.Offset)

	setting.logger.Zap.Info("rows ", rows)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.TaskType{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.TaskType,
			&response.Kegiatan,
			&response.Period,
			&response.Range,
			&response.Upload,
			&response.TasklistMax,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT COUNT(*) FROM admin_setting WHERE status = 'Aktif' ` + kegiatan + ` ` + keyword
	err = setting.dbRaw.DB.QueryRow(paginationQuery).Scan(&totalRows)

	result := float64(totalRows) / float64(request.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalRows, err
}

func (setting AdminSettingRepository) GetOne(id int64) (responses models.TaskType, err error) {
	err = setting.db.DB.Raw(`
		SELECT admin_setting.id 'id', admin_setting.task_type 'task_type', admin_setting.kegiatan 'kegiatan', 
		admin_setting.period 'period', admin_setting.range 'range', admin_setting.upload 'upload', 
		admin_setting.tasklist_max 'tasklist_max' FROM admin_setting WHERE id = ?`, id).Find(&responses).Error

	if err != nil {
		setting.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (setting AdminSettingRepository) GetRole(id int64) (responses []models.TaskTypeRoleResponse, err error) {
	err = setting.db.DB.Raw(`
		SELECT admin_setting_role.id 'id', admin_setting_role.id_setting 'id_setting', admin_setting_role.kostl 'kostl', admin_setting_role.orgeh 'orgeh', admin_setting_role.hilfm 'hilfm', admin_setting_role.stell 'stell_tx', admin_setting_role.jg 'jg', admin_setting_role.tipe_uker 'tipe_uker', admin_setting_role.role 'role' FROM admin_setting_role WHERE id_setting = ?`, id).Find(&responses).Error

	if err != nil {
		setting.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (setting AdminSettingRepository) CheckAvailability(request models.TaskTypeCheckRequest) (response models.TaskTypeCheckResponse, err error) {
	err = setting.db.DB.Raw(`
		SELECT COUNT(*) as 'total' FROM admin_setting WHERE task_type = ? AND status = 'Aktif'`, request.TaskType).Find(&response).Error

	if err != nil {
		setting.logger.Zap.Error(err)
		return response, err
	}
	return response, err
}
