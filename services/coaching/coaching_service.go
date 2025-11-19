package coaching

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/coaching"

	"gorm.io/gorm"

	coachingRepo "riskmanagement/repository/coaching"

	fileModel "riskmanagement/models/filemanager"
	filemanager "riskmanagement/services/filemanager"

	msUker "riskmanagement/repository/msuker"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
)

var (
	// timeNow = lib.GetTimeNow("timestime")
	UUID = uuid.NewString()
)

type CoachingDefinition interface {
	WithTrx(trxHandle *gorm.DB) CoachingService
	GetAll() (responses []models.CoachingResponse, err error)
	GetOne(id int64) (responses models.CoachingResponsesGetOneString, status bool, err error)
	Store(request models.CoachingRequest) (status bool, message string, err error)
	StoreDraft(request models.CoachingRequest) (status bool, err error)
	Delete(request *models.CoachingRequestUpdate) (responses bool, err error)
	DeleteCoachingActivity(request *models.CoachingActRequest) (status bool, err error)
	UpdateAllCoaching(request *models.CoachingResponseMaintain) (status bool, message string, err error)
	UpdateDraft(request *models.CoachingResponseMaintain) (status bool, err error)
	FilterCoaching(request models.CoachingFilterRequest) (response []models.CoachingResponseData, pagination lib.Pagination, err error)
	GetNoPelaporan(request models.NoPalaporanRequest) (response []models.NoPelaporanResponse, err error)
	GetData() (responses []models.CoachingResponse, err error)
	GetDataWithPagination(request models.CoachingPagination) (responses []models.CoachingResponseData, pagination lib.Pagination, err error)
	DeleteMapPeserta(request *models.CoachingMapPeserta) (status bool, err error)
	CoachingReportFilter(request *models.CoachingFilterReportRequest) (responses models.CoachingReportResponse, totalRows int64, err error)
	CoachingFinalReportFilter(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterReportFinalResponse, totalRows int64, err error)
	CoachingReportDetail(request models.CoachingReportDetailRequest) (responses models.CoachingReportDetailResponse, err error)
	CoachingReportMateriList(request models.CoachingReportMateriRequest) (responses []models.CoachingDetailMateriResponse, err error)
	CoachingReportFilterByUkerAllActivity(request models.CoachingFilterReportRequest) (responses []models.CoachingFilterByUkerAllActivityReportResponse, totalRows int, totalData int, err error)
	CoachingReportByUkerFilter(request models.CoachingFilterReportByUker) (responses []models.CoachingReportFilteredByUkerResponse, totalRows int64, err error)
	CoachingReportFilterByUkerComplete(request models.CoachingFilterReportByUker) (responses []models.CoachingFilterReportFinalResponse, totalRows int, err error)
	CoachingReportList(request models.CoachingReportListRequest) (responses []models.CoachingReportListFinalResponse, totalRows int, err error)
	CoachingFrekuensiRpt(request models.FrekuensiCoachingRequest) (responses []models.FrekuensiCoachingResponse, totalRows int, err error)
}

type CoachingService struct {
	db               lib.Database
	logger           logger.Logger
	coachingRepo     coachingRepo.CoachingDefinition
	coachingActivity coachingRepo.CoachingActivityDefinition
	MapPeserta       coachingRepo.CoachingMapPesertaDefinition
	filemanager      filemanager.FileManagerDefinition
	msUker           msUker.MsUkerDefinition
}

func NewCoachingService(
	db lib.Database,
	logger logger.Logger,
	coachingRepo coachingRepo.CoachingDefinition,
	coachingActivity coachingRepo.CoachingActivityDefinition,
	MapPeserta coachingRepo.CoachingMapPesertaDefinition,
	filemanager filemanager.FileManagerDefinition,
	msUker msUker.MsUkerDefinition,
) CoachingDefinition {
	return CoachingService{
		db:               db,
		logger:           logger,
		coachingRepo:     coachingRepo,
		coachingActivity: coachingActivity,
		MapPeserta:       MapPeserta,
		filemanager:      filemanager,
		msUker:           msUker,
	}
}

// Delete implements CoachingDefinition
func (coaching CoachingService) Delete(request *models.CoachingRequestUpdate) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := coaching.db.DB.Begin()

	getOneCoaching, exist, err := coaching.GetOne(request.ID)
	if err != nil {
		coaching.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	updateDataCoaching := &models.CoachingUpdateDelete{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02b", //selesai
		Action:        "updateDelete",
		Deleted:       true,
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"status",
		"action",
		"updated_at",
	}
	_, err = coaching.coachingRepo.Delete(updateDataCoaching, include, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneCoaching", getOneCoaching)
		tx.Commit()
		return true, err
	}

	return false, err
}

// DeleteMapPeserta implements CoachingDefinition
func (coaching CoachingService) DeleteMapPeserta(request *models.CoachingMapPeserta) (status bool, err error) {
	tx := coaching.db.DB.Begin()

	err = coaching.MapPeserta.DeleteByID(request.ID, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteCoachingActivity implements CoachingDefinition
func (coaching CoachingService) DeleteCoachingActivity(request *models.CoachingActRequest) (status bool, err error) {
	tx := coaching.db.DB.Begin()
	err = coaching.coachingRepo.DeleteCoachingActivity(request.ID, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetAll implements CoachingDefinition
func (coaching CoachingService) GetAll() (responses []models.CoachingResponse, err error) {
	return coaching.coachingRepo.GetAll()
}

// GetOne implements CoachingDefinition
func (coaching CoachingService) GetOne(id int64) (responses models.CoachingResponsesGetOneString, status bool, err error) {
	dataCoaching, err := coaching.coachingRepo.GetOne(id)
	fmt.Println(dataCoaching)

	if dataCoaching.ID != 0 {
		fmt.Println("Bukan 0")

		activity, err := coaching.coachingActivity.GetOneActivity(dataCoaching.ID)
		peserta, err := coaching.MapPeserta.GetByIDCoaching(dataCoaching.ID)

		responses = models.CoachingResponsesGetOneString{
			ID:             dataCoaching.ID,
			NoPelaporan:    dataCoaching.NoPelaporan,
			REGION:         dataCoaching.REGION,
			RGDESC:         dataCoaching.RGDESC,
			MAINBR:         dataCoaching.MAINBR,
			MBDESC:         dataCoaching.MBDESC,
			BRANCH:         dataCoaching.BRANCH,
			BRDESC:         dataCoaching.BRDESC,
			JenisPeserta:   dataCoaching.JenisPeserta,
			JabatanPeserta: dataCoaching.JabatanPeserta,
			Peserta:        peserta,
			JumlahPeserta:  dataCoaching.JumlahPeserta,
			ListPeserta:    dataCoaching.ListPeserta,
			ActivityID:     dataCoaching.ActivityID,
			SubActivityID:  dataCoaching.SubActivityID,
			ProductID:      dataCoaching.ProductID,
			MakerID:        dataCoaching.MakerID,
			MakerDesc:      dataCoaching.MakerDesc,
			MakerDate:      dataCoaching.MakerDate,
			LastMakerID:    dataCoaching.LastMakerID,
			LastMakerDesc:  dataCoaching.LastMakerDesc,
			LastMakerDate:  dataCoaching.LastMakerDate,
			Status:         dataCoaching.Status,
			Action:         dataCoaching.Action,
			Deleted:        dataCoaching.Deleted,
			Activity:       activity,
			UpdatedAt:      dataCoaching.UpdatedAt,
			CreatedAt:      dataCoaching.CreatedAt,
		}

		return responses, true, err
	}
	return responses, false, err
}

// Store implements CoachingDefinition
func (coaching CoachingService) StoreDraft(request models.CoachingRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := coaching.db.DB.Begin()

	fmt.Println("Data => ", request)

	reqCoaching := &models.Coaching{
		NoPelaporan:    request.NoPelaporan,
		REGION:         request.REGION,
		RGDESC:         request.RGDESC,
		MAINBR:         request.MAINBR,
		MBDESC:         request.MBDESC,
		BRANCH:         request.BRANCH,
		BRDESC:         request.BRDESC,
		JenisPeserta:   request.JenisPeserta,
		JabatanPeserta: request.JabatanPeserta,
		JumlahPeserta:  request.JumlahPeserta,
		ListPeserta:    request.ListPeserta,
		ActivityID:     request.ActivityID,
		SubActivityID:  request.SubActivityID,
		ProductID:      request.ProductID,
		MakerID:        request.MakerID,
		MakerDesc:      request.MakerDesc,
		MakerDate:      &timeNow,
		LastMakerID:    request.LastMakerID,
		LastMakerDesc:  request.LastMakerDesc,
		LastMakerDate:  request.LastMakerDate,
		Status:         "01a",
		Action:         "Draft",
		CreatedAt:      &timeNow,
	}

	dataCoaching, err := coaching.coachingRepo.Store(reqCoaching, tx)
	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	fmt.Println("dataCoaching : ", dataCoaching)

	// if len(request.Peserta) != 0 {
	// 	for _, value := range request.Peserta {
	// 		_, err = coaching.MapPeserta.Store(&models.CoachingMapPeserta{
	// 			IDCoaching:  dataCoaching.ID,
	// 			PERNR:       value.PERNR,
	// 			NamaPeserta: value.NamaPeserta,
	// 			SteelTx:     value.SteelTx,
	// 		}, tx)

	// 		if err != nil {
	// 			tx.Rollback()
	// 			coaching.logger.Zap.Error(err)
	// 			return false, err
	// 		}
	// 	}
	// } else {
	// 	tx.Rollback()
	// 	coaching.logger.Zap.Error(err)
	// 	return false, err
	// }

	if len(request.Activity) != 0 {
		for _, value := range request.Activity {
			_, err = coaching.coachingActivity.Store(&models.CoachingActivity{
				CoachingID:        dataCoaching.ID,
				RiskIssueID:       value.RiskIssueID,
				RiskIssue:         value.RiskIssue,
				RiskIssueCode:     value.RiskIssueCode,
				JudulMateri:       value.JudulMateri,
				RiskIndicatorID:   value.RiskIndicatorID,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				CreatedAt:         &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				coaching.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// Store implements CoachingDefinition
func (coaching CoachingService) Store(request models.CoachingRequest) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := coaching.db.DB.Begin()

	if request.NoPelaporan != "" {
		reqCoaching := &models.Coaching{
			NoPelaporan:    request.NoPelaporan,
			REGION:         request.REGION,
			RGDESC:         request.RGDESC,
			MAINBR:         request.MAINBR,
			MBDESC:         request.MBDESC,
			BRANCH:         request.BRANCH,
			BRDESC:         request.BRDESC,
			JenisPeserta:   request.JenisPeserta,
			JabatanPeserta: request.JabatanPeserta,
			JumlahPeserta:  request.JumlahPeserta,
			ListPeserta:    request.ListPeserta,
			ActivityID:     request.ActivityID,
			SubActivityID:  request.SubActivityID,
			ProductID:      request.ProductID,
			MakerID:        request.MakerID,
			MakerDesc:      request.MakerDesc,
			MakerDate:      &timeNow,
			LastMakerID:    request.LastMakerID,
			LastMakerDesc:  request.LastMakerDesc,
			LastMakerDate:  request.LastMakerDate,
			// Status:        "01a",
			// Action:        "Selesai",
			// change 23/11/2023
			Status:    request.Status,
			Action:    request.Action,
			CreatedAt: &timeNow,
		}

		dataCoaching, err := coaching.coachingRepo.Store(reqCoaching, tx)
		if err != nil {
			tx.Rollback()
			coaching.logger.Zap.Error(err)
			message = "Error transaction database !"
			return false, message, err
		}

		fmt.Println("dataCoaching : ", dataCoaching)

		// if len(request.Peserta) != 0 {
		// 	for _, value := range request.Peserta {
		// 		_, err = coaching.MapPeserta.Store(&models.CoachingMapPeserta{
		// 			IDCoaching:  dataCoaching.ID,
		// 			PERNR:       value.PERNR,
		// 			NamaPeserta: value.NamaPeserta,
		// 			SteelTx:     value.SteelTx,
		// 		}, tx)

		// 		if err != nil {
		// 			tx.Rollback()
		// 			coaching.logger.Zap.Error(err)
		// 			message = "Error transaction database !"
		// 			return false, message, err
		// 		}
		// 	}
		// } else {
		// 	tx.Rollback()
		// 	coaching.logger.Zap.Error(err)
		// 	message = "Data peserta kosong"
		// 	return false, message, err
		// }

		if len(request.Activity) != 0 {
			for _, value := range request.Activity {
				_, err = coaching.coachingActivity.Store(&models.CoachingActivity{
					CoachingID:        dataCoaching.ID,
					RiskIssueID:       value.RiskIssueID,
					RiskIssue:         value.RiskIssue,
					RiskIssueCode:     value.RiskIssueCode,
					TitleMateries:     value.TitleMateries,
					JudulMateri:       value.JudulMateri,
					RiskIndicatorID:   value.RiskIndicatorID,
					RekomendasiMateri: value.RekomendasiMateri,
					MateriTambahan:    value.MateriTambahan,
					CreatedAt:         &timeNow,
				}, tx)

				if err != nil {
					tx.Rollback()
					coaching.logger.Zap.Error(err)
					message = "Error transaction database !"
					return false, message, err
				}
			}
		} else {
			tx.Rollback()
			coaching.logger.Zap.Error(err)
			message = "Data materi kosong !"
			return false, message, err
		}

		checkPekerja, err := coaching.msUker.CheckJumlahPekerja(request.BRANCH)
		fmt.Println("Jumlah Pekerja Uker => ", checkPekerja)

		if request.JenisPeserta == "peruker" && request.JumlahPeserta != checkPekerja && request.JumlahPeserta != int64(len(request.Peserta)) {
			tx.Rollback()
			fmt.Println("Jumlah Pekerja tidak sama dengan jumlah pekerja uker")
			message = "Jumlah Pekerja tidak sama dengan jumlah pekerja uker !"
			return false, message, err
		}

		tx.Commit()
		if request.Action == "Draft" {
			message = "Berhasil menyimpan draft"
		} else {
			message = "Input data berhasil"
		}
		return true, message, err
	} else {
		tx.Rollback()
		message = "Data Gagal disimpan, nomor pelaporan kosong!"
		return false, message, err
	}

}

// UpdateAllCoaching implements CoachingDefinition
func (coaching CoachingService) UpdateAllCoaching(request *models.CoachingResponseMaintain) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := coaching.db.DB.Begin()

	updateCoaching := &models.CoachingUpdateActivity{
		ID:             request.ID,
		REGION:         request.REGION,
		RGDESC:         request.RGDESC,
		MAINBR:         request.MAINBR,
		MBDESC:         request.MBDESC,
		BRANCH:         request.BRANCH,
		BRDESC:         request.BRDESC,
		UnitKerja:      request.UnitKerja,
		JenisPeserta:   request.JenisPeserta,
		JabatanPeserta: request.JabatanPeserta,
		JumlahPeserta:  request.JumlahPeserta,
		ListPeserta:    request.ListPeserta,
		ActivityID:     request.ActivityID,
		SubActivityID:  request.SubActivityID,
		ProductID:      request.ProductID,
		LastMakerID:    request.LastMakerID,
		LastMakerDesc:  request.LastMakerDesc,
		LastMakerDate:  &timeNow,
		Deleted:        false,
		// Action:        "Update",
		// Status:        "02b",
		Action:    request.Action,
		Status:    request.Status,
		UpdatedAt: &timeNow,
	}

	include := []string{
		"REGION",
		"RGDESC",
		"MAINBR",
		"MBDESC",
		"BRANCH",
		"BRDESC",
		"unit_kerja",
		"peserta",
		"jumlah_peserta",
		"activity_id",
		"sub_activity_id",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"action",
		"status",
		"updated_at",
	}

	_, err = coaching.coachingRepo.UpdateAllCoaching(updateCoaching, include, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		message = "Error transaction database !"
		return false, message, err
	}

	// if len(request.Peserta) != 0 {
	// 	for _, value := range request.Peserta {
	// 		_, err = coaching.MapPeserta.Store(&models.CoachingMapPeserta{
	// 			ID:          value.ID,
	// 			IDCoaching:  request.ID,
	// 			PERNR:       value.PERNR,
	// 			NamaPeserta: value.NamaPeserta,
	// 			SteelTx:     value.SteelTx,
	// 		}, tx)

	// 		if err != nil {
	// 			tx.Rollback()
	// 			coaching.logger.Zap.Error(err)
	// 			message = "Error transaction database !"
	// 			return false, message, err
	// 		}
	// 	}
	// } else {
	// 	tx.Rollback()
	// 	coaching.logger.Zap.Error(err)
	// 	message = "Error transaction database !"
	// 	return false, message, err
	// }

	if len(request.Activity) != 0 {
		for _, value := range request.Activity {
			updateCoachinAct := &models.CoachingActivity{
				ID:                value.ID,
				CoachingID:        request.ID,
				RiskIssueID:       value.RiskIssueID,
				RiskIssue:         value.RiskIssue,
				RiskIssueCode:     value.RiskIssueCode,
				JudulMateri:       value.JudulMateri,
				RiskIndicatorID:   value.RiskIndicatorID,
				RekomendasiMateri: value.RekomendasiMateri,
				TitleMateries:     value.TitleMateries,
				MateriTambahan:    value.MateriTambahan,
				UpdatedAt:         &timeNow,
				CreatedAt:         value.CreatedAt,
			}

			fmt.Println("rows ====> ", updateCoachinAct)
			_, err = coaching.coachingActivity.Store(updateCoachinAct, tx)

			if err != nil {
				tx.Rollback()
				coaching.logger.Zap.Error(err)
				message = "Error transaction database !"
				return false, message, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			coaching.logger.Zap.Error(err)
			message = "Error transaction database !"
			return false, message, err
		}
	}

	checkPekerja, err := coaching.msUker.CheckJumlahPekerja(request.BRANCH)
	fmt.Println("Jumlah Pekerja Uker => ", checkPekerja)

	if request.JenisPeserta == "peruker" && request.JumlahPeserta != checkPekerja && request.JumlahPeserta != int64(len(request.Peserta)) {
		tx.Rollback()
		fmt.Println("Jumlah Pekerja tidak sama dengan jumlah pekerja uker")
		message = "Jumlah Pekerja tidak sama dengan jumlah pekerja uker !"
		return false, message, err
	}

	tx.Commit()
	message = "Update Data Berhasil"
	return true, message, err
}

func (coaching CoachingService) UpdateDraft(request *models.CoachingResponseMaintain) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := coaching.db.DB.Begin()

	updateCoaching := &models.CoachingUpdateActivity{
		ID:            request.ID,
		REGION:        request.REGION,
		RGDESC:        request.RGDESC,
		MAINBR:        request.MAINBR,
		MBDESC:        request.MBDESC,
		BRANCH:        request.BRANCH,
		BRDESC:        request.BRDESC,
		UnitKerja:     request.UnitKerja,
		JenisPeserta:  request.JenisPeserta,
		JumlahPeserta: request.JumlahPeserta,
		ActivityID:    request.ActivityID,
		SubActivityID: request.SubActivityID,
		ProductID:     request.ProductID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Deleted:       false,
		Action:        "Draft",
		Status:        "01a",
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"REGION",
		"RGDESC",
		"MAINBR",
		"MBDESC",
		"BRANCH",
		"BRDESC",
		"unit_kerja",
		"peserta",
		"jumlah_peserta",
		"activity_id",
		"sub_activity_id",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"action",
		"status",
		"updated_at",
	}

	_, err = coaching.coachingRepo.UpdateAllCoaching(updateCoaching, include, tx)

	if err != nil {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	if len(request.Peserta) != 0 {
		for _, value := range request.Peserta {
			_, err = coaching.MapPeserta.Store(&models.CoachingMapPeserta{
				ID:          value.ID,
				IDCoaching:  request.ID,
				PERNR:       value.PERNR,
				NamaPeserta: value.NamaPeserta,
				SteelTx:     value.SteelTx,
			}, tx)

			if err != nil {
				tx.Rollback()
				coaching.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		coaching.logger.Zap.Error(err)
		return false, err
	}

	if len(request.Activity) != 0 {
		for _, value := range request.Activity {
			updateCoachinAct := &models.CoachingActivity{
				ID:                value.ID,
				CoachingID:        request.ID,
				RiskIssueID:       value.RiskIssueID,
				RiskIssue:         value.RiskIssue,
				RiskIssueCode:     value.RiskIssueCode,
				JudulMateri:       value.JudulMateri,
				RiskIndicatorID:   value.RiskIndicatorID,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				UpdatedAt:         &timeNow,
				CreatedAt:         value.CreatedAt,
			}

			fmt.Println("rows ====> ", updateCoachinAct)
			_, err = coaching.coachingActivity.Store(updateCoachinAct, tx)

			if err != nil {
				tx.Rollback()
				coaching.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			coaching.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

// WithTrx implements CoachingDefinition
func (coaching CoachingService) WithTrx(trxHandle *gorm.DB) CoachingService {
	coaching.coachingRepo = coaching.coachingRepo.WithTrx(trxHandle)
	return coaching
}

// FilterCoaching implements CoachingDefinition
func (coaching CoachingService) FilterCoaching(request models.CoachingFilterRequest) (responses []models.CoachingResponseData, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataCoach, totalRows, totalData, err := coaching.coachingRepo.FilterCoaching(&request)
	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataCoach {
		materi, err := coaching.coachingRepo.GetJudulMateri(response.ID)
		if err != nil {
			coaching.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.CoachingResponseData{
			ID:          response.ID,
			NoPelaporan: response.NoPelaporan,
			UnitKerja:   response.UnitKerja,
			Aktifitas:   response.Aktifitas,
			Materi:      materi,
			StatusCoach: response.Status,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	fmt.Println("pagination =>", pagination)

	return responses, pagination, err
}

// GetDataWithPagination implements CoachingDefinition
func (coaching CoachingService) GetDataWithPagination(request models.CoachingPagination) (responses []models.CoachingResponseData, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataCoach, totalRows, totalData, err := coaching.coachingRepo.GetDataWithPagination(&request)
	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		coaching.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataCoach {
		materi, err := coaching.coachingRepo.GetJudulMateri(response.ID)
		if err != nil {
			coaching.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.CoachingResponseData{
			ID:          response.ID,
			NoPelaporan: response.NoPelaporan,
			Aktifitas:   response.Aktifitas,
			Materi:      materi,
			UnitKerja:   response.UnitKerja,
			StatusCoach: response.Status,
		})
	}

	fmt.Println()

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// GetNoPelaporan implements CoachingDefinition
func (coaching CoachingService) GetNoPelaporan(request models.NoPalaporanRequest) (responses []models.NoPelaporanResponse, err error) {
	dataCoaching, err := coaching.coachingRepo.GetNoPelaporan(&request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataCoaching {
		responses = append(responses, models.NoPelaporanResponse{
			ORGEH:       request.ORGEH,
			NoPelaporan: response.NoPelaporan.String,
		})
	}

	return responses, err
}

// GetData implements CoachingDefinition
func (coaching CoachingService) GetData() (responses []models.CoachingResponse, err error) {
	return coaching.coachingRepo.GetData()
}

func (coaching CoachingService) CoachingReportFilter(request *models.CoachingFilterReportRequest) (responses models.CoachingReportResponse, totalRows int64, err error) {
	var dataArr []models.CoachingFilterReportResponse
	// totalData := int64(0)
	dataCoaching, totalAktivitas, totalRows, err := coaching.coachingRepo.CoachingReportFilter(request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataCoaching {
		dataArr = append(dataArr, models.CoachingFilterReportResponse{
			Id:    response.Id,
			Code:  response.Code,
			Name:  response.Name,
			Total: response.Total,
		})

		// totalData = totalAktivitas
	}

	responses = models.CoachingReportResponse{
		Data:      dataArr,
		TotalData: totalAktivitas,
	}

	fmt.Println("================== Total Data")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (coaching CoachingService) CoachingFinalReportFilter(request *models.CoachingFilterReportRequest) (responses []models.CoachingFilterReportFinalResponse, totalRows int64, err error) {
	dataCoaching, totalRows, err := coaching.coachingRepo.CoachingFinalReportFilter(request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataCoaching {
		responses = append(responses, models.CoachingFilterReportFinalResponse{
			Id:        response.Id,
			Date:      response.Date,
			BRANCH:    response.BRANCH,
			BRDESC:    response.BRDESC,
			Activity:  response.Activity,
			Product:   response.Product,
			RiskIssue: response.RiskIssue,
			Materi:    response.Materi,
		})
	}

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (coaching CoachingService) CoachingReportByUkerFilter(request models.CoachingFilterReportByUker) (responses []models.CoachingReportFilteredByUkerResponse, totalRows int64, err error) {
	dataCoaching, totalRows, err := coaching.coachingRepo.CoachingReportByUkerFilter(&request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataCoaching {
		responses = append(responses, models.CoachingReportFilteredByUkerResponse{
			REGION:          response.REGION,
			RGDESC:          response.RGDESC,
			MAINBR:          response.MAINBR,
			MBDESC:          response.MBDESC,
			BRANCH:          response.BRANCH,
			BRDESC:          response.BRDESC,
			TOTALCOACHING:   response.TOTALCOACHING,
			TOTALBRC:        response.TOTALBRC,
			PERCENTCOACHING: response.PERCENTCOACHING,
		})
	}

	fmt.Println("================== Responses")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (coaching CoachingService) CoachingReportFilterByUkerComplete(request models.CoachingFilterReportByUker) (responses []models.CoachingFilterReportFinalResponse, totalRows int, err error) {
	dataCoaching, totalRows, err := coaching.coachingRepo.CoachingReportFilterByUkerComplete(&request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataCoaching {
		responses = append(responses, models.CoachingFilterReportFinalResponse{
			Id:        response.Id.Int64,
			Date:      response.Date.String,
			BRANCH:    response.BRANCH.String,
			BRDESC:    response.BRDESC.String,
			Activity:  response.Activity.String,
			Product:   response.Product.String,
			RiskIssue: response.RiskIssue.String,
			Materi:    response.Materi.String,
		})
	}

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (coaching CoachingService) CoachingReportFilterByUkerAllActivity(request models.CoachingFilterReportRequest) (responses []models.CoachingFilterByUkerAllActivityReportResponse, totalRows int, totalData int, err error) {
	dataCoaching, _, _, err := coaching.coachingRepo.CoachingReportFilterByUkerAllActivity(&request)

	fmt.Println(dataCoaching)
	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, totalRows, totalData, err
	}

	return responses, totalRows, totalData, err
}

func (coaching CoachingService) CoachingReportDetail(request models.CoachingReportDetailRequest) (responses models.CoachingReportDetailResponse, err error) {
	dataCoaching, err := coaching.coachingRepo.CoachingReportDetail(&request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, err
	}

	responses = dataCoaching

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, err
}

func (coaching CoachingService) CoachingReportMateriList(request models.CoachingReportMateriRequest) (responses []models.CoachingDetailMateriResponse, err error) {
	dataCoaching, err := coaching.coachingRepo.CoachingReportMateriList(&request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataCoaching {
		minioLink, err := coaching.filemanager.GetFile(fileModel.FileManagerRequest{
			Subdir:   response.Path.String,
			Filename: response.Filename.String,
		})
		if err != nil {
			coaching.logger.Zap.Error(err)
			return responses, err
		}

		responses = append(responses, models.CoachingDetailMateriResponse{
			ID:           response.ID.Int64,
			NamaLampiran: response.NamaLampiran.String,
			Filename:     response.Filename,
			Path:         minioLink.MinioPath,
		})
	}

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, err
}

// CoachingReportList implements CoachingDefinition
func (coaching CoachingService) CoachingReportList(request models.CoachingReportListRequest) (responses []models.CoachingReportListFinalResponse, totalRows int, err error) {
	fmt.Println("masuk-service 4", &request)

	dataCoaching, totalRows, err := coaching.coachingRepo.CoachingReportList(&request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataCoaching {
		// materi, _ := coaching.coachingActivity.GetAktifitasReport(response.ID)
		// peserta, _ := coaching.MapPeserta.GetPesertaReport(response.ID)

		responses = append(responses, models.CoachingReportListFinalResponse{
			NoPelaporan:    response.NoPelaporan,
			RGDESC:         response.RGDESC,
			MBDESC:         response.MBDESC,
			BRANCH:         response.BRANCH,
			BRDESC:         response.BRDESC,
			JudulMateri:    response.JudulMateri,
			RincianMateri:  response.RincianMateri,
			JumlahPeserta:  response.JumlahPeserta,
			JabatanPeserta: response.JabatanPeserta,
			JenisPeserta:   response.JenisPeserta,
			Peserta:        response.Peserta,
			Aktifitas:      response.Aktifitas,
			SubAktifitas:   response.SubAktifitas,
			IsuRisiko:      response.IsuRisiko,
			RiskIndicator:  response.RiskIndicator,
			MakerID:        response.MakerID,
			Status:         response.Status,
		})
	}

	fmt.Println("response service")
	// fmt.Println(responses)

	return responses, totalRows, err
}

func (coaching CoachingService) CoachingFrekuensiRpt(request models.FrekuensiCoachingRequest) (responses []models.FrekuensiCoachingResponse, totalRows int, err error) {
	frekuensiRpt, totalRows, err := coaching.coachingRepo.CoachingFrekuensiRpt(&request)

	if err != nil {
		coaching.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, value := range frekuensiRpt {
		responses = append(responses, models.FrekuensiCoachingResponse{
			Aktivitas:     value.Aktivitas,
			Produk:        value.Produk,
			RiskEvent:     value.RiskEvent,
			RiskIndicator: value.RiskIndicator,
			Jumlah:        value.Jumlah,
		})
	}

	return responses, totalRows, err
}
