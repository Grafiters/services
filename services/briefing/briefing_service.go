package briefing

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/briefing"

	briefingRepo "riskmanagement/repository/briefing"

	fileModel "riskmanagement/models/filemanager"
	filemanager "riskmanagement/services/filemanager"

	msUker "riskmanagement/repository/msuker"

	"github.com/google/uuid"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

var (
	UUID = uuid.NewString()
)

type BriefingDefinition interface {
	WithTrx(trxHandle *gorm.DB) BriefingService
	GetAll() (responses []models.BriefingResponse, err error)
	GetData() (responses []models.BriefingResponse, err error)
	GetDataWithPagination(request models.BriefingPagination) (responses []models.BriefingResponseData, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.BriefingResponseGetOneString, status bool, err error)
	Store(request models.BriefingRequest) (status bool, message string, err error)
	StoreDraft(request models.BriefingRequest) (status bool, err error)
	Delete(request *models.BriefingRequestUpdate) (responses bool, err error)
	DeleteBriefingMateri(request *models.BriefMateriRequest) (status bool, err error)
	UpdateAllBrief(request *models.BriefingResponseMaintain) (status bool, message string, err error)
	UpdateDraft(request *models.BriefingResponseMaintain) (status bool, err error)
	FilterBriefing(requests models.BriefingFilterRequest) (responses []models.BriefingResponseData, pagination lib.Pagination, err error)
	GetNoPelaporan(request models.NoPelaporanRequest) (responses []models.NoPelaporanResponse, err error)
	BriefingReportFilter(request models.BriefingFilterReport) (responses models.BriefingReportResponse, totalRows int64, err error)
	BriefingReportFilterComplete(request models.BriefingFilterReport) (responses []models.BriefingFilterReportFinalResponse, totalRows int, err error)
	BriefingReportDetail(request models.BriefingReportDetailRequest) (responses models.BriefingReportDetailResponse, err error)
	BriefingReportMateriList(request models.BriefingReportMateriRequest) (responses []models.BriefingDetailMateriResponse, err error)
	DeleteMapPeserta(request *models.BriefingMapPeserta) (status bool, err error)
	BriefingReportByUkerFilter(request models.BriefingFilterReportByUker) (responses []models.BriefingReportFilteredByUkerResponse, totalRows int64, err error)
	BriefingReportFilterByUkerComplete(request models.BriefingFilterReportByUker) (responses []models.BriefingFilterReportFinalResponse, totalRows int, err error)
	BriefingReportList(request models.BriefingReportListRequest) (responses []models.BriefingReportListFinalResponse, totalRows int, err error)
	BriefingFrekuensiRpt(request models.FrekuensiBriefingRequest) (responses []models.FrekuensiBriefingResponse, totalRows int, err error)
}

type BriefingService struct {
	db             lib.Database
	logger         logger.Logger
	briefingRepo   briefingRepo.BriefingDefinition
	briefingMateri briefingRepo.BriefingMateriDefinition
	MapPeserta     briefingRepo.BriefingMapPesertaDefinition
	filemanager    filemanager.FileManagerDefinition
	msUker         msUker.MsUkerDefinition
}

func NewBriefingService(
	db lib.Database,
	logger logger.Logger,
	briefingRepo briefingRepo.BriefingDefinition,
	briefingMateri briefingRepo.BriefingMateriDefinition,
	MapPeserta briefingRepo.BriefingMapPesertaDefinition,
	filemanager filemanager.FileManagerDefinition,
	msUker msUker.MsUkerDefinition,
) BriefingDefinition {
	return BriefingService{
		db:             db,
		logger:         logger,
		briefingRepo:   briefingRepo,
		briefingMateri: briefingMateri,
		MapPeserta:     MapPeserta,
		filemanager:    filemanager,
		msUker:         msUker,
	}
}

// Delete implements BriefingDefinition
func (briefing BriefingService) Delete(request *models.BriefingRequestUpdate) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := briefing.db.DB.Begin()

	getOneBriefing, exist, err := briefing.GetOne(request.ID)
	if err != nil {
		briefing.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	updateDataBriefing := &models.BriefingUpdateDelete{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Deleted:       true,
		Action:        "updateDelete",
		Status:        "02b", //selesai
		UpdatedAt:     &timeNow,
	}

	_, err = briefing.briefingRepo.Delete(updateDataBriefing,
		[]string{
			"last_maker_id",
			"last_maker_desc",
			"last_maker_date",
			"deleted",
			"action",
			"status",
			"updated_at",
		}, tx)

	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOneBriefing", getOneBriefing)
		tx.Commit()
		return true, err
	}
	return false, err
}

// DeleteMapPeserta implements BriefingDefinition
func (briefing BriefingService) DeleteMapPeserta(request *models.BriefingMapPeserta) (status bool, err error) {
	tx := briefing.db.DB.Begin()

	err = briefing.MapPeserta.DeleteByID(request.ID, tx)

	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// DeleteBriefingMateri implements BriefingDefinition
func (briefing BriefingService) DeleteBriefingMateri(request *models.BriefMateriRequest) (status bool, err error) {
	tx := briefing.db.DB.Begin()
	err = briefing.briefingRepo.DeleteBriefingMateri(request.ID, tx)

	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}
	tx.Commit()
	return true, err
}

// GetAll implements BriefingDefinition
func (briefing BriefingService) GetAll() (responses []models.BriefingResponse, err error) {
	return briefing.briefingRepo.GetAll()
}

// GetOne implements BriefingDefinition
func (briefing BriefingService) GetOne(id int64) (responses models.BriefingResponseGetOneString, status bool, err error) {
	dataBriefing, err := briefing.briefingRepo.GetOne(id)
	fmt.Println(dataBriefing)
	if dataBriefing.ID != 0 {
		fmt.Println("Bukan 0")

		materi, err := briefing.briefingMateri.GetOneBriefing(dataBriefing.ID)
		// peserta, err := briefing.MapPeserta.GetByIDBriefing(dataBriefing.ID)

		responses = models.BriefingResponseGetOneString{
			ID:             dataBriefing.ID,
			NoPelaporan:    dataBriefing.NoPelaporan,
			REGION:         dataBriefing.REGION,
			RGDESC:         dataBriefing.RGDESC,
			MAINBR:         dataBriefing.MAINBR,
			MBDESC:         dataBriefing.MBDESC,
			BRANCH:         dataBriefing.BRANCH,
			BRDESC:         dataBriefing.BRDESC,
			JenisPeserta:   dataBriefing.JenisPeserta,
			JabatanPeserta: dataBriefing.JabatanPeserta,
			// Peserta:       peserta,
			JumlahPeserta: dataBriefing.JumlahPeserta,
			ListPeserta:   dataBriefing.ListPeserta,
			MakerID:       dataBriefing.MakerID,
			MakerDesc:     dataBriefing.MakerDesc,
			MakerDate:     dataBriefing.MakerDate,
			LastMakerID:   dataBriefing.LastMakerID,
			LastMakerDesc: dataBriefing.LastMakerDesc,
			LastMakerDate: dataBriefing.LastMakerDate,
			Status:        dataBriefing.Status,
			Action:        dataBriefing.Action,
			Deleted:       dataBriefing.Deleted,
			Materi:        materi,
			CreatedAt:     dataBriefing.CreatedAt,
			UpdatedAt:     dataBriefing.UpdatedAt,
		}
		return responses, true, err
	}
	return responses, false, err
}

// StoreDraft implements BriefingDefinition
func (briefing BriefingService) StoreDraft(request models.BriefingRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := briefing.db.DB.Begin()

	//Input Briefing
	reqBriefing := &models.Briefing{
		NoPelaporan:   request.NoPelaporan,
		REGION:        request.REGION,
		RGDESC:        request.RGDESC,
		MAINBR:        request.MAINBR,
		MBDESC:        request.MBDESC,
		BRANCH:        request.BRANCH,
		BRDESC:        request.BRDESC,
		JenisPeserta:  request.JenisPeserta,
		JumlahPeserta: request.JumlahPeserta,
		MakerID:       request.MakerID,
		MakerDesc:     request.MakerDesc,
		MakerDate:     &timeNow,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "01a",
		Action:        "Draft",
		CreatedAt:     &timeNow,
	}

	dataBriefing, err := briefing.briefingRepo.Store(reqBriefing, tx)
	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	fmt.Println("dataBriefing", dataBriefing)

	//Input MapPesertaBriefing
	// if request.JenisPeserta == "perorangan" {
	if len(request.Peserta) != 0 {
		for _, value := range request.Peserta {
			_, err = briefing.MapPeserta.Store(&models.BriefingMapPeserta{
				IDBriefing:  dataBriefing.ID,
				PERNR:       value.PERNR,
				NamaPeserta: value.NamaPeserta,
				SteelTx:     value.SteelTx,
			}, tx)

			if err != nil {
				tx.Rollback()
				briefing.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}
	// }

	//Input Briefing Materi
	if len(request.Materi) != 0 {
		for _, value := range request.Materi {
			_, err = briefing.briefingMateri.Store(&models.BriefingMateri{
				BriefingID:        dataBriefing.ID,
				ActivityID:        value.ActivityID,
				SubActivityID:     value.SubActivityID,
				ProductID:         value.ProductID,
				JudulMateri:       value.JudulMateri,
				RiskIssueCode:     value.RiskIssueCode,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				CreatedAt:         &timeNow,
			}, tx)

			if err != nil {
				tx.Rollback()
				briefing.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// Store implements BriefingDefinition
func (briefing BriefingService) Store(request models.BriefingRequest) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := briefing.db.DB.Begin()

	fmt.Println("isi request.NoPelaporan: ", request.NoPelaporan)

	if request.NoPelaporan != "" {
		//Input Briefing
		reqBriefing := &models.Briefing{
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
			MakerID:        request.MakerID,
			MakerDesc:      request.MakerDesc,
			MakerDate:      &timeNow,
			LastMakerID:    request.LastMakerID,
			LastMakerDesc:  request.LastMakerDesc,
			LastMakerDate:  &timeNow,
			// Status:        "01a",
			// Action:        "Selesai",
			// change 23/11/2023 by panji
			Status:    request.Status,
			Action:    request.Action,
			CreatedAt: &timeNow,
		}

		dataBriefing, err := briefing.briefingRepo.Store(reqBriefing, tx)
		if err != nil {
			tx.Rollback()
			briefing.logger.Zap.Error(err)
			message = "Error transaction database !"
			return false, message, err
		}

		fmt.Println("dataBriefing", dataBriefing)

		//Input MapPesertaBriefing
		// if request.JenisPeserta == "perorangan" {
		// if len(request.Peserta) != 0 {
		// 	for _, value := range request.Peserta {
		// 		_, err = briefing.MapPeserta.Store(&models.BriefingMapPeserta{
		// 			IDBriefing:  dataBriefing.ID,
		// 			PERNR:       value.PERNR,
		// 			NamaPeserta: value.NamaPeserta,
		// 			SteelTx:     value.SteelTx,
		// 		}, tx)

		// 		if err != nil {
		// 			tx.Rollback()
		// 			briefing.logger.Zap.Error(err)
		// 			message = "Error transaction database !"
		// 			return false, message, err
		// 		}
		// 	}
		// } else {
		// 	tx.Rollback()
		// 	briefing.logger.Zap.Error(err)
		// 	message = "Data peserta kosong!"
		// 	return false, message, err
		// }
		// }

		//Input Briefing Materi
		if len(request.Materi) != 0 {
			for _, value := range request.Materi {
				_, err = briefing.briefingMateri.Store(&models.BriefingMateri{
					BriefingID:        dataBriefing.ID,
					ActivityID:        value.ActivityID,
					SubActivityID:     value.SubActivityID,
					ProductID:         value.ProductID,
					TitleMateries:     value.TitleMateries,
					JudulMateri:       value.JudulMateri,
					RiskIssueCode:     value.RiskIssueCode,
					RekomendasiMateri: value.RekomendasiMateri,
					MateriTambahan:    value.MateriTambahan,
					CreatedAt:         &timeNow,
				}, tx)

				if err != nil {
					tx.Rollback()
					briefing.logger.Zap.Error(err)
					message = "Error transaction database !"
					return false, message, err
				}
			}
		} else {
			tx.Rollback()
			briefing.logger.Zap.Error(err)
			message = "Data materi kosong !"
			return false, message, err
		}

		checkPekerja, err := briefing.msUker.CheckJumlahPekerja(request.BRANCH)
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
			message = "Input Data Berhasil !"
		}
		return true, message, err
	} else {
		tx.Rollback()
		// briefing.logger.Zap(err)
		message = "Data Gagal disimpan, nomor pelaporan kosong!"
		return false, message, err
	}
}

// WithTrx implements BriefingDefinition
func (briefing BriefingService) WithTrx(trxHandle *gorm.DB) BriefingService {
	briefing.briefingRepo = briefing.briefingRepo.WithTrx(trxHandle)
	return briefing
}

// GetData implements BriefingDefinition
func (briefing BriefingService) GetData() (responses []models.BriefingResponse, err error) {
	return briefing.briefingRepo.GetData()
}

// UpdateAllBrief implements BriefingDefinition
func (briefing BriefingService) UpdateAllBrief(request *models.BriefingResponseMaintain) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := briefing.db.DB.Begin()

	updateBriefing := &models.BriefingUpdateMateri{
		ID:             request.ID,
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
		LastMakerID:    request.LastMakerID,
		LastMakerDesc:  request.LastMakerDesc,
		LastMakerDate:  &timeNow,
		Deleted:        false,
		// Action:        "Update",
		// Status:        "02b",
		// change 23/11/2023 by panji
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
		"peserta",
		"jumlah_peserta",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"action",
		"status",
		"updated_at",
	}

	_, err = briefing.briefingRepo.UpdateAllBrief(updateBriefing, include, tx)

	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		message = "Error transaction database !"
		return false, message, err
	}

	//Input MapPesertaBriefing
	// if request.JenisPeserta == "perorangan" {
	// if len(request.Peserta) != 0 {
	// 	for _, value := range request.Peserta {
	// 		_, err = briefing.MapPeserta.Store(&models.BriefingMapPeserta{
	// 			ID:          value.ID,
	// 			IDBriefing:  request.ID,
	// 			PERNR:       value.PERNR,
	// 			NamaPeserta: value.NamaPeserta,
	// 			SteelTx:     value.SteelTx,
	// 		}, tx)

	// 		if err != nil {
	// 			tx.Rollback()
	// 			briefing.logger.Zap.Error(err)
	// 			message = "Error transaction database !"
	// 			return false, message, err
	// 		}
	// 	}
	// } else {
	// 	tx.Rollback()
	// 	briefing.logger.Zap.Error(err)
	// 	message = "Error transaction database !"
	// 	return false, message, err
	// }
	// }

	if len(request.Materi) != 0 {
		for _, value := range request.Materi {
			updateMateriBrief := &models.BriefingMateriUpdate{
				ID:                value.ID,
				BriefingID:        request.ID, //updateBriefing.ID
				ActivityID:        value.ActivityID,
				SubActivityID:     value.SubActivityID,
				ProductID:         value.ProductID,
				TitleMateries:     value.TitleMateries,
				JudulMateri:       value.JudulMateri,
				RiskIssueCode:     value.RiskIssueCode,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				UpdatedAt:         &timeNow,
			}

			// fmt.Println("ID =>", request.ID)
			// fmt.Println("request => ", updateMateriBrief)
			_, err = briefing.briefingMateri.Update(updateMateriBrief, tx)

			if err != nil {
				tx.Rollback()
				briefing.logger.Zap.Error(err)
				message = "Error transaction database !"
				return false, message, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			briefing.logger.Zap.Error(err)
			message = "Error transaction database !"
			return false, message, err
		}
	}

	checkPekerja, err := briefing.msUker.CheckJumlahPekerja(request.BRANCH)
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
		message = "Input Data Berhasil !"
	}

	return true, message, err
}

// UpdateDraft implements BriefingDefinition
func (briefing BriefingService) UpdateDraft(request *models.BriefingResponseMaintain) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := briefing.db.DB.Begin()

	updateBriefing := &models.BriefingUpdateMateri{
		ID:            request.ID,
		REGION:        request.REGION,
		RGDESC:        request.RGDESC,
		MAINBR:        request.MAINBR,
		MBDESC:        request.MBDESC,
		BRANCH:        request.BRANCH,
		BRDESC:        request.BRDESC,
		JenisPeserta:  request.JenisPeserta,
		JumlahPeserta: request.JumlahPeserta,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Deleted:       false,
		Status:        "01a",
		Action:        "Draft",
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"REGION",
		"RGDESC",
		"MAINBR",
		"MBDESC",
		"BRANCH",
		"BRDESC",
		"peserta",
		"jumlah_peserta",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"deleted",
		"action",
		"status",
		"updated_at",
	}

	_, err = briefing.briefingRepo.UpdateAllBrief(updateBriefing, include, tx)

	if err != nil {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}

	//Input MapPesertaBriefing
	// if request.JenisPeserta == "perorangan" {
	if len(request.Peserta) != 0 {
		for _, value := range request.Peserta {
			_, err = briefing.MapPeserta.Store(&models.BriefingMapPeserta{
				ID:          value.ID,
				IDBriefing:  request.ID,
				PERNR:       value.PERNR,
				NamaPeserta: value.NamaPeserta,
				SteelTx:     value.SteelTx,
			}, tx)

			if err != nil {
				tx.Rollback()
				briefing.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		briefing.logger.Zap.Error(err)
		return false, err
	}
	// }

	if len(request.Materi) != 0 {
		for _, value := range request.Materi {
			updateMateriBrief := &models.BriefingMateriUpdate{
				ID:                value.ID,
				BriefingID:        request.ID, //updateBriefing.ID
				ActivityID:        value.ActivityID,
				SubActivityID:     value.SubActivityID,
				ProductID:         value.ProductID,
				JudulMateri:       value.JudulMateri,
				RiskIssueCode:     value.RiskIssueCode,
				RekomendasiMateri: value.RekomendasiMateri,
				MateriTambahan:    value.MateriTambahan,
				UpdatedAt:         &timeNow,
			}

			_, err = briefing.briefingMateri.Update(updateMateriBrief, tx)

			if err != nil {
				tx.Rollback()
				briefing.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			briefing.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

// FilterBriefin implements BriefingDefinition
func (briefing BriefingService) FilterBriefing(requests models.BriefingFilterRequest) (responses []models.BriefingResponseData, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(requests.Page, requests.Limit, requests.Order, requests.Sort)
	requests.Offset = offset
	requests.Order = order
	requests.Sort = sort
	dataBrief, totalRows, totalData, err := briefing.briefingRepo.FilterBriefing(&requests)

	fmt.Println("totalRows ===>", totalRows)
	fmt.Println("totalData ===>", totalData)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataBrief {
		judul_materi, err := briefing.briefingRepo.GetJudulMateri(response.ID)
		if err != nil {
			briefing.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.BriefingResponseData{
			ID:          response.ID,
			NoPelaporan: response.NoPelaporan,
			UnitKerja:   response.UnitKerja,
			JudulMateri: judul_materi,
			StatusBrf:   response.Status,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	fmt.Println("pagination ==>", pagination)
	return responses, pagination, err
}

// GetDataWithPagination implements BriefingDefinition
func (brf BriefingService) GetDataWithPagination(request models.BriefingPagination) (responses []models.BriefingResponseData, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	dataBrief, totalRows, totalData, err := brf.briefingRepo.GetDataWithPagination(&request)
	if err != nil {
		brf.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataBrief {
		judul_materi, err := brf.briefingRepo.GetJudulMateri(response.ID)
		if err != nil {
			brf.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.BriefingResponseData{
			ID:          response.ID,
			NoPelaporan: response.NoPelaporan,
			UnitKerja:   response.UnitKerja,
			JudulMateri: judul_materi,
			StatusBrf:   response.Status,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetNoPelaporan implements BriefingDefinition
func (briefing BriefingService) GetNoPelaporan(request models.NoPelaporanRequest) (responses []models.NoPelaporanResponse, err error) {
	dataBriefing, err := briefing.briefingRepo.GetNoPelaporan(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataBriefing {
		responses = append(responses, models.NoPelaporanResponse{
			ORGEH:       request.ORGEH,
			NoPelaporan: response.NoPelaporan.String,
		})
	}

	return responses, err
}

func (briefing BriefingService) BriefingReportFilter(request models.BriefingFilterReport) (responses models.BriefingReportResponse, totalRows int64, err error) {
	var dataArr []models.BriefingFilterReportResponse
	// totalData := int64(0)
	dataBriefing, totalAktivitas, totalRows, err := briefing.briefingRepo.BriefingReportFilter(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataBriefing {
		dataArr = append(dataArr, models.BriefingFilterReportResponse{
			Id:    response.Id,
			Code:  response.Code,
			Name:  response.Name,
			Total: response.Total,
		})

		// totalData = totalAktivitas
	}

	responses = models.BriefingReportResponse{
		Data:      dataArr,
		TotalData: totalAktivitas,
	}

	fmt.Println("================== Total Data")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (briefing BriefingService) BriefingReportByUkerFilter(request models.BriefingFilterReportByUker) (responses []models.BriefingReportFilteredByUkerResponse, totalRows int64, err error) {
	dataBriefing, totalRows, err := briefing.briefingRepo.BriefingReportByUkerFilter(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataBriefing {
		responses = append(responses, models.BriefingReportFilteredByUkerResponse{
			REGION:          response.REGION,
			RGDESC:          response.RGDESC,
			MAINBR:          response.MAINBR,
			MBDESC:          response.MBDESC,
			BRANCH:          response.BRANCH,
			BRDESC:          response.BRDESC,
			TOTALBRIEFING:   response.TOTALBRIEFING,
			TOTALBRC:        response.TOTALBRC,
			PERCENTBRIEFING: response.PERCENTBRIEFING,
		})
	}

	fmt.Println("================== Responses")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (briefing BriefingService) BriefingReportFilterComplete(request models.BriefingFilterReport) (responses []models.BriefingFilterReportFinalResponse, totalRows int, err error) {
	dataBriefing, totalRows, err := briefing.briefingRepo.BriefingReportFilterComplete(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataBriefing {
		responses = append(responses, models.BriefingFilterReportFinalResponse{
			Id:        response.Id,
			Date:      response.Date,
			BRANCH:    response.BRANCH,
			BRDESC:    response.BRDESC,
			Activity:  response.Activity,
			Product:   response.Product,
			RiskIssue: response.RiskIssue,
		})
	}

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, totalRows, err
}

func (briefing BriefingService) BriefingReportDetail(request models.BriefingReportDetailRequest) (responses models.BriefingReportDetailResponse, err error) {
	dataBriefing, err := briefing.briefingRepo.BriefingReportDetail(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, err
	}

	responses = dataBriefing
	// for _, response := range dataBriefing {
	// 	responses = append(responses, models.BriefingFilterReportFinalResponse{
	// 		Id:       	   		response.Id.Int64,
	// 		Date:      	   		response.Date.String,
	// 		UnitKerja:          response.UnitKerja.String,
	// 		Activity: 		   	response.Activity.String,
	// 		Product: 		   	response.Product.String,
	// 		RiskIssue: 		   	response.RiskIssue.String,
	// 	})
	// }

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, err
}

func (briefing BriefingService) BriefingReportMateriList(request models.BriefingReportMateriRequest) (responses []models.BriefingDetailMateriResponse, err error) {
	dataBriefing, err := briefing.briefingRepo.BriefingReportMateriList(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataBriefing {
		minioLink, err := briefing.filemanager.GetFile(fileModel.FileManagerRequest{
			Subdir:   response.Path.String,
			Filename: response.Filename.String,
		})
		if err != nil {
			briefing.logger.Zap.Error(err)
			return responses, err
		}

		responses = append(responses, models.BriefingDetailMateriResponse{
			ID:           response.ID.Int64,
			NamaLampiran: response.NamaLampiran.String,
			Filename:     response.Filename.String,
			Path:         minioLink.MinioPath,
		})
	}

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, err
}

func (briefing BriefingService) BriefingReportFilterByUkerComplete(request models.BriefingFilterReportByUker) (responses []models.BriefingFilterReportFinalResponse, totalRows int, err error) {
	dataBriefing, totalRows, err := briefing.briefingRepo.BriefingReportFilterByUkerComplete(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataBriefing {
		responses = append(responses, models.BriefingFilterReportFinalResponse{
			Id:        response.Id.Int64,
			Date:      response.Date.String,
			BRANCH:    response.BRANCH.String,
			BRDESC:    response.BRDESC.String,
			Activity:  response.Activity.String,
			Product:   response.Product.String,
			RiskIssue: response.RiskIssue.String,
		})
	}

	fmt.Println("response service")
	fmt.Println(responses)

	return responses, totalRows, err
}

// BriefingReportList implements BriefingDefinition
func (briefing BriefingService) BriefingReportList(request models.BriefingReportListRequest) (responses []models.BriefingReportListFinalResponse, totalRows int, err error) {
	dataBriefing, totalRows, err := briefing.briefingRepo.BriefingReportList(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range dataBriefing {
		// materi, _ := briefing.briefingMateri.GetMateriReport(response.ID)
		// peserta, _ := briefing.MapPeserta.GetPesertaReport(response.ID)

		responses = append(responses, models.BriefingReportListFinalResponse{
			NoPelaporan:    response.NoPelaporan,
			RGDESC:         response.RGDESC,
			MBDESC:         response.MBDESC,
			BRANCH:         response.BRANCH,
			BRDESC:         response.BRDESC,
			JudulMateri:    response.JudulMateri,
			RiskEvent:      response.RiskEvent,
			RincianMateri:  response.RincianMateri,
			Aktivitas:      response.Aktivitas,
			JumlahPeserta:  response.JumlahPeserta,
			JabatanPeserta: response.JabatanPeserta,
			JenisPeserta:   response.JenisPeserta,
			Peserta:        response.Peserta,
			MakerID:        response.MakerID,
			Status:         response.Status,
		})
	}

	return responses, totalRows, err
}

func (briefing BriefingService) BriefingFrekuensiRpt(request models.FrekuensiBriefingRequest) (responses []models.FrekuensiBriefingResponse, totalRows int, err error) {
	frekuensiRpt, totalRows, err := briefing.briefingRepo.BriefingFrekuensiRpt(&request)

	if err != nil {
		briefing.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, value := range frekuensiRpt {
		responses = append(responses, models.FrekuensiBriefingResponse{
			Aktivitas: value.Aktivitas,
			Produk:    value.Produk,
			RiskEvent: value.RiskEvent,
			Jumlah:    value.Jumlah,
		})
	}

	return responses, totalRows, err
}
