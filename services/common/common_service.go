package common

import (
	commonRepo "riskmanagement/repository/common"

	"gitlab.com/golang-package-library/logger"

	"riskmanagement/lib"
	models "riskmanagement/models/common"
	common "riskmanagement/repository/common"
)

type CommonDefinition interface {
	FilterPnNama(request models.KeywordRequest) (responses []models.PnNamaResult, err error)
	FilterKanwil(request models.KeywordRequest) (responses []models.KanwilResult, err error)
	FilterKanca(request models.KeywordRequest) (responses []models.KancaResult, err error)
	FilterUker(request models.KeywordRequest) (responses []models.UkerResult, err error)
	FilterRiskEventByActifityAndPoduct(request models.RiskEventRequest) (responses []models.RiskEventResult, err error)
	FilterRiskIndikatorByRiskEvent(request models.RiskIndikatorRequest) (responses []models.RiskIndikatorResult, err error)

	FilterRRMHead(request models.RRMHeadRequest) (responses []models.PnNamaResult, err error)
	FilterPimpinanUker(request models.PimpinanUkerRequest) (responses []models.PnNamaResult, err error)

	GetApprovalResponse(request models.CommonRequest) (response []models.PnNamaResult, err error)
	GetAllORDMember() (response []models.ORDMember, err error)

	// MQ Enhance
	GetMstDataOption() (response []models.MasterDataOption, err error)

	SearchBrc(request models.BrcKeywordRequest) (responses []models.PnNamaResult, pagination lib.Pagination, err error)
}

type CommonService struct {
	db         lib.Database
	logger     logger.Logger
	commonRepo common.CommonDefinition
}

// GetMstDataOption implements CommonDefinition.
func (c CommonService) GetMstDataOption() (response []models.MasterDataOption, err error) {
	data, err := c.commonRepo.GetMstDataOption()

	if err != nil {
		c.logger.Zap.Error(err)
		return response, err
	}

	return data, nil
}

// FilterPimpinanUker implements CommonDefinition.
func (c CommonService) FilterPimpinanUker(request models.PimpinanUkerRequest) (responses []models.PnNamaResult, err error) {
	data, err := c.commonRepo.GetPimpinanUker(request)

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return data, nil
}

// FilterRRMHead implements CommonDefinition.
func (c CommonService) FilterRRMHead(request models.RRMHeadRequest) (responses []models.PnNamaResult, err error) {
	data, err := c.commonRepo.GetRRMHead(request)

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return data, nil
}

func (c CommonService) FilterKanca(request models.KeywordRequest) (responses []models.KancaResult, err error) {
	c.logger.Zap.Info(request)
	data, err := c.commonRepo.GetKanca(request)

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return data, nil
}

func (c CommonService) FilterUker(request models.KeywordRequest) (responses []models.UkerResult, err error) {
	c.logger.Zap.Info(request)
	data, err := c.commonRepo.GetUker(request)

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return data, nil
}

func (c CommonService) FilterKanwil(request models.KeywordRequest) (responses []models.KanwilResult, err error) {
	c.logger.Zap.Info(request)
	data, err := c.commonRepo.GetKanwil(request)

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return data, nil
}

func (c CommonService) FilterPnNama(request models.KeywordRequest) (responses []models.PnNamaResult, err error) {
	c.logger.Zap.Info(request)
	response, err := c.commonRepo.GetNPAndNama(request)

	for _, data := range response {
		c.logger.Zap.Info(data.Sname)
		c.logger.Zap.Info(data.Pernr)
	}

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return response, nil
}

func (c CommonService) FilterRiskEventByActifityAndPoduct(request models.RiskEventRequest) (responses []models.RiskEventResult, err error) {
	c.logger.Zap.Info(request)

	data, err := c.commonRepo.GetRiskEventByActifityAndPoduct(request)

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return data, nil
}

func (c CommonService) FilterRiskIndikatorByRiskEvent(request models.RiskIndikatorRequest) (responses []models.RiskIndikatorResult, err error) {
	c.logger.Zap.Info(request)

	data, err := c.commonRepo.GetRiskIndikatorByRiskEvent(request)

	if err != nil {
		c.logger.Zap.Error(err)
		return responses, err
	}

	return data, nil
}

func (c CommonService) GetApprovalResponse(request models.CommonRequest) (response []models.PnNamaResult, err error) {
	data, err := c.commonRepo.GetApprovalResponse(request)

	if err != nil {
		return response, err
	}

	return data, err
}

func (c CommonService) GetAllORDMember() (response []models.ORDMember, err error) {
	return c.commonRepo.GetAllORDMember()
}

func NewCommonService(
	db lib.Database,
	logger logger.Logger,
	commonRepo commonRepo.CommonDefinition,
) CommonDefinition {
	return CommonService{
		db:         db,
		logger:     logger,
		commonRepo: commonRepo,
	}
}

// SearchBrc implements CommonDefinition.
func (c CommonService) SearchBrc(request models.BrcKeywordRequest) (responses []models.PnNamaResult, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataBrc, totalRows, totalData, err := c.commonRepo.SearchBrc(request)

	if err != nil {
		c.logger.Zap.Error("Service : ", err)
		return responses, pagination, err
	}

	for _, response := range dataBrc {
		responses = append(responses, models.PnNamaResult{
			Pernr: response.Pernr,
			Sname: response.Sname,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}
