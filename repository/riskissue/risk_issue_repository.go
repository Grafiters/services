package riskissue

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type RiskIssueDefinition interface {
	GetAll() (responses []models.RiskIssueResponse, err error)
	GetAllWithPaginate(request *models.Paginate) (responses []models.RiskIssueResponse, totalData int, totalRows int, err error)
	GetOne(id int64) (responses models.RiskIssueResponse, err error)
	Store(request *models.RiskIssue, tx *gorm.DB) (responses *models.RiskIssue, err error)
	Update(request *models.RiskIssueUpdate, include []string, tx *gorm.DB) (responses bool, err error)
	DeleteMapProses(id int64, tx *gorm.DB) (err error)
	DeleteMapEvent(id int64, tx *gorm.DB) (err error)
	DeleteMapProduct(id int64, tx *gorm.DB) (err error)
	DeleteMapKejadian(id int64, tx *gorm.DB) (err error)
	DeleteMapLiniBisnis(id int64, tx *gorm.DB) (err error)
	DeleteMapAktifitas(id int64, tx *gorm.DB) (err error)
	DeleteMapControl(id int64, tx *gorm.DB) (err error)
	DeleteMapIndicator(id int64, tx *gorm.DB) (err error)
	Delete(request *models.RiskIssueDeleteRequest, include []string, tx *gorm.DB) (response bool, err error)
	GetKode() (responses []models.KodeResponsNull, err error)
	SearchRiskIssue(request *models.KeywordRequest) (responses []models.RiskIssueResponses, totalRows int, totalData int, err error)
	SearchRiskIssueWithoutSub(request *models.RiskIssueWithoutSub) (responses []models.RiskIssueResponses, totalRows int, totalData int, err error)
	FilterRiskIssue(request *models.FilterRiskIssueRequest) (responses []models.RiskIssueFilterResponses, totalRows int, totalData int, err error)
	WithTrx(trxHandle *gorm.DB) RiskIssueRepository
	GetRiskIssueByActivity(id int64) (responses []models.RiskIssueResponseByActivityNull, err error)
	GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateriNull, err error)
	GetMateriByCode(request models.RiskIssueCode) (responses []models.ListMateriNull, err error)
	GetRiskIssueByActivityID(id int64) (responses []models.RiskIssueResponseByActivityNull, err error)
	GetRiskEventName(id int64) (responses models.RiskIssueName, err error)

	GenerateNewCode() (string, error)
	UpdateStatus(id int64, status bool) error
	BulkCreateRiskEvent(items []models.RiskIssue, tx *gorm.DB) error
}

type RiskIssueRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewRiskIssueRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) RiskIssueDefinition {
	return RiskIssueRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// WithTrx implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) WithTrx(trxHandle *gorm.DB) RiskIssueRepository {
	if trxHandle == nil {
		riskIssue.logger.Zap.Error("transaction Database not found in gin context")
		return riskIssue
	}

	riskIssue.db.DB = trxHandle
	return riskIssue
}

// GetAllWithPaginate implements RiskIssueDefinition
func (RI RiskIssueRepository) GetAllWithPaginate(request *models.Paginate) (responses []models.RiskIssueResponse, totalData int, totalRows int, err error) {
	baseQuery := `
		SELECT 
			ri.id,
			ri.risk_type_id,
			ri.risk_issue_code,
			ri.risk_issue,
			ri.deskripsi,
			ri.kategori_risiko,
			ri.status,
			ri.likelihood,
			ri.impact,
			ri.delete_flag,
			ri.created_at,
			ri.updated_at
		FROM risk_issue ri
		WHERE ri.delete_flag != 1
	`

	countQuery := `
		SELECT count(*)
		FROM risk_issue ri
		WHERE ri.delete_flag != 1
	`

	var whereParams []interface{}
	var whereClause string

	// Dynamic Filtering
	if request.Search != "" {
		whereClause += " AND (ri.risk_issue LIKE ? OR ri.risk_issue_code LIKE ? OR ri.deskripsi LIKE ?) "
		searchVal := "%" + request.Search + "%"
		whereParams = append(whereParams, searchVal, searchVal, searchVal)
	}

	if request.RiskTypeID != 0 {
		whereClause += " AND ri.risk_type_id = ? "
		whereParams = append(whereParams, request.RiskTypeID)
	}

	if request.RiskIssueCode != "" {
		whereClause += " AND ri.risk_issue_code LIKE ? "
		whereParams = append(whereParams, "%"+request.RiskIssueCode+"%")
	}

	if request.RiskIssue != "" {
		whereClause += " AND ri.risk_issue LIKE ? "
		whereParams = append(whereParams, "%"+request.RiskIssue+"%")
	}

	if len(request.Status) > 0 {
		var mappedStatus []int

		for _, s := range request.Status {
			switch strings.ToLower(s) {
			case "active", "aktif", "true", "1":
				mappedStatus = append(mappedStatus, 1)
			case "inactive", "nonaktif", "false", "0":
				mappedStatus = append(mappedStatus, 0)
			}
		}

		if len(mappedStatus) > 0 {
			whereClause += " AND ri.status IN (?) "
			whereParams = append(whereParams, mappedStatus)
		}
	}

	if request.KategoriRisiko != "" {
		whereClause += " AND ri.kategori_risiko = ? "
		whereParams = append(whereParams, request.KategoriRisiko)
	}

	if request.CreatedAt != nil {
		whereClause += " AND DATE(ri.created_at) = DATE(?) "
		whereParams = append(whereParams, *request.CreatedAt)
	}

	if request.UpdatedAt != nil {
		whereClause += " AND DATE(ri.updated_at) = DATE(?) "
		whereParams = append(whereParams, *request.UpdatedAt)
	}

	// Append WHERE clause
	finalQuery := baseQuery + whereClause
	finalCountQuery := countQuery + whereClause

	// Sorting
	orderBy := "ri.id"
	if request.Order != "" {
		orderBy = request.Order
	}
	sortDirection := "ASC"
	if strings.ToUpper(request.Sort) == "DESC" {
		sortDirection = "DESC"
	}

	finalQuery += " ORDER BY " + orderBy + " " + sortDirection
	finalQuery += " LIMIT ? OFFSET ?"

	whereParams = append(whereParams, request.Limit, request.Offset)

	// Execute main query
	rows, err := RI.db.DB.Raw(finalQuery, whereParams...).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	for rows.Next() {
		var issue models.RiskIssueResponse
		RI.db.DB.ScanRows(rows, &issue)
		responses = append(responses, issue)
	}

	// Count query
	err = RI.dbRaw.DB.QueryRow(finalCountQuery, whereParams[:len(whereParams)-2]...).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	return responses, totalRows, totalData, nil
}

// GetAll implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetAll() (responses []models.RiskIssueResponse, err error) {
	return responses, riskIssue.db.DB.Raw(`SELECT * FROM risk_issue WHERE delete_flag != '1'`).Find(&responses).Error
}

// GetOne implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetOne(id int64) (responses models.RiskIssueResponse, err error) {
	err = riskIssue.db.DB.Raw(`SELECT * FROM risk_issue WHERE id = ?`, id).Find(&responses).Error

	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}
func (riskIssue RiskIssueRepository) Store(request *models.RiskIssue, tx *gorm.DB) (responses *models.RiskIssue, err error) {

	var RiskIssueCode string

	// === 1. Jika request code kosong → langsung generate kode baru ===
	if request.RiskIssueCode == "" {
		RiskIssueCode, err = riskIssue.GenerateNewCode()
		if err != nil {
			return nil, err
		}

	} else {

		// === 2. Jika tidak kosong, cek apakah sudah ada di database ===
		var count int64
		err = riskIssue.db.DB.
			Table("risk_issue").
			Where("risk_issue_code = ?", request.RiskIssueCode).
			Count(&count).Error
		if err != nil {
			return nil, err
		}

		if count == 0 {
			// Tidak ada → aman dipakai
			RiskIssueCode = request.RiskIssueCode
		} else {
			// Ada → harus generate baru
			RiskIssueCode, err = riskIssue.GenerateNewCode()
			if err != nil {
				return nil, err
			}
		}
	}

	// === 3. Persist data ===
	input := &models.RiskIssue{
		RiskTypeID:     request.RiskTypeID,
		RiskIssueCode:  RiskIssueCode,
		RiskIssue:      request.RiskIssue,
		Deskripsi:      request.Deskripsi,
		KategoriRisiko: request.KategoriRisiko,
		Status:         request.Status,
		Likelihood:     request.Likelihood,
		Impact:         request.Impact,
		DeleteFlag:     request.DeleteFlag,
		CreatedAt:      request.CreatedAt,
	}

	err = tx.Create(input).Error // lebih tepat dari Save untuk insert
	return input, err
}

func (riskIssue RiskIssueRepository) GenerateNewCode() (string, error) {

	// Base: RE.20250127
	baseCode := "RE." + lib.GetTimeNow("date2")

	var count int64
	err := riskIssue.db.DB.
		Table("risk_issue").
		Where("risk_issue_code LIKE ?", baseCode+"%").
		Count(&count).Error

	if err != nil {
		return "", err
	}

	// Generate 4 digit sequence
	seq := fmt.Sprintf("%04d", count+1)

	return baseCode + "." + seq, nil
}

// Update implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) Update(request *models.RiskIssueUpdate, include []string, tx *gorm.DB) (responses bool, err error) {
	return true, tx.Save(&request).Error
}

// DeleteMapAktifitas implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) DeleteMapAktifitas(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapAktifitas{}).Error
}

// DeleteMapEvent implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) DeleteMapEvent(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapEvent{}).Error
}

// DeleteMapKejadian implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) DeleteMapKejadian(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapKejadian{}).Error
}

// DeleteMapLiniBisnis implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) DeleteMapLiniBisnis(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapLiniBisnis{}).Error
}

// DeleteMapProduct implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) DeleteMapProduct(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapProduct{}).Error
}

// DeleteMapProses implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) DeleteMapProses(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapProses{}).Error
}

// DeleteMapControl implements RiskIssueDefinition
func (RiskIssueRepository) DeleteMapControl(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapControl{}).Error
}

// DeleteMapIndicator implements RiskIssueDefinition
func (RiskIssueRepository) DeleteMapIndicator(id int64, tx *gorm.DB) (err error) {
	return tx.Where("id = ?", id).Delete(&models.MapIndicator{}).Error
}

// GetKode implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetKode() (responses []models.KodeResponsNull, err error) {
	rowsCheckCode, err := riskIssue.dbRaw.DB.Query(`SELECT COUNT(*) 'count' FROM risk_issue`)

	query := ""

	checkErr(err)

	if checkCount(rowsCheckCode) == 0 {
		query = `SELECT RIGHT(CONCAT("0000",(count(*) + 1)), 4) 'kode' FROM risk_issue`
	} else {
		query = `SELECT 
					RIGHT(CONCAT("0000",(T.risk_issue_code + 1)), 4) 'kode'
				FROM(
					SELECT
						CAST(SUBSTRING_INDEX(ri.risk_issue_code,'.', -1) as DECIMAL) 'risk_issue_code'
					FROM risk_issue ri 
					ORDER BY ri .id DESC LIMIT 1) 
				AS T`
	}

	riskIssue.logger.Zap.Info(query)
	rows, err := riskIssue.dbRaw.DB.Query(query)

	riskIssue.logger.Zap.Info("rows", rows)
	for err != nil {
		return responses, err
	}

	response := models.KodeResponsNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.Kode,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// SearchRiskIssue implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) SearchRiskIssue(request *models.KeywordRequest) (responses []models.RiskIssueResponses, totalRows int, totalData int, err error) {
	db := riskIssue.db.DB.Table("risk_issue_map_aktifitas rima")

	db = db.Select(`
			ri.id  'ID',
			ri.risk_type_id 'risk_type_id',
			ri.risk_issue_code  'risk_issue_code',
			ri.risk_issue 'risk_issue'
		`).
		Joins(`INNER JOIN risk_issue ri ON rima.id_risk_issue = ri.id`).
		Where(`ri.delete_flag = 0`).
		Where(`rima.aktifitas = ?`, request.Aktivitas).
		Group(`ri.id`)
		// Where(`rima.sub_aktifitas = ?`, request.SubActivityID)

	if request.Keyword != "" {
		db = db.Where(`CONCAT(ri.risk_issue_code,' - ', ri.risk_issue) LIKE ?`, fmt.Sprintf("%%%s%%", request.Keyword))
	}

	var count int64
	db.Count(&count)

	totalData = int(count)
	db.Limit(request.Limit).Offset(request.Offset)

	err = db.Find(&responses).Error
	// calculate the total pages
	totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	return responses, totalRows, totalData, err
}

// Delete implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) Delete(request *models.RiskIssueDeleteRequest, include []string, tx *gorm.DB) (response bool, err error) {
	return true, tx.Save(&request).Error
}

// FilterRiskIssue implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) FilterRiskIssue(request *models.FilterRiskIssueRequest) (responses []models.RiskIssueFilterResponses, totalRows int, totalData int, err error) {
	db := riskIssue.db.DB

	queryBuilder := db.Model(&responses).
		Select(`
			DISTINCT 
			risk_issue.id 'id',
			risk_issue.risk_type_id 'risk_type_id',
			risk_issue.risk_issue_code 'risk_issue_code',
			risk_issue.risk_issue 'risk_issue',
			risk_issue.kategori_risiko 'kategori_risiko',
			risk_issue.status 'status'
		`).
		Joins("JOIN risk_issue_map_product mprod ON mprod.id_risk_issue = risk_issue.id").
		Joins("JOIN risk_issue_map_lini_bisnis mlb ON mlb.id_risk_issue = risk_issue.id").
		Joins("JOIN risk_issue_map_aktifitas makt ON makt.id_risk_issue = risk_issue.id").
		Joins("JOIN risk_issue_map_proses mpros ON mpros.id_risk_issue = risk_issue.id").
		Joins("JOIN risk_issue_map_event mevent ON mevent.id_risk_issue = risk_issue.id").
		Joins("JOIN risk_issue_map_kejadian mkejadian ON mkejadian.id_risk_issue = risk_issue.id").
		Where("risk_issue.delete_flag != 1").
		Order("risk_issue.id DESC")

	if request.Kode != "" {
		queryBuilder = queryBuilder.Where("risk_issue.risk_issue_code like ?", fmt.Sprintf("%%%s%%", request.Kode))
	}

	if request.RiskIssue != "" {
		queryBuilder = queryBuilder.Where("risk_issue.risk_issue like ?", fmt.Sprintf("%%%s%%", request.RiskIssue))
	}

	if request.RiskTypeID != "" {
		queryBuilder = queryBuilder.Where("risk_issue.risk_type_id = ?", request.RiskTypeID)
	}

	if request.KategoriRisiko != "" {
		queryBuilder = queryBuilder.Where("risk_issue.kategori_risiko = ?", request.KategoriRisiko)
	}

	if !request.Status {
		queryBuilder = queryBuilder.Where("risk_issue.status = ?", false)
	}

	if request.Status {
		queryBuilder = queryBuilder.Where("risk_issue.status = ?", true)
	}

	if request.Product != "" {
		queryBuilder = queryBuilder.Where("mprod.product = ?", request.Product)
	}

	if request.LiniBisnis != "" {
		queryBuilder = queryBuilder.Where("mlb.lini_bisnis_lv1 = ?", request.LiniBisnis)
	}

	if request.Aktifitas != "" {
		queryBuilder = queryBuilder.Where("makt.aktifitas = ?", request.Aktifitas)
	}

	if request.Proses != "" {
		queryBuilder = queryBuilder.Where("mpros.mega_proses = ?", request.Proses)
	}

	if request.EventType != "" {
		queryBuilder = queryBuilder.Where("mevent.event_type_lv1 = ?", request.EventType)
	}

	if request.Kejadian != "" {
		queryBuilder = queryBuilder.Where("mkejadian.penyebab_kejadian_lv1 = ?", request.Kejadian)
	}

	var count int64
	queryBuilder.
		Group("risk_issue.id").
		Count(&count)

	totalData = int(count)

	queryBuilder.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses)

	// calculate the total pages
	totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	return responses, totalRows, totalData, err
}

// GetRekomendasiMateri implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateriNull, err error) {
	query := `SELECT 
				mapFiles.id,
				mapFiles.id_indicator,
				mapFiles.nama_lampiran,
				mapFiles.path,
				mapFiles.filename
			FROM risk_issue issue
			INNER JOIN risk_issue_map_indicator mapIndicator ON mapIndicator.id_risk_issue = issue.id
			INNER JOIN risk_indicator_map_files mapFiles ON mapFiles.id_indicator = mapIndicator.id_indicator
			WHERE issue.id = ?`

	// riskIssue.logger.Zap.Info(query)
	rows, err := riskIssue.dbRaw.DB.Query(query, id)
	if err != nil {
		return responses, err
	}
	defer rows.Close()

	// riskIssue.logger.Zap.Info("rows =>", rows)

	if err != nil {
		return responses, err
	}

	response := models.RekomendasiMateriNull{}
	for rows.Next() {
		err = rows.Scan(
			&response.ID,
			&response.IDIndicator,
			&response.NamaLampiran,
			&response.Path,
			&response.Filename,
		)

		if err != nil {
			return responses, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetRiskIssueByActivity implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetRiskIssueByActivity(id int64) (responses []models.RiskIssueResponseByActivityNull, err error) {
	query := `SELECT 
				ri.id,
				ri.risk_issue_code,
				ri.risk_issue
			FROM sub_activity subs 
			INNER JOIN risk_issue_map_aktifitas mapAct ON mapAct.sub_aktifitas = subs.id
			INNER JOIN risk_issue ri ON ri.id = mapAct.id_risk_issue
			WHERE subs.id = ?`

	// riskIssue.logger.Zap.Info(query)
	rows, err := riskIssue.dbRaw.DB.Query(query, id)
	if err != nil {
		return responses, err
	}

	defer rows.Close()

	// riskIssue.logger.Zap.Info("rows =>", rows)

	if err != nil {
		return responses, err
	}

	response := models.RiskIssueResponseByActivityNull{}
	for rows.Next() {
		err = rows.Scan(
			&response.ID,
			&response.RiskIssueCode,
			&response.RiskIssue,
		)

		if err != nil {
			return responses, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetMateriByCode implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetMateriByCode(request models.RiskIssueCode) (responses []models.ListMateriNull, err error) {
	query := `SELECT
				rimf.* 
			FROM risk_issue ri
			INNER JOIN risk_issue_map_indicator rimi ON rimi.id_risk_issue = ri.id 
			INNER JOIN risk_indicator_map_files rimf ON rimf.id_indicator = rimi.id_indicator 
			WHERE ri.risk_issue_code = ?`

	riskIssue.logger.Zap.Info(query)
	rows, err := riskIssue.dbRaw.DB.Query(query, request.RiskIssueCode)
	if err != nil {
		return responses, err
	}

	defer rows.Close()

	riskIssue.logger.Zap.Info("rows =>", rows)

	if err != nil {
		return responses, err
	}

	response := models.ListMateriNull{}
	for rows.Next() {
		err = rows.Scan(
			&response.ID,
			&response.IDIndicator,
			&response.NamaLampiran,
			&response.NomorLampiran,
			&response.JenisFile,
			&response.Path,
			&response.Filename,
		)

		if err != nil {
			return responses, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// GetRiskIssueByActivityID implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) GetRiskIssueByActivityID(id int64) (responses []models.RiskIssueResponseByActivityNull, err error) {
	query := `SELECT 
					DISTINCT 
					ri.id,
					ri.risk_issue_code,
					ri.risk_issue
				FROM risk_issue ri
				JOIN risk_issue_map_aktifitas rima ON rima.id_risk_issue = ri.id 
				WHERE rima.aktifitas = ?`

	riskIssue.logger.Zap.Info(query)
	rows, err := riskIssue.dbRaw.DB.Query(query, id)

	// riskIssue.logger.Zap.Info("rows =>", rows)

	if err != nil {
		return responses, err
	}

	response := models.RiskIssueResponseByActivityNull{}
	for rows.Next() {
		_ = rows.Scan(
			&response.ID,
			&response.RiskIssueCode,
			&response.RiskIssue,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

func (riskIssue RiskIssueRepository) GetRiskEventName(id int64) (responses models.RiskIssueName, err error) {
	db := riskIssue.db.DB

	err = db.Table("risk_issue").
		Select("risk_issue").
		Where("id = ?", id).
		First(&responses).Error

	if err != nil {
		return responses, err
	}

	return responses, err
}

// SearchRiskIssueWithoutSub implements RiskIssueDefinition.
func (riskIssue RiskIssueRepository) SearchRiskIssueWithoutSub(request *models.RiskIssueWithoutSub) (responses []models.RiskIssueResponses, totalRows int, totalData int, err error) {
	searchProd := riskIssue.db.DB.Table(`risk_issue ri`)

	searchProd.Select(`
		DISTINCT
		ri.id 'ID',
		ri.risk_type_id 'risk_type_id',
		ri.risk_issue_code  'risk_issue_code',
		ri.risk_issue 'risk_issue'
	`).Joins(`left join risk_issue_map_product rimp on rimp.id_risk_issue = ri.id`).Where(`rimp.product = ?`, request.ProductID)

	if request.Keyword != "" {
		searchProd.Where(`ri.risk_issue LIKE ?`, fmt.Sprintf("%%%s%%", request.Keyword))
	}

	var countProd int64
	searchProd.Count(&countProd)

	totalRowsProd := int(countProd)

	if totalRowsProd == 0 {
		searchAct := riskIssue.db.DB.Table(`risk_issue ri`)

		searchAct.Select(`
			DISTINCT
			ri.id 'ID',
			ri.risk_type_id 'risk_type_id',
			ri.risk_issue_code  'risk_issue_code',
			ri.risk_issue 'risk_issue'
		`).Joins(`left join risk_issue_map_aktifitas rima on rima.id_risk_issue = ri.id`).Where(`rima.aktifitas = ?`, request.ActivityID)

		if request.Keyword != "" {
			searchAct.Where(`ri.risk_issue LIKE ?`, fmt.Sprintf("%%%s%%", request.Keyword))
		}

		var countAct int64
		searchAct.Count(&countAct)
		totalRowsAct := int(countAct)

		searchAct.Limit(request.Limit).Offset(request.Offset)

		errAct := searchAct.Find(&responses).Error

		// calculate the total pages
		totalPagesAct := int(math.Ceil(float64(totalRowsAct) / float64(request.Limit)))

		return responses, totalPagesAct, totalRowsAct, errAct
	}

	searchProd.Limit(request.Limit).Offset(request.Offset)

	errProd := searchProd.Find(&responses).Error

	// calculate the total pages
	totalPagesProd := int(math.Ceil(float64(totalRowsProd) / float64(request.Limit)))

	return responses, totalPagesProd, totalRowsProd, errProd
}

// Update data risk control for status only
func (rc RiskIssueRepository) UpdateStatus(id int64, status bool) error {
	err := rc.db.DB.Table("risk_issue").Where("id = ?", id).Update("status", status).Error

	return err
}

func (LI RiskIssueRepository) BulkCreateRiskEvent(items []models.RiskIssue, tx *gorm.DB) error {
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
