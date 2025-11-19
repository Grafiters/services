package common

import (
	"fmt"
	"math"
	"riskmanagement/lib"
	models "riskmanagement/models/common"
	"time"

	"gitlab.com/golang-package-library/logger"
)

type CommonDefinition interface {
	GetNPAndNama(request models.KeywordRequest) (responses []models.PnNamaResult, err error)
	GetKanwil(request models.KeywordRequest) (responses []models.KanwilResult, err error)
	GetKanca(request models.KeywordRequest) (responses []models.KancaResult, err error)
	GetUker(request models.KeywordRequest) (responses []models.UkerResult, err error)
	GetRiskEventByActifityAndPoduct(request models.RiskEventRequest) (responses []models.RiskEventResult, err error)
	GetRiskIndikatorByRiskEvent(request models.RiskIndikatorRequest) (responses []models.RiskIndikatorResult, err error)

	GetRRMHead(request models.RRMHeadRequest) (responses []models.PnNamaResult, err error)
	GetPimpinanUker(request models.PimpinanUkerRequest) (responses []models.PnNamaResult, err error)
	GetApprovalResponse(request models.CommonRequest) (response []models.PnNamaResult, err error)
	GetAllORDMember() (response []models.ORDMember, err error)

	// MQ Enhance
	GetMstDataOption() (response []models.MasterDataOption, err error)

	SearchBrc(request models.BrcKeywordRequest) (responses []models.PnNamaResult, totalRows int, totalData int, err error)
}

type CommonRepository struct {
	db      lib.Database
	dbRaw   lib.Databases
	logger  logger.Logger
	timeout time.Duration
}

// GetMstDataOption implements CommonDefinition.
func (c CommonRepository) GetMstDataOption() (responses []models.MasterDataOption, err error) {
	db := c.db.DB

	err = db.Table(`tbl_master_data`).Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return nil, err
	}

	return responses, err
}

func NewCommonRepository(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
) CommonDefinition {
	return CommonRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// GetPimpinanUker implements CommonDefinition.
func (c CommonRepository) GetPimpinanUker(request models.PimpinanUkerRequest) (responses []models.PnNamaResult, err error) {
	db := c.db.DB

	query := db.Table(`pa0001_eof`).
		Select(`PERNR "pernr", SNAME "sname"`).
		Where(`BRANCH = ? `, request.BRANCH).
		Where(`HILFM in ('014','019','057')`)

	err = query.Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

// GetRRMHead implements CommonDefinition.
func (c CommonRepository) GetRRMHead(request models.RRMHeadRequest) (responses []models.PnNamaResult, err error) {
	db := c.db.DB

	query := db.Table(`pa0001_eof`).
		Select(`PERNR "pernr", SNAME "sname"`).
		Where(`BTRTL_TX LIKE '%RO%'`).
		Where(`TIPE_UKER IN ('KW')`).
		Where(`STELL_TX  LIKE "%RISK MANAGEMENT & COMPLIANCE%" OR STELL_TX LIKE "%REGIONAL RISK MANAGEMENT HEAD%"`).
		Where(`BTRTL = ?`, request.BTRTL).
		Where(`SNAME like ? OR PERNR like ?`, "%"+request.Keyword+"%", "%"+request.Keyword+"%")

		//TBA
		//urc
		// if request.HILFM == "034" {
		// 	query.Where(`STELL_TX LIKE "%BRANCH RISK & COMPLIANCE%" OR STELL_TX  LIKE "%RISK MANAGEMENT & COMPLIANCE%" OR STELL_TX LIKE "%REGIONAL RISK MANAGEMENT HEAD%"`)
		// } else {
		// 	query.Where(`STELL_TX  LIKE "%RISK MANAGEMENT & COMPLIANCE%" OR STELL_TX LIKE "%REGIONAL RISK MANAGEMENT HEAD%"`)
		// }

		// Where(`BTRTL = ?`, request.BTRTL).

		// Where(db.Where(`PERNR LIKE '%?%'`, request.Keyword).Or(`SNAME LIKE '%?%'`, request.Keyword))

	err = query.Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return responses, err
}

func (c CommonRepository) GetUker(request models.KeywordRequest) (responses []models.UkerResult, err error) {
	sQuery := ""
	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {

		sQuery = `
				SELECT
					uku.BRANCH 'kode',
					uku.BRDESC 'nama'
				FROM uker_kelolaan_user uku  
				WHERE uku.pn = ? AND concat(uku.BRANCH,uku.RGDESC) like ?
				AND (ISNULL(uku.expired_at) OR uku.expired_at <= NOW())`

		//sQuery = `select BRANCH as kode, BRDESC as nama  from dwh_branch db where BRANCH = (SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ? and 1 != ?)`
	} else if request.HILFM == "147" {
		sQuery = `select db2.BRANCH as kode, db2.BRDESC as nama from dwh_branch db2 where REGION  in (
						select db.REGION from dwh_branch db where BRANCH = 
							(
								SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ?
							)
						) and concat(db2.BRANCH,db2.RGDESC) like ?`
	} else if request.HILFM == "159" || request.HILFM == "160" || request.HILFM == "161" || request.HILFM == "162" || request.HILFM == "157" || request.HILFM == "158" {
		sQuery = `select db2.BRANCH as kode, db2.BRDESC as nama from dwh_branch db2 where 1 != ? and concat(db2.BRANCH,db2.BRDESC) like ?`
	}

	sQuery = sQuery + ` LIMIT ? OFFSET ? `
	err = c.db.DB.Raw(sQuery, request.PERN,
		"%"+request.Keyword+"%", request.Limit, request.Offset).Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (c CommonRepository) GetKanca(request models.KeywordRequest) (responses []models.KancaResult, err error) {
	sQuery := ""
	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		sQuery = `select MBDESC as nama  from dwh_branch db where BRANCH = (SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ? AND 1 != ?)`
	} else if request.HILFM == "147" {
		sQuery = `select distinct db2.MBDESC as nama from dwh_branch db2 where REGION  in (
						select db.REGION from dwh_branch db where BRANCH = 
							(
								SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ?
							)
						) and db2.RGDESC like ?`
	} else if request.HILFM == "159" || request.HILFM == "160" || request.HILFM == "161" || request.HILFM == "162" || request.HILFM == "157" || request.HILFM == "158" {
		sQuery = `select distinct db2.MBDESC as nama from dwh_branch db2 where 1 != ? and db2.MBDESC like ?`
	}

	sQuery = sQuery + ` LIMIT ? OFFSET ? `
	err = c.db.DB.Raw(sQuery, request.PERN,
		"%"+request.Keyword+"%", request.Limit, request.Offset).Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (c CommonRepository) GetKanwil(request models.KeywordRequest) (responses []models.KanwilResult, err error) {
	sQuery := ""
	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		sQuery = `select RGDESC as nama  from dwh_branch db where BRANCH = (SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ? AND 1 != ?)`
	} else if request.HILFM == "147" {
		sQuery = `select distinct db2.RGDESC as nama from dwh_branch db2 where REGION  in (
						select db.REGION from dwh_branch db where BRANCH = 
							(
								SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ?
							)
						) and db2.RGDESC like ?`
	} else if request.HILFM == "159" || request.HILFM == "160" || request.HILFM == "161" || request.HILFM == "162" || request.HILFM == "157" || request.HILFM == "158" {
		sQuery = `select distinct db2.RGDESC as nama from dwh_branch db2 where 1 != ? and db2.RGDESC like ?`
	}

	sQuery = sQuery + ` LIMIT ? OFFSET ? `
	err = c.db.DB.Raw(sQuery, request.PERN,
		"%"+request.Keyword+"%", request.Limit, request.Offset).Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (c CommonRepository) GetNPAndNama(request models.KeywordRequest) (responses []models.PnNamaResult, err error) {
	sQuery := ""
	if request.HILFM == "033" || request.HILFM == "034" || request.HILFM == "228" {
		sQuery = `select pe.pernr as pernr, pe.sname as sname from pa0001_eof pe where pe.PERNR  = ? `
	} else if request.HILFM == "174" {
		sQuery = `SELECT pe2.PERNR as pernr,pe2.SNAME as sname  from pa0001_eof pe2 where CAST(pe2.BRANCH  AS DECIMAL) in (
		SELECT db2.BRANCH from dwh_branch db2 where db2.REGION = (
		select db.REGION  from dwh_branch db	where db.BRANCH = (
	      SELECT CAST(pe.BRANCH  AS DECIMAL)  FROM pa0001_eof pe WHERE pe.PERNR = ?
	      )
	      )
	      ) and pe2.HILFM in ('033','034','228') and concat(pe2.SNAME, pe2.PERNR) like ? `
	} else if request.HILFM == "159" || request.HILFM == "160" || request.HILFM == "161" || request.HILFM == "162" || request.HILFM == "163" || request.HILFM == "157" || request.HILFM == "158" {
		sQuery = `select pe.pernr, pe.sname from pa0001_eof pe where pe.HILFM in ('033','034','228') and concat(pe.SNAME, pe.PERNR) like ?`
	}

	sQuery = sQuery + " LIMIT ? OFFSET ? "
	if request.HILFM == "174" {
		err = c.db.DB.Raw(sQuery, request.PERN,
			"%"+request.Keyword+"%", request.Limit, request.Offset).Scan(&responses).Error

		c.logger.Zap.Info(responses[0].Sname)
		if err != nil {
			c.logger.Zap.Error(err)
			return responses, err
		}
	} else if request.HILFM == "034" || request.HILFM == "033" || request.HILFM == "228" {
		err = c.db.DB.Raw(sQuery, request.PERN, request.Limit, request.Offset).Scan(&responses).Error

		c.logger.Zap.Info(responses[0].Sname)
		if err != nil {
			c.logger.Zap.Error(err)
			return responses, err
		}
	} else {
		err = c.db.DB.Raw(sQuery,
			"%"+request.Keyword+"%", request.Limit, request.Offset).Scan(&responses).Error

		c.logger.Zap.Info(responses[0].Sname)
		if err != nil {
			c.logger.Zap.Error(err)
			return responses, err
		}
	}

	return responses, err
}

func (c CommonRepository) GetRiskEventByActifityAndPoduct(request models.RiskEventRequest) (responses []models.RiskEventResult, err error) {
	sQuery := `select DISTINCT ri.id, ri.risk_issue_code , ri.risk_issue from risk_issue ri 
                          	join risk_issue_map_aktifitas rima on ri.id  = rima.id_risk_issue
                          	join risk_issue_map_product rimp on ri.id = rimp.id_risk_issue 
                          	where rima.aktifitas = ?
                          	and rimp.product = ?`

	err = c.db.DB.Raw(sQuery, request.ActivityID, request.ProductID).Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (c CommonRepository) GetRiskIndikatorByRiskEvent(request models.RiskIndikatorRequest) (responses []models.RiskIndikatorResult, err error) {
	sQuery := ` select DISTINCT ri.id,ri.risk_indicator_code,ri.risk_indicator  from risk_indicator ri 
                         join risk_issue_map_indicator rimi on rimi.id_indicator = ri.id
                         join risk_issue ri3 on ri3.id = rimi.id_indicator
                         where ri3.id = ?`

	err = c.db.DB.Raw(sQuery, request.RiskEventID).Scan(&responses).Error

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}
	return responses, err
}

func (c CommonRepository) GetApprovalResponse(request models.CommonRequest) (response []models.PnNamaResult, err error) {
	db := c.db.DB

	query := db.Table(`pa0001_eof`).
		Select(`PERNR "pernr", SNAME "sname"`).
		// Where("KOSTL = 'PS21014'").
		Where(`SNAME like ? OR PERNR like ?`, "%"+request.Keyword+"%", "%"+request.Keyword+"%").
		Limit(10)

	err = query.Scan(&response).Error
	if err != nil {
		return response, err
	}

	return response, err
}

func (c CommonRepository) GetAllORDMember() (response []models.ORDMember, err error) {
	db := c.db.DB

	query := db.Table(`pa0001_eof`).
		Select(`PERNR "pernr", SNAME "sname"`).
		Where(`KOSTL = 'PS21014'`)

	err = query.Scan(&response).Error
	if err != nil {
		return response, err
	}

	return response, err
}

// SearchBrc implements CommonDefinition.
func (c CommonRepository) SearchBrc(request models.BrcKeywordRequest) (responses []models.PnNamaResult, totalRows int, totalData int, err error) {
	db := c.db.DB.Table("mst_uker_kelolaan muk")

	keyword := fmt.Sprintf("%%%s%%", request.Keyword)
	query := db.Select(`muk.pn 'pernr', muk.sname`).
		Joins("inner join uker_kelolaan_user uku ON uku.pn = muk.pn").
		Where(`CONCAT(muk.pn, muk.sname) like ?`, keyword).Group("muk.pn")

	if request.TipeUker != "KP" {
		query = query.Where(`uku.REGION = ?`, request.Region)
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

	err = db.Limit(request.Limit).Offset(request.Offset).Scan(&responses).Error

	if err != nil {
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}
