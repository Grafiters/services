package taskassignment

import (
	"fmt"
	"regexp"
	"riskmanagement/lib"
	models "riskmanagement/models/taskassignment"
	"strconv"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type TaskAssignmentsDefinition interface {
	StoreData(request models.Task, tx *gorm.DB) (reponses models.Task, err error)
	StoreFile(request models.TaskFile, tx *gorm.DB) (reponses bool, err error)
	CheckTableExist(request models.CheckTableRequest) (response models.TemplateResponse, err error)
	GetDataTask(request models.TaskFilterRequest) (responses []models.TaskResponses, totalData int64, err error)
	GetDataTaskDetail(id int64) (responses models.TaskResponses, err error)
	GetDetailTematik(request models.DataTematikRequest) (responses models.DataTematikResponse, totalData int64, err error)
	GetTaskApprovalList(request models.TaskApprovalRequest) (responses []models.TaskResponses, totalData int64, err error)
	GetRejectionNotes(id int64) (response models.TaskRejectionNotes, err error)

	DeleteLampiran(request models.DataTematikRequest) (response bool, err error)
	DeleteTaskUker(id int64) (response bool, err error)
	GetBranchList(id int64) (response []string, err error)
	MyTasklist(request models.TaskFilterRequest) (responses []models.MyTasklistResponse, totalData int64, err error)
	MyTasklistDetail(id int64) (responses models.MyTasklistResponse, err error)
	// Add by Panji 02-01-2025
	GenerateNoTask(orgeh string) (response string, err error)

	GetAll() (responses []models.Task, err error)
	GetOneById(id int64) (response models.Task, err error)
	StoreTasklist(request models.Task) (response models.Task, err error)
	UpdateTasklist(request models.Task) (response models.Task, err error)
	DeleteTasklist(id int64) (err error)
	GetDWH() (responses []models.DWHBranch, err error)
	ValidateData(data models.ValidateLaporanRAPDTO) (isValid bool)
	CreateTableLampiranRAP(tableName string, columns []string) (err error)
	ValidatFileToDB(tableName string, columns []string) (isExist bool, err error)
	BatchStoreLampiranRAP(tableName string, headers []string, columns [][]string, id int64) (err error)
	BatchStoreTasklistUker(request []models.TasklistUker) error
	StoreTasklistRejected(request models.TasklistRejected) (err error)
	GetTaskType(id int64) (responses models.AdminSetting, err error)
	GetDataBRC(request []string, id int64) (responses []models.GetBRCResponse, err error)
	BatchStoreTasklistToday(request []models.TasklistsToday) (err error)
	BatchStoreNotifTask(request []models.TasklistNotif) (err error)
	DeleteTasklistUker(id int64) (err error)

	InsertLampiranIndicator(request models.LampiranIndicatorRequest) (err error)
	GetMyTasklistTotal(request models.RequestMyTasklist) (Count int64, err error)
}

type TaskAssignmentsRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func ConvertToSnakeCase(str string) string {
	// Replace all non-alphanumeric characters with underscores
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	snake := re.ReplaceAllString(str, "_")

	// Convert to lowercase
	snake = strings.ToLower(snake)

	// Trim any leading or trailing underscores
	snake = strings.Trim(snake, "_")

	return snake
}

func NewTaskAssignmentsRepository(db lib.Database, dbRaw lib.Databases, logger logger.Logger) TaskAssignmentsDefinition {
	return TaskAssignmentsRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

func (r TaskAssignmentsRepository) GetAll() (responses []models.Task, err error) {
	return responses, r.db.DB.Find(&responses).Error
}

func (r TaskAssignmentsRepository) GetOneById(id int64) (responses models.Task, err error) {
	err = r.db.DB.Where("id = ?", id).Find(&responses).Error
	return responses, err
}

func (r TaskAssignmentsRepository) GetDWH() (responses []models.DWHBranch, err error) {
	query := "SELECT REGION, RGDESC, MAINBR, MBDESC, BRANCH, BRDESC FROM dwh_branch"

	r.logger.Zap.Info(query)
	rows, err := r.dbRaw.DB.Query(query)
	if err != nil {
		return responses, err
	}
	defer rows.Close()

	response := models.DWHBranch{}
	for rows.Next() {
		_ = rows.Scan(
			&response.REGION,
			&response.RGDESC,
			&response.MAINBR,
			&response.BRDESC,
			&response.BRANCH,
			&response.BRDESC,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
	// return responses, submajorproses.db.DB.Find(&responses).Error
}

// StoreData implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) StoreData(request models.Task, tx *gorm.DB) (reponses models.Task, err error) {
	return request, tx.Save(&request).Error
}

func (r TaskAssignmentsRepository) StoreTasklist(request models.Task) (response models.Task, err error) {
	if err := r.db.DB.Create(&request).Error; err != nil {
		return models.Task{}, err
	}
	return request, nil
}

// StoreFile implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) StoreFile(request models.TaskFile, tx *gorm.DB) (reponses bool, err error) {
	err = tx.Create(&request).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return false, err
	}

	return true, nil
}

func (r TaskAssignmentsRepository) UpdateTasklist(request models.Task) (response models.Task, err error) {
	if err := r.db.DB.Model(&models.Task{}).Where("id = ?", request.ID).Updates(&request).Error; err != nil {
		return models.Task{}, err
	}

	return request, nil
}

// CheckTableExist implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) CheckTableExist(request models.CheckTableRequest) (response models.TemplateResponse, err error) {
	tableName := "lampiran_rap_" + strconv.FormatInt(request.RiskIssue, 10) + "_" + strconv.FormatInt(request.RiskIndicator, 10)

	var columns []string
	err = r.db.DB.Table("information_schema.columns").
		Select("COLUMN_NAME").
		Where("table_name = ? AND table_schema = DATABASE()", tableName).
		Order("ORDINAL_POSITION").
		Pluck("COLUMN_NAME", &columns).Error

	if err != nil {
		r.logger.Zap.Error("Error CheckTableExist", err)
		return response, err
	}

	columnsList := make([]string, len(columns))
	for i, column := range columns {
		// columnsList[i] = ConvertToSnakeCase(column)
		columnsList[i] = column
	}

	joinsColumns := strings.Join(columnsList, ",")

	if len(columns) < 1 {
		response.Status = ""
		response.Columns = ""
	} else {
		response.Status = "ada"
		response.Columns = joinsColumns
	}

	return response, err
}

func (r TaskAssignmentsRepository) DeleteTasklist(id int64) (err error) {
	return r.db.DB.Where("id = ?", id).Delete(&models.Task{}).Error
}

// GenerateNoTask implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GenerateNoTask(orgeh string) (response string, err error) {
	kode := "TSK-"
	today := lib.GetTimeNow("date2")

	if orgeh != "" {
		kode += "%" + orgeh + "-" + today + "%"
	}

	query := r.db.DB.Table("tasklists").
		Select(`RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'response'`).
		Where("no_tasklist LIKE ?", kode).Order("id desc").Limit(1)

	err = query.Scan(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, err
}

// GetDataTask implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GetDataTask(request models.TaskFilterRequest) (responses []models.TaskResponses, totalData int64, err error) {
	db := r.db.DB.Table("tasklists t")

	query := db.Select(`
			ROW_NUMBER() OVER (ORDER BY t.id DESC) AS 'no',
			t.id,
			t.no_tasklist,
			t.nama_tasklist,
			t.activity_id,
			t.product_id,
			t.product_name,
			t.risk_issue_id,
			t.risk_issue,
			t.risk_indicator_id,
			t.risk_indicator,
			t.task_type,
			t.task_type_name,
			t.start_date,
			t.end_date,
			t.approval_status,
			t.period,
			t.status,
			t.status_file 
		`).
		Joins(`LEFT JOIN tasklists_uker tu ON tu.tasklist_id = t.id`).
		Joins(`JOIN activity a ON a.id = t.activity_id`).
		Where(`t.rap = 0`).
		Group(`t.id`)

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where(`tu.REGION = ?`, request.Kanwil)
	}

	if request.Kanca != "all" && request.Kanca != "" {
		mainbrs := strings.Split(request.Kanca, ",")
		query = query.Where(`tu.MAINBR in (?)`, mainbrs)
	}

	if request.Uker != "all" && request.Uker != "" {
		branches := strings.Split(request.Uker, ",")
		query = query.Where(`tu.BRANCH in (?)`, branches)
	}

	if request.Aktifitas != 0 {
		query = query.Where(`t.activity_id = ?`, request.Aktifitas)
	}

	if request.Produck != 0 {
		query = query.Where(`t.product_id = ?`, request.Produck)
	}

	if request.RiskEvent != 0 {
		query = query.Where(`t.risk_issue_id = ?`, request.RiskEvent)
	}

	if request.RiskIndicator != 0 {
		query = query.Where(`t.risk_indicator_id = ?`, request.RiskIndicator)
	}

	if request.TaskType != 0 {
		query = query.Where(`t.task_type = ?`, request.TaskType)
	}

	if request.StatusApproval != "All" && request.StatusApproval != "" {
		query = query.Where(`t.approval_status = ?`, request.StatusApproval)
	}

	if request.StatusTask != "" {
		query = query.Where(`t.status = ?`, request.StatusTask)
	}

	if request.TglAwal != "" {
		query = query.Where("t.created_at >= ?", request.TglAwal)
	}

	if request.TglAkhir != "" {
		query = query.Where("t.created_at <= ?", request.TglAkhir+" 23:59:59")
	}

	if err = query.Count(&totalData).Error; err != nil {
		r.logger.Zap.Error("Error counting records:", err)
		return
	}

	if request.Limit != 0 {
		query = query.Limit(int(request.Limit))
	}

	if request.Offset != 0 {
		query = query.Offset(int(request.Offset))
	}

	err = query.Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return nil, totalData, err
	}

	return responses, totalData, err
}

// GetTaskApprovalList implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GetTaskApprovalList(request models.TaskApprovalRequest) (responses []models.TaskResponses, totalData int64, err error) {
	db := r.db.DB.Table("tasklists t")

	query := db.Select(`
			ROW_NUMBER() OVER (ORDER BY t.id DESC) AS 'no',
			t.id,
			t.no_tasklist,
			t.nama_tasklist,
			t.activity_id,
			t.product_id,
			t.product_name,
			t.risk_issue_id,
			t.risk_issue,
			t.risk_indicator_id,
			t.risk_indicator,
			t.task_type,
			t.task_type_name,
			t.start_date,
			t.end_date,
			t.approval_status,
			t.period,
			t.status,
			t.status_file,
			t.validation,
			t.approval
		`).
		Joins(`LEFT JOIN tasklists_uker tu ON tu.tasklist_id = t.id`).
		Joins(`JOIN activity a ON a.id = t.activity_id`).
		// Where(`t.rap = 0`).
		Order(`t.id DESC`).
		Group(`t.id`)

	if request.Kanwil != "all" && request.Kanwil != "" {
		query = query.Where(`tu.REGION = ?`, request.Kanwil)
	}

	if request.Kanca != "all" && request.Kanca != "" {
		mainbrs := strings.Split(request.Kanca, ",")
		query = query.Where(`tu.MAINBR in (?)`, mainbrs)
	}

	if request.Uker != "all" && request.Uker != "" {
		branches := strings.Split(request.Uker, ",")
		query = query.Where(`tu.BRANCH in (?)`, branches)
	}

	if request.Aktifitas != 0 {
		query = query.Where(`t.activity_id = ?`, request.Aktifitas)
	}

	if request.Produck != 0 {
		query = query.Where(`t.product_id = ?`, request.Produck)
	}

	if request.RiskEvent != 0 {
		query = query.Where(`t.risk_issue_id = ?`, request.RiskEvent)
	}

	if request.RiskIndicator != 0 {
		query = query.Where(`t.risk_indicator_id = ?`, request.RiskIndicator)
	}

	if request.TaskType != 0 {
		query = query.Where(`t.task_type = ?`, request.TaskType)
	}

	if request.StatusApproval != "All" && request.StatusApproval != "" {
		query = query.Where(`t.approval_status = ?`, request.StatusApproval)
	}

	if request.StatusTask != "" {
		query = query.Where(`t.status = ?`, request.StatusTask)
	}

	if request.Validator != "" {
		query = query.Where(`(t.validation = ? OR t.approval = ?)`, request.Validator, request.Validator)
	}

	if request.TglAwal != "" {
		query = query.Where("t.created_at >= ?", request.TglAwal)
	}

	if request.TglAkhir != "" {
		query = query.Where("t.created_at <= ?", request.TglAkhir+" 23:59:59")
	}

	if err = query.Count(&totalData).Error; err != nil {
		r.logger.Zap.Error("Error counting records:", err)
		return
	}

	if request.Limit != 0 {
		query = query.Limit(int(request.Limit))
	}

	if request.Offset != 0 {
		query = query.Offset(int(request.Offset))
	}

	err = query.Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return nil, totalData, err
	}

	return responses, totalData, err
}

// GetDataTaskDetail implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GetDataTaskDetail(id int64) (responses models.TaskResponses, err error) {
	db := r.db.DB.Table(`tasklists`).Where(`id = ?`, id)

	err = db.Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

func (r TaskAssignmentsRepository) DeleteLampiran(request models.DataTematikRequest) (responses bool, err error) {
	NamaTable := "lampiran_rap_" + strconv.FormatInt(request.RiskEvent, 10) + "_" + strconv.FormatInt(request.RiskIndicator, 10)

	db := r.db.DB.Table(NamaTable)

	err = db.Where("tasklist_id = ?", request.Id).Delete(nil).Error

	if err != nil {
		return false, err
	}

	return true, err
}

// DeleteTaskUker implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) DeleteTaskUker(id int64) (response bool, err error) {
	db := r.db.DB.Table(`tasklists_uker`)

	err = db.Where("tasklist_id = ?", id).Delete(nil).Error

	if err != nil {
		return false, err
	}

	return true, err
}

// GetRejectionNotes implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GetRejectionNotes(id int64) (response models.TaskRejectionNotes, err error) {
	db := r.db.DB.Table(`tasklists_rejected`).Where(`tasklist_id = ?`, id).Order("id DESC").Limit(1)

	err = db.Scan(&response).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return response, err
	}

	return response, nil
}

// GetDetailTematik implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GetDetailTematik(request models.DataTematikRequest) (responses models.DataTematikResponse, totalData int64, err error) {
	NamaTable := "lampiran_rap_" + strconv.FormatInt(request.RiskEvent, 10) + "_" + strconv.FormatInt(request.RiskIndicator, 10)

	var columns []string
	err = r.db.DB.Table("information_schema.columns").
		Select("COLUMN_NAME").
		Where("table_name = ? AND table_schema = DATABASE()", NamaTable).
		Order("ORDINAL_POSITION").
		Pluck("COLUMN_NAME", &columns).Error

	if err != nil {
		fmt.Println("Error fetching column names:", err)
		return responses, totalData, err
	}

	columnList := make([]string, len(columns))
	for i, column := range columns {
		columnList[i] = column
	}

	joinedColumns := strings.Join(columnList, ",")

	var rawData []map[string]interface{}

	selectColumns := "`" + strings.Join(columnList, "`, `") + "`"
	query := r.db.DB.Table(NamaTable).
		Select(selectColumns).
		Where(`tasklist_id = ?`, request.Id)

	if request.Region != "all" && request.Region != "" {
		query = query.Where(`REGION = ?`, request.Region)
	}

	if request.Branch != "" {
		query = query.Where(`BRANCH = ?`, request.Branch)
	}

	query = query.Count(&totalData)

	err = query.Limit(int(request.Limit)).Offset(int(request.Offset)).Find(&rawData).Error

	var columnsData []interface{}
	for _, row := range rawData {
		orderedRow := make(map[string]interface{})

		for _, column := range columns {
			orderedRow[column] = row[column]
		}

		columnsData = append(columnsData, orderedRow)
	}

	responses = models.DataTematikResponse{
		Columns:     joinedColumns,
		ColumnsData: columnsData,
	}

	return responses, totalData, err
}

// GetBranchList implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GetBranchList(id int64) (branches []string, err error) {
	db := r.db.DB.Table("tasklists_uker")

	err = db.Select("BRANCH").Where(`tasklist_id = ?`, id).Find(&branches).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return branches, err
	}

	return branches, err
}

func (r TaskAssignmentsRepository) ValidateData(data models.ValidateLaporanRAPDTO) (isValid bool) {
	var count int64
	r.db.DB.
		Model(&models.DWHBranch{}).
		Where("LOWER(region) = ?", data.REGION).
		// Where("LOWER(rgdesc) = ?", strings.ToLower(data.RGDESC)).
		Where("LOWER(mainbr) = ?", data.MAINBR).
		Where("LOWER(mbdesc) = ?", strings.ToLower(data.MBDESC)).
		Where("branch = ?", data.BRANCH).
		Where("LOWER(brdesc) = ?", strings.ToLower(data.BRDESC)).
		Count(&count)

	if count > 0 {
		return true
	} else {
		return false
	}
}

func (r TaskAssignmentsRepository) ValidatFileToDB(tableName string, columns []string) (isExist bool, err error) {
	for i, col := range columns {
		col = strings.ToLower(strings.ReplaceAll(col, " ", "_"))

		// Check if the column already has a data type
		if !strings.Contains(col, " ") {
			columns[i] = col
		}
	}

	JoinColumns := strings.Join(columns, "', '")
	tblScheme, err := lib.GetVarEnv("DBName")

	fmt.Println("TABLE SCHEME ===>", tblScheme)
	fmt.Println("column ===>", JoinColumns)
	fmt.Println("TABLE NAME ===>", tableName)

	if err != nil {
		return false, err
	}

	query := "SELECT COUNT(*) FROM INFORMATION_SCHEMA.columns WHERE TABLE_SCHEMA = '" + tblScheme + "' AND TABLE_NAME = ? AND COLUMN_NAME IN ('" + JoinColumns + "')"
	var columnCount int
	err = r.dbRaw.DB.QueryRow(query, tableName).Scan(&columnCount)
	if err != nil {
		return false, fmt.Errorf("format kolom file tidak sesuai")
	}

	if columnCount == len(columns) {
		return true, nil
	} else if columnCount == 0 {
		return false, nil
	} else {
		return false, fmt.Errorf("format kolom file tidak sesuai")
	}
}

func (r TaskAssignmentsRepository) CreateTableLampiranRAP(tableName string, columns []string) (err error) {
	headers := make([]string, len(columns))
	copy(headers, columns)
	for i, col := range headers {
		col = strings.ToLower(strings.ReplaceAll(col, " ", "_"))

		// Check if the column already has a data type
		if !strings.Contains(col, " ") {
			headers[i] = col + " VARCHAR(255)"
		}
	}

	// columns = append([]string{"tasklist_id", "status"}, columns...)
	headers = append([]string{"id INT PRIMARY KEY AUTO_INCREMENT", "tasklist_id varchar(255)", "status varchar(255)", "no_verifikasi varchar(255)"}, headers...)

	// Create table
	// columnQuery := strings.Join(columns, " varchar(255), ") + " varchar(255)"
	columnQuery := strings.Join(headers, ", ")

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, columnQuery)
	_, err = r.dbRaw.DB.Exec(query)
	if err != nil {
		return err
	}

	// Create index
	indexQueries := []string{
		fmt.Sprintf("CREATE INDEX idx_REGION ON %s (REGION);", tableName),
		fmt.Sprintf("CREATE INDEX idx_MAINBR ON %s (MAINBR);", tableName),
		fmt.Sprintf("CREATE INDEX idx_BRANCH ON %s (BRANCH);", tableName),
		fmt.Sprintf("CREATE INDEX idx_REGION_MAINBR_BRANCH ON %s (REGION, MAINBR, BRANCH);", tableName),
	}

	for _, query := range indexQueries {
		_, err = r.dbRaw.DB.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r TaskAssignmentsRepository) BatchStoreLampiranRAP(tableName string, headers []string, columns [][]string, id int64) (err error) {
	columnNames := append([]string{}, headers...)

	for i, col := range columnNames {
		col = strings.ToLower(strings.ReplaceAll(col, " ", "_"))

		// Check if the column already has a data type
		if !strings.Contains(col, " ") {
			columnNames[i] = col
		}
	}

	columnNames = append([]string{"tasklist_id", "status"}, columnNames...)

	fmt.Println("COLUMN NAMES ===>", columnNames)

	columnNamesQuery := "(" + strings.Join(columnNames, ", ") + ")"

	fmt.Println("COLUMN NAMES QUERY ===>", columnNamesQuery)

	query := "INSERT INTO " + tableName + " " + columnNamesQuery + " VALUES "
	for _, column := range columns {
		query += "('" + strconv.FormatInt(id, 10) + "','', '" + strings.Join(column, "','") + "'), "
	}
	query = strings.TrimSuffix(query, ", ")

	err = r.db.DB.Exec(query).Error
	if err != nil {
		return fmt.Errorf("format kolom file tidak sesuai")
	}
	return nil
}

func (r TaskAssignmentsRepository) BatchStoreTasklistUker(request []models.TasklistUker) error {
	if err := r.db.DB.Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func (r TaskAssignmentsRepository) StoreTasklistRejected(request models.TasklistRejected) (err error) {
	if err = r.db.DB.Table("tasklists_rejected").Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func (r TaskAssignmentsRepository) GetTaskType(id int64) (responses models.AdminSetting, err error) {
	err = r.db.DB.Table("admin_setting").Where("id = ?", id).First(&responses).Error
	if err != nil {
		return models.AdminSetting{}, err
	}
	return responses, nil
}

func (r TaskAssignmentsRepository) GetDataBRC(request []string, id int64) (responses []models.GetBRCResponse, err error) {
	// err = r.db.DB.Table("uker_kelolaan_user uku").
	// 	Select(`
	// 		uku.pn,
	// 		uku.SNAME,
	// 		uku.REGION,
	// 		uku.RGDESC,
	// 		uku.MAINBR,
	// 		uku.MBDESC,
	// 		uku.BRANCH,
	// 		uku.BRDESC,
	// 		tu.jumlah_nominatif
	// 	`).
	// 	Joins(`JOIN tasklists_uker tu ON tu.BRANCH = uku.BRANCH`).
	// 	Where("uku.BRANCH IN (?)", request).
	// 	Where(`tu.tasklist_id = ?`, id).
	// 	Find(&responses).Error

	err = r.db.DB.Table(`tasklists_uker`).Select(`
		REGION,
		RGDESC,
		MAINBR,
		MBDESC,
		BRANCH,
		BRDESC,
		jumlah_nominatif
	`).Where(`BRANCH IN (?)`, request).Where(`tasklist_id = ?`, id).Find(&responses).Error

	if err != nil {
		return []models.GetBRCResponse{}, err
	}

	return responses, nil
}

func (r TaskAssignmentsRepository) BatchStoreTasklistToday(request []models.TasklistsToday) (err error) {
	if err := r.db.DB.Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func (r TaskAssignmentsRepository) BatchStoreNotifTask(request []models.TasklistNotif) (err error) {
	if err := r.db.DB.Create(&request).Error; err != nil {
		return err
	}
	return nil
}

func (r TaskAssignmentsRepository) DeleteTasklistUker(id int64) (err error) {
	err = r.db.DB.Table("tasklists_uker").Where("tasklist_id = ?", id).Delete(&models.TasklistUker{}).Error
	return err
}

// MyTasklist implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) MyTasklist(request models.TaskFilterRequest) (responses []models.MyTasklistResponse, totalData int64, err error) {
	db := r.db.DB.Table(`tasklists_today tt`)

	query := db.Select(`
			ROW_NUMBER() OVER (ORDER BY tt.id DESC) AS 'no',
			tt.id,
			t.no_tasklist,
			t.nama_tasklist,
        	DATE(t.created_at) 'periode',
			tt.kegiatan, 
			tt.BRANCH,
			tt.BRDESC,
			tt.risk_issue,
			tt.risk_indicator,
			tt.task_type_name,
			tt.period,
			tt.start_date,
			tt.end_date,
			tt.sample 'jumlah_nominatif',
			tt.progres 'progress'
		`).Joins(`JOIN tasklists t ON t.id = tt.tasklist_id`).
		Where(`tt.isReady = 1`).
		Group(`tt.BRANCH, tt.tasklist_id`)

	if request.Branches != "" {
		branches := strings.Split(request.Branches, ",")
		query = query.Where(`tt.BRANCH in (?)`, branches)
	}

	if request.TaskType != 0 {
		query = query.Where(`tt.task_type = ?`, request.TaskType)
	}

	if err = query.Count(&totalData).Error; err != nil {
		r.logger.Zap.Error("Error counting records:", err)
		return
	}

	if request.Limit != 0 {
		query = query.Limit(int(request.Limit))
	}

	if request.Offset != 0 {
		query = query.Offset(int(request.Offset))
	}

	err = query.Scan(&responses).Error

	if err != nil {
		r.logger.Zap.Error(err)
		return nil, totalData, err
	}

	return responses, totalData, err
}

func (r TaskAssignmentsRepository) MyTasklistDetail(id int64) (responses models.MyTasklistResponse, err error) {
	db := r.db.DB.Table(`tasklists_today tt`)

	query := db.Select(`
			tt.tasklist_id 'id',
			tt.task_type_name,
			a.name 'activity_name',
			tt.product 'product_name',
			tt.kegiatan, 
			tt.BRANCH,
			tt.BRDESC,
			tt.risk_issue_id,
			tt.risk_issue,
			tt.risk_indicator_id,
			tt.risk_indicator
		`).
		Joins(`JOIN activity a ON a.id = tt.activity_id `).
		Where(`tt.id = ?`, id)

	err = query.Scan(&responses).Error
	if err != nil {
		return responses, err
	}

	return responses, nil
}

// InsertLampiranIndicator implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) InsertLampiranIndicator(request models.LampiranIndicatorRequest) (err error) {
	if err = r.db.DB.Table("lampiran_indikator").Create(&request).Error; err != nil {
		return err
	}
	return nil
}

// GetMyTasklistTotal implements TaskAssignmentsDefinition.
func (r TaskAssignmentsRepository) GetMyTasklistTotal(request models.RequestMyTasklist) (total int64, err error) {
	type TasklistToday struct {
		TasklistID string
		Branch     string
		IsReady    bool
		// Tambahkan field lain jika diperlukan
	}
	db := r.db.DB.Table(`tasklists_today tt`)

	branches := strings.Split(request.Branch, ",")

	subquery := db.Model(&TasklistToday{}).
		Select("tasklist_id").
		Where(`tt.BRANCH in (?)`, branches).
		Where(`tt.isReady = 1`).
		Group(`tt.BRANCH, tt.tasklist_id`)

	var count int64
	resultDB := r.db.DB
	err = resultDB.Table(`(?) as myTask`, subquery).Count(&count).Error

	if err != nil {
		return count, err
	}

	total = count

	return total, nil
}
