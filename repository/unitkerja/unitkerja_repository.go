package unitkerja

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/unitkerja"
	"strings"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type UnitKerjaDefinition interface {
	GetAll() (responses []models.UnitKerjaResponse, err error)
	GetOne(id int64) (responses models.UnitKerjaResponse, err error)
	Store(request *models.UnitKerjaRequest) (responses bool, err error)
	Update(request *models.UnitKerjaRequest) (responses bool, err error)
	Delete(id int64) (err error)
	WithTrx(trxHandle *gorm.DB) UnitKerjaRepository
	GetRegionList(request *models.RegionRequest) (responses []models.RegionList, err error)
	GetMainbrList(request *models.MainbrRequest) (responses []models.MainbrList, err error)
	GetBranchList(request *models.BranchRequest) (responses []models.BranchList, err error)

	GetRegionName(REGION string) (name models.UkerName, err error)
	GetMainbrName(MAINBR string) (name models.UkerName, err error)
	GetBranchName(BRANCH string) (name models.UkerName, err error)
	GetMainbrKWList(request *models.MainbrKWRequest) (responses []models.MainbrList, err error)
	GetEmployeeRegion(request *models.EmployeeRegionRequest) (responses models.EmployeeRegionResponse, err error)

	CekPgsActive(PERNR string, BRANCH string) (count int, err error)

	// Batch3
	// CekIsBRC(PERNR string) (response bool, err error)
	CekIsBRC(HILFM string) (response bool, err error)
	GetDetailUker(BRANCH string) (response models.DetailUkerResponse, err error)

	// For Disaster Map
	GetMapRegionList(request *models.MapLocationRequest) (response []models.MapRegionOffice, err error)
	GetMapBranchList(request *models.MapLocationRequest) (response []models.MapBranchOffice, err error)
	GetMapUnitList(request *models.MapLocationRequest) (response []models.MapUnitOffice, err error)
}

type UnitKerjaRepository struct {
	db      lib.Database
	dbRaw   lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewUnitKerjaRepository(db lib.Database, dbRaw lib.Database, logger logger.Logger) UnitKerjaDefinition {
	return UnitKerjaRepository{
		db:      db,
		dbRaw:   dbRaw,
		logger:  logger,
		timeout: time.Second * 100,
	}
}

// Delete implements ActicityDefinition
func (unitKerja UnitKerjaRepository) Delete(id int64) (err error) {
	return unitKerja.db.DB.Where("id = ?", id).Delete(&models.UnitKerjaResponse{}).Error
}

// GetAll implements ActicityDefinition
func (unitKerja UnitKerjaRepository) GetAll() (responses []models.UnitKerjaResponse, err error) {
	return responses, unitKerja.db.DB.Find(&responses).Error
}

// GetOne implements ActicityDefinition
func (unitKerja UnitKerjaRepository) GetOne(id int64) (responses models.UnitKerjaResponse, err error) {
	return responses, unitKerja.db.DB.Where("id = ?", id).Find(&responses).Error
}

// Store implements ActicityDefinition
func (unitKerja UnitKerjaRepository) Store(request *models.UnitKerjaRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	fmt.Println("repo = ", models.UnitKerjaRequest{
		KodeUker:   request.KodeUker,
		NamaUker:   request.NamaUker,
		KodeCabang: request.KodeCabang,
		NamaCabang: request.NamaCabang,
		KanwilID:   request.KanwilID,
		KodeKanwil: request.KodeKanwil,
		Kanwil:     request.Kanwil,
		Status:     request.Status,
		CreatedAt:  &timeNow,
	})
	err = unitKerja.db.DB.Save(&models.UnitKerjaRequest{
		KodeUker:   request.KodeUker,
		NamaUker:   request.NamaUker,
		KodeCabang: request.KodeCabang,
		NamaCabang: request.NamaCabang,
		KanwilID:   request.KanwilID,
		KodeKanwil: request.KodeKanwil,
		Kanwil:     request.Kanwil,
		Status:     request.Status,
		CreatedAt:  &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// Update implements ActicityDefinition
func (unitKerja UnitKerjaRepository) Update(request *models.UnitKerjaRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	err = unitKerja.db.DB.Save(&models.UnitKerjaRequest{
		ID:         request.ID,
		KodeUker:   request.KodeUker,
		NamaUker:   request.NamaUker,
		KodeCabang: request.KodeCabang,
		NamaCabang: request.NamaCabang,
		KanwilID:   request.KanwilID,
		KodeKanwil: request.KodeKanwil,
		Kanwil:     request.Kanwil,
		Status:     request.Status,
		CreatedAt:  request.CreatedAt,
		UpdatedAt:  &timeNow,
	}).Error
	if err != nil {
		return false, err
	}

	return true, err
}

// WithTrx implements ActicityDefinition
func (unitKerja UnitKerjaRepository) WithTrx(trxHandle *gorm.DB) UnitKerjaRepository {
	if trxHandle == nil {
		unitKerja.logger.Zap.Error("transaction Database not found in gin context")
		return unitKerja
	}
	unitKerja.db.DB = trxHandle
	return unitKerja
}

func (unitKerja UnitKerjaRepository) GetRegionList(request *models.RegionRequest) (responses []models.RegionList, err error) {
	fmt.Println("=========== masuk")

	db := unitKerja.db.DB

	if request.WERKS != "" {
		db = db.Table("uker_kelolaan_user").
			Select(`REGION, RGDESC 'BRDESC'`).
			Where(`pn = ?`, request.PERNR).
			Group(`REGION`)

		err = db.Scan(&responses).Error
	} else {
		db = db.Table("dwh_branch").
			Select(`
				REGION,
				CASE 
					WHEN REGION = 'Y' THEN RGDESC 
					WHEN REGION = 'Z' THEN RGDESC 
					ELSE BRDESC 
				END 'BRDESC'
			`).
			Where("BRDESC LIKE 'KANWIL%' OR REGION in ('Y', 'Z')")

		if request.TipeUker != "KP" && request.TipeUker != "" {
			REGION := unitKerja.CheckRegion(request.BRANCH)
			// fmt.Println("Region ====>", REGION)
			db = db.Where("REGION = ?", REGION)
		}

		err = db.Group("REGION").
			Scan(&responses).Error
	}

	// fmt.Println("=========== error", err)

	if err != nil {
		return responses, err
	}

	return responses, err
}

func (unitKerja UnitKerjaRepository) GetMainbrList(request *models.MainbrRequest) (responses []models.MainbrList, err error) {
	db := unitKerja.db.DB

	isBRC := request.FlagBrc

	if isBRC || request.WERKS == "LN00" {
		db = db.Table("uker_kelolaan_user").
			Select("MAINBR, MBDESC as 'BRDESC'").
			Where("REGION = ?", request.REGION).
			Where("pn = ?", request.PERNR).
			Group("MAINBR")

		err = db.Scan(&responses).Error

		if err != nil {
			return responses, err
		}

		ada, err := unitKerja.CekPgsActive(request.PERNR, "")
		if err != nil {
			return responses, err
		}

		if ada > 0 {
			fmt.Println("ada")
			var additionalResponses []models.MainbrList

			pgsQuery := unitKerja.db.DB.Table("pgs_user").
				Select(`MAINBR, MBDESC as 'BRDESC'`).
				Where(`pn = ?`, request.PERNR).
				Where(`delete_flag = 0`).
				Where(`status IN ("02a")`).
				Where(`action = 'Active'`)

			pgsQuery.Group("MAINBR").Find(&additionalResponses)

			responses = append(responses, additionalResponses...)
		}
	} else {
		db = db.Table("dwh_branch").
			Select("MAINBR, BRDESC").
			Where("REGION = ?", request.REGION).
			Where("MBDESC = BRDESC").
			Group("MAINBR")
		if request.TipeUker == "KC" {
			db.Where("MAINBR = BRANCH").Where("BRANCH = CAST(? as float)", request.BRANCH)
		}

		if request.TipeUker == "UN" {
			db.Where("MAINBR = ?", request.MAINBR)
		}

		err = db.Scan(&responses).Error

		if err != nil {
			return responses, err
		}

		ada, err := unitKerja.CekPgsActive(request.PERNR, "")
		if err != nil {
			return responses, err
		}

		if ada > 0 {
			fmt.Println("ada")
			var additionalResponses []models.MainbrList

			pgsQuery := unitKerja.db.DB.Table("pgs_user").
				Select(`MAINBR, MBDESC as 'BRDESC'`).
				Where(`pn = ?`, request.PERNR).
				Where(`delete_flag = 0`).
				Where(`status IN ("02a")`).
				Where(`action = 'Active'`)

			pgsQuery.Group("MAINBR").Find(&additionalResponses)

			responses = append(responses, additionalResponses...)
		}
	}

	return responses, nil
}

func (unitKerja UnitKerjaRepository) GetBranchList(request *models.BranchRequest) (responses []models.BranchList, err error) {
	db := unitKerja.db.DB

	isBrc := request.FlagBrc

	if isBrc || request.WERKS == "LN00" {
		mainbrs := strings.Split(request.MAINBR, ",")

		db = db.Table("uker_kelolaan_user").
			Select("BRANCH, BRDESC").
			Where("REGION = ?", request.REGION).
			Where("MAINBR in (?)", mainbrs).
			Where("pn = ?", request.PERNR)

		err = db.Find(&responses).Error

		if err != nil {
			return responses, err
		}

		ada, err := unitKerja.CekPgsActive(request.PERNR, "")
		if err != nil {
			return responses, err
		}

		fmt.Println("PGS =>", ada)

		if ada > 0 {
			fmt.Println("ada")
			var additionalResponses []models.BranchList

			pgsQuery := unitKerja.db.DB.Table("pgs_user").
				Select(`BRANCH, BRDESC`).
				Where(`pn = ?`, request.PERNR).
				Where(`delete_flag = 0`).
				Where(`status IN ("02a")`).
				Where(`action = 'Active'`).
				Where("REGION = ?", request.REGION).
				Where("MAINBR = ?", request.MAINBR)

			pgsQuery.Group("MAINBR").Find(&additionalResponses)

			responses = append(responses, additionalResponses...)
		}
	} else {
		db = db.Table("dwh_branch").
			Select("BRANCH, BRDESC").
			Where("REGION = ?", request.REGION).
			Where("MAINBR = ?", request.MAINBR)

		if request.TipeUker == "UN" {
			db.Where("BRANCH = CAST(? as float)", request.BRANCH)
		}

		err = db.Scan(&responses).Error

		if err != nil {
			return responses, err
		}

		ada, err := unitKerja.CekPgsActive(request.PERNR, "")
		if err != nil {
			return responses, err
		}

		fmt.Println("PGS =>", ada)

		if ada > 0 {
			fmt.Println("ada")
			var additionalResponses []models.BranchList

			pgsQuery := unitKerja.db.DB.Table("pgs_user").
				Select(`BRANCH, BRDESC`).
				Where(`pn = ?`, request.PERNR).
				Where(`delete_flag = 0`).
				Where(`status IN ("02a")`).
				Where(`action = 'Active'`).
				Where("REGION = ?", request.REGION).
				Where("MAINBR = ?", request.MAINBR)

			pgsQuery.Group("MAINBR").Find(&additionalResponses)

			responses = append(responses, additionalResponses...)
		}
	}

	return responses, nil
}

func (unitKerja UnitKerjaRepository) GetRegionName(REGION string) (name models.UkerName, err error) {
	db := unitKerja.db.DB

	err = db.Table("dwh_branch").
		Select("BRDESC").
		Where("REGION = ?", REGION).
		Where("BRDESC LIKE 'KANWIL%'").
		Where("BRDESC NOT LIKE ?", "KANWIL Banda Aceh").
		First(&name).Error

	if err != nil {
		return name, err
	}

	return name, err
}

func (unitKerja UnitKerjaRepository) GetMainbrName(MAINBR string) (name models.UkerName, err error) {
	db := unitKerja.db.DB

	err = db.Table("dwh_branch").
		Select("BRDESC").
		Where("MAINBR = ?", MAINBR).
		Where("MBDESC = BRDESC").
		Where("BRDESC LIKE 'kc%' OR 'kanca%'").
		First(&name).Error

	if err != nil {
		return name, err
	}

	return name, err
}

func (unitKerja UnitKerjaRepository) GetBranchName(BRANCH string) (name models.UkerName, err error) {
	db := unitKerja.db.DB

	branches := strings.Split(BRANCH, ",")

	err = db.Table("dwh_branch").
		Select("BRDESC").
		Where("BRANCH in (?)", branches).
		First(&name).Error

	if err != nil {
		return name, err
	}

	return name, err
}

func (unitKerja UnitKerjaRepository) CheckRegion(REGION string) (response string) {
	db := unitKerja.db.DB

	db = db.Table("dwh_branch").Select("REGION 'response'").Where(`BRANCH = CAST(? as float)`, REGION).Find(&response)

	return response
}

func (unitKerja UnitKerjaRepository) GetMainbrKWList(request *models.MainbrKWRequest) (responses []models.MainbrList, err error) {
	db := unitKerja.db.DB

	err = db.Table("dwh_branch").
		Select("MAINBR, BRDESC").
		Where("REGION = ?", request.REGION).
		Where("MBDESC = BRDESC").
		Group("MAINBR").
		Scan(&responses).Error

	if err != nil {
		return responses, err
	}

	return responses, err
}

func (unitKerja UnitKerjaRepository) GetEmployeeRegion(request *models.EmployeeRegionRequest) (responses models.EmployeeRegionResponse, err error) {
	return responses, unitKerja.db.DB.Where("BRANCH = ?", request.Branch).Find(&responses).Error
}

// versioning 24/10/2023 by panji
// CekPgsActive implements PgsUserDefinition
func (unitKerja UnitKerjaRepository) CekPgsActive(PERNR string, BRANCH string) (count int, err error) {
	db := unitKerja.db.DB

	db = db.Table("pgs_user").
		Select(`COUNT(*) 'count'`).
		Where(`delete_flag = 0`).
		Where(`status != '03a'`)
		// Where(`status IN ("02a","01a","01b")`).
		// Where(`action IN ("Active", "New Request", "UpdateData")`)

	if BRANCH != "" {
		db = db.Where(`pn = ? OR BRANCH = ?`, PERNR, BRANCH)
	} else {
		db = db.Where(`pn = ?`, PERNR)
	}

	err = db.Scan(&count).Error

	if err != nil {
		return count, err
	}

	return count, err
}

func (uk UnitKerjaRepository) CekIsBRC(HILFM string) (response bool, err error) {
	var Count int64
	db := uk.db.DB.Table("mst_parameter_search_brc").
		Where("params_value = ?", HILFM)

	err = db.Count(&Count).Error

	if err != nil {
		return false, err
	}

	if Count < 1 {
		return false, nil
	}

	return true, nil
}

// GetDetailUker implements UnitKerjaDefinition.
func (uk UnitKerjaRepository) GetDetailUker(BRANCH string) (responses models.DetailUkerResponse, err error) {
	db := uk.db.DB.Table("dwh_branch").
		Select(`REGION,RGDESC, MAINBR, MBDESC ,BRANCH, BRDESC`).Where(`BRANCH = CAST(? as float)`, BRANCH)

	err = db.Scan(&responses).Error

	if err != nil {
		return responses, err
	}

	return responses, nil
}

// GetMapRegionList implements UnitKerjaDefinition.
func (uk UnitKerjaRepository) GetMapRegionList(request *models.MapLocationRequest) (response []models.MapRegionOffice, err error) {
	db := uk.db.DB

	db = db.Table("dwh_branch db").
		Select(`
				db.REGION 'region',
				db.BRANCH 'region_code',
				UPPER(db.BRDESC)'region_name',
				dbl.longitude 'longitude',
				dbl.latitude 'latitude',
				dbl.address 'address'`).
		Joins(`INNER JOIN dwh_branch_location dbl ON db.BRANCH = dbl.branch`).
		Where("BRDESC LIKE 'KANWIL%'")

	if request.LevelUker != "KP" {
		if request.KodeRegion != "" {
			db = db.Where("db.REGION = ?", request.KodeRegion)
		} else {
			return nil, fmt.Errorf("field kode_region is required")
		}
	}

	if request.Keyword != "" {
		like := fmt.Sprintf("%%%s%%", strings.ToUpper(request.Keyword))
		db = db.Where("UPPER(db.BRDESC) LIKE ?", like)
	}

	err = db.Group("db.REGION").Scan(&response).Error

	if err != nil {
		return response, err
	}

	return response, nil
}

// GetMapBranchList implements UnitKerjaDefinition.
func (unitKerja UnitKerjaRepository) GetMapBranchList(request *models.MapLocationRequest) (response []models.MapBranchOffice, err error) {
	db := unitKerja.db.DB

	db = db.Table("dwh_branch db").
		Select(`
				db.BRUNIT,
				db.MAINBR 'branch',
				db.BRANCH 'branch_code',
				UPPER(db.BRDESC)'branch_name',
				dbl.longitude 'longitude',
				dbl.latitude 'latitude',
				dbl.address 'address'`).
		Joins(`INNER JOIN dwh_branch_location dbl ON db.BRANCH = dbl.branch`).
		Where(`db.BRANCH = db.MAINBR`).
		Where(`db.BRUNIT = 'B'`)

	if request.LevelUker != "KP" {
		if request.KodeRegion != "" {
			db = db.Where("db.REGION = ?", request.KodeRegion)

			if request.KodeBranch != "" {
				db = db.Where("db.MAINBR = ?", request.KodeBranch)
			}
		} else {
			return nil, fmt.Errorf("field kode_region is required")
		}
	}

	if request.Keyword != "" {
		like := fmt.Sprintf("%%%s%%", strings.ToUpper(request.Keyword))
		db = db.Where("CONCAT(db.BRANCH, ' - ', UPPER(db.BRDESC)) LIKE ?", like)
	}

	err = db.Scan(&response).Error

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetMapUnitList implements UnitKerjaDefinition.
func (unitKerja UnitKerjaRepository) GetMapUnitList(request *models.MapLocationRequest) (response []models.MapUnitOffice, err error) {
	db := unitKerja.db.DB

	db = db.Table("dwh_branch db").
		Select(`
				db.BRANCH 'unit',
				db.BRANCH 'unit_code',
				UPPER(db.BRDESC)'unit_name',
				dbl.longitude 'longitude',
				dbl.latitude 'latitude',
				dbl.address 'address'`).
		Joins(`INNER JOIN dwh_branch_location dbl ON db.BRANCH = dbl.branch`)

	if request.FlagAll {
		db = db.Where(`db.BRUNIT in ('K','S','U','B')`)
	} else {
		db = db.Where(`db.BRUNIT in ('K','S','U')`)
	}

	if request.LevelUker != "KP" {
		if request.KodeRegion != "" {
			db = db.Where("db.REGION = ?", request.KodeRegion)
		} else {
			return nil, fmt.Errorf("field kode_region is required")
		}

		if request.KodeBranch != "" {
			db = db.Where("db.MAINBR = ?", request.KodeBranch)
		}
	}

	if request.Keyword != "" {
		like := fmt.Sprintf("%%%s%%", strings.ToUpper(request.Keyword))
		db = db.Where("CONCAT(db.BRANCH, ' - ', UPPER(db.BRDESC)) LIKE ?", like)
	}

	err = db.Order("db.BRANCH ASC").Scan(&response).Error

	if err != nil {
		return nil, err
	}

	return response, err
}
