package laporan

import (
	"errors"
	"riskmanagement/lib"
	models "riskmanagement/models/laporan"
	repository "riskmanagement/repository/laporan"
	"strconv"

	"gitlab.com/golang-package-library/logger"
)

type LaporanServicesDefinition interface {
	//GetLaporanHistoriTaskDataVerifikasi
	GetLaporanHistoriTaskDataVerifikasi(request models.HistoriTaskDataVerifikasiPagianted) (response []models.HistoriTaskDataVerifikasiResult, pagination lib.Pagination, err error)
	GetLaporanHistoriTaskDataVerifikasiDetail(request models.HistoriTaskDataVerifikasiDetailRequest) (response []models.HistoriTaskDataVerifikasiDetailResult, err error)
	GetLaporanHistoriTaskDataVerifikasiDownload(request models.HistoriTaskDataVerifikasiDownload) (response []models.HistoriTaskDataVerifikasiResult, err error)
	GetLaporanPerhitunganPersentasePenyelesaian(request models.PerhitunganPersentasePenyelesaianPagianted) (response []interface{}, pagination lib.Pagination, err error)
	GetLaporanPerhitunganPersentasePenyelesaianDownload(request models.PerhitunganPersentasePenyelesaianDownload) (response []interface{}, err error)
	GetRiskEventOnTaskList() (response []models.RiskEventOnTaskList, err error)
	GetMonitoringJob(request models.JobMonitoringRequest) (responses []models.JobMonitoringResponse, pagination lib.Pagination, err error)
	GetNamaJob(request *models.SearchNamaJobReq) (responses []models.SearchNamaJobRes, err error)
	GetActivityDaily(request models.ActivityDailyRequest) (responses []models.ActivityDailyResponse, pagination lib.Pagination, err error)
	GetActivityDailyDetail(request models.ActivityDailyDetailRequest) (responses []models.ActivityDailyDetailResponse, pagination lib.Pagination, err error)
	GetLaporanPerhitunganBriefing(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanBriefingResponse, pagination lib.Pagination, err error)
	GetLaporanPerhitunganBriefingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanBriefingResponse, err error)
	GetLaporanPerhitunganCoaching(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanCoachingResponse, pagination lib.Pagination, err error)
	GetLaporanPerhitunganCoachingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanCoachingResponse, err error)
	GetLaporanPerhitunganVerifikasi(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanVerifikasiResponse, pagination lib.Pagination, err error)
	GetLaporanPerhitunganVerifikasiDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanVerifikasiResponse, err error)
}

type LaporanServices struct {
	logger     logger.Logger
	repository repository.LaporanDefinition
}

func NewLaporanServices(
	logger logger.Logger,
	repository repository.LaporanDefinition,
) LaporanServicesDefinition {
	return LaporanServices{
		logger:     logger,
		repository: repository,
	}
}

func (l LaporanServices) GetRiskEventOnTaskList() (response []models.RiskEventOnTaskList, err error) {
	data, err := l.repository.GetRiskEventOnTaskList()

	if err != nil {
		l.logger.Zap.Error(err)
		return nil, err
	}

	return data, nil
}

func (l LaporanServices) GetLaporanHistoriTaskDataVerifikasiDetail(request models.HistoriTaskDataVerifikasiDetailRequest) (responses []models.HistoriTaskDataVerifikasiDetailResult, err error) {
	data, err := l.repository.GetListVerifikasiByTaskID(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, err
	}

	// l.logger.Zap.Info(data)

	for _, response := range data {
		var persenPerbaikan float64

		if response.JumlahDataAnomali != 0 {
			selesai, _ := strconv.ParseFloat(response.StatusPerbaikanSelesai, 64)
			persenPerbaikan = selesai / float64(response.JumlahDataAnomali) * 100
		} else {
			persenPerbaikan = 0
		}

		responses = append(responses, models.HistoriTaskDataVerifikasiDetailResult{
			Branch:                 response.Branch,
			UnitKerja:              response.UnitKerja,
			Kanwil:                 response.Kanwil,
			Kanca:                  response.Kanca,
			NoPelaporan:            response.NoPelaporan,
			Aktifitas:              response.Aktifitas,
			SubAktifitas:           response.SubAktifitas,
			InformasiLainnya:       response.InformasiLainnya,
			StatusPerbaikan:        response.StatusPerbaikan,
			Maker:                  response.Maker,
			RiskIssueId:            response.RiskIssueId,
			RiskIssue:              response.RiskIssue,
			HasilVerifikasi:        response.HasilVerifikasi,
			JumlahDataDiverifikasi: response.JumlahDataAnomali,
			ButuhPerbaikan:         response.ButuhPerbaikan,
			YangHarusDiperbaiki:    response.YangHarusDiperbaiki,
			RtlUker:                response.RtlUker,
			StatusPerbaikanSelesai: response.StatusPerbaikanSelesai,
			StatusPerbaikanProses:  response.StatusPerbaikanProses,
			PersentasePerbaikan:    int64(persenPerbaikan),
			BatasWaktuPerbaikan:    response.BatasWaktuPerbaikan,
			IndikasiFraud:          response.IndikasiFraud,
		})
	}

	return responses, err

}

func (l LaporanServices) GetLaporanHistoriTaskDataVerifikasi(request models.HistoriTaskDataVerifikasiPagianted) (responses []models.HistoriTaskDataVerifikasiResult, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	l.logger.Zap.Info(request)
	data, totalData, totalRows, err := l.repository.GetTasklist(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		responses = append(responses, models.HistoriTaskDataVerifikasiResult{
			ID:             response.ID,
			PN:             response.PN,
			Nama:           response.Nama,
			NoTasklist:     response.NoTasklist,
			NamaTasklist:   response.NamaTasklist,
			Region:         response.Region,
			Kanwil:         response.Kanwil,
			Mainbr:         response.Mainbr,
			Kanca:          response.Kanca,
			Branch:         response.Branch,
			Uker:           response.Uker,
			Aktifitas:      response.Aktifitas,
			Product:        response.Product,
			Indikator:      response.Indikator,
			JenisTask:      response.JenisTask,
			Period:         response.Period,
			TanggalMulai:   response.TanggalMulai,
			TanggalAkhir:   response.TanggalAkhir,
			StatusApproval: response.StatusApproval,
			Status:         response.Status,
			Kegiatan:       response.Kegiatan,
			RiskIssue:      response.RiskIssue,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (l LaporanServices) GetLaporanHistoriTaskDataVerifikasiDownload(request models.HistoriTaskDataVerifikasiDownload) (responses []models.HistoriTaskDataVerifikasiResult, err error) {

	l.logger.Zap.Info(request)
	data, err := l.repository.GetTasklistDownload(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range data {
		// l.logger.Zap.Error(response)
		responses = append(responses, models.HistoriTaskDataVerifikasiResult{
			ID:             response.ID,
			PN:             response.PN,
			NoTasklist:     response.NoTasklist,
			NamaTasklist:   response.NamaTasklist,
			Nama:           response.Nama,
			Kanwil:         response.Kanwil,
			Kanca:          response.Kanca,
			Uker:           response.Uker,
			Aktifitas:      response.Aktifitas,
			Product:        response.Product,
			Kegiatan:       response.Kegiatan,
			RiskIssue:      response.RiskIssue,
			Indikator:      response.Indikator,
			JenisTask:      response.JenisTask,
			TanggalMulai:   response.TanggalMulai,
			TanggalAkhir:   response.TanggalAkhir,
			StatusApproval: response.StatusApproval,
			Status:         response.Status,
		})
	}

	return responses, err
}

func (l LaporanServices) GetLaporanPerhitunganPersentasePenyelesaian(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []interface{}, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	request.Page = page
	request.Limit = limit

	var totalData int

	if request.JenisReport == "0" {
		data, totalRows, totalData, err := l.repository.GetPerhitunganPersentasePenyelesaianPerPekerja(request)

		if err != nil {
			l.logger.Zap.Error(err)
			return nil, pagination, err
		}

		for _, response := range data {
			// l.logger.Zap.Info(response.Aktifitas)
			var persenVerif float64
			persenVerif = 0

			if response.JumlahDataAnomali != 0 {
				persenVerif = float64(response.JumlahDataVerifikasi) / float64(response.JumlahDataAnomali) * 100
			}

			responses = append(responses, models.LaporanPerPekerjaResult{
				ID:                          response.ID,
				Pn:                          response.Pn,
				Nama:                        response.Nama,
				Aktifitas:                   response.Aktifitas,
				Produk:                      response.Produk,
				Indikator:                   response.Indikator,
				JenisTask:                   response.JenisTask,
				TanggalMulai:                response.TanggalMulai,
				TanggalSelesai:              response.TanggalSelesai,
				JumlahDataAnomali:           response.JumlahDataAnomali,
				JumlahDataVerifikasi:        response.JumlahDataVerifikasi,
				JumlahDataPerluTindaklanjut: response.JumlahDataPerluTindaklanjut,
				JumlahDataSudahTindaklanjut: response.JumlahDataSudahTindaklanjut,
				PersenSudahVerifikasi:       persenVerif,
				PersenSudahTindaklanjut:     0,
			})
			// l.logger.Zap.Info(responses)
		}

		pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

		return responses, pagination, err
	} else if request.JenisReport == "1" {
		data, totalRows, err := l.repository.GetPerhitunganPersentasePenyelesaianPerUker(request)

		if err != nil {
			l.logger.Zap.Error(err)
			return nil, pagination, err
		}

		for _, response := range data {
			responses = append(responses, models.LaporanPerUkerResult{
				ID:                          response.ID,
				Kanwil:                      response.Kanwil,
				Kanca:                       response.Kanca,
				Uker:                        response.Uker,
				Aktifitas:                   response.Aktifitas,
				Produk:                      response.Produk,
				Indikator:                   response.Indikator,
				JenisTask:                   response.JenisTask,
				TanggalMulai:                response.TanggalMulai,
				TanggalSelesai:              response.TanggalSelesai,
				JumlahDataAnomali:           response.JumlahDataAnomali,
				JumlahDataVerifikasi:        response.JumlahDataVerifikasi,
				JumlahDataPerluTidaklanjut:  response.JumlahDataPerluTidaklanjut,
				JumlahDataSudahTindaklanjut: response.JumlahDataSudahTindaklanjut,
				PersenSudahVerifikasi:       0,
				PersenSudahTindaklanjut:     0,
			})
		}

		pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
		return responses, pagination, err
	} else if request.JenisReport == "2" {
		data, totalRows, err := l.repository.GetPerhitunganPersentasePenyelesaianPerPekerjaUker(request)

		if err != nil {
			l.logger.Zap.Error(err)
			return nil, pagination, err
		}

		for _, response := range data {
			responses = append(responses, models.LaporanPerPekerjaUkerResult{
				ID:                         response.ID,
				Pn:                         response.Pn,
				Nama:                       response.Nama,
				Kanwil:                     response.Kanwil,
				Kanca:                      response.Kanca,
				Uker:                       response.Uker,
				Aktifitas:                  response.Aktifitas,
				Produk:                     response.Produk,
				Indikator:                  response.Indikator,
				JenisTask:                  response.JenisTask,
				TanggalMulai:               response.TanggalMulai,
				TanggalSelesai:             response.TanggalSelesai,
				JumlahDataAnomali:          response.JumlahDataAnomali,
				JumlahDataVerifikasi:       response.JumlahDataVerifikasi,
				JumlahDataPerluTidaklanjut: response.JumlahDataPerluTidaklanjut,
				PersenSudahVerifikasi:      response.PersenSudahVerifikasi,
				PersenSudahTindaklanjut:    response.PersenSudahTindaklanjut,
			})
		}

		pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
		return responses, pagination, err
	} else {
		return nil, pagination, errors.New("invalid jenis report")
	}
}

func (l LaporanServices) GetMonitoringJob(request models.JobMonitoringRequest) (responses []models.JobMonitoringResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalData, totalRows, err := l.repository.GetMonitoringJob(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		responses = append(responses, models.JobMonitoringResponse{
			Tanggal:         response.Tanggal,
			NamaJob:         response.NamaJob,
			Proses:          response.Proses,
			StatusProses:    response.StatusProses,
			DeskripsiStatus: response.DeskripsiStatus,
		})
		l.logger.Zap.Info(responses)
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (l LaporanServices) GetNamaJob(request *models.SearchNamaJobReq) (responses []models.SearchNamaJobRes, err error) {
	dataPekerja, err := l.repository.GetNamaJob(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataPekerja {
		responses = append(responses, models.SearchNamaJobRes{
			Name: response.Name,
		})
	}

	return responses, err
}

func (l LaporanServices) GetActivityDaily(request models.ActivityDailyRequest) (responses []models.ActivityDailyResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalData, totalRows, err := l.repository.GetActivityDaily(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		listUkerReq := models.UkerListRequest{
			PERNR: response.PERNR,
		}

		listUkerData, err := l.repository.GetUkerList(listUkerReq)

		if err != nil {
			l.logger.Zap.Error(err)
			return responses, pagination, err
		}

		// persenReq := models.PersentaseTotalRequest{
		// 	PERNR:      response.PERNR,
		// 	Persentase: request.Persentase,
		// }

		// persentase, err := l.repository.GetTasklistPersentase(persenReq)

		// if err != nil {
		// 	l.logger.Zap.Error(err)
		// 	return responses, pagination, err
		// }

		floatAngka, err := strconv.ParseFloat(response.Persentase, 64)
		if err != nil {
			l.logger.Zap.Error(err)
			return responses, pagination, err
		}

		responses = append(responses, models.ActivityDailyResponse{
			PERNR:        response.PERNR,
			Nama:         response.Nama,
			Kanwil:       response.Kanwil,
			UkerKelolaan: listUkerData,
			Kegiatan:     response.Kegiatan,
			ActivityID:   response.ActivityID,
			ProductID:    response.ProductID,
			RiskIssueID:  response.RiskIssueID,
			Sample:       response.Sample,
			Progres:      response.Progres,
			Persentase:   int64(floatAngka),
			RiskEvent:    response.RiskEvent,
		})
		l.logger.Zap.Info(responses)
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (l LaporanServices) GetActivityDailyDetail(request models.ActivityDailyDetailRequest) (responses []models.ActivityDailyDetailResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	data, totalData, totalRows, err := l.repository.GetActivityDailyDetail(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		// listUkerReq := models.UkerListRequest{
		// 	PERNR: response.PERNR,
		// }

		// listUkerData, err := l.repository.GetUkerList(listUkerReq)

		// if err != nil {
		// 	l.logger.Zap.Error(err)
		// 	return responses, pagination, err
		// }

		responses = append(responses, models.ActivityDailyDetailResponse{
			PERNR:  response.PERNR,
			Nama:   response.Nama,
			Kanwil: response.Kanwil,
			Kanca:  response.Kanca,
			// UkerKelolaan: listUkerData,
			UnitKerja:       response.UnitKerja,
			Kegiatan:        response.Kegiatan,
			RiskEvent:       response.RiskEvent,
			TaskType:        response.TaskType,
			Period:          response.Period,
			Sample:          response.Sample,
			Progres:         response.Progres,
			Persentase:      response.Persentase,
			AssignedCreated: response.AssignedCreated,
			StartDate:       response.StartDate,
			EndDate:         response.EndDate,
		})
		l.logger.Zap.Info(responses)
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (l LaporanServices) GetLaporanPerhitunganPersentasePenyelesaianDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []interface{}, err error) {
	if request.JenisReport == "0" {
		data, err := l.repository.GetPerhitunganPersentasePenyelesaianPerPekerjaDownload(request)

		if err != nil {
			l.logger.Zap.Error(err)
			return nil, err
		}

		for _, response := range data {
			// l.logger.Zap.Info(response.Aktifitas)
			var persenVerif float64
			persenVerif = 0

			if response.JumlahDataAnomali != 0 {
				persenVerif = float64(response.JumlahDataVerifikasi) / float64(response.JumlahDataAnomali) * 100
			}
			responses = append(responses, models.LaporanPerPekerjaResult{
				ID:                          response.ID,
				Pn:                          response.Pn,
				Nama:                        response.Nama,
				Aktifitas:                   response.Aktifitas,
				Produk:                      response.Produk,
				Indikator:                   response.Indikator,
				JenisTask:                   response.JenisTask,
				TanggalMulai:                response.TanggalMulai,
				TanggalSelesai:              response.TanggalSelesai,
				JumlahDataAnomali:           response.JumlahDataAnomali,
				JumlahDataVerifikasi:        response.JumlahDataVerifikasi,
				JumlahDataPerluTindaklanjut: response.JumlahDataPerluTindaklanjut,
				JumlahDataSudahTindaklanjut: response.JumlahDataSudahTindaklanjut,
				PersenSudahVerifikasi:       persenVerif,
				PersenSudahTindaklanjut:     0,
			})
			// l.logger.Zap.Info(responses)
		}

		return responses, err
	} else if request.JenisReport == "1" {
		data, err := l.repository.GetPerhitunganPersentasePenyelesaianPerUkerDownload(request)

		if err != nil {
			l.logger.Zap.Error(err)
			return nil, err
		}

		for _, response := range data {
			responses = append(responses, models.LaporanPerUkerResult{
				ID:                          response.ID,
				Kanwil:                      response.Kanwil,
				Kanca:                       response.Kanca,
				Uker:                        response.Uker,
				Aktifitas:                   response.Aktifitas,
				Produk:                      response.Produk,
				Indikator:                   response.Indikator,
				JenisTask:                   response.JenisTask,
				TanggalMulai:                response.TanggalMulai,
				TanggalSelesai:              response.TanggalSelesai,
				JumlahDataAnomali:           response.JumlahDataAnomali,
				JumlahDataVerifikasi:        response.JumlahDataVerifikasi,
				JumlahDataPerluTidaklanjut:  response.JumlahDataPerluTidaklanjut,
				JumlahDataSudahTindaklanjut: response.JumlahDataSudahTindaklanjut,
				PersenSudahVerifikasi:       0,
				PersenSudahTindaklanjut:     0,
			})
		}

		return responses, err
	} else if request.JenisReport == "2" {
		data, err := l.repository.GetPerhitunganPersentasePenyelesaianPerPekerjaUkerDownload(request)

		if err != nil {
			l.logger.Zap.Error(err)
			return nil, err
		}

		for _, response := range data {
			responses = append(responses, models.LaporanPerPekerjaUkerResult{
				ID:                         response.ID,
				Pn:                         response.Pn,
				Nama:                       response.Nama,
				Kanwil:                     response.Kanwil,
				Kanca:                      response.Kanca,
				Uker:                       response.Uker,
				Aktifitas:                  response.Aktifitas,
				Produk:                     response.Produk,
				Indikator:                  response.Indikator,
				JenisTask:                  response.JenisTask,
				TanggalMulai:               response.TanggalMulai,
				TanggalSelesai:             response.TanggalSelesai,
				JumlahDataAnomali:          response.JumlahDataAnomali,
				JumlahDataVerifikasi:       response.JumlahDataVerifikasi,
				JumlahDataPerluTidaklanjut: response.JumlahDataPerluTidaklanjut,
				PersenSudahVerifikasi:      0,
				PersenSudahTindaklanjut:    0,
			})
		}

		return responses, err
	} else {
		return nil, errors.New("invalid jenis report")
	}
}

// 28-08-2023
func (l LaporanServices) GetLaporanPerhitunganVerifikasi(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanVerifikasiResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	request.Page = page
	request.Limit = limit

	data, totalRow, totalData, err := l.repository.GetPerhitunganVerifikasi(request)
	if err != nil {
		l.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		var persenVerif float64
		var persenTindakLanjut float64

		if response.JumlahDataAnomali != 0 {
			persenVerif = float64(response.JumlahDataVerifikasi) / float64(response.JumlahDataAnomali) * 100
		} else {
			persenVerif = 0
		}

		// if response.ButuhPerbaikan == "Ya" {
		// 	persenTindakLanjut = (float64(response.JumlahDataSudahTindaklanjut) / float64(response.JumlahDataPerluTindaklanjut)) * 100
		// } else {
		// 	persenTindakLanjut = 100
		// }

		if response.JumlahDataPerluTindaklanjut != 0 {
			persenTindakLanjut = (float64(response.JumlahDataSudahTindaklanjut) / float64(response.JumlahDataVerifikasi)) * 100
		} else {
			persenTindakLanjut = 100
		}

		//note Jumlah
		responses = append(responses, models.LaporanVerifikasiResponse{
			ID:                          response.ID,
			Pn:                          response.Pn,
			Nama:                        response.Nama,
			Kanwil:                      response.Kanwil,
			NoTasklist:                  response.NoTasklist,
			NamaTasklist:                response.NamaTasklist,
			Kanca:                       response.Kanca,
			Uker:                        response.Uker,
			JenisTask:                   response.JenisTask,
			Kegiatan:                    response.Kegiatan,
			Aktifitas:                   response.Aktifitas,
			Produk:                      response.Produk,
			RiskIssue:                   response.RiskIssue,
			Indikator:                   response.Indikator,
			TanggalMulai:                response.TanggalMulai,
			TanggalSelesai:              response.TanggalSelesai,
			JumlahDataAnomali:           response.JumlahDataAnomali,
			JumlahDataVerifikasi:        response.JumlahDataVerifikasi,
			PersenSudahVerifikasi:       persenVerif,
			JumlahDataPerluTindaklanjut: response.JumlahDataPerluTindaklanjut,
			JumlahDataSudahTindaklanjut: response.JumlahDataSudahTindaklanjut,
			PersenSudahTindaklanjut:     persenTindakLanjut,
			JumlahKegiatanDilakukan:     response.JumlahKegiatanDilakukan,
		})

		// l.logger.Zap.Info(responses)
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRow, totalData)
	return responses, pagination, nil
}

// GetLaporanPerhitunganVerifikasiDownload implements LaporanServicesDefinition.
func (l LaporanServices) GetLaporanPerhitunganVerifikasiDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanVerifikasiResponse, err error) {
	data, err := l.repository.GetPerhitunganVerifikasiDownload(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return nil, err
	}

	for _, response := range data {
		var persenVerif float64
		var persenTindakLanjut float64

		if response.JumlahDataAnomali != 0 {
			persenVerif = float64(response.JumlahDataVerifikasi) / float64(response.JumlahDataAnomali) * 100
		} else {
			persenVerif = 0
		}

		if response.JumlahDataPerluTindaklanjut != 0 {
			persenTindakLanjut = (float64(response.JumlahDataSudahTindaklanjut) / float64(response.JumlahDataVerifikasi)) * 100
		} else {
			persenTindakLanjut = 100
		}

		responses = append(responses, models.LaporanVerifikasiResponse{
			ID:                          response.ID,
			Pn:                          response.Pn,
			Nama:                        response.Nama,
			NoTasklist:                  response.NoTasklist,
			NamaTasklist:                response.NamaTasklist,
			Kanwil:                      response.Kanwil,
			Kanca:                       response.Kanca,
			Uker:                        response.Uker,
			JenisTask:                   response.JenisTask,
			Kegiatan:                    response.Kegiatan,
			Aktifitas:                   response.Aktifitas,
			RiskIssue:                   response.RiskIssue,
			Produk:                      response.Produk,
			Indikator:                   response.Indikator,
			TanggalMulai:                response.TanggalMulai,
			TanggalSelesai:              response.TanggalSelesai,
			JumlahDataAnomali:           response.JumlahDataAnomali,
			JumlahDataVerifikasi:        response.JumlahDataVerifikasi,
			PersenSudahVerifikasi:       persenVerif,
			JumlahDataPerluTindaklanjut: response.JumlahDataPerluTindaklanjut,
			JumlahDataSudahTindaklanjut: response.JumlahDataSudahTindaklanjut,
			PersenSudahTindaklanjut:     persenTindakLanjut,
			JumlahKegiatanDilakukan:     response.JumlahKegiatanDilakukan,
		})

		// l.logger.Zap.Info(responses)
	}

	return responses, err
}

// GetLaporanPerhitunganBriefing implements LaporanServicesDefinition.
func (l LaporanServices) GetLaporanPerhitunganBriefing(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanBriefingResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	request.Page = page
	request.Limit = limit

	data, totalRow, totalData, err := l.repository.GetPerhitunganBriefing(request)
	if err != nil {
		l.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		responses = append(responses, models.LaporanBriefingResponse{
			ID:                              response.ID,
			Pn:                              response.Pn,
			Nama:                            response.Nama,
			Kanwil:                          response.Kanwil,
			Kanca:                           response.Kanca,
			Uker:                            response.Uker,
			JenisTask:                       response.JenisTask,
			Kegiatan:                        response.Kegiatan,
			Aktifitas:                       response.Aktifitas,
			Produk:                          response.Produk,
			RiskIssue:                       response.RiskIssue,
			Indikator:                       response.Indikator,
			TanggalMulai:                    response.TanggalMulai,
			TanggalSelesai:                  response.TanggalSelesai,
			JumlahDataAnomali:               "-",
			JumlahDataVerifikasi:            "-",
			PersenSudahVerifikasi:           "-",
			JumlahDataPerluTindaklanjut:     "-",
			JumlahDataYangSudahTindaklanjut: "-",
			PersenSudahTindaklanjut:         "-",
			JumlahKegiatanDilakukan:         response.JumlahKegiatanDilakukan,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRow, totalData)
	return responses, pagination, nil
}

// GetLaporanPerhitunganBriefingDownload implements LaporanServicesDefinition.
func (l LaporanServices) GetLaporanPerhitunganBriefingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanBriefingResponse, err error) {
	data, err := l.repository.GetPerhitunganBriefingDownload(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return nil, err
	}

	for _, response := range data {
		responses = append(responses, models.LaporanBriefingResponse{
			ID:                              response.ID,
			Pn:                              response.Pn,
			Nama:                            response.Nama,
			Kanwil:                          response.Kanwil,
			Kanca:                           response.Kanca,
			Uker:                            response.Uker,
			JenisTask:                       response.JenisTask,
			Kegiatan:                        response.Kegiatan,
			Aktifitas:                       response.Aktifitas,
			Produk:                          response.Produk,
			RiskIssue:                       response.RiskIssue,
			Indikator:                       response.Indikator,
			TanggalMulai:                    response.TanggalMulai,
			TanggalSelesai:                  response.TanggalSelesai,
			JumlahDataAnomali:               "-",
			JumlahDataVerifikasi:            "-",
			PersenSudahVerifikasi:           "-",
			JumlahDataPerluTindaklanjut:     "-",
			JumlahDataYangSudahTindaklanjut: "-",
			PersenSudahTindaklanjut:         "-",
			JumlahKegiatanDilakukan:         response.JumlahKegiatanDilakukan,
		})
	}

	return responses, err
}

// GetLaporanPerhitunganCoaching implements LaporanServicesDefinition.
func (l LaporanServices) GetLaporanPerhitunganCoaching(request models.PerhitunganPersentasePenyelesaianPagianted) (responses []models.LaporanCoachingResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	request.Page = page
	request.Limit = limit

	data, totalRow, totalData, err := l.repository.GetPerhitunganCoaching(request)
	if err != nil {
		l.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range data {
		responses = append(responses, models.LaporanCoachingResponse{
			ID:                              response.ID,
			Pn:                              response.Pn,
			Nama:                            response.Nama,
			Kanwil:                          response.Kanwil,
			Kanca:                           response.Kanca,
			Uker:                            response.Uker,
			JenisTask:                       response.JenisTask,
			Kegiatan:                        response.Kegiatan,
			Aktifitas:                       response.Aktifitas,
			RiskIssue:                       response.RiskIssue,
			Produk:                          response.Produk,
			Indikator:                       response.Indikator,
			TanggalMulai:                    response.TanggalMulai,
			TanggalSelesai:                  response.TanggalSelesai,
			JumlahDataAnomali:               "-",
			JumlahDataVerifikasi:            "-",
			PersenSudahVerifikasi:           "-",
			JumlahDataPerluTindaklanjut:     "-",
			JumlahDataYangSudahTindaklanjut: "-",
			PersenSudahTindaklanjut:         "-",
			JumlahKegiatanDilakukan:         response.JumlahKegiatanDilakukan,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRow, totalData)
	return responses, pagination, nil
}

// GetLaporanPerhitunganCoachingDownload implements LaporanServicesDefinition.
func (l LaporanServices) GetLaporanPerhitunganCoachingDownload(request models.PerhitunganPersentasePenyelesaianDownload) (responses []models.LaporanCoachingResponse, err error) {
	data, err := l.repository.GetPerhitunganCoachingDownload(request)

	if err != nil {
		l.logger.Zap.Error(err)
		return nil, err
	}

	for _, response := range data {
		responses = append(responses, models.LaporanCoachingResponse{
			ID:                              response.ID,
			Pn:                              response.Pn,
			Nama:                            response.Nama,
			Kanwil:                          response.Kanwil,
			Kanca:                           response.Kanca,
			Uker:                            response.Uker,
			JenisTask:                       response.JenisTask,
			Kegiatan:                        response.Kegiatan,
			Aktifitas:                       response.Aktifitas,
			Produk:                          response.Produk,
			RiskIssue:                       response.RiskIssue,
			Indikator:                       response.Indikator,
			TanggalMulai:                    response.TanggalMulai,
			TanggalSelesai:                  response.TanggalSelesai,
			JumlahDataAnomali:               "-",
			JumlahDataVerifikasi:            "-",
			PersenSudahVerifikasi:           "-",
			JumlahDataPerluTindaklanjut:     "-",
			JumlahDataYangSudahTindaklanjut: "-",
			PersenSudahTindaklanjut:         "-",
			JumlahKegiatanDilakukan:         response.JumlahKegiatanDilakukan,
		})
	}

	return responses, err
}
