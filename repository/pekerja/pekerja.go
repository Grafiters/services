package pekerja

import (
	"fmt"
	lib "riskmanagement/lib"
	"strings"
	"time"

	models "riskmanagement/models/pekerja"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type PekerjaDefinition interface {
	WithTrx(trxHandle *gorm.DB) PekerjaRepository
	GetApproval(request *models.RequestApproval) (response []models.DataPekerjaResponse, err error)
	GetAllPekerjaBranch(request *models.PekerjaUkerRequest) (response []models.DataPekerjaResponse, err error)
}

type PekerjaRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewPekerjaRepository(
	db lib.Database,
	logger logger.Logger,
) PekerjaDefinition {
	return PekerjaRepository{
		db:      db,
		logger:  logger,
		timeout: 0,
	}
}

// WithTrx implements PekerjaDefinition.
func (p PekerjaRepository) WithTrx(trxHandle *gorm.DB) PekerjaRepository {
	if trxHandle == nil {
		p.logger.Zap.Error("transaction Database not found in gin context")
		return p
	}

	p.db.DB = trxHandle
	return p
}

// GetAllPekerjaBranc implements PekerjaDefinition.
func (p PekerjaRepository) GetAllPekerjaBranch(request *models.PekerjaUkerRequest) (response []models.DataPekerjaResponse, err error) {
	where1 := ""
	where2 := ""
	args1 := []interface{}{}
	args2 := []interface{}{}

	if request.Branch != "" {
		branches := strings.Split(request.Branch, ",")

		where1 = "CAST(pe.BRANCH AS UNSIGNED) in  (?)"
		args1 = append(args1, branches)
		where2 = "os.UnitKerja in (?)"
		args2 = append(args2, branches)
	}

	if request.Kostl != "" && request.Werks != "" {
		where1 = "pe.KOSTL = ? AND pe.WERKS = ?"
		args1 = append(args1, request.Kostl, request.Werks)
		where2 = "os.kostl = ? AND os.werks = ?"
		args2 = append(args2, request.Kostl, request.Werks)
	}

	query := `SELECT
				pe.PERNR 'pernr',
				pe.SNAME 'sname',
				pe.STELL 'stell',
				pe.STELL_TX 'stell_tx',
				CAST(pe.BRANCH AS UNSIGNED) AS branch 
			FROM
				pa0001_eof pe
			WHERE ` + where1 + `
			UNION
			SELECT
				os.PersonalNumber 'pernr',
				os.Nama 'sname',
				os.Posisi 'stell',
				os.desc_posisi 'stell_tx',
				os.UnitKerja AS 'branch' 
			FROM
				pekerja_outsource os 
			WHERE ` + where2

	// fmt.Println(query)
	db := p.db.DB.Raw(query, append(args1, args2...)...)

	err = db.Scan(&response).Error

	if err != nil {
		p.logger.Zap.Error("Error scan column")
		return response, err
	}

	return response, nil
}

// GetApproval implements PekerjaDefinition.
func (p PekerjaRepository) GetApproval(request *models.RequestApproval) (response []models.DataPekerjaResponse, err error) {
	db := p.db.DB.Table("pa0001_eof pe").
		Select(`
			PERNR 'pernr',
			SNAME 'sname',
			STELL_TX 'stell_tx',
			CAST( BRANCH AS UNSIGNED ) AS 'branch'`).
		Where(`(pe.JGPG LIKE 'JG07%' 
			OR pe.JGPG LIKE 'JG08%' 
			OR pe.JGPG LIKE 'JG09%' 
			OR pe.JGPG LIKE 'JG10%' 
			OR pe.JGPG LIKE 'JG11%' 
			OR pe.JGPG LIKE 'JG12%'
			OR pe.JGPG LIKE 'JG13%'
			OR pe.JGPG LIKE 'JG14%'
			OR pe.JGPG LIKE 'JG15%'
			OR pe.JGPG LIKE 'JG16%')`).
		Where("concat(pe.SNAME, pe.PERNR) like ?", fmt.Sprintf("%%%s%%", request.Keyword)).
		Limit(int(request.Limit)).Offset(int(request.Offset))

	err = db.Scan(&response).Error
	if err != nil {
		p.logger.Zap.Error("Error scan column")
		return response, err
	}

	return response, err
}
