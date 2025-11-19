package krid

import (
	"encoding/base64"
	"fmt"
	"reflect"
	models "riskmanagement/models/krid"
	repo "riskmanagement/repository/riskindicator"
	"strings"

	libOPRA "riskmanagement/lib"

	lib "gitlab.com/golang-package-library/goresums"

	"gitlab.com/golang-package-library/logger"
)

type KridDefinition interface {
	GetDetailIndikator(request *models.HeaderRequest) (response []models.Krid, err error)
	GetAllParameterIndikator() (response []models.KeyRiskIndicator, err error)
	SearchIndikatorKRI(request *models.KeywordSearch) (response []models.KeyRiskIndicator, err error)
	SearchIndikatorKRIEdit(request *models.KeywordSearchEdit) (response []models.KeyRiskIndicator, err error)
}

type KridService struct {
	db     libOPRA.Database
	logger logger.Logger
	repo   repo.RiskIndicatorDefinition
}

func NewKridService(
	db libOPRA.Database,
	logger logger.Logger,
	repo repo.RiskIndicatorDefinition,
) KridDefinition {
	return KridService{
		db:     db,
		logger: logger,
		repo:   repo,
	}
}

// GetDetailIndikator implements KridDefinition
func (krid KridService) GetDetailIndikator(request *models.HeaderRequest) (response []models.Krid, err error) {
	baseUrlKRID, err := libOPRA.GetVarEnv("KRIDUrl")
	if err != nil {
		return nil, err
	}

	usernameKRID, err := libOPRA.GetVarEnv("KRIDUsername")
	if err != nil {
		return nil, err
	}

	credsKRID, err := libOPRA.GetVarEnv("KRIDPassword")
	if err != nil {
		return nil, err
	}

	options := lib.Options{
		BaseUrl: baseUrlKRID,
		SSL:     false,
		Payload: models.HeaderRequest{
			Request: request.Request,
		},
		Method: "POST",
		Auth:   false,
	}

	basicAuth := usernameKRID + ":" + credsKRID
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(basicAuth))

	auth := lib.Auth{
		Authorization: "Basic " + encodedAuth,
	}

	options.BaseUrl = baseUrlKRID + "krid/GetDetailIndikator"
	responseObject, err := lib.AuthBearer(options, auth)
	if err != nil {
		krid.logger.Zap.Error(err)
		return response, err
	}

	// fmt.Println("responsesObjectJwt", responseObject)
	// responseMessage := responseObject["RESPONSE"].(map[string]interface{})["RESPONSE_MESSAGE"].(string)
	responseMessage := responseObject["response"].(map[string]interface{})["responseMessage"].(string)

	// fmt.Println("statusResponse ====>", responseMessage)

	// responseData := responseObject["RESPONSE"]
	responseData := responseObject["response"]

	// data := ""
	status := fmt.Sprint(responseMessage)
	fmt.Println("status response", reflect.TypeOf(status))

	dataResponse := []models.Krid{}

	if status == "Success" {
		krid.logger.Zap.Info("getDetailIndikator")
		fmt.Println("==========================")
		fmt.Println("======Data Indicator======")
		for _, data := range responseData.(map[string]interface{}) {
			if fmt.Sprint(reflect.TypeOf(data)) == "[]interface {}" {
				for _, dataObj := range data.([]interface{}) {

					// aktifitas := dataObj.(map[string]interface{})["AKTIVITAS"]
					// idIndikator := dataObj.(map[string]interface{})["ID_Indikator"]
					// periode := dataObj.(map[string]interface{})["PERIODE"]
					// produk := dataObj.(map[string]interface{})["PRODUK"]
					// riskIssue := dataObj.(map[string]interface{})["RISK_ISSUE"]
					// unitKerja := dataObj.(map[string]interface{})["UNIT_KERJA"]
					aktifitas := dataObj.(map[string]interface{})["aktivitas"]
					idIndikator := dataObj.(map[string]interface{})["idIndikator"]
					periode := dataObj.(map[string]interface{})["periode"]
					produk := dataObj.(map[string]interface{})["produk"]
					riskIssue := dataObj.(map[string]interface{})["riskIssue"]
					unitKerja := dataObj.(map[string]interface{})["unitKerja"]

					// var structFields []reflect.StructField

					// renderContent := []models.Content{}
					Content := []models.Content{}

					for _, dataContent := range dataObj.(map[string]interface{}) {
						if fmt.Sprint(reflect.TypeOf(dataContent)) == "[]interface {}" {
							// fmt.Println(len(dataContent.([]interface{})))
							for _, content := range dataContent.([]interface{}) {
								Content = append(Content, content.(map[string]interface{}))
							}
						}
					}

					subData := models.Krid{
						Periode:     periode.(string),
						UnitKerja:   unitKerja.(string),
						IdIndikator: idIndikator.(string),
						Aktivitas:   aktifitas.(string),
						Produk:      produk.(string),
						RiskIssue:   riskIssue.(string),
						Content:     Content,
					}

					dataResponse = append(dataResponse, subData)
				}
			}
		}
	}

	// fmt.Println(dataResponse)
	return dataResponse, err
}

func (krid KridService) GetAllParameterIndikator() (response []models.KeyRiskIndicator, err error) {
	// fmt.Println("request", request.Request)
	baseUrlKRID, err := libOPRA.GetVarEnv("KRIDUrl")
	if err != nil {
		return nil, err
	}

	usernameKRID, err := libOPRA.GetVarEnv("KRIDUsername")
	if err != nil {
		return nil, err
	}

	credsKRID, err := libOPRA.GetVarEnv("KRIDPassword")
	if err != nil {
		return nil, err
	}

	options := lib.Options{
		BaseUrl: baseUrlKRID,
		SSL:     false,
		Method:  "POST",
		Auth:    false,
	}

	basicAuth := usernameKRID + ":" + credsKRID
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(basicAuth))

	auth := lib.Auth{
		Authorization: "Basic " + encodedAuth,
	}

	options.BaseUrl = baseUrlKRID + "krid/GetAllParameterIndikator"
	responseObject, err := lib.AuthBearer(options, auth)
	if err != nil {
		krid.logger.Zap.Error(err)
		return response, err
	}

	// fmt.Println("responsesObjectJwt", responseObject)
	// responseMessage := responseObject["RESPONSE"].(map[string]interface{})["RESPONSE_MESSAGE"].(string)
	responseMessage := responseObject["response"].(map[string]interface{})["responseMessage"].(string)

	// fmt.Println("statusResponse ====>", responseMessage)

	responseData := responseObject["response"]

	// data := ""
	status := fmt.Sprint(responseMessage)
	fmt.Println("status response", reflect.TypeOf(status))

	dataResponse := []models.KeyRiskIndicator{}

	if status == "Success" {
		krid.logger.Zap.Info("GetAllParameterIndikator")
		fmt.Println("===================================")
		fmt.Println("======Data Key Risk Indicator======")
		for _, data := range responseData.(map[string]interface{}) {
			if fmt.Sprint(reflect.TypeOf(data)) == "[]interface {}" {
				for _, dataObj := range data.([]interface{}) {
					// fmt.Println("dataObj  =====>", dataObj)

					// id := dataObj.(map[string]interface{})["ID"]
					// keyRiskIndicator := dataObj.(map[string]interface{})["KEY_RISK_INDICATOR"]
					// aktivitas := dataObj.(map[string]interface{})["AKTIVITAS"]
					// produk := dataObj.(map[string]interface{})["PRODUK"]
					// jenisIndikator := dataObj.(map[string]interface{})["JENIS_INDIKATOR"]
					// indikasiRisiko := dataObj.(map[string]interface{})["INDIKASI_RISIKO"]
					id := dataObj.(map[string]interface{})["id"]
					keyRiskIndicator := dataObj.(map[string]interface{})["keyRiskIndicator"]
					aktivitas := dataObj.(map[string]interface{})["aktivitas"]
					produk := dataObj.(map[string]interface{})["produk"]
					jenisIndikator := dataObj.(map[string]interface{})["jenisIndikator"]
					indikasiRisiko := dataObj.(map[string]interface{})["indikasiRisiko"]

					// fmt.Println(id, keyRiskIndicator)
					subData := models.KeyRiskIndicator{
						ID:                   id.(string),
						KodeKeyRiskIndicator: id.(string),
						KeyRiskIndicator:     keyRiskIndicator.(string),
						Aktifitas:            aktivitas.(string),
						Produk:               produk.(string),
						JenisIndicator:       jenisIndikator.(string),
						IndikasiRisiko:       indikasiRisiko.(string),
					}

					dataResponse = append(dataResponse, subData)
				}
			}
		}
	}

	// fmt.Println(dataResponse)
	return dataResponse, err
}

// SearchIndikatorKRI implements KridDefinition
func (krid KridService) SearchIndikatorKRI(request *models.KeywordSearch) (response []models.KeyRiskIndicator, err error) {
	data, err := krid.GetAllParameterIndikator()
	if err != nil {
		return response, err
	}

	for _, item := range data {
		if strings.ToLower(item.Aktifitas) == strings.ToLower(request.Aktifitas) && strings.ToLower(item.Produk) == strings.ToLower(request.Produk) {
			if request.Keyword != "" {
				if strings.Contains(strings.ToLower(item.KeyRiskIndicator), strings.ToLower(request.Keyword)) {
					// fmt.Println("filtered data :", item.KeyRiskIndicator)
					response = append(response, item)
				}
			}
		}
	}

	return response, err
}

func (krid KridService) SearchIndikatorKRIEdit(request *models.KeywordSearchEdit) (response []models.KeyRiskIndicator, err error) {
	data, err := krid.GetAllParameterIndikator()
	if err != nil {
		return response, err
	}

	activity, err := krid.repo.GetByID(request.Aktifitas)

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range data {
		if strings.ToLower(item.Aktifitas) == strings.ToLower(activity.Name) && strings.ToLower(item.Produk) == strings.ToLower(request.Produk) {
			if request.Keyword != "" {
				if strings.Contains(strings.ToLower(item.KeyRiskIndicator), strings.ToLower(request.Keyword)) {
					// fmt.Println("filtered data :", item.KeyRiskIndicator)
					response = append(response, item)
				}
			}
		}
	}

	return response, err
}
