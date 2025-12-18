package riskissue

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"riskmanagement/dto"
	"riskmanagement/lib"
	modelsControl "riskmanagement/models/riskcontrol"
	modelsIndicator "riskmanagement/models/riskindicator"
	models "riskmanagement/models/riskissue"
	"riskmanagement/repository/eventtypelv1"
	"riskmanagement/repository/eventtypelv2"
	"riskmanagement/repository/eventtypelv3"
	"riskmanagement/repository/incident"
	"riskmanagement/repository/penyebabkejadianlv3"
	"riskmanagement/repository/product"
	"riskmanagement/repository/riskcontrol"
	"riskmanagement/repository/riskindicator"
	riskissue "riskmanagement/repository/riskissue"
	"riskmanagement/repository/risktype"
	"riskmanagement/repository/subincident"
	"riskmanagement/services/arlords"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"

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
	ListRiskIssue(request models.ListRiskIssueRequest) (responses []models.ListRiskIssueResponse, err error)
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
	UpdateStatus(id int64) (bool, error)
	Template() ([]byte, string, error)
	PreviewData(pernr string, data [][]string) (dto.PreviewFileImport[[27]string], error)
	ImportData(pernr string, data [][]string) error
	Download(pernr, format string) ([]byte, string, error)
	GetMappingControlPaginate(id int64, filter modelsControl.Paginate) (response []models.MapControlResponseFinal, paginate lib.Pagination, err error)
	GetMapIndicatorWithPaginate(id int, filter modelsIndicator.Paginate) (response []models.MapIndicatorResponseFinal, pagination lib.Pagination, err error)
	GetRiskCategories(id []int64) ([]string, error)
}

type RiskIssueService struct {
	db             lib.Database
	dbRaw          lib.Databases
	logger         logger.Logger
	arlodsService  arlords.ArlordsServiceDefinition
	riskissueRepo  riskissue.RiskIssueDefinition
	mapAktifitas   riskissue.MapAktifitasDefinition
	mapEvent       riskissue.MapEventDefinition
	mapLiniBisnis  riskissue.MapLiniBisnisDefinition
	mapKejadian    riskissue.MapKejadianDefinition
	mapProduct     riskissue.MapProductDefinition
	mapProses      riskissue.MapProsesDefinition
	mapControl     riskissue.MapControlDefinition
	mapIndicator   riskissue.MapIndicatorDefinition
	eventTypelvl1  eventtypelv1.EventTypeLv1Definition
	eventTypeLvl2  eventtypelv2.EventTypeLv2Definition
	eventTypeLvl3  eventtypelv3.EventTypeLv3Definition
	incident       incident.IncidentDefinition
	subIncident    subincident.SubIncidentDefinition
	subsubIncident penyebabkejadianlv3.PenyebabKejadianLv3Definition
	product        product.ProductDefinition
	riskType       risktype.RiskTypeDefinition
	riskControl    riskcontrol.RiskControlDefinition
	riskIndicator  riskindicator.RiskIndicatorDefinition
}

func NewRiskIssueService(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
	arlodsService arlords.ArlordsServiceDefinition,
	riskissueRepo riskissue.RiskIssueDefinition,
	mapAktifitas riskissue.MapAktifitasDefinition,
	mapEvent riskissue.MapEventDefinition,
	mapLiniBisnis riskissue.MapLiniBisnisDefinition,
	mapKejadian riskissue.MapKejadianDefinition,
	mapProduct riskissue.MapProductDefinition,
	mapProses riskissue.MapProsesDefinition,
	mapControl riskissue.MapControlDefinition,
	mapIndicator riskissue.MapIndicatorDefinition,
	eventTypelvl1 eventtypelv1.EventTypeLv1Definition,
	eventTypeLvl2 eventtypelv2.EventTypeLv2Definition,
	eventTypeLvl3 eventtypelv3.EventTypeLv3Definition,
	incident incident.IncidentDefinition,
	subIncident subincident.SubIncidentDefinition,
	subsubIncident penyebabkejadianlv3.PenyebabKejadianLv3Definition,
	product product.ProductDefinition,
	riskType risktype.RiskTypeDefinition,
	riskControl riskcontrol.RiskControlDefinition,
	riskIndicator riskindicator.RiskIndicatorDefinition,
) RiskIssueDefinition {
	return RiskIssueService{
		db:             db,
		dbRaw:          dbRaw,
		logger:         logger,
		arlodsService:  arlodsService,
		riskissueRepo:  riskissueRepo,
		mapAktifitas:   mapAktifitas,
		mapEvent:       mapEvent,
		mapLiniBisnis:  mapLiniBisnis,
		mapKejadian:    mapKejadian,
		mapProduct:     mapProduct,
		mapProses:      mapProses,
		mapControl:     mapControl,
		mapIndicator:   mapIndicator,
		eventTypelvl1:  eventTypelvl1,
		eventTypeLvl2:  eventTypeLvl2,
		eventTypeLvl3:  eventTypeLvl3,
		incident:       incident,
		subIncident:    subIncident,
		subsubIncident: subsubIncident,
		product:        product,
		riskType:       riskType,
		riskControl:    riskControl,
		riskIndicator:  riskIndicator,
	}
}

// GetAll implements RiskIssueDefinition
func (riskIssue RiskIssueService) GetAll() (responses []models.RiskIssueResponse, err error) {
	return riskIssue.riskissueRepo.GetAll()
}

// GetAllWithPaginate implements RiskIssueDefinition
func (ri RiskIssueService) GetAllWithPaginate(request models.Paginate) (responses []models.RiskIssueResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Limit = limit
	request.Page = page
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
	}

	// Update MapEvent
	if len(request.MapEvent) != 0 {
		for _, value := range request.MapEvent {

			lvl1 := strings.TrimSpace(value.EventTypeLv1)
			lvl2 := strings.TrimSpace(value.EventTypeLv2)
			lvl3 := strings.TrimSpace(value.EventTypeLv3)

			if lvl1 == "" {
				continue
			}

			if lvl3 != "" && lvl2 == "" {
				continue
			}

			updateEvent := &models.MapEvent{
				IDRiskIssue:  dataRiskIssue.ID,
				EventTypeLv1: lvl1,
				EventTypeLv2: lvl2,
				EventTypeLv3: lvl3,
			}

			_, err = riskIssue.mapEvent.Update(updateEvent, tx)
			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	}

	// Update MapKejadian (Cause)
	if len(request.MapKejadian) != 0 {
		for _, value := range request.MapKejadian {

			lvl1 := strings.TrimSpace(value.PenyebabKejadianLv1)
			lvl2 := strings.TrimSpace(value.PenyebabKejadianLv2)
			lvl3 := strings.TrimSpace(value.PenyebabKejadianLv3)

			if lvl1 == "" {
				continue
			}

			if lvl3 != "" && lvl2 == "" {
				continue
			}
			updateKejadian := &models.MapKejadian{
				IDRiskIssue:         dataRiskIssue.ID,
				PenyebabKejadianLv1: lvl1,
				PenyebabKejadianLv2: lvl2,
				PenyebabKejadianLv3: lvl3,
			}

			_, err = riskIssue.mapKejadian.Update(updateKejadian, tx)
			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
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
	}

	tx.Commit()
	return true, err
}

// Update implements RiskIssueDefinition
func (riskIssue RiskIssueService) Update(request *models.RiskIssueRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	exists, err := riskIssue.riskissueRepo.GetOne(request.ID)
	if err != nil {
		riskIssue.logger.Zap.Error("Error when query existing issue: %s", err)
		return false, err
	}

	tx := riskIssue.db.DB.Begin()

	updateRiskIssue := &models.RiskIssueUpdate{
		ID:             request.ID,
		RiskTypeID:     request.RiskTypeID,
		RiskIssueCode:  exists.RiskIssueCode,
		RiskIssue:      request.RiskIssue,
		Deskripsi:      request.Deskripsi,
		KategoriRisiko: request.KategoriRisiko,
		Status:         request.Status,
		Likelihood:     request.Likelihood,
		Impact:         request.Likelihood,
		UpdatedAt:      &timeNow,
	}

	include := []string{
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
	}

	// Update MapEvent
	if len(request.MapEvent) != 0 {
		for _, value := range request.MapEvent {

			lvl1 := strings.TrimSpace(value.EventTypeLv1)
			lvl2 := strings.TrimSpace(value.EventTypeLv2)
			lvl3 := strings.TrimSpace(value.EventTypeLv3)

			if lvl1 == "" {
				continue
			}

			if lvl3 != "" && lvl2 == "" {
				continue
			}

			updateEvent := &models.MapEvent{
				ID:           value.ID,
				IDRiskIssue:  request.ID,
				EventTypeLv1: lvl1,
				EventTypeLv2: lvl2,
				EventTypeLv3: lvl3,
			}

			_, err = riskIssue.mapEvent.Update(updateEvent, tx)
			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
	}

	// Update MapKejadian (Cause)
	if len(request.MapKejadian) != 0 {
		for _, value := range request.MapKejadian {

			lvl1 := strings.TrimSpace(value.PenyebabKejadianLv1)
			lvl2 := strings.TrimSpace(value.PenyebabKejadianLv2)
			lvl3 := strings.TrimSpace(value.PenyebabKejadianLv3)

			if lvl1 == "" {
				continue
			}

			if lvl3 != "" && lvl2 == "" {
				continue
			}

			updateKejadian := &models.MapKejadian{
				ID:                  value.ID,
				IDRiskIssue:         request.ID,
				PenyebabKejadianLv1: lvl1,
				PenyebabKejadianLv2: lvl2,
				PenyebabKejadianLv3: lvl3,
			}

			_, err = riskIssue.mapKejadian.Update(updateKejadian, tx)
			if err != nil {
				tx.Rollback()
				riskIssue.logger.Zap.Error(err)
				return false, err
			}
		}
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
func (riskIssue RiskIssueService) MappingRiskControl(
	request models.MappingControlRequest,
) (bool, error) {

	tx := riskIssue.db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ================= VALIDASI REQUEST =================
	if len(request.MapControl) == 0 {
		err := errors.New("map_control is empty")
		riskIssue.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	// ================= GET EXISTING MAPPING =================
	dataControl, err := riskIssue.mapControl.GetOneDataByID(request.ID)
	if err != nil {
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	// ================= BUILD LOOKUP MAP =================
	// key format: riskIssueID_controlID
	existingMap := make(map[string]struct{})
	for _, v := range dataControl {
		key := fmt.Sprintf("%d_%d", v.IDRiskIssue, v.IDControl)
		existingMap[key] = struct{}{}
	}

	var insertedCount int

	// ================= INSERT ONLY NEW DATA =================
	for _, value := range request.MapControl {

		key := fmt.Sprintf("%d_%d", request.ID, value.IDControl)

		// skip jika sudah ada
		if _, exists := existingMap[key]; exists {
			continue
		}

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

		insertedCount++
	}

	// ================= VALIDASI DATA BARU =================
	if insertedCount == 0 {
		err := errors.New("no new control mapping inserted, all controls already mapped")
		tx.Rollback()
		riskIssue.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, nil
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

func (riskIssue RiskIssueService) GetMappingControlPaginate(id int64, filter modelsControl.Paginate) (response []models.MapControlResponseFinal, paginate lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(filter.Page, filter.Limit, filter.Order, filter.Sort)
	filter.Limit = limit
	filter.Page = page
	filter.Offset = offset
	filter.Order = order
	filter.Sort = sort

	if filter.Order == "" {
		filter.Order = "id"
	}

	data, total, err := riskIssue.mapControl.GetWithPagination(int(id), filter)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return response, paginate, err
	}

	paginate = lib.SetPaginationResponse(filter.Page, filter.Limit, total, total)
	return data, paginate, nil
}

func (riskIssue RiskIssueService) ListRiskIssue(request models.ListRiskIssueRequest) (responses []models.ListRiskIssueResponse, err error) {
	dataRiskIssue, err := riskIssue.riskissueRepo.RiskIssueByIndicator(request)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return responses, err
	}

	return dataRiskIssue, err
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

func (riskIssue RiskIssueService) GetMapIndicatorWithPaginate(id int, filter modelsIndicator.Paginate) (response []models.MapIndicatorResponseFinal, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(filter.Page, filter.Limit, filter.Order, filter.Sort)
	filter.Limit = limit
	filter.Page = page
	filter.Offset = offset
	filter.Order = order
	filter.Sort = sort

	if filter.Order == "" {
		filter.Order = "id"
	}

	data, total, err := riskIssue.mapIndicator.GetWithPaginate(id, filter)
	if err != nil {
		riskIssue.logger.Zap.Error(err)
		return response, pagination, err
	}

	pagination = lib.SetPaginationResponse(filter.Page, filter.Limit, total, total)
	return data, pagination, nil
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
		existingData, err := riskIssue.mapIndicator.GetOneDataByID(request.ID)
		if err != nil {
			tx.Rollback()
			riskIssue.logger.Zap.Error(err)
			return false, err
		}

		existingMap := make(map[string]struct{})
		for _, v := range existingData {
			key := fmt.Sprintf("%d_%d", v.IDRiskIssue, v.IDIndicator)
			existingMap[key] = struct{}{}
		}

		var insertedCount int

		for _, value := range request.MapIndicator {

			key := fmt.Sprintf("%d_%d", request.ID, value.IDIndicator)

			// skip jika sudah ada
			if _, exists := existingMap[key]; exists {
				continue
			}

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

			insertedCount++
		}

		if insertedCount == 0 {
			err := errors.New("no new mapping inserted, all indicators already mapped")
			tx.Rollback()
			riskIssue.logger.Zap.Error(err)
			return false, err
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

func (riskIssue RiskIssueService) UpdateStatus(id int64) (bool, error) {
	var (
		status bool = true
	)
	data, err := riskIssue.riskissueRepo.GetOne(id)
	if err != nil {
		return status, err
	}

	if data.Status {
		status = false
	}

	err = riskIssue.riskissueRepo.UpdateStatus(id, status)

	return status, err
}

func (riskIssue RiskIssueService) Template() ([]byte, string, error) {
	f := excelize.NewFile()
	sheet := "Template"

	f.SetSheetName("Sheet1", sheet)

	sheetIndex, err := f.GetSheetIndex(sheet)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get sheet index: %w", err)
	}
	f.SetActiveSheet(sheetIndex)

	headers := []string{
		"Tipe Risiko",
		"Risk Event",
		"Deskripsi",
		"Kategori Risiko",
		"Map Event - Event Type Lv1",
		"Map Event - Event Type Lv2",
		"Map Event - Event Type Lv3",
		"Map Kejadian - Penyebab Kejadian Lv1",
		"Map Kejadian - Penyebab Kejadian Lv2",
		"Map Kejadian - Penyebab Kejadian Lv3",
		"Map Product - Product",
		"Likelihood",
		"Impact",
		"Code Control",
		"Risk Control",
		"Code Indicator",
		"Risk Indicator",
		"Code Business Cycle",
		"Business Cycle",
		"Code Sub Business Cycle",
		"Sub Business Cycle",
		"Code Process",
		"Process",
		"Code Sub Process",
		"Sub Process",
		"Code Activity",
		"Activity",
	}

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			return nil, "", fmt.Errorf("failed to set cell: %w", err)
		}
	}

	// optional: set lebar kolom
	for i := 1; i <= len(headers); i++ {
		col, _ := excelize.ColumnNumberToName(i)
		f.SetColWidth(sheet, col, col, 25)
	}

	// simpan ke buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", fmt.Errorf("failed to write excel: %w", err)
	}

	return buf.Bytes(), "risk_event_template.xlsx", nil
}

func (riskIssue RiskIssueService) PreviewData(pernr string, data [][]string) (dto.PreviewFileImport[[27]string], error) {
	issues, err := riskIssue.riskissueRepo.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query risk issue: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	// Mapping code -> ID
	riskIssueMap := make(map[string]int64)
	for _, e := range issues {
		riskIssueMap[e.RiskIssueCode] = e.ID
	}

	riskType, err := riskIssue.riskType.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query risk type: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	riskTypeMap := make(map[string]bool, len(riskType))
	riskTypeIDMap := make(map[string]int64, len(riskType))
	for _, a := range riskType {
		key := strings.ToLower(a.RiskTypeCode)
		riskTypeMap[key] = true
		riskTypeIDMap[key] = a.ID
	}

	eventLv1, err := riskIssue.eventTypelvl1.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query event type lvl 1: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	eventTypeLv1Map := make(map[string]bool, len(eventLv1))
	eventTypeLv1IDMap := make(map[string]int64, len(eventLv1))
	for _, a := range eventLv1 {
		key := strings.ToLower(a.KodeEventType)
		eventTypeLv1Map[key] = true
		eventTypeLv1IDMap[key] = a.ID
	}

	eventLv2, err := riskIssue.eventTypeLvl2.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query event type lvl 2: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	eventTypelv2Map := make(map[string]bool, len(eventLv2))
	eventTypelv2IDMap := make(map[string]int64, len(eventLv2))
	eventLv2Parent := make(map[string]string)

	for _, a := range eventLv2 {
		key := strings.ToLower(a.KodeEventTypeLv2)
		parent := strings.ToLower(a.IDEventTypeLv1)

		eventTypelv2Map[key] = true
		eventTypelv2IDMap[key] = a.ID
		eventLv2Parent[key] = parent
	}

	eventLv3, err := riskIssue.eventTypeLvl3.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query event type lvl 3: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	eventTypelv3Map := make(map[string]bool, len(eventLv3))
	eventTypelv3IDMap := make(map[string]int64, len(eventLv3))
	eventLv3Parent := make(map[string]string)

	for _, a := range eventLv3 {
		key := strings.ToLower(a.KodeEventTypeLv3)
		parent := strings.ToLower(a.IDEventTypeLv2)

		eventTypelv3Map[key] = true
		eventTypelv3IDMap[key] = a.ID
		eventLv3Parent[key] = parent
	}

	incident, err := riskIssue.incident.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query incident: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	incidentMap := make(map[string]bool, len(incident))
	incidentIDMap := make(map[string]int64, len(incident))
	for _, a := range incident {
		key := strings.ToLower(a.KodeKejadian)
		incidentMap[key] = true
		incidentIDMap[key] = a.ID
	}

	subIncident, err := riskIssue.subIncident.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query subIncident: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	subIncidentParentMap := make(map[string]string)
	subIncidentMap := make(map[string]bool)
	subIncidentIDMap := make(map[string]int64)

	for _, a := range subIncident {
		key := strings.ToLower(a.KodeSubKejadian)
		parent := strings.ToLower(a.KodeKejadian)

		subIncidentMap[key] = true
		subIncidentIDMap[key] = a.ID
		subIncidentParentMap[key] = parent
	}

	subsubIncident, err := riskIssue.subsubIncident.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query subSubIncident: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	subsubIncidentMap := make(map[string]bool)
	subsubIncidentIDMap := make(map[string]int64)
	subsubIncidentParentMap := make(map[string]string)

	for _, a := range subsubIncident {
		key := strings.ToLower(a.KodePenyebabKejadianLv3)
		parent := strings.ToLower(a.KodeSubKejadian)

		subsubIncidentMap[key] = true
		subsubIncidentIDMap[key] = a.ID
		subsubIncidentParentMap[key] = parent
	}

	product, err := riskIssue.product.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query product: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	productMap := make(map[string]bool)
	productIDMap := make(map[string]int64)
	for _, a := range product {
		key := strings.ToLower(a.KodeProduct)
		productMap[key] = true
		productIDMap[key] = a.ID
	}

	control, err := riskIssue.riskControl.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query risk control: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	controlMap := make(map[string]bool)
	controlIDMap := make(map[string]int64)
	for _, a := range control {
		key := strings.ToLower(a.Kode)
		controlMap[key] = true
		controlIDMap[key] = a.ID
	}

	indicator, err := riskIssue.riskIndicator.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query risk indicator: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	indicatorMap := make(map[string]bool)
	indicatorIDMap := make(map[string]int64)
	for _, a := range indicator {
		key := strings.ToLower(a.RiskIndicatorCode)
		indicatorMap[key] = true
		indicatorIDMap[key] = a.ID
	}

	mapEvent, err := riskIssue.mapEvent.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event event lvl: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	mapEventLvlMap := make(map[string]bool)
	for _, v := range mapEvent {
		lvl1 := strings.ToLower(v.EventTypeLv1)
		lvl2 := strings.ToLower(v.EventTypeLv2)
		lvl3 := strings.ToLower(v.EventTypeLv3)
		key := fmt.Sprintf("%d|%s|%s|%s", v.IDRiskIssue, lvl1, lvl2, lvl3)
		mapEventLvlMap[key] = true
	}

	mapKejadian, err := riskIssue.mapKejadian.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event kejadian lvl: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}
	mapKejadianLvlMap := make(map[string]bool)
	for _, v := range mapKejadian {
		lvl1 := strings.ToLower(v.PenyebabKejadianLv1)
		lvl2 := strings.ToLower(v.PenyebabKejadianLv2)
		lvl3 := strings.ToLower(v.PenyebabKejadianLv3)
		key := fmt.Sprintf("%d|%s|%s|%s", v.IDRiskIssue, lvl1, lvl2, lvl3)
		mapKejadianLvlMap[key] = true
	}

	mapProduct, err := riskIssue.mapProduct.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event product: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}
	mapProductMap := make(map[string]bool)
	for _, v := range mapProduct {
		key := fmt.Sprintf("%d|%d", v.IDRiskIssue, v.Product)
		mapProductMap[key] = true
	}
	mapControl, err := riskIssue.mapControl.GetAllWithRelation()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event control: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}
	mapControlMap := make(map[string]bool)
	for _, v := range mapControl {
		control := strings.ToLower(v.Kode)
		key := fmt.Sprintf("%d|%s", v.IDRiskIssue, control)
		mapControlMap[key] = true
	}
	mapIndicator, err := riskIssue.mapIndicator.GetAllWithRelation()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event indicator: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}
	mapIndicatorMap := make(map[string]bool)
	for _, v := range mapIndicator {
		indicator := strings.ToLower(v.Kode)
		key := fmt.Sprintf("%d|%s", v.IDRiskIssue, indicator)
		mapIndicatorMap[key] = true
	}

	businessProces, err := riskIssue.arlodsService.BulkGetBusinessProcessByActivity(pernr, nil)
	if err != nil {
		riskIssue.logger.Zap.Error("errored query business process: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	businessProcessMap := make(map[string]bool)
	for _, v := range businessProces.Data {
		key := fmt.Sprintf("%s|%s|%s|%s|%s", v.BusinessCycleCode, v.SubBusinessCycleCode, v.ProcessCode, v.SubProcessCode, v.ActivityCode)
		businessProcessMap[strings.ToLower(key)] = true
	}

	mapBusinessProcess, err := riskIssue.arlodsService.BulkGetMappingEventBusinessProcess(pernr, nil)
	if err != nil {
		riskIssue.logger.Zap.Error("errored query mapping business process: %s", err)
		return dto.PreviewFileImport[[27]string]{}, err
	}

	mapBusinessProcessMap := make(map[string]bool)
	for _, v := range mapBusinessProcess.Data {
		key := fmt.Sprintf("%s|%s|%s|%s|%s|%s", v.RiskEventID, v.BusinessCode, v.SubBusinessCode, v.ProcessCode, v.SubProcessCode, v.ActivityCode)
		mapBusinessProcessMap[strings.ToLower(key)] = true
	}

	headers := [27]string{
		"Tipe Risiko",
		"Risk Event",
		"Deskripsi",
		"Kategori Risiko",
		"Map Event - Event Type Lv1",
		"Map Event - Event Type Lv2",
		"Map Event - Event Type Lv3",
		"Map Kejadian - Penyebab Kejadian Lv1",
		"Map Kejadian - Penyebab Kejadian Lv2",
		"Map Kejadian - Penyebab Kejadian Lv3",
		"Map Product - Product",
		"Likelihood",
		"Impact",
		"Code Control",
		"Risk Control",
		"Code Indicator",
		"Risk Indicator",
		"Code Business Cycle",
		"Business Cycle",
		"Code Sub Business Cycle",
		"Sub Business Cycle",
		"Code Process",
		"Process",
		"Code Sub Process",
		"Sub Process",
		"Code Activity",
		"Activity",
	}

	previewFile := dto.PreviewFileImport[[27]string]{}
	body := []dto.PreviewFile[[27]string]{}

	eventIDs := make([]string, 0)
	mapping := make(map[string]dto.MappingEvent, 0)

	for i, row := range data {
		if i == 0 {
			continue
		}
		riskType := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(row[0], "-")))
		eventIDs = append(eventIDs, riskType)
	}

	for i, row := range data {
		validation := ""
		col := [27]string{}
		event := dto.MappingEvent{}

		if i == 0 {
			if len(row) < 27 {
				return dto.PreviewFileImport[[27]string]{}, fmt.Errorf("invalid header format risk event")
			}

			for i, v := range headers {
				if strings.TrimSpace(row[i]) != v {
					return dto.PreviewFileImport[[27]string]{}, fmt.Errorf("header kolom ke-%d invalid format, diharapkan '%s', diterima '%s'", i+1, v, row[i])
				}
			}
			previewFile.Header = headers
			continue
		}

		parse := lib.ParseStringToArray(row[1], "|")
		var existingEvent int64 = 0
		riskEventCode := parse[0]
		if val, ok := riskIssueMap[riskEventCode]; ok {
			validation += fmt.Sprintf("Risk Event Sudah terdaftar: %s", row[1])
			existingEvent = val
		}

		businessCycleStr := lib.ParseStringToArray(row[17], ";")
		subBusinessCycleStr := lib.ParseStringToArray(row[19], ";")
		processStr := lib.ParseStringToArray(row[21], ";")
		subProcessStr := lib.ParseStringToArray(row[23], ";")
		activityStr := lib.ParseStringToArray(row[25], ";")

		riskType := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(row[0], "-")))
		eventLv1 := lib.ParseStringToArray(row[4], ";")
		eventLv2 := lib.ParseStringToArray(row[5], ";")
		eventLv3 := lib.ParseStringToArray(row[6], ";")
		incident := lib.ParseStringToArray(row[7], ";")
		subIncident := lib.ParseStringToArray(row[8], ";")
		subsubIncident := lib.ParseStringToArray(row[9], ";")
		product := lib.ParseStringToArray(row[10], ";")
		control := lib.ParseStringToArray(row[13], ";")
		indicator := lib.ParseStringToArray(row[15], ";")

		if _, ok := riskTypeMap[riskType]; !ok {
			validation += fmt.Sprintf("Risk Type tidak terdaftar: %s; ", riskType)
		}
		eventIDs = append(eventIDs, riskType)

		evnLv1map, evnLv2map, evnLv3map := []string{}, []string{}, []string{}
		for i := range eventLv1 {
			lv1 := normalizeEventCode(eventLv1[i])
			lv2 := normalizeEventCode(eventLv2[i])
			lv3 := normalizeEventCode(eventLv3[i])

			// === cek apakah sudah termapping ===
			isAlreadyMapped := false
			if existingEvent > 0 {
				key := fmt.Sprintf("%d|%s|%s|%s", existingEvent, lv1, lv2, lv3)
				if mapEventLvlMap[key] {
					validation += fmt.Sprintf(
						"event LEVEL '%s' - '%s' - '%s' sudah termapping; ",
						lv1, lv2, lv3,
					)
					isAlreadyMapped = true
				}
			}

			// kalau sudah termapping → SKIP SEMUA INSERT
			if isAlreadyMapped {
				continue
			}

			// === LV1 ===
			if !eventTypeLv1Map[lv1] {
				validation += fmt.Sprintf(
					"Event LV1 tidak terdaftar: %s; ",
					eventLv1[i],
				)
			} else {
				evnLv1map = append(evnLv1map, lv1)
			}

			// === LV2 ===
			if !eventTypelv2Map[lv2] {
				validation += fmt.Sprintf(
					"Event LV2 tidak terdaftar: %s; ",
					eventLv2[i],
				)
			} else {
				evnLv2map = append(evnLv2map, lv2)
			}

			// === LV3 ===
			if !eventTypelv3Map[lv3] {
				validation += fmt.Sprintf(
					"Event LV3 tidak terdaftar: %s; ",
					eventLv3[i],
				)
			} else {
				evnLv3map = append(evnLv3map, lv3)
			}

			// === parent validation ===
			if parent, ok := eventLv2Parent[lv2]; ok && parent != lv1 {
				validation += fmt.Sprintf(
					"Event LV2 '%s' tidak terkait LV1 '%s'; ",
					eventLv2[i], eventLv1[i],
				)
			}

			if parent, ok := eventLv3Parent[lv3]; ok && parent != lv2 {
				validation += fmt.Sprintf(
					"Event LV3 '%s' tidak terkait LV2 '%s'; ",
					eventLv3[i], eventLv2[i],
				)
			}
		}

		event.EventLV1 = evnLv1map
		event.EventLv2 = evnLv2map
		event.EventLv3 = evnLv3map

		// Incident
		incLv1Map, incLv2Map, incLv3Map := []string{}, []string{}, []string{}
		for i := range incident {
			lv1 := normalizeEventCode(incident[i])
			lv2 := normalizeEventCode(subIncident[i])
			lv3 := normalizeEventCode(subsubIncident[i])

			// === cek sudah termapping ===
			isAlreadyMapped := false
			if existingEvent > 0 {
				key := fmt.Sprintf("%d|%s|%s|%s", existingEvent, lv1, lv2, lv3)
				if mapKejadianLvlMap[key] {
					validation += fmt.Sprintf(
						"penyebab kejadian LEVEL '%s' - '%s' - '%s' sudah termapping; ",
						lv1, lv2, lv3,
					)
					isAlreadyMapped = true
				}
			}

			// kalau sudah termapping → SKIP SEMUA APPEND
			if isAlreadyMapped {
				continue
			}

			// === LV1 ===
			if !incidentMap[lv1] {
				validation += fmt.Sprintf(
					"Incident LV1 tidak terdaftar: %s; ",
					incident[i],
				)
			} else {
				incLv1Map = append(incLv1Map, lv1)
			}

			// === LV2 ===
			if lv2 != "" {
				if !subIncidentMap[lv2] {
					validation += fmt.Sprintf(
						"Incident LV2 tidak terdaftar: %s; ",
						subIncident[i],
					)
				} else {
					incLv2Map = append(incLv2Map, lv2)

					if parent, ok := subIncidentParentMap[lv2]; ok && parent != lv1 {
						validation += fmt.Sprintf(
							"Incident LV2 '%s' tidak terkait LV1 '%s'; ",
							subIncident[i], incident[i],
						)
					}
				}
			}

			// === LV3 ===
			if lv3 != "" {
				if !subsubIncidentMap[lv3] {
					validation += fmt.Sprintf(
						"Incident LV3 tidak terdaftar: %s; ",
						subsubIncident[i],
					)
				} else {
					incLv3Map = append(incLv3Map, lv3)

					if parent, ok := subsubIncidentParentMap[lv3]; ok && parent != lv2 {
						validation += fmt.Sprintf(
							"Incident LV3 '%s' tidak terkait LV2 '%s'; ",
							subsubIncident[i], subIncident[i],
						)
					}
				}
			}
		}

		event.Incident = incLv1Map
		event.SubIncident = incLv2Map
		event.SubSubIncident = incLv3Map

		productMapped := []string{}
		for _, p := range product {
			p = normalizeEventCode(p)
			if p == "" {
				continue
			}

			if !productMap[p] {
				validation += fmt.Sprintf("Product tidak terdaftar: %s; ", p)
			} else {
				if existingEvent > 0 {
					key := fmt.Sprintf("%d|%s", existingEvent, p)
					if productMap[key] {
						validation += fmt.Sprintf("Product '%s' sudah termapping", p)
					}
				}
				productMapped = append(productMapped, p)
			}
		}
		event.ProductIDs = append(event.ProductIDs, productMapped...)

		for _, v := range control {
			c := normalizeEventCode(v)
			if c == "" {
				continue
			}
			if !controlMap[c] {
				validation += fmt.Sprintf("Risk Control tidak terdaftar: %s", c)
			} else {
				if existingEvent > 0 {
					key := fmt.Sprintf("%d|%s", existingEvent, c)
					if mapControlMap[key] {
						validation += fmt.Sprintf("Risk Control '%s' sudah termapping", v)
					}
				}
			}
		}

		for _, v := range indicator {
			c := normalizeEventCode(v)
			if c == "" {
				continue
			}

			if !indicatorMap[c] {
				validation += fmt.Sprintf("Risk Indicator tidak terdaftar: %s", c)
			} else {
				if existingEvent > 0 {
					key := fmt.Sprintf("%d|%s", existingEvent, c)
					if mapIndicatorMap[key] {
						validation += fmt.Sprintf("Risk Indicator '%s' sudah termapping", v)
					}
				}
			}
		}

		for i := range activityStr {
			key := fmt.Sprintf("%s|%s|%s|%s|%s", businessCycleStr[i], subBusinessCycleStr[i], processStr[i], subProcessStr[i], activityStr[i])
			riskIssue.logger.Zap.Debug(businessProcessMap[strings.ToLower(key)])
			riskIssue.logger.Zap.Debug(key)
			if !businessProcessMap[strings.ToLower(key)] {
				validation += fmt.Sprintf("Maaf business process  busniss cycle '%s' sub busniss cycle'%s' process '%s' sub process '%s' activity '%s'  tidak terdaftar silahkan check kembali; ", businessCycleStr[i], subBusinessCycleStr[i], processStr[i], subProcessStr[i], activityStr[i])
			} else {
				if existingEvent > 0 {
					mapKey := fmt.Sprintf("%d|%s", existingEvent, key)
					if mapBusinessProcessMap[strings.ToLower(mapKey)] {
						validation += "maaf data business process sudah termapping"
					}
				}
			}
		}

		mapping[riskType] = event

		for z := range col {
			if z < len(row) {
				col[z] = row[z]
			}
		}

		body = append(body, dto.PreviewFile[[27]string]{
			PerRow:     col,
			Validation: validation,
		})
	}

	previewFile.Body = body

	return previewFile, nil
}
func (riskIssue RiskIssueService) ImportData(pernr string, data [][]string) error {
	// ================================
	// STEP 1 — Ambil reference data
	// ================================
	issues, err := riskIssue.riskissueRepo.GetAll()
	if err != nil {
		return fmt.Errorf("failed get risk event: %v", err)
	}

	// Mapping code -> ID
	riskIssueMap := make(map[string]int64)
	for _, e := range issues {
		riskIssueMap[e.RiskIssueCode] = e.ID
	}

	riskTypes, err := riskIssue.riskType.GetAll()
	if err != nil {
		return fmt.Errorf("failed get risk type: %v", err)
	}
	riskTypeMap := make(map[string]int64)
	for _, r := range riskTypes {
		riskTypeMap[strings.ToLower(r.RiskTypeCode)] = r.ID
	}

	eventLv1, _ := riskIssue.eventTypelvl1.GetAll()
	eventLv1Map := make(map[string]bool)
	for _, e := range eventLv1 {
		eventLv1Map[strings.ToLower(e.KodeEventType)] = true
	}

	eventLv2, _ := riskIssue.eventTypeLvl2.GetAll()
	eventLv2Map := make(map[string]bool)
	eventLv2Parent := make(map[string]string)
	for _, e := range eventLv2 {
		key := strings.ToLower(e.KodeEventTypeLv2)
		eventLv2Map[key] = true
		eventLv2Parent[key] = strings.ToLower(e.IDEventTypeLv1)
	}

	eventLv3, _ := riskIssue.eventTypeLvl3.GetAll()
	eventLv3Map := make(map[string]bool)
	eventLv3Parent := make(map[string]string)
	for _, e := range eventLv3 {
		key := strings.ToLower(e.KodeEventTypeLv3)
		eventLv3Map[key] = true
		eventLv3Parent[key] = strings.ToLower(e.IDEventTypeLv2)
	}

	incident, _ := riskIssue.incident.GetAll()
	incidentMap := make(map[string]bool)
	for _, i := range incident {
		incidentMap[strings.ToLower(i.KodeKejadian)] = true
	}

	subIncident, _ := riskIssue.subIncident.GetAll()
	subIncidentMap := make(map[string]bool)
	subIncidentParentMap := make(map[string]string)
	for _, i := range subIncident {
		key := strings.ToLower(i.KodeSubKejadian)
		subIncidentMap[key] = true
		subIncidentParentMap[key] = strings.ToLower(i.KodeKejadian)
	}

	subsubIncident, _ := riskIssue.subsubIncident.GetAll()
	subsubIncidentMap := make(map[string]bool)
	subsubIncidentParentMap := make(map[string]string)
	for _, i := range subsubIncident {
		key := strings.ToLower(i.KodePenyebabKejadianLv3)
		subsubIncidentMap[key] = true
		subsubIncidentParentMap[key] = strings.ToLower(i.KodeSubKejadian)
	}

	product, _ := riskIssue.product.GetAll()
	productMap := make(map[string]int64)
	for _, p := range product {
		productMap[strings.ToLower(p.KodeProduct)] = p.ID
	}

	control, err := riskIssue.riskControl.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query risk control: %s", err)
		return err
	}
	controlMap := make(map[string]bool)
	controlIDMap := make(map[string]int64)
	for _, a := range control {
		key := strings.ToLower(a.Kode)
		controlMap[key] = true
		controlIDMap[key] = a.ID
	}

	indicator, err := riskIssue.riskIndicator.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query risk indicator: %s", err)
		return err
	}
	indicatorMap := make(map[string]bool)
	indicatorIDMap := make(map[string]int64)
	for _, a := range indicator {
		key := strings.ToLower(a.RiskIndicatorCode)
		indicatorMap[key] = true
		indicatorIDMap[key] = a.ID
	}

	mapEvent, err := riskIssue.mapEvent.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event event lvl: %s", err)
		return err
	}

	mapEventLvlMap := make(map[string]bool)
	for _, v := range mapEvent {
		lvl1 := strings.ToLower(v.EventTypeLv1)
		lvl2 := strings.ToLower(v.EventTypeLv2)
		lvl3 := strings.ToLower(v.EventTypeLv3)
		key := fmt.Sprintf("%d|%s|%s|%s", v.IDRiskIssue, lvl1, lvl2, lvl3)
		mapEventLvlMap[key] = true
	}

	mapKejadian, err := riskIssue.mapKejadian.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event kejadian lvl: %s", err)
		return err
	}
	mapKejadianLvlMap := make(map[string]bool)
	for _, v := range mapKejadian {
		lvl1 := strings.ToLower(v.PenyebabKejadianLv1)
		lvl2 := strings.ToLower(v.PenyebabKejadianLv2)
		lvl3 := strings.ToLower(v.PenyebabKejadianLv3)
		key := fmt.Sprintf("%d|%s|%s|%s", v.IDRiskIssue, lvl1, lvl2, lvl3)
		mapKejadianLvlMap[key] = true
	}

	mapProduct, err := riskIssue.mapProduct.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event product: %s", err)
		return err
	}
	mapProductMap := make(map[string]bool)
	for _, v := range mapProduct {
		key := fmt.Sprintf("%d|%d", v.IDRiskIssue, v.Product)
		mapProductMap[key] = true
	}
	mapControl, err := riskIssue.mapControl.GetAllWithRelation()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event control: %s", err)
		return err
	}
	mapControlMap := make(map[string]bool)
	for _, v := range mapControl {
		control := strings.ToLower(v.Kode)
		key := fmt.Sprintf("%d|%s", v.IDRiskIssue, control)
		mapControlMap[key] = true
	}
	mapIndicator, err := riskIssue.mapIndicator.GetAllWithRelation()
	if err != nil {
		riskIssue.logger.Zap.Error("errored query map risk event indicator: %s", err)
		return err
	}
	mapIndicatorMap := make(map[string]bool)
	for _, v := range mapIndicator {
		indicator := strings.ToLower(v.Kode)
		key := fmt.Sprintf("%d|%s", v.IDRiskIssue, indicator)
		mapIndicatorMap[key] = true
	}

	businessProces, err := riskIssue.arlodsService.BulkGetBusinessProcessByActivity(pernr, nil)
	if err != nil {
		riskIssue.logger.Zap.Error("errored query business process: %s", err)
		return err
	}

	businessProcessMap := make(map[string]dto.BusinessHierarchyFlatResponse)
	for _, v := range businessProces.Data {
		key := fmt.Sprintf("%s|%s|%s|%s|%s", v.BusinessCycleCode, v.SubBusinessCycleCode, v.ProcessCode, v.SubProcessCode, v.ActivityCode)
		businessProcessMap[strings.ToLower(key)] = v
	}

	mapBusinessProcess, err := riskIssue.arlodsService.BulkGetMappingEventBusinessProcess(pernr, nil)
	if err != nil {
		riskIssue.logger.Zap.Error("errored query mapping business process: %s", err)
		return err
	}

	mapBusinessProcessMap := make(map[string]bool)
	for _, v := range mapBusinessProcess.Data {
		key := fmt.Sprintf("%s|%s|%s|%s|%s|%s", v.RiskEventID, v.BusinessCode, v.SubBusinessCode, v.ProcessCode, v.SubProcessCode, v.ActivityCode)
		mapBusinessProcessMap[strings.ToLower(key)] = true
	}

	headers := [27]string{
		"Tipe Risiko",
		"Risk Event",
		"Deskripsi",
		"Kategori Risiko",
		"Map Event - Event Type Lv1",
		"Map Event - Event Type Lv2",
		"Map Event - Event Type Lv3",
		"Map Kejadian - Penyebab Kejadian Lv1",
		"Map Kejadian - Penyebab Kejadian Lv2",
		"Map Kejadian - Penyebab Kejadian Lv3",
		"Map Product - Product",
		"Likelihood",
		"Impact",
		"Code Control",
		"Risk Control",
		"Code Indicator",
		"Risk Indicator",
		"Code Business Cycle",
		"Business Cycle",
		"Code Sub Business Cycle",
		"Sub Business Cycle",
		"Code Process",
		"Process",
		"Code Sub Process",
		"Sub Process",
		"Code Activity",
		"Activity",
	}

	// ================================
	// STEP 3 — Proses import row
	// ================================
	newRiskEvents := []models.RiskIssue{}
	newMapEvent := make(map[string]*models.MapEvent)
	newMapKejadian := make(map[string]*models.MapKejadian)
	newMapProduct := make(map[string]*models.MapProduct)
	newMapControl := make(map[string]*models.MapControl)
	newMapIndicator := make(map[string]*models.MapIndicator)
	mappingBusinessProcess := []dto.MapingBusinessProcessRequest{}
	timeNow := lib.GetTimeNow("timestime")
	for i, row := range data {
		if i == 0 {
			continue
		}

		riskTypeCode := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(row[0], "-")))
		riskTypeID := int64(0)
		if id, ok := riskTypeMap[riskTypeCode]; ok {
			riskTypeID = id
		}

		if riskTypeID == 0 {
			if len(row) < 27 {
				return fmt.Errorf("invalid header format risk event")
			}

			for i, v := range headers {
				if strings.TrimSpace(row[i]) != v {
					return fmt.Errorf("header kolom ke-%d invalid format, diharapkan '%s', diterima '%s'", i+1, v, row[i])
				}
			}
			continue
		}

		parse := lib.ParseStringToArray(row[1], "|")
		var existingEvent int64 = 0
		riskEventCode := parse[0]
		if val, ok := riskIssueMap[riskEventCode]; ok {
			existingEvent = val
		} else {
			newRiskEvents = append(newRiskEvents, models.RiskIssue{
				RiskTypeID:     riskTypeID,
				RiskIssueCode:  riskEventCode,
				RiskIssue:      parse[1],
				Deskripsi:      row[2],
				KategoriRisiko: row[3],
				Status:         true,
				DeleteFlag:     false,
				CreatedAt:      &timeNow,
			})
		}

		evnLv1 := lib.ParseStringToArray(row[4], ";")
		evnLv2 := lib.ParseStringToArray(row[5], ";")
		evnLv3 := lib.ParseStringToArray(row[6], ";")
		for i := range evnLv3 {
			lv1 := normalizeEventCode(evnLv1[i])
			lv2 := normalizeEventCode(evnLv2[i])
			lv3 := normalizeEventCode(evnLv3[i])

			// === cek sudah termapping ===
			if existingEvent > 0 {
				key := fmt.Sprintf("%d|%s|%s|%s", existingEvent, lv1, lv2, lv3)
				if mapEventLvlMap[key] {
					continue
				}
			}

			if parent, ok := eventLv2Parent[lv2]; ok && parent != lv1 {
				continue
			}
			if parent, ok := eventLv3Parent[lv3]; ok && parent != lv2 {
				continue
			}

			lv1ID := strings.ToUpper(lv1)
			lv2ID := strings.ToUpper(lv2)
			lv3ID := strings.ToUpper(lv3)

			// === insert baru ===
			newMapEvent[riskEventCode] = &models.MapEvent{
				EventTypeLv1: lv1ID,
				EventTypeLv2: lv2ID,
				EventTypeLv3: lv3ID,
			}
		}

		riskIssue.logger.Zap.Debug(len(newMapEvent))

		incLv1 := lib.ParseStringToArray(row[7], ";")
		incLv2 := lib.ParseStringToArray(row[8], ";")
		incLv3 := lib.ParseStringToArray(row[9], ";")
		for i := range incLv3 {
			lv1 := normalizeEventCode(incLv1[i])
			lv2 := normalizeEventCode(incLv2[i])
			lv3 := normalizeEventCode(incLv3[i])

			if lv1 == "" {
				continue
			}

			if existingEvent > 0 {
				key := fmt.Sprintf("%d|%s|%s|%s", existingEvent, lv1, lv2, lv3)
				if mapKejadianLvlMap[key] {
					continue
				}
			}

			if !incidentMap[lv1] {
				continue
			}

			if lv2 != "" {
				if !subIncidentMap[lv2] {
					continue
				}
				if parent, ok := subIncidentParentMap[lv2]; ok && parent != lv1 {
					continue
				}
			}

			if lv3 != "" {
				if !subsubIncidentMap[lv3] {
					continue
				}
				if parent, ok := subsubIncidentParentMap[lv3]; ok && parent != lv2 {
					continue
				}
			}

			lv1ID := incLv1[i]
			lv2ID := incLv2[i]
			lv3ID := incLv3[i]

			mapKey := fmt.Sprintf("%s|%s|%s|%s", riskEventCode, lv1ID, lv2ID, lv3ID)
			if _, exists := newMapKejadian[mapKey]; exists {
				continue
			}

			// === simpan mapping baru ===
			newMapKejadian[riskEventCode] = &models.MapKejadian{
				PenyebabKejadianLv1: lv1,
				PenyebabKejadianLv2: lv2,
				PenyebabKejadianLv3: lv3,
			}
		}

		product := lib.ParseStringToArray(row[10], ";")
		for _, v := range product {
			p := normalizeEventCode(v)
			if p == "" {
				continue
			}

			var productId int64 = 0
			if val, ok := productMap[strings.ToLower(p)]; !ok {
				continue
			} else {
				productId = val
			}

			newMapProduct[riskEventCode] = &models.MapProduct{
				Product: productId,
			}
		}

		businessCycleStr := lib.ParseStringToArray(row[17], ";")
		subBusinessCycleStr := lib.ParseStringToArray(row[19], ";")
		processStr := lib.ParseStringToArray(row[21], ";")
		subProcessStr := lib.ParseStringToArray(row[23], ";")
		activityStr := lib.ParseStringToArray(row[25], ";")
		for i := range activityStr {
			key := fmt.Sprintf("%s|%s|%s|%s|%s", businessCycleStr[i], subBusinessCycleStr[i], processStr[i], subProcessStr[i], activityStr[i])
			riskIssue.logger.Zap.Debug(businessProcessMap[strings.ToLower(key)])
			riskIssue.logger.Zap.Debug(key)
			if val, ok := businessProcessMap[strings.ToLower(key)]; ok {
				if existingEvent > 0 {
					mapKey := fmt.Sprintf("%d|%s", existingEvent, key)
					riskIssue.logger.Zap.Debug(mapBusinessProcessMap)
					riskIssue.logger.Zap.Debug(mapKey)
					if mapBusinessProcessMap[strings.ToLower(mapKey)] {
						continue
					}
				}

				mappingBusinessProcess = append(mappingBusinessProcess, dto.MapingBusinessProcessRequest{
					RiskEventID:      riskEventCode,
					ActivityID:       val.ActivityID,
					BusinessCycle:    val.BusinessCycleID,
					SubBusinessCycle: val.SubBusinessCycleID,
					Process:          val.ProcessID,
					SubProcess:       val.SubProcessID,
				})
			}
		}

		controlCodes := lib.ParseStringToArray(row[13], ";")
		indicatorCodes := lib.ParseStringToArray(row[15], ";")
		for _, c := range controlCodes {
			riskIssue.logger.Zap.Debug(c)
			c = strings.ToLower(strings.TrimSpace(c))
			if c == "" {
				continue
			}
			if id, ok := controlIDMap[c]; ok {
				newMapControl[riskEventCode] = &models.MapControl{
					IDControl: id,
					IsChecked: true,
				}
			}
		}

		for _, ind := range indicatorCodes {
			ind = strings.ToLower(strings.TrimSpace(ind))
			if ind == "" {
				continue
			}
			if id, ok := indicatorIDMap[ind]; ok {
				newMapIndicator[riskEventCode] = &models.MapIndicator{
					IDIndicator: id,
					IsChecked:   true,
				}
			}
		}
	}

	tx := riskIssue.db.DB.Begin()

	if len(newRiskEvents) > 0 {
		if err := riskIssue.riskissueRepo.BulkCreateRiskEvent(newRiskEvents, tx); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed bulk create risk events: %v", err)
		}
	}

	issue, err := riskIssue.riskissueRepo.GetAllWithTx(tx)
	if err != nil {
		return fmt.Errorf("failed get risk event: %v", err)
	}

	riskEventMap := make(map[string]int64)
	for _, e := range issue {
		riskEventMap[e.RiskIssueCode] = e.ID
	}

	riskIssue.logger.Zap.Debug(riskEventMap)
	riskIssue.logger.Zap.Debug(controlIDMap)
	newMapEventReq := make([]*models.MapEvent, 0)
	if len(newMapEvent) > 0 {
		for riskIssueCode, v := range newMapEvent {
			riskIssueID, ok := riskEventMap[riskIssueCode]
			if !ok {
				continue
			}

			v.IDRiskIssue = riskIssueID
			newMapEventReq = append(newMapEventReq, v)
		}
		err := riskIssue.mapEvent.BulkCreate(newMapEventReq, tx)
		if err != nil {
			riskIssue.logger.Zap.Error(err)
			tx.Rollback()
			return err
		}
	}

	newMapKejadianReq := make([]*models.MapKejadian, 0)
	if len(newMapKejadian) > 0 {
		for riskIssueCode, v := range newMapKejadian {
			riskIssueID, ok := riskEventMap[riskIssueCode]
			if !ok {
				continue
			}

			v.IDRiskIssue = riskIssueID
			newMapKejadianReq = append(newMapKejadianReq, v)
		}
		err := riskIssue.mapKejadian.BulkCreate(newMapKejadianReq, tx)
		if err != nil {
			riskIssue.logger.Zap.Error(err)
			tx.Rollback()
			return err
		}
	}
	newMapProductReq := make([]*models.MapProduct, 0)
	if len(newMapProduct) > 0 {
		for riskIssueCode, v := range newMapProduct {
			riskIssueID, ok := riskEventMap[riskIssueCode]
			if !ok {
				continue
			}

			v.IDRiskIssue = riskIssueID
			newMapProductReq = append(newMapProductReq, v)
		}
		err := riskIssue.mapProduct.BulkCreate(newMapProductReq, tx)
		if err != nil {
			riskIssue.logger.Zap.Error(err)
			tx.Rollback()
			return err
		}
	}
	newMapControlReq := make([]*models.MapControl, 0)
	if len(newMapControl) > 0 {
		for riskIssueCode, v := range newMapControl {
			riskIssueID, ok := riskEventMap[riskIssueCode]
			if !ok {
				continue
			}

			v.IDRiskIssue = riskIssueID
			newMapControlReq = append(newMapControlReq, v)
		}
		err := riskIssue.mapControl.BulkCreate(newMapControlReq, tx)
		if err != nil {
			riskIssue.logger.Zap.Error(err)
			tx.Rollback()
			return err
		}
	}
	newMapIndicatorReq := make([]*models.MapIndicator, 0)
	if len(newMapIndicator) > 0 {
		for riskIssueCode, v := range newMapIndicator {
			riskIssueID, ok := riskEventMap[riskIssueCode]
			if !ok {
				continue
			}

			v.IDRiskIssue = riskIssueID
			newMapIndicatorReq = append(newMapIndicatorReq, v)
		}
		err := riskIssue.mapIndicator.BulkCreate(newMapIndicatorReq, tx)
		if err != nil {
			riskIssue.logger.Zap.Error(err)
			tx.Rollback()
			return err
		}
	}

	if len(mappingBusinessProcess) > 0 {
		for _, v := range mappingBusinessProcess {
			riskIssueID, ok := riskEventMap[v.RiskEventID]
			if !ok {
				continue
			}
			v.RiskEventID = strconv.Itoa(int(riskIssueID))
		}

		err := riskIssue.arlodsService.BulkCreateMappingBusinessProcess(pernr, mappingBusinessProcess)
		if err != nil {
			riskIssue.logger.Zap.Error(err)
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (riskIssue RiskIssueService) Download(pernr, format string) ([]byte, string, error) {
	data, err := riskIssue.riskissueRepo.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when try to query risk issue: ", err)
		return nil, "", err
	}

	if len(data) == 0 {
		return nil, "", nil
	}

	switch format {
	case "csv":
		return riskIssue.exportCSV(pernr, data)
	case "xlsx":
		return riskIssue.exportExcel(pernr, data)
	case "pdf":
		return riskIssue.exportPDF(pernr, data)
	default:
		return nil, "", fmt.Errorf("unsupported format export file")
	}
}

func (riskIssue RiskIssueService) exportPDF(pernr string, data []models.RiskIssueResponse) ([]byte, string, error) {
	headers := []string{
		"Tipe Risiko",
		"Risk Event",
		"Deskripsi",
		"Kategori Risiko",
		"Map Event - Event Type Lv1",
		"Map Event - Event Type Lv2",
		"Map Event - Event Type Lv3",
		"Map Kejadian - Penyebab Kejadian Lv1",
		"Map Kejadian - Penyebab Kejadian Lv2",
		"Map Kejadian - Penyebab Kejadian Lv3",
		"Map Product - Product",
		"Likelihood",
		"Impact",
		"Code Control",
		"Risk Control",
		"Code Indicator",
		"Risk Indicator",
		"Code Business Cycle",
		"Business Cycle",
		"Code Sub Business Cycle",
		"Sub Business Cycle",
		"Code Process",
		"Process",
		"Code Sub Process",
		"Sub Process",
		"Code Activity",
		"Activity",
	}

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetAutoPageBreak(false, 10)
	pdf.SetMargins(10, 10, 10) // margin kiri, atas, kanan
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, "Risk Control Report", "", 1, "C", false, 0, "")

	colWidths := []float64{
		20, 20, 35, 25, 25,
		25, 25, 25, 25, 25,
		25, 25, 25, 25, 25,
		25, 25, 25, 25, 25,
		25, 25, 25, 25, 25,
		25, 25,
	}

	printHeader := func() {
		pdf.SetFillColor(200, 200, 200)
		pdf.SetFont("Arial", "B", 10)

		lineHeight := 5.0

		// Hitung tinggi maksimum header (berdasarkan wrapping)
		maxHeight := 0.0
		for i, h := range headers {
			lines := pdf.SplitLines([]byte(h), colWidths[i])
			hh := float64(len(lines)) * lineHeight
			if hh > maxHeight {
				maxHeight = hh
			}
		}

		xStart := pdf.GetX()
		yStart := pdf.GetY()

		// Cetak setiap header cell
		for i, h := range headers {
			x := pdf.GetX()
			y := pdf.GetY()

			// Gambar border kotak
			pdf.Rect(x, y, colWidths[i], maxHeight, "DF") // DF = fill + border

			// Cetak text dengan wrapping di tengah vertikal
			lines := pdf.SplitLines([]byte(h), colWidths[i])
			textHeight := float64(len(lines)) * lineHeight
			yOffset := (maxHeight - textHeight) / 2

			pdf.SetXY(x, y+yOffset)
			pdf.MultiCell(colWidths[i], lineHeight, h, "", "C", false)
			pdf.SetXY(x+colWidths[i], yStart)
		}

		pdf.SetXY(xStart, yStart+maxHeight)
		pdf.SetFont("Arial", "", 9)
	}

	printHeader()
	_, pageHeight := pdf.GetPageSize()
	marginBottom := 15.0

	getRowHeight := func(row []string) float64 {
		maxHeight := 0.0
		lineHeight := 5.0
		for i, txt := range row {
			lines := pdf.SplitLines([]byte(txt), colWidths[i])
			h := float64(len(lines)) * lineHeight
			if h > maxHeight {
				maxHeight = h
			}
		}
		return maxHeight
	}

	riskType, err := riskIssue.riskType.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping risk type: ", err)
		return nil, "", err
	}

	riskTypeMap := make(map[int64]string, 0)
	for _, v := range riskType {
		riskTypeMap[v.ID] = fmt.Sprintf("%s - %s", v.RiskTypeCode, v.RiskType)
	}

	for _, v := range data {
		riskType := ""
		if v, ok := riskTypeMap[v.RiskTypeID]; ok {
			riskType = v
		}

		likelihood := ""
		impact := ""
		if v.Likelihood != nil {
			likelihood = *v.Likelihood
		}
		if v.Impact != nil {
			impact = *v.Impact
		}

		business, err := riskIssue.arlodsService.GetMappingEventBusinessProess(pernr, int(v.ID))
		if err != nil {
			riskIssue.logger.Zap.Error("Errored when get mapping: ", err)
			return nil, "", err
		}

		event, err := riskIssue.mapEvent.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		var (
			eventTypeLv1s []string
			eventTypeLv2s []string
			eventTypeLv3s []string
		)

		if len(event) > 0 {
			for _, event := range event {

				if event.EventTypeLv1 != "" || event.EventTypeLv1Desc != "" {
					eventTypeLv1s = append(
						eventTypeLv1s,
						fmt.Sprintf("%s - %s", event.EventTypeLv1, event.EventTypeLv1Desc),
					)
				}

				if event.EventTypeLv2 != "" || event.EventTypeLv2Desc != "" {
					eventTypeLv2s = append(
						eventTypeLv2s,
						fmt.Sprintf("%s - %s", event.EventTypeLv2, event.EventTypeLv2Desc),
					)
				}

				if event.EventTypeLv3 != "" || event.EventTypeLv3Desc != "" {
					eventTypeLv3s = append(
						eventTypeLv3s,
						fmt.Sprintf("%s - %s", event.EventTypeLv3, event.EventTypeLv3Desc),
					)
				}
			}
		}

		var (
			penyebabLv1s []string
			penyebabLv2s []string
			penyebabLv3s []string
		)

		kejadian, err := riskIssue.mapKejadian.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		if len(kejadian) > 0 {
			for _, k := range kejadian {

				if k.PenyebabKejadianLv1 != "" || k.PenyebabKejadianLv1Desc != "" {
					penyebabLv1s = append(
						penyebabLv1s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv1,
							k.PenyebabKejadianLv1Desc,
						),
					)
				}

				if k.PenyebabKejadianLv2 != "" || k.PenyebabKejadianLv2Desc != "" {
					penyebabLv2s = append(
						penyebabLv2s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv2,
							k.PenyebabKejadianLv2Desc,
						),
					)
				}

				if k.PenyebabKejadianLv3 != "" || k.PenyebabKejadianLv3Desc != "" {
					penyebabLv3s = append(
						penyebabLv3s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv3,
							k.PenyebabKejadianLv3Desc,
						),
					)
				}
			}
		}

		product, err := riskIssue.mapProduct.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		var productNames []string
		if len(product) > 0 {
			for _, v := range product {
				productNames = append(productNames, v.ProductDesc)
			}
		}

		var (
			businessCycles        []string
			businessCycleNames    []string
			subBusinessCycles     []string
			subBusinessCycleNames []string
			activityIDs           []string
			activityNames         []string
			processIDs            []string
			processNames          []string
			subProcessIDs         []string
			subProcessNames       []string
		)

		if len(business.Data.List) > 0 {
			for _, v := range business.Data.List {
				businessCycles = append(businessCycles, v.BusinessCode)
				businessCycleNames = append(businessCycleNames, v.BusinessCycleName)
				subBusinessCycles = append(subBusinessCycles, v.SubBusinessCode)
				subBusinessCycleNames = append(subBusinessCycleNames, v.SubBusinessCycleName)
				activityIDs = append(activityIDs, v.ActivityCode)
				activityNames = append(activityNames, v.ActivityName)
				processIDs = append(processIDs, v.ProcessCode)
				processNames = append(processNames, v.ProcessName)
				subProcessIDs = append(subProcessIDs, v.SubProcessCode)
				subProcessNames = append(subProcessNames, v.SubProcessName)
			}
		}

		var (
			controlCodes []string
			controlNames []string
		)

		control, err := riskIssue.mapControl.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		if len(control) > 0 {
			for _, c := range control {

				if strings.TrimSpace(c.Kode) != "" {
					controlCodes = append(controlCodes, c.Kode)
				}

				if strings.TrimSpace(c.RiskControl) != "" {
					controlNames = append(controlNames, c.RiskControl)
				}
			}
		}

		var (
			indicatorCodes []string
			indicatorNames []string
		)

		indicator, err := riskIssue.mapIndicator.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		if len(indicator) > 0 {
			for _, i := range indicator {

				if strings.TrimSpace(i.Kode) != "" {
					indicatorCodes = append(indicatorCodes, i.Kode)
				}

				if strings.TrimSpace(i.RiskIndicator) != "" {
					indicatorNames = append(indicatorNames, i.RiskIndicator)
				}
			}
		}

		row := []string{
			riskType,
			fmt.Sprintf("%s - %s", v.RiskIssueCode, v.RiskIssue),
			v.Deskripsi,
			v.KategoriRisiko,
			strings.Join(eventTypeLv1s, ";"),
			strings.Join(eventTypeLv2s, ";"),
			strings.Join(eventTypeLv3s, ";"),
			strings.Join(penyebabLv1s, ";"),
			strings.Join(penyebabLv2s, ";"),
			strings.Join(penyebabLv3s, ";"),
			strings.Join(productNames, ";"),
			likelihood,
			impact,
			strings.Join(controlCodes, ";"),
			strings.Join(controlNames, ";"),
			strings.Join(indicatorCodes, ";"),
			strings.Join(indicatorNames, ";"),
			strings.Join(businessCycles, ";"),
			strings.Join(businessCycleNames, ";"),
			strings.Join(subBusinessCycles, ";"),
			strings.Join(subBusinessCycleNames, ";"),
			strings.Join(processIDs, ";"),
			strings.Join(processNames, ";"),
			strings.Join(subProcessIDs, ";"),
			strings.Join(subProcessNames, ";"),
			strings.Join(activityIDs, ";"),
			strings.Join(activityNames, ";"),
		}

		rowHeight := getRowHeight(row)
		xStart := pdf.GetX()
		yStart := pdf.GetY()

		// Check page break
		if yStart+rowHeight+marginBottom > pageHeight {
			pdf.AddPage()
			printHeader()
			xStart = pdf.GetX()
			yStart = pdf.GetY()
		}

		// Print each cell with MultiCell and border
		for i, txt := range row {
			x := pdf.GetX()
			y := pdf.GetY()

			pdf.Rect(x, y, colWidths[i], rowHeight, "D")
			pdf.MultiCell(colWidths[i], 5, txt, "", "L", false)
			pdf.SetXY(x+colWidths[i], yStart)
		}
		pdf.SetXY(xStart, yStart+rowHeight)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate PDF: %w", err)
	}

	fileName := fmt.Sprintf("risk_event_%s.pdf", time.Now().Format("20060102_150405"))

	return buf.Bytes(), fileName, nil
}

func (riskIssue RiskIssueService) exportExcel(pernr string, data []models.RiskIssueResponse) ([]byte, string, error) {
	headers := []string{
		"Tipe Risiko",
		"Risk Event",
		"Deskripsi",
		"Kategori Risiko",
		"Map Event - Event Type Lv1",
		"Map Event - Event Type Lv2",
		"Map Event - Event Type Lv3",
		"Map Kejadian - Penyebab Kejadian Lv1",
		"Map Kejadian - Penyebab Kejadian Lv2",
		"Map Kejadian - Penyebab Kejadian Lv3",
		"Map Product - Product",
		"Likelihood",
		"Impact",
		"Code Control",
		"Risk Control",
		"Code Indicator",
		"Risk Indicator",
		"Code Business Cycle",
		"Business Cycle",
		"Code Sub Business Cycle",
		"Sub Business Cycle",
		"Code Process",
		"Process",
		"Code Sub Process",
		"Sub Process",
		"Code Activity",
		"Activity",
	}
	f := excelize.NewFile()
	sheet := "risk-issue"

	f.SetSheetName("Sheet1", sheet)

	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			return nil, "", fmt.Errorf("failed to set cell: %w", err)
		}
	}

	riskType, err := riskIssue.riskType.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping risk type: ", err)
		return nil, "", err
	}

	riskTypeMap := make(map[int64]string, 0)
	for _, v := range riskType {
		riskTypeMap[v.ID] = fmt.Sprintf("%s - %s", v.RiskTypeCode, v.RiskType)
	}

	for idx, v := range data {
		riskType := ""
		if v, ok := riskTypeMap[v.RiskTypeID]; ok {
			riskType = v
		}

		business, err := riskIssue.arlodsService.GetMappingEventBusinessProess(pernr, int(v.ID))
		if err != nil {
			riskIssue.logger.Zap.Error("Errored when get mapping: ", err)
			return nil, "", err
		}

		event, err := riskIssue.mapEvent.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		var (
			eventTypeLv1s []string
			eventTypeLv2s []string
			eventTypeLv3s []string
		)

		if len(event) > 0 {
			for _, event := range event {

				if event.EventTypeLv1 != "" || event.EventTypeLv1Desc != "" {
					eventTypeLv1s = append(
						eventTypeLv1s,
						fmt.Sprintf("%s - %s", event.EventTypeLv1, event.EventTypeLv1Desc),
					)
				}

				if event.EventTypeLv2 != "" || event.EventTypeLv2Desc != "" {
					eventTypeLv2s = append(
						eventTypeLv2s,
						fmt.Sprintf("%s - %s", event.EventTypeLv2, event.EventTypeLv2Desc),
					)
				}

				if event.EventTypeLv3 != "" || event.EventTypeLv3Desc != "" {
					eventTypeLv3s = append(
						eventTypeLv3s,
						fmt.Sprintf("%s - %s", event.EventTypeLv3, event.EventTypeLv3Desc),
					)
				}
			}
		}

		var (
			penyebabLv1s []string
			penyebabLv2s []string
			penyebabLv3s []string
		)

		kejadian, err := riskIssue.mapKejadian.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		if len(kejadian) > 0 {
			for _, k := range kejadian {

				if k.PenyebabKejadianLv1 != "" || k.PenyebabKejadianLv1Desc != "" {
					penyebabLv1s = append(
						penyebabLv1s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv1,
							k.PenyebabKejadianLv1Desc,
						),
					)
				}

				if k.PenyebabKejadianLv2 != "" || k.PenyebabKejadianLv2Desc != "" {
					penyebabLv2s = append(
						penyebabLv2s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv2,
							k.PenyebabKejadianLv2Desc,
						),
					)
				}

				if k.PenyebabKejadianLv3 != "" || k.PenyebabKejadianLv3Desc != "" {
					penyebabLv3s = append(
						penyebabLv3s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv3,
							k.PenyebabKejadianLv3Desc,
						),
					)
				}
			}
		}

		product, err := riskIssue.mapProduct.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		var productNames []string
		if len(product) > 0 {
			for _, v := range product {
				productNames = append(productNames, v.ProductDesc)
			}
		}

		var (
			businessCycles        []string
			businessCycleNames    []string
			subBusinessCycles     []string
			subBusinessCycleNames []string
			activityIDs           []string
			activityNames         []string
			processIDs            []string
			processNames          []string
			subProcessIDs         []string
			subProcessNames       []string
		)

		if len(business.Data.List) > 0 {
			for _, v := range business.Data.List {
				businessCycles = append(businessCycles, v.BusinessCode)
				businessCycleNames = append(businessCycleNames, v.BusinessCycleName)
				subBusinessCycles = append(subBusinessCycles, v.SubBusinessCode)
				subBusinessCycleNames = append(subBusinessCycleNames, v.SubBusinessCycleName)
				activityIDs = append(activityIDs, v.ActivityCode)
				activityNames = append(activityNames, v.ActivityName)
				processIDs = append(processIDs, v.ProcessCode)
				processNames = append(processNames, v.ProcessName)
				subProcessIDs = append(subProcessIDs, v.SubProcessCode)
				subProcessNames = append(subProcessNames, v.SubProcessName)
			}
		}

		var (
			controlCodes []string
			controlNames []string
		)

		control, err := riskIssue.mapControl.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		if len(control) > 0 {
			for _, c := range control {

				if strings.TrimSpace(c.Kode) != "" {
					controlCodes = append(controlCodes, c.Kode)
				}

				if strings.TrimSpace(c.RiskControl) != "" {
					controlNames = append(controlNames, c.RiskControl)
				}
			}
		}

		var (
			indicatorCodes []string
			indicatorNames []string
		)

		indicator, err := riskIssue.mapIndicator.GetOneDataByID(v.ID)
		if err != nil {
			riskIssue.logger.Zap.Debug(err)
			continue
		}

		if len(indicator) > 0 {
			for _, i := range indicator {

				if strings.TrimSpace(i.Kode) != "" {
					indicatorCodes = append(indicatorCodes, i.Kode)
				}

				if strings.TrimSpace(i.RiskIndicator) != "" {
					indicatorNames = append(indicatorNames, i.RiskIndicator)
				}
			}
		}

		likelihood := ""
		impact := ""
		if v.Likelihood != nil {
			likelihood = *v.Likelihood
		}
		if v.Impact != nil {
			impact = *v.Impact
		}

		f.SetCellValue(sheet, fmt.Sprintf("A%d", idx+2), riskType)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", idx+2), fmt.Sprintf("%s - %s", v.RiskIssueCode, v.RiskIssue))
		f.SetCellValue(sheet, fmt.Sprintf("C%d", idx+2), v.Deskripsi)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", idx+2), v.KategoriRisiko)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", idx+2), strings.Join(eventTypeLv1s, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", idx+2), strings.Join(eventTypeLv2s, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", idx+2), strings.Join(eventTypeLv3s, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("H%d", idx+2), strings.Join(penyebabLv1s, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("I%d", idx+2), strings.Join(penyebabLv2s, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("J%d", idx+2), strings.Join(penyebabLv3s, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("K%d", idx+2), strings.Join(productNames, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("L%d", idx+2), likelihood)
		f.SetCellValue(sheet, fmt.Sprintf("M%d", idx+2), impact)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", idx+2), strings.Join(controlCodes, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("O%d", idx+2), strings.Join(controlNames, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("P%d", idx+2), strings.Join(indicatorCodes, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", idx+2), strings.Join(indicatorNames, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("R%d", idx+2), strings.Join(businessCycles, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("S%d", idx+2), strings.Join(businessCycleNames, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("T%d", idx+2), strings.Join(subBusinessCycles, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("U%d", idx+2), strings.Join(subBusinessCycleNames, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("V%d", idx+2), strings.Join(processIDs, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("W%d", idx+2), strings.Join(processNames, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("XA%d", idx+2), strings.Join(subProcessIDs, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("Y%d", idx+2), strings.Join(subProcessNames, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("Z%d", idx+2), strings.Join(activityIDs, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("AA%d", idx+2), strings.Join(activityNames, ";"))
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", fmt.Errorf("failed to write excel file: %w", err)
	}

	fileName := fmt.Sprintf("risk_event_%s.xlsx", time.Now().Format("20060102_150405"))

	return buf.Bytes(), fileName, nil
}

func (riskissue RiskIssueService) exportCSV(pernr string, data []models.RiskIssueResponse) ([]byte, string, error) {
	headers := []string{
		"Tipe Risiko",
		"Risk Event",
		"Deskripsi",
		"Kategori Risiko",
		"Map Event - Event Type Lv1",
		"Map Event - Event Type Lv2",
		"Map Event - Event Type Lv3",
		"Map Kejadian - Penyebab Kejadian Lv1",
		"Map Kejadian - Penyebab Kejadian Lv2",
		"Map Kejadian - Penyebab Kejadian Lv3",
		"Map Product - Product",
		"Likelihood",
		"Impact",
		"Code Control",
		"Risk Control",
		"Code Indicator",
		"Risk Indicator",
		"Code Business Cycle",
		"Business Cycle",
		"Code Sub Business Cycle",
		"Sub Business Cycle",
		"Code Process",
		"Process",
		"Code Sub Process",
		"Sub Process",
		"Code Activity",
		"Activity",
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	if err := writer.Write(headers); err != nil {
		return nil, "", fmt.Errorf("failed to write csv header: %w", err)
	}

	riskType, err := riskissue.riskType.GetAll()
	if err != nil {
		riskissue.logger.Zap.Error("Errored when get mapping risk type: ", err)
		return nil, "", err
	}

	riskTypeMap := make(map[int64]string, 0)
	for _, v := range riskType {
		riskTypeMap[v.ID] = fmt.Sprintf("%s - %s", v.RiskTypeCode, v.RiskType)
	}

	for _, d := range data {
		riskType := ""
		if v, ok := riskTypeMap[d.RiskTypeID]; ok {
			riskType = v
		}

		likelihood := ""
		impact := ""
		if d.Likelihood != nil {
			likelihood = *d.Likelihood
		}
		if d.Impact != nil {
			impact = *d.Impact
		}

		business, err := riskissue.arlodsService.GetMappingEventBusinessProess(pernr, int(d.ID))
		if err != nil {
			riskissue.logger.Zap.Error("Errored when get mapping: ", err)
			return nil, "", err
		}

		event, err := riskissue.mapEvent.GetOneDataByID(d.ID)
		if err != nil {
			riskissue.logger.Zap.Debug(err)
			continue
		}

		var (
			eventTypeLv1s []string
			eventTypeLv2s []string
			eventTypeLv3s []string
		)

		if len(event) > 0 {
			for _, event := range event {

				if event.EventTypeLv1 != "" || event.EventTypeLv1Desc != "" {
					eventTypeLv1s = append(
						eventTypeLv1s,
						fmt.Sprintf("%s - %s", event.EventTypeLv1, event.EventTypeLv1Desc),
					)
				}

				if event.EventTypeLv2 != "" || event.EventTypeLv2Desc != "" {
					eventTypeLv2s = append(
						eventTypeLv2s,
						fmt.Sprintf("%s - %s", event.EventTypeLv2, event.EventTypeLv2Desc),
					)
				}

				if event.EventTypeLv3 != "" || event.EventTypeLv3Desc != "" {
					eventTypeLv3s = append(
						eventTypeLv3s,
						fmt.Sprintf("%s - %s", event.EventTypeLv3, event.EventTypeLv3Desc),
					)
				}
			}
		}

		var (
			penyebabLv1s []string
			penyebabLv2s []string
			penyebabLv3s []string
		)

		kejadian, err := riskissue.mapKejadian.GetOneDataByID(d.ID)
		if err != nil {
			riskissue.logger.Zap.Debug(err)
			continue
		}

		if len(kejadian) > 0 {
			for _, k := range kejadian {

				if k.PenyebabKejadianLv1 != "" || k.PenyebabKejadianLv1Desc != "" {
					penyebabLv1s = append(
						penyebabLv1s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv1,
							k.PenyebabKejadianLv1Desc,
						),
					)
				}

				if k.PenyebabKejadianLv2 != "" || k.PenyebabKejadianLv2Desc != "" {
					penyebabLv2s = append(
						penyebabLv2s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv2,
							k.PenyebabKejadianLv2Desc,
						),
					)
				}

				if k.PenyebabKejadianLv3 != "" || k.PenyebabKejadianLv3Desc != "" {
					penyebabLv3s = append(
						penyebabLv3s,
						fmt.Sprintf(
							"%s - %s",
							k.PenyebabKejadianLv3,
							k.PenyebabKejadianLv3Desc,
						),
					)
				}
			}
		}

		product, err := riskissue.mapProduct.GetOneDataByID(d.ID)
		if err != nil {
			riskissue.logger.Zap.Debug(err)
			continue
		}

		var productNames []string
		if len(product) > 0 {
			for _, v := range product {
				productNames = append(productNames, v.ProductDesc)
			}
		}

		var (
			businessCycles        []string
			businessCycleNames    []string
			subBusinessCycles     []string
			subBusinessCycleNames []string
			activityIDs           []string
			activityNames         []string
			processIDs            []string
			processNames          []string
			subProcessIDs         []string
			subProcessNames       []string
		)

		if len(business.Data.List) > 0 {
			for _, v := range business.Data.List {
				businessCycles = append(businessCycles, v.BusinessCode)
				businessCycleNames = append(businessCycleNames, v.BusinessCycleName)
				subBusinessCycles = append(subBusinessCycles, v.SubBusinessCode)
				subBusinessCycleNames = append(subBusinessCycleNames, v.SubBusinessCycleName)
				activityIDs = append(activityIDs, v.ActivityCode)
				activityNames = append(activityNames, v.ActivityName)
				processIDs = append(processIDs, v.ProcessCode)
				processNames = append(processNames, v.ProcessName)
				subProcessIDs = append(subProcessIDs, v.SubProcessCode)
				subProcessNames = append(subProcessNames, v.SubProcessName)
			}
		}

		var (
			controlCodes []string
			controlNames []string
		)

		control, err := riskissue.mapControl.GetOneDataByID(d.ID)
		if err != nil {
			riskissue.logger.Zap.Debug(err)
			continue
		}

		if len(control) > 0 {
			for _, c := range control {

				if strings.TrimSpace(c.Kode) != "" {
					controlCodes = append(controlCodes, c.Kode)
				}

				if strings.TrimSpace(c.RiskControl) != "" {
					controlNames = append(controlNames, c.RiskControl)
				}
			}
		}

		var (
			indicatorCodes []string
			indicatorNames []string
		)

		indicator, err := riskissue.mapIndicator.GetOneDataByID(d.ID)
		if err != nil {
			riskissue.logger.Zap.Debug(err)
			continue
		}

		if len(indicator) > 0 {
			for _, i := range indicator {

				if strings.TrimSpace(i.Kode) != "" {
					indicatorCodes = append(indicatorCodes, i.Kode)
				}

				if strings.TrimSpace(i.RiskIndicator) != "" {
					indicatorNames = append(indicatorNames, i.RiskIndicator)
				}
			}
		}

		row := []string{
			riskType,
			fmt.Sprintf("%s - %s", d.RiskIssueCode, d.RiskIssue),
			d.Deskripsi,
			d.KategoriRisiko,
			strings.Join(eventTypeLv1s, ";"),
			strings.Join(eventTypeLv2s, ";"),
			strings.Join(eventTypeLv3s, ";"),
			strings.Join(penyebabLv1s, ";"),
			strings.Join(penyebabLv2s, ";"),
			strings.Join(penyebabLv3s, ";"),
			strings.Join(productNames, ";"),
			likelihood,
			impact,
			strings.Join(controlCodes, ";"),
			strings.Join(controlNames, ";"),
			strings.Join(indicatorCodes, ";"),
			strings.Join(indicatorNames, ";"),
			strings.Join(businessCycles, ";"),
			strings.Join(businessCycleNames, ";"),
			strings.Join(subBusinessCycles, ";"),
			strings.Join(subBusinessCycleNames, ";"),
			strings.Join(processIDs, ";"),
			strings.Join(processNames, ";"),
			strings.Join(subProcessIDs, ";"),
			strings.Join(subProcessNames, ";"),
			strings.Join(activityIDs, ";"),
			strings.Join(activityNames, ";"),
		}

		if err := writer.Write(row); err != nil {
			return nil, "", fmt.Errorf("failed to write csv row: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, "", fmt.Errorf("failed to flush csv data: %w", err)
	}

	fileName := fmt.Sprintf("risk_issue_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func extractBusinessProcessNodes(
	bc dto.BusinessProcessMap,
	cache map[string]dto.BusinessProcessNode,
) {

	// Loop SBC (sub business cycle)
	for _, sbc := range bc.BusinessProcessMap {

		// Loop Process
		for _, p := range sbc.BusinessProcessMap {

			// Loop SubProcess
			for _, sp := range p.BusinessProcessMap {

				// Loop Activity
				for _, ac := range sp.BusinessProcessMap {
					// Simpan ke cache (key: activity code lowercase)
					cache[strings.ToLower(ac.Code)] = dto.BusinessProcessNode{
						ActivityCode:         ac.Code,
						ActivityName:         ac.Name,
						SubProcessCode:       sp.Code,
						SubProcessName:       sp.Name,
						ProcessCode:          p.Code,
						ProcessName:          p.Name,
						SubBusinessCycleCode: sbc.Code,
						SubBusinessCycleName: sbc.Name,
						BusinessCycleCode:    bc.Code,
						BusinessCycleName:    bc.Name,
						ActivityID:           ac.ID,
						SubProcessID:         sp.ID,
						ProcessID:            p.ID,
						SubBusinessCycleID:   sbc.ID,
						BusinessCycleID:      bc.ID,
					}
				}
			}
		}
	}
}

func (riskIssue RiskIssueService) GetRiskCategories(id []int64) ([]string, error) {
	data, err := riskIssue.riskissueRepo.GetRiskCategories(id)
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when try to query risk categories: ", err)
		return nil, err
	}

	return data, nil
}

func normalizeEventCode(s string) string {
	s = strings.TrimSpace(s)

	// 1. Format aman: "CODE - DESC"
	if idx := strings.Index(s, " - "); idx != -1 {
		s = s[:idx]
		return strings.ToLower(strings.TrimSpace(s))
	}

	// 2. Fallback: "CODE-DESC" (ambil sebelum dash pertama yang diikuti huruf/spasi)
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			// pastikan setelah '-' bukan angka (biar EVT001-1 tetap utuh)
			if i+1 < len(s) && (s[i+1] == ' ' || (s[i+1] >= 'A' && s[i+1] <= 'Z') || (s[i+1] >= 'a' && s[i+1] <= 'z')) {
				s = s[:i]
				break
			}
		}
	}

	return strings.ToLower(strings.TrimSpace(s))
}
