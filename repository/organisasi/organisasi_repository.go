package organisasi

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/organisasi"
	"strings"

	"gitlab.com/golang-package-library/logger"
)

type OrganisasiDefinition interface {
	GetCostCenter(request models.CostCenterRequest) (responses []models.CostCenterResponse, err error)
	GetOrgUnit(request models.DepartmentRequest) (responses []models.DepartmentResponse, err error)
	GetHilfm(request models.JabatanRequest) (responses []models.JabatanResponse, err error)
}

type OrganisasiRepository struct {
	db     lib.Database
	logger logger.Logger
}

// GetJabatan implements OrganisasiDefinition.
func NewOrganisasiRepository(
	db lib.Database,
	logger logger.Logger,
) OrganisasiDefinition {
	return OrganisasiRepository{
		db:     db,
		logger: logger,
	}
}

// getCostCenter implements OrganisasiDefinition.
func (o OrganisasiRepository) GetCostCenter(request models.CostCenterRequest) (responses []models.CostCenterResponse, err error) {
	werksID, err := lib.GetVarEnv("werksID")
	if err != nil {
		return nil, fmt.Errorf("error getting WerksID: %w", err)
	}

	werksIds := strings.Split(werksID, ",")

	var ors []string
	var vals []interface{}
	for _, s := range werksIds {
		s = strings.TrimSpace(s)
		if s != "" {
			ors = append(ors, "werks LIKE ?")
			vals = append(vals, s+"%")
		}
	}

	db := o.db.DB.Table(`mst_werks`).
		Where(strings.Join(ors, " OR "), vals...)

	err = db.Scan(&responses).Error

	if err != nil {
		return nil, fmt.Errorf("error getting data: %w", err)
	}

	return responses, err
}

// GetOrgUnit implements OrganisasiDefinition.
func (o OrganisasiRepository) GetOrgUnit(request models.DepartmentRequest) (responses []models.DepartmentResponse, err error) {
	db := o.db.DB.Table("mst_dept").Limit(request.Limit).Offset(request.Offset)

	if request.Type != "" {
		db = db.Where("tipe_uker = ?", request.Type)
	}

	if request.Werks != "" {
		db = db.Where("werks = ?", request.Werks)
	}

	// Keyword := fmt.Sprintf("%%%s%%", request.Keyword)
	if request.Keyword != "" {
		// Kalau keyword mengandung spasi panjang → anggap sebagai full phrase search
		if strings.Contains(request.Keyword, " ") {
			// Gunakan fulltext search (IN BOOLEAN MODE + kutip)
			phrase := fmt.Sprintf("\"%s\"", request.Keyword)
			db = db.Where("MATCH(orgeh, orgeh_tx) AGAINST(? IN BOOLEAN MODE)", phrase)
		} else {
			// Kalau keyword cuma 1 kata → coba LIKE untuk lebih fleksibel
			like := fmt.Sprintf("%%%s%%", request.Keyword)
			db = db.Where("orgeh LIKE ? OR orgeh_tx LIKE ?", like, like)
		}
	}

	err = db.Scan(&responses).Error

	if err != nil {
		return nil, fmt.Errorf("error getting data: %w", err)
	}

	return responses, err
}

func (o OrganisasiRepository) GetHilfm(request models.JabatanRequest) (responses []models.JabatanResponse, err error) {
	db := o.db.DB.Table("mst_hilfm").Limit(request.Limit).Offset(request.Offset)

	// Keyword := fmt.Sprintf("%%%s%%", request.Keyword)
	if request.Keyword != "" {
		if strings.Contains(request.Keyword, " ") {
			// Gunakan fulltext search (IN BOOLEAN MODE + kutip)
			phrase := fmt.Sprintf("\"%s\"", request.Keyword)
			db = db.Where("MATCH(hilfm, htext) AGAINST(? IN BOOLEAN MODE)", phrase)
		} else {
			// Kalau keyword cuma 1 kata → coba LIKE untuk lebih fleksibel
			like := fmt.Sprintf("%%%s%%", request.Keyword)
			db = db.Where("hilfm LIKE ? OR htext LIKE ?", like, like)
		}
	}

	err = db.Scan(&responses).Error

	if err != nil {
		return nil, fmt.Errorf("error getting data: %w", err)
	}

	return responses, err
}
