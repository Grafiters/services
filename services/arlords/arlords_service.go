package arlords

import (
	"fmt"
	"net/url"
	"riskmanagement/dto"
	"riskmanagement/lib"
	jwt "riskmanagement/services/auth"
	"strconv"

	"gitlab.com/golang-package-library/logger"
)

type ArlordsServiceDefinition interface {
	GetControlAttribute(pernr string, code []string) (response dto.Response[dto.DtoRiskControlAttributeResponse], err error)
	CreateMappingEvent(pernr string, req dto.BulkMappingRiskEventRequest) (err error)
	GetAllMappingRiskEvent(pernr string, eventID []string) (response dto.Response[[]dto.DetailMappingEvent], err error)
	GetHirearcyBusinessProcess(pernr string, activity string) (response dto.Response[dto.HierarchyPagination], err error)
	GetMappingControlAttribute(pernr string, id int) (response dto.Response[[]dto.ListAttributeMap], err error)
	GetTopicAnomaly(pernr string) (response dto.Response[dto.TopicListWithPaginate], err error)
	GetMapHeaderIndicator(pernr string, indicatorID int) (response dto.ListIndicatorHeaderResponse, err error)
	GetHeader(pernr, topic string) (response dto.Response[[]string], err error)
	GetUker(pernr string, IndicatorID int) (response dto.Response[[]dto.UkerDevision], err error)
	CreateSetUker(pernr string, req []dto.UpdateSelectedUker) (err error)
	CreateSetHeader(pernr string, req []dto.IndicatorHeader) (err error)
	GetBusinessCycle(pernr string) (response dto.Response[dto.BusinessProcessPagination], err error)
	GetMappingEventBusinessProess(pernr string, issueID int) (response dto.Response[dto.MappingRiskEventBusinesProcessRelationRespnse], err error)
	BulkGetBusinessProcessByActivity(pernr string, code []string) (response dto.Response[[]dto.BusinessHierarchyFlatResponse], err error)
	BulkGetMappingEventBusinessProcess(pernr string, code []string) (response dto.Response[[]dto.MappingRiskEventBusinesProcessRelation], err error)
	BulkCreateMappingBusinessProcess(pernr string, data []dto.MapingBusinessProcessRequest) (err error)
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

func (as ArlordService) GetTopicAnomaly(pernr string) (response dto.Response[dto.TopicListWithPaginate], err error) {
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

	pathUrl := fmt.Sprintf("%s/topic", baseUrl)
	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to save mapping cause: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) GetMappingControlAttribute(pernr string, id int) (response dto.Response[[]dto.ListAttributeMap], err error) {
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

	pathUrl := fmt.Sprintf("%s/control-attribute/risk-control/%s/list", baseUrl, strconv.Itoa(id))

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

func (as ArlordService) GetHeader(pernr, topic string) (response dto.Response[[]string], err error) {
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

	pathUrl := fmt.Sprintf("%s/verification/anomalies/anomaly-attributes/%s", baseUrl, topic)

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get data business process hierarcy: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) GetMapHeaderIndicator(pernr string, indicatorID int) (response dto.ListIndicatorHeaderResponse, err error) {
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

	pathUrl := fmt.Sprintf("%s/indicator/%d/header", baseUrl, indicatorID)
	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get data map indicator header: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) GetUker(pernr string, IndicatorID int) (response dto.Response[[]dto.UkerDevision], err error) {
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

	u, _ := url.Parse(baseUrl + "/uker/getAll")
	q := u.Query()
	if IndicatorID != 0 {
		q.Add("indicator_id", strconv.Itoa(IndicatorID))
	}

	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get data uker: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) CreateSetUker(pernr string, req []dto.UpdateSelectedUker) (err error) {
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

	u, _ := url.Parse(baseUrl + "/indicator/bulk-update-indicator-uker")
	pathUrl := u.String()

	request := dto.RequestSelectedUker{
		Data: req,
	}

	var response dto.HttpResResponse
	err = lib.MakeRequest("POST", pathUrl, headers, request, &response)
	if err != nil {
		as.logger.Zap.Error("Error when to save indicator uker: %s ", err)
		return err
	}

	return nil
}

func (as ArlordService) CreateSetHeader(pernr string, req []dto.IndicatorHeader) (err error) {
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

	u, _ := url.Parse(baseUrl + "/indicator/bulk-create-header")
	pathUrl := u.String()

	request := dto.IndicatorHeaderRequest{
		Data: req,
	}

	var response dto.HttpResResponse
	err = lib.MakeRequest("POST", pathUrl, headers, request, &response)
	if err != nil {
		as.logger.Zap.Error("Error when to save mapping indicator header: %s ", err)
		return err
	}

	return nil
}

func (as ArlordService) GetBusinessCycle(pernr string) (response dto.Response[dto.BusinessProcessPagination], err error) {
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

	u, _ := url.Parse(baseUrl + "/activity/")
	q := u.Query()
	q.Add("limit", "10000")

	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get data uker: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) GetMappingEventBusinessProess(pernr string, issueID int) (response dto.Response[dto.MappingRiskEventBusinesProcessRelationRespnse], err error) {
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

	path := fmt.Sprintf("/event/%d/business-process", issueID)
	u, _ := url.Parse(baseUrl + path)
	q := u.Query()
	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get data mapping business: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) BulkGetBusinessProcessByActivity(pernr string, code []string) (response dto.Response[[]dto.BusinessHierarchyFlatResponse], err error) {
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

	path := "/business-process/bulk-get-process"
	u, _ := url.Parse(baseUrl + path)
	q := u.Query()
	for _, v := range code {
		q.Add("activity_code", v)
	}
	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get bulk data business: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) BulkGetMappingEventBusinessProcess(pernr string, code []string) (response dto.Response[[]dto.MappingRiskEventBusinesProcessRelation], err error) {
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

	path := "/event/bulk-get-mapping-business-process"
	u, _ := url.Parse(baseUrl + path)
	q := u.Query()
	for _, v := range code {
		q.Add("activity_code", v)
	}
	u.RawQuery = q.Encode()
	pathUrl := u.String()

	err = lib.MakeRequest("GET", pathUrl, headers, nil, &response)
	if err != nil {
		as.logger.Zap.Error("Error when request to get bulk data business: %s", err)
		return response, err
	}

	return response, nil
}

func (as ArlordService) BulkCreateMappingBusinessProcess(pernr string, data []dto.MapingBusinessProcessRequest) (err error) {
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

	u, _ := url.Parse(baseUrl + "/event//business-process")
	pathUrl := u.String()

	var response dto.HttpResResponse
	err = lib.MakeRequest("POST", pathUrl, headers, data, &response)
	if err != nil {
		as.logger.Zap.Error("Error when to save mapping risk event: %s ", err)
		return err
	}

	return nil
}
