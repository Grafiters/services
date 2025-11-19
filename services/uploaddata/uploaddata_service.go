package uploaddata

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/uploaddata"
	repository "riskmanagement/repository/uploaddata"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
)

var (
	UUID = uuid.NewString()
)

type UploadDataDefinition interface {
	UploadRiskControl(request models.UploadControlRequest) (status bool, message string, err error)
	UploadRisknIndicator(request models.UploadIndicatorRequest) (status bool, message string, err error)
	UploadRiskEvent(request models.UploadRiskIssueRequest) (status bool, message string, err error)
}

type UploadDataService struct {
	db     lib.Database
	logger logger.Logger
	repo   repository.UploadDataDefinition
}

func NewUploadDataService(
	db lib.Database,
	logger logger.Logger,
	repo repository.UploadDataDefinition,
) UploadDataDefinition {
	return UploadDataService{
		db:     db,
		logger: logger,
		repo:   repo,
	}
}

// UploadRiskControl implements UploadDataDefinition.
func (u UploadDataService) UploadRiskControl(request models.UploadControlRequest) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")
	tx := u.db.DB.Begin()
	if request.JenisData != "Risk Control" {
		tx.Rollback()
		message = "Format template tidak sesuai !"
		return false, message, err
	} else {
		if len(request.ExcelData) != 0 {
			numbered, _ := u.repo.GetKodeRiskControl()

			for i, value := range request.ExcelData {
				ada, err := u.repo.CekRiskControl(value.RiskControl)

				if ada > 0 {
					tx.Rollback()
					message = "Data duplicate pada baris ke " + strconv.Itoa(i+1)
					return false, message, err
				} else {
					// fmt.Println("numbered =>", "C"+strconv.Itoa(numbered))
					var kode string
					kode = "C" + strconv.Itoa(numbered)

					var status bool

					if strings.ToLower(value.Status) == "aktif" {
						status = true
					} else {
						status = false
					}

					_, err = u.repo.UploadRiskControl(models.RiskControlRequest{
						Kode:        kode,
						RiskControl: value.RiskControl,
						ControlType: value.ControlType,
						Nature:      value.Nature,
						KeyControl:  value.KeyControl,
						Deskripsi:   value.Deskripsi,
						Status:      status,
						CreatedAt:   &timeNow,
					}, tx)

					numbered++
					if err != nil {
						tx.Rollback()
						message = "Gagal menyimpan data"
						return false, message, err
					}
				}
			}
		} else {
			tx.Rollback()
			message = "Data Kosong !"
			return false, message, err
		}
	}

	tx.Commit()
	message = "Data Berhasil di upload"
	return true, message, err
}

// UploadRisknIndicator implements UploadDataDefinition.
func (u UploadDataService) UploadRisknIndicator(request models.UploadIndicatorRequest) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := u.db.DB.Begin()

	if request.JenisData != "Risk Indicator" {
		tx.Rollback()
		message = "Format Template tidak sesuai !"
		return false, message, err
	} else {
		if len(request.ExcelData) != 0 {
			for i, value := range request.ExcelData {
				ada, _ := u.repo.CekRiskIndicator(value.RiskIndicatorCode, value.RiskIndicator)

				if ada > 0 {
					tx.Rollback()
					message = "Data duplicate pada baris ke " + strconv.Itoa(i+1)
					return false, message, err
				} else {
					var status bool

					if strings.ToLower(value.Status) == "aktif" {
						status = true
					} else {
						status = false
					}

					_, err = u.repo.UploadRisknIndicator(models.RiskIndicatorRequest{
						RiskIndicatorCode: value.RiskIndicatorCode,
						RiskIndicator:     value.RiskIndicator,
						ActivityID:        value.ActivityID,
						ProductID:         value.ProductID,
						Deskripsi:         value.Deskripsi,
						Satuan:            value.Satuan,
						Sifat:             value.Sifat,
						SLAVerifikasi:     value.SLAVerifikasi,
						SLATindakLanjut:   value.SLATindakLanjut,
						SumberData:        value.SumberData,
						SumberDataText:    value.SumberDataText,
						PeriodePemantauan: value.PeriodePemantauan,
						Owner:             value.Owner,
						KPI:               value.KPI,
						StatusIndikator:   value.StatusIndikator,
						DataSourceAnomaly: value.DataSourceAnomaly,
						Status:            status,
						CreatedAt:         &timeNow,
					}, tx)

					if err != nil {
						tx.Rollback()
						message = "Gagal menyimpan data"
						return false, message, err
					}
				}
			}
		} else {
			tx.Rollback()
			message = "Data Kosong !"
			return false, message, err
		}
	}

	tx.Commit()
	message = "Data Berhasil di upload"
	return true, message, err
}

func (u UploadDataService) UploadRiskEvent(request models.UploadRiskIssueRequest) (status bool, message string, err error) {
	timeNow := lib.GetTimeNow("timestime")
	today := lib.GetTimeNow("date2")

	tx := u.db.DB.Begin()

	if request.JenisData != "Risk Event" {
		tx.Rollback()
		message = "Format Template tidak sesuai !"
		return false, message, err
	} else {
		if len(request.ExcelData) != 0 {
			kode := "RE." + today

			numbered, _ := u.repo.GetCounterEvent(kode)
			for i, value := range request.ExcelData {
				RiskIssueCode := "RE." + today + "." + ConcatNumber(numbered+1)

				ada, _ := u.repo.CekRiskEvent(value.RiskIssue)

				if ada > 0 {
					tx.Rollback()
					message = "Data duplicate pada baris ke " + strconv.Itoa(i+1)
					return false, message, err
				} else {
					fmt.Println("Risk Issue Kode =>", RiskIssueCode)
					var status bool
					if strings.ToLower(value.Status) == "aktif" {
						status = true
					} else {
						status = false
					}

					requestEvent := &models.RiskEvent{
						RiskTypeID:     value.RiskTypeID,
						RiskIssueCode:  RiskIssueCode,
						RiskIssue:      value.RiskIssue,
						Deskripsi:      value.Deskripsi,
						KategoriRisiko: value.KategoriRisiko,
						Status:         status,
						Likelihood:     value.Likelihood,
						Impact:         value.Impact,
						DeleteFlag:     false,
						CreatedAt:      &timeNow,
					}

					riskEvent, err := u.repo.UploadRiskIssue(requestEvent, tx)

					fmt.Println("Output =>", riskEvent)

					if err != nil {
						tx.Rollback()
						message = "Gagal menyimpan data"
						return false, message, err
					}

					if len(value.MapProses) == 0 {
						tx.Rollback()
						message = "Data Proses Kosong !"
						return false, message, err
					}

					for _, proses := range value.MapProses {
						validProses1, _ := u.repo.ValidasiMegaProses(proses.IDMegaProses)

						if !validProses1 {
							tx.Rollback()
							message = "Penulisan Mega Proses salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						validProses2, _ := u.repo.ValidasiMajorProses(proses.IDMajorProses)

						if !validProses2 {
							tx.Rollback()
							message = "Penulisan Major Proses salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						validProses3, _ := u.repo.ValidasiSubMajorProses(proses.IdSubMajorProses)

						if !validProses3 {
							tx.Rollback()
							message = "Penulisan Sub Major Proses salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						// valid
						_, err = u.repo.MapProses(&models.MapProsesRequest{
							IDRiskIssue:    riskEvent.ID,
							MegaProses:     proses.IDMegaProses,
							MajorProses:    proses.IDMajorProses,
							SubMajorProses: proses.IdSubMajorProses,
						}, tx)

						if err != nil {
							tx.Rollback()
							message = "Gagal menyimpan data proses"
							return false, message, err
						}
					}

					if len(value.MapEvent) == 0 {
						tx.Rollback()
						message = "Data event type kosong !"
						return false, message, err
					}

					for _, event := range value.MapEvent {
						validEvent1, _ := u.repo.ValidasiEventLv1(event.IDEventTypeLv1)

						if !validEvent1 {
							tx.Rollback()
							message = "Penulisan Event Type 1 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						validEvent2, _ := u.repo.ValidasiEventLv2(event.IDEventTypeLv2)

						if !validEvent2 {
							tx.Rollback()
							message = "Penulisan Event Type 2 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						validEvent3, _ := u.repo.ValidasiEventLv3(event.IDEventTypeLv3)

						if !validEvent3 {
							tx.Rollback()
							message = "Penulisan Event Type 3 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						_, err = u.repo.MapEvent(&models.MapEventRequest{
							IDRiskIssue:  riskEvent.ID,
							EventTypeLv1: event.IDEventTypeLv1,
							EventTypeLv2: event.IDEventTypeLv2,
							EventTypeLv3: event.IDEventTypeLv3,
						}, tx)

						if err != nil {
							tx.Rollback()
							message = "Gagal menyimpan data event type"
							return false, message, err
						}
					}

					if len(value.MapKejadian) == 0 {
						tx.Rollback()
						message = "Data penyebab kejadian kosong !"
						return false, message, err
					}

					for _, kejadian := range value.MapKejadian {
						ValidasiKejadianLv1, _ := u.repo.ValidasiKejadianLv1(kejadian.IDPenyebabKejadianLv1)

						if !ValidasiKejadianLv1 {
							tx.Rollback()
							message = "Penulisan Penyebab Kejadian 1 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						ValidasiKejadianLv2, _ := u.repo.ValidasiKejadianLv2(kejadian.IDPenyebabKejadianLv2)

						if !ValidasiKejadianLv2 {
							tx.Rollback()
							message = "Penulisan Penyebab Kejadian 2 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						ValidasiKejadianLv3, _ := u.repo.ValidasiKejadianLv3(kejadian.IDPenyebabKejadianLv3)

						if !ValidasiKejadianLv3 {
							tx.Rollback()
							message = "Penulisan Penyebab Kejadian 3 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						_, err = u.repo.MapKejadian(&models.MapKejadianRequest{
							IDRiskIssue:         riskEvent.ID,
							PenyebabKejadianLv1: kejadian.IDPenyebabKejadianLv1,
							PenyebabKejadianLv2: kejadian.IDPenyebabKejadianLv2,
							PenyebabKejadianLv3: kejadian.IDPenyebabKejadianLv3,
						}, tx)

						if err != nil {
							tx.Rollback()
							message = "Gagal menyimpan data penyebab kejadian"
							return false, message, err
						}
					}

					if len(value.MapProduct) == 0 {
						tx.Rollback()
						message = "Data produk kosong !"
						return false, message, err
					}

					for _, produk := range value.MapProduct {
						_, err = u.repo.MapProduct(&models.MapProductRequest{
							IDRiskIssue: riskEvent.ID,
							Product:     produk.ProductId,
						}, tx)

						if err != nil {
							tx.Rollback()
							message = "Gagal menyimpan data produk"
							return false, message, err
						}
					}

					if len(value.MapLiniBisnis) == 0 {
						tx.Rollback()
						message = "Data lini bisnis kosong !"
						return false, message, err
					}

					for _, liniBisnis := range value.MapLiniBisnis {
						validLB1, _ := u.repo.ValidasiLiniBisnisLv1(liniBisnis.IdLiniBisnisLv1)

						if !validLB1 {
							tx.Rollback()
							message = "Penulisan Lini Bisnis 1 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						validLB2, _ := u.repo.ValidasiLiniBisnisLv2(liniBisnis.IdLiniBisnisLv2)

						if !validLB2 {
							tx.Rollback()
							message = "Penulisan Lini Bisnis 2 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						validLB3, _ := u.repo.ValidasiLiniBisnisLv3(liniBisnis.IdLiniBisnisLv3)

						if !validLB3 {
							tx.Rollback()
							message = "Penulisan Lini Bisnis 3 salah pada baris ke " + strconv.Itoa(i+1)
							return false, message, err
						}

						_, err = u.repo.MapLiniBisnis(&models.MapLiniBisnisRequest{
							IDRiskIssue:   riskEvent.ID,
							LiniBisnisLv1: liniBisnis.IdLiniBisnisLv1,
							LiniBisnisLv2: liniBisnis.IdLiniBisnisLv2,
							LiniBisnisLv3: liniBisnis.LiniBisnisLv3,
						}, tx)

						if err != nil {
							tx.Rollback()
							message = "Gagal menyimpan data lini bisnis"
							return false, message, err
						}
					}

					if len(value.MapAktifitas) == 0 {
						tx.Rollback()
						message = "Data aktivitas kosong !"
						return false, message, err
					}

					// fmt.Println("Map Aktifitas =>", value.MapAktifitas)

					for _, aktifitas := range value.MapAktifitas {
						_, err = u.repo.MapAktifitas(&models.MapAktifitasRequest{
							IDRiskIssue:  riskEvent.ID,
							Aktifitas:    aktifitas.ActivityID,
							SubAktifitas: aktifitas.SubActivityID,
						}, tx)

						if err != nil {
							tx.Rollback()
							message = "Gagal menyimpan data aktivitas"
							return false, message, err
						}
					}

					numbered++
				}
			}
		} else {
			tx.Rollback()
			message = "Data Kosong !"
			return false, message, err
		}
	}

	tx.Commit()
	message = "Data Berhasil di upload"
	return true, message, err
}

func ConcatNumber(number int) string {
	numStr := strconv.Itoa(number)
	// Prepend zeros and then take the rightmost 4 characters
	formattedStr := "0000" + numStr
	formattedStr = formattedStr[len(formattedStr)-4:]

	return formattedStr
}
