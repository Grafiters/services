package verifikasireportrealisasi

import (
	"encoding/json"
	"log"
	"riskmanagement/lib"
	models "riskmanagement/models/verifikasireportrealisasi"
	verifikasiReportRealisasi "riskmanagement/repository/verifikasireportrealisasi"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
	minio "gitlab.com/golang-package-library/minio"
)

var (
	UUID = uuid.NewString()
)

type VerifikasiReportRealisasiServiceDefinition interface {
	ReportRealisasiKreditListFilter(request models.ReportRealisasiKreditListRequest) (responses []models.ReportRealisasiKreditListResponse, totalRows int64, err error)
	ReportRealisasiKreditSummaryFilter(request models.ReportRealisasiKreditSummaryRequest) (responses []models.ReportRealisasiKreditSummaryResponse, totalRows int64, err error)
	GetAllSegmentRealisasiKredit(request models.SegmentRealisasiKreditRequest) (responses []models.SegmentRealisasiKreditResponse, totalRows int64, err error)
}

type VerifikasiReportRealisasiService struct {
	db                        lib.Database
	minio                     minio.Minio
	logger                    logger.Logger
	verifikasiReportRealisasi verifikasiReportRealisasi.VerifikasiReportRealisasiDefinition
}

func NewVerifikasiReportRealisasiService(
	db lib.Database,
	minio minio.Minio,
	logger logger.Logger,
	verifikasiReportRealisasi verifikasiReportRealisasi.VerifikasiReportRealisasiDefinition,
) VerifikasiReportRealisasiServiceDefinition {
	return VerifikasiReportRealisasiService{
		db:                        db,
		minio:                     minio,
		logger:                    logger,
		verifikasiReportRealisasi: verifikasiReportRealisasi,
	}
}

func (vrr VerifikasiReportRealisasiService) ReportRealisasiKreditListFilter(request models.ReportRealisasiKreditListRequest) (responses []models.ReportRealisasiKreditListResponse, totalRows int64, err error) {
	data, totalRows, err := vrr.verifikasiReportRealisasi.ReportRealisasiKreditListFilter(&request)

	if err != nil {
		vrr.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range data {
		var dataRealisasi interface{}
		if response.DataRealisasi != "" {
			err := json.Unmarshal([]byte(response.DataRealisasi), &dataRealisasi)
			if err != nil {
				log.Fatalf("Error parsing JSON: %s", err)
			}
		}

		var kriteriaData []string
		if response.KriteriaData != "" {
			kriteriaData = strings.Split(response.KriteriaData[1:len(response.KriteriaData)-1], ",")
		}

		responses = append(responses, models.ReportRealisasiKreditListResponse{
			VerifikasiID:     response.VerifikasiID,
			NoPelaporan:      response.NoPelaporan,
			REGION:           response.REGION,
			RGDESC:           response.RGDESC,
			MAINBR:           response.MAINBR,
			MBDESC:           response.MBDESC,
			BRANCH:           response.BRANCH,
			BRDESC:           response.BRDESC,
			ActivityId:       response.ActivityId,
			ActivityName:     response.ActivityName,
			ProductId:        response.ProductId,
			ProductName:      response.ProductName,
			PeriodeData:      response.PeriodeData,
			RestruckFlag:     response.RestruckFlag,
			ButuhPerbaikan:   response.ButuhPerbaikan,
			KriteriaData:     kriteriaData,
			HasilVerifikasi:  response.HasilVerifikasi,
			KunjunganNasabah: response.KunjunganNasabah,
			TglKunjungan:     response.TglKunjungan,
			Segment:          response.Segment,
			CreatedId:        response.CreatedId,
			CreatedDesc:      response.CreatedDesc,
			DataRealisasi:    dataRealisasi,
			StatusVerifikasi: response.StatusVerifikasi,
			LampiranName:     response.LampiranName,
			LampiranPath:     response.LampiranPath,
		})
	}

	return responses, totalRows, err
}

func (vrr VerifikasiReportRealisasiService) ReportRealisasiKreditSummaryFilter(request models.ReportRealisasiKreditSummaryRequest) (responses []models.ReportRealisasiKreditSummaryResponse, totalRows int64, err error) {
	data, totalRows, err := vrr.verifikasiReportRealisasi.ReportRealisasiKreditSummaryFilter(&request)

	if err != nil {
		vrr.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	for _, response := range data {
		var dataRealisasi interface{}
		if response.DataRealisasi != "" {
			err := json.Unmarshal([]byte(response.DataRealisasi), &dataRealisasi)
			if err != nil {
				log.Fatalf("Error parsing JSON: %s", err)
			}
		}

		var kriteriaData []string
		if response.KriteriaData != "" {
			rawData := strings.Split(response.KriteriaData[1:len(response.KriteriaData)-1], ",")
			for _, value := range rawData {
				if value != "" && value != " " {
					kriteriaData = append(kriteriaData, value)
				}
			}
		}

		responses = append(responses, models.ReportRealisasiKreditSummaryResponse{
			TotalVerifikasi:  response.TotalVerifikasi,
			ProductId:        response.ProductId,
			ProductName:      response.ProductName,
			CreatedId:        response.CreatedId,
			CreatedDesc:      response.CreatedDesc,
			REGION:           response.REGION,
			RGDESC:           response.RGDESC,
			MAINBR:           response.MAINBR,
			MBDESC:           response.MBDESC,
			BRANCH:           response.BRANCH,
			BRDESC:           response.BRDESC,
			StatusVerifikasi: response.StatusVerifikasi,
			DataRealisasi:    dataRealisasi,
			Efektif:          response.Efektif,
			NonEfektif:       response.NonEfektif,
			KriteriaData:     kriteriaData,
		})
	}

	return responses, totalRows, err
}

func (vrr VerifikasiReportRealisasiService) GetAllSegmentRealisasiKredit(request models.SegmentRealisasiKreditRequest) (responses []models.SegmentRealisasiKreditResponse, totalRows int64, err error) {
	data, totalRows, err := vrr.verifikasiReportRealisasi.GetAllSegmentRealisasiKredit(&request)

	if err != nil {
		vrr.logger.Zap.Error(err)
		return responses, totalRows, err
	}

	return data, totalRows, err
}
