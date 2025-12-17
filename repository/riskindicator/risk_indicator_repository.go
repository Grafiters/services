package riskindicator

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/riskindicator"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskIndicatorDefinition interface {
	WithTrx(trxHandle *gorm.DB) RiskIndicatorRepository
	GetAll() (responses []models.RiskIndicatorResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.RiskIndicatorResponse, totalData int, totalRows int, err error)
	GetOne(id int64) (responses models.RiskIndicatorResponse, err error)
	GetByID(id int64) (responses models.ActivityResponse, err error)
	Store(request *models.RiskIndicator, tx *gorm.DB) (responses *models.RiskIndicator, err error)
	Update(request *models.RiskIndicator, include []string, tx *gorm.DB) (responses bool, err error)
	Delete(request *models.UpdateDelete, include []string, tx *gorm.DB) (response bool, err error)
	SearchRiskIndicatorByIssue(request *models.SearchRequest) (responses []models.RiskIndicatorResponse, totalRows int, totalData int, err error)
	GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateriNull, err error)
	SearchRiskIndicatorBySource(req models.KeyRiskBySourceRequest) (res []models.RiskIndicator, totalData int, err error)
	SearchRiskIndicatorKRID(requests *models.KeyRiskRequest) (responses []models.RiskIndicatorKRIDResponseNull, totalRows int, toatalData int, err error)
	StoreRiskIndicatorKRID(request *models.RiskIndicatorKRID, tx *gorm.DB) (response *models.RiskIndicatorKRID, err error)
	GetKode() (response []models.KodeResponseNull, err error)
	TruncateTable() (response bool, err error)
	FilterRiskIndicator(request *models.FilterRequest) (responses []models.RiskIndicatorResponse, totalData int, totalRows int, err error)
	GetDataThreshold(id int64) (responses []models.ThresholdIndicator, err error)
	GetIndicatorByAktivityProduct(request *models.IndicatorRequest) (responses []models.IndikatorResponse, err error)

	// Batch 3
	SearchRiskIndicatorTematik(request *models.SearchRequest) (responses []models.IndicatorTematikResponse, err error)
	GetTematikData(request *models.TematikDataRequest) (responses []byte, err error)
	// GetTematikData(request *models.TematikDataRequest) (responses models.TematikDataResponse, err error)

	GetMateriIfFinish(request *models.RequestMateriIfFinish) (response []models.RekomendasiMateri, err error)
	BulkCreateRiskIndicator(items []models.RiskIndicator, tx *gorm.DB) error
	UpdateStatus(id int64, status bool) error
}

type RiskIndicatorRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewRiskIndicatorRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) RiskIndicatorDefinition {
	return RiskIndicatorRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) Delete(request *models.UpdateDelete, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(request).Error
}

// GetKode implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) GetKode() (response []models.KodeResponseNull, err error) {
	query := `SELECT 
				RIGHT(CONCAT("0000",(T.risk_indicator_code + 1)), 4) 'kode'
			FROM(
				SELECT
					CAST(SUBSTRING_INDEX(ri.risk_indicator_code,'.', -1) as DECIMAL) 'risk_indicator_code'
				FROM risk_indicator ri
				ORDER BY ri.id DESC LIMIT 1) 
			AS T`

	LI.logger.Zap.Info(query)
	rows, err := LI.dbRaw.DB.Query(query)
	defer rows.Close()

	LI.logger.Zap.Info("rows", rows)
	for err != nil {
		return response, err
	}

	kodeResponse := models.KodeResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&kodeResponse.Kode,
		)

		response = append(response, kodeResponse)
	}

	if err = rows.Err(); err != nil {
		return response, err
	}

	return response, err
}

// GetAll implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) GetAll() (responses []models.RiskIndicatorResponse, err error) {
	return responses, LI.db.DB.Where("delete_flag != 1").Find(&responses).Error
}

// GetAllWithPaginate implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) GetAllWithPaginate(request *models.Paginate) (
	responses []models.RiskIndicatorResponse,
	totalPage int,
	totalRows int,
	err error,
) {

	// ambil whereQuery + args dari fungsi terpisah
	whereQuery, args := buildRiskIndicatorFilterQuery(request)

	orderField := "ri.id"
	if request.Order != "" {
		orderField = "ri." + request.Order
	}

	LI.logger.Zap.Debug(request)
	limit := request.Limit
	offset := request.Offset

	query := fmt.Sprintf(`
		SELECT 
			ri.id,
			ri.risk_indicator_code,
			ri.risk_indicator,
			ri.activity_id,
			ri.product_id,
			ri.deskripsi,
			ri.satuan,
			ri.sifat,
			ri.business_cycle_activity,
			ri.batasan,
			ri.kondisi,
			ri.type,
			ri.sla_verifikasi,
			ri.sla_tindak_lanjut,
			ri.sumber_data,
			ri.sumber_data_text,
			ri.periode_pemantauan,
			ri.owner,
			ri.kpi,
			ri.status_indikator,
			ri.data_source_anomaly,
			ri.status,
			ri.created_at,
			ri.updated_at
		FROM risk_indicator ri
		%s
		ORDER BY %s %s
		LIMIT ? OFFSET ?
	`, whereQuery, orderField, request.Sort)

	argsWithPagination := append(args, limit, offset)

	rows, err := LI.db.DB.Raw(query, argsWithPagination...).Rows()
	if err != nil {
		return nil, 0, 0, err
	}
	defer rows.Close()

	LI.logger.Zap.Debug(rows)

	for rows.Next() {
		var item models.RiskIndicatorResponse
		LI.db.DB.ScanRows(rows, &item)
		responses = append(responses, item)
	}

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM risk_indicator ri
		%s
	`, whereQuery)

	err = LI.dbRaw.DB.QueryRow(countQuery, args...).Scan(&totalRows)
	if err != nil {
		return nil, 0, 0, err
	}

	// hitung total halaman
	result := float64(totalRows) / float64(request.Limit)
	totalPage = int(math.Ceil(result))

	return responses, totalPage, totalRows, nil
}

// GetOne implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) GetOne(id int64) (responses models.RiskIndicatorResponse, err error) {
	err = LI.db.DB.Raw(`SELECT * FROM risk_indicator WHERE id = ?`, id).Find(&responses).Error

	if err != nil {
		LI.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

// Store implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) Store(request *models.RiskIndicator, tx *gorm.DB) (responses *models.RiskIndicator, err error) {
	return request, tx.Save(&request).Error
}

// Update implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) Update(request *models.RiskIndicator, include []string, tx *gorm.DB) (responses bool, err error) {
	err = tx.
		Model(&models.RiskIndicator{}).
		Where("id = ?", request.ID).
		Select(include).
		Updates(request).Error

	return err == nil, err
}

// WithTrx implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) WithTrx(trxHandle *gorm.DB) RiskIndicatorRepository {
	if trxHandle == nil {
		LI.logger.Zap.Error("transaction Database not found in gin context")
		return LI
	}

	LI.db.DB = trxHandle
	return LI
}

// SearchRiskIndicatorByIssue implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) SearchRiskIndicatorByIssue(request *models.SearchRequest) (responses []models.RiskIndicatorResponse, totalRows int, totalData int, err error) {
	db := LI.db.DB.Table(`risk_indicator ri`)

	db = db.Select(`
		ri.id 'id',
		ri.risk_indicator_code 'risk_indicator_code',
		ri.risk_indicator 'risk_indicator'
	`).Joins(`JOIN risk_issue_map_indicator rimi ON rimi.id_indicator = ri.id`).
		Where(`ri.delete_flag = 0`).
		Where(`rimi.id_risk_issue = ?`, request.RiskIssueId)

	if request.Keyword != "" {
		db = db.Where(`CONCAT(ri.risk_indicator_code, ri.risk_indicator) LIKE ?`, fmt.Sprintf("%%%s%%", request.Keyword))
	}

	var count int64
	db.Count(&count)

	totalRows = int(count)
	db.Limit(request.Limit).Offset(request.Offset)

	err = db.Find(&responses).Error
	// calculate the total pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(request.Limit)))
	return responses, totalPages, totalRows, err
}

// GetRekomendasiMateri implements RiskIndicatorDefinition
func (LI RiskIndicatorRepository) GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateriNull, err error) {
	query := `SELECT 
				rimf.id 'id',
				rimf.id_indicator 'id_indicator',
				rimf.nama_lampiran 'nama_lampiran',
				rimf.filename 'filename',
				rimf.path 'path'
			FROM risk_indicator_map_files rimf WHERE rimf.id_indicator  = ?`

	rows, err := LI.dbRaw.DB.Query(query, id)
	defer rows.Close()

	if err != nil {
		return responses, err
	}

	response := models.RekomendasiMateriNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.IDIndicator,
			&response.NamaLampiran,
			&response.Filename,
			&response.Path,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

func (repo RiskIndicatorRepository) SearchRiskIndicatorBySource(req models.KeyRiskBySourceRequest) (res []models.RiskIndicator, totalData int, err error) {
	db := repo.db.DB.Table("risk_indicator")

	db = db.Select(`id, risk_indicator_code, risk_indicator`).Where("data_source_anomaly = ?", req.Source)
	keyword := `%` + req.Keyword + `%`
	if req.Keyword != "" {
		db = db.Where("CONCAT(risk_indicator_code, risk_indicator) LIKE ?", keyword)
	}
	var count int64
	db.Count(&count)
	totalData = int(count)

	db = db.Limit(req.Limit).Offset(req.Offset)
	err = db.Find(&res).Error
	return res, totalData, err
}

// SearchRiskIndicatorKRID implements RiskIndicatorDefinition
func (repo RiskIndicatorRepository) SearchRiskIndicatorKRID(requests *models.KeyRiskRequest) (responses []models.RiskIndicatorKRIDResponseNull, totalRows int, totalData int, err error) {
	where := ""

	keyword := fmt.Sprintf("%%%s%%", requests.Keyword)

	if requests.Keyword != "" {
		where += " WHERE CONCAT(rik.kode_key_risk_indicator, rik.key_risk_indicator) LIKE ?"
	}

	query := `SELECT * FROM risk_indicator_krid rik` + where + ` LIMIT ? OFFSET ?`

	repo.logger.Zap.Info(query)
	rows, err := repo.dbRaw.DB.Query(query, keyword, requests.Limit, requests.Offset)
	defer rows.Close()

	repo.logger.Zap.Info("rows =>", rows)

	if err != nil {
		return responses, totalRows, totalData, err
	}

	response := models.RiskIndicatorKRIDResponseNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.KodeKeyRiskIndicator,
			&response.KeyRiskIndicator,
			&response.Aktifitas,
			&response.Produk,
			&response.JenisIndicator,
			&response.IndikasiRisiko,
		)
		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT count(*) FROM risk_indicator_krid rik` + where

	err = repo.dbRaw.DB.QueryRow(paginationQuery, keyword).Scan(&totalRows)

	result := float64(totalRows) / float64(requests.Limit)
	resultFinal := int(math.Ceil(result))

	return responses, resultFinal, totalData, err
}

// StoreRiskIndicatorKRID implements RiskIndicatorDefinition
func (RiskIndicatorRepository) StoreRiskIndicatorKRID(request *models.RiskIndicatorKRID, tx *gorm.DB) (response *models.RiskIndicatorKRID, err error) {
	return request, tx.Save(&request).Error
}

// TruncateTable implements RiskIndicatorDefinition
func (repo RiskIndicatorRepository) TruncateTable() (response bool, err error) {
	// panic("TRUNCATE TABLE riskmanagement.risk_indicator_krid;")
	err = repo.db.DB.Raw(`TRUNCATE TABLE riskmanagement.risk_indicator_krid`).Find(true).Error
	if err != nil {
		return false, err
	}

	return true, err
}

func (repo RiskIndicatorRepository) GetByID(id int64) (responses models.ActivityResponse, err error) {
	return responses, repo.db.DB.Where("id = ?", id).Find(&responses).Error
}

// FilterRiskIndicator implements RiskIndicatorDefinition
func (ri RiskIndicatorRepository) FilterRiskIndicator(request *models.FilterRequest) (responses []models.RiskIndicatorResponse, totalData int, totalRows int, err error) {
	db := ri.db.DB

	db = db.Table("risk_indicator ri").
		Select(`
			ri.id 'id',
			ri.risk_indicator_code 'risk_indicator_code',
			ri.risk_indicator 'risk_indicator',
			ri.activity_id 'activity_id',
			ri.product_id 'product_id',
			ri.deskripsi 'deskripsi',
			ri.satuan 'satuan',
			ri.sifat 'sifat',
			ri.sla_verifikasi 'sla_verifikasi',
			ri.sla_tindak_lanjut 'sla_tindak_lanjut',
			ri.sumber_data 'sumber_data',
			ri.periode_pemantauan 'periode_pemantauan',
			ri.owner 'owner',
			ri.kpi 'kpi',
			ri.status 'status',
			ri.created_at 'created_at',
			ri.updated_at 'updated_at'
		`).
		Where("delete_flag != 1").
		Where("status = ?", request.Status).Order("ri.created_at ASC")

	if request.Kode != "" {
		db = db.Where("ri.risk_indicator_code like ?", fmt.Sprintf("%%%s%%", request.Kode))
	}

	if request.Indikator != "" {
		db = db.Where("ri.risk_indicator like ?", fmt.Sprintf("%%%s%%", request.Indikator))
	}

	var count int64
	db.Count(&count)

	totalData = int(count)
	fmt.Println("TotalRows =>", totalData)

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error

	totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))

	return responses, totalRows, totalData, err
}

// GetDataThreshold implements RiskIndicatorDefinition
func (kri RiskIndicatorRepository) GetDataThreshold(id int64) (responses []models.ThresholdIndicator, err error) {
	db := kri.db.DB

	db = db.Table(`risk_indicator ri`).
		Select(`
			ri.id 'index',
			ri.risk_indicator_code 'id',
			ri.risk_indicator 'key_risk_indicator',
			a.name 'aktivitas',
			p.product 'produk',
			SUBSTRING(ri.sifat, 1, INSTR(ri.sifat, '/') - 1) 'jenis_indikator',
			"" AS 'indikasi_risiko',
			ri.deskripsi 'deskripsi',
			ri.sla_verifikasi 'sla_verifikasi',
			ri.sla_tindak_lanjut 'sla_tl',
			"" AS 'risk_awareness',
			"" As 'data_source',
			ri.satuan 'parameter',
			ri.status_indikator 'status_indikator',
			CASE 
				WHEN ri.status = 1 THEN 'aktif'
				ELSE 'non aktif '
			END 'is_aktif'`).
		Joins(`LEFT JOIN activity a ON a.id = ri.activity_id`).
		Joins(`LEFT JOIN product p ON p.id = ri.product_id`).
		Where(`ri.id = ?`, id)

	err = db.Find(&responses).Error

	return responses, err
}

// GetIndicatorByAktivityProduct implements RiskIndicatorDefinition
func (ri RiskIndicatorRepository) GetIndicatorByAktivityProduct(request *models.IndicatorRequest) (responses []models.IndikatorResponse, err error) {
	db := ri.db.DB

	db = db.Table(`risk_indicator`).
		Select(`
			id,
			risk_indicator_code,
			risk_indicator
		`).
		Where(`activity_id = ?`, request.Aktivitas).
		Where(`product_id = ?`, request.Produk)

	err = db.Find(&responses).Error

	return responses, err
}

// SearchRiskIndicatorTematik implements RiskIndicatorDefinition.
func (repo RiskIndicatorRepository) SearchRiskIndicatorTematik(request *models.SearchRequest) (responses []models.IndicatorTematikResponse, err error) {
	db := repo.db.DB.Table("lampiran_indikator")

	db = db.Select(`
			risk_indicator_id 'id',
			risk_indicator_desc 'risk_indicator',
			nama_table
		`).Where("risk_issue_id = ?", request.RiskIssueId)

	if request.Keyword != "" {
		keyword := fmt.Sprintf("%%%s%%", request.Keyword)
		db = db.Where("risk_indicator_desc LIKE ?", keyword)
	}

	err = db.Limit(request.Limit).Offset(request.Offset).Scan(&responses).Error

	return responses, err
}

/* split string
// GetTematikData implements RiskIndicatorDefinition.
func (LI RiskIndicatorRepository) GetTematikData(request *models.TematikDataRequest) (responses models.TematikDataResponse, err error) {
	if request.NamaTable == "" {
		return models.TematikDataResponse{}, fmt.Errorf("Table can't be null")
	}

	var columns []string
	queryColumn := `
			SELECT COLUMN_NAME
			FROM INFORMATION_SCHEMA.COLUMNS
			WHERE TABLE_NAME = ? AND TABLE_SCHEMA = DATABASE() AND COLUMN_NAME NOT IN ('id')`

	// Retrieve column names from the specified table
	LI.db.DB.Raw(queryColumn, request.NamaTable).Scan(columns)
	jsonFields := make([]string, len(columns))
	for i, column := range columns {
		jsonFields[i] = fmt.Sprintf("lampiran.`%s`", column) // Null as a string; replace 'null' with NULL for JSON null
	}

	header := strings.Join(columns, "|")

	query := fmt.Sprintf(
		"SELECT CONCAT_WS('|', %s) AS 'stringData' FROM `%s` lampiran"+
			" JOIN tasklists_uker tu ON tu.tasklist_id = lampiran.tasklist_id"+
			" WHERE DATE(periode_data) = '%s'",
		strings.Join(jsonFields, ", "),
		request.NamaTable,
		request.PeriodeData)

	var stringData []string
	row := LI.db.DB.Raw(query).Row() // Get the row containing the JSON result

	err = row.Scan(&stringData)

	if err != nil {
		fmt.Println("Error executing JSON query:", err)
		return responses, err
	}

	responses = models.TematikDataResponse{
		Header: header,
		Data:   stringData, // Assign 'data' to responses.Data
	}

	return responses, err
}
*/

// GetTematikData implements RiskIndicatorDefinition.
func (LI RiskIndicatorRepository) GetTematikData(request *models.TematikDataRequest) (responses []byte, err error) {
	if request.NamaTable != "" {
		fmt.Println("masuk repo")
		var columns []string

		queryColumn := `
			SELECT COLUMN_NAME
			FROM INFORMATION_SCHEMA.COLUMNS
			WHERE TABLE_NAME = ? AND TABLE_SCHEMA = DATABASE() ORDER BY(ORDINAL_POSITION)
			`

		// Retrieve column names from the specified table
		LI.db.DB.Raw(queryColumn, request.NamaTable).Scan(&columns)

		println("Columns =>", columns)

		// Build JSON field strings for each column
		jsonFields := make([]string, len(columns))
		for i, column := range columns {
			jsonFields[i] = fmt.Sprintf("'%s', IFNULL(lampiran.`%s`, 'null')", column, column) // Null as a string; replace 'null' with NULL for JSON null
		}

		// Create the JSON query using JSON_OBJECT and JSON_ARRAYAGG
		jsonQuery := fmt.Sprintf(
			"SELECT JSON_ARRAYAGG(JSON_OBJECT(%s)) AS json_data FROM `%s` lampiran"+
				" JOIN tasklists_uker tu ON tu.tasklist_id = lampiran.tasklist_id"+
				" JOIN tasklists_lampiran tl ON tl.tasklists_id = lampiran .tasklist_id"+
				" WHERE DATE(tl.created_at) = '%s'"+
				" AND tu.BRANCH = '%s'",
			// " WHERE DATE(periode_data) = '%s'",
			strings.Join(jsonFields, ", "),
			request.NamaTable,
			request.PeriodeData,
			request.UnitKerja,
		)

		var jsonData []byte
		row := LI.db.DB.Raw(jsonQuery).Row() // Get the row containing the JSON result

		err = row.Scan(&jsonData)

		if err != nil {
			fmt.Println("Error executing JSON query:", err)
			return nil, err
		}

		return jsonData, nil
	}

	return responses, nil
}

// GetMateriIfFinish implements RiskIndicatorDefinition.
func (LI RiskIndicatorRepository) GetMateriIfFinish(request *models.RequestMateriIfFinish) (response []models.RekomendasiMateri, err error) {
	ids := strings.Split(request.RequestId, ",")
	db := LI.db.DB.Table("risk_indicator_map_files document").
		Select(`
			document.id 'id',
			document.id_indicator 'id_indicator',
			document.nama_lampiran 'nama_lampiran',
			document.filename 'filename',
			document.path 'path'
		`).Where(`document.id in (?)`, ids)

	err = db.Scan(&response).Error

	return response, err
}

func (li RiskIndicatorRepository) UpdateStatus(id int64, status bool) error {
	err := li.db.DB.Table("risk_indicator").Where("id = ?", id).Update("status", status).Error

	return err
}

func (LI RiskIndicatorRepository) BulkCreateRiskIndicator(items []models.RiskIndicator, tx *gorm.DB) error {
	if len(items) == 0 {
		return nil
	}

	batchSize := 100
	if len(items) > 1000 {
		batchSize = 500
	} else if len(items) > 5000 {
		batchSize = 1000
	}

	return tx.CreateInBatches(items, batchSize).Error
}

func buildRiskIndicatorFilterQuery(request *models.Paginate) (string, []interface{}) {
	var (
		whereQuery strings.Builder
		args       []interface{}
	)

	whereQuery.WriteString(" WHERE ri.delete_flag != 1 ")

	// Search global
	if request.Search != "" {
		whereQuery.WriteString(`
			AND (
				ri.risk_indicator_code LIKE ? OR
				ri.risk_indicator LIKE ? OR
				ri.deskripsi LIKE ? OR
				ri.owner LIKE ? OR
				ri.kpi LIKE ?
			)
		`)
		s := "%" + request.Search + "%"
		args = append(args, s, s, s, s, s)
	}

	// Filter created_at
	if request.CreatedAt != "" {
		whereQuery.WriteString(" AND DATE(ri.created_at) = ? ")
		args = append(args, request.CreatedAt)
	}

	if request.Code != "" {
		whereQuery.WriteString(" AND ri.risk_indicator_code LIKE ? ")
		args = append(args, "%"+request.Code+"%")
	}

	// Filter name (risk_indicator)
	if request.Name != "" {
		whereQuery.WriteString(" AND ri.risk_indicator LIKE ? ")
		args = append(args, "%"+request.Name+"%")
	}

	// Filter status_indikator
	if request.Status != "" {
		whereQuery.WriteString(" AND ri.status_indikator = ? ")
		args = append(args, request.Status)
	}

	if request.Batasan != "" {
		whereQuery.WriteString(" AND ri.batasan = ? ")
		args = append(args, request.Batasan)
	}

	// Active / Inactive
	if request.Active {
		whereQuery.WriteString(" AND ri.status = 1 ")
	}
	if request.Inactive {
		whereQuery.WriteString(" AND ri.status = 0 ")
	}

	return whereQuery.String(), args
}
