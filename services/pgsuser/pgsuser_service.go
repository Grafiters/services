package pgsuser

import (
	"database/sql"
	"fmt"

	"riskmanagement/lib"
	// modelsMU "riskmanagement/models/managementuser"
	// modelRole "riskmanagement/models/mstrole"
	models "riskmanagement/models/pgsuser"
	managementUser "riskmanagement/repository/managementuser"
	mstRole "riskmanagement/repository/mstrole"
	pgsUser "riskmanagement/repository/pgsuser"
	services "riskmanagement/services/auth"

	"github.com/google/uuid"

	"gitlab.com/golang-package-library/logger"
)

var (
	UUID = uuid.NewString()
)

type PgsUserDefinition interface {
	GetAll(makerID string) (responses []models.PgsUserResponses, err error)
	GetOne(id int64) (responses models.PgsUserResponseOne, status bool, err error)
	Store(request models.PgsUserRequest) (responses bool, err error)
	Update(request *models.PgsUserRequestUpdate) (responses bool, err error)
	GetPgsApproval(request models.Paginate) (responses []models.PgsApprovalResponse, pagination lib.Pagination, err error)
	ApprovePgsUser(request *models.PgsUpdateApproval) (responses bool, err error)
	RejectPgsUser(request *models.PgsUpdateApproval) (responses bool, err error)
	// Login(request models.LoginRequest) (responses interface{}, err error)
	// GetMenu(request modelsMU.MenuRequest) (responses []modelsMU.MenuResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.PgsUserResponses, pagination lib.Pagination, err error)
	Delete(request *models.UpdateDelete) (response bool, err error)
	SearchPekerjaByPn(PERNR string) (responses models.UserResponseLocal, err error)
}

type PgsUserService struct {
	db             lib.Database
	dbRaw          lib.Databases
	logger         logger.Logger
	pgsUserRepo    pgsUser.PgsUserDefinition
	approval       pgsUser.PgsUserApprovalDefinition
	managementUser managementUser.ManagementUserDefinition
	jwtService     services.JWTAuthService
	MstRole        mstRole.MstRoleDefinition
	// serviceMenu servicesMenu.ManagementUserDefinition
}

func NewPgsUserService(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
	pgsUserRepo pgsUser.PgsUserDefinition,
	approval pgsUser.PgsUserApprovalDefinition,
	managementUser managementUser.ManagementUserDefinition,
	jwtService services.JWTAuthService,
	MstRole mstRole.MstRoleDefinition,
	// serviceMenu servicesMenu.ManagementUserDefinition,
) PgsUserDefinition {
	return PgsUserService{
		db:             db,
		dbRaw:          dbRaw,
		logger:         logger,
		pgsUserRepo:    pgsUserRepo,
		approval:       approval,
		managementUser: managementUser,
		jwtService:     jwtService,
		MstRole:        MstRole,
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Delete implements PgsUserDefinition
func (pgs PgsUserService) Delete(request *models.UpdateDelete) (response bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := pgs.db.DB.Begin()

	getOneData, exist, err := pgs.GetOne(request.ID)
	if err != nil {
		pgs.logger.Zap.Error(err)
		tx.Rollback()
		return false, err
	}

	updateDataIndicator := &models.UpdateDelete{
		ID:         request.ID,
		DeleteFlag: true,
		UpdatedAt:  &timeNow,
	}

	_, err = pgs.pgsUserRepo.Delete(updateDataIndicator,
		[]string{
			"delete_flag",
			"updated_at",
		}, tx)

	if err != nil {
		tx.Rollback()
		pgs.logger.Zap.Error(err)
		return false, err
	}

	if exist {
		fmt.Println("getOne", getOneData)
		tx.Commit()
		return true, err
	}

	return false, err
}

// GetAll implements PgsUserDefinition
func (pgsUser PgsUserService) GetAll(makerID string) (responses []models.PgsUserResponses, err error) {
	return pgsUser.pgsUserRepo.GetAll(makerID)
}

// GetAllWithPaginate implements PgsUserDefinition
func (pgs PgsUserService) GetAllWithPaginate(request models.Paginate) (responses []models.PgsUserResponses, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPgs, totalRows, totalData, err := pgs.pgsUserRepo.GetAllWithPaginate(&request)
	if err != nil {
		pgs.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPgs {
		responses = append(responses, models.PgsUserResponses{
			ID:            response.ID,
			PN:            response.PN,
			NamaPekerja:   response.NamaPekerja,
			UnitKerja:     response.UnitKerja,
			JabatanPgs:    response.JabatanPgs,
			PeriodeAwal:   response.PeriodeAwal,
			PeriodeAkhir:  response.PeriodeAkhir,
			MakerID:       response.MakerID,
			MakerDesc:     response.MakerDesc,
			MakerDate:     response.MakerDate,
			LastMakerID:   response.LastMakerDesc,
			LastMakerDesc: response.LastMakerDesc,
			LastMakerDate: response.LastMakerDate,
			Status:        response.Status,
			Action:        response.Action,
			CreatedAt:     response.CreatedAt,
			UpdatedAt:     response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements PgsUserDefinition
func (pgsUser PgsUserService) GetOne(id int64) (responses models.PgsUserResponseOne, status bool, err error) {
	dataPgs, err := pgsUser.pgsUserRepo.GetOne(id)
	if dataPgs.ID != 0 {
		approval, err := pgsUser.approval.GeOneApprovalByID(dataPgs.ID)

		responses = models.PgsUserResponseOne{
			ID:            dataPgs.ID,
			PN:            dataPgs.PN,
			NamaPekerja:   dataPgs.NamaPekerja,
			UnitKerja:     dataPgs.UnitKerja,
			REGION:        dataPgs.REGION,
			RGDESC:        dataPgs.RGDESC,
			RGNAME:        dataPgs.RGNAME,
			MAINBR:        dataPgs.MAINBR,
			MBDESC:        dataPgs.MBDESC,
			MBNAME:        dataPgs.MBNAME,
			BRANCH:        dataPgs.BRANCH,
			BRDESC:        dataPgs.BRDESC,
			BRNAME:        dataPgs.BRNAME,
			JabatanPgs:    dataPgs.JabatanPgs,
			PeriodeAwal:   dataPgs.PeriodeAwal,
			PeriodeAkhir:  dataPgs.PeriodeAkhir,
			MakerID:       dataPgs.MakerID,
			MakerDesc:     dataPgs.MakerDesc,
			MakerDate:     dataPgs.MakerDate,
			LastMakerID:   dataPgs.LastMakerID,
			LastMakerDesc: dataPgs.LastMakerDesc,
			LastMakerDate: dataPgs.LastMakerDate,
			Status:        dataPgs.Status,
			Action:        dataPgs.Action,
			Approval:      approval,
			CreatedAt:     dataPgs.CreatedAt,
			UpdatedAt:     dataPgs.UpdatedAt,
		}

		return responses, true, err
	}

	return responses, false, err
}

// Store implements PgsUserDefinition
func (pgsUser PgsUserService) Store(request models.PgsUserRequest) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	cekPgsExist, err := pgsUser.pgsUserRepo.CekPgsActive(request.PN, request.BRANCH)

	if cekPgsExist < 1 {
		fmt.Println("massokk")
		fmt.Println("jumlah ===>", cekPgsExist)
		tx := pgsUser.db.DB.Begin()

		reqPgsUser := &models.PgsUser{
			PN:           request.PN,
			NamaPekerja:  request.NamaPekerja,
			UnitKerja:    request.UnitKerja,
			REGION:       request.REGION,
			RGDESC:       request.RGDESC,
			RGNAME:       request.RGNAME,
			MAINBR:       request.MAINBR,
			MBDESC:       request.MBDESC,
			MBNAME:       request.MBNAME,
			BRANCH:       request.BRANCH,
			BRDESC:       request.BRDESC,
			BRNAME:       request.BRNAME,
			JabatanPgs:   request.JabatanPgs,
			PeriodeAwal:  request.PeriodeAwal,
			PeriodeAkhir: request.PeriodeAkhir,
			MakerID:      request.MakerID,
			MakerDesc:    request.MakerDesc,
			MakerDate:    &timeNow,
			Status:       "01a",
			Action:       "New Request",
			CreatedAt:    &timeNow,
		}

		dataPgsUser, err := pgsUser.pgsUserRepo.Store(reqPgsUser, tx)
		fmt.Println("data => 01b", dataPgsUser)

		if err != nil {
			tx.Rollback()
			pgsUser.logger.Zap.Error(err)
			return false, err
		}

		if len(request.Approval) != 0 {
			for _, value := range request.Approval {
				_, err = pgsUser.approval.Store(&models.PgsUserApproval{
					IDPgsUser:      dataPgsUser.ID,
					ApprovalID:     value.ApprovalID,
					ApprovalDesc:   value.ApprovalDesc,
					ApprovalStatus: "0",
				}, tx)

				if err != nil {
					tx.Rollback()
					pgsUser.logger.Zap.Error(err)
					return false, err
				}
			}
		} else {
			if err != nil {
				tx.Rollback()
				pgsUser.logger.Zap.Error(err)
				return false, err
			}
		}

		tx.Commit()
		return true, err
	}

	fmt.Println("gagal")
	return false, err

}

// Update implements PgsUserDefinition
func (pgsUser PgsUserService) Update(request *models.PgsUserRequestUpdate) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := pgsUser.db.DB.Begin()

	updatePgs := &models.PgsUserRequestMaintainance{
		ID:            request.ID,
		PN:            request.PN,
		NamaPekerja:   request.NamaPekerja,
		REGION:        request.REGION,
		RGDESC:        request.RGDESC,
		RGNAME:        request.RGNAME,
		MAINBR:        request.MAINBR,
		MBDESC:        request.MBDESC,
		MBNAME:        request.MBNAME,
		BRANCH:        request.BRANCH,
		BRDESC:        request.BRDESC,
		BRNAME:        request.BRNAME,
		JabatanPgs:    request.JabatanPgs,
		PeriodeAwal:   request.PeriodeAwal,
		PeriodeAkhir:  request.PeriodeAkhir,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "01b",
		Action:        "UpdateData",
		UpdatedAt:     &timeNow,
	}

	include := []string{
		"pn",
		"nama_pekerja",
		"unit_kerja",
		"jabatan_pgs",
		"periode_awal",
		"periode_akhir",
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"updated_at",
		"status",
		"action",
	}

	_, err = pgsUser.pgsUserRepo.Update(updatePgs, include, tx)

	if err != nil {
		tx.Rollback()
		pgsUser.logger.Zap.Error(err)
		return false, err
	}

	if len(request.Approval) != 0 {
		for _, value := range request.Approval {
			_, err = pgsUser.approval.Store(&models.PgsUserApproval{
				ID:             value.ID,
				IDPgsUser:      request.ID,
				ApprovalID:     value.ApprovalID,
				ApprovalDesc:   value.ApprovalDesc,
				ApprovalStatus: "0",
			}, tx)

			if err != nil {
				tx.Rollback()
				pgsUser.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			pgsUser.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

// GetPgsApproval implements PgsUserDefinition
func (pgsUser PgsUserService) GetPgsApproval(request models.Paginate) (responses []models.PgsApprovalResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataPgs, totalRows, totalData, err := pgsUser.pgsUserRepo.GetPgsApproval(&request)
	if err != nil {
		pgsUser.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataPgs {
		responses = append(responses, models.PgsApprovalResponse{
			ID:             response.ID.Int64,
			PN:             response.PN.String,
			NamaPekerja:    response.NamaPekerja.String,
			UnitKerja:      response.UnitKerja.String,
			JabatanPgs:     response.JabatanPgs.String,
			IDApproval:     response.IDApproval.Int64,
			Approval:       response.Approval.String,
			ApprovalDate:   response.ApprovalDate.String,
			ApprovalStatus: response.ApprovalStatus.Bool,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// ApprovePgsUser implements PgsUserDefinition
func (pgsUser PgsUserService) ApprovePgsUser(request *models.PgsUpdateApproval) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := pgsUser.db.DB.Begin()

	updateApproval := &models.PgsUpdateRequest{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02a",
		Action:        "Active",
	}

	include := []string{
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"updated_at",
		"action",
		"status",
	}

	_, err = pgsUser.pgsUserRepo.ApprovePgsUser(updateApproval, include, tx)

	if err != nil {
		tx.Rollback()
		pgsUser.logger.Zap.Error(err)
		return false, err
	}

	if len(request.Approval) != 0 {
		for _, value := range request.Approval {
			_, err = pgsUser.approval.ApprovalUpdate(&models.ApprovalUpdate{
				ID:             value.ID,
				ApprovalDesc:   request.LastMakerDesc,
				ApprovalDate:   timeNow,
				ApprovalStatus: "1",
			}, tx)

			if err != nil {
				tx.Rollback()
				pgsUser.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			pgsUser.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

// RejectPgsUser implements PgsUserDefinition
func (pgsUser PgsUserService) RejectPgsUser(request *models.PgsUpdateApproval) (responses bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := pgsUser.db.DB.Begin()

	updateApproval := &models.PgsUpdateRequest{
		ID:            request.ID,
		LastMakerID:   request.LastMakerID,
		LastMakerDesc: request.LastMakerDesc,
		LastMakerDate: &timeNow,
		Status:        "02b",
		Action:        "Reject",
	}

	include := []string{
		"last_maker_id",
		"last_maker_desc",
		"last_maker_date",
		"updated_at",
		"action",
		"status",
	}

	_, err = pgsUser.pgsUserRepo.ApprovePgsUser(updateApproval, include, tx)

	if err != nil {
		tx.Rollback()
		pgsUser.logger.Zap.Error(err)
		return false, err
	}

	if len(request.Approval) != 0 {
		for _, value := range request.Approval {
			_, err = pgsUser.approval.ApprovalUpdate(&models.ApprovalUpdate{
				ID:             value.ID,
				ApprovalDesc:   request.LastMakerDesc,
				ApprovalDate:   timeNow,
				ApprovalStatus: "2",
			}, tx)

			if err != nil {
				tx.Rollback()
				pgsUser.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		if err != nil {
			tx.Rollback()
			pgsUser.logger.Zap.Error(err)
			return false, err
		}
	}

	tx.Commit()
	return true, err
}

func (pekerja PgsUserService) SearchPekerjaByPn(PERNR string) (responses models.UserResponseLocal, err error) {
	responses, err = pekerja.pgsUserRepo.SearchPekerjaByPn(PERNR)

	return responses, err
}
