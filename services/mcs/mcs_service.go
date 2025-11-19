package repMcs

import (
	"fmt"
	"reflect"
	models "riskmanagement/models/mcs"
	libOPRA "riskmanagement/lib"

	lib "gitlab.com/golang-package-library/goresums"

	"gitlab.com/golang-package-library/logger"
)

type McsDefinition interface {
	GetUker(request *models.McsRequest) (response []models.UkerResponse, err error)
	GetPIC(request *models.McsRequest) (response []models.PICResponse, err error)
	GetJabatan(request *models.McsRequest) (response []models.JabatanResponse, err error)
	GetOrgeh(request *models.McsRequest) (response []models.OrgehResponse, err error)
}

type McsService struct {
	logger logger.Logger
}

func NewMcsService(logger logger.Logger) McsDefinition {
	return McsService{
		logger: logger,
	}
}

// GetUker implements McsDefinition
func (mcs McsService) GetUker(request *models.McsRequest) (response []models.UkerResponse, err error) {
	fmt.Println("request", request)
	oneGateUrl, err := libOPRA.GetVarEnv("OnegateURL")
	if err != nil {
        return nil, err
    }

	oneGateClientId, err := libOPRA.GetVarEnv("OnegateClientID")
	if err != nil {
        return nil, err
    }

	oneGateSecret, err := libOPRA.GetVarEnv("OnegateSecret")
	if err != nil {
        return nil, err
    }

	jwt := ""
	options := lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Clientid:     oneGateClientId,
			Clientsecret: oneGateSecret,
			Keyword:      request.Keyword,
			Limit:        request.Limit,
			Offset:       request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth := lib.Auth{
		Authorization: "Bearer " + jwt,
	}

	options.BaseUrl = oneGateUrl + "api/v1/client_auth/request_token"
	responseObjectJwt, err := lib.AuthBearer(options, auth)
	if err != nil {
		mcs.logger.Zap.Error(err)
		return response, err
	}
	fmt.Println("responseObjectJwt", responseObjectJwt)
	statusResponseJwt := responseObjectJwt["success"]
	dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"].(string)

	fmt.Println("statusResponseJwt", statusResponseJwt)
	fmt.Println("dataResponseJwt", dataResponseJwt)
	fmt.Println("===============================================")
	fmt.Println("====================JWT AUTH===================")

	fmt.Println("request", request)
	jwt = ""
	options = lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Keyword: request.Keyword,
			Limit:   request.Limit,
			Offset:  request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth = lib.Auth{
		Authorization: "Bearer " + dataResponseJwt,
	}

	dataResponse := []models.UkerResponse{}
	status := fmt.Sprint(statusResponseJwt)
	fmt.Println("status response", reflect.TypeOf(status))
	if status == "true" {
		mcs.logger.Zap.Info("Search Uker")
		options.BaseUrl = oneGateUrl + "api/v1/uker/searchUker"
		responseObjectSession, err := lib.AuthBearer(options, auth)
		if err != nil {
			mcs.logger.Zap.Error(err)
			return response, err
		}
		fmt.Println(responseObjectSession)
		statusResponseSession := responseObjectSession["success"]
		dataResponseSession := responseObjectSession["message"]

		fmt.Println("statusResponseSession", statusResponseSession)
		fmt.Println("dataResponseSession", dataResponseSession)
		fmt.Println("==========================================================")
		fmt.Println("====================DATA MCS==============================")

		fmt.Println("check interface or []interface", reflect.TypeOf(dataResponseSession))
		fmt.Println("dataResponseSession", fmt.Sprint(reflect.TypeOf(dataResponseSession)))

		if fmt.Sprint(reflect.TypeOf(dataResponseSession)) == "[]interface {}" {
			for _, data := range dataResponseSession.([]interface{}) {
				BRNAME := data.(map[string]interface{})["brname"]
				BRANCH := data.(map[string]interface{})["branch"]

				if fmt.Sprint(reflect.TypeOf(BRANCH)) == "float64" {
					subData := models.UkerResponse{
						BRNAME: BRNAME.(string),
						BRANCH: fmt.Sprint(int(BRANCH.(float64))),
					}
					dataResponse = append(dataResponse, subData)
				} else {
					subData := models.UkerResponse{
						BRNAME: BRNAME.(string),
						BRANCH: BRANCH.(string),
					}
					dataResponse = append(dataResponse, subData)

				}
			}
		}

	}

	fmt.Println("dataResponse", dataResponse)
	return dataResponse, err
}

// GetPIC implements McsDefinition
func (mcs McsService) GetPIC(request *models.McsRequest) (response []models.PICResponse, err error) {
	oneGateUrl, err := libOPRA.GetVarEnv("OnegateURL")
	if err != nil {
        return nil, err
    }

	oneGateClientId, err := libOPRA.GetVarEnv("OnegateClientID")
	if err != nil {
        return nil, err
    }

	oneGateSecret, err := libOPRA.GetVarEnv("OnegateSecret")
	if err != nil {
        return nil, err
    }

	fmt.Println("request", request)
	jwt := ""
	options := lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Clientid:     oneGateClientId,
			Clientsecret: oneGateSecret,
			Keyword:      request.Keyword,
			Limit:        request.Limit,
			Offset:       request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth := lib.Auth{
		Authorization: "Bearer " + jwt,
	}

	options.BaseUrl = oneGateUrl + "api/v1/client_auth/request_token"
	responseObjectJwt, err := lib.AuthBearer(options, auth)
	if err != nil {
		mcs.logger.Zap.Error(err)
		return response, err
	}
	fmt.Println("responseObjectJwt", responseObjectJwt)
	statusResponseJwt := responseObjectJwt["success"]
	dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"].(string)

	fmt.Println("statusResponseJwt", statusResponseJwt)
	fmt.Println("dataResponseJwt", dataResponseJwt)
	fmt.Println("===============================================")
	fmt.Println("====================JWT AUTH===================")

	fmt.Println("request", request)
	jwt = ""
	options = lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Keyword: request.Keyword,
			Limit:   request.Limit,
			Offset:  request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth = lib.Auth{
		Authorization: "Bearer " + dataResponseJwt,
	}

	dataResponse := []models.PICResponse{}
	status := fmt.Sprint(statusResponseJwt)
	fmt.Println("status response", reflect.TypeOf(status))
	if status == "true" {
		mcs.logger.Zap.Info("Search Pekerja")
		options.BaseUrl = oneGateUrl + "api/v1/pekerja/searchPekerja"
		responseObjectSession, err := lib.AuthBearer(options, auth)
		if err != nil {
			mcs.logger.Zap.Error(err)
			return response, err
		}
		fmt.Println(responseObjectSession)
		statusResponseSession := responseObjectSession["success"]
		dataResponseSession := responseObjectSession["message"]

		fmt.Println("statusResponseSession", statusResponseSession)
		fmt.Println("dataResponseSession", dataResponseSession)
		fmt.Println("==========================================================")
		fmt.Println("====================DATA MCS==============================")

		fmt.Println("check interface or []interface", reflect.TypeOf(dataResponseSession))
		fmt.Println("dataResponseSession", fmt.Sprint(reflect.TypeOf(dataResponseSession)))

		if fmt.Sprint(reflect.TypeOf(dataResponseSession)) == "[]interface {}" {
			for _, data := range dataResponseSession.([]interface{}) {
				PERNR := data.(map[string]interface{})["PERNR"]
				HTEXT := data.(map[string]interface{})["HTEXT"]
				NAMA := data.(map[string]interface{})["NAMA"]

				if fmt.Sprint(reflect.TypeOf(PERNR)) == "float64" {
					subData := models.PICResponse{
						PERNR: fmt.Sprint(int(PERNR.(float64))),
						HTEXT: HTEXT.(string),
						NAMA:  NAMA.(string),
					}
					dataResponse = append(dataResponse, subData)
				} else {
					subData := models.PICResponse{
						PERNR: PERNR.(string),
						HTEXT: HTEXT.(string),
						NAMA:  NAMA.(string),
					}
					dataResponse = append(dataResponse, subData)
				}
			}
		}

	}

	fmt.Println("dataResponse", dataResponse)
	return dataResponse, err
}

// GetJabatan implements McsDefinition
func (mcs McsService) GetJabatan(request *models.McsRequest) (response []models.JabatanResponse, err error) {
	fmt.Println("request", request)
	oneGateUrl, err := libOPRA.GetVarEnv("OnegateURL")
	if err != nil {
        return nil, err
    }

	oneGateClientId, err := libOPRA.GetVarEnv("OnegateClientID")
	if err != nil {
        return nil, err
    }

	oneGateSecret, err := libOPRA.GetVarEnv("OnegateSecret")
	if err != nil {
        return nil, err
    }

	jwt := ""
	options := lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Clientid:     oneGateClientId,
			Clientsecret: oneGateSecret,
			Keyword:      request.Keyword,
			Limit:        request.Limit,
			Offset:       request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth := lib.Auth{
		Authorization: "Bearer " + jwt,
	}

	options.BaseUrl = oneGateUrl + "api/v1/client_auth/request_token"
	responseObjectJwt, err := lib.AuthBearer(options, auth)
	if err != nil {
		mcs.logger.Zap.Error(err)
		return response, err
	}
	fmt.Println("responseObjectJwt", responseObjectJwt)
	statusResponseJwt := responseObjectJwt["success"]
	dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"].(string)

	fmt.Println("statusResponseJwt", statusResponseJwt)
	fmt.Println("dataResponseJwt", dataResponseJwt)
	fmt.Println("===============================================")
	fmt.Println("====================JWT AUTH===================")

	fmt.Println("request", request)
	jwt = ""
	options = lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Keyword: request.Keyword,
			Limit:   request.Limit,
			Offset:  request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth = lib.Auth{
		Authorization: "Bearer " + dataResponseJwt,
	}

	dataResponse := []models.JabatanResponse{}
	status := fmt.Sprint(statusResponseJwt)
	fmt.Println("status response", reflect.TypeOf(status))
	if status == "true" {
		mcs.logger.Zap.Info("Search Jabatan")
		options.BaseUrl = oneGateUrl + "api/v1/pekerja/searchJabatan"
		responseObjectSession, err := lib.AuthBearer(options, auth)
		if err != nil {
			mcs.logger.Zap.Error(err)
			return response, err
		}
		fmt.Println(responseObjectSession)
		statusResponseSession := responseObjectSession["success"]
		dataResponseSession := responseObjectSession["message"]
		fmt.Println("statusResponseSession", statusResponseSession)
		fmt.Println("dataResponseSession", dataResponseSession)
		fmt.Println("==========================================================")
		fmt.Println("====================DATA MCS==============================")

		fmt.Println("check interface or []interface", reflect.TypeOf(dataResponseSession))
		fmt.Println("dataResponseSession", fmt.Sprint(reflect.TypeOf(dataResponseSession)))

		if fmt.Sprint(reflect.TypeOf(dataResponseSession)) == "[]interface {}" {
			for _, data := range dataResponseSession.([]interface{}) {
				HILFM := data.(map[string]interface{})["HILFM"]
				HTEXT := data.(map[string]interface{})["HTEXT"]

				if fmt.Sprint(reflect.TypeOf(HILFM)) == "floast64" {
					subData := models.JabatanResponse{
						HILFM: fmt.Sprint(int(HILFM.(float64))),
						HTEXT: HTEXT.(string),
					}
					dataResponse = append(dataResponse, subData)
				} else {
					subData := models.JabatanResponse{
						HILFM: HILFM.(string),
						HTEXT: HTEXT.(string),
					}
					dataResponse = append(dataResponse, subData)
				}
			}
		}
	}
	fmt.Println("dataResponse", dataResponse)
	return dataResponse, err
}

// GetOrgeh implements McsDefinition
func (mcs McsService) GetOrgeh(request *models.McsRequest) (response []models.OrgehResponse, err error) {
	fmt.Println("request", request)
	oneGateUrl, err := libOPRA.GetVarEnv("OnegateURL")
	if err != nil {
        return nil, err
    }

	oneGateClientId, err := libOPRA.GetVarEnv("OnegateClientID")
	if err != nil {
        return nil, err
    }

	oneGateSecret, err := libOPRA.GetVarEnv("OnegateSecret")
	if err != nil {
        return nil, err
    }

	jwt := ""
	options := lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Clientid:     oneGateClientId,
			Clientsecret: oneGateSecret,
			Keyword:      request.Keyword,
			Limit:        request.Limit,
			Offset:       request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth := lib.Auth{
		Authorization: "Bearer " + jwt,
	}

	options.BaseUrl = oneGateUrl + "api/v1/client_auth/request_token"
	responseObjectJwt, err := lib.AuthBearer(options, auth)
	if err != nil {
		mcs.logger.Zap.Error(err)
		return response, err
	}
	fmt.Println("responseObjectJwt", responseObjectJwt)
	statusResponseJwt := responseObjectJwt["success"]
	dataResponseJwt := responseObjectJwt["message"].(map[string]interface{})["token"].(map[string]interface{})["token"].(string)

	fmt.Println("statusResponseJwt", statusResponseJwt)
	fmt.Println("dataResponseJwt", dataResponseJwt)
	fmt.Println("===============================================")
	fmt.Println("====================JWT AUTH===================")

	fmt.Println("request", request)
	jwt = ""
	options = lib.Options{
		BaseUrl: oneGateUrl,
		SSL:     false,
		Payload: models.McsRequest{
			Keyword: request.Keyword,
			Limit:   request.Limit,
			Offset:  request.Offset,
		},
		Method: "POST",
		Auth:   false,
	}

	auth = lib.Auth{
		Authorization: "Bearer " + dataResponseJwt,
	}

	dataResponse := []models.OrgehResponse{}
	status := fmt.Sprint(statusResponseJwt)
	fmt.Println("status response", reflect.TypeOf(status))
	if status == "true" {
		mcs.logger.Zap.Info("Search orgeh")
		options.BaseUrl = oneGateUrl + "api/v1/pekerja/searchOrgeh"
		responseObjectSession, err := lib.AuthBearer(options, auth)
		if err != nil {
			mcs.logger.Zap.Error(err)
			return response, err
		}
		fmt.Println(responseObjectSession)
		statusResponseSession := responseObjectSession["success"]
		dataResponseSession := responseObjectSession["message"]
		fmt.Println("statusResponseSession", statusResponseSession)
		fmt.Println("dataResponseSession", dataResponseSession)
		fmt.Println("==========================================================")
		fmt.Println("====================DATA MCS==============================")

		fmt.Println("check interface or []interface", reflect.TypeOf(dataResponseSession))
		fmt.Println("dataResponseSession", fmt.Sprint(reflect.TypeOf(dataResponseSession)))

		if fmt.Sprint(reflect.TypeOf(dataResponseSession)) == "[]interface {}" {
			for _, data := range dataResponseSession.([]interface{}) {
				// WERKS := data.(map[string]interface{})["WERKS"]
				// WERKSTX := data.(map[string]interface{})["WERKS_TX"]
				// BTRTL := data.(map[string]interface{})["BTRTL"]
				// BTRTLTX := data.(map[string]interface{})["BTRTL_TX"]
				// KOSTL := data.(map[string]interface{})["KOSTL"]
				// KOSTLTX := data.(map[string]interface{})["KOSTL_TX"]
				ORGEH := data.(map[string]interface{})["ORGEH"]
				ORGEHTX := data.(map[string]interface{})["ORGEH_TX"]

				if fmt.Sprint(reflect.TypeOf(ORGEH)) == "float64" {
					subData := models.OrgehResponse{
						// WERKS:   WERKS.(string),
						// WERKSTX: WERKSTX.(string),
						// BTRTL:   BTRTL.(string),
						// BTRTLTX: BTRTLTX.(string),
						// KOSTL:   KOSTL.(string),
						// KOSTLTX: KOSTLTX.(string),
						ORGEH:   fmt.Sprint(int(ORGEH.(float64))),
						ORGEHTX: ORGEHTX.(string),
					}

					dataResponse = append(dataResponse, subData)
				} else {
					subData := models.OrgehResponse{
						// WERKS:   WERKS.(string),
						// WERKSTX: WERKSTX.(string),
						// BTRTL:   BTRTL.(string),
						// BTRTLTX: BTRTLTX.(string),
						// KOSTL:   KOSTL.(string),
						// KOSTLTX: KOSTLTX.(string),
						ORGEH:   ORGEH.(string),
						ORGEHTX: ORGEHTX.(string),
					}

					dataResponse = append(dataResponse, subData)
				}

			}
		}
	}
	fmt.Println("dataResponse", dataResponse)
	return dataResponse, err
}
