package realisasiservice

import (
	lib "riskmanagement/lib"
	models "riskmanagement/models/RealisasiModels"
	jwt "riskmanagement/services/auth"

	"gitlab.com/golang-package-library/logger"
)

type RealisasiDefinition interface {
	GetDataParameter(request *models.ParameterGetHeaderRequest) (response models.ResponseData, err error)
	StoreDataParameter(request *models.ParameterStoreHeaderRequest) (response models.StoreResponse, err error)
	GetDataRevisiUker(request *models.RevisiUkerGetHeaderRequest) (response models.ResponseData, err error)
	StoreDataRevisiUker(request *models.RevisiUkerStoreHeaderRequest) (response models.StoreResponse, err error)
	DeleteDataRevisiUker(request *models.RevisiUkerDeleteHeaderRequest) (response models.StoreResponse, err error)
	GetDataRealisasi(request *models.RealisasiGetHeaderRequest) (response models.ResponseSampleData, err error)
	UpdateFlagVerifikasi(request *models.RealisasiUpdateFlagRequest) (response models.UpdateFlagResponse, err error)
}

type RealisasiService struct {
	logger     logger.Logger
	jwtService jwt.JWTAuthService
}

func NewRealisasiService(
	logger logger.Logger,
	jwtService jwt.JWTAuthService,
) RealisasiDefinition {
	return RealisasiService{
		logger:     logger,
		jwtService: jwtService,
	}
}

// GetDataParameter implements RealisasiDefinition.
func (r RealisasiService) GetDataParameter(request *models.ParameterGetHeaderRequest) (response models.ResponseData, err error) {
	baseUrl, err := lib.GetVarEnv("RealisasiUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "/parameter/getData"

	token := r.jwtService.CreateRealisasiToken(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	requestBody := models.ParameterBodyFilterRequest{
		PeriodeAwal:  request.Request.PeriodeAwal,
		PeriodeAkhir: request.Request.PeriodeAkhir,
		Limit:        request.Request.Limit,
		Offset:       request.Request.Offset,
		Pernr:        request.Pernr,
	}

	err = lib.MakeRequest("POST", Url, headers, requestBody, &response)

	if err != nil {
		r.logger.Zap.Error(err)
	}

	return response, err
}

// StoreDataParameter implements RealisasiDefinition.
func (r RealisasiService) StoreDataParameter(request *models.ParameterStoreHeaderRequest) (response models.StoreResponse, err error) {
	baseUrl, err := lib.GetVarEnv("RealisasiUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "/parameter/store"

	token := r.jwtService.CreateRealisasiToken(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	requestBody := models.ParameterBodyRequest{
		NilaiNpl:          request.Request.NilaiNpl,
		NilaiDpk:          request.Request.NilaiDpk,
		Rugi:              request.Request.Rugi,
		PeriodeKeragaan:   request.Request.PeriodeKeragaan,
		LastUpdatePernr:   request.Request.LastUpdatePernr,
		LastUpdateSname:   request.Request.LastUpdateSname,
		LastUpdateStelltx: request.Request.LastUpdateStelltx,
		Pernr:             request.Pernr,
	}

	err = lib.MakeRequest("POST", Url, headers, requestBody, &response)

	if err != nil {
		r.logger.Zap.Error(err)
	}

	return response, err
}

// GetDataRevisiUker implements RealisasiDefinition.
func (r RealisasiService) GetDataRevisiUker(request *models.RevisiUkerGetHeaderRequest) (response models.ResponseData, err error) {
	baseUrl, err := lib.GetVarEnv("RealisasiUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "/revisiuker/getData"

	token := r.jwtService.CreateRealisasiToken(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	requestBody := models.RevisUkerBodyRequest{
		REGION:      request.Request.REGION,
		MAINBR:      request.Request.MAINBR,
		BRANCH:      request.Request.BRANCH,
		JenisRevisi: request.Request.JenisRevisi,
		Limit:       request.Request.Limit,
		Offset:      request.Request.Offset,
		Pernr:       request.Pernr,
	}

	err = lib.MakeRequest("POST", Url, headers, requestBody, &response)

	if err != nil {
		r.logger.Zap.Error(err)
	}

	return response, err
}

// DeleteDataRevisiUker implements RealisasiDefinition.
func (r RealisasiService) DeleteDataRevisiUker(request *models.RevisiUkerDeleteHeaderRequest) (response models.StoreResponse, err error) {
	baseUrl, err := lib.GetVarEnv("RealisasiUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "/revisiuker/delete"

	token := r.jwtService.CreateRealisasiToken(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	requestBody := models.RevisiUkerDeleteRequest{
		Id:    request.Request.Id,
		Pernr: request.Pernr,
	}

	err = lib.MakeRequest("POST", Url, headers, requestBody, &response)

	if err != nil {
		r.logger.Zap.Error(err)
	}

	return response, err
}

// StoreDataRevisiUker implements RealisasiDefinition.
func (r RealisasiService) StoreDataRevisiUker(request *models.RevisiUkerStoreHeaderRequest) (response models.StoreResponse, err error) {
	baseUrl, err := lib.GetVarEnv("RealisasiUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "/revisiuker/store"

	token := r.jwtService.CreateRealisasiToken(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	requestBody := models.RevisiUkerStoreRequest{
		REGION:      request.Request.REGION,
		RGDESC:      request.Request.RGDESC,
		MAINBR:      request.Request.MAINBR,
		MBDESC:      request.Request.MBDESC,
		BRANCH:      request.Request.BRANCH,
		BRDESC:      request.Request.BRDESC,
		JenisRevisi: request.Request.JenisRevisi,
		UpdateId:    request.Request.UpdateId,
		UpdateName:  request.Request.UpdateName,
		UpdateStell: request.Request.UpdateStell,
		Pernr:       request.Pernr,
	}

	err = lib.MakeRequest("POST", Url, headers, requestBody, &response)

	if err != nil {
		r.logger.Zap.Error(err)
	}

	return response, err
}

// GetDataRealisasi implements RealisasiDefinition.
func (r RealisasiService) GetDataRealisasi(request *models.RealisasiGetHeaderRequest) (response models.ResponseSampleData, err error) {
	baseUrl, err := lib.GetVarEnv("RealisasiUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "/datarealisasi/getData"

	token := r.jwtService.CreateRealisasiToken(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	requestBody := models.RealisasiBodyRequest{
		Periode:  request.Request.Periode,
		Branch:   request.Request.Branch,
		Restruck: request.Request.Restruck,
		Limit:    request.Request.Limit,
		Offset:   request.Request.Offset,
		Pernr:    request.Pernr,
	}

	err = lib.MakeRequest("POST", Url, headers, requestBody, &response)

	if err != nil {
		r.logger.Zap.Error(err)
	}

	return response, err
}

// UpdateFlagVerifikasi implements RealisasiDefinition.
func (r RealisasiService) UpdateFlagVerifikasi(request *models.RealisasiUpdateFlagRequest) (response models.UpdateFlagResponse, err error) {
	baseUrl, err := lib.GetVarEnv("RealisasiUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "/datarealisasi/updateFlag"

	token := r.jwtService.CreateRealisasiToken(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	requestBody := models.RequestUpdateFlag{
		Id:             request.Request.Id,
		VerifikasiFlag: request.Request.VerifikasiFlag,
		Pernr:          request.Pernr,
	}

	err = lib.MakeRequest("POST", Url, headers, requestBody, &response)

	if err != nil {
		r.logger.Zap.Error(err)
	}

	return response, err
}
