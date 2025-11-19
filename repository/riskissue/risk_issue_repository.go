package riskissue

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
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
	rows, err := RI.db.DB.Raw(`
	SELECT	
		ri.id 'id',
		ri.risk_type_id 'risk_type_id',
		ri.risk_issue_code 'risk_issue_code',
		ri.risk_issue 'risk_issue',
		ri.deskripsi 'deskripsi',
		ri.kategori_risiko 'kategori_risiko',
		ri.status 'status',
		ri.likelihood 'likelihood',
		ri.impact 'impact',
		ri.delete_flag 'deleted_flag',
		ri.created_at 'created_at',
		ri.updated_at 'updated_at'
	FROM risk_issue ri
	WHERE ri.delete_flag != 1 ORDER BY ri.id ASC LIMIT ? OFFSET ?
	`, request.Limit, request.Offset).Rows()
	if err != nil {
		return responses, totalRows, totalData, err
	}

	defer rows.Close()
	var issue models.RiskIssueResponse

	for rows.Next() {
		RI.db.DB.ScanRows(rows, &issue)
		responses = append(responses, issue)
	}

	paginateQuery := `SELECT 
						count(*)
					FROM risk_issue WHERE delete_flag != 1`
	err = RI.dbRaw.DB.QueryRow(paginateQuery).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	return responses, totalRows, totalData, err
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

// Store implements RiskIssueDefinition
func (riskIssue RiskIssueRepository) Store(request *models.RiskIssue, tx *gorm.DB) (responses *models.RiskIssue, err error) {
	// return request, tx.Save(&request).Error
	checkRiskCode, err := riskIssue.dbRaw.DB.Query(`SELECT COUNT(*) FROM risk_issue WHERE risk_issue_code = ?`, request.RiskIssueCode)

	RiskIssueCode := ""

	checkErr(err)

	if checkCount(checkRiskCode) == 0 {
		RiskIssueCode = request.RiskIssueCode
		fmt.Println("checkRisk ==>", RiskIssueCode)
	} else {
		// fmt.Println("count", checkCount(checkRiskCode))

		// var count int
		RiskIssueCode = "RE." + lib.GetTimeNow("date2")
		queryCount := riskIssue.db.DB.Table("risk_issue").
			Select(`COUNT(*) 'count'`).
			Where("risk_issue_code LIKE ?", fmt.Sprintf("%s%%", RiskIssueCode))

		var count int64
		queryCount.
			Count(&count)

		counter := count

		str := fmt.Sprintf("%04d", counter+1)

		RiskIssueCode += "." + str

		fmt.Println("checkRisk gak nol ==>", RiskIssueCode)

	}

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

	err = tx.Save(input).Error

	return input, err
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

	if request.Status == false {
		queryBuilder = queryBuilder.Where("risk_issue.status = ?", false)
	}

	if request.Status == true {
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
