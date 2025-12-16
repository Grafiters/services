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
	riskIssueMap := make(map[string]string)
	for _, e := range issues {
		riskIssueMap[e.RiskIssueCode] = strconv.FormatInt(e.ID, 10)
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

	cacheBusinessProcess := make(map[string]dto.BusinessProcessNode, 0)
	eventIDs := make([]string, 0)
	mapping := make(map[string]dto.MappingEvent, 0)

	for i, row := range data {
		if i == 0 {
			continue
		}
		riskType := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(row[0], "-")))
		eventIDs = append(eventIDs, riskType)
	}

	mappingResp, err := riskIssue.arlodsService.GetAllMappingRiskEvent(pernr, eventIDs)
	if err != nil {
		return dto.PreviewFileImport[[27]string]{}, fmt.Errorf("gagal mengambil mapping event: %v", err)
	}

	// Buat map riskEventID -> DetailMappingEvent
	mappingByEventID := make(map[string]dto.DetailMappingEvent)
	for _, item := range mappingResp.Data {
		mappingByEventID[strings.ToLower(item.RiskEventID)] = item
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
		riskEventCode := parse[0]
		if _, ok := riskIssueMap[riskEventCode]; ok {
			validation += fmt.Sprintf("Risk Event Sudah terdaftar: %s", row[1])
		}

		// ================================
		// STEP 1 — PERSIAPAN FIELD BP
		// ================================
		businessCycleStr := row[17]
		subBusinessCycleStr := row[19]
		processStr := row[21]
		subProcessStr := row[23]
		activityStr := row[25]

		activities := lib.ParseStringToArray(activityStr, ";")
		// ================================
		// STEP 2 — CACHE BUSINESS PROCESS
		// ================================
		for _, act := range activities {
			resp, err := riskIssue.arlodsService.GetHirearcyBusinessProcess(pernr, act)
			if err != nil {
				return dto.PreviewFileImport[[27]string]{}, err
			}
			if len(resp.Data.List) == 0 {
				validation += fmt.Sprintf("Activity '%s' tidak ditemukan di hierarchy; ", act)
				continue
			}
			for _, bc := range resp.Data.List {
				extractBusinessProcessNodes(bc, cacheBusinessProcess)
			}
		}
		validation += ValidateBPWithMultiValue(
			businessCycleStr,
			subBusinessCycleStr,
			processStr,
			subProcessStr,
			activityStr,
			cacheBusinessProcess,
		)

		// ================================
		// STEP 3 — PARSE FIELD LAINNYA
		// ================================
		riskType := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(row[0], "-")))
		eventLv1 := lib.ParseStringToArray(row[4], ";")
		eventLv2 := lib.ParseStringToArray(row[5], ";")
		eventLv3 := lib.ParseStringToArray(row[6], ";")
		incident := lib.ParseStringToArray(row[7], ";")
		subIncident := lib.ParseStringToArray(row[8], ";")
		subsubIncident := lib.ParseStringToArray(row[9], ";")
		product := lib.ParseStringToArray(row[10], ";")

		if _, ok := riskTypeMap[riskType]; !ok {
			validation += fmt.Sprintf("Risk Type tidak terdaftar: %s; ", riskType)
		}
		eventIDs = append(eventIDs, riskType)

		evnLv1map, evnLv2map, evnLv3map := []string{}, []string{}, []string{}
		for i := range eventLv1 {
			lv1 := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(eventLv1[i], "-")))
			lv2 := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(eventLv2[i], "-")))
			lv3 := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(eventLv3[i], "-")))

			if !eventTypeLv1Map[lv1] {
				validation += fmt.Sprintf("Event LV1 tidak terdaftar: %s; ", row[4])
			} else {
				evnLv1map = append(evnLv1map, lv1)
			}
			if !eventTypelv2Map[lv2] {
				validation += fmt.Sprintf("Event LV2 tidak terdaftar: %s; ", row[5])
			} else {
				evnLv2map = append(evnLv2map, lv2)
			}
			if !eventTypelv3Map[lv3] {
				validation += fmt.Sprintf("Event LV3 tidak terdaftar: %s; ", row[6])
			} else {
				evnLv3map = append(evnLv3map, lv3)
			}

			if eventLv2Parent[lv2] != lv1 {
				validation += fmt.Sprintf("Event LV2 '%s' tidak terkait LV1 '%s'; ", row[5], row[4])
			}
			if eventLv3Parent[lv3] != lv2 {
				validation += fmt.Sprintf("Event LV3 '%s' tidak terkait LV2 '%s'; ", row[6], row[5])
			}
		}

		event.EventLV1 = evnLv1map
		event.EventLv2 = evnLv2map
		event.EventLv3 = evnLv3map

		// Incident
		incLv1Map, incLv2Map, incLv3Map := []string{}, []string{}, []string{}
		for i := range incident {
			lv1 := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(incident[i], "-")))
			lv2 := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(subIncident[i], "-")))
			lv3 := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(subsubIncident[i], "-")))

			if !incidentMap[lv1] {
				validation += fmt.Sprintf("Incident LV1 tidak terdaftar: %s; ", row[7])
			} else {
				incLv1Map = append(incLv1Map, lv1)
			}
			if lv2 != "" {
				if !incidentMap[lv2] {
					validation += fmt.Sprintf("Incident LV2 tidak terdaftar: %s; ", row[8])
				} else {
					incLv2Map = append(incLv2Map, lv2)
					if subIncidentParentMap[lv2] != lv1 {
						validation += fmt.Sprintf("Incident LV2 '%s' tidak terkait LV1 '%s'; ", row[8], row[7])
					}
				}
			}
			if lv3 != "" {
				if !incidentMap[lv3] {
					validation += fmt.Sprintf("Incident LV3 tidak terdaftar: %s; ", row[9])
				} else {
					incLv3Map = append(incLv3Map, lv3)
					if subsubIncidentParentMap[lv3] != lv2 {
						validation += fmt.Sprintf("Incident LV3 '%s' tidak terkait LV2 '%s'; ", row[9], row[8])
					}
				}
			}
		}
		event.Incident = incLv1Map
		event.SubIncident = incLv2Map
		event.SubSubIncident = incLv3Map

		// Product
		productMapped := []string{}
		for _, p := range product {
			p = strings.ToLower(strings.TrimSpace(p))
			if p == "" {
				continue
			}
			if !productMap[p] {
				validation += fmt.Sprintf("Product tidak terdaftar: %s; ", p)
			} else {
				productMapped = append(productMapped, p)
			}
		}
		event.ProductIDs = append(event.ProductIDs, productMapped...)

		// ================================
		// STEP 4 — VALIDASI TERHADAP MAPPING EXISTING
		// ================================
		if mappedEvent, exists := mappingByEventID[riskType]; exists {
			// MappingEvent
			for _, m := range mappedEvent.MappingDetail.MappingEvent {
				for _, v := range evnLv1map {
					if v == "" {
						continue
					}
					if v == strings.ToLower(m.EventTypeLvl1) {
						validation += fmt.Sprintf("Event LV1 '%s' sudah termapping; ", v)
					}
				}
				for _, v := range evnLv2map {
					if v == "" {
						continue
					}
					if v == strings.ToLower(m.EventTypeLvl2) {
						validation += fmt.Sprintf("Event LV2 '%s' sudah termapping; ", v)
					}
				}
				for _, v := range evnLv3map {
					if v == "" {
						continue
					}
					if v == strings.ToLower(m.EventTypeLvl3) {
						validation += fmt.Sprintf("Event LV3 '%s' sudah termapping; ", v)
					}
				}
			}
			// MappingCause
			for _, m := range mappedEvent.MappingDetail.MappingCause {
				for _, v := range incLv1Map {
					if v == "" {
						continue
					}
					if v == strings.ToLower(m.Incident) {
						validation += fmt.Sprintf("Incident LV1 '%s' sudah termapping; ", v)
					}
				}
				for _, v := range incLv2Map {
					if v == "" {
						continue
					}
					if v == strings.ToLower(m.SubIncident) {
						validation += fmt.Sprintf("Incident LV2 '%s' sudah termapping; ", v)
					}
				}
				for _, v := range incLv3Map {
					if v == "" {
						continue
					}
					if v == strings.ToLower(m.SubSubIncident) {
						validation += fmt.Sprintf("Incident LV3 '%s' sudah termapping; ", v)
					}
				}
			}
			// MappingProduct
			for _, m := range mappedEvent.MappingDetail.MappingProduct {
				for _, v := range productMapped {
					if v == "" {
						continue
					}
					if v == strings.ToLower(m.ProductID) {
						validation += fmt.Sprintf("Product '%s' sudah termapping; ", v)
					}
				}
			}
		}

		mapping[riskType] = event

		// ================================
		// STEP 5 — COPY DATA
		// ================================
		riskIssue.logger.Zap.Debug(row)
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
	riskIssueMap := make(map[string]string)
	for _, e := range issues {
		riskIssueMap[e.RiskIssueCode] = strconv.FormatInt(e.ID, 10)
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
	eventLv1Map := make(map[string]int64)
	for _, e := range eventLv1 {
		eventLv1Map[strings.ToLower(e.KodeEventType)] = e.ID
	}

	eventLv2, _ := riskIssue.eventTypeLvl2.GetAll()
	eventLv2Map := make(map[string]int64)
	eventLv2Parent := make(map[string]string)
	for _, e := range eventLv2 {
		key := strings.ToLower(e.KodeEventTypeLv2)
		eventLv2Map[key] = e.ID
		eventLv2Parent[key] = strings.ToLower(e.IDEventTypeLv1)
	}

	eventLv3, _ := riskIssue.eventTypeLvl3.GetAll()
	eventLv3Map := make(map[string]int64)
	eventLv3Parent := make(map[string]string)
	for _, e := range eventLv3 {
		key := strings.ToLower(e.KodeEventTypeLv3)
		eventLv3Map[key] = e.ID
		eventLv3Parent[key] = strings.ToLower(e.IDEventTypeLv2)
	}

	incident, _ := riskIssue.incident.GetAll()
	incidentMap := make(map[string]int64)
	for _, i := range incident {
		incidentMap[strings.ToLower(i.KodeKejadian)] = i.ID
	}

	subIncident, _ := riskIssue.subIncident.GetAll()
	subIncidentMap := make(map[string]int64)
	subIncidentParentMap := make(map[string]string)
	for _, i := range subIncident {
		key := strings.ToLower(i.KodeSubKejadian)
		subIncidentMap[key] = i.ID
		subIncidentParentMap[key] = strings.ToLower(i.KodeKejadian)
	}

	subsubIncident, _ := riskIssue.subsubIncident.GetAll()
	subsubIncidentMap := make(map[string]int64)
	subsubIncidentParentMap := make(map[string]string)
	for _, i := range subsubIncident {
		key := strings.ToLower(i.KodePenyebabKejadianLv3)
		subsubIncidentMap[key] = i.ID
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

	// ================================
	// STEP 2 — Ambil mapping yang sudah ada
	// ================================
	eventIDs := []string{}
	for i, row := range data {
		if i == 0 {
			continue
		}
		riskType := strings.ToLower(lib.SafeFirst(lib.ParseStringToArray(row[0], "-")))
		eventIDs = append(eventIDs, riskType)
	}

	mappingResp, err := riskIssue.arlodsService.GetAllMappingRiskEvent(pernr, eventIDs)
	if err != nil {
		return fmt.Errorf("failed get existing mapping: %v", err)
	}
	mappingByEventID := make(map[string]dto.DetailMappingEvent)
	for _, item := range mappingResp.Data {
		mappingByEventID[strings.ToLower(item.RiskEventID)] = item
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
	newMappingRequests := []dto.MappingRiskEventRequest{}

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
		riskEventCode := parse[0]
		if _, ok := riskIssueMap[riskEventCode]; ok {
			continue
		}

		// Buat RiskIssue baru
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

		// Parsing Event LV1-LV3
		evnLv1 := lib.ParseStringToArray(row[4], ";")
		evnLv2 := lib.ParseStringToArray(row[5], ";")
		evnLv3 := lib.ParseStringToArray(row[6], ";")
		mappingEvent := []dto.MappingLVLRequest{}
		for j := range evnLv1 {
			eventLv3 := strconv.Itoa(int(eventLv3Map[lib.ParseStringToArray(evnLv3[j], " - ")[0]]))
			if _, ok := eventLv3Map[eventLv3]; !ok {
				mappingEvent = append(mappingEvent, dto.MappingLVLRequest{
					RiskEventID: riskEventCode,
					Lvl1:        strconv.Itoa(int(eventLv1Map[lib.ParseStringToArray(evnLv1[j], " - ")[0]])),
					Lvl2:        strconv.Itoa(int(eventLv2Map[lib.ParseStringToArray(evnLv2[j], " - ")[0]])),
					Lvl3:        eventLv3,
				})
			}
		}

		// Parsing Cause/Incident LV1-LV3
		incLv1 := lib.ParseStringToArray(row[7], ";")
		incLv2 := lib.ParseStringToArray(row[8], ";")
		incLv3 := lib.ParseStringToArray(row[9], ";")
		mappingCause := []dto.MappingLVLRequest{}
		for j := range incLv1 {
			lv3 := strconv.Itoa(int(eventLv3Map[lib.ParseStringToArray(incLv3[j], " - ")[0]]))
			if _, ok := subsubIncidentMap[lv3]; !ok {
				mappingCause = append(mappingCause, dto.MappingLVLRequest{
					RiskEventID: riskEventCode,
					Lvl1:        strconv.Itoa(int(incidentMap[lib.ParseStringToArray(incLv1[j], " - ")[0]])),
					Lvl2:        strconv.Itoa(int(subIncidentMap[lib.ParseStringToArray(incLv2[j], " - ")[0]])),
					Lvl3:        strings.ToLower(lv3),
				})
			}
		}

		// Parsing Product
		prod := lib.ParseStringToArray(row[10], ";")
		mappingProduct := []dto.MappingProductRiskEventRequest{}
		for _, p := range prod {
			if strings.TrimSpace(p) == "" {
				continue
			}
			if _, ok := productMap[p]; !ok {
				mappingProduct = append(mappingProduct, dto.MappingProductRiskEventRequest{
					RiskEventID: riskEventCode,
					ProductID:   strconv.Itoa(int(productMap[p])),
				})
			}
		}

		businessCycleStr := row[17]
		subBusinessCycleStr := row[19]
		processStr := row[21]
		subProcessStr := row[23]
		activityStr := row[25]

		activities := lib.ParseStringToArray(activityStr, ";")

		// Cache untuk hierarki business process
		cacheBusinessProcess := make(map[string]dto.BusinessProcessNode)

		// Validasi tiap activity dengan panggilan ke service
		for _, act := range activities {
			act = strings.TrimSpace(act)
			if act == "" {
				continue
			}

			resp, err := riskIssue.arlodsService.GetHirearcyBusinessProcess(pernr, act)
			if err != nil {
				continue
			}

			// jika tidak ada node, skip
			if len(resp.Data.List) == 0 {
				continue
			}

			for _, bc := range resp.Data.List {
				extractBusinessProcessNodes(bc, cacheBusinessProcess)
			}
		}

		// Validasi kombinasi multi-value business process
		validation := ValidateBPWithMultiValue(
			businessCycleStr,
			subBusinessCycleStr,
			processStr,
			subProcessStr,
			activityStr,
			cacheBusinessProcess,
		)

		// Jika lolos validasi, baru dibuat mapping Business Process
		mappingBusinessProcess := []dto.MappingRiskEventBusinesProcess{}
		if validation == "" {
			for _, act := range activities {
				act = strings.TrimSpace(act)
				if act == "" {
					continue
				}

				// Ambil node dari cache
				node, ok := cacheBusinessProcess[strings.ToLower(act)]
				if !ok {
					continue // skip jika tidak ada di cache
				}

				mappingBusinessProcess = append(mappingBusinessProcess, dto.MappingRiskEventBusinesProcess{
					RiskEventID:      riskEventCode, // dari hasil bulk insert RiskIssue
					ActivityID:       node.ActivityID,
					BusinessCycle:    node.BusinessCycleID,
					SubBusinessCycle: node.SubBusinessCycleID,
					ProcessID:        node.ProcessID,
					SubProcessID:     node.SubProcessID,
				})
			}
		}

		// Parsing Control & Indicator
		controlCodes := lib.ParseStringToArray(row[11], ";")
		indicatorCodes := lib.ParseStringToArray(row[12], ";")
		mappingControlIndicator := []dto.RiskEventControlMutateInput{}

		for _, c := range controlCodes {
			c = strings.ToLower(strings.TrimSpace(c))
			if c == "" {
				continue
			}
			if id, ok := controlIDMap[c]; ok {
				mappingControlIndicator = append(mappingControlIndicator, dto.RiskEventControlMutateInput{
					RiskEventID: riskEventCode,
					TypeEvent:   "event",
					RiskID:      []string{strconv.Itoa(int(id))},
				})
			}
		}

		for _, ind := range indicatorCodes {
			ind = strings.ToLower(strings.TrimSpace(ind))
			if ind == "" {
				continue
			}
			if id, ok := indicatorIDMap[ind]; ok {
				mappingControlIndicator = append(mappingControlIndicator, dto.RiskEventControlMutateInput{
					RiskEventID: riskEventCode,
					TypeEvent:   "indicator",
					RiskID:      []string{strconv.Itoa(int(id))},
				})
			}
		}

		// Gabungkan semua mapping
		newMappingRequests = append(newMappingRequests, dto.MappingRiskEventRequest{
			MappingRiskEvent:        mappingEvent,
			MappingCauseRiskEvent:   mappingCause,
			MappingProductRiskEvent: mappingProduct,
			MappingBusinessProcess:  mappingBusinessProcess,
			MappingIndicatorControl: mappingControlIndicator,
		})
	}

	tx := riskIssue.db.DB.Begin()
	// ================================
	// STEP 4 — Bulk insert RiskIssue
	// ================================
	if len(newRiskEvents) > 0 {
		if err := riskIssue.riskissueRepo.BulkCreateRiskEvent(newRiskEvents, tx); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed bulk create risk events: %v", err)
		}
	}

	issue, err := riskIssue.riskissueRepo.GetAll()
	if err != nil {
		return fmt.Errorf("failed get risk event: %v", err)
	}

	// Mapping code -> ID
	riskEventMap := make(map[string]string)
	for _, e := range issue {
		riskEventMap[e.RiskIssueCode] = strconv.FormatInt(e.ID, 10)
	}

	// ================================
	// STEP 5 — Update RiskEventID di semua mapping
	// ================================
	for i := range newMappingRequests {
		// Mapping Event
		for j := range newMappingRequests[i].MappingRiskEvent {
			code := newMappingRequests[i].MappingRiskEvent[j].RiskEventID
			if id, ok := riskEventMap[code]; ok {
				newMappingRequests[i].MappingRiskEvent[j].RiskEventID = id
			}
		}
		for j, v := range newMappingRequests[i].MappingRiskEvent {
			code := newMappingRequests[i].MappingRiskEvent[j].RiskEventID
			if id, ok := riskEventMap[code]; ok {
				i64, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
				}
				updateEvent := &models.MapEvent{
					IDRiskIssue:  i64,
					EventTypeLv1: v.Lvl1,
					EventTypeLv2: v.Lvl2,
					EventTypeLv3: v.Lvl3,
				}

				_, err = riskIssue.mapEvent.Update(updateEvent, tx)
				if err != nil {
					tx.Rollback()
					riskIssue.logger.Zap.Error(err)
					return err
				}
			}
		}
		// Mapping Cause
		for j, v := range newMappingRequests[i].MappingCauseRiskEvent {
			code := newMappingRequests[i].MappingCauseRiskEvent[j].RiskEventID
			if id, ok := riskEventMap[code]; ok {
				i64, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
				}
				updateKejadian := &models.MapKejadian{
					IDRiskIssue:         i64,
					PenyebabKejadianLv1: v.Lvl1,
					PenyebabKejadianLv2: v.Lvl2,
					PenyebabKejadianLv3: v.Lvl3,
				}

				_, err = riskIssue.mapKejadian.Update(updateKejadian, tx)
				if err != nil {
					tx.Rollback()
					riskIssue.logger.Zap.Error(err)
					return err
				}
			}
		}
		// Mapping Product
		for j, v := range newMappingRequests[i].MappingProductRiskEvent {
			code := newMappingRequests[i].MappingProductRiskEvent[j].RiskEventID
			if id, ok := riskEventMap[code]; ok {
				i64, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
				}
				pi64, err := strconv.ParseInt(v.ProductID, 10, 64)
				if err != nil {
				}
				_, err = riskIssue.mapProduct.Store(&models.MapProduct{
					IDRiskIssue: i64,
					Product:     pi64,
				}, tx)

				if err != nil {
					tx.Rollback()
					riskIssue.logger.Zap.Error(err)
					return err
				}
				newMappingRequests[i].MappingProductRiskEvent[j].RiskEventID = id
			}
		}

		for j, v := range newMappingRequests[i].MappingIndicatorControl {
			code := newMappingRequests[i].MappingIndicatorControl[j].RiskEventID
			if id, ok := riskEventMap[code]; ok {
				i64, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
				}

				if v.TypeEvent == "control" {
					for _, val := range v.RiskID {
						contrlID, err := strconv.ParseInt(val, 10, 64)
						if err != nil {
						}
						req := &models.MapControl{
							IDRiskIssue: i64,
							IDControl:   contrlID,
							IsChecked:   false,
						}

						_, err = riskIssue.mapControl.Store(req, tx)
						if err != nil {
							tx.Rollback()
							riskIssue.logger.Zap.Error(err)
							return err
						}
					}

				}

				if v.TypeEvent == "indicator" {
					for _, val := range v.RiskID {
						contrlID, err := strconv.ParseInt(val, 10, 64)
						if err != nil {
						}
						req := &models.MapIndicator{
							IDRiskIssue: i64,
							IDIndicator: contrlID,
							IsChecked:   false,
						}

						_, err = riskIssue.mapIndicator.Store(req, tx)
						if err != nil {
							tx.Rollback()
							riskIssue.logger.Zap.Error(err)
							return err
						}
					}

				}

			}
		}
		// Mapping Business Process
		for j := range newMappingRequests[i].MappingBusinessProcess {
			code := newMappingRequests[i].MappingBusinessProcess[j].RiskEventID
			if id, ok := riskEventMap[code]; ok {
				newMappingRequests[i].MappingBusinessProcess[j].RiskEventID = id
			}
		}
	}

	// ================================
	// STEP 6 — Create Mapping Event
	// ================================
	if len(newMappingRequests) > 0 {
		req := dto.BulkMappingRiskEventRequest{
			Data: newMappingRequests,
		}
		if err := riskIssue.arlodsService.CreateMappingEvent(pernr, req); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed create mapping event: %v", err)
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
		"Status",
		"Create Time",
		"Update Time",
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
		25, 25, 25, 25, 25,
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

	control, err := riskIssue.riskControl.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping event: ", err)
		return nil, "", err
	}

	controlMap := make(map[string]modelsControl.RiskControlResponse, 0)
	for _, c := range control {
		controlMap[strconv.Itoa(int(c.ID))] = c
	}

	indicator, err := riskIssue.riskIndicator.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping indicator: ", err)
		return nil, "", err
	}

	indicatorMap := make(map[string]modelsIndicator.RiskIndicatorResponse, 0)
	for _, c := range indicator {
		indicatorMap[strconv.Itoa(int(c.ID))] = c
	}

	eventID := make([]string, 0)
	for _, v := range data {
		eventID = append(eventID, strconv.Itoa(int(v.ID)))
	}

	mapping, err := riskIssue.arlodsService.GetAllMappingRiskEvent(pernr, eventID)
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping: ", err)
		return nil, "", err
	}

	mappingMap := make(map[string]dto.BulkRiskEventDetail, 0)
	for _, v := range mapping.Data {
		mappingMap[v.RiskEventID] = v.MappingDetail
	}

	for _, v := range data {
		status := "Aktif"
		if !v.Status {
			status = "Inactif"
		}
		createTime := lib.FormatDatePtr(v.CreatedAt)
		updateTime := lib.FormatDatePtr(v.UpdatedAt)

		riskType := ""
		if v, ok := riskTypeMap[v.RiskTypeID]; ok {
			riskType = v
		}

		mappingDetail, ok := mappingMap[strconv.Itoa(int(v.ID))]

		var eventTypeLv1, eventTypeLv2, eventTypeLv3 []string
		if ok && len(mappingDetail.MappingEvent) > 0 {
			for _, m := range mappingDetail.MappingEvent {
				eventTypeLv1 = append(eventTypeLv1, m.EventTypeLvl1)
				eventTypeLv2 = append(eventTypeLv2, m.EventTypeLvl2)
				eventTypeLv3 = append(eventTypeLv3, m.EventTypeLvl3)
			}
		} else {
			eventTypeLv1 = []string{""}
			eventTypeLv2 = []string{""}
			eventTypeLv3 = []string{""}
		}

		var causeLv1, causeLv2, causeLv3 []string
		if ok && len(mappingDetail.MappingCause) > 0 {
			for _, m := range mappingDetail.MappingCause {
				causeLv1 = append(causeLv1, m.Incident)
				causeLv2 = append(causeLv2, m.SubIncident)
				causeLv3 = append(causeLv3, m.SubSubIncident)
			}
		} else {
			causeLv1 = []string{""}
			causeLv2 = []string{""}
			causeLv3 = []string{""}
		}

		var productCodes []string
		if ok && len(mappingDetail.MappingProduct) > 0 {
			for _, m := range mappingDetail.MappingProduct {
				productCodes = append(productCodes, m.ProductID)
			}
		} else {
			productCodes = []string{""}
		}

		var codeControls, nameControls []string
		if ok && len(mappingDetail.Controls) > 0 {
			for _, v := range mappingDetail.Controls {
				if val, ok := controlMap[v.ID]; ok {
					codeControls = append(codeControls, val.Kode)
					nameControls = append(nameControls, val.RiskControl)
				}
			}
		}
		if len(codeControls) == 0 {
			codeControls = []string{""}
			nameControls = []string{""}
		}

		var codeIndicators, nameIndicators []string
		if ok && len(mappingDetail.Indicators) > 0 {
			for _, v := range mappingDetail.Indicators {
				if val, ok := indicatorMap[v.ID]; ok {
					codeIndicators = append(codeIndicators, val.RiskIndicatorCode)
					nameIndicators = append(nameIndicators, val.RiskIndicator)
				}
			}
		}
		if len(codeIndicators) == 0 {
			codeIndicators = []string{""}
			nameIndicators = []string{""}
		}

		likelihood := ""
		impact := ""
		if v.Likelihood != nil {
			likelihood = *v.Likelihood
		}
		if v.Impact != nil {
			impact = *v.Impact
		}

		row := []string{
			riskType,
			fmt.Sprintf("%s - %s", v.RiskIssueCode, v.RiskIssue),
			v.Deskripsi,
			v.KategoriRisiko,
			strings.Join(eventTypeLv1, ";"),
			strings.Join(eventTypeLv2, ";"),
			strings.Join(eventTypeLv3, ";"),
			strings.Join(causeLv1, ";"),
			strings.Join(causeLv2, ";"),
			strings.Join(causeLv3, ";"),
			strings.Join(productCodes, ";"),
			likelihood,
			impact,
			status,
			createTime,
			updateTime,
			strings.Join(codeControls, ";"),
			strings.Join(nameControls, ";"),
			strings.Join(codeIndicators, ";"),
			strings.Join(nameIndicators, ";"),
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
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
		"Status",
		"Create Time",
		"Update Time",
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

	control, err := riskIssue.riskControl.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping event: ", err)
		return nil, "", err
	}

	controlMap := make(map[string]modelsControl.RiskControlResponse, 0)
	for _, c := range control {
		controlMap[strconv.Itoa(int(c.ID))] = c
	}

	indicator, err := riskIssue.riskIndicator.GetAll()
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping indicator: ", err)
		return nil, "", err
	}

	indicatorMap := make(map[string]modelsIndicator.RiskIndicatorResponse, 0)
	for _, c := range indicator {
		indicatorMap[strconv.Itoa(int(c.ID))] = c
	}

	eventID := make([]string, 0)
	for _, v := range data {
		eventID = append(eventID, strconv.Itoa(int(v.ID)))
	}

	mapping, err := riskIssue.arlodsService.GetAllMappingRiskEvent(pernr, eventID)
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when get mapping: ", err)
		return nil, "", err
	}

	mappingMap := make(map[string]dto.BulkRiskEventDetail, 0)
	for _, v := range mapping.Data {
		mappingMap[v.RiskEventID] = v.MappingDetail
	}

	for idx, v := range data {
		status := "Aktif"
		if !v.Status {
			status = "Inactif"
		}
		createTime := lib.FormatDatePtr(v.CreatedAt)
		updateTime := lib.FormatDatePtr(v.UpdatedAt)

		riskType := ""
		if v, ok := riskTypeMap[v.RiskTypeID]; ok {
			riskType = v
		}

		mappingDetail, ok := mappingMap[strconv.Itoa(int(v.ID))]

		var eventTypeLv1, eventTypeLv2, eventTypeLv3 []string
		if ok && len(mappingDetail.MappingEvent) > 0 {
			for _, m := range mappingDetail.MappingEvent {
				eventTypeLv1 = append(eventTypeLv1, m.EventTypeLvl1)
				eventTypeLv2 = append(eventTypeLv2, m.EventTypeLvl2)
				eventTypeLv3 = append(eventTypeLv3, m.EventTypeLvl3)
			}
		} else {
			eventTypeLv1 = []string{""}
			eventTypeLv2 = []string{""}
			eventTypeLv3 = []string{""}
		}

		var causeLv1, causeLv2, causeLv3 []string
		if ok && len(mappingDetail.MappingCause) > 0 {
			for _, m := range mappingDetail.MappingCause {
				causeLv1 = append(causeLv1, m.Incident)
				causeLv2 = append(causeLv2, m.SubIncident)
				causeLv3 = append(causeLv3, m.SubSubIncident)
			}
		} else {
			causeLv1 = []string{""}
			causeLv2 = []string{""}
			causeLv3 = []string{""}
		}

		var productCodes []string
		if ok && len(mappingDetail.MappingProduct) > 0 {
			for _, m := range mappingDetail.MappingProduct {
				productCodes = append(productCodes, m.ProductID)
			}
		} else {
			productCodes = []string{""}
		}

		var codeControls, nameControls []string
		if ok && len(mappingDetail.Controls) > 0 {
			for _, v := range mappingDetail.Controls {
				if val, ok := controlMap[v.ID]; ok {
					codeControls = append(codeControls, val.Kode)
					nameControls = append(nameControls, val.RiskControl)
				}
			}
		}
		if len(codeControls) == 0 {
			codeControls = []string{""}
			nameControls = []string{""}
		}

		var codeIndicators, nameIndicators []string
		if ok && len(mappingDetail.Indicators) > 0 {
			for _, v := range mappingDetail.Indicators {
				if val, ok := indicatorMap[v.ID]; ok {
					codeIndicators = append(codeIndicators, val.RiskIndicatorCode)
					nameIndicators = append(nameIndicators, val.RiskIndicator)
				}
			}
		}
		if len(codeIndicators) == 0 {
			codeIndicators = []string{""}
			nameIndicators = []string{""}
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
		f.SetCellValue(sheet, fmt.Sprintf("E%d", idx+2), strings.Join(eventTypeLv1, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", idx+2), strings.Join(eventTypeLv2, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", idx+2), strings.Join(eventTypeLv3, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("H%d", idx+2), strings.Join(causeLv1, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("I%d", idx+2), strings.Join(causeLv2, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("J%d", idx+2), strings.Join(causeLv3, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("K%d", idx+2), strings.Join(productCodes, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("L%d", idx+2), likelihood)
		f.SetCellValue(sheet, fmt.Sprintf("M%d", idx+2), impact)
		f.SetCellValue(sheet, fmt.Sprintf("N%d", idx+2), status)
		f.SetCellValue(sheet, fmt.Sprintf("O%d", idx+2), createTime)
		f.SetCellValue(sheet, fmt.Sprintf("P%d", idx+2), updateTime)
		f.SetCellValue(sheet, fmt.Sprintf("Q%d", idx+2), strings.Join(codeControls, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("R%d", idx+2), strings.Join(nameControls, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("S%d", idx+2), strings.Join(codeIndicators, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("T%d", idx+2), strings.Join(nameIndicators, ";"))
		f.SetCellValue(sheet, fmt.Sprintf("U%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("V%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("W%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("X%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("Y%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("Z%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("AA%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("AB%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("AC%d", idx+2), "")
		f.SetCellValue(sheet, fmt.Sprintf("AD%d", idx+2), "")
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
		"Create Time",
		"Update Time",
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

	riskType, err := riskissue.riskType.GetAll()
	if err != nil {
		riskissue.logger.Zap.Error("Errored when get mapping risk type: ", err)
		return nil, "", err
	}

	riskTypeMap := make(map[int64]string, 0)
	for _, v := range riskType {
		riskTypeMap[v.ID] = fmt.Sprintf("%s - %s", v.RiskTypeCode, v.RiskType)
	}

	control, err := riskissue.riskControl.GetAll()
	if err != nil {
		riskissue.logger.Zap.Error("Errored when get mapping event: ", err)
		return nil, "", err
	}

	controlMap := make(map[string]modelsControl.RiskControlResponse, 0)
	for _, c := range control {
		controlMap[strconv.Itoa(int(c.ID))] = c
	}

	indicator, err := riskissue.riskIndicator.GetAll()
	if err != nil {
		riskissue.logger.Zap.Error("Errored when get mapping indicator: ", err)
		return nil, "", err
	}

	indicatorMap := make(map[string]modelsIndicator.RiskIndicatorResponse, 0)
	for _, c := range indicator {
		indicatorMap[strconv.Itoa(int(c.ID))] = c
	}

	if err := writer.Write(headers); err != nil {
		return nil, "", fmt.Errorf("failed to write csv header: %w", err)
	}

	eventID := make([]string, 0)
	for _, v := range data {
		eventID = append(eventID, strconv.Itoa(int(v.ID)))
	}

	mapping, err := riskissue.arlodsService.GetAllMappingRiskEvent(pernr, eventID)
	if err != nil {
		riskissue.logger.Zap.Error("Errored when get mapping: ", err)
		return nil, "", err
	}

	mappingMap := make(map[string]dto.BulkRiskEventDetail, 0)
	for _, v := range mapping.Data {
		mappingMap[v.RiskEventID] = v.MappingDetail
	}

	for _, d := range data {
		createTime := lib.FormatDatePtr(d.CreatedAt)
		updateTime := lib.FormatDatePtr(d.UpdatedAt)

		riskType := ""
		if v, ok := riskTypeMap[d.RiskTypeID]; ok {
			riskType = v
		}

		mappingDetail, ok := mappingMap[strconv.Itoa(int(d.ID))]

		var eventTypeLv1, eventTypeLv2, eventTypeLv3 []string
		if ok && len(mappingDetail.MappingEvent) > 0 {
			for _, m := range mappingDetail.MappingEvent {
				eventTypeLv1 = append(eventTypeLv1, m.EventTypeLvl1)
				eventTypeLv2 = append(eventTypeLv2, m.EventTypeLvl2)
				eventTypeLv3 = append(eventTypeLv3, m.EventTypeLvl3)
			}
		} else {
			eventTypeLv1 = []string{""}
			eventTypeLv2 = []string{""}
			eventTypeLv3 = []string{""}
		}

		var causeLv1, causeLv2, causeLv3 []string
		if ok && len(mappingDetail.MappingCause) > 0 {
			for _, m := range mappingDetail.MappingCause {
				causeLv1 = append(causeLv1, m.Incident)
				causeLv2 = append(causeLv2, m.SubIncident)
				causeLv3 = append(causeLv3, m.SubSubIncident)
			}
		} else {
			causeLv1 = []string{""}
			causeLv2 = []string{""}
			causeLv3 = []string{""}
		}

		var productCodes []string
		if ok && len(mappingDetail.MappingProduct) > 0 {
			for _, m := range mappingDetail.MappingProduct {
				productCodes = append(productCodes, m.ProductID)
			}
		} else {
			productCodes = []string{""}
		}

		var codeControls, nameControls []string
		if ok && len(mappingDetail.Controls) > 0 {
			for _, v := range mappingDetail.Controls {
				if val, ok := controlMap[v.ID]; ok {
					codeControls = append(codeControls, val.Kode)
					nameControls = append(nameControls, val.RiskControl)
				}
			}
		}
		if len(codeControls) == 0 {
			codeControls = []string{""}
			nameControls = []string{""}
		}

		var codeIndicators, nameIndicators []string
		if ok && len(mappingDetail.Indicators) > 0 {
			for _, v := range mappingDetail.Indicators {
				if val, ok := indicatorMap[v.ID]; ok {
					codeIndicators = append(codeIndicators, val.RiskIndicatorCode)
					nameIndicators = append(nameIndicators, val.RiskIndicator)
				}
			}
		}
		if len(codeIndicators) == 0 {
			codeIndicators = []string{""}
			nameIndicators = []string{""}
		}

		likelihood := ""
		impact := ""
		if d.Likelihood != nil {
			likelihood = *d.Likelihood
		}
		if d.Impact != nil {
			impact = *d.Impact
		}

		row := []string{
			riskType,
			fmt.Sprintf("%s - %s", d.RiskIssueCode, d.RiskIssue),
			d.Deskripsi,
			d.KategoriRisiko,
			strings.Join(eventTypeLv1, ";"),
			strings.Join(eventTypeLv2, ";"),
			strings.Join(eventTypeLv3, ";"),
			strings.Join(causeLv1, ";"),
			strings.Join(causeLv2, ";"),
			strings.Join(causeLv3, ";"),
			strings.Join(productCodes, ";"),
			likelihood,
			impact,
			createTime,
			updateTime,
			strings.Join(codeControls, ";"),
			strings.Join(nameControls, ";"),
			strings.Join(codeIndicators, ";"),
			strings.Join(nameIndicators, ";"),
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
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

func ValidateBPWithMultiValue(
	businessCycleStr string,
	subBusinessCycleStr string,
	processStr string,
	subProcessStr string,
	activityStr string,
	cache map[string]dto.BusinessProcessNode,
) string {

	bcs := strings.Split(businessCycleStr, ";")
	sbcs := strings.Split(subBusinessCycleStr, ";")
	ps := strings.Split(processStr, ";")
	sps := strings.Split(subProcessStr, ";")
	acts := strings.Split(activityStr, ";")

	// --- Cek jumlah harus sama ---
	n := len(acts)
	if len(bcs) != n || len(sbcs) != n || len(ps) != n || len(sps) != n {
		return "Jumlah elemen BP tidak konsisten; "
	}

	var validation string

	// --- Loop berdasarkan index ---
	for i := range n {

		act := strings.TrimSpace(acts[i])
		sp := strings.TrimSpace(sps[i])
		p := strings.TrimSpace(ps[i])
		sbc := strings.TrimSpace(sbcs[i])
		bc := strings.TrimSpace(bcs[i])

		// --- 1. Check Activity ada dalam cache ---
		node, ok := cache[act]
		if !ok {
			validation += fmt.Sprintf("Activity '%s' tidak ditemukan; ", act)
			continue
		}

		// --- 2. VALIDASI PARENT SESUAI URUTAN ---

		// SubProcess
		if node.SubProcessCode != sp {
			validation += fmt.Sprintf("Sub Process '%s' tidak sesuai untuk Activity '%s'; ", sp, act)
		}

		// Process
		if node.ProcessCode != p {
			validation += fmt.Sprintf("Process '%s' tidak sesuai untuk Activity '%s'; ", p, act)
		}

		// Sub Business Cycle
		if node.SubBusinessCycleCode != sbc {
			validation += fmt.Sprintf("Sub Business Cycle '%s' tidak sesuai untuk Activity '%s'; ", sbc, act)
		}

		// Business Cycle
		if node.BusinessCycleCode != bc {
			validation += fmt.Sprintf("Business Cycle '%s' tidak sesuai untuk Activity '%s'; ", bc, act)
		}
	}

	return validation
}

func (riskIssue RiskIssueService) GetRiskCategories(id []int64) ([]string, error) {
	data, err := riskIssue.riskissueRepo.GetRiskCategories(id)
	if err != nil {
		riskIssue.logger.Zap.Error("Errored when try to query risk categories: ", err)
		return nil, err
	}

	return data, nil
}
