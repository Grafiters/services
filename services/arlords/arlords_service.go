package arlords

import (
	"fmt"
	"net/url"
	"riskmanagement/dto"
	"riskmanagement/lib"
	jwt "riskmanagement/services/auth"

	"gitlab.com/golang-package-library/logger"
)

type ArlordsServiceDefinition interface {
	GetControlAttribute(pernr string, code []string) (response dto.Response[dto.DtoRiskControlAttributeResponse], err error)
	CreateMappingEvent(pernr string, req dto.BulkMappingRiskEventRequest) (err error)
	GetAllMappingRiskEvent(pernr string, eventID []string) (response dto.Response[[]dto.DetailMappingEvent], err error)
	GetHirearcyBusinessProcess(pernr string, activity string) (response dto.Response[dto.HierarchyPagination], err error)
}

type ArlordService struct {
	logger     logger.Logger
	jwtService jwt.JWTAuthService
}

func NewArlordService(
	logger logger.Logger,
	jwtService jwt.JWTAuthService,
) ArlordsServiceDefinition {
	return &ArlordService{
		logger:     logger,
		jwtService: jwtService,
	}
}

var baseUrl = ""

func (as ArlordService) GetControlAttribute(pernr string, code []string) (response dto.Response[dto.DtoRiskControlAttributeResponse], err error) {
	err = InitURL()
	if err != nil {
		return response, fmt.Errorf("failed to init module: %s", err)
	}
	query := dto.ControlAttributeFiler{
		CodeIDs: code,
		Status:  "active",
	}

	authToken := as.jwtService.CreateArlordsToken(pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + authToken,
		"Content-Type":  "application/json",
		"pernr":         pernr,
	}

	u, _ := url.Parse(baseUrl + "/control-attribute/")

	q := u.Query()
	for _, c := range query.CodeIDs {
		q.Add("code_ids", c)
	}
	if query.Status != "" {
		q.Add("status", query.Status)
	}

	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to save mapping cause: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) CreateMappingEvent(pernr string, req dto.BulkMappingRiskEventRequest) (err error) {
	err = InitURL()
	if err != nil {
		return fmt.Errorf("failed to init module: %s", err)
	}

	authToken := as.jwtService.CreateArlordsToken(pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + authToken,
		"Content-Type":  "application/json",
		"pernr":         pernr,
	}
	u, _ := url.Parse(baseUrl + "/event/bulk-create-mapping")
	pathUrl := u.String()

	var response dto.HttpResResponse
	err = lib.MakeRequest("POST", pathUrl, headers, req, &response)
	if err != nil {
		as.logger.Zap.Error("Error when to save mapping risk event: %s ", err)
		return err
	}

	return nil

}

func (as ArlordService) GetAllMappingRiskEvent(pernr string, eventID []string) (response dto.Response[[]dto.DetailMappingEvent], err error) {
	err = InitURL()
	if err != nil {
		return response, fmt.Errorf("failed to init module: %s", err)
	}

	authToken := as.jwtService.CreateArlordsToken(pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + authToken,
		"Content-Type":  "application/json",
		"pernr":         pernr,
	}

	u, _ := url.Parse(baseUrl + "/event/bulk-get-mapping")
	q := u.Query()
	for _, c := range eventID {
		q.Add("event_id", c)
	}

	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to save mapping cause: %s", err)
		return response, err
	}

	return response, nil
}

func InitURL() error {
	url, err := lib.GetVarEnv("ArlordsUrl")
	if err != nil {
		return fmt.Errorf("errored when got url arlods: %s", err)
	}

	baseUrl = url
	return nil
}

func (as ArlordService) GetHirearcyBusinessProcess(pernr string, activity string) (response dto.Response[dto.HierarchyPagination], err error) {
	err = InitURL()
	if err != nil {
		return response, fmt.Errorf("failed to init module: %s", err)
	}

	authToken := as.jwtService.CreateArlordsToken(pernr)
	headers := map[string]string{
		"Authorization": "Bearer " + authToken,
		"Content-Type":  "application/json",
		"pernr":         pernr,
	}

	u, _ := url.Parse(baseUrl + "/business-process/hierarchy")
	q := u.Query()
	q.Add("activity", activity)

	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get data business process hierarcy: %s", err)
		return response, err
	}

	return response, nil
}
