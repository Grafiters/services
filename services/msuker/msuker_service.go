package msuker

import (
	"riskmanagement/lib"
	models "riskmanagement/models/msuker"
	repository "riskmanagement/repository/msuker"

	"gitlab.com/golang-package-library/logger"
)

type MsUkerDefinition interface {
	GetAll() (responses []models.MsUkerResponse, err error)
	GetUkerByBranch(branchid int64) (responses []models.MsUkerResponse, err error)
	SearchUker(request models.KeywordRequest) (responses []models.SearchResponse, pagination lib.Pagination, err error)
	GetUkerPerRegion(request models.Region) (responses []models.MsUkerResponse, err error)
	SearchPeserta(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error)
	GetPekerjaByBranch(request models.BranchCodeInduk) (responses []models.MsPekerjaResponse, err error)
	GetPekerjaByRegion(request *models.SearchPNByRegionReq) (responses []models.SearchPNByRegionRes, err error)
	SearchJabatan(request models.KeywordRequest) (responses []models.Jabatan, pagination lib.Pagination, err error)
	SearchUkerByRegionPekerja(request models.KeyRequest) (responses []models.SearchResponse, pagination lib.Pagination, err error)
	SearchPekerjaPerUker(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error)
	SearchSigner(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error)
	SearchRMC(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error)
	SearchPelakuFraud(request models.KeywordRequest) (responses []models.MsPelaku, pagination lib.Pagination, err error)
	SearchBrcUrcPerRegion(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error)

	// Batch3
	ListingJabatanPerUker(request models.ListJabatanRequest) (responses []models.ListJabatanResponse, err error)
	GetPekerjaBranchByHILFM(request models.BranchByHilfmRequest) (responses []models.MsPeserta, err error)
	SearchRRMHead(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error)
	SearchPekerjaOrd(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error)
}

type MsUkerService struct {
	db         lib.Database
	logger     logger.Logger
	repository repository.MsUkerDefinition
}

func NewMsUkerService(
	db lib.Database,
	logger logger.Logger,
	repository repository.MsUkerDefinition,
) MsUkerDefinition {
	return MsUkerService{
		db:         db,
		logger:     logger,
		repository: repository,
	}
}

// GetAll implements MsUkerDefinition
func (msUker MsUkerService) GetAll() (responses []models.MsUkerResponse, err error) {
	return msUker.repository.GetAll()
}

// GetUkerByBranch implements MsUkerDefinition
func (msUker MsUkerService) GetUkerByBranch(brancid int64) (responses []models.MsUkerResponse, err error) {
	return msUker.repository.GetUkerByBranch(brancid)
}

// GetUkerByRegion implements MsUkerDefinition
func (msUker MsUkerService) GetUkerPerRegion(request models.Region) (responses []models.MsUkerResponse, err error) {
	dwhBranch, err := msUker.repository.GetUkerPerRegion(&request)

	if err != nil {
		msUker.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dwhBranch {
		responses = append(responses, models.MsUkerResponse{
			SRCSYSID: response.SRCSYSID,
			BRUNIT:   response.BRUNIT,
			REGION:   response.REGION,
			RGDESC:   response.RGDESC,
			RGNAME:   response.RGNAME,
			MAINBR:   response.MAINBR,
			MBDESC:   response.MBDESC,
			MBNAME:   response.MBNAME,
			SUBBR:    response.SUBBR,
			SBDESC:   response.SBDESC,
			SBNAME:   response.SBNAME,
			BRANCH:   response.BRANCH,
			BRDESC:   response.BRDESC,
			BRNAME:   response.BRNAME,
			BIBR:     response.BIBR,
		})
	}

	return responses, err
}

// SearchUker implements MsUkerDefinition
func (msUker MsUkerService) SearchUker(request models.KeywordRequest) (responses []models.SearchResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataUker, totalRows, totalData, err := msUker.repository.SearchUker(&request)
	if err != nil {
		msUker.logger.Zap.Error(err)
		return responses, pagination, err
	}

	// Check if totalData are valid
	if totalData < 0 {
		msUker.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataUker {
		responses = append(responses, models.SearchResponse{
			SRCSYSID: response.SRCSYSID.String,
			BRUNIT:   response.BRUNIT.String,
			REGION:   response.REGION.String,
			RGDESC:   response.RGDESC.String,
			RGNAME:   response.RGNAME.String,
			MAINBR:   response.MAINBR.Int64,
			MBDESC:   response.MBDESC.String,
			MBNAME:   response.MBNAME.String,
			SUBBR:    response.SUBBR.Int64,
			SBDESC:   response.SBDESC.String,
			SBNAME:   response.SBNAME.String,
			BRANCH:   response.BRANCH.String,
			BRDESC:   response.BRDESC.String,
			BRNAME:   response.BRNAME.String,
			BIBR:     response.BIBR.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// SearchUkerByRegionPekerja implements MsUkerDefinition
func (msUker MsUkerService) SearchUkerByRegionPekerja(request models.KeyRequest) (responses []models.SearchResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataUker, totalRows, totalData, err := msUker.repository.SearchUkerByRegionPekerja(&request)
	if err != nil {
		msUker.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		msUker.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataUker {
		responses = append(responses, models.SearchResponse{
			SRCSYSID: response.SRCSYSID,
			BRUNIT:   response.BRUNIT,
			REGION:   response.REGION,
			RGDESC:   response.RGDESC,
			RGNAME:   response.RGNAME,
			MAINBR:   response.MAINBR,
			MBDESC:   response.MBDESC,
			MBNAME:   response.MBNAME,
			SUBBR:    response.SUBBR,
			SBDESC:   response.SBDESC,
			SBNAME:   response.SBNAME,
			BRANCH:   response.BRANCH,
			BRDESC:   response.BRDESC,
			BRNAME:   response.BRNAME,
			BIBR:     response.BIBR,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// SearchJabatan implements MsUkerDefinition
func (msUker MsUkerService) SearchJabatan(request models.KeywordRequest) (responses []models.Jabatan, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataJabatan, totalRows, totalData, err := msUker.repository.SearchJabatan(&request)
	if err != nil {
		msUker.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		msUker.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataJabatan {
		responses = append(responses, models.Jabatan{
			HILFM:   response.HILFM,
			HTEXT:   response.HTEXT,
			STELLTX: response.STELLTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// SearchPeserta implements MsUkerDefinition
func (serv MsUkerService) SearchPeserta(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPeserta, totalRows, totalData, err := serv.repository.SearchPeserta(&request)
	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR.String,
			SNAME:   response.SNAME.String,
			STEELTX: response.STEELTX.String,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetPekerjaByBranch implements MsUkerDefinition
func (serv MsUkerService) GetPekerjaByBranch(request models.BranchCodeInduk) (responses []models.MsPekerjaResponse, err error) {
	dataPeserta, err := serv.repository.GetPekerjaByBranch(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPekerjaResponse{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
			BRANCH:  response.BRANCH,
		})
	}

	return responses, err
}

func (serv MsUkerService) GetPekerjaByRegion(request *models.SearchPNByRegionReq) (responses []models.SearchPNByRegionRes, err error) {
	dataPekerja, err := serv.repository.GetPekerjaByRegion(request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataPekerja {
		responses = append(responses, models.SearchPNByRegionRes{
			PERNR: response.PERNR,
			SNAME: response.SNAME,
		})
	}

	return responses, err
}

// SearchPekerjaPerUker implements MsUkerDefinition
func (serv MsUkerService) SearchPekerjaPerUker(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	//  serv.repository.
	dataPeserta, totalRows, totalData, err := serv.repository.SearchPekerjaPerUker(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// SearchSigner implements MsUkerDefinition
func (serv MsUkerService) SearchSigner(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	//  serv.repository.
	dataPeserta, totalRows, totalData, err := serv.repository.SearchSigner(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// SearchRMC implements MsUkerDefinition
func (serv MsUkerService) SearchRMC(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	//  serv.repository.
	dataPeserta, totalRows, totalData, err := serv.repository.SearchRMC(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// SearchPelakuFraud implements MsUkerDefinition
func (serv MsUkerService) SearchPelakuFraud(request models.KeywordRequest) (responses []models.MsPelaku, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	//  serv.repository.
	dataPeserta, totalRows, totalData, err := serv.repository.SearchPelakuFraud(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPelaku{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// SearchBrcUrcPerRegion implements MsUkerDefinition
func (serv MsUkerService) SearchBrcUrcPerRegion(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	//  serv.repository.
	dataPeserta, totalRows, totalData, err := serv.repository.SearchBrcUrcPerRegion(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// ListingJabatanPerUker implements MsUkerDefinition.
func (msUker MsUkerService) ListingJabatanPerUker(request models.ListJabatanRequest) (responses []models.ListJabatanResponse, err error) {
	listingJabatan, err := msUker.repository.ListingJabatanPerUker(&request)

	if err != nil {
		msUker.logger.Zap.Error()
		return responses, err
	}

	for _, response := range listingJabatan {
		responses = append(responses, models.ListJabatanResponse{
			HILFM:  response.HILFM,
			HTEXT:  response.HTEXT,
			Jumlah: response.Jumlah,
		})
	}

	return responses, err
}

// GetPekerjaBranchByHILFM implements MsUkerDefinition.
func (msUker MsUkerService) GetPekerjaBranchByHILFM(request models.BranchByHilfmRequest) (responses []models.MsPeserta, err error) {
	listingPekerja, err := msUker.repository.GetPekerjaBranchByHILFM(request)

	if err != nil {
		msUker.logger.Zap.Error()
		return responses, err
	}

	for _, response := range listingPekerja {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	return responses, err
}

// SearchRRMHead implements MsUkerDefinition.
func (serv MsUkerService) SearchRRMHead(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	//  serv.repository.
	dataPeserta, totalRows, totalData, err := serv.repository.SearchRRMHead(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPeserta {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}

// SearchPekerjaOrd implements MsUkerDefinition.
func (serv MsUkerService) SearchPekerjaOrd(request models.KeywordRequest) (responses []models.MsPeserta, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	//  serv.repository.
	dataPekerja, totalRows, totalData, err := serv.repository.SearchPekerjaOrd(&request)

	if err != nil {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		serv.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPekerja {
		responses = append(responses, models.MsPeserta{
			PERNR:   response.PERNR,
			SNAME:   response.SNAME,
			STEELTX: response.STEELTX,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)

	return responses, pagination, err
}
