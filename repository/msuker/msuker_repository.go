package msuker

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/msuker"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type MsUkerDefinition interface {
	WithTrx(trxHandle *gorm.DB) MsUkerRepository
	GetAll() (responses []models.MsUkerResponse, err error)
	GetUkerByBranch(branchid int64) (responses []models.MsUkerResponse, err error)
	SearchUker(request *models.KeywordRequest) (responses []models.MsUkerResponseNull, totalRows int, totalData int, err error)
	GetUkerPerRegion(request *models.Region) (responses []models.MsUkerResponse, err error)
	SearchPeserta(request *models.KeywordRequest) (responses []models.MsPesertaNull, totalRows int, totalData int, err error)
	GetPekerjaByBranch(request *models.BranchCodeInduk) (responses []models.MsPekerjaResponse, err error)
	GetPekerjaByRegion(request *models.SearchPNByRegionReq) (responses []models.SearchPNByRegionRes, err error)
	SearchJabatan(request *models.KeywordRequest) (responses []models.Jabatan, totalRows int, totalData int, err error)
	SearchUkerByRegionPekerja(request *models.KeyRequest) (responses []models.MsUkerResponse, totalRows int, totalData int, err error)
	SearchPekerjaPerUker(request *models.KeywordRequest) (responses []models.MsPekerjaResponse, totalRows int, totalData int, err error)
	SearchSigner(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error)
	SearchRMC(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error)
	SearchPelakuFraud(request *models.KeywordRequest) (responses []models.MsPelaku, totalRows int, totalData int, err error)
	SearchBrcUrcPerRegion(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error)
	// Batch 3
	CheckJumlahPekerja(BRANCH string) (Total int64, err error)
	ListingJabatanPerUker(request *models.ListJabatanRequest) (responses []models.ListJabatanResponse, err error)
	GetPekerjaBranchByHILFM(request models.BranchByHilfmRequest) (responses []models.MsPeserta, err error)
	SearchRRMHead(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error)
	SearchPekerjaOrd(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error)
}

type MsUkerRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

func NewMsUkerRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) MsUkerDefinition {
	return MsUkerRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetAll implements MsUkerDefinition
func (msUker MsUkerRepository) GetAll() (responses []models.MsUkerResponse, err error) {
	return responses, msUker.db.DB.Find(&responses).Error
}

// GetUkerByBranch implements MsUkerDefinition
func (msUker MsUkerRepository) GetUkerByBranch(branchid int64) (responses []models.MsUkerResponse, err error) {
	return responses, msUker.db.DB.Where("BRANCH = ?", branchid).Find(&responses).Error
}

// SearchJabatan implements MsUkerDefinition
func (msUker MsUkerRepository) SearchJabatan(request *models.KeywordRequest) (responses []models.Jabatan, totalRows int, totalData int, err error) {
	where := ""

	keyword := fmt.Sprintf("%%%s%%", request.Keyword)

	if request.Keyword != "" {
		where += ` WHERE CONCAT(pe.HILFM, pe.HTEXT) LIKE ?`
	}

	query := `SELECT 
					pe.HILFM,
					pe.HTEXT,
					pe.STELL_TX 'STELLTX'
				FROM pa0001_eof pe` + where + ` GROUP BY pe.HILFM ORDER BY pe.HTEXT ASC LIMIT ? OFFSET ?`
	msUker.logger.Zap.Info(query)
	rows, err := msUker.dbRaw.DB.Query(query, keyword, request.Limit, request.Offset)
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	msUker.logger.Zap.Info("rows ", rows)

	response := models.Jabatan{}
	for rows.Next() {
		err = rows.Scan(
			&response.HILFM,
			&response.HTEXT,
			&response.STELLTX,
		)
		if err != nil {
			return responses, totalRows, totalData, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT 
							count(t.hilfm)
						FROM (
						SELECT 
							DISTINCT 
							pe.HILFM,
							pe.HTEXT 
						FROM pa0001_eof pe ` + where + `) as t`

	err = msUker.dbRaw.DB.QueryRow(paginationQuery, keyword).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// SearchUker implements MsUkerDefinition
func (msUker MsUkerRepository) SearchUker(request *models.KeywordRequest) (responses []models.MsUkerResponseNull, totalRows int, totalData int, err error) {
	where := ""

	keyword := fmt.Sprintf("%%%s%%", request.Keyword)

	if request.Keyword != "" {
		where += ` WHERE concat("BRANCH","BRNAME","BRDESC") LIKE ?`
	}

	query := `SELECT * FROM dwh_branch` + where + ` LIMIT ? OFFSET ?`
	msUker.logger.Zap.Info(query)
	rows, err := msUker.dbRaw.DB.Query(query, keyword, request.Limit, request.Offset)
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()
	msUker.logger.Zap.Info("rows ", rows)

	response := models.MsUkerResponseNull{}
	for rows.Next() {
		err = rows.Scan(
			&response.SRCSYSID,
			&response.BRUNIT,
			&response.REGION,
			&response.RGDESC,
			&response.RGNAME,
			&response.MAINBR,
			&response.MBDESC,
			&response.MBNAME,
			&response.SUBBR,
			&response.SBDESC,
			&response.SBNAME,
			&response.BRANCH,
			&response.BRDESC,
			&response.BRNAME,
			&response.BIBR,
		)
		if err != nil {
			return responses, totalRows, totalData, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT count(*) FROM dwh_branch` + where

	err = msUker.dbRaw.DB.QueryRow(paginationQuery, keyword, request.Limit).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// SearchUkerByRegionPekerja implements MsUkerDefinition
func (msUker MsUkerRepository) SearchUkerByRegionPekerja(request *models.KeyRequest) (responses []models.MsUkerResponse, totalRows int, totalData int, err error) {
	firstString := request.PN[0:1]
	where := ""

	keyword := fmt.Sprintf("%%%s%%", request.Keyword)

	if firstString == "0" {
		where = ` WHERE REGION = (
			SELECT db2.REGION
			FROM dwh_branch db2
			WHERE BRANCH = (
				SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ?
			)
		)`
	} else {
		where = ` WHERE REGION = (
			SELECT db2.REGION
			FROM dwh_branch db2
			WHERE BRANCH = (
				SELECT po.UnitKerja FROM pekerja_outsource po WHERE po.PersonalNumber  = ?
			)
		)`
	}

	if request.Keyword != "" {
		where += " AND concat(BRANCH,BRNAME,BRDESC) like '%" + request.Keyword + "%'"
	}

	query := `SELECT * FROM dwh_branch` + where + ` LIMIT ? OFFSET ?`
	msUker.logger.Zap.Info(query)
	rows, err := msUker.dbRaw.DB.Query(query, request.PN, keyword, request.Limit, request.Offset)
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	msUker.logger.Zap.Info("rows ", rows)

	response := models.MsUkerResponse{}
	for rows.Next() {
		err = rows.Scan(
			&response.SRCSYSID,
			&response.BRUNIT,
			&response.REGION,
			&response.RGDESC,
			&response.RGNAME,
			&response.MAINBR,
			&response.MBDESC,
			&response.MBNAME,
			&response.SUBBR,
			&response.SBDESC,
			&response.SBNAME,
			&response.BRANCH,
			&response.BRDESC,
			&response.BRNAME,
			&response.BIBR,
		)
		if err != nil {
			return responses, totalRows, totalData, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, totalRows, totalData, err
	}

	paginationQuery := `SELECT count(*) FROM dwh_branch` + where

	err = msUker.dbRaw.DB.QueryRow(paginationQuery, request.PN, keyword).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetUkerPerRegion implements MsUkerDefinition
func (msUker MsUkerRepository) GetUkerPerRegion(request *models.Region) (responses []models.MsUkerResponse, err error) {
	query := `SELECT * FROM dwh_branch WHERE REGION = ?`

	rows, err := msUker.dbRaw.DB.Query(query, request.REGION)
	defer rows.Close()

	if err != nil {
		return responses, err
	}

	response := models.MsUkerResponse{}
	for rows.Next() {
		_ = rows.Scan(
			&response.SRCSYSID,
			&response.BRUNIT,
			&response.REGION,
			&response.RGDESC,
			&response.RGNAME,
			&response.MAINBR,
			&response.MBDESC,
			&response.MBNAME,
			&response.SUBBR,
			&response.SBDESC,
			&response.SBNAME,
			&response.BRANCH,
			&response.BRDESC,
			&response.BRNAME,
			&response.BIBR,
		)

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return responses, err
	}

	return responses, err
}

// SearchPeserta implements MsUkerDefinition
func (repo MsUkerRepository) SearchPeserta(request *models.KeywordRequest) (responses []models.MsPesertaNull, totalRows int, totalData int, err error) {
	wherePA := ""
	wherePO := ""

	keyword := fmt.Sprintf("%%%s%%", request.Keyword)

	if request.Keyword != "" {
		wherePA += ` WHERE concat(SNAME, PERNR) LIKE ?`
		wherePO += ` WHERE concat(po.SNAME, po.PERNR) LIKE ?`
	}

	query := `SELECT PERNR,SNAME,STELL_TX FROM pa0001_eof` + wherePA + `
			UNION 
			SELECT * FROM 
			(SELECT PersonalNumber "PERNR", Nama "SNAME", desc_posisi "STELL_TX" FROM pekerja_outsource) as po` + wherePO + ` LIMIT ? OFFSET ?`

	repo.logger.Zap.Info(query)

	rows, err := repo.dbRaw.DB.Query(query, keyword, keyword, request.Limit, request.Offset)
	if err != nil {
		return responses, totalRows, totalData, err
	}
	defer rows.Close()

	// repo.logger.Zap.Info("rows =>", rows)

	response := models.MsPesertaNull{}
	for rows.Next() {
		err = rows.Scan(
			&response.PERNR,
			&response.SNAME,
			&response.STEELTX,
		)
		if err != nil {
			return responses, totalRows, totalData, err
		}

		responses = append(responses, response)
	}

	paginationQuery := `SELECT SUM(pa_count) FROM (SELECT count(*) 'pa_count' FROM pa0001_eof
	` + wherePA + `
	UNION 
	SELECT count(*) FROM (SELECT Nama 'SNAME', PersonalNumber 'PERNR',desc_posisi 'STELL_TX' FROM pekerja_outsource) AS po
	` + wherePO + `) AS Peserta`

	err = repo.dbRaw.DB.QueryRow(paginationQuery, keyword, keyword).Scan(&totalData)
	if err != nil {
		return responses, totalRows, totalData, err
	}

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}
	return responses, totalRows, totalData, err
}

// GetPekerjaByBranch implements MsUkerDefinition
func (repo MsUkerRepository) GetPekerjaByBranch(request *models.BranchCodeInduk) (responses []models.MsPekerjaResponse, err error) {
	where1 := ""
	where2 := ""
	args1 := []interface{}{}
	args2 := []interface{}{}

	if request.BRANCH != "" {
		where1 = "CAST(BRANCH AS UNSIGNED) = ?"
		args1 = append(args1, request.BRANCH)
		where2 = "UnitKerja = ?"
		args2 = append(args2, request.BRANCH)
	}

	if request.KOSTL != "" && request.WERKS != "" {
		where1 = "KOSTL = ? AND WERKS = ?"
		args1 = append(args1, request.KOSTL, request.WERKS)
		where2 = "kostl = ? AND werks = ?"
		args2 = append(args2, request.KOSTL, request.WERKS)
	}

	sql := `SELECT
				PERNR,
				SNAME,
				STELL_TX 'STEELTX',
				CAST( BRANCH AS UNSIGNED ) AS BRANCH 
			FROM
				pa0001_eof 
			WHERE ` + where1 + `
			UNION
			SELECT
				PersonalNumber 'PERNR',
				Nama 'SNAME',
				desc_posisi 'STEELTX',
				UnitKerja AS 'BRANCH' 
			FROM
				pekerja_outsource os 
			WHERE ` + where2

	// repo.logger.Zap.Info(query)

	db := repo.db.DB.Raw(sql, append(args1, args2...)...)

	err = db.Scan(&responses).Error

	return responses, err
}

func (repo MsUkerRepository) GetPekerjaByRegion(request *models.SearchPNByRegionReq) (responses []models.SearchPNByRegionRes, err error) {
	hilfms := strings.Split(request.ParameterBrc, ",")

	db := repo.db.DB.Table("pa0001_eof pe").
		Select(`pe.PERNR, pe.SNAME`).
		Joins(`JOIN dwh_branch db on db.BRANCH = TRIM(LEADING '0' FROM pe.BRANCH)`).
		Where(`db.REGION = ?`, request.REGION).
		Where(`pe.HILFM in (?)`, hilfms).
		Where("concat(SNAME, PERNR) like ?", fmt.Sprintf("%%%s%%", request.Keyword))

	err = db.Scan(&responses).Error

	return responses, err
}

// WithTrx implements MsUkerDefinition
func (msUker MsUkerRepository) WithTrx(trxHandle *gorm.DB) MsUkerRepository {
	if trxHandle == nil {
		msUker.logger.Zap.Error("transaction Database not found in gin context.")
		return msUker
	}
	msUker.db.DB = trxHandle
	return msUker
}

// SearchPekerjaPerRegion implements MsUkerDefinition
func (repo MsUkerRepository) SearchPekerjaPerUker(request *models.KeywordRequest) (responses []models.MsPekerjaResponse, totalRows int, totalData int, err error) {
	db := repo.db.DB
	wherePA := ""
	wherePO := ""
	args1 := []interface{}{}
	args2 := []interface{}{}

	keyword := fmt.Sprintf("%%%s%%", request.Keyword)

	if request.Keyword != "" && request.BRANCH != "" {
		wherePA += ` WHERE concat(SNAME, PERNR) LIKE ? AND CAST( BRANCH AS UNSIGNED ) = ?`
		args1 = append(args1, keyword, request.BRANCH)
		wherePO += ` WHERE concat(PersonalNumber, Nama) LIKE ? AND UnitKerja = ?`
		args2 = append(args2, keyword, request.BRANCH)
	}

	if request.Keyword == "" && request.BRANCH != "" {
		wherePA += ` WHERE CAST( BRANCH AS UNSIGNED ) = ?`
		args1 = append(args1, request.BRANCH)
		wherePO += ` WHERE UnitKerja = ?`
		args2 = append(args2, request.BRANCH)
	}

	if request.Keyword != "" && request.KOSTL != "" {
		wherePA += ` WHERE concat(SNAME, PERNR) LIKE ? AND KOSTL = ?`
		args1 = append(args1, keyword, request.KOSTL)
		wherePO += ` WHERE concat(PersonalNumber, Nama) LIKE ? AND kostl = ?`
		args1 = append(args1, keyword, request.KOSTL)
	}

	if request.Keyword == "" && request.KOSTL != "" {
		wherePA += ` WHERE KOSTL = ?`
		args1 = append(args1, request.KOSTL)
		wherePO += ` WHERE kostl = ?`
		args1 = append(args1, request.KOSTL)
	}

	sql := `SELECT PERNR,SNAME,STELL_TX 'STEELTX' FROM pa0001_eof ` + wherePA + `
		UNION
		SELECT PersonalNumber "PERNR", Nama "SNAME", desc_posisi "STEELTX"
		FROM pekerja_outsource` + wherePO + ` LIMIT ? OFFSET ?`

	argsFinal := append(args1, args2...)
	argsFinal = append(argsFinal, request.Limit, request.Offset)

	db = db.Raw(sql, argsFinal...)

	err = db.Scan(&responses).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalData = len(responses) // Simplistic; not actual count of all rows matching criteria
	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit))) // Calculate the total number of pages
	}
	return responses, totalRows, totalData, err
}

// SearchSigner implements MsUkerDefinition
func (repo MsUkerRepository) SearchSigner(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error) {
	db := repo.db.DB

	jgpg, err := lib.GetVarEnv("signerJG")
	if err != nil {
		return responses, totalRows, totalData, fmt.Errorf("error getting RRM STELLs: %w", err)
	}

	jgpgs := strings.Split(jgpg, ",")

	var ors []string
	var vals []interface{}
	for _, s := range jgpgs {
		s = strings.TrimSpace(s)
		if s != "" {
			ors = append(ors, "JGPG LIKE ?")
			vals = append(vals, s+"%")
		}
	}

	db = db.Table("pa0001_eof").
		Select(`PERNR,SNAME,STELL_TX as 'STEELTX'`).
		Where("concat(SNAME, PERNR) like ?", fmt.Sprintf("%%%s%%", request.Keyword)).
		Where(strings.Join(ors, " OR "), vals...)
		// Where("JGPG LIKE 'JG11%' OR JGPG LIKE 'JG12%' OR JGPG LIKE 'JG13%' OR JGPG LIKE 'JG14%' OR JGPG LIKE 'JG15%' OR JGPG LIKE 'JG16%'")

	if request.TipeUker != "" {
		db = db.Where("TIPE_UKER = ?", request.TipeUker)
	}

	if request.BRANCH != "" {
		if request.TipeUker == "KP" {
			db = db.Where("BRANCH = ?", request.BRANCH).
				Where("KOSTL = ?", request.KOSTL)
		} else {
			db = db.Where("BRANCH = ?", request.BRANCH)
		}
	}

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalData = int(count)
	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error

	if err != nil {
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}

// SearchRMC implements MsUkerDefinition
func (repo MsUkerRepository) SearchRMC(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error) {
	db := repo.db.DB

	db = db.Table("pa0001_eof").
		Select(`PERNR,SNAME,STELL_TX as 'STEELTX'`).
		Where("concat(SNAME, PERNR) like ?", fmt.Sprintf("%%%s%%", request.Keyword))

	stells, err := lib.GetVarEnv("signerRMC")
	if err != nil {
		return responses, totalRows, totalData, fmt.Errorf("error getting RRM STELLs: %w", err)
	}

	stellTx := strings.Split(stells, ",")

	if request.TipeUker != "" {
		if request.TipeUker == "KP" {
			db = db.
				// Where("JGPG LIKE 'JG11%' OR JGPG LIKE 'JG12%' OR JGPG LIKE 'JG13%' OR JGPG LIKE 'JG14%' OR JGPG LIKE 'JG15%' OR JGPG LIKE 'JG16%'").
				Where(`TIPE_UKER = "KP"`).
				Where("KOSTL = ?", request.KOSTL)
		} else {
			var ors []string
			var vals []interface{}
			for _, s := range stellTx {
				s = strings.TrimSpace(s)
				if s != "" {
					ors = append(ors, "STELL_TX LIKE ?")
					vals = append(vals, "%"+s+"%")
				}
			}

			db = db.Where(`BTRTL_TX  LIKE '%RO%'`).
				Where(`TIPE_UKER IN ('KW')`).
				Where(strings.Join(ors, " OR "), vals...).
				Where(`BTRTL = ?`, request.BTRTL)
		}
	}

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalData = int(count)
	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error

	if err != nil {
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}

// SearchPelakuFraud implements MsUkerDefinition
func (repo MsUkerRepository) SearchPelakuFraud(request *models.KeywordRequest) (responses []models.MsPelaku, totalRows int, totalData int, err error) {
	wherePA := ""
	wherePO := ""
	argsPA := []interface{}{}
	argsPO := []interface{}{}

	keyword := fmt.Sprintf("%%%s%%", request.Keyword)

	// #Lock By Uker
	// if request.BRANCH != "" {
	// 	wherePA = ` concat(SNAME, PERNR) LIKE ? AND CAST( BRANCH AS UNSIGNED ) = ?`
	// 	argsPA = append(argsPA, keyword, request.BRANCH)
	// 	wherePO = ` concat(Nama, PersonalNumber) LIKE ? AND UnitKerja = ?`
	// 	argsPO = append(argsPO, keyword, request.BRANCH)
	// }

	// if request.KOSTL != "" {
	// 	wherePA = ` concat(SNAME, PERNR) LIKE ? AND KOSTL = ?`
	// 	argsPA = append(argsPA, keyword, request.KOSTL)
	// 	wherePO = ` concat(Nama, PersonalNumber) LIKE ? AND kostl = ?`
	// 	argsPO = append(argsPO, keyword, request.KOSTL)
	// }

	// #unlock by uker
	wherePA = ` concat(SNAME, PERNR) LIKE ?`
	argsPA = append(argsPA, keyword)
	wherePO = ` concat(Nama, PersonalNumber) LIKE ?`
	argsPO = append(argsPO, keyword)

	sql := `SELECT PERNR,SNAME, STELL_TX 'STEELTX' FROM pa0001_eof 
			WHERE ` + wherePA + `
			UNION
			SELECT PersonalNumber 'PERNR', Nama 'SNAME', desc_posisi 'STEELTX' FROM pekerja_outsource os 
			WHERE ` + wherePO + ` LIMIT ? OFFSET ?`

	argsFinal := append(argsPA, argsPO...)
	argsFinal = append(argsFinal, request.Limit, request.Offset)

	db := repo.db.DB.Raw(sql, argsFinal...)

	err = db.Scan(&responses).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalData = len(responses) // This might need proper querying to get the real count

	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	return responses, totalRows, totalData, err
}

// SearchBrcUrcPerRegion implements MsUkerDefinition
func (repo MsUkerRepository) SearchBrcUrcPerRegion(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error) {
	db := repo.db.DB

	db = db.Table("pa0001_eof").
		Select(`PERNR,SNAME,STELL_TX as 'STEELTX'`).
		Where("concat(SNAME, PERNR) like ?", fmt.Sprintf("%%%s%%", request.Keyword))

	hilfms := strings.Split(request.ParameterBrc, ",")

	if request.KOSTL != "PS21014" {
		db = db.
			Where("HILFM in (?)", hilfms).Where(`BTRTL = ?`, request.BTRTL)
	}

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalData = int(count)
	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	fmt.Println("Limit =>", request.Limit)
	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Scan(&responses).Error

	if err != nil {
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}

// CheckJumlahPekerja implements MsUkerDefinition.
func (msUker MsUkerRepository) CheckJumlahPekerja(BRANCH string) (Total int64, err error) {
	query := `SELECT SUM(TotalPekerjaUker.Total) 'Total' FROM(
				SELECT count(*) 'Total' FROM pa0001_eof WHERE CAST(BRANCH AS float) = ?
				UNION 
				SELECT count(*) 'Total' FROM pekerja_outsource WHERE CAST(brname AS float) = ?
			) TotalPekerjaUker`

	msUker.logger.Zap.Info(query)

	// QueryRow is used since the query is expected to return a single value.
	err = msUker.dbRaw.DB.QueryRow(query, BRANCH, BRANCH).Scan(&Total)

	if err != nil {
		return 0, err
	}

	return Total, nil
}

// ListingJabatanPerUker implements MsUkerDefinition.
func (msUker MsUkerRepository) ListingJabatanPerUker(request *models.ListJabatanRequest) (responses []models.ListJabatanResponse, err error) {
	where1 := ""
	where2 := ""
	args1 := []interface{}{}
	args2 := []interface{}{}

	if request.BRANCH != "" {
		where1 = "CAST(BRANCH AS UNSIGNED) = ?"
		args1 = append(args1, request.BRANCH)
		where2 = "UnitKerja = ?"
		args2 = append(args2, request.BRANCH)
	}

	if request.KOSTL != "" && request.WERKS != "" {
		where1 = "KOSTL = ? AND WERKS = ?"
		args1 = append(args1, request.KOSTL, request.WERKS)
		where2 = "kostl = ? AND werks = ?"
		args2 = append(args2, request.KOSTL, request.WERKS)
	}

	sql := `SELECT HILFM, HTEXT , count(*) 'Jumlah' FROM pa0001_eof pe 
			WHERE ` + where1 + `
			GROUP BY HILFM 
			UNION 
			SELECT
				posisi 'HILFM',
				desc_posisi 'HTEXT',	
				COUNT(*) Jumlah
			FROM pekerja_outsource po WHERE ` + where2 + `  GROUP BY posisi`

	db := msUker.db.DB.Raw(sql, append(args1, args2...)...)

	err = db.Scan(&responses).Error

	return responses, err
}

// GetPekerjaBranchByHILFM implements MsUkerDefinition.
func (repo MsUkerRepository) GetPekerjaBranchByHILFM(request models.BranchByHilfmRequest) (responses []models.MsPeserta, err error) {
	var where1Parts []string
	var where2Parts []string
	var args1 []interface{}
	var args2 []interface{}

	if request.BRANCH != "" {
		where1Parts = append(where1Parts, "CAST(BRANCH AS UNSIGNED) = ?")
		args1 = append(args1, request.BRANCH)

		where2Parts = append(where2Parts, "UnitKerja = ?")
		args2 = append(args2, request.BRANCH)
	}

	// Kondisi berdasarkan KOSTL dan WERKS
	if request.KOSTL != "" && request.WERKS != "" {
		where1Parts = append(where1Parts, "KOSTL = ?", "WERKS = ?")
		args1 = append(args1, request.KOSTL, request.WERKS)

		where2Parts = append(where2Parts, "kostl = ?", "werks = ?")
		args2 = append(args2, request.KOSTL, request.WERKS)
	}

	// Tambahkan HILFM jika tidak kosong
	if request.HILFM != "" {
		where1Parts = append(where1Parts, "HILFM = ?")
		args1 = append(args1, request.HILFM)

		where2Parts = append(where2Parts, "posisi = ?")
		args2 = append(args2, request.HILFM)
	}

	// Gabungkan kondisi where menjadi string SQL
	where1 := strings.Join(where1Parts, " AND ")
	where2 := strings.Join(where2Parts, " AND ")

	sql := `SELECT PERNR,SNAME, STELL_TX 'STEELTX' FROM pa0001_eof 
			WHERE ` + where1 + `
			UNION
			SELECT PersonalNumber 'PERNR', Nama 'SNAME', desc_posisi 'STEELTX' FROM pekerja_outsource os 
			WHERE ` + where2

	db := repo.db.DB.Raw(sql, append(args1, args2...)...)

	err = db.Scan(&responses).Error

	return responses, err
}

// SearchRRMHead implements MsUkerDefinition.
func (msUker MsUkerRepository) SearchRRMHead(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error) {
	db := msUker.db.DB.Table(`pa0001_eof`)

	db = db.Select(`PERNR,SNAME,STELL_TX as 'STEELTX'`).
		Where("concat(SNAME, PERNR) like ?", fmt.Sprintf("%%%s%%", request.Keyword))

	stells, err := lib.GetVarEnv("signerRRM")
	if err != nil {
		return responses, totalRows, totalData, fmt.Errorf("error getting RRM STELLs: %w", err)
	}

	stellTx := strings.Split(stells, ",")

	if request.TipeUker != "" {
		if request.TipeUker == "KP" {
			//  'JG11%' OR JGPG LIKE 'JG12%' OR  JGPG LIKE 'JG13%' OR
			db = db.
				// Where("JGPG LIKE JGPG LIKE 'JG14%' OR JGPG LIKE 'JG15%' OR JGPG LIKE 'JG16%'").
				Where(`TIPE_UKER = "KP"`).
				Where("KOSTL = ?", request.KOSTL)
		} else {
			var ors []string
			var vals []interface{}
			for _, s := range stellTx {
				s = strings.TrimSpace(s)
				if s != "" {
					ors = append(ors, "STELL_TX LIKE ?")
					vals = append(vals, "%"+s+"%")
				}
			}

			db = db.Where(`BTRTL_TX  LIKE '%RO%'`).
				Where(`TIPE_UKER IN ('KW')`).
				Where(strings.Join(ors, " OR "), vals...).
				Where(`BTRTL = ?`, request.BTRTL).
				Order(`JGPG DESC`)
		}
	}

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalData = int(count)
	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error

	if err != nil {
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}

// SearchPekerjaOrd implements MsUkerDefinition.
func (msUker MsUkerRepository) SearchPekerjaOrd(request *models.KeywordRequest) (responses []models.MsPeserta, totalRows int, totalData int, err error) {
	db := msUker.db.DB.Table(`pa0001_eof`)

	kostl, err := lib.GetVarEnv("signerORD")

	if err != nil {
		return responses, totalRows, totalData, fmt.Errorf("error getting ORD KOSTL: %w", err)
	}

	kostls := strings.Split(kostl, ",")

	db = db.Select(`PERNR,SNAME,STELL_TX as 'STEELTX'`).
		Where("concat(SNAME, PERNR) like ?", fmt.Sprintf("%%%s%%", request.Keyword)).
		Where("KOSTL in (?)", kostls)

	var count int64
	err = db.Count(&count).Error
	if err != nil {
		return responses, totalRows, totalData, err
	}

	totalData = int(count)
	if totalData > 0 {
		totalRows = int(math.Ceil(float64(totalData) / float64(request.Limit)))
	}

	err = db.
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&responses).Error

	if err != nil {
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}
