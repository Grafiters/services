package mstrole

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/mstrole"
	roleRepo "riskmanagement/repository/mstrole"

	"github.com/google/uuid"
	"gitlab.com/golang-package-library/logger"
)

var (
	UUID = uuid.NewString()
)

type MstRoleDefinition interface {
	GetAll() (responses []models.MstRoleResponse, err error)
	GetAllWithPaginate(request models.Paginate) (responses []models.MstRoleResponse, pagination lib.Pagination, err error)
	GetOne(id int64) (responses models.MstRoleResponseOne, status bool, err error)
	Store(request models.MstRoleRequest) (status bool, err error)
	Update(request *models.MstRoleRequest) (status bool, err error)
	Delete(request *models.MstRoleRequestDelete) (status bool, err error)
}

type MstRoleService struct {
	db          lib.Database
	logger      logger.Logger
	mstRoleRepo roleRepo.MstRoleDefinition
	mstRoleMenu roleRepo.MstRoleMenuDefinition
}

func NewMstRoleService(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
	mstRoleRepo roleRepo.MstRoleDefinition,
	mstRoleMenu roleRepo.MstRoleMenuDefinition,
) MstRoleDefinition {
	return MstRoleService{
		db:          db,
		logger:      logger,
		mstRoleRepo: mstRoleRepo,
		mstRoleMenu: mstRoleMenu,
	}
}

// GetAll implements MstRoleDefinition
func (mstRole MstRoleService) GetAll() (responses []models.MstRoleResponse, err error) {
	return mstRole.mstRoleRepo.GetAll()
}

// GetAllWithPaginate implements MstRoleDefinition
func (msRole MstRoleService) GetAllWithPaginate(request models.Paginate) (responses []models.MstRoleResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataRoles, totalRows, totalData, err := msRole.mstRoleRepo.GetAllWithPaginate(&request)
	if err != nil {
		msRole.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataRoles {
		responses = append(responses, models.MstRoleResponse{
			ID:         response.ID,
			RoleName:   response.RoleName,
			Menu:       response.Menu,
			DeleteFlag: response.DeleteFlag,
			CreatedAt:  response.CreatedAt,
			UpdatedAt:  response.UpdatedAt,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements MstRoleDefinition
func (mstRole MstRoleService) GetOne(id int64) (responses models.MstRoleResponseOne, status bool, err error) {
	dataRole, err := mstRole.mstRoleRepo.GetOne(id)

	if dataRole.ID != 0 {
		// menus, err := mstRole.mstRoleMenu.GetByIDRole(dataRole.ID)
		responses = models.MstRoleResponseOne{
			ID:             dataRole.ID,
			RoleName:       dataRole.RoleName,
			AddonPernr:     dataRole.AddonPernr,
			Menu:           dataRole.Menu,
			AdditionalMenu: dataRole.AdditionalMenu,
			DeleteFlag:     dataRole.DeleteFlag,
			// Menu:       menus,
			CreatedAt: dataRole.CreatedAt,
			UpdatedAt: dataRole.UpdatedAt,
		}

		return responses, true, err
	}

	return responses, false, err
}

// Store implements MstRoleDefinition
func (mstRole MstRoleService) Store(request models.MstRoleRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")
	tx := mstRole.db.DB.Begin()

	reqInsert := &models.MstRole{
		RoleName:       request.RoleName,
		AddonPernr:     request.AddonPernr,
		Menu:           request.Menu,
		AdditionalMenu: request.AdditionalMenu,
		DeleteFlag:     false,
		CreatedAt:      &timeNow,
	}

	dataInsert, err := mstRole.mstRoleRepo.Store(reqInsert, tx)
	fmt.Println(dataInsert)

	if err != nil {
		tx.Rollback()
		mstRole.logger.Zap.Error(err)
		return false, err
	}

	// fmt.Println(dataInsert.ID)

	// if len(request.Menu) != 0 {
	// 	for _, value := range request.Menu {
	// 		// fmt.Println(dataInsert.ID, value.IDMenu)
	// 		_, err = mstRole.mstRoleMenu.Store(&models.MstRoleMapMenu{
	// 			IDRole: dataInsert.ID,
	// 			IDMenu: value.IDMenu,
	// 		}, tx)

	// 		if err != nil {
	// 			tx.Rollback()
	// 			mstRole.logger.Zap.Error(err)
	// 			return false, err
	// 		}
	// 	}
	// } else {
	// 	tx.Rollback()
	// 	mstRole.logger.Zap.Error(err)
	// 	return false, err
	// }

	tx.Commit()
	return true, err

}

// Update implements MstRoleDefinition
func (mstRole MstRoleService) Update(request *models.MstRoleRequest) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := mstRole.db.DB.Begin()

	reqUpdate := &models.MstRoleRequestUpdate{
		ID:             request.ID,
		AddonPernr:     request.AddonPernr,
		RoleName:       request.RoleName,
		Menu:           request.Menu,
		AdditionalMenu: request.AdditionalMenu,
		UpdatedAt:      &timeNow,
	}

	include := []string{
		"id",
		"role_name",
		"updated_at",
	}

	_, err = mstRole.mstRoleRepo.Update(reqUpdate, include, tx)

	if err != nil {
		tx.Rollback()
		mstRole.logger.Zap.Error(err)
		return false, err
	}

	// if len(request.Menu) != 0 {
	// 	for _, value := range request.Menu {
	// 		_, err = mstRole.mstRoleMenu.Store(&models.MstRoleMapMenu{
	// 			ID:     value.ID,
	// 			IDRole: request.ID,
	// 			IDMenu: value.IDMenu,
	// 		}, tx)

	// 		if err != nil {
	// 			tx.Rollback()
	// 			mstRole.logger.Zap.Error(err)
	// 			return false, err
	// 		}
	// 	}
	// } else {
	// 	tx.Rollback()
	// 	mstRole.logger.Zap.Error(err)
	// 	return false, err
	// }

	tx.Commit()
	return true, err
	// panic("uniplained")

}

// Delete implements MstRoleDefinition
func (mstRole MstRoleService) Delete(request *models.MstRoleRequestDelete) (status bool, err error) {
	timeNow := lib.GetTimeNow("timestime")

	tx := mstRole.db.DB.Begin()

	reqDelete := &models.MstRoleRequestDelete{
		ID:         request.ID,
		DeleteFlag: true,
		UpdatedAt:  &timeNow,
	}

	include := []string{
		"id",
		"deleted_flag",
		"updated_at",
	}

	_, err = mstRole.mstRoleRepo.Delete(reqDelete, include, tx)

	if err != nil {
		tx.Rollback()
		mstRole.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}
