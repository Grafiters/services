package laporan

import (
	"bytes"
	"net/http"
	"riskmanagement/lib"
	models "riskmanagement/models/laporan"
	service "riskmanagement/services/laporan"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"gitlab.com/golang-package-library/logger"
)

type LaporanController struct {
	logger  logger.Logger
	service service.LaporanServicesDefinition
}

func NewLaporanController(LaporanService service.LaporanServicesDefinition, logger logger.Logger) LaporanController {
	return LaporanController{
		logger:  logger,
		service: LaporanService,
	}
}

//GetLaporanHistoriTaskDataVerifikasi

func (laporan LaporanController) GetLaporanPerhitunganPersentasePenyelesaian(c *gin.Context) {

	data := models.PerhitunganPersentasePenyelesaianPagianted{}

	if err := c.Bind(&data); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	result, pagination, err := laporan.service.GetLaporanPerhitunganPersentasePenyelesaian(data)
	if err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	// fmt.Println("cek", pagination)
	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, pagination)
	// lib.ReturnToJson(c, 200, "200", "tes json", result)
}

func (laporan LaporanController) GetLaporanHistoriTaskDataVerifikasi(c *gin.Context) {
	data := models.HistoriTaskDataVerifikasiPagianted{}

	if err := c.Bind(&data); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	laporan.logger.Zap.Info(data)

	result, totalRow, err := laporan.service.GetLaporanHistoriTaskDataVerifikasi(data)
	if err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, totalRow)
}

func (laporan LaporanController) GetMonitoringJob(c *gin.Context) {
	data := models.JobMonitoringRequest{}

	if err := c.Bind(&data); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	result, totalRow, err := laporan.service.GetMonitoringJob(data)
	if err != nil {
		laporan.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, totalRow)
}

func (laporan LaporanController) GetNamaJob(c *gin.Context) {
	requests := models.SearchNamaJobReq{}

	if err := c.Bind(&requests); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai", "")
		return
	}

	data, err := laporan.service.GetNamaJob(&requests)
	if err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "500", "Internal Error", err.Error())
		return
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", data)
}

func (laporan LaporanController) GetActivityDaily(c *gin.Context) {
	data := models.ActivityDailyRequest{}

	if err := c.Bind(&data); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	result, totalRow, err := laporan.service.GetActivityDaily(data)
	if err != nil {
		laporan.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, totalRow)
}

func (laporan LaporanController) GetActivityDailyDetail(c *gin.Context) {
	requests := models.ActivityDailyDetailRequest{}

	if err := c.ShouldBindJSON(&requests); err != nil {
		laporan.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai : "+err.Error(), "")
		return
	}
	data, totalRow, err := laporan.service.GetActivityDailyDetail(requests)
	if err != nil {
		laporan.logger.Zap.Error()
		lib.ReturnToJson(c, 200, "500", "Internal Error", "")
	}

	lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", data, totalRow)
}

func (laporan LaporanController) GetLaporanHistoriTaskDataVerifikasiDetail(c *gin.Context) {
	request := models.HistoriTaskDataVerifikasiDetailRequest{}

	if err := c.Bind(&request); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input Tidak Sesuai: "+err.Error(), "")
		return
	}

	// taskID := c.Param("taskID")

	// i, err := strconv.ParseInt(taskID, 10, 64)
	// if err != nil {
	// 	lib.ReturnToJson(c, 400, "400", "Inquery data berhasil", err)
	// 	return
	// }

	result, err := laporan.service.GetLaporanHistoriTaskDataVerifikasiDetail(request)
	if err != nil {
		laporan.logger.Zap.Error(err)
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (laporan LaporanController) GetLaporanHistoriTaskDataVerifikasiDownload(c *gin.Context) {

	data := models.HistoriTaskDataVerifikasiDownload{}

	if err := c.Bind(&data); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	laporan.logger.Zap.Info(data)

	result, err := laporan.service.GetLaporanHistoriTaskDataVerifikasiDownload(data)
	if err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	file := excelize.NewFile()

	sheet1Name := "Sheet One"
	file.SetSheetName(file.GetSheetName(1), sheet1Name)

	file.SetCellValue(sheet1Name, "A1", "No")
	file.SetCellValue(sheet1Name, "B1", "PN")
	file.SetCellValue(sheet1Name, "C1", "Nama")
	file.SetCellValue(sheet1Name, "D1", "No Tasklist")
	file.SetCellValue(sheet1Name, "E1", "Nama Tasklist")
	file.SetCellValue(sheet1Name, "F1", "Kanwil")
	file.SetCellValue(sheet1Name, "G1", "Kanca")
	file.SetCellValue(sheet1Name, "H1", "Uker Supervisi")
	file.SetCellValue(sheet1Name, "I1", "Aktifitas")
	file.SetCellValue(sheet1Name, "J1", "Produk")
	file.SetCellValue(sheet1Name, "K1", "Kegiatan")
	file.SetCellValue(sheet1Name, "L1", "Risk Event")
	file.SetCellValue(sheet1Name, "M1", "Indikator")
	file.SetCellValue(sheet1Name, "N1", "Tanggal Mulai")
	file.SetCellValue(sheet1Name, "O1", "Jenis Task")
	file.SetCellValue(sheet1Name, "P1", "Tanggal Selesai")
	file.SetCellValue(sheet1Name, "Q1", "Status Approval")
	file.SetCellValue(sheet1Name, "R1", "Status")
	row := 2
	baris := 1
	for _, response := range result {
		laporan.logger.Zap.Info("LOOP REPORT")
		laporan.logger.Zap.Info(response.Nama)
		file.SetCellValue(sheet1Name, "A"+strconv.Itoa(row), strconv.Itoa(baris))
		file.SetCellValue(sheet1Name, "B"+strconv.Itoa(row), response.PN)
		file.SetCellValue(sheet1Name, "C"+strconv.Itoa(row), response.Nama)
		file.SetCellValue(sheet1Name, "D"+strconv.Itoa(row), response.NoTasklist)
		file.SetCellValue(sheet1Name, "E"+strconv.Itoa(row), response.Nama)
		file.SetCellValue(sheet1Name, "F"+strconv.Itoa(row), response.Kanwil)
		file.SetCellValue(sheet1Name, "G"+strconv.Itoa(row), response.Kanca)
		file.SetCellValue(sheet1Name, "H"+strconv.Itoa(row), response.Uker)
		file.SetCellValue(sheet1Name, "I"+strconv.Itoa(row), response.Aktifitas)
		file.SetCellValue(sheet1Name, "J"+strconv.Itoa(row), response.Product)
		file.SetCellValue(sheet1Name, "K"+strconv.Itoa(row), response.Kegiatan)
		file.SetCellValue(sheet1Name, "L"+strconv.Itoa(row), response.RiskIssue)
		file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), response.Indikator)
		file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), response.TanggalMulai)
		file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), response.JenisTask)
		file.SetCellValue(sheet1Name, "P"+strconv.Itoa(row), response.TanggalAkhir)
		file.SetCellValue(sheet1Name, "Q"+strconv.Itoa(row), response.StatusApproval)
		file.SetCellValue(sheet1Name, "R"+strconv.Itoa(row), response.Status)
		row = row + 1
		baris = baris + 1
	}

	var b bytes.Buffer
	if err := file.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

func (laporan LaporanController) GetRiskEventOnTaskList(c *gin.Context) {

	result, err := laporan.service.GetRiskEventOnTaskList()
	if err != nil {
		laporan.logger.Zap.Error(err)
	}

	lib.ReturnToJson(c, 200, "200", "Inquery data berhasil", result)
}

func (laporan LaporanController) GetLaporanPerhitunganPersentasePenyelesaianDownload(c *gin.Context) {

	data := models.PerhitunganPersentasePenyelesaianDownload{}
	// laporan.logger.Zap.Info("ASU DONWLOAD")
	if err := c.Bind(&data); err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	file := excelize.NewFile()

	sheet1Name := "Sheet One"
	file.SetSheetName(file.GetSheetName(1), sheet1Name)

	result, err := laporan.service.GetLaporanPerhitunganPersentasePenyelesaianDownload(data)
	if err != nil {
		laporan.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	sFileName := ""

	if data.JenisReport == "0" {
		file.SetCellValue(sheet1Name, "A1", "No")
		file.SetCellValue(sheet1Name, "B1", "PN")
		file.SetCellValue(sheet1Name, "C1", "Nama")
		file.SetCellValue(sheet1Name, "D1", "Jenis Task")
		file.SetCellValue(sheet1Name, "E1", "Aktifitas")
		file.SetCellValue(sheet1Name, "F1", "Produk")
		file.SetCellValue(sheet1Name, "G1", "Indikator")
		file.SetCellValue(sheet1Name, "H1", "Tanggal Mulai")
		file.SetCellValue(sheet1Name, "I1", "Tanggal Selesai")
		file.SetCellValue(sheet1Name, "J1", "Jumlah Anomaly Data")
		file.SetCellValue(sheet1Name, "K1", "Jumlah Data Sudah Dilakukan Verifikasi")
		file.SetCellValue(sheet1Name, "L1", "Jumlah Verifikasi yang Perlu ditindaklanjuti")
		file.SetCellValue(sheet1Name, "M1", "Jmlh yang sudah ditindaklanjutil")
		file.SetCellValue(sheet1Name, "N1", "% Sudah Dilakukan Verifikasi")
		file.SetCellValue(sheet1Name, "O1", "% yang Sudah Ditindaklanjuti")
		sFileName = "LaporanPerhitunganPersentasePenyelesaian_PerPekerja.xlsx"
		row := 2
		baris := 1
		for _, item := range result {

			dd, ok := item.(models.LaporanPerPekerjaResult)
			if ok {
				file.SetCellValue(sheet1Name, "A"+strconv.Itoa(row), strconv.Itoa(baris))
				file.SetCellValue(sheet1Name, "B"+strconv.Itoa(row), dd.Pn)
				file.SetCellValue(sheet1Name, "C"+strconv.Itoa(row), dd.Nama)
				file.SetCellValue(sheet1Name, "D"+strconv.Itoa(row), dd.JenisTask)
				file.SetCellValue(sheet1Name, "E"+strconv.Itoa(row), dd.Aktifitas)
				file.SetCellValue(sheet1Name, "F"+strconv.Itoa(row), dd.Produk)
				file.SetCellValue(sheet1Name, "G"+strconv.Itoa(row), dd.Indikator)
				file.SetCellValue(sheet1Name, "H"+strconv.Itoa(row), dd.TanggalMulai)
				file.SetCellValue(sheet1Name, "I"+strconv.Itoa(row), dd.TanggalSelesai)
				file.SetCellValue(sheet1Name, "J"+strconv.Itoa(row), dd.JumlahDataAnomali)
				file.SetCellValue(sheet1Name, "K"+strconv.Itoa(row), dd.JumlahDataVerifikasi)
				file.SetCellValue(sheet1Name, "L"+strconv.Itoa(row), dd.JumlahDataPerluTindaklanjut)
				file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), dd.JumlahDataSudahTindaklanjut)
				file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), dd.PersenSudahVerifikasi)
				file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), "-")

			}
			row = row + 1
			baris = baris + 1
		}

	} else if data.JenisReport == "1" {
		file.SetCellValue(sheet1Name, "A1", "No")
		file.SetCellValue(sheet1Name, "B1", "Kanwil")
		file.SetCellValue(sheet1Name, "C1", "Kanca")
		file.SetCellValue(sheet1Name, "D1", "Jenis Task")
		file.SetCellValue(sheet1Name, "E1", "Aktifitas")
		file.SetCellValue(sheet1Name, "F1", "Produk")
		file.SetCellValue(sheet1Name, "G1", "Indikator")
		file.SetCellValue(sheet1Name, "H1", "Tanggal Mulai")
		file.SetCellValue(sheet1Name, "I1", "Tanggal Selesai")
		file.SetCellValue(sheet1Name, "J1", "Jumlah Anomaly Data")
		file.SetCellValue(sheet1Name, "K1", "Jumlah Data Sudah Dilakukan Verifikasi")
		file.SetCellValue(sheet1Name, "L1", "Jumlah Verifikasi yang Perlu ditindaklanjuti")
		file.SetCellValue(sheet1Name, "M1", "Jmlh yang sudah ditindaklanjutil")
		file.SetCellValue(sheet1Name, "N1", "% Sudah Dilakukan Verifikasi")
		file.SetCellValue(sheet1Name, "O1", "% yang Sudah Ditindaklanjuti")

		sFileName = "LaporanPerhitunganPersentasePenyelesaian_PerUker.xlsx"

		row := 2
		baris := 1
		for _, item := range result {
			dd, ok := item.(models.LaporanPerUkerResult)
			if ok {
				file.SetCellValue(sheet1Name, "A"+strconv.Itoa(row), strconv.Itoa(baris))
				file.SetCellValue(sheet1Name, "B"+strconv.Itoa(row), dd.Kanwil)
				file.SetCellValue(sheet1Name, "C"+strconv.Itoa(row), dd.Kanca)
				file.SetCellValue(sheet1Name, "D"+strconv.Itoa(row), dd.JenisTask)
				file.SetCellValue(sheet1Name, "E"+strconv.Itoa(row), dd.Aktifitas)
				file.SetCellValue(sheet1Name, "F"+strconv.Itoa(row), dd.Produk)
				file.SetCellValue(sheet1Name, "G"+strconv.Itoa(row), dd.Indikator)
				file.SetCellValue(sheet1Name, "H"+strconv.Itoa(row), dd.TanggalMulai)
				file.SetCellValue(sheet1Name, "I"+strconv.Itoa(row), dd.TanggalSelesai)
				file.SetCellValue(sheet1Name, "J"+strconv.Itoa(row), "-")
				file.SetCellValue(sheet1Name, "K"+strconv.Itoa(row), dd.JumlahDataVerifikasi)
				file.SetCellValue(sheet1Name, "L"+strconv.Itoa(row), dd.JumlahDataPerluTidaklanjut)
				file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), "-")
				file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), "-")
				file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), "-")

			}
			row = row + 1
			baris = baris + 1
		}
	} else if data.JenisReport == "2" {
		file.SetCellValue(sheet1Name, "A1", "No")
		file.SetCellValue(sheet1Name, "B1", "Pn")
		file.SetCellValue(sheet1Name, "C1", "Nama")
		file.SetCellValue(sheet1Name, "D1", "Kanwil")
		file.SetCellValue(sheet1Name, "E1", "Kanca")
		file.SetCellValue(sheet1Name, "F1", "Jenis Task")
		file.SetCellValue(sheet1Name, "G1", "Aktifitas")
		file.SetCellValue(sheet1Name, "H1", "Produk")
		file.SetCellValue(sheet1Name, "I1", "Indikator")
		file.SetCellValue(sheet1Name, "J1", "Tanggal Mulai")
		file.SetCellValue(sheet1Name, "K1", "Tanggal Selesai")
		file.SetCellValue(sheet1Name, "L1", "Jumlah Anomaly Data")
		file.SetCellValue(sheet1Name, "M1", "Jumlah Data Sudah Dilakukan Verifikasi")
		file.SetCellValue(sheet1Name, "N1", "Jumlah Verifikasi yang Perlu ditindaklanjuti")
		file.SetCellValue(sheet1Name, "O1", "Jmlh yang sudah ditindaklanjutil")
		file.SetCellValue(sheet1Name, "P1", "% Sudah Dilakukan Verifikasi")
		file.SetCellValue(sheet1Name, "Q1", "% yang Sudah Ditindaklanjuti")

		sFileName = "LaporanPerhitunganPersentasePenyelesaian_PerPekerjaPerUker.xlsx"

		row := 2
		baris := 1
		for _, item := range result {

			dd, ok := item.(models.LaporanPerPekerjaUkerResult)
			if ok {
				file.SetCellValue(sheet1Name, "A"+strconv.Itoa(row), strconv.Itoa(baris))
				file.SetCellValue(sheet1Name, "B"+strconv.Itoa(row), dd.Pn)
				file.SetCellValue(sheet1Name, "C"+strconv.Itoa(row), dd.Nama)
				file.SetCellValue(sheet1Name, "D"+strconv.Itoa(row), dd.Kanwil)
				file.SetCellValue(sheet1Name, "E"+strconv.Itoa(row), dd.Kanca)
				file.SetCellValue(sheet1Name, "F"+strconv.Itoa(row), dd.JenisTask)
				file.SetCellValue(sheet1Name, "G"+strconv.Itoa(row), dd.Aktifitas)
				file.SetCellValue(sheet1Name, "H"+strconv.Itoa(row), dd.Produk)
				file.SetCellValue(sheet1Name, "I"+strconv.Itoa(row), dd.Indikator)
				file.SetCellValue(sheet1Name, "J"+strconv.Itoa(row), dd.TanggalMulai)
				file.SetCellValue(sheet1Name, "K"+strconv.Itoa(row), dd.TanggalSelesai)
				file.SetCellValue(sheet1Name, "L"+strconv.Itoa(row), "-")
				file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), dd.JumlahDataVerifikasi)
				file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), dd.JumlahDataPerluTidaklanjut)
				file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), "-")
				file.SetCellValue(sheet1Name, "P"+strconv.Itoa(row), "-")
				file.SetCellValue(sheet1Name, "Q"+strconv.Itoa(row), "-")

			}
			row = row + 1
			baris = baris + 1
		}
	}

	var b bytes.Buffer
	if err := file.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	downloadName := time.Now().UTC().Format(sFileName)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

func (l LaporanController) GetLaporanPerhitunganPersentasePenyelesaianBaru(c *gin.Context) {
	data := models.PerhitunganPersentasePenyelesaianPagianted{}

	if err := c.Bind(&data); err != nil {
		l.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	if data.Kegiatan == "Verifikasi" {
		result, pagination, err := l.service.GetLaporanPerhitunganVerifikasi(data)
		if err != nil {
			lib.ReturnToJson(c, 200, "400", "Error :"+err.Error(), false)
			return
		}

		lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, pagination)
	} else if data.Kegiatan == "Briefing" {
		result, pagination, err := l.service.GetLaporanPerhitunganBriefing(data)
		if err != nil {
			lib.ReturnToJson(c, 200, "400", "Error :"+err.Error(), false)
			return
		}

		lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, pagination)
	} else if data.Kegiatan == "Coaching" {
		result, pagination, err := l.service.GetLaporanPerhitunganCoaching(data)
		if err != nil {
			lib.ReturnToJson(c, 200, "400", "Error :"+err.Error(), false)
			return
		}

		lib.ReturnToJsonWithPaginate(c, 200, "200", "Inquery data berhasil", result, pagination)
	} else {
		lib.ReturnToJson(c, 200, "400", "Kegiatan tidak ada", false)
		return
	}
}

func (l LaporanController) GetLaporanPerhitunganPersentasePenyelesaianBaruDownload(c *gin.Context) {
	data := models.PerhitunganPersentasePenyelesaianDownload{}

	if err := c.Bind(&data); err != nil {
		l.logger.Zap.Error(err)
		lib.ReturnToJson(c, 200, "400", "Input tidak sesuai : "+err.Error(), false)
		return
	}

	file := excelize.NewFile()

	sheet1Name := "Sheet One"
	file.SetSheetName(file.GetSheetName(1), sheet1Name)

	sFileName := ""
	if data.Kegiatan == "Verifikasi" {
		result, err := l.service.GetLaporanPerhitunganVerifikasiDownload(data)
		if err != nil {
			l.logger.Zap.Error(err)
			lib.ReturnToJson(c, 200, "400", "Error : "+err.Error(), false)
			return
		}

		// l.logger.Zap.Info("downloadku", result)

		file.SetCellValue(sheet1Name, "A1", "No")
		file.SetCellValue(sheet1Name, "B1", "PN")
		file.SetCellValue(sheet1Name, "C1", "Nama")
		file.SetCellValue(sheet1Name, "D1", "Kanwil")
		file.SetCellValue(sheet1Name, "E1", "Kanca")
		file.SetCellValue(sheet1Name, "F1", "Uker")
		file.SetCellValue(sheet1Name, "G1", "Kegiatan")
		file.SetCellValue(sheet1Name, "H1", "Aktifitas")
		file.SetCellValue(sheet1Name, "I1", "Produk")
		file.SetCellValue(sheet1Name, "J1", "Risk Event")
		file.SetCellValue(sheet1Name, "K1", "Indikator")
		file.SetCellValue(sheet1Name, "L1", "Tanggal Mulai")
		file.SetCellValue(sheet1Name, "M1", "Tanggal Selesai")
		file.SetCellValue(sheet1Name, "N1", "Jumlah Data Anomali")
		file.SetCellValue(sheet1Name, "O1", "Jumlah Data Sudah Dilakukan Verifikasi")
		file.SetCellValue(sheet1Name, "P1", "Jumlah Data Perlu Tindak Lanjut")
		file.SetCellValue(sheet1Name, "Q1", "Jumlah Data Yang Sudah Tindak Lanjut")
		file.SetCellValue(sheet1Name, "R1", "% Sudah Verifikasi")
		file.SetCellValue(sheet1Name, "S1", "% Sudah Tindak Lanjut")
		file.SetCellValue(sheet1Name, "T1", "Jumlah Kegiatan Dilakukan")

		sFileName = "LaporanPerhitunganVerifikasi.xlsx"
		row := 2
		baris := 1
		for _, item := range result {

			file.SetCellValue(sheet1Name, "A"+strconv.Itoa(row), strconv.Itoa(baris))
			file.SetCellValue(sheet1Name, "B"+strconv.Itoa(row), item.Pn)
			file.SetCellValue(sheet1Name, "C"+strconv.Itoa(row), item.Nama)
			file.SetCellValue(sheet1Name, "D"+strconv.Itoa(row), item.Kanwil)
			file.SetCellValue(sheet1Name, "E"+strconv.Itoa(row), item.Kanca)
			file.SetCellValue(sheet1Name, "F"+strconv.Itoa(row), item.Uker)
			file.SetCellValue(sheet1Name, "G"+strconv.Itoa(row), item.Kegiatan)
			file.SetCellValue(sheet1Name, "H"+strconv.Itoa(row), item.Aktifitas)
			file.SetCellValue(sheet1Name, "I"+strconv.Itoa(row), item.Produk)
			file.SetCellValue(sheet1Name, "J"+strconv.Itoa(row), item.RiskIssue)
			file.SetCellValue(sheet1Name, "K"+strconv.Itoa(row), item.Indikator)
			file.SetCellValue(sheet1Name, "L"+strconv.Itoa(row), item.TanggalMulai)
			file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), item.TanggalSelesai)
			file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), item.JumlahDataAnomali)
			file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), item.JumlahDataVerifikasi)
			file.SetCellValue(sheet1Name, "P"+strconv.Itoa(row), item.JumlahDataPerluTindaklanjut)
			file.SetCellValue(sheet1Name, "Q"+strconv.Itoa(row), item.JumlahDataSudahTindaklanjut)
			file.SetCellValue(sheet1Name, "R"+strconv.Itoa(row), item.PersenSudahVerifikasi)
			file.SetCellValue(sheet1Name, "S"+strconv.Itoa(row), item.PersenSudahTindaklanjut)

			// file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), item.JumlahDataVerifikasi)
			// file.SetCellValue(sheet1Name, "P"+strconv.Itoa(row), item.JumlahDataPerluTindaklanjut)
			// file.SetCellValue(sheet1Name, "Q"+strconv.Itoa(row), item.JumlahDataYangSudahTindaklanjut)
			// file.SetCellValue(sheet1Name, "R"+strconv.Itoa(row), strconv.FormatFloat(item.PersenSudahVerifikasi, 'f', 0, 64)+"%")
			// file.SetCellValue(sheet1Name, "S"+strconv.Itoa(row), strconv.FormatFloat(item.PersenSudahTindaklanjut, 'f', 0, 64)+"%")
			file.SetCellValue(sheet1Name, "T"+strconv.Itoa(row), item.JumlahKegiatanDilakukan)

			row = row + 1
			baris = baris + 1
		}

		var b bytes.Buffer
		if err := file.Write(&b); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		downloadName := time.Now().UTC().Format(sFileName)

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+downloadName)
		c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
	} else if data.Kegiatan == "Briefing" {
		result, err := l.service.GetLaporanPerhitunganBriefingDownload(data)
		if err != nil {
			l.logger.Zap.Error(err)
			lib.ReturnToJson(c, 200, "400", "Error : "+err.Error(), false)
			return
		}

		// l.logger.Zap.Info("downloadku", result)

		file.SetCellValue(sheet1Name, "A1", "No")
		file.SetCellValue(sheet1Name, "B1", "PN")
		file.SetCellValue(sheet1Name, "C1", "Nama")
		file.SetCellValue(sheet1Name, "D1", "Kanwil")
		file.SetCellValue(sheet1Name, "E1", "Kanca")
		file.SetCellValue(sheet1Name, "F1", "Uker")
		file.SetCellValue(sheet1Name, "G1", "Kegiatan")
		file.SetCellValue(sheet1Name, "H1", "Aktifitas")
		file.SetCellValue(sheet1Name, "I1", "Produk")
		file.SetCellValue(sheet1Name, "J1", "Risk Event")
		file.SetCellValue(sheet1Name, "K1", "Indikator")
		file.SetCellValue(sheet1Name, "L1", "Tanggal Mulai")
		file.SetCellValue(sheet1Name, "M1", "Tanggal Selesai")
		// file.SetCellValue(sheet1Name, "M1", "Jumlah Data Anomali")
		// file.SetCellValue(sheet1Name, "N1", "Jumlah Data Sudah Dilakukan Verifikasi")
		// file.SetCellValue(sheet1Name, "O1", "Jumlah Data Perlu Tindak Lanjut")
		// file.SetCellValue(sheet1Name, "P1", "Jumlah Data Yang Sudah Tindak Lanjut")
		// file.SetCellValue(sheet1Name, "Q1", "% Sudah Verifikasi")
		// file.SetCellValue(sheet1Name, "R1", "% Sudah Tindak Lanjut")
		file.SetCellValue(sheet1Name, "N1", "Jumlah Kegiatan Dilakukan")

		sFileName = "LaporanPerhitunganBriefing.xlsx"
		row := 2
		baris := 1
		for _, item := range result {

			file.SetCellValue(sheet1Name, "A"+strconv.Itoa(row), strconv.Itoa(baris))
			file.SetCellValue(sheet1Name, "B"+strconv.Itoa(row), item.Pn)
			file.SetCellValue(sheet1Name, "C"+strconv.Itoa(row), item.Nama)
			file.SetCellValue(sheet1Name, "D"+strconv.Itoa(row), item.Kanwil)
			file.SetCellValue(sheet1Name, "E"+strconv.Itoa(row), item.Kanca)
			file.SetCellValue(sheet1Name, "F"+strconv.Itoa(row), item.Uker)
			file.SetCellValue(sheet1Name, "G"+strconv.Itoa(row), item.Kegiatan)
			file.SetCellValue(sheet1Name, "H"+strconv.Itoa(row), item.Aktifitas)
			file.SetCellValue(sheet1Name, "I"+strconv.Itoa(row), item.Produk)
			file.SetCellValue(sheet1Name, "J"+strconv.Itoa(row), item.RiskIssue)
			file.SetCellValue(sheet1Name, "K"+strconv.Itoa(row), item.Indikator)
			file.SetCellValue(sheet1Name, "L"+strconv.Itoa(row), item.TanggalMulai)
			file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), item.TanggalSelesai)
			// file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), item.JumlahDataAnomali)
			// file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), item.JumlahDataVerifikasi)
			// file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), item.JumlahDataPerluTindaklanjut)
			// file.SetCellValue(sheet1Name, "P"+strconv.Itoa(row), item.JumlahDataYangSudahTindaklanjut)
			// file.SetCellValue(sheet1Name, "Q"+strconv.Itoa(row), item.PersenSudahVerifikasi+"%")
			// file.SetCellValue(sheet1Name, "R"+strconv.Itoa(row), item.PersenSudahTindaklanjut+"%")
			file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), item.JumlahKegiatanDilakukan)

			row = row + 1
			baris = baris + 1
		}

		var b bytes.Buffer
		if err := file.Write(&b); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		downloadName := time.Now().UTC().Format(sFileName)

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+downloadName)
		c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
	} else if data.Kegiatan == "Coaching" {
		result, err := l.service.GetLaporanPerhitunganCoachingDownload(data)
		if err != nil {
			l.logger.Zap.Error(err)
			lib.ReturnToJson(c, 200, "400", "Error : "+err.Error(), false)
			return
		}

		// l.logger.Zap.Info("downloadku", result)

		file.SetCellValue(sheet1Name, "A1", "No")
		file.SetCellValue(sheet1Name, "B1", "PN")
		file.SetCellValue(sheet1Name, "C1", "Nama")
		file.SetCellValue(sheet1Name, "D1", "Kanwil")
		file.SetCellValue(sheet1Name, "E1", "Kanca")
		file.SetCellValue(sheet1Name, "F1", "Uker")
		file.SetCellValue(sheet1Name, "G1", "Kegiatan")
		file.SetCellValue(sheet1Name, "H1", "Aktifitas")
		file.SetCellValue(sheet1Name, "I1", "Produk")
		file.SetCellValue(sheet1Name, "J1", "Risk Event")
		file.SetCellValue(sheet1Name, "K1", "Indikator")
		file.SetCellValue(sheet1Name, "L1", "Tanggal Mulai")
		file.SetCellValue(sheet1Name, "M1", "Tanggal Selesai")
		// file.SetCellValue(sheet1Name, "M1", "Jumlah Data Anomali")
		// file.SetCellValue(sheet1Name, "N1", "Jumlah Data Sudah Dilakukan Verifikasi")
		// file.SetCellValue(sheet1Name, "O1", "Jumlah Data Perlu Tindak Lanjut")
		// file.SetCellValue(sheet1Name, "P1", "Jumlah Data Yang Sudah Tindak Lanjut")
		// file.SetCellValue(sheet1Name, "Q1", "% Sudah Verifikasi")
		// file.SetCellValue(sheet1Name, "R1", "% Sudah Tindak Lanjut")
		file.SetCellValue(sheet1Name, "N1", "Jumlah Kegiatan Dilakukan")

		sFileName = "LaporanPerhitunganCoaching.xlsx"
		row := 2
		baris := 1
		for _, item := range result {

			file.SetCellValue(sheet1Name, "A"+strconv.Itoa(row), strconv.Itoa(baris))
			file.SetCellValue(sheet1Name, "B"+strconv.Itoa(row), item.Pn)
			file.SetCellValue(sheet1Name, "C"+strconv.Itoa(row), item.Nama)
			file.SetCellValue(sheet1Name, "D"+strconv.Itoa(row), item.Kanwil)
			file.SetCellValue(sheet1Name, "E"+strconv.Itoa(row), item.Kanca)
			file.SetCellValue(sheet1Name, "F"+strconv.Itoa(row), item.Uker)
			file.SetCellValue(sheet1Name, "G"+strconv.Itoa(row), item.Kegiatan)
			file.SetCellValue(sheet1Name, "H"+strconv.Itoa(row), item.Aktifitas)
			file.SetCellValue(sheet1Name, "I"+strconv.Itoa(row), item.Produk)
			file.SetCellValue(sheet1Name, "J"+strconv.Itoa(row), item.RiskIssue)
			file.SetCellValue(sheet1Name, "K"+strconv.Itoa(row), item.Indikator)
			file.SetCellValue(sheet1Name, "L"+strconv.Itoa(row), item.TanggalMulai)
			file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), item.TanggalSelesai)
			// file.SetCellValue(sheet1Name, "M"+strconv.Itoa(row), item.JumlahDataAnomali)
			// file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), item.JumlahDataVerifikasi)
			// file.SetCellValue(sheet1Name, "O"+strconv.Itoa(row), item.JumlahDataPerluTindaklanjut)
			// file.SetCellValue(sheet1Name, "P"+strconv.Itoa(row), item.JumlahDataYangSudahTindaklanjut)
			// file.SetCellValue(sheet1Name, "Q"+strconv.Itoa(row), item.PersenSudahVerifikasi+"%")
			// file.SetCellValue(sheet1Name, "R"+strconv.Itoa(row), item.PersenSudahTindaklanjut+"%")
			file.SetCellValue(sheet1Name, "N"+strconv.Itoa(row), item.JumlahKegiatanDilakukan)

			row = row + 1
			baris = baris + 1
		}

		var b bytes.Buffer
		if err := file.Write(&b); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		downloadName := time.Now().UTC().Format(sFileName)

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+downloadName)
		c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
	} else {

	}
}
