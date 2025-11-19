package riskissue

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/riskissue"
	riskissue "riskmanagement/repository/riskissue"

	"github.com/google/uuid"

	"gitlab.com/golang-package-library/logger"
)

var (
	UUID = uuid.NewString()
)

type RiskIssueDefinition interface {
	GetAll() (responses []models.RiskIssueResponse, err error)
	GetOne(id int64) (responses models.RiskIssueResponseGetOne, status bool, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.RiskIssueResponse, pagination lib.Pagination, err error)
	Store(request models.RiskIssueRequest) (responses bool, err error)
	Update(request *models.RiskIssueRequest) (status bool, err error)
	DeleteMapProses(request *models.MapProses) (status bool, err error)
	DeleteMapEvent(request *models.MapEvent) (status bool, err error)
	DeleteMapKejadian(request *models.MapKejadian) (status bool, err error)
	DeleteMapLiniBisnis(request *models.MapLiniBisnis) (status bool, err error)
	DeleteMapProduct(request *models.MapProduct) (status bool, err error)
	DeleteMapAktifitas(request *models.MapAktifitas) (status bool, err error)
	DeleteMapControl(request *models.MapControl) (status bool, err error)
	DeleteMapIndicator(request *models.MapIndicator) (status bool, err error)
	GetKode() (responses []models.KodeRespon, err error)
	MappingRiskControl(request models.MappingControlRequest) (responses bool, err error)
	GetMappingControlbyID(id int64) (responses models.RiskIssueResponseGetOne, err error)
	SearchRiskIssue(requests models.KeywordRequest) (responses []models.RiskIssueResponses, pagination lib.Pagination, err error)
	SearchRiskIssueWithoutSub(request models.RiskIssueWithoutSub) (responses []models.RiskIssueResponses, pagination lib.Pagination, err error)
	MappingRiskIndicator(request models.MappingIndicatorRequest) (responses bool, err error)
	GetMappingIndicatorbyID(id int64) (responses models.RiskIssueResponseGetOne, err error)
	Delete(request *models.RiskIssueDeleteRequest) (responses bool, err error)
	FilterRiskIssue(request models.FilterRiskIssueRequest) (responses []models.RiskIssueFilterResponses, pagination lib.Pagination, err error)
	GetRiskIssueByActivity(id int64) (responses []models.RiskIssueResponseByActivity, err error)
	GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateri, err error)
	GetMateriByCode(request models.RiskIssueCode) (responses []models.ListMateri, err error)
	GetRiskIssueByActivityID(id int64) (responses []models.RiskIssueResponseByActivity, err error)
	GetRiskEventName(id int64) (name string, err error)
}

type RiskIssueService struct {
	db            lib.Database
	dbRaw         lib.Databases
	logger        logger.Logger
	riskissueRepo riskissue.RiskIssueDefinition
	mapAktifitas  riskissue.MapAktifitasDefinition
	mapEvent      riskissue.MapEventDefinition
	mapLiniBisnis riskissue.MapLiniBisnisDefinition
	mapKejadian   riskissue.MapKejadianDefinition
	mapProduct    riskissue.MapProductDefinition
	mapProses     riskissue.MapProsesDefinition
	mapControl    riskissue.MapControlDefinition
	mapIndicator  riskissue.MapIndicatorDefinition
}

func NewRiskIssueService(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
	riskissueRepo riskissue.RiskIssueDefinition,
	mapAktifitas riskissue.MapAktifitasDefinition,
	mapEvent riskissue.MapEventDefinition,
	mapLiniBisnis riskissue.MapLiniBisnisDefinition,
	mapKejadian riskissue.MapKejadianDefinition,
	mapProduct riskissue.MapProductDefinition,
	mapProses riskissue.MapProsesDefinition,
	mapControl riskissue.MapControlDefinition,
	mapIndicator riskissue.MapIndicatorDefinition,
) RiskIssueDefinition {
	return RiskIssueService{
		db:            db,
		dbRaw:         dbRaw,
		logger:        logger,
		riskissueRepo: riskissueRepo,
		mapAktifitas:  mapAktifitas,
		mapEvent:      mapEvent,
		mapLiniBisnis: mapLiniBisnis,
		mapKejadian:   mapKejadian,
		mapProduct:    mapProduct,
		mapProses:     mapProses,
		mapControl:    mapControl,
		mapIndicator:  mapIndicator,
	}
}

// GetAll implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetAll() (responses []models.RiskIssueResponse, err error) {
	return riskIssue.riskissueRepo.GetAll()
}

// GetAllWithPaginate implements RiskIssueDefinition
func (ri RiskIssueService) GetAllWithPaginate(request models.Paginate) (responses []models.RiskIssueResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPgs, totalRows, totalData, err := ri.riskissueRepo.GetAllWithPaginate(&request)
	if err != nil {
		ri.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPgs {
		responses = append(responses, models.RiskIssueResponse{
			ID:             response.ID,
			RiskTypeID:     response.RiskTypeID,
			RiskIssueCode:  response.RiskIssueCode,
			RiskIssue:      response.RiskIssue,
			Deskripsi:      response.Deskripsi,
			KategoriRisiko: response.KategoriRisiko,
			Status:         response.Status,
			Likelihood:     response.Likelihood,
			Impact:         response.Impact,
			DeleteFlag:     response.DeleteFlag,
			CreatedAt:      response.CreatedAt,
			UpdatedAt:      response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetOne(id int64) (responses models.RiskIssueResponseGetOne, status bool, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.GetOne(id)
	if dataRiskIssue.ID != 0 {
		fmt.Println("bukan 0")

		data_aktifitas, err := riskIssue.mapAktifitas.GetOneDataByID(dataRiskIssue.ID)
		data_event, err := riskIssue.mapEvent.GetOneDataByID(dataRiskIssue.ID)
		data_lini_bisnis, err := riskIssue.mapLiniBisnis.GetOneDataByID(dataRiskIssue.ID)
		data_kejadian, err := riskIssue.mapKejadian.GetOneDataByID(dataRiskIssue.ID)
		data_product, err := riskIssue.mapProduct.GetOneDataByID(dataRiskIssue.ID)
		data_proses, err := riskIssue.mapProses.GetOneDataByID(dataRiskIssue.ID)

		fmt.Println(data_proses)

		responses = models.RiskIssueResponseGetOne{
			ID:             dataRiskIssue.ID,
			RiskTypeID:     dataRiskIssue.RiskTypeID,
			RiskIssueCode:  dataRiskIssue.RiskIssueCode,
			RiskIssue:      dataRiskIssue.RiskIssue,
			Deskripsi:      dataRiskIssue.Deskripsi,
			KategoriRisiko: dataRiskIssue.KategoriRisiko,
			Status:         dataRiskIssue.Status,
			Likelihood:     dataRiskIssue.Likelihood,
			Impact:         dataRiskIssue.Impact,
			MapProses:      data_proses,
			MapEvent:       data_event,
			MapKejadian:    data_kejadian,
			MapProduct:     data_product,
			MapLiniBisnis:  data_lini_bisnis,
			MapAktifitas:   data_aktifitas,
			CreatedAt:      dataRiskIssue.CreatedAt,
			UpdatedAt:      dataRiskIssue.UpdatedAt,
		}

		return responses, true, err
	}

	return responses, false, err
}

// Store implements RiskIssueDefinition
func (riskIssue RiskIssueService) Store(request models.RiskIssueRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := riskIssue.db.DB.Begin()

	reqRiskIssue := &models.RiskIssue{
		RiskTypeID:     request.RiskTypeID,
		RiskIssueCode:  request.RiskIssueCode,
		RiskIssue:      request.RiskIssue,
		Deskripsi:      request.Deskripsi,
		KategoriRisiko: request.KategoriRisiko,
		Status:         request.Status,
		Likelihood:     request.Likelihood,
		Impact:         request.Impact,
		DeleteFlag:     false,
		CreatedAt:      &timeNow,
	}

	dataRiskIssue, err := riskIssue.riskissueRepo.Store(reqRiskIssue, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Input Map Proses
	if len(request.MapProses) != 0 {
		for _, value := range request.MapProses {
			_, err = riskIssue.mapProses.Store(&models.MapProses{
				IDRiskIssue:    dataRiskIssue.ID,
				MegaProses:     value.MegaProses,
				MajorProses:    value.MajorProses,
				SubMajorProses: value.SubMajorProses,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Input MapEvent
	if len(request.MapEvent) != 0 {
		for _, value := range request.MapEvent {
			_, err = riskIssue.mapEvent.Store(&models.MapEvent{
				IDRiskIssue:  dataRiskIssue.ID,
				EventTypeLv1: value.EventTypeLv1,
				EventTypeLv2: value.EventTypeLv2,
				EventTypeLv3: value.EventTypeLv3,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Input MapKejadian
	if len(request.MapKejadian) != 0 {
		for _, value := range request.MapKejadian {
			_, err = riskIssue.mapKejadian.Store(&models.MapKejadian{
				IDRiskIssue:         dataRiskIssue.ID,
				PenyebabKejadianLv1: value.PenyebabKejadianLv1,
				PenyebabKejadianLv2: value.PenyebabKejadianLv2,
				PenyebabKejadianLv3: value.PenyebabKejadianLv3,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Input MapProduct
	if len(request.MapProduct) != 0 {
		for _, value := range request.MapProduct {
			_, err = riskIssue.mapProduct.Store(&models.MapProduct{
				IDRiskIssue: dataRiskIssue.ID,
				Product:     value.Product,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//input Map Lini Bisnis
	if len(request.MapLiniBisnis) != 0 {
		for _, value := range request.MapLiniBisnis {
			_, err = riskIssue.mapLiniBisnis.Store(&models.MapLiniBisnis{
				IDRiskIssue:   dataRiskIssue.ID,
				LiniBisnisLv1: value.LiniBisnisLv1,
				LiniBisnisLv2: value.LiniBisnisLv2,
				LiniBisnisLv3: value.LiniBisnisLv3,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//input Map aktifitas
	if len(request.MapAktifitas) != 0 {
		for _, value := range request.MapAktifitas {
			_, err = riskIssue.mapAktifitas.Store(&models.MapAktifitas{
				IDRiskIssue:  dataRiskIssue.ID,
				Aktifitas:    value.Aktifitas,
				SubAktifitas: value.SubAktifitas,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// Update implements RiskIssueDefinition
func (riskIssue RiskIssueService) Update(request *models.RiskIssueRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := riskIssue.db.DB.Begin()

	updateRiskIssue := &models.RiskIssueUpdate{
		ID:             request.ID,
		RiskTypeID:     request.RiskTypeID,
		RiskIssueCode:  request.RiskIssueCode,
		RiskIssue:      request.RiskIssue,
		Deskripsi:      request.Deskripsi,
		KategoriRisiko: request.KategoriRisiko,
		Status:         request.Status,
		Likelihood:     request.Likelihood,
		Impact:         request.Likelihood,
		UpdatedAt:      &timeNow,
	}

	include := []string{
		"id",
		"risk_type_id",
		"risk_issue_code",
		"risk_issue",
		"deskripsi",
		"kategori_risiko",
		"status",
		"likelihood",
		"impact",
		"updated_at",
	}

	_, err = riskIssue.riskissueRepo.Update(updateRiskIssue, include, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Update MapProses
	if len(request.MapProses) != 0 {
		for _, value := range request.MapProses {
			updateProses := &models.MapProses{
				ID:             value.ID,
				IDRiskIssue:    request.ID,
				MegaProses:     value.MegaProses,
				MajorProses:    value.MajorProses,
				SubMajorProses: value.SubMajorProses,
			}

			_, err = riskIssue.mapProses.Update(updateProses, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Update MapEvent
	if len(request.MapEvent) != 0 {
		for _, value := range request.MapEvent {
			updateEvent := &models.MapEvent{
				ID:           value.ID,
				IDRiskIssue:  request.ID,
				EventTypeLv1: value.EventTypeLv1,
				EventTypeLv2: value.EventTypeLv2,
				EventTypeLv3: value.EventTypeLv3,
			}

			_, err = riskIssue.mapEvent.Update(updateEvent, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Update MapKejadian
	if len(request.MapKejadian) != 0 {
		for _, value := range request.MapKejadian {
			updateKejadian := &models.MapKejadian{
				ID:                  value.ID,
				IDRiskIssue:         request.ID,
				PenyebabKejadianLv1: value.PenyebabKejadianLv1,
				PenyebabKejadianLv2: value.PenyebabKejadianLv2,
				PenyebabKejadianLv3: value.PenyebabKejadianLv3,
			}

			_, err = riskIssue.mapKejadian.Update(updateKejadian, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Update MapProduct
	if len(request.MapProduct) != 0 {
		for _, value := range request.MapProduct {
			updateProduct := &models.MapProduct{
				ID:          value.ID,
				IDRiskIssue: request.ID,
				Product:     value.Product,
			}

			_, err = riskIssue.mapProduct.Update(updateProduct, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//Update MapLiniBisnis
	if len(request.MapLiniBisnis) != 0 {
		for _, value := range request.MapLiniBisnis {
			updateLiniBisnis := &models.MapLiniBisnis{
				ID:            value.ID,
				IDRiskIssue:   request.ID,
				LiniBisnisLv1: value.LiniBisnisLv1,
				LiniBisnisLv2: value.LiniBisnisLv2,
				LiniBisnisLv3: value.LiniBisnisLv3,
			}

			_, err = riskIssue.mapLiniBisnis.Update(updateLiniBisnis, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	//MapAktifitas
	if len(request.MapAktifitas) != 0 {
		for _, value := range request.MapAktifitas {
			updateAktifitas := &models.MapAktifitas{
				ID:           value.ID,
				IDRiskIssue:  request.ID,
				Aktifitas:    value.Aktifitas,
				SubAktifitas: value.SubAktifitas,
			}

			_, err = riskIssue.mapAktifitas.Update(updateAktifitas, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapAktifitas implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapAktifitas(request *models.MapAktifitas) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapAktifitas(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapEvent implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapEvent(request *models.MapEvent) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapEvent(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapKejadian implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapKejadian(request *models.MapKejadian) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapKejadian(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapLiniBisnis implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapLiniBisnis(request *models.MapLiniBisnis) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapLiniBisnis(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapProduct implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapProduct(request *models.MapProduct) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapProduct(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapControl implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapControl(request *models.MapControl) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapControl(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapIndicator implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapIndicator(request *models.MapIndicator) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapIndicator(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteMapProses implements RiskIssueDefinition
func (riskIssue RiskIssueService) DeleteMapProses(request *models.MapProses) (status bool, err error) {
	tx := riskIssue.db.DB.Begin()

	err = riskIssue.riskissueRepo.DeleteMapProses(request.ID, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetKode implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetKode() (responses []models.KodeRespon, err error) {
	datariskIssue, err := riskIssue.riskissueRepo.GetKode()
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range datariskIssue {
		responses = append(responses, models.KodeRespon{
			Kode: response.Kode.String,
		})
	}

	return responses, err
}

// MappingRiskControl implements RiskIssueDefinition
func (riskIssue RiskIssueService) MappingRiskControl(request models.MappingControlRequest) (responses bool, err error) {
	tx := riskIssue.db.DB.Begin()

	if len(request.MapControl) != 0 {
		for _, value := range request.MapControl {
			_, err = riskIssue.mapControl.Store(&models.MapControl{
				ID:          value.ID,
				IDRiskIssue: request.ID,
				IDControl:   value.IDControl,
				IsChecked:   value.IsChecked,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetMappingControlbyID implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetMappingControlbyID(id int64) (responses models.RiskIssueResponseGetOne, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.GetOne(id)
	if dataRiskIssue.ID != 0 {
		fmt.Println("bukan 0")
		dataControl, err := riskIssue.mapControl.GetOneDataByID(dataRiskIssue.ID)

		responses = models.RiskIssueResponseGetOne{
			ID:             dataRiskIssue.ID,
			RiskTypeID:     dataRiskIssue.RiskTypeID,
			RiskIssueCode:  dataRiskIssue.RiskIssueCode,
			RiskIssue:      dataRiskIssue.RiskIssue,
			Deskripsi:      dataRiskIssue.Deskripsi,
			KategoriRisiko: dataRiskIssue.KategoriRisiko,
			Status:         dataRiskIssue.Status,
			Likelihood:     dataRiskIssue.Likelihood,
			Impact:         dataRiskIssue.Impact,
			MapControl:     dataControl,
			CreatedAt:      dataRiskIssue.CreatedAt,
			UpdatedAt:      dataRiskIssue.UpdatedAt,
		}

		return responses, err
	}

	return responses, err

}

// SearchRiskIssue implements RiskIssueDefinition
func (riskIssue RiskIssueService) SearchRiskIssue(request models.KeywordRequest) (responses []models.RiskIssueResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataRiskIssue, totalRows, totalData, err := riskIssue.riskissueRepo.SearchRiskIssue(&request)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataRiskIssue {
		responses = append(responses, models.RiskIssueResponses{
			ID:            response.ID,
			RiskTypeID:    response.RiskTypeID,
			RiskIssueCode: response.RiskIssueCode,
			RiskIssue:     response.RiskIssue,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// SearchRiskIssueWithoutSub implements RiskIssueDefinition.
func (riskIssue RiskIssueService) SearchRiskIssueWithoutSub(request models.RiskIssueWithoutSub) (responses []models.RiskIssueResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataRiskIssue, totalRows, totalData, err := riskIssue.riskissueRepo.SearchRiskIssueWithoutSub(&request)

	fmt.Println("totalRows =>", totalRows)
	fmt.Println("tottalData =>", totalData)

	// fmt.Println("data =>", dataRiskIssue)

	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataRiskIssue {
		responses = append(responses, models.RiskIssueResponses{
			ID:            response.ID,
			RiskTypeID:    response.RiskTypeID,
			RiskIssueCode: response.RiskIssueCode,
			RiskIssue:     response.RiskIssue,
		})
	}

	fmt.Println("service response => ", responses)
	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err

}

// GetMappingIndicatorbyID implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetMappingIndicatorbyID(id int64) (responses models.RiskIssueResponseGetOne, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.GetOne(id)
	if dataRiskIssue.ID != 0 {
		fmt.Println("bukan 0")
		dataIndicator, err := riskIssue.mapIndicator.GetOneDataByID(dataRiskIssue.ID)

		fmt.Println(dataIndicator)

		responses = models.RiskIssueResponseGetOne{
			ID:             dataRiskIssue.ID,
			RiskTypeID:     dataRiskIssue.RiskTypeID,
			RiskIssueCode:  dataRiskIssue.RiskIssueCode,
			RiskIssue:      dataRiskIssue.RiskIssue,
			Deskripsi:      dataRiskIssue.Deskripsi,
			KategoriRisiko: dataRiskIssue.KategoriRisiko,
			Status:         dataRiskIssue.Status,
			Likelihood:     dataRiskIssue.Likelihood,
			Impact:         dataRiskIssue.Impact,
			MapIndicator:   dataIndicator,
			CreatedAt:      dataRiskIssue.CreatedAt,
			UpdatedAt:      dataRiskIssue.UpdatedAt,
		}

		return responses, err
	}

	return responses, err
}

// MappingRiskIndicator implements RiskIssueDefinition
func (riskIssue RiskIssueService) MappingRiskIndicator(request models.MappingIndicatorRequest) (responses bool, err error) {
	tx := riskIssue.db.DB.Begin()

	if len(request.MapIndicator) != 0 {
		for _, value := range request.MapIndicator {
			_, err = riskIssue.mapIndicator.Store(&models.MapIndicator{
				ID:          value.ID,
				IDRiskIssue: request.ID,
				IDIndicator: value.IDIndicator,
				IsChecked:   value.IsChecked,
			}, tx)

			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// Delete implements RiskIssueDefinition
func (riskIssue RiskIssueService) Delete(request *models.RiskIssueDeleteRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := riskIssue.db.DB.Begin()

	getRiskIssue, exist, err := riskIssue.GetOne(request.ID)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	updateRiskIssue := &models.RiskIssueDeleteRequest{
		ID:         request.ID,
		DeleteFlag: true,
		UpdatedAt:  &timeNow,
	}

	include := []string{
		"delete_flag",
		"updated_at",
	}

	_, err = riskIssue.riskissueRepo.Delete(updateRiskIssue, include, tx)

	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("Risk Issue =>", getRiskIssue)
		tx.Commit()
		return true, err
	}

	return false, err
}

// FilterRiskIssue implements RiskIssueDefinition
func (riskIssue RiskIssueService) FilterRiskIssue(request models.FilterRiskIssueRequest) (responses []models.RiskIssueFilterResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataRiskIssue, totalRows, totalData, err := riskIssue.riskissueRepo.FilterRiskIssue(&request)

	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataRiskIssue {
		responses = append(responses, models.RiskIssueFilterResponses{
			ID:             response.ID,
			RiskTypeID:     response.RiskTypeID,
			RiskIssueCode:  response.RiskIssueCode,
			RiskIssue:      response.RiskIssue,
			KategoriRisiko: response.KategoriRisiko,
			Status:         response.Status,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetRekomendasiMateri implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetRekomendasiMateri(id int64) (responses []models.RekomendasiMateri, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.GetRekomendasiMateri(id)

	if err != nil {
		riskIssue.logger.Zap.Error()
		return responses, err
	}

	for _, response := range dataRiskIssue {
		responses = append(responses, models.RekomendasiMateri{
			ID:           response.ID.Int64,
			IDIndicator:  response.IDIndicator.Int64,
			NamaLampiran: response.NamaLampiran.String,
			Path:         response.Path.String,
			Filename:     response.Filename.String,
		})
	}

	return responses, err
}

// GetRiskIssueByActivity implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetRiskIssueByActivity(id int64) (responses []models.RiskIssueResponseByActivity, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.GetRiskIssueByActivity(id)

	if err != nil {
		riskIssue.logger.Zap.Error()
		return responses, err
	}

	for _, response := range dataRiskIssue {
		responses = append(responses, models.RiskIssueResponseByActivity{
			ID:            response.ID.Int64,
			RiskIssueCode: response.RiskIssueCode.String,
			RiskIssue:     response.RiskIssue.String,
		})
	}

	return responses, err
}

// GetMateriByCode implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetMateriByCode(request models.RiskIssueCode) (responses []models.ListMateri, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.GetMateriByCode(request)

	if err != nil {
		riskIssue.logger.Zap.Error()
		return responses, err
	}

	for _, response := range dataRiskIssue {
		responses = append(responses, models.ListMateri{
			ID:            response.ID.Int64,
			IDIndicator:   response.IDIndicator.Int64,
			NamaLampiran:  response.NamaLampiran.String,
			NomorLampiran: response.NomorLampiran.String,
			JenisFile:     response.JenisFile.String,
			Path:          response.Path.String,
			Filename:      response.Filename.String,
		})
	}

	return responses, err
}

// GetRiskIssueByActivityID implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetRiskIssueByActivityID(id int64) (responses []models.RiskIssueResponseByActivity, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.GetRiskIssueByActivityID(id)

	if err != nil {
		riskIssue.logger.Zap.Error()
		return responses, err
	}

	for _, response := range dataRiskIssue {
		responses = append(responses, models.RiskIssueResponseByActivity{
			ID:            response.ID.Int64,
			RiskIssueCode: response.RiskIssueCode.String,
			RiskIssue:     response.RiskIssue.String,
		})
	}

	return responses, err
}

// GetRiskIssueByActivityID implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetRiskEventName(id int64) (name string, err error) {
	var data models.RiskIssueName
	data, err = riskIssue.riskissueRepo.GetRiskEventName(id)
	if err != nil {
		return "", err
	}

	name = data.RiskIssue

	return name, err
}
