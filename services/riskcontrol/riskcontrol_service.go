package riskcontrol

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"regexp"
	"riskmanagement/dto"
	"riskmanagement/lib"
	modelsOrganisasi "riskmanagement/models/organisasi"
	models "riskmanagement/models/riskcontrol"
	modelsControlAttribute "riskmanagement/models/riskcontrolattribute"
	repoOrganisasi "riskmanagement/repository/organisasi"
	repository "riskmanagement/repository/riskcontrol"
	repoIssue "riskmanagement/repository/riskissue"
	"riskmanagement/services/arlords"
	jwt "riskmanagement/services/auth"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
	"gitlab.com/golang-package-library/logger"
)

type RiskControlDefinition interface {
	GetAll() (responses []models.RiskControlResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.RiskControlResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.RiskControlResponse, err error)
	Store(request *models.RiskControlRequest) (err error)
	Update(request *models.RiskControlRequest) (err error)
	Delete(id int64) (err error)
	GetKodeRiskControl() (responses []models.KodeRiskControl, err error)
	GenCode() (string, error)
	UpdateStatus(id int64) (stats bool, err error)
	Preview(pernr string, data [][]string) (dto.PreviewFileImport[[10]string], error)
	Template() ([]byte, string, error)
	ImportData(pernr string, data [][]string) error
	Download(pernr, format string) (blob []byte, name string, err error)
	SearchRiskControlByIssue(request models.KeywordRequest) (response []models.RiskControlResponses, pagination lib.Pagination, err error)
}

type RiskControlService struct {
	logger        logger.Logger
	repository    repository.RiskControlDefinition
	jwtService    jwt.JWTAuthService
	arlodsService arlords.ArlordsServiceDefinition
	orgRepo       repoOrganisasi.OrganisasiDefinition
	repoIssue     repoIssue.MapControlDefinition
	db            lib.Database
}

func NewRiskControService(
	db lib.Database,
	logger logger.Logger,
	arlodsService arlords.ArlordsServiceDefinition,
	repository repository.RiskControlDefinition,
	orgRepo repoOrganisasi.OrganisasiDefinition,
	repoIssue repoIssue.MapControlDefinition,
	jwtService jwt.JWTAuthService,
) RiskControlDefinition {
	return RiskControlService{
		db:            db,
		logger:        logger,
		arlodsService: arlodsService,
		repository:    repository,
		orgRepo:       orgRepo,
		repoIssue:     repoIssue,
		jwtService:    jwtService,
	}
}

// Delete implements RiskControlDefinition
func (riskControl RiskControlService) Delete(id int64) (err error) {
	maps, _, err := riskControl.repoIssue.GetWithPagination(int(id), models.Paginate{
		Order:  "ID",
		Sort:   "DESC",
		Offset: 1,
		Limit:  10,
		Page:   0,
	})

	if err != nil {
		riskControl.logger.Zap.Error(err)
		return err
	}

	if len(maps) > 0 {
		return fmt.Errorf("data risk control masih sudah termapping ke risk event dan tidak bisa dihapus")
	}

	return riskControl.repository.Delete(id)
}

// GetAllWithPaginate implements RiskControlDefinition
func (rc RiskControlService) GetAllWithPaginate(request models.Paginate) (responses []models.RiskControlResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort
	request.Limit = limit

	dataPgs, totalData, totalRows, err := rc.repository.GetAllWithPaginate(&request)
	if err != nil {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	responses = dataPgs

	pagination = lib.SetPaginationResponse(page, limit, int(totalRows), int(totalData))
	return responses, pagination, err
}

// GetAll implements RiskControlDefinition
func (riskControl RiskControlService) GetAll() (responses []models.RiskControlResponse, err error) {
	return riskControl.repository.GetAll()
}

// GetOne implements RiskControlDefinition
func (riskControl RiskControlService) GetOne(id int64) (responses models.RiskControlResponse, err error) {
	// return riskControl.repository.GetOne(id)
	data, err := riskControl.repository.GetOne(id)

	return data, err
}

// Store implements RiskControlDefinition
func (riskControl RiskControlService) Store(request *models.RiskControlRequest) (err error) {
	timeNow := lib.GetTimeNow("timestime")
	if request.OwnerGroup != "" && request.Owner != "" {
		err = riskControl.ValidationOwner(request)
		if err != nil {
			riskControl.logger.Zap.Error(err)
			return fmt.Errorf("owner is invalid: %s", err)
		}
	}
	transform := request.ParseRequest()
	if transform.Kode == "" {
		code, err := riskControl.repository.GenLastCode()
		if err != nil {
			riskControl.logger.Zap.Error(err)
			return err
		}

		transform.Kode = code
	}

	transform.Status = true
	transform.CreatedAt = &timeNow

	status, err := riskControl.repository.Store(&transform)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements RiskControlDefinition
func (riskControl RiskControlService) Update(request *models.RiskControlRequest) (err error) {
	timeNow := lib.GetTimeNow("timestime")
	if request.OwnerGroup != "" && request.Owner != "" {
		err = riskControl.ValidationOwner(request)
		if err != nil {
			riskControl.logger.Zap.Error(err)
			return fmt.Errorf("owner is invalid: %s", err)
		}
	}

	exists, err := riskControl.repository.GetOne(request.ID)
	if err != nil {
		riskControl.logger.Zap.Error(err)
		return err
	}

	transform := request.ParseRequest()

	transform.Status = exists.Status
	transform.CreatedAt = exists.CreatedAt
	transform.UpdatedAt = &timeNow
	status, err := riskControl.repository.Update(&transform)
	if !status || err != nil {
		return err
	}

	return nil
}

// GetKodeRiskControl implements RiskControlDefinition
func (riskControl RiskControlService) GetKodeRiskControl() (responses []models.KodeRiskControl, err error) {
	dataRC, err := riskControl.repository.GetKodeRiskControl()

	if err != nil {
		riskControl.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range dataRC {
		responses = append(responses, models.KodeRiskControl{
			KodeRiskControl: response.KodeRiskControl,
		})
	}

	return responses, err
}

// SearchRiskControlByIssue implements RiskControlDefinition
func (rc RiskControlService) SearchRiskControlByIssue(request models.KeywordRequest) (responses []models.RiskControlResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataControl, totalRows, totalData, err := rc.repository.SearchRiskControlByIssue(&request)
	if err != nil {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	if totalData < 0 {
		rc.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataControl {
		responses = append(responses, models.RiskControlResponses{
			ID:          response.ID,
			Kode:        response.Kode,
			RiskControl: response.RiskControl,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

func (rc RiskControlService) UpdateStatus(id int64) (stats bool, err error) {
	var (
		status bool = true
	)
	data, err := rc.repository.GetOne(id)
	if err != nil {
		return status, err
	}

	if data.Status {
		status = false
	}

	err = rc.repository.UpdateStatus(id, status)

	return status, err
}

func (rc RiskControlService) ValidationOwner(request *models.RiskControlRequest) error {
	switch request.OwnerGroup {
	case "jabatan":
		return rc.ValidateJabatan(request.Owner)
	case "departemen":
		return rc.ValidateDepartemen(request.Owner)
	default:
		return fmt.Errorf("invalid owner group")
	}

}

func (rc RiskControlService) ValidateJabatan(jabatanID string) error {
	data, err := rc.orgRepo.GetHilfm(modelsOrganisasi.JabatanRequest{
		Keyword: jabatanID,
		Limit:   0,
		Offset:  0,
	})

	if err != nil {
		rc.logger.Zap.Error(`Error checking data jabatan`, err)
		return err
	}

	if len(data) > 0 {
		return nil
	}

	return fmt.Errorf("data jabatan not found")
}

func (rc RiskControlService) ValidateDepartemen(departemenID string) error {
	data, err := rc.orgRepo.GetOrgUnit(modelsOrganisasi.DepartmentRequest{
		Keyword: departemenID,
		Limit:   0,
		Offset:  0,
	})

	if err != nil {
		rc.logger.Zap.Error(`Error checking data jabatan`, err)
		return err
	}

	if len(data) > 0 {
		return nil
	}

	return fmt.Errorf("data departemen not found")
}

func (rc RiskControlService) GenCode() (string, error) {
	code, err := rc.repository.GenLastCode()
	if err != nil {
		rc.logger.Zap.Error(err)
		return "", err
	}

	return code, nil
}

func (rc RiskControlService) Preview(pernr string, data [][]string) (dto.PreviewFileImport[[10]string], error) {
	jabatan, err := rc.orgRepo.GetHilfm(modelsOrganisasi.JabatanRequest{})
	if err != nil {
		rc.logger.Zap.Error("Errored to query jabatan: %s", err)
		return dto.PreviewFileImport[[10]string]{}, err
	}

	departemen, err := rc.orgRepo.GetOrgUnit(modelsOrganisasi.DepartmentRequest{})
	if err != nil {
		rc.logger.Zap.Error("Errored to query departemen: %s", err)
		return dto.PreviewFileImport[[10]string]{}, err
	}

	control, err := rc.GetAll()
	if err != nil {
		rc.logger.Zap.Error("Errored to query risk control: %s", err)
		return dto.PreviewFileImport[[10]string]{}, err
	}

	previewFile := dto.PreviewFileImport[[10]string]{}
	body := []dto.PreviewFile[[10]string]{}
	cacheControlAttribute := make(map[string]bool)

	for index, row := range data {
		if index == 0 {
			if len(row) < 10 {
				return dto.PreviewFileImport[[10]string]{}, fmt.Errorf("invalid header format risk control")
			}

			previewFile.Header = [10]string{
				row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9],
			}
			continue
		}

		var col [10]string
		validation := ""

		var (
			exists    bool = false
			jabExists bool = false
			depExists bool = false
		)
		riskControlCode := row[0]

		for _, c := range control {
			if strings.EqualFold(c.RiskControl, riskControlCode) {
				exists = true
				break
			}
		}

		if strings.ToLower(row[6]) == "jabatan" {
			if row[7] == "" && row[6] == "" {
				jabExists = true
			}

			if !IsValidCodeName(row[7]) {
				validation += fmt.Sprintf("Owner invalid format, format must be <code> - <name>: %s; ", row[7])
			}

			for _, c := range jabatan {
				jab := lib.ParseStringToArray(row[7], "-")
				if strings.EqualFold(c.Hilfm, jab[0]) {
					jabExists = true
					break
				}
			}

			if !jabExists {
				validation += fmt.Sprintf("Owner tidak terdaftar: %s; ", row[7])
			}
		}

		if strings.ToLower(row[6]) == "departemen" {
			if row[7] == "" && row[6] == "" {
				depExists = true
			}

			if !IsValidCodeName(row[7]) {
				validation += fmt.Sprintf("Owner invalid format, format must be <code> - <name>: %s; ", row[7])
			}

			for _, c := range departemen {
				jab := lib.ParseStringToArray(row[7], "-")
				if strings.EqualFold(c.Orgeh, jab[0]) {
					depExists = true
					break
				}
			}

			if !depExists {
				validation += fmt.Sprintf("Owner tidak terdaftar: %s; ", row[7])
			}
		}

		if exists {
			validation += fmt.Sprintf("Nama risk control: %s sudah terdaftar; ", row[0])
		}

		parseCodeControAttribute := lib.ParseStringToArray(row[9], ";")

		// Step 1: cari code yang belum ada dalam cache
		var needFetch []string
		for _, code := range parseCodeControAttribute {
			if _, ok := cacheControlAttribute[code]; !ok {
				needFetch = append(needFetch, code)
			}
		}

		// Step 2: Jika ada yg belum di-cache → fetch API
		if len(needFetch) > 0 {
			controlAttribute, err := rc.arlodsService.GetControlAttribute("", needFetch)
			if err != nil {
				rc.logger.Zap.Error("failed to request data control attribute %s", err)
				continue
			}

			rc.logger.Zap.Debug(needFetch)
			rc.logger.Zap.Debug(controlAttribute)
			// Code yang valid → set true
			for _, item := range controlAttribute.Data.List {
				cacheControlAttribute[item.Code] = true
			}
			rc.logger.Zap.Debug(cacheControlAttribute)

			// Code yang tidak muncul dalam response → invalid
			for _, code := range needFetch {
				if _, ok := cacheControlAttribute[code]; !ok {
					cacheControlAttribute[code] = false
				}
			}
		}

		// Step 3: Validasi dari cache (super cepat)
		rc.logger.Zap.Debug(cacheControlAttribute)
		for _, code := range parseCodeControAttribute {
			if !cacheControlAttribute[code] {
				validation += fmt.Sprintf("Code control attribute %s Tidak terdaftar atau masih non active; ", code)
			}
		}

		for i := range 10 {
			if i < len(row) {
				col[i] = row[i]
			}
		}

		body = append(body, dto.PreviewFile[[10]string]{
			PerRow:     col,
			Validation: validation,
		})
	}

	previewFile.Body = body

	return previewFile, nil

}

func (rc RiskControlService) Template() ([]byte, string, error) {
	f := excelize.NewFile()
	sheet := "Template"

	f.SetSheetName("Sheet1", sheet)

	sheetIndex, err := f.GetSheetIndex(sheet)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get sheet index: %w", err)
	}
	f.SetActiveSheet(sheetIndex)

	// header
	headers := []string{
		"Risk Control",
		"Control Type",
		"Nature",
		"Key Control",
		"Deskripsi",
		"Control Owner Level",
		"Owner Group",
		"Control Owner",
		"Control Document",
		"Code Control Attribute",
		"Control Attribute",
	}

	// tulis header
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			return nil, "", fmt.Errorf("failed to set cell: %w", err)
		}
	}

	// optional: set lebar kolom
	for i := 1; i <= len(headers); i++ {
		col, _ := excelize.ColumnNumberToName(i)
		f.SetColWidth(sheet, col, col, 25)
	}

	// contoh data
	exampleData := [][]string{
		{"RC-001", "Operational", "Financial", "Key1", "Contoh deskripsi control", "Head Office", "Jabatan", "code - name", "Doc-001", "abc;def", "<Name1>; <Name2>"},
		{"RC-002", "Compliance", "Legal", "Key2", "Deskripsi control kedua", "Regional Office", "Departemen", "code - name", "Doc-002", "ghi;jkl", "<Name3>; <Name4>"},
	}

	// tulis data
	for rowIndex, row := range exampleData {
		for colIndex, val := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2) // +2 karena header di baris 1
			if err := f.SetCellValue(sheet, cell, val); err != nil {
				return nil, "", fmt.Errorf("failed to set example data: %w", err)
			}
		}
	}

	// simpan ke buffer
	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", fmt.Errorf("failed to write excel: %w", err)
	}

	return buf.Bytes(), "risk_control_template.xlsx", nil
}

func (rc RiskControlService) ImportData(pernr string, data [][]string) error {
	timeNow := lib.GetTimeNow("timestime")
	newRecord := make([]models.RiskControlRequest, 0)
	attribute := make(map[string][]string)

	jabatan, err := rc.orgRepo.GetHilfm(modelsOrganisasi.JabatanRequest{})
	if err != nil {
		rc.logger.Zap.Error("Errored to query jabatan: %s", err)
		return err
	}

	departemen, err := rc.orgRepo.GetOrgUnit(modelsOrganisasi.DepartmentRequest{})
	if err != nil {
		rc.logger.Zap.Error("Errored to query departemen: %s", err)
		return err
	}

	control, err := rc.GetAll()
	if err != nil {
		rc.logger.Zap.Error("Errored to query risk control: %s", err)
		return err
	}

	genCode, err := rc.repository.GenLastCode()
	if err != nil {
		rc.logger.Zap.Error("Errored to query generate code: %s", err)
		return err
	}

	cacheControlAttribute := make(map[string]bool)
	code := ""
	for i, v := range data {
		if i == 0 {
			continue
		}

		var (
			exists    bool = false
			jabExists bool = false
		)

		parseCodeControAttribute := lib.ParseStringToArray(v[9], ";")
		for _, c := range control {
			if strings.EqualFold(c.RiskControl, v[0]) {
				exists = true
				break
			}
		}

		if strings.ToLower(v[6]) == "jabatan" {
			if v[7] == "" && v[6] == "" {
				jabExists = true
			}

			if IsValidCodeName(v[7]) {
				continue
			}
			for _, c := range jabatan {
				if v[7] == "" {
					jabExists = true
					continue
				}
				jab := lib.ParseStringToArray(v[7], "-")
				if strings.EqualFold(c.Hilfm, jab[0]) {
					jabExists = true
					break
				}
			}
		}

		if strings.ToLower(v[6]) == "departemen" {
			if v[7] == "" && v[6] == "" {
				jabExists = true
			}
			if !IsValidCodeName(v[7]) {
				continue
			}
			for _, c := range departemen {
				if v[7] == "" {
					jabExists = true
				}
				jab := lib.ParseStringToArray(v[7], "-")
				if strings.EqualFold(c.Orgeh, jab[0]) {
					jabExists = true
					break
				}
			}
		}

		if !exists && jabExists {
			if i == 1 {
				code = genCode
			} else {
				code, err = GenerateRunningCode(code, i)
				if err != nil {
					rc.logger.Zap.Error("invalid generate code %s", err)
					return err
				}
			}

			newRecord = append(newRecord, models.RiskControlRequest{
				Kode:        code,
				RiskControl: v[0],
				ControlType: v[1],
				Nature:      v[2],
				KeyControl:  v[3],
				Deskripsi:   v[4],
				Status:      true,
				OwnerLvl:    v[5],
				OwnerGroup:  v[6],
				Owner:       v[7],
				Document:    v[8],
				CreatedAt:   &timeNow,
			})
		}

		// Step 1: cari code yang belum ada dalam cache
		var needFetch []string
		for _, codes := range parseCodeControAttribute {
			if _, ok := cacheControlAttribute[codes]; !ok {
				needFetch = append(needFetch, codes)
			}
		}

		// Step 2: Jika ada yg belum di-cache → fetch API
		if len(needFetch) > 0 {
			controlAttribute, err := rc.arlodsService.GetControlAttribute("", needFetch)
			if err != nil {
				rc.logger.Zap.Error("failed to request data control attribute %s", err)
				continue
			}

			// Code yang valid → set true
			for _, item := range controlAttribute.Data.List {
				cacheControlAttribute[item.Code] = true
			}

			// Code yang tidak muncul dalam response → invalid
			for _, code := range needFetch {
				if _, ok := cacheControlAttribute[code]; !ok {
					cacheControlAttribute[code] = false
				}
			}
		}

		// Step 3: Validasi dari cache (super cepat)
		var validCodes []string
		for _, codeA := range parseCodeControAttribute {
			if !cacheControlAttribute[codeA] {
				continue
			}

			validCodes = append(validCodes, codeA)
		}

		exists = false
		for _, c := range control {
			if strings.EqualFold(c.RiskControl, v[0]) {
				attribute[c.Kode] = validCodes
				exists = true
				break
			}
		}

		if !exists {
			attribute[code] = validCodes
		}

	}

	tx := rc.db.DB.Begin()

	err = rc.repository.BulkCreateRiskControl(newRecord, tx)
	if err != nil {
		tx.Rollback()
		rc.logger.Zap.Error("cannot create risk control data: %s ", err)
		return err
	}

	err = rc.RequestAttributeStore(pernr, attribute)
	if err != nil {
		tx.Rollback()
		rc.logger.Zap.Error(err)
		return err
	}

	tx.Commit()
	return nil
}

func (rc RiskControlService) RequestAttributeStore(pernr string, data map[string][]string) error {
	attrReqest := make([]models.RiskControlAttributeRequest, 0)
	baseUrl, err := lib.GetVarEnv("ArlordsUrl")
	if err != nil {
		return fmt.Errorf("errored when got url arlods: %s", err)
	}

	url := baseUrl + "/control/bulk-create-attribute"

	authToken := rc.jwtService.CreateArlordsToken(pernr)

	headers := map[string]string{
		"Authorization": "Bearer " + authToken,
		"Content-Type":  "application/json",
		"pernr":         pernr,
	}

	rc.logger.Zap.Debug(data)
	for i, v := range data {
		attrReqest = append(attrReqest, models.RiskControlAttributeRequest{
			ControlID: i,
			Attribute: v,
		})
	}

	var response modelsControlAttribute.HttpResResponse

	requestBody := models.RiskControlAttributeRequestBody{
		Data: attrReqest,
	}

	err = lib.MakeRequest("POST", url, headers, requestBody, &response)
	if err != nil {
		rc.logger.Zap.Error("Error when request to save mapping cause: %s", err)
		return err
	}

	return nil
}

func (rc RiskControlService) Download(pernr, format string) (blob []byte, name string, err error) {
	data, err := rc.repository.GetAll()
	if err != nil {
		rc.logger.Zap.Error("Errored when try to query: %s", err)
		return nil, "", err
	}

	if len(data) == 0 {
		return nil, "", nil
	}

	switch format {
	case "csv":
		return rc.exportCsv(pernr, data)
	case "xlsx":
		return rc.exportExcel(pernr, data)
	case "pdf":
		return rc.exportPDF(pernr, data)
	default:
		return nil, "", fmt.Errorf("unsupported format export file")
	}
}

func OwnerLvl(lvl string) string {
	list := map[string]string{
		"head office":     "ho",
		"regional office": "ro",
		"branch office":   "bo",
		"kcp":             "kcp",
		"unit kerja":      "uk",
	}

	lvlInfo, ok := list[strings.ToLower(lvl)]
	if !ok {
		return fmt.Sprintf("invalid lvl of owner: %s", lvl)
	}

	return lvlInfo
}

func GenerateRunningCode(baseCode string, index int) (string, error) {
	parts := strings.Split(baseCode, "-")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid code format")
	}

	prefix := parts[0]
	numberStr := parts[1]
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return "", err
	}

	nextNumber := number + index
	newCode := fmt.Sprintf("%s-%05d", prefix, nextNumber)
	return newCode, nil
}

func (rc RiskControlService) exportPDF(pernr string, data []models.RiskControlResponse) (blob []byte, name string, err error) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetAutoPageBreak(false, 10)
	pdf.SetMargins(10, 10, 10) // margin kiri, atas, kanan
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, "Risk Control Report", "", 1, "C", false, 0, "")

	headers := []string{"Control", "Type", "Nature", "Key", "Description", "Owner Level", "Owner Group", "Owner", "Document", "Code Control Attribute", "Name Control Attribute"}
	colWidths := []float64{
		35, 20, 20, 10,
		30, 20, 20, 25, 20,
		30, 30,
	}

	printHeader := func() {
		pdf.SetFillColor(200, 200, 200)
		pdf.SetFont("Arial", "B", 10)

		lineHeight := 5.0

		// Hitung tinggi maksimum header (berdasarkan wrapping)
		maxHeight := 0.0
		for i, h := range headers {
			lines := pdf.SplitLines([]byte(h), colWidths[i])
			hh := float64(len(lines)) * lineHeight
			if hh > maxHeight {
				maxHeight = hh
			}
		}

		xStart := pdf.GetX()
		yStart := pdf.GetY()

		// Cetak setiap header cell
		for i, h := range headers {
			x := pdf.GetX()
			y := pdf.GetY()

			// Gambar border kotak
			pdf.Rect(x, y, colWidths[i], maxHeight, "DF") // DF = fill + border

			// Cetak text dengan wrapping di tengah vertikal
			lines := pdf.SplitLines([]byte(h), colWidths[i])
			textHeight := float64(len(lines)) * lineHeight
			yOffset := (maxHeight - textHeight) / 2

			pdf.SetXY(x, y+yOffset)
			pdf.MultiCell(colWidths[i], lineHeight, h, "", "C", false)
			pdf.SetXY(x+colWidths[i], yStart)
		}

		pdf.SetXY(xStart, yStart+maxHeight)
		pdf.SetFont("Arial", "", 9)
	}

	printHeader()

	_, pageHeight := pdf.GetPageSize()
	marginBottom := 15.0

	getRowHeight := func(row []string) float64 {
		maxHeight := 0.0
		lineHeight := 5.0
		for i, txt := range row {
			lines := pdf.SplitLines([]byte(txt), colWidths[i])
			h := float64(len(lines)) * lineHeight
			if h > maxHeight {
				maxHeight = h
			}
		}
		return maxHeight
	}

	for _, v := range data {
		attribute, err := rc.arlodsService.GetMappingControlAttribute(pernr, int(v.ID))
		if err != nil {
			continue
		}

		code, name := BuildMaping(attribute.Data)
		row := []string{
			v.RiskControl,
			v.ControlType,
			v.Nature,
			v.KeyControl,
			v.Deskripsi,
			ReverseOwnerLvl(v.OwnerLvl),
			v.OwnerGroup,
			v.Owner,
			v.Document,
			code,
			name,
		}

		rowHeight := getRowHeight(row)
		xStart := pdf.GetX()
		yStart := pdf.GetY()

		if yStart+rowHeight+marginBottom > pageHeight {
			pdf.AddPage()
			printHeader()
			xStart = pdf.GetX()
			yStart = pdf.GetY()
		}

		for i, txt := range row {
			x := pdf.GetX()
			y := pdf.GetY()

			pdf.Rect(x, y, colWidths[i], rowHeight, "D")
			pdf.MultiCell(colWidths[i], 5, txt, "", "L", false)
			pdf.SetXY(x+colWidths[i], yStart)
		}
		pdf.SetXY(xStart, yStart+rowHeight)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate PDF: %w", err)
	}

	fileName := fmt.Sprintf("risk_control_%s.pdf", time.Now().Format("20060102_150405"))

	return buf.Bytes(), fileName, nil
}

func (rc RiskControlService) exportExcel(pernr string, data []models.RiskControlResponse) (blob []byte, name string, err error) {
	f := excelize.NewFile()
	sheet := "risk-control"

	f.SetSheetName("Sheet1", sheet)
	headers := []string{"Control", "Type", "Nature", "Key", "Description", "Owner Level", "Owner Group", "Owner", "Document", "Code Control Attribute", "Name Control Attribute"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		if err := f.SetCellValue(sheet, cell, h); err != nil {
			return nil, "", fmt.Errorf("failed to set cell: %w", err)
		}
	}

	for idx, v := range data {
		attribute, err := rc.arlodsService.GetMappingControlAttribute(pernr, int(v.ID))
		if err != nil {
			continue
		}

		code, name := BuildMaping(attribute.Data)

		f.SetCellValue("risk-control", fmt.Sprintf("A%d", idx+2), v.RiskControl)
		f.SetCellValue("risk-control", fmt.Sprintf("B%d", idx+2), v.ControlType)
		f.SetCellValue("risk-control", fmt.Sprintf("C%d", idx+2), v.Nature)
		f.SetCellValue("risk-control", fmt.Sprintf("D%d", idx+2), v.KeyControl)
		f.SetCellValue("risk-control", fmt.Sprintf("E%d", idx+2), v.Deskripsi)
		f.SetCellValue("risk-control", fmt.Sprintf("F%d", idx+2), ReverseOwnerLvl(v.OwnerLvl))
		f.SetCellValue("risk-control", fmt.Sprintf("G%d", idx+2), v.OwnerGroup)
		f.SetCellValue("risk-control", fmt.Sprintf("H%d", idx+2), v.Owner)
		f.SetCellValue("risk-control", fmt.Sprintf("I%d", idx+2), v.Document)
		f.SetCellValue("risk-control", fmt.Sprintf("J%d", idx+2), code)
		f.SetCellValue("risk-control", fmt.Sprintf("K%d", idx+2), name)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, "", fmt.Errorf("failed to write excel file: %w", err)
	}

	fileName := fmt.Sprintf("risk_control_%s.xlsx", time.Now().Format("20060102_150405"))

	return buf.Bytes(), fileName, nil

}

func (rc RiskControlService) exportCsv(pernr string, data []models.RiskControlResponse) (blob []byte, name string, err error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	headers := []string{"Control", "Type", "Nature", "Key", "Description", "Owner Level", "Owner Group", "Owner", "Document", "Code Control Attribute", "Name Control Attribute"}
	if err := writer.Write(headers); err != nil {
		return nil, "", fmt.Errorf("failed to write csv header: %w", err)
	}

	for _, v := range data {
		attribute, err := rc.arlodsService.GetMappingControlAttribute(pernr, int(v.ID))
		if err != nil {
			continue
		}

		code, name := BuildMaping(attribute.Data)

		row := []string{
			v.RiskControl,
			v.ControlType,
			v.Nature,
			v.KeyControl,
			v.Deskripsi,
			ReverseOwnerLvl(v.OwnerLvl),
			v.OwnerGroup,
			v.Owner,
			v.Document,
			code,
			name,
		}

		if err := writer.Write(row); err != nil {
			return nil, "", fmt.Errorf("failed to write csv row: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, "", fmt.Errorf("failed to flush csv data: %w", err)
	}

	fileName := fmt.Sprintf("risk_control_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil

}

func ReverseOwnerLvl(lvl string) string {
	list := map[string]string{
		"ho":  "head office",
		"ro":  "regional office",
		"bo":  "branch office",
		"kcp": "kcp",
		"uk":  "unit kerja",
	}

	lvlInfo, ok := list[strings.ToLower(lvl)]
	if !ok {
		return fmt.Sprintf("invalid reverse lvl of owner: %s", lvl)
	}

	return lvlInfo
}

func BuildMaping(data []dto.ListAttributeMap) (string, string) {
	var (
		code []string
		name []string
	)
	for _, v := range data {
		code = append(code, v.Code)
		name = append(name, v.Name)
	}

	return strings.Join(code, "; "), strings.Join(name, "; ")
}

func IsValidCodeName(s string) bool {
	// format: CODE - NAME
	// CODE  : huruf/angka/strip
	// NAME  : huruf, spasi, angka
	pattern := `^[A-Za-z0-9\-]+ - [A-Za-z0-9 ]+$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
