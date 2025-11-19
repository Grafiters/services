package generateExcels

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"riskmanagement/lib"
	AuditTrail "riskmanagement/models/audittrail"
	briefingModel "riskmanagement/models/briefing"
	coachingModel "riskmanagement/models/coaching"
	RiskIndicator "riskmanagement/models/riskindicator"
	verifModels "riskmanagement/models/verifikasi"
	models "riskmanagement/models/verifikasireportrealisasi"
	verifRealpinModels "riskmanagement/models/verifikasireportrealisasi"
	"time"

	// fileRepo "riskmanagement/repository/files"
	// verifRepo "riskmanagement/repository/verifikasi"

	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"

	// "gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"

	"encoding/json"
	objJSON "encoding/json"
	// "encoding/base64"
	// "io/ioutil"
)

func dateFormatter(dateString string) string {
	// Parse the date string
	t, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	// Format the date
	formattedDate := t.Format("2006-01-02")
	return formattedDate
}

func SaveRequestInfo(db *lib.Database, job Job) (insertId int, err error) {
	data := &job.GenerateInfo

	result := db.DB.Create(data)
	if result.Error != nil {
		return 0, result.Error
	}

	insertedID := data.ID // Assuming the primary key field is named "ID"

	if result.RowsAffected > 0 {
		fmt.Println("insertedID", insertedID)
		// The record was successfully inserted
		// Access the inserted ID using the insertedID variable
	} else {
		// The record was not inserted
		// Handle the failure case
	}

	return insertedID, err
}

func convertJSONToExcel(job Job) (isSuccess bool, errFilePath error) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	// Create a new style with border
	borderStyle := &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Size: 11,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
			WrapText:   true,
		},
	}

	// Set the style for a cell
	style, errstyle := f.NewStyle(borderStyle)
	if errstyle != nil {
		fmt.Println("Error:", errstyle)
		fmt.Println("errExcel NewStyle", errstyle)
	}

	numberStyle := &excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
		},
		Font: &excelize.Font{
			Size: 11,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
			WrapText:   true,
		},

		NumFmt: 4,
	}

	nbrStyle, errstyle := f.NewStyle(numberStyle)
	if errstyle != nil {
		fmt.Println("Error:", errstyle)
		fmt.Println("errExcel NewStyle", errstyle)
	}

	//set column name
	fieldReport := make(map[int]string)
	// isKriteria := make(map[int]interface{})

	for colIndex, columnName := range job.ColumnNames {
		column, err := excelize.ColumnNumberToName(colIndex + 1)
		if err != nil {
			// Handle the error
			fmt.Println("Error:", err)
			fmt.Println("errExcel ColumnNumberToName", err)

			return false, err
		}

		if job.ReportId == 2 {
			colStart := "12" // for reportId 2
			cell := column + colStart
			f.SetCellValue(sheet, cell, columnName)
			f.SetCellStyle(sheet, cell, cell, style)
			f.SetColWidth(sheet, column, column, 40)
		} else if job.ReportId == 3 {
			colStart := "13"
			cell := column + colStart
			f.SetCellValue(sheet, cell, columnName)
			f.SetCellStyle(sheet, cell, cell, style)
			f.SetColWidth(sheet, column, column, 40)
		} else if job.ReportId == 4 {
			colStart := "15"
			cell := column + colStart
			f.SetCellValue(sheet, cell, columnName)
			f.SetCellStyle(sheet, cell, cell, style)
			f.SetColWidth(sheet, column, column, 40)
		} else if job.ReportId == 5 {
			colStart := "10"
			cell := column + colStart
			f.SetCellValue(sheet, cell, columnName)
			f.SetCellStyle(sheet, cell, cell, style)
			f.SetColWidth(sheet, column, column, 40)
		} else if job.ReportId == 6 {
			colStart := "1"
			cell := column + colStart
			f.SetCellValue(sheet, cell, columnName)
			// f.SetCellStyle(sheet, cell, cell, style)
			// f.SetColWidth(sheet, column, column, 40)
		} else if job.ReportId == 7 {

			colStart := "9"

			cell := column + colStart
			cell2 := column + "10"
			if columnName != "BRIEFING" && columnName != "COACHING" && columnName != "VERIFIKASI" {
				f.SetCellValue(sheet, cell, columnName)
				f.MergeCell(sheet, cell, cell2)
			} else if columnName == "BRIEFING" {
				f.SetCellValue(sheet, "H"+colStart, columnName)
				f.MergeCell(sheet, "H"+colStart, "J"+colStart)
			} else if columnName == "COACHING" {
				f.SetCellValue(sheet, "K"+colStart, columnName)
				f.MergeCell(sheet, "K"+colStart, "M"+colStart)
			} else if columnName == "VERIFIKASI" {
				f.SetCellValue(sheet, "N"+colStart, columnName)
				f.MergeCell(sheet, "N"+colStart, "P"+colStart)
			}

			//BCV Count
			f.SetCellValue(sheet, "H10", "DRAFT")
			f.SetCellValue(sheet, "I10", "SELESAI")
			f.SetCellValue(sheet, "J10", "TOTAL")
			f.SetCellValue(sheet, "K10", "DRAFT")
			f.SetCellValue(sheet, "L10", "SELESAI")
			f.SetCellValue(sheet, "M10", "TOTAL")
			f.SetCellValue(sheet, "N10", "DRAFT")
			f.SetCellValue(sheet, "O10", "SELESAI")
			f.SetCellValue(sheet, "P10", "TOTAL")

			f.SetCellStyle(sheet, "A9", "P10", style)
			f.SetColWidth(sheet, "A", "P", 40)

		} else if job.ReportId == 8 {
			colStart := "6"
			cell := column + colStart
			f.SetCellValue(sheet, cell, columnName)
			f.SetCellStyle(sheet, cell, cell, style)
			f.SetColWidth(sheet, "A", "D", 40)
		} else if job.ReportId == 9 {
			// indexKolom := colIndex
			colStart := "9"
			cell := column + colStart

			if colIndex < 21 {
				fmt.Println("Key =>", cell, columnName)
				f.SetCellValue(sheet, cell, columnName)
				f.SetCellStyle(sheet, cell, cell, style)
				f.SetColWidth(sheet, column, column, 40)
				// fieldReport[colIndex] = columnName
				switch columnName {
				case "Nomor Pelaporan":
					fieldReport[colIndex] = "no_pelaporan"
				case "AKTIVITAS":
					fieldReport[colIndex] = "activity_name"
				case "NO REKENING":
					fieldReport[colIndex] = "data_realisasi"
				case "RESTRUK":
					fieldReport[colIndex] = "restruck_flag"
				case "NAMA":
					fieldReport[colIndex] = "data_realisasi"
				case "KANWIL":
					fieldReport[colIndex] = "RGDESC"
				case "KANCA":
					fieldReport[colIndex] = "MBDESC"
				case "UKER":
					fieldReport[colIndex] = "BRDESC"
				case "CIF":
					fieldReport[colIndex] = "data_realisasi"
				case "TANGGAL REALISASI":
					fieldReport[colIndex] = "data_realisasi"
				case "TANGGAL JATUH TEMPO":
					fieldReport[colIndex] = "data_realisasi"
				case "SEGMEN":
					fieldReport[colIndex] = "segment"
				case "PRODUK":
					fieldReport[colIndex] = "product_name"
				case "PN PEMRAKARSA":
					fieldReport[colIndex] = "data_realisasi"
				case "NAMA PEMRAKARSA":
					fieldReport[colIndex] = "data_realisasi"
				case "OUTSTANDING":
					fieldReport[colIndex] = "data_realisasi"
				case "PLAFOND":
					fieldReport[colIndex] = "data_realisasi"
				case "LOAN_TYPE":
					fieldReport[colIndex] = "data_realisasi"
				case "PEMUTUS":
					fieldReport[colIndex] = "pemutus"
				case "SUDAH DILAKUKAN VERIFIKASI":
					fieldReport[colIndex] = "status_verifikasi"
				case "EFEKTIF":
					fieldReport[colIndex] = "butuh_perbaikan"
				}
			} else {

				processed := false
				for _, data := range job.JSONData {
					if processed {
						break // Stop further processing once data is loaded
					}
					// fmt.Println(data["list_kriteria"])
					if listKriteria, ok := data["list_kriteria"].([]interface{}); ok {
						for _, item := range listKriteria {
							colIndex++
							if columnName == "KRITERIA" {
								// Each item is a map with keys "id" and "kriteria"
								if kriteriaMap, ok := item.(map[string]interface{}); ok {
									if kriteria, exists := kriteriaMap["kriteria"]; exists {
										// fieldReport[colIndex-1] = kriteria.(string)
										fieldReport[colIndex-1] = "kriteria|" + strconv.FormatFloat(kriteriaMap["id_criteria"].(float64), 'f', -1, 64)
										fmt.Println("Key =>", cell, kriteria)
										f.SetCellValue(sheet, cell, kriteria)
										f.SetCellStyle(sheet, cell, cell, style)
										f.SetColWidth(sheet, column, column, 40)

										column, err = excelize.ColumnNumberToName(colIndex + 1)
										if err != nil {
											log.Fatalf("Error getting column name: %v", err)
										}
										cell = column + colStart
									}
								}
							}
						}

						processed = true
					}
				}

				// fmt.Println("Kolom KE => ", colIndex)
				column, err = excelize.ColumnNumberToName(colIndex)
				if err != nil {
					log.Fatalf("Error getting column name: %v", err)
				}
				cell = column + colStart

				if columnName != "KRITERIA" {
					colIndex -= 1
					// fieldReport[colIndex] = columnName
					fmt.Println("Key =>", cell, columnName)
					f.SetCellValue(sheet, cell, columnName)
					f.SetCellStyle(sheet, cell, cell, style)
					f.SetColWidth(sheet, column, column, 40)

					switch columnName {
					case "HASIL VERIFIKASI":
						fieldReport[colIndex] = "hasil_verifikasi"
					case "KUNJUNGAN NASABAH":
						fieldReport[colIndex] = "kunjungan_nasabah"
					case "TANGGAL KUNJUNGAN":
						fieldReport[colIndex] = "tgl_kunjungan"
					case "PN BRC/URC":
						fieldReport[colIndex] = "created_id"
					case "NAMA BRC/URC":
						fieldReport[colIndex] = "created_desc"
					}
				}
				colIndex++
			}
		} else if job.ReportId == 10 {
			colStart := "6"
			cell := column + colStart
			f.SetCellValue(sheet, cell, columnName)
			f.SetCellStyle(sheet, cell, cell, style)
			f.SetColWidth(sheet, "A", "N", 40)
		}
	}

	// set header
	isSuccess, err := SetHeaderExcel(f, sheet, job)
	fmt.Println(isSuccess)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return false, err
	}

	// set body

	// fmt.Println("action =>", isSuccess)
	// fmt.Println("JSONDATA =>", job.JSONData)
	// fmt.Println("type Data =>", reflect.TypeOf(job.JSONData))

	if isSuccess {
		// json array loop
		if job.ReportId == 1 {
			for i, json := range job.JSONData {
				row := i + 6
				indexJson := 0

				//object loop
				for key, _ := range json {
					// fmt.Println("key ", key)

					if key != "id" {
						indexJson += 1
						// strIndex := strconv.Itoa(indexJson)

						fields := map[int]string{
							1:  "no_pelaporan",
							2:  "BRANCH",
							3:  "BRDESC",
							4:  "activity_name",
							5:  "sub_activity_name",
							6:  "product_name",
							7:  "risk_issue",
							8:  "risk_indicator",
							9:  "incident_cause_code",
							10: "incident_cause_name",
							11: "sub_incident_cause_code",
							12: "sub_incident_cause_name",
							13: "verification_result",
							14: "data_source",
						}

						if fieldName, ok := fields[indexJson]; ok {

							dataToInput, _ := json[fieldName]
							column, err := excelize.ColumnNumberToName(indexJson)
							if err != nil {
								// Handle the error
								fmt.Println("Error:", err)
								return false, err
							}

							strRow := strconv.Itoa(row)
							cell := column + strRow

							f.SetCellValue(sheet, cell, dataToInput)
							f.SetCellStyle(sheet, cell, cell, style)
						}
					}
				}
			}
		} else if job.ReportId == 2 {
			for i, json := range job.JSONData {
				row := i + 13
				indexJson := 0

				// khusus reportid = 2
				for idx := 0; idx < 15; idx++ {
					indexJson += 1

					fields := map[int]string{
						1: "number",
						2: "BRANCH",
						3: "BRDESC",
						4: "MBDESC",
						5: "RGDESC",
						6: "no_pelaporan",
						// 7:  "materi", //judul materi
						// 8:  "materi", //rincian materi
						7:  "judul_materi", //judul materi
						8:  "risk_event",
						9:  "rincian_materi", //rincian materi
						10: "jumlah_peserta",
						11: "jabatan_peserta",
						12: "peserta",   //nama peserta
						13: "aktivitas", // aktifitas
						14: "status",
						15: "maker_id",
					}

					if fieldName, ok := fields[indexJson]; ok {
						dataToInput, _ := json[fieldName]
						column, err := excelize.ColumnNumberToName(indexJson)
						if err != nil {
							// Handle the error
							fmt.Println("errExcel ColumnNumberToName", err)
							fmt.Println("Error:", err)
							return false, err
						}

						strRow := strconv.Itoa(row)
						cell := column + strRow

						if indexJson == 1 {
							f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
						} else if indexJson == 7 || indexJson == 8 || indexJson == 9 || indexJson == 13 {
							if dataToInput != nil {
								dataStr := dataToInput.(string)

								// Split dataStr by "&n/" string
								parts := strings.Split(dataStr, "&n/")
								bulletData := ""
								bulletSymbol := "• "
								// Iterate over each part and set it to the corresponding cell
								for _, part := range parts {
									// Calculate the cell for each part
									bulletData += bulletSymbol + " " + part + "\n"
									f.SetCellValue(sheet, cell, bulletData)
								}
							} else {
								f.SetCellValue(sheet, cell, "-")
							}
						} else if indexJson == 12 {
							if dataToInput.(string) != "" {
								var peserta []interface{}

								// // Unmarshal the JSON string into the 'employees' slice
								err := objJSON.Unmarshal([]byte(dataToInput.(string)), &peserta)
								if err != nil {
									log.Fatalf("Error occurred during unmarshalling: %v", err)
								}
								// fmt.Println("Peserta =>", peserta)

								bulletData := ""
								bulletSymbol := "• "
								for _, emp := range peserta {
									valMap := emp.(map[string]interface{})

									sname, ok := valMap["SNAME"].(string)
									if !ok {
										log.Println("Error: SNAME is not a string")
										continue // Skip this iteration if the type assertion fails
									}

									bulletData += bulletSymbol + " " + sname + "\n"
									// fmt.Println("Peserta =>", sname)
								}

								f.SetCellValue(sheet, cell, bulletData)
							} else {
								f.SetCellValue(sheet, cell, "-")
							}

						} else {
							f.SetCellValue(sheet, cell, dataToInput)
						}

						f.SetCellStyle(sheet, cell, cell, style)
					}
				}
			}
		} else if job.ReportId == 3 {
			for i, json := range job.JSONData {
				row := i + 14
				indexJson := 0

				//object loop
				for idx := 0; idx < 17; idx++ {
					indexJson += 1
					fields := map[int]string{
						1: "number",
						2: "BRANCH",
						3: "BRDESC",
						4: "MBDESC",
						5: "RGDESC",
						6: "no_pelaporan",
						// 7:  "materi", //judul materi
						// 8:  "materi", //rincian materi
						7:  "judul_materi",
						8:  "rincian_materi",
						9:  "jumlah_peserta",
						10: "jabatan_peserta",
						11: "peserta", //nama peserta
						12: "aktifitas",
						13: "sub_aktifitas",
						// 13: "materi", //isu_risiko
						14: "isu_risiko",
						15: "risk_indicator",
						16: "status",
						17: "maker_id",
					}

					if fieldName, ok := fields[indexJson]; ok {

						dataToInput, _ := json[fieldName]

						column, err := excelize.ColumnNumberToName(indexJson)
						if err != nil {
							// Handle the error
							fmt.Println("Error:", err)
							return false, err
						}

						strRow := strconv.Itoa(row)
						cell := column + strRow

						if indexJson == 1 {
							f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
						} else if indexJson == 7 || indexJson == 8 || indexJson == 14 || indexJson == 15 {
							if dataToInput != nil {
								dataStr := dataToInput.(string)

								// Split dataStr by "&n/" string
								parts := strings.Split(dataStr, "&n/")
								bulletData := ""
								bulletSymbol := "• "
								// Iterate over each part and set it to the corresponding cell
								for _, part := range parts {
									// Calculate the cell for each part
									bulletData += bulletSymbol + " " + part + "\n"
									f.SetCellValue(sheet, cell, bulletData)
								}
							} else {
								f.SetCellValue(sheet, cell, "-")
							}
						} else if indexJson == 11 {
							if dataToInput.(string) != "" {
								var peserta []interface{}

								// Unmarshal the JSON string into the 'employees' slice
								err := objJSON.Unmarshal([]byte(dataToInput.(string)), &peserta)
								if err != nil {
									log.Fatalf("Error occurred during unmarshalling: %v", err)
								}
								// fmt.Println("Peserta =>", peserta)

								bulletData := ""
								bulletSymbol := "• "
								for _, emp := range peserta {
									valMap := emp.(map[string]interface{})

									sname, ok := valMap["SNAME"].(string)
									if !ok {
										log.Println("Error: SNAME is not a string")
										continue // Skip this iteration if the type assertion fails
									}

									bulletData += bulletSymbol + " " + sname + "\n"
									// fmt.Println("Peserta =>", sname)
								}

								f.SetCellValue(sheet, cell, bulletData)
							} else {
								f.SetCellValue(sheet, cell, "-")
							}

						} else {
							f.SetCellValue(sheet, cell, dataToInput)
						}

						f.SetCellStyle(sheet, cell, cell, style)
					}
				}
			}
		} else if job.ReportId == 4 {
			for i, json := range job.JSONData {
				row := i + 16
				indexJson := 0

				//object loop
				for idx := 0; idx < 25; idx++ {
					indexJson += 1
					fields := map[int]string{
						1:  "number",
						2:  "BRANCH",
						3:  "BRDESC",
						4:  "MBDESC",
						5:  "RGDESC",
						6:  "no_pelaporan",
						7:  "aktifitas",
						8:  "sub_aktifitas",
						9:  "informasi_lain",
						10: "status_perbaikan_konsolidasi",
						11: "maker",
						12: "risk_issue_code",
						13: "risk_issue",
						14: "risk_indicator",
						15: "risk_control",
						16: "hasil_verifikasi",
						17: "jumlah_data_yg_diverifikasi",
						18: "butuh_perbaikan",
						19: "jumlah_data_yg_harus_diperbaiki",
						20: "rtl_user",
						21: "status_perbaikan_selesai",
						22: "status_perbaikan_proses",
						23: "presentase_perbaikan",
						24: "batas_waktu_perbaikan",
						25: "indikasi_fraud",
					}

					if fieldName, ok := fields[indexJson]; ok {

						dataToInput, _ := json[fieldName]

						column, err := excelize.ColumnNumberToName(indexJson)
						if err != nil {
							// Handle the error
							fmt.Println("Error:", err)
							return false, err
						}

						strRow := strconv.Itoa(row)
						cell := column + strRow

						if indexJson == 1 {
							f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
						} else {
							f.SetCellValue(sheet, cell, dataToInput)
						}
						f.SetCellStyle(sheet, cell, cell, style)
					}
				}
			}
		} else if job.ReportId == 5 {
			for i, json := range job.JSONData {
				row := i + 11
				indexJson := 0

				for idx := 0; idx < 11; idx++ {
					indexJson += 1
					fields := map[int]string{
						1:  "number",
						2:  "tanggal",
						3:  "pn",
						4:  "nama_brc_urc",
						5:  "Kanwil",
						6:  "Kanca",
						7:  "Uker",
						8:  "no_pelaporan",
						9:  "aktifitas",
						10: "ip_address",
						11: "lokasi",
					}

					if fieldName, ok := fields[indexJson]; ok {
						dataToInput, _ := json[fieldName]

						column, err := excelize.ColumnNumberToName(indexJson)
						if err != nil {
							fmt.Println("Error: ", err)
							return false, err
						}

						strRow := strconv.Itoa(row)
						cell := column + strRow

						if indexJson == 1 {
							f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
						} else {
							f.SetCellValue(sheet, cell, dataToInput)
						}

						f.SetCellStyle(sheet, cell, cell, style)
					}
				}
			}
		} else if job.ReportId == 6 {
			for i, json := range job.JSONData {
				row := i + 2
				indexJson := 0

				for idx := 0; idx < 64; idx++ {
					indexJson += 1
					fields := map[int]string{
						1:  "id",
						2:  "key_risk_indicator",
						3:  "aktivitas",
						4:  "produk",
						5:  "jenis_indikator",
						6:  "indikasi_risiko",
						7:  "deskripsi",
						8:  "sla_verifikasi",
						9:  "sla_tl",
						10: "risk_awarness",
						11: "data_source",
						12: "map_threshold", //kck_1_min
						13: "map_threshold", //kck_2_min
						14: "map_threshold", //kck_3_min
						15: "map_threshold", //kck_4_min
						16: "map_threshold", //kck_5_min
						17: "map_threshold", //kc_1_min
						18: "map_threshold", //kc_2_min
						19: "map_threshold", //kc_3_min
						20: "map_threshold", //kc_4_min
						21: "map_threshold", //kc_5_min
						22: "map_threshold", //kcp_1_min
						23: "map_threshold", //kcp_2_min
						24: "map_threshold", //kcp_3_min
						25: "map_threshold", //kcp_4_min
						26: "map_threshold", //kcp_5_min
						27: "map_threshold", //un_1_min
						28: "map_threshold", //un_2_min
						29: "map_threshold", //un_3_min
						30: "map_threshold", //un_4_min
						31: "map_threshold", //un_5_min
						32: "map_threshold", //kk_1_min
						33: "map_threshold", //kk_2_min
						34: "map_threshold", //kk_3_min
						35: "map_threshold", //kk_4_min
						36: "map_threshold", //kk_5_min
						37: "map_threshold", //kck_1_max
						38: "map_threshold", //kck_2_max
						39: "map_threshold", //kck_3_max
						40: "map_threshold", //kck_4_max
						41: "map_threshold", //kck_5_max
						42: "map_threshold", //kc_1_max
						43: "map_threshold", //kc_2_max
						44: "map_threshold", //kc_3_max
						45: "map_threshold", //kc_4_max
						46: "map_threshold", //kc_5_max
						47: "map_threshold", //kcp_1_max
						48: "map_threshold", //kcp_2_max
						49: "map_threshold", //kcp_3_max
						50: "map_threshold", //kcp_4_max
						51: "map_threshold", //kcp_5_max
						52: "map_threshold", //un_1_max
						53: "map_threshold", //un_2_max
						54: "map_threshold", //un_3_max
						55: "map_threshold", //un_4_max
						56: "map_threshold", //un_5_max
						57: "map_threshold", //kk_1_max
						58: "map_threshold", //kk_2_max
						59: "map_threshold", //kk_3_max
						60: "map_threshold", //kk_4_max
						61: "map_threshold", //kk_5_max
						62: "parameter",
						63: "status_indikator",
						64: "is_aktif",
					}

					if fieldName, ok := fields[indexJson]; ok {
						dataToInput, _ := json[fieldName]
						column, err := excelize.ColumnNumberToName(indexJson)

						if err != nil {
							fmt.Println("Error:", err)
							return false, err
						}

						strRow := strconv.Itoa(row)
						cell := column + strRow

						if indexJson == 12 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)
								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold = valMap["kck_1_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 13 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_2_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 14 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_3_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 15 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_4_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 16 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_5_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 17 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_1_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 18 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_2_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 19 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_3_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 20 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_4_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 21 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_5_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 22 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_1_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 23 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_2_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 24 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_3_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 25 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_4_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 26 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_5_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 27 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_1_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 28 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_2_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 29 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_3_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 30 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_4_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 31 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_5_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 32 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_1_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 33 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_2_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 34 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_3_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 35 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_4_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 36 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_5_min"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 37 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_1_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 38 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_2_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 39 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_3_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 40 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_4_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 41 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kck_5_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 42 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_1_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 43 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_2_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 44 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_3_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 45 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_4_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 46 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kc_5_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 47 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_1_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 48 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_2_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 49 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_3_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 50 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_4_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 51 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kcp_5_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 52 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_1_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 53 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_2_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 54 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_3_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 55 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_4_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 56 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["un_5_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 57 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_1_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 58 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_2_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 59 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_3_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 60 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_4_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else if indexJson == 61 {
							if dataToInput != nil {
								thresholdArr := dataToInput.([]interface{})
								threshold := float64(0)

								for _, value := range thresholdArr {
									valMap := value.(map[string]interface{})
									threshold += valMap["kk_5_max"].(float64)
								}

								f.SetCellValue(sheet, cell, threshold)
							} else {
								f.SetCellValue(sheet, cell, "0")
							}
						} else {
							f.SetCellValue(sheet, cell, dataToInput)
						}

						// f.SetCellStyle(sheet, cell, cell, style)
					}
				}
			}
		} else if job.ReportId == 7 {
			for i, json := range job.JSONData {
				row := i + 11
				indexJson := 0

				for idx := 0; idx < 16; idx++ {
					indexJson += 1
					fields := map[int]string{
						1:  "number",
						2:  "pernr",
						3:  "brc",
						4:  "BRANCH",
						5:  "BRDESC",
						6:  "MBDESC",
						7:  "RGDESC",
						8:  "b_draft",
						9:  "b_finish",
						10: "b_total",
						11: "c_draft",
						12: "c_finish",
						13: "c_total",
						14: "v_draft",
						15: "v_finish",
						16: "v_total",
					}

					if fieldName, ok := fields[indexJson]; ok {
						dataToInput, _ := json[fieldName]

						column, err := excelize.ColumnNumberToName(indexJson)
						if err != nil {
							fmt.Println("Error: ", err)
							return false, err
						}

						strRow := strconv.Itoa(row)
						cell := column + strRow

						if indexJson == 1 {
							f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
						} else {
							f.SetCellValue(sheet, cell, dataToInput)
						}

						f.SetCellStyle(sheet, cell, cell, style)
					}
				}
			}
		} else if job.ReportId == 8 {
			for i, json := range job.JSONData {
				row := i + 7
				indexJson := 0
				// Check if "risk_event" field is nil for this row
				riskEventValue, riskEventExists := json["risk_event"]
				if riskEventExists && riskEventValue == "" {
					// fmt.Println("risk_indicator")

					// If "risk_event" is nil, use the alternative fields
					fields := map[int]string{
						1: "number",
						2: "risk_indicator",
						3: "module",
						4: "count",
					}

					for idx := 0; idx < 4; idx++ {
						indexJson++
						if fieldName, ok := fields[indexJson]; ok {
							dataToInput, _ := json[fieldName]

							column, err := excelize.ColumnNumberToName(indexJson)
							if err != nil {
								fmt.Println("Error: ", err)
								return
							}

							strRow := strconv.Itoa(row)
							cell := column + strRow

							if indexJson == 1 {
								f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
							} else {
								f.SetCellValue(sheet, cell, dataToInput)
							}

							f.SetCellStyle(sheet, cell, cell, style)
						}
					}
				} else {
					// fmt.Println("risk_event")
					// If "risk_event" is not nil, use the regular fields
					fields := map[int]string{
						1: "number",
						2: "risk_event",
						3: "module",
						4: "count",
					}

					for idx := 0; idx < 5; idx++ {
						indexJson++
						if fieldName, ok := fields[indexJson]; ok {
							dataToInput, _ := json[fieldName]

							column, err := excelize.ColumnNumberToName(indexJson)
							if err != nil {
								fmt.Println("Error: ", err)
								return
							}

							strRow := strconv.Itoa(row)
							cell := column + strRow

							if indexJson == 1 {
								f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
							} else {
								f.SetCellValue(sheet, cell, dataToInput)
							}

							f.SetCellStyle(sheet, cell, cell, style)
						}
					}
				}
			}
		} else if job.ReportId == 9 {
			fmt.Println("JSON CONVERTER RPT REALPIN")
			fmt.Println("FieldReport", fieldReport)

			for i, data := range job.JSONData {
				row := i + 10
				indexJson := 0
				for idx := 0; idx < len(fieldReport); idx++ {
					indexJson++

					if fieldName, ok := fieldReport[idx]; ok {
						dataToInput, _ := data[fieldName]
						column, err := excelize.ColumnNumberToName(indexJson)
						if err != nil {
							// Handle the error
							fmt.Println("errExcel ColumnNumberToName", err)
							fmt.Println("Error:", err)
							return false, err
						}

						strRow := strconv.Itoa(row)
						cell := column + strRow

						if listData, ok := dataToInput.(map[string]interface{}); ok {
							if indexJson == 3 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["NOMOR_REKENING"])
								f.SetCellValue(sheet, cell, listData["NOMOR_REKENING"])
							} else if indexJson == 5 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["NAMA_KREDITUR"])
								f.SetCellValue(sheet, cell, listData["NAMA_KREDITUR"])
							} else if indexJson == 9 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["CIFNO"])
								f.SetCellValue(sheet, cell, listData["CIFNO"])
							} else if indexJson == 10 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["periode"])
								f.SetCellValue(sheet, cell, dateFormatter(listData["periode"].(string)))
							} else if indexJson == 11 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["MATDT"])
								f.SetCellValue(sheet, cell, dateFormatter(listData["MATDT"].(string)))
							} else if indexJson == 12 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["SEGMENT"])
								f.SetCellValue(sheet, cell, listData["SEGMENT"])
							} else if indexJson == 14 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["pn_pemrakarsa"])
								f.SetCellValue(sheet, cell, listData["pn_pemrakarsa"])
							} else if indexJson == 15 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["nama_pemrakarsa"])
								f.SetCellValue(sheet, cell, listData["nama_pemrakarsa"])
							} else if indexJson == 16 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["PLAFOND"])
								f.SetCellValue(sheet, cell, listData["PLAFOND"].(float64))
							} else if indexJson == 17 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["PLAFOND"])
								f.SetCellValue(sheet, cell, listData["PLAFOND"].(float64))
							} else if indexJson == 18 {
								fmt.Println("cell =>", indexJson, cell, fieldName, listData["LOAN_TYPE"])
								f.SetCellValue(sheet, cell, listData["LOAN_TYPE"])
							} else {
								fmt.Println("")
								fmt.Println("cell =>", cell)
							}
						} else {
							if fieldName == "status_verifikasi" || fieldName == "butuh_perbaikan" || fieldName == "restruck_flag" || fieldName == "kunjungan_nasabah" {
								kondisi := ""
								if dataToInput.(float64) == 1 {
									kondisi = "Ya"
								} else {
									kondisi = "Tidak"
								}

								fmt.Println("cell =>", cell, fieldName, kondisi)
								f.SetCellValue(sheet, cell, kondisi)
							} else if fieldName == "tgl_kunjungan" {
								fmt.Println("cell =>", cell, fieldName, dataToInput)
								f.SetCellValue(sheet, cell, dateFormatter(dataToInput.(string)))
							} else if strings.Contains(fieldName, "kriteria") {

								parts := strings.Split(fieldName, "|")
								id := parts[1]
								// fmt.Println("Kriteria", id)
								found := false
								for _, arrID := range data["kriteria_data"].([]interface{}) {
									if id == arrID {
										found = true
										break
									}
								}
								var kriteria string
								if found {
									kriteria = "Ya"
								} else {
									kriteria = "Tidak"
								}

								fmt.Println("cell =>", cell, fieldName, kriteria)
								f.SetCellValue(sheet, cell, kriteria)
							} else {
								fmt.Println("cell =>", cell, fieldName, dataToInput)
								f.SetCellValue(sheet, cell, dataToInput)
							}

						}

						// f.SetCellStyle(sheet, cell, cell, style)
						f.SetCellStyle(sheet, cell, cell, nbrStyle)
					}
				}

			}

		} else if job.ReportId == 10 {
			req := job.GenerateInfo.JSONPARAMS
			var requestParams verifRealpinModels.ReportRealisasiKreditSummaryRequest
			if err := objJSON.Unmarshal([]byte(req), &requestParams); err != nil {
				fmt.Println("Error - unmarshal from json string to json struct:", err)
			}

			fields := []string{
				0: "-",
			}
			if containsString(requestParams.GroupBy, "produk") {
				fields = append(fields, "product_name")
			}

			if containsString(requestParams.GroupBy, "pn-pemrakarsa") {
				fields = append(fields, "data_realisasi.pn_pemrakarsa")
			}

			if containsString(requestParams.GroupBy, "regional-office") {
				fields = append(fields, "RGDESC")
			}

			if containsString(requestParams.GroupBy, "branch-office") {
				fields = append(fields, "MBDESC")
			}

			if containsString(requestParams.GroupBy, "unit-kerja") {
				fields = append(fields, "BRDESC")
			}

			if containsString(requestParams.GroupBy, "pn-brc-urc") {
				fields = append(fields, "created_id", "created_desc")
			}

			if containsString(requestParams.GroupBy, "efektifitas") {
				fields = append(fields, "efektif", "non_efektif")
			}

			if containsString(requestParams.GroupBy, "status-verifikasi") {
				fields = append(fields, "status_verifikasi")
			}

			if containsString(requestParams.GroupBy, "criteria") {
				fields = append(fields, "mst_kriteria")
			}

			// filter group

			startAt := 7
			for i, json := range job.JSONData {
				if json != nil {
					var jmlKriteria int
					if json["mst_kriteria"] != nil {
						jmlKriteria = len(json["mst_kriteria"].(map[string]interface{}))
					}
					for idx, field := range fields {
						var dataToInput interface{}
						if field == "mst_kriteria" {
							i_tmp := i
							for key, value := range json[field].(map[string]interface{}) {

								//Master Data Kriteria
								criteria := value
								column, err := excelize.ColumnNumberToName(idx + 1)
								if err != nil {
									fmt.Println("Error: ", err)
									return
								}
								strRow := strconv.Itoa(startAt + i_tmp)
								cell := column + strRow
								f.SetCellValue(sheet, cell, criteria)
								f.SetCellStyle(sheet, cell, cell, style)

								//Kriteria Value
								criteriaYes := countValueInArray(json["kriteria_data"].([]interface{}), key)
								column, err = excelize.ColumnNumberToName(idx + 2)
								if err != nil {
									fmt.Println("Error: ", err)
									return
								}
								cell = column + strRow
								f.SetCellValue(sheet, cell, criteriaYes)
								f.SetCellStyle(sheet, cell, cell, style)

								//Kriteria Value
								criteriaNo := json["total_verifikasi"].(float64) - float64(criteriaYes)
								column, err = excelize.ColumnNumberToName(idx + 3)
								if err != nil {
									fmt.Println("Error: ", err)
									return
								}
								cell = column + strRow
								f.SetCellValue(sheet, cell, criteriaNo)
								f.SetCellStyle(sheet, cell, cell, style)
								i_tmp++
							}
						} else {
							//jika data_realisasi, ambil object childnya
							if len(field) > 14 && field[0:15] == "data_realisasi." {
								raw := json[field[0:14]].(map[string]interface{})
								dataToInput = raw[field[15:]]
							} else {
								dataToInput = json[field]
							}

							column, err := excelize.ColumnNumberToName(idx + 1)
							if err != nil {
								fmt.Println("Error: ", err)
								return
							}

							strRow := strconv.Itoa(startAt + i)
							cell := column + strRow

							if idx == 0 {
								f.SetCellValue(sheet, cell, strconv.Itoa(i+1))
							} else {
								f.SetCellValue(sheet, cell, dataToInput)
							}
							f.SetCellStyle(sheet, cell, cell, style)

							//jika ada kriteria, merge field sesuai length nya
							if jmlKriteria > 0 {
								strTarget := strconv.Itoa(startAt + i + jmlKriteria - 1)
								cellTarget := column + strTarget
								f.MergeCell(sheet, cell, cellTarget)
								f.SetCellStyle(sheet, cell, cellTarget, style)
							}
						}
					}
					startAt = startAt + (jmlKriteria - 1)
				}
			}
		}
	}

	err = f.SaveAs(job.ExcelPath)
	if err != nil {
		fmt.Println("errExcel saveas", err)
		return false, err
	}

	// Set file permissions (example: read and write permissions for the owner, read-only for others)
	err = os.Chmod(job.ExcelPath, 0644)
	if err != nil {
		fmt.Println("errExcel Failed to set file permissions:", err)
		return false, err
	}

	return true, errFilePath
}

func containsString(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func countValueInArray(arr []interface{}, value string) int {
	count := 0
	for _, v := range arr {
		if v.(string) == value {
			count++
		}
	}
	return count
}

func uploadToMinIO(minio *minio.Minio, job Job) (minioLink string, filename string, status bool, err error) {
	var minioPath string
	bucketName := os.Getenv("BUCKET_NAME")
	subdir := "generatedExcels"

	timeNow := lib.GetTimeNow("timestime")
	filename = timeNow + "-" + job.FileName + ".xlsx"

	bucketExist := minio.BucketExist(minio.Client(), bucketName)

	uuid := uuid.New()
	minioPath = subdir + "/" + lib.GetTimeNow("year") + "/" + lib.GetTimeNow("month") + "/" + lib.GetTimeNow("day") + "/" + uuid.String() + "/" + filename

	fmt.Println("Bucket => ", bucketExist)
	if bucketExist {
		_, err := os.Open(job.ExcelPath)
		if err != nil {
			return minioPath, filename, false, err
		}

		contentType := `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`

		_, err = minio.UploadObject(minio.Client(), bucketName, minioPath, job.ExcelPath, contentType)
		if err != nil {
			return minioPath, filename, false, err
		}

	} else {
		minio.MakeBucket(minio.Client(), bucketName, "")

		_, err := os.Open(job.ExcelPath)
		if err != nil {
			return minioPath, filename, false, err

		}

		contentType := `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
		fmt.Println("contentType", contentType)

		status, err = minio.UploadObject(minio.Client(), bucketName, minioPath, job.ExcelPath, contentType)

		fmt.Println("Error => ", err)
		fmt.Println("STatus upload =>", status)

		if err != nil {
			return minioPath, filename, false, err
		}
	}

	// minioPath = GetFile(minio, minioPath, filename) // get file url
	// minioPath, _ = FileToBase64(minioPath) // turn file url to base64 download url

	return minioPath, filename, true, err
}

func saveMinIOLinkToDB(db *lib.Database, job Job, minioLink string, filename string, insertId int) (isSuccess bool, err error) {
	var UUID = uuid.NewString()

	// save link to db
	dataDownload := &DownloadUrl{
		DOWNLOADURL: UUID,
		FILEPATH:    minioLink,
		FILENAME:    filename,
	}

	resultDownload := db.DB.Create(dataDownload)
	if resultDownload.Error != nil {
		return false, resultDownload.Error
	}

	dataExport := GenerateInfo{
		ID: insertId,
	}

	resultExport := db.DB.Model(&dataExport).Updates(GenerateInfo{
		DOWNLOADURL: UUID,
		FILEDESC:    job.FileName,
		FILENAME:    filename,
	})

	if resultExport.Error != nil {
		return false, resultExport.Error
	}

	return true, err
}

func GetFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func SetHeaderExcel(f *excelize.File, sheet string, job Job) (isSuccess bool, err error) {
	// Create a new style with data body
	dataStyle := &excelize.Style{
		Font: &excelize.Font{
			Size: 12,
		},
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	}

	// Set the style for a cell
	style, errstyle := f.NewStyle(dataStyle)
	if errstyle != nil {
		fmt.Println("Error:", errstyle)
		fmt.Println("errExcel NewStyle", errstyle)
	}

	// Create a new style with data title
	titleStyle := &excelize.Style{
		Font: &excelize.Font{
			Size: 18,
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	}

	// Set the style for a cell
	style2, errstyle2 := f.NewStyle(titleStyle)
	if errstyle2 != nil {
		fmt.Println("Error:", errstyle2)
		fmt.Println("errExcel NewStyle", errstyle2)
	}

	indexJson := 0
	if job.ReportId == 2 {
		// set report title
		f.SetCellValue(sheet, "A1", "Laporan List Briefing")
		f.SetCellStyle(sheet, "A1", "A2", style2)
		// set column name
		startColumn := 3 // for column cell
		keyData := [...]string{"Nomor Pelaporan: ", "Kanwil: ", "Kanca: ", "Unit Kerja: ", "Aktivitas: ", "Judul Materi: ", "Status: ", "Periode: "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	} else if job.ReportId == 3 {
		// set report title
		f.SetCellValue(sheet, "A1", "Laporan List Coaching")
		f.SetCellStyle(sheet, "A1", "A2", style2)
		// set column name
		startColumn := 3 // for column cell
		keyData := [...]string{"Nomor Pelaporan: ", "Kanwil: ", "Kanca: ", "Unit Kerja: ", "Aktivitas: ", "Risk Event: ", "Judul Materi: ", "Status: ", "Periode: "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	} else if job.ReportId == 4 {
		// set report title
		f.SetCellValue(sheet, "A1", "Laporan List Verifikasi")
		f.SetCellStyle(sheet, "A1", "A2", style2)
		// set column name
		startColumn := 3 // for column cell
		keyData := [...]string{"Nomor Pelaporan: ", "Nama BRC/URC: ", "Kanwil: ", "Kanca: ", "Unit Kerja: ", "Risk Event: ", "Risk Indicator: ", "Indikasi Fraud: ", "Status: ", "Periode: "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	} else if job.ReportId == 5 {
		// set report title
		f.SetCellValue(sheet, "A1", "Laporan Audit Trail")
		f.SetCellStyle(sheet, "A1", "A2", style2)
		// set column name
		startColumn := 3 // for column cell
		keyData := [...]string{"Nama BRC/URC : ", "Aktivitas : ", "Kanwil : ", "Kanca : ", "Unit Kerja : ", "Periode : "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	} else if job.ReportId == 6 {
		isSuccess = true
	} else if job.ReportId == 7 {
		// set report title
		f.SetCellValue(sheet, "A1", "Laporan Rekapitulasi BCV")
		f.SetCellStyle(sheet, "A1", "A2", style2)

		startColumn := 3
		keyData := [...]string{"Kanwil: ", "Kanca: ", "Unit Kerja: ", "Nama BRC/URC: ", "Periode: "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	} else if job.ReportId == 8 {
		// set report title
		f.SetCellValue(sheet, "A1", "Laporan Rekomendasi Risk Event & Risk Indicator")
		f.MergeCell(sheet, "A1", "B1")
		f.SetCellStyle(sheet, "A1", "A2", style2)

		startColumn := 3
		keyData := [...]string{"Jenis Data: ", "Periode: "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	} else if job.ReportId == 9 {
		f.SetCellValue(sheet, "A1", "Laporan List Verifikasi Realisasi ")
		f.MergeCell(sheet, "A1", "B1")
		f.SetCellStyle(sheet, "A1", "A2", style2)

		startColumn := 3

		keyData := [...]string{"Jenis Report : ", "Kanwil : ", "Kanca : ", "Unit Kerja : "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	} else if job.ReportId == 10 {
		// set report title
		f.SetCellValue(sheet, "A1", "Laporan Realisasi Kredit")
		f.MergeCell(sheet, "A1", "B1")
		f.SetCellStyle(sheet, "A1", "A2", style2)

		startColumn := 2
		keyData := [...]string{"Jenis Report : ", "Kanwil : ", "Kanca : ", "Unit Kerja : "}

		for _, _ = range job.Headers {

			cellKey := "A" + strconv.Itoa(startColumn) // A3 A4 A5 A6 ...
			cellVal := "B" + strconv.Itoa(startColumn) // B3 B4 B5 B6 ...

			f.SetCellValue(sheet, cellKey, keyData[indexJson])
			f.SetCellValue(sheet, cellVal, job.Headers[keyData[indexJson]])

			f.SetCellStyle(sheet, cellKey, cellKey, style)
			f.SetCellStyle(sheet, cellVal, cellVal, style)

			startColumn += 1
			indexJson += 1
		}

		isSuccess = true
	}

	return isSuccess, err
}

func UpdateStatusDownload(db *lib.Database, job Job, insertedID int, rptStatus string) (isSuccess bool, err error) {
	dataExport := GenerateInfo{
		ID: insertedID,
	}

	result := db.DB.Model(&dataExport).Updates(GenerateInfo{
		RPTSTATUS: rptStatus,
	})

	if result.Error != nil {
		return false, result.Error
	}

	return true, err
}

func UpdateDataFailUpload(db *lib.Database, job Job, insertedID int, filename string, rptStatus string) (isSuccess bool, err error) {
	dataExport := GenerateInfo{
		ID: insertedID,
	}

	result := db.DB.Model(&dataExport).Updates(GenerateInfo{
		FILENAME:  filename,
		FILEDESC:  job.FileName,
		RPTSTATUS: rptStatus,
	})

	if result.Error != nil {
		return false, result.Error
	}

	return true, err
}

func UpdateRequestInfo(db *lib.Database, job Job) (status bool, err error) {
	fmt.Println("Action.go ->", job.FileName)
	dataExport := GenerateInfo{
		ID: job.ID,
	}

	result := db.DB.Model(&dataExport).Updates(GenerateInfo{
		FILEDESC: job.FileName,
	})

	if result.Error != nil {
		return false, result.Error
	}

	return true, err
}

func Generator(db *lib.Database, request Job) (JsonData []map[string]interface{}, GeneratorStatus bool, GenerateErr error) {

	if request.ReportId == 2 {
		fmt.Println("Brifing List Generator")
		req := request.GenerateInfo.JSONPARAMS

		dataReport, err := RptListBriefing(db, req)

		if err != nil {
			fmt.Println("Error Get data : ", err)
			return JsonData, false, err
		}

		var dataMapping []briefingModel.BriefingReportListFinalResponse

		for _, value := range dataReport {
			dataMapping = append(dataMapping, briefingModel.BriefingReportListFinalResponse{
				NoPelaporan:    value.NoPelaporan,
				RGDESC:         value.RGDESC,
				MBDESC:         value.MBDESC,
				BRANCH:         value.BRANCH,
				BRDESC:         value.BRDESC,
				JudulMateri:    value.JudulMateri,
				RiskEvent:      value.RiskEvent,
				RincianMateri:  value.RincianMateri,
				Aktivitas:      value.Aktivitas,
				JumlahPeserta:  value.JumlahPeserta,
				JabatanPeserta: value.JabatanPeserta,
				JenisPeserta:   value.JenisPeserta,
				Peserta:        value.Peserta,
				MakerID:        value.MakerID,
				Status:         value.Status,
			})
		}

		jsonData, errMarshal := json.Marshal(dataMapping)
		if errMarshal != nil {
			fmt.Println("Error :", errMarshal)
			return JsonData, false, errMarshal
		}

		errUnMarshal := json.Unmarshal(jsonData, &JsonData)
		if errUnMarshal != nil {
			fmt.Println("Error:", errUnMarshal)
			return JsonData, false, errUnMarshal
		}

	} else if request.ReportId == 3 {
		fmt.Println("Coaching List Generator")
		req := request.GenerateInfo.JSONPARAMS
		fmt.Println("Masuk Action Generator =>", req)

		dataReport, err := RptListCoaching(db, req)

		if err != nil {
			fmt.Println("Error Get data : ", err)
			return JsonData, false, err
		}

		var dataMapping []coachingModel.CoachingReportListFinalResponse

		for _, value := range dataReport {
			dataMapping = append(dataMapping, coachingModel.CoachingReportListFinalResponse{
				NoPelaporan:    value.NoPelaporan,
				RGDESC:         value.RGDESC,
				MBDESC:         value.MBDESC,
				BRANCH:         value.BRANCH,
				BRDESC:         value.BRDESC,
				JudulMateri:    value.JudulMateri,
				RincianMateri:  value.RincianMateri,
				JumlahPeserta:  value.JumlahPeserta,
				JabatanPeserta: value.JabatanPeserta,
				JenisPeserta:   value.JenisPeserta,
				Peserta:        value.Peserta,
				Aktifitas:      value.Aktifitas,
				SubAktifitas:   value.SubAktifitas,
				IsuRisiko:      value.IsuRisiko,
				RiskIndicator:  value.RiskIndicator,
				MakerID:        value.MakerID,
				Status:         value.Status,
			})
		}

		jsonData, errMarshal := json.Marshal(dataMapping)
		if errMarshal != nil {
			fmt.Println("Error :", errMarshal)
			return JsonData, false, errMarshal
		}

		errUnMarshal := json.Unmarshal(jsonData, &JsonData)
		if errUnMarshal != nil {
			fmt.Println("Error:", errUnMarshal)
			return JsonData, false, errUnMarshal
		}

	} else if request.ReportId == 4 {
		fmt.Println("Verifikasi List Generator")

		req := request.GenerateInfo.JSONPARAMS
		fmt.Println("Masuk Action Generator =>", req)

		dataReport, err := RptListVerifikasi(db, req)

		if err != nil {
			fmt.Println("Error Get data : ", err)
			return JsonData, false, err
		}

		var presentase_perbaikan float64
		var dataMapping []verifModels.VerifikasiReportListResponse

		for _, value := range dataReport {
			if value.ButuhPerbaikan == "Tidak" {
				presentase_perbaikan = (float64(value.StatusPerbaikanSelesai) / float64(value.JumlahDataYgDiverifikasi)) * 100
			} else {
				presentase_perbaikan = (float64(value.StatusPerbaikanSelesai) / float64(value.JumlahDataYgHarusDiperbaiki)) * 100
			}

			dataMapping = append(dataMapping, verifModels.VerifikasiReportListResponse{
				ID:                          value.ID,
				Periode:                     value.Periode,
				RGDESC:                      value.RGDESC,
				MBDESC:                      value.MBDESC,
				BRANCH:                      value.BRANCH,
				BRDESC:                      value.BRDESC,
				NoPelaporan:                 value.NoPelaporan,
				Aktifitas:                   value.Aktifitas,
				SubAktifitas:                value.SubAktifitas,
				InformasiLain:               value.InformasiLain,
				StatusPerbaikanKonsolidasi:  value.StatusPerbaikanKonsolidasi,
				Maker:                       value.Maker,
				RiskIssueCode:               value.RiskIssueCode,
				RiskIssue:                   value.RiskIssue,
				RiskIndicator:               value.RiskIndicator,
				RiskControl:                 value.RiskControl,
				HasilVerifikasi:             value.HasilVerifikasi,
				JumlahDataYgDiverifikasi:    value.JumlahDataYgDiverifikasi,
				ButuhPerbaikan:              value.ButuhPerbaikan,
				JumlahDataYgHarusDiperbaiki: value.JumlahDataYgHarusDiperbaiki,
				RTLUser:                     value.RTLUser,
				StatusPerbaikanSelesai:      value.StatusPerbaikanSelesai,
				StatusPerbaikanProses:       value.StatusPerbaikanSelesai,
				PresentasePerbaikan:         int(presentase_perbaikan),
				BatasWaktuPerbaikan:         value.BatasWaktuPerbaikan,
				IndikasiFraud:               value.IndikasiFraud,
				Filename:                    value.Filename,
				Filepath:                    value.Filepath,
			})

		}

		jsonData, errMarshal := json.Marshal(dataMapping)
		if errMarshal != nil {
			fmt.Println("Error :", errMarshal)
			return JsonData, false, errMarshal
		}

		errUnMarshal := json.Unmarshal(jsonData, &JsonData)
		if errUnMarshal != nil {
			fmt.Println("Error:", errUnMarshal)
			return JsonData, false, errUnMarshal
		}
	} else if request.ReportId == 5 {
		fmt.Println("Audit Trail Generator")
		req := request.GenerateInfo.JSONPARAMS

		dataReport, err := RptAuditTrail(db, req)

		if err != nil {
			fmt.Println("Error Get Data : ", err)
			return JsonData, false, err
		}

		var dataMapping []AuditTrail.AuditTrailResponse

		for _, value := range dataReport {
			dataMapping = append(dataMapping, AuditTrail.AuditTrailResponse{
				ID:          value.ID,
				Tanggal:     value.Tanggal,
				PN:          value.PN,
				NamaBrcUrc:  value.NamaBrcUrc,
				Kanwil:      value.Kanwil,
				Kanca:       value.Kanca,
				Uker:        value.Uker,
				NoPelaporan: value.NoPelaporan,
				Aktifitas:   value.Aktifitas,
				IpAddress:   value.IpAddress,
				Lokasi:      value.Lokasi,
			})
		}

		jsonData, errMarshal := json.Marshal(dataMapping)
		if errMarshal != nil {
			fmt.Println("Error :", errMarshal)
			return JsonData, false, errMarshal
		}

		errUnMarshal := json.Unmarshal(jsonData, &JsonData)
		if errUnMarshal != nil {
			fmt.Println("Error:", errUnMarshal)
			return JsonData, false, errUnMarshal
		}
	} else if request.ReportId == 6 {
		fmt.Println("Setting Threshold")
		req := request.GenerateInfo.JSONPARAMS

		dataReport, err := GetDataThreshold(db, req)

		if err != nil {
			fmt.Println("Err Get Data : ", err)
			return JsonData, false, err
		}

		var dataMapping []RiskIndicator.ThresholdIndicatorResponse

		for _, value := range dataReport {
			dataThrehold, err := GetThreshold(db, value.Index)

			if err != nil {
				fmt.Println("Err Get Data Threshold : ", err)
				return JsonData, false, err
			}

			dataMapping = append(dataMapping, RiskIndicator.ThresholdIndicatorResponse{
				Index:            value.Index,
				Id:               value.Id,
				KeyRiskIndicator: value.KeyRiskIndicator,
				Aktivitas:        value.Aktivitas,
				Produk:           value.Produk,
				JenisIndikator:   value.JenisIndikator,
				IndikasiRisiko:   value.IndikasiRisiko,
				Deskripsi:        value.Deskripsi,
				SlaVerifikasi:    value.SlaVerifikasi,
				SlaTl:            value.SlaTl,
				RiskAwarness:     value.RiskAwarness,
				DataSource:       value.DataSource,
				Parameter:        value.Parameter,
				StatusIndikator:  value.StatusIndikator,
				IsAktif:          value.IsAktif,
				MapThreshold:     dataThrehold,
			})
		}

		jsonData, errMarshal := json.Marshal(dataMapping)
		if errMarshal != nil {
			fmt.Println("Error :", errMarshal)
			return JsonData, false, errMarshal
		}

		errUnMarshal := json.Unmarshal(jsonData, &JsonData)
		if errUnMarshal != nil {
			fmt.Println("Error:", errUnMarshal)
			return JsonData, false, errUnMarshal
		}

	} else if request.ReportId == 7 {
		fmt.Println("Rekapitulasi BCV Generator")

		req := request.GenerateInfo.JSONPARAMS

		dataReport, err := RptRekapitulasiBCV(db, req)

		if err != nil {
			fmt.Println("Error Get Data : ", err)
			return JsonData, false, err
		}

		var dataMapping []verifModels.RptRekapitulasiBCVResponse

		for _, value := range dataReport {
			dataMapping = append(dataMapping, verifModels.RptRekapitulasiBCVResponse{
				Pernr:   value.Pernr,
				BRC:     value.BRC,
				BRANCH:  value.BRANCH,
				BRDESC:  value.BRDESC,
				MBDESC:  value.MBDESC,
				RGDESC:  value.RGDESC,
				BDraft:  value.BDraft,
				BFinish: value.BFinish,
				BTotal:  value.BTotal,
				CDraft:  value.CDraft,
				CFinish: value.CFinish,
				CTotal:  value.CTotal,
				VDraft:  value.VDraft,
				VFinish: value.VFinish,
				VTotal:  value.VTotal,
			})
		}

		jsonData, errMarshal := json.Marshal(dataMapping)
		if errMarshal != nil {
			fmt.Println("Error :", errMarshal)
			return JsonData, false, errMarshal
		}

		errUnMarshal := json.Unmarshal(jsonData, &JsonData)
		if errUnMarshal != nil {
			fmt.Println("Error:", errUnMarshal)
			return JsonData, false, errUnMarshal
		}
	} else if request.ReportId == 8 {
		fmt.Println("Rekomdaasi Risk Generator")
		req := request.GenerateInfo.JSONPARAMS

		dataReport, err := RptRekomendasiRiskRequest(db, req)

		if err != nil {
			fmt.Println("Error Get Data : ", err)
			return JsonData, false, err
		}

		var dataMapping []verifModels.RptRekomendasiRiskResponse

		for _, value := range dataReport {
			dataMapping = append(dataMapping, verifModels.RptRekomendasiRiskResponse{
				RiskEvent:     value.RiskEvent,
				RiskIndicator: value.RiskIndicator,
				Module:        value.Module,
				Count:         value.Count,
			})
		}

		jsonData, errMarshal := json.Marshal(dataMapping)
		if errMarshal != nil {
			fmt.Println("Error :", errMarshal)
			return JsonData, false, errMarshal
		}

		errUnMarshal := json.Unmarshal(jsonData, &JsonData)
		if errUnMarshal != nil {
			fmt.Println("Error:", errUnMarshal)
			return JsonData, false, errUnMarshal
		}

	} else if request.ReportId == 9 {
		fmt.Println("Rpt Verifikasi Realisasi List Generator")
		req := request.GenerateInfo.JSONPARAMS

		dataReport, err := RptRealisasiKreditList(db, req)

		if err != nil {
			fmt.Println("Error Get data : ", err)
			return JsonData, false, err
		}

		var dataMapping []verifRealpinModels.GeneratorRealPinRptList

		listKriteria, err := GetKriteriaByPeriodeList(db, req)

		if err != nil {
			log.Fatalf("Error parsing JSON : %s", err)
			return JsonData, false, err
		}

		for _, value := range dataReport {
			var dataRealisasi interface{}
			if value.DataRealisasi != "" {
				err := json.Unmarshal([]byte(value.DataRealisasi), &dataRealisasi)
				if err != nil {
					log.Fatalf("Error parsing JSON : %s", err)
					return JsonData, false, err
				}
			}

			var kriteriaData []string
			if value.KriteriaData != "" {
				kriteriaData = strings.Split(value.KriteriaData[1:len(value.KriteriaData)-1], ",")

				fmt.Println("kriteria data =>", kriteriaData)
			}

			dataMapping = append(dataMapping, verifRealpinModels.GeneratorRealPinRptList{
				NoPelaporan:      value.NoPelaporan,
				REGION:           value.REGION,
				RGDESC:           value.RGDESC,
				MAINBR:           value.MAINBR,
				MBDESC:           value.MBDESC,
				BRANCH:           value.BRANCH,
				BRDESC:           value.BRDESC,
				ActivityId:       value.ActivityId,
				ActivityName:     value.ActivityName,
				ProductId:        value.ProductId,
				ProductName:      value.ProductName,
				PeriodeData:      value.PeriodeData,
				RestruckFlag:     value.RestruckFlag,
				ButuhPerbaikan:   value.ButuhPerbaikan,
				KriteriaData:     kriteriaData,
				Segment:          value.Segment,
				CreatedId:        value.CreatedId,
				CreatedDesc:      value.CreatedDesc,
				DataRealisasi:    dataRealisasi,
				StatusVerifikasi: value.StatusVerifikasi,
				HasilVerifikasi:  value.HasilVerifikasi,
				KunjunganNasabah: value.KunjunganNasabah,
				TglKunjungan:     value.TglKunjungan,
				ListKriteria:     listKriteria,
			})

			jsonData, errMarshal := json.Marshal(dataMapping)

			// fmt.Println("JSONDATA =>", jsonData)

			if errMarshal != nil {
				fmt.Println("Error :", errMarshal)
				return JsonData, false, errMarshal
			}

			errUnMarshal := json.Unmarshal(jsonData, &JsonData)
			if errUnMarshal != nil {
				fmt.Println("Error:", errUnMarshal)
				return JsonData, false, errUnMarshal
			}
		}
	} else if request.ReportId == 10 {
		fmt.Println("Rpt Verifikasi Realisasi Summary Generator")
		req := request.GenerateInfo.JSONPARAMS

		dataReport, err := RptRealisasiKreditSummary(db, req)

		if err != nil {
			fmt.Println("Error Get data : ", err)
			return JsonData, false, err
		}

		dataMstKriteria, err := GetKriteriaByPeriodeSummary(db, req)
		if err != nil {
			fmt.Println("Error Get data : ", err)
			return JsonData, false, err
		}

		mstKriteria := make(map[int64]string)
		for _, value := range dataMstKriteria {
			mstKriteria[value.IdCriteria] = value.Kriteria
		}

		var dataMapping []verifRealpinModels.ReportRealisasiKreditSummaryDownloadResponse

		for _, value := range dataReport {
			var dataRealisasi interface{}
			if value.DataRealisasi != "" {
				err := json.Unmarshal([]byte(value.DataRealisasi), &dataRealisasi)
				if err != nil {
					log.Fatalf("Error parsing JSON: %s", err)
				}
			}

			var kriteriaData []string
			if value.KriteriaData != "" {
				rawData := strings.Split(value.KriteriaData[1:len(value.KriteriaData)-1], ",")
				for _, value := range rawData {
					if value != "" && value != " " {
						kriteriaData = append(kriteriaData, value)
					}
				}
			}

			dataMapping = append(dataMapping, models.ReportRealisasiKreditSummaryDownloadResponse{
				TotalVerifikasi:  value.TotalVerifikasi,
				ProductId:        value.ProductId,
				ProductName:      value.ProductName,
				CreatedId:        value.CreatedId,
				CreatedDesc:      value.CreatedDesc,
				REGION:           value.REGION,
				RGDESC:           value.RGDESC,
				MAINBR:           value.MAINBR,
				MBDESC:           value.MBDESC,
				BRANCH:           value.BRANCH,
				BRDESC:           value.BRDESC,
				StatusVerifikasi: value.StatusVerifikasi,
				DataRealisasi:    dataRealisasi,
				Efektif:          value.Efektif,
				NonEfektif:       value.NonEfektif,
				KriteriaData:     kriteriaData,
				MstKriteria:      mstKriteria,
			})

			jsonData, errMarshal := json.Marshal(dataMapping)

			if errMarshal != nil {
				fmt.Println("Error :", errMarshal)
				return JsonData, false, errMarshal
			}

			errUnMarshal := json.Unmarshal(jsonData, &JsonData)
			if errUnMarshal != nil {
				fmt.Println("Error:", errUnMarshal)
				return JsonData, false, errUnMarshal
			}
		}
	} else {
		return JsonData, false, fmt.Errorf("generator not available")
	}

	return JsonData, true, GenerateErr
}

func DeleteTempFile(job Job) (status bool, err error) {
	// Open the file to ensure it's not being used by another process
	// file, err := os.Open(job.ExcelPath)
	file, err := os.OpenFile(job.ExcelPath, os.O_RDWR, 0666)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Close the file before deleting
	file.Close()

	err = os.Remove(job.ExcelPath)

	if err != nil {
		return false, err
	}

	return true, err
}
