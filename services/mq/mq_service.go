package mq

import (
	"encoding/json"
	"fmt"
	"riskmanagement/lib"
	menuModels "riskmanagement/models/menu"
	models "riskmanagement/models/mq"

	menuRepo "riskmanagement/repository/menu"

	"github.com/google/uuid"

	"gitlab.com/golang-package-library/logger"

	auth "riskmanagement/services/auth"
)

var (
	timeNow = lib.GetTimeNow("timestime")
	UUID    = uuid.NewString()
)

type MQDefinition interface {
	// Type
	GetTypeCode(request models.TypeRequest) (response interface{}, err error)
	GetAll(request models.TypeRequest) (response interface{}, err error)
	GetAllSent(request models.TypeRequest) (response interface{}, err error)
	GetTypeList(request models.TypeRequest) (response interface{}, err error)
	GetTypeListAll(request models.TypeRequest) (response interface{}, err error)
	GetTypeListAktif(request models.TypeRequest) (response interface{}, err error)
	Filter(request models.TypeRequest) (response interface{}, err error)
	FilterSent(request models.TypeRequest) (response interface{}, err error)
	GetOne(request models.TypeRequest) (response interface{}, err error)
	Store(request models.TypeRequest) (response interface{}, err error)
	Update(request models.TypeRequest) (response interface{}, err error)
	Delete(request models.TypeRequest) (response interface{}, err error)
	SetStatusType(request models.TypeRequest) (response interface{}, err error)
	GetAllORDMember(request models.TypeRequest) (response interface{}, err error)
	GetDataQuestForm(request models.TypeRequest) (response interface{}, err error)
	SendQuest(request models.TypeRequest) (response interface{}, err error)
	ApproverUpdate(request models.ApproverRequest) (response interface{}, err error)
	GenerateQuestionnaireRequest(request models.GenerateQuestionnaireRequest) (response interface{}, err error)

	WeightTotal(request models.GenerateRequest) (response interface{}, err error)
	UpdatePartWeight(request models.PartWeightRequest) (response interface{}, err error)

	// Part
	GetPartCode(request models.PartRequest) (response interface{}, err error)
	GetAllPart(request models.PartRequest) (response interface{}, err error)
	FilterPart(request models.PartRequest) (response interface{}, err error)
	StorePart(request models.PartRequest) (response interface{}, err error)
	DeletePart(request models.PartRequest) (response interface{}, err error)
	UpdatePart(request models.PartRequest) (response interface{}, err error)
	GetOnePart(request models.PartRequest) (response interface{}, err error)

	//SubPart
	StoreSubPart(request models.SubPart) (response interface{}, err error)
	UpdateSubPart(request models.SubPartEditRequest) (response interface{}, err error)
	DeleteSubPart(request models.SubPartListRequest) (response interface{}, err error)
	GetSubPartList(request models.SubPartListRequest) (response interface{}, err error)
	GetSubPartKode(request models.SubPartListRequest) (response interface{}, err error)
	GetDetailSubPart(request models.SubPartListRequest) (response interface{}, err error)

	//Questionnaire
	StoreQuestion(request models.RequestQuestion) (response interface{}, err error)
	UpdateQuestion(request models.RequestQuestionUpdate) (response interface{}, err error)
	DeleteQuestion(request models.RequestQuestion) (response interface{}, err error)
	GetQuestionList(request models.QuestionListRequest) (response interface{}, err error)
	GetKodeQuestionnaire(request models.QuestionListRequest) (response interface{}, err error)
	GetDetailQuestionnaire(request models.QuestionListRequest) (response interface{}, err error)

	//Common
	CommonGetType(request models.CommonRequest) (response interface{}, err error)
	CommonGetPart(request models.CommonRequest) (response interface{}, err error)
	CommonGetSubPart(request models.CommonRequest) (response interface{}, err error)
	GetDetailType(request models.CommonRequest) (response interface{}, err error)
	GetApprovalResponse(request models.CommonRequest) (response interface{}, err error)

	//Menu
	GetAllMenu(request models.MenuRequest) (response interface{}, err error)
	GetAllMstMenu(request models.MenuRequest) (response interface{}, err error)

	// Linkcage
	StoreLinkcage(request models.LinkcageRequest) (response interface{}, err error)
	GetActive(request models.LinkcageRequest) (response interface{}, err error)
	GetAllLinkcage(request models.LinkcageRequest) (response interface{}, err error)
	GetOneLinkcage(request models.LinkcageRequest) (response interface{}, err error)
	SetStatus(request models.LinkcageRequest) (response interface{}, err error)
	UpdateLinkcage(request models.LinkcageRequest) (response interface{}, err error)
	DeleteLinkcage(request models.LinkcageRequest) (response interface{}, err error)

	// Response User
	GetResponseUserList(request models.RequestResponseUserList) (response interface{}, err error)
	GetResponseApprovalList(request models.RequestResponseApprovalList) (response interface{}, err error)
	StoreResponseUser(request models.RequestUserHistory) (response interface{}, err error)
	UpdateResponseUser(request models.RequestUserHistory) (response interface{}, err error)
	GenerateQuestWithAnswer(request models.GenerateRequest) (response interface{}, err error)
	ApproveResponseUser(request models.ApprovalUpdate) (response interface{}, err error)
	RejectResponseUser(request models.RejectedUpdate) (response interface{}, err error)
	GetRejectResponse(request models.RejectedUpdate) (response interface{}, err error)
	GetNilaiAkhir(request models.GenerateRequest) (response interface{}, err error)
	GenerateQuestPerPage(request models.GenerateRequest) (response interface{}, err error)
	GenerateQuestPerPagePreview(request models.GenerateRequest) (response interface{}, err error)
	ProcessGeneratePagination(request models.RequestUserHistory) (response interface{}, err error)
	CancelResponseUser(request models.UpdateResponseUserHistory) (response interface{}, err error)
	GetPartPagination(request models.GenerateRequest) (response interface{}, err error)
	GetPartPaginationDraft(request models.GenerateRequest) (response interface{}, err error)
	DisableQuestionByPart(request models.RequestPartid) (response interface{}, err error)
	GenerateQuestForApprover(request models.GenerateRequest) (response interface{}, err error)            // Perbaikan DAST
	GenerateQuestPerPagePreviewApprover(request models.GenerateRequest) (response interface{}, err error) // perbaikan DAST // untuk template questionnaire

	//report
	GetReportList(request models.RequestReportList) (response interface{}, err error)
	GenerateReportPerPage(request models.GenerateRequest) (response interface{}, err error)
	GetNamaRespondenList(request models.GenerateRequest) (response interface{}, err error)
	ResponseDownload(request models.ReportListQuery) (response interface{}, err error)
	GetSummary(request models.GenerateRequest) (response interface{}, err error)
}

type MQService struct {
	db         lib.Database
	dbRaw      lib.Databases
	logger     logger.Logger
	jwtService auth.JWTAuthService
	menuRepo   menuRepo.MenuDefinition
}

func NewMQService(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
	jwtService auth.JWTAuthService,
	menuRepo menuRepo.MenuDefinition,
) MQDefinition {
	return MQService{
		db:         db,
		dbRaw:      dbRaw,
		logger:     logger,
		jwtService: jwtService,
		menuRepo:   menuRepo,
	}
}

func (mq MQService) GetTypeCode(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getTypeCode"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetAll(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getAll"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetAllSent(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getAllSent"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) Filter(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/filter"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	fmt.Println("Url :", Url)
	fmt.Println("header :", headers)
	fmt.Println("request :", request)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	jsonStr, err := json.Marshal(response)
	if err != nil {
		mq.logger.Zap.Error("Terjadi kesalahan : ", err)
	}
	// // fmt.Printf("Tipe data: %T\n", response)
	var resp models.ResponseFilter
	if err = json.Unmarshal(jsonStr, &resp); err != nil {
		mq.logger.Zap.Error("Terjadi kesalahan => ", err)
		return
	}

	var responses []models.TypeResponse
	for _, item := range resp.Data {

		mstMenuRRM, err := mq.menuRepo.GetMstMenuRRM(item.MstMenu.IDParent)
		if err != nil {
			mq.logger.Zap.Error(err)
		}
		responses = append(responses, models.TypeResponse{
			ID:                  item.ID,
			MenuID:              item.MenuID,
			Code:                item.Code,
			Name:                item.Name,
			PERNR:               item.PERNR,
			SNAME:               item.SNAME,
			MstMenu:             item.MstMenu,
			MstMenuRRM:          models.MstMenuResponse(mstMenuRRM),
			ApproveDate:         item.ApproveDate,
			ApprovedType:        item.ApprovedType,
			QuestionnaireStatus: item.QuestionnaireStatus,
			ApprovalStatus:      item.ApprovalStatus,
			Status:              item.Status,
			CreatedAt:           item.CreatedAt,
		})
	}

	response = models.ResponseFilter{
		Data:       responses,
		Message:    resp.Message,
		Status:     resp.Status,
		Pagination: resp.Pagination,
	}

	return response, err
}

func (mq MQService) FilterSent(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/filterSent"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetTypeList(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getTypeList"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetTypeListAll(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getTypeListAll"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetTypeListAktif(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getTypeListAktif"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetOne(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getOne"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) Store(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/store"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	//cek apakah nama menu sama dengan menu existing

	if request.MenuType != "none" {
		checkMenu, err := mq.menuRepo.CheckMenuParentExist(request.Title)
		if err != nil {
			mq.logger.Zap.Error(err)
		}

		//jika kosong atau true
		if checkMenu {
			err = lib.MakeRequest("POST", Url, headers, request, &response)

			if err != nil {
				mq.logger.Zap.Error(err)
			}
		} else {
			response = models.ResponseStore{
				Data:    false,
				Message: "Nama menu sudah digunakan",
				Status:  "500",
			}
		}
	} else {
		err = lib.MakeRequest("POST", Url, headers, request, &response)

		if err != nil {
			mq.logger.Zap.Error(err)
		}
	}

	return response, err
}

func (mq MQService) Update(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/update"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	if request.MenuType != "none" {
		//cek apakah nama menu sama dengan menu existing
		checkMenu, err := mq.menuRepo.CheckMenuParentExist(request.Title)
		if err != nil {
			mq.logger.Zap.Error(err)
		}

		if checkMenu {
			err = lib.MakeRequest("POST", Url, headers, request, &response)

			if err != nil {
				mq.logger.Zap.Error(err)
			}
		} else {
			response = models.ResponseStore{
				Data:    false,
				Message: "Nama menu sudah digunakan",
				Status:  "500",
			}
		}
	} else {
		err = lib.MakeRequest("POST", Url, headers, request, &response)

		if err != nil {
			mq.logger.Zap.Error(err)
		}
	}

	return response, err
}

func (mq MQService) Delete(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/delete"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) SetStatusType(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/setStatus"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	//set status menu

	menuStatus := ""
	if request.Status == "Aktif" {
		menuStatus = "1"
	} else if request.Status == "Non Aktif" {
		menuStatus = "0"
	}

	_, err = mq.menuRepo.SetStatus(&menuModels.MstMenu{
		IDMenu: request.MenuID,
		Status: menuStatus,
	})

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetAllORDMember(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getAllORDMember"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetDataQuestForm(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/getDataQuestForm"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) SendQuest(request models.TypeRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/sendQuest"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) ApproverUpdate(request models.ApproverRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/approverUpdate"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GenerateQuestionnaireRequest(request models.GenerateQuestionnaireRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/generateQuestionnaireRequest"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) WeightTotal(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/weightTotal"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) UpdatePartWeight(request models.PartWeightRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/updatePartWeight"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetPartCode(request models.PartRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "part/getPartCode"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetAllPart(request models.PartRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "part/getAll"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) FilterPart(request models.PartRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "part/filter"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) StorePart(request models.PartRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "part/store"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) DeletePart(request models.PartRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "part/delete"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) UpdatePart(request models.PartRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "part/update"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetOnePart(request models.PartRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "part/getOne"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) StoreSubPart(request models.SubPart) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "subpart/store"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) UpdateSubPart(request models.SubPartEditRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "subpart/update"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) DeleteSubPart(request models.SubPartListRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "subpart/delete"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetSubPartList(request models.SubPartListRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "subpart/list"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetSubPartKode(request models.SubPartListRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "subpart/kode"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetDetailSubPart(request models.SubPartListRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "subpart/detail"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) StoreQuestion(request models.RequestQuestion) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "questionnaire/store"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) UpdateQuestion(request models.RequestQuestionUpdate) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "questionnaire/update"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) DeleteQuestion(request models.RequestQuestion) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "questionnaire/delete"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetQuestionList(request models.QuestionListRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "questionnaire/list"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetKodeQuestionnaire(request models.QuestionListRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "questionnaire/kode"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetDetailQuestionnaire(request models.QuestionListRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "questionnaire/detail"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) CommonGetType(request models.CommonRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "common/type"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) CommonGetPart(request models.CommonRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "common/part"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) CommonGetSubPart(request models.CommonRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "common/sub-part"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetDetailType(request models.CommonRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "common/detail-type"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetApprovalResponse(request models.CommonRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "common/pn-approval"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetAllMenu(request models.MenuRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "menu/getAll"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetAllMstMenu(request models.MenuRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "menu/getAllMstMenu"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) StoreLinkcage(request models.LinkcageRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "linkcage/store"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetActive(request models.LinkcageRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "linkcage/getActive"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetAllLinkcage(request models.LinkcageRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "linkcage/getAll"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetOneLinkcage(request models.LinkcageRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "linkcage/getOne"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) SetStatus(request models.LinkcageRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "linkcage/setStatus"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) UpdateLinkcage(request models.LinkcageRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "linkcage/update"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) DeleteLinkcage(request models.LinkcageRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "linkcage/delete"

	token := mq.jwtService.CreateTokenByPN(request.PERNR)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

// =============================Response User=========================
func (mq MQService) GetResponseUserList(request models.RequestResponseUserList) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/list"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetResponseApprovalList(request models.RequestResponseApprovalList) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/approval-list"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) StoreResponseUser(request models.RequestUserHistory) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/store"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GenerateQuestWithAnswer(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/generate"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) UpdateResponseUser(request models.RequestUserHistory) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/update"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) ApproveResponseUser(request models.ApprovalUpdate) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/approve"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) RejectResponseUser(request models.RejectedUpdate) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/reject"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetRejectResponse(request models.RejectedUpdate) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/reject-list"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetNilaiAkhir(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/nilai-akhir"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GenerateQuestPerPage(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/generate-perpage"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

// Perbaikan DAST
func (mq MQService) GenerateQuestForApprover(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/generate-perpage-approval"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GenerateQuestPerPagePreview(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/generate-perpage"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

// perbaikan DAST
func (mq MQService) GenerateQuestPerPagePreviewApprover(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "type/generate-perpage-approval"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) ProcessGeneratePagination(request models.RequestUserHistory) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/pagination-process"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) CancelResponseUser(request models.UpdateResponseUserHistory) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/cancel"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetPartPagination(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/get-part-pagination"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetPartPaginationDraft(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/get-part-draft-pagination"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) DisableQuestionByPart(request models.RequestPartid) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "response-user/disable-quest"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

// ============================ Report ===================================
func (mq MQService) GetReportList(request models.RequestReportList) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "report/list"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GenerateReportPerPage(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "report/generate"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetNamaRespondenList(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "report/responden"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) ResponseDownload(request models.ReportListQuery) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")

	if err != nil {
		return response, err
	}

	Url := baseUrl + "report/downloadv2"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	fmt.Println(response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}

func (mq MQService) GetSummary(request models.GenerateRequest) (response interface{}, err error) {
	baseUrl, err := lib.GetVarEnv("MQUrl")
	if err != nil {
		return response, err
	}

	Url := baseUrl + "report/getSummary"

	token := mq.jwtService.CreateTokenByPN(request.Pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	err = lib.MakeRequest("POST", Url, headers, request, &response)

	fmt.Println(response)

	if err != nil {
		mq.logger.Zap.Error(err)
	}

	return response, err
}
