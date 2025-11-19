package mstkriteria

import (
	"database/sql"
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/mstkriteria"
	repository "riskmanagement/repository/mstkriteria"

	"gitlab.com/golang-package-library/logger"
)

type MstKriteriaDefinition interface {
	GetAll(request models.FilterRequest) (responses []models.MstKriteriaResponse, err error)
	GetAllWithPaginate(request models.FilterRequest) (responses []models.MstKriteriaResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.MstKriteriaResponse, err error)
	GetKodeCriteria() (responses []models.KodeMstKriteria, err error)
	Store(request *models.MstKriteriaRequest) (status bool, msg string, err error)
	Update(request *models.MstKriteriaRequest) (status bool, err error)
	Delete(id int64) (err error)

	// add by panji 18/12/2024
	GetCriteriaById(request models.CriteriaRequestById) (responses []models.MstKriteriaResponse, err error)
	GetCriteriaByPeriode(request models.PeriodeRequest) (responses []models.MstKriteriaHistoryResponses, err error)
}

type MstKriteriaService struct {
	db         lib.Database
	dbRaw      lib.Databases
	logger     logger.Logger
	repository repository.MstKriteriaDefinition
}

func NewMstKriteriaService(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
	repository repository.MstKriteriaDefinition,
) MstKriteriaDefinition {
	return MstKriteriaService{
		db:         db,
		dbRaw:      dbRaw,
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements MstKriteriaDefinition
func (mstKriteria MstKriteriaService) GetAll(request models.FilterRequest) (responses []models.MstKriteriaResponse, err error) {
	return mstKriteria.repository.GetAll(request)
}

// GetAllWithPaginate implements MstKriteriaDefinition
func (mstKriteria MstKriteriaService) GetAllWithPaginate(request models.FilterRequest) (responses []models.MstKriteriaResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalRows, totalData, err := mstKriteria.repository.GetAllWithPaginate(&request)

	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		mstKriteria.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		responses = append(responses, models.MstKriteriaResponse{
			ID:           response.ID,
			KodeCriteria: response.KodeCriteria,
			Criteria:     response.Criteria,
			Restruck:     response.Restruck,
			Status:       response.Status,
			ActiveDate:   response.ActiveDate,
			InactiveDate: response.InactiveDate,
			CreatedAt:    response.CreatedAt,
			CreatedBy:    response.CreatedBy,
			CreatedDesc:  response.CreatedDesc,
			EnabledDate:  response.EnabledDate,
			EnabledBy:    response.EnabledBy,
			EnabledDesc:  response.EnabledDesc,
			DisabledDate: response.DisabledDate,
			DisabledBy:   response.DisabledBy,
			DisabledDesc: response.DisabledDesc,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements MstKriteriaDefinition
func (mstKriteria MstKriteriaService) GetOne(id int64) (responses models.MstKriteriaResponse, err error) {
	return mstKriteria.repository.GetOne(id)
}

// Store implements MstKriteriaDefinition
func (mstKriteria MstKriteriaService) Store(request *models.MstKriteriaRequest) (status bool, msg string, err error) {
	checkId, err := mstKriteria.dbRaw.DB.Query("SELECT COUNT(*) FROM mst_kriteria WHERE kode_criteria = ?", request.KodeCriteria)
	if err != nil {
		return false, msg, err
	}

	rowsCheck, err := mstKriteria.dbRaw.DB.Query("SELECT COUNT(*) FROM mst_kriteria WHERE criteria = ?", request.Criteria)
	if err != nil {
		return false, msg, err
	}

	if checkCount(checkId) > 0 {
		msg = "Duplicate Code : '" + request.KodeCriteria + "'"
		return false, msg, err
	} else if checkCount(rowsCheck) > 0 {
		msg = "Duplicate Criteria : '" + request.Criteria + "'"
		return false, msg, err
	}

	timeNow := lib.GetTimeNow("timestime")
	tx := mstKriteria.db.DB.Begin()

	result, err := mstKriteria.repository.Store(&models.MstKriteria{
		KodeCriteria: request.KodeCriteria,
		Criteria:     request.Criteria,
		Restruck:     request.Restruck,
		Status:       request.Status,
		ActiveDate:   request.ActiveDate,
		InactiveDate: request.InactiveDate,
		CreatedAt:    &timeNow,
		CreatedBy:    request.CreatedBy,
		CreatedDesc:  request.CreatedDesc,
	}, tx)

	if err != nil {
		tx.Rollback()
		msg = "Input data gagal"
		return false, msg, err
	}

	// fmt.Println(result)

	hasil, err := mstKriteria.repository.StoreHistory(&models.MstKriteriaHistory{
		IdCriteria:   result.ID,
		KodeCriteria: request.KodeCriteria,
		Criteria:     request.Criteria,
		Restruck:     request.Restruck,
		Status:       request.Status,
		CreatedAt:    &timeNow,
		CreatedBy:    &request.CreatedBy,
		CreatedDesc:  &request.CreatedDesc,
	}, tx)

	if !hasil && err != nil {
		tx.Rollback()
		fmt.Println("Error :", err)
		msg = "Input data gagal"
		return false, msg, err
	}

	tx.Commit()
	msg = "Input data berhasil"
	return true, msg, err
}

// Update implements MstKriteriaDefinition
func (mstKriteria MstKriteriaService) Update(request *models.MstKriteriaRequest) (status bool, err error) {

	rowsCheck, err := mstKriteria.dbRaw.DB.Query("SELECT COUNT(*) FROM mst_kriteria WHERE id != ? AND criteria = ?", request.ID, request.Criteria)
	if err != nil {
		return false, err
	}

	timeNow := lib.GetTimeNow("timestime")
	tx := mstKriteria.db.DB.Begin()

	if checkCount(rowsCheck) < 1 {
		status, err := mstKriteria.repository.Update(request, tx)
		if !status || err != nil {
			tx.Rollback()
			return status, err
		}

		if request.Status == 1 {
			hasil, err := mstKriteria.repository.StoreHistory(&models.MstKriteriaHistory{
				IdCriteria:   request.ID,
				KodeCriteria: request.KodeCriteria,
				Criteria:     request.Criteria,
				Restruck:     request.Restruck,
				Status:       request.Status,
				ActiveDate:   request.ActiveDate,
				CreatedAt:    &timeNow,
				CreatedBy:    &request.EnabledBy,
				CreatedDesc:  &request.EnabledDesc,
			}, tx)

			if !hasil && err != nil {
				tx.Rollback()
				return false, err
			}
		}

		tx.Commit()
		return true, err
	}

	return false, nil
}

// Delete implements MstKriteriaDefinition
func (mstKriteria MstKriteriaService) Delete(id int64) (err error) {
	return mstKriteria.repository.Delete(id)
}

// GetKodeSubActivity implements SubActivityDefinition
func (mstKriteria MstKriteriaService) GetKodeCriteria() (responses []models.KodeMstKriteria, err error) {
	dataSubs, err := mstKriteria.repository.GetKodeCriteria()
	if err != nil {
		mstKriteria.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataSubs {
		responses = append(responses, models.KodeMstKriteria{
			KodeCriteria: response.KodeCriteria,
		})
	}

	return responses, err
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetCriteriaById implements MstKriteriaDefinition.
func (mstKriteria MstKriteriaService) GetCriteriaById(request models.CriteriaRequestById) (responses []models.MstKriteriaResponse, err error) {
	return mstKriteria.repository.GetCriteriaById(request)
}

// GetCriteriaByPeriode implements MstKriteriaDefinition.
func (mstKriteria MstKriteriaService) GetCriteriaByPeriode(request models.PeriodeRequest) (responses []models.MstKriteriaHistoryResponses, err error) {
	return mstKriteria.repository.GetCriteriaByPeriode(request)
}
