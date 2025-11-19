package managementuser

import (
	"fmt"
	"riskmanagement/lib"
	models "riskmanagement/models/managementuser"
	"riskmanagement/models/mstrole"
	managementUser "riskmanagement/repository/managementuser"
	repository "riskmanagement/repository/managementuser"
	mstRole "riskmanagement/repository/mstrole"
	pgsRepo "riskmanagement/repository/pgsuser"
	"riskmanagement/repository/unitkerja"
	services "riskmanagement/services/auth"

	"gitlab.com/golang-package-library/logger"
)

// type ManagementUserDefinition interface {
// 	GetAll() (responses []models.ManagementUserResponse, err error)
// 	GetOne(id int64) (responses models.ManagementUserResponse, err error)
// 	Store(request *models.ManagementUserRequest) (response bool, err error)
// 	Update(request *models.ManagementUserRequest) (response bool, err error)
// 	Delete(id int64) (err error)
// 	MappingMenu(request models.MappingMenuRequest) (responses bool, err error)
// 	GetMappingMenu(id int64) (responses models.ManagementUserResponses, err error)
// 	GetAllMenu() (responses []models.Menu, err error)
// 	DeleteMappingMenu(request *models.MapMenu) (status bool, err error)
// 	GetMenu(request models.MenuRequest) (responses []models.MenuResponse, err error)
// 	GetAllWithPaginate(request models.Paginate) (responses []models.ManagementUserFinResponse, pagination lib.Pagination, err error)
// 	GetUkerKelolaan(request models.UkerKelolaanUserRequest) (responses []models.UkerKelolaanUserResponse, err error)
// 	GetTreeMenu() (responses []models.MenuResponse, err error)
// 	GetLevelUker() (responses []models.LevelUkerResponse, err error)
// 	GetJabatanRole() (responses []models.JabatanRolesResponse, err error)
// }

type ManagementUserService struct {
	db             lib.Database
	dbRaw          lib.Databases
	logger         logger.Logger
	repository     repository.ManagementUserDefinition
	mapMenu        repository.MapMenuDefinition
	jwtService     services.JWTAuthService
	MstRole        mstRole.MstRoleDefinition
	managementUser managementUser.ManagementUserDefinition
	pgsRepo        pgsRepo.PgsUserDefinition
	UnitKerja      unitkerja.UnitKerjaDefinition
}

func NewManagementUserService(
	db lib.Database,
	dbRaw lib.Databases,
	logger logger.Logger,
	repository repository.ManagementUserDefinition,
	mapMenu repository.MapMenuDefinition,
	jwtService services.JWTAuthService,
	MstRole mstRole.MstRoleDefinition,
	managementUser managementUser.ManagementUserDefinition,
	pgsRepo pgsRepo.PgsUserDefinition,
	UnitKerja unitkerja.UnitKerjaDefinition,
) ManagementUserService {
	return ManagementUserService{
		db:             db,
		dbRaw:          dbRaw,
		logger:         logger,
		repository:     repository,
		mapMenu:        mapMenu,
		jwtService:     jwtService,
		MstRole:        MstRole,
		managementUser: managementUser,
		pgsRepo:        pgsRepo,
		UnitKerja:      UnitKerja,
	}
}

// Delete implements  ManagementUserDefinition
func (managementuser ManagementUserService) Delete(id int64) (err error) {
	return managementuser.repository.Delete(id)
}

// GetAll implements  ManagementUserDefinition
func (managementuser ManagementUserService) GetAll() (responses []models.ManagementUserResponse, err error) {
	return managementuser.repository.GetAll()
}

// GetAllWithPaginate implements ManagementUserDefinition
func (mu ManagementUserService) GetAllWithPaginate(request models.Paginate) (responses []models.ManagementUserFinResponse, pagination lib.Pagination, err error) {
	offset, page, limit, order, sort := lib.SetPaginationParameter(request.Page, request.Limit, request.Order, request.Sort)
	request.Offset = offset
	request.Order = order
	request.Sort = sort

	dataRoles, totalRows, totalData, err := mu.repository.GetAllWithPaginate(&request)
	if err != nil {
		mu.logger.Zap.Error(err)
		return responses, pagination, err
	}

	for _, response := range dataRoles {
		responses = append(responses, models.ManagementUserFinResponse{
			ID:        response.ID,
			Role:      response.Role,
			LevelUker: response.LevelUker,
			StellTx:   response.StellTx,
			Jgpg:      response.Jgpg,
		})
	}

	pagination = lib.SetPaginationResponse(page, limit, totalRows, totalData)
	return responses, pagination, err
}

// GetOne implements  ManagementUserDefinition
func (managementuser ManagementUserService) GetOne(id int64) (responses models.ManagementUserResponse, err error) {
	return managementuser.repository.GetOne(id)
}

// Store implements  ManagementUserDefinition
func (managementuser ManagementUserService) Store(request *models.ManagementUserRequest) (response bool, err error) {
	response, err = managementuser.repository.Store(request)
	return response, err
}

// Update implements  ManagementUserDefinition
func (managementuser ManagementUserService) Update(request *models.ManagementUserRequest) (response bool, err error) {
	response, err = managementuser.repository.Update(request)
	return response, err
}

// MappingMenu implements ManagementUserDefinition
func (managementuser ManagementUserService) MappingMenu(request models.MappingMenuRequest) (responses bool, err error) {
	tx := managementuser.db.DB.Begin()

	if len(request.MapMenu) != 0 {
		for _, value := range request.MapMenu {
			_, err = managementuser.mapMenu.Store(&models.MapMenu{
				ID:         value.ID,
				IDJabatan:  request.ID,
				IDMenu:     value.IDMenu,
				Keterangan: value.Keterangan,
			}, tx)

			if err != nil {
				tx.Rollback()
				managementuser.logger.Zap.Error(err)
				return false, err
			}
		}
	} else {
		tx.Rollback()
		managementuser.logger.Zap.Error(err)
		return false, err
	}

	tx.Commit()
	return true, err
}

// GetMappingMenu implements ManagementUserDefinition
func (managementuser ManagementUserService) GetMappingMenu(id int64) (responses models.ManagementUserResponses, err error) {
	dataMenu, err := managementuser.repository.GetOne(id)
	if dataMenu.ID != 0 {
		dataMenus, err := managementuser.mapMenu.GetOneDataByID(dataMenu.ID)

		responses = models.ManagementUserResponses{
			ID:        dataMenu.ID,
			RoleID:    dataMenu.RoleID,
			LevelUker: dataMenu.LevelUker,
			LevelID:   dataMenu.LevelID,
			MapMenu:   dataMenus,
		}

		return responses, err
	}

	return responses, err
}

// GetAllMenu implements ManagementUserDefinition
func (managementuser ManagementUserService) GetAllMenu() (responses []models.Menu, err error) {
	return managementuser.repository.GetAllMenu()
}

// DeleteMappingMenu implements ManagementUserDefinition
func (managementuser ManagementUserService) DeleteMappingMenu(request *models.MapMenu) (status bool, err error) {
	tx := managementuser.db.DB.Begin()

	err = managementuser.repository.DeleteMappingMenu(request.ID, tx)
	if err != nil {
		tx.Rollback()
		managementuser.logger.Zap.Error(err)
		return false, err
	}
	tx.Commit()
	return true, err
}

// GetMenu implements ManagementUserDefinition
func (mu ManagementUserService) GetMenu(request models.MenuListRequest) (list_menu []models.MenuResponse, additional_menu []models.AdditionalMenuResponse, err error) {
	menuList, err := mu.MstRole.GetMenuList(mstrole.MenuListRequest{
		TIPEUKER: request.TIPEUKER,
		HILFM:    request.HILFM,
		ORGEH:    request.ORGEH,
		PERNR:    request.PERNR,
		KOSTL:    request.KOSTL,
		StellTx:  request.StellTx,
		Jgpg:     request.Jgpg,
	})

	if err != nil {
		mu.logger.Zap.Error(err)
		return list_menu, additional_menu, err

	}

	menuRoleQuest, err := mu.MstRole.GetMenuListQuestionnaire(mstrole.MenuListRequest{
		TIPEUKER: request.TIPEUKER,
		HILFM:    request.HILFM,
		ORGEH:    request.ORGEH,
		PERNR:    request.PERNR,
		KOSTL:    request.KOSTL,
		StellTx:  request.StellTx,
		Jgpg:     request.Jgpg,
	})

	if err != nil {
		mu.logger.Zap.Error(err)
		return list_menu, additional_menu, err
	}

	//get accessible Menu
	// menuList.Menu = menuList.Menu + "," + menuRoleQuest.Menu
	listMenu := menuList.Menu + "," + menuRoleQuest.Menu

	menus, err := mu.repository.GetMenu(menuList.Menu)
	if err != nil {
		mu.logger.Zap.Error(err)
		return list_menu, additional_menu, err
	}

	questMenus, err := mu.repository.GetMenuQuestionnaire(menuRoleQuest.Menu)

	if err != nil {
		mu.logger.Zap.Error(err)
		return list_menu, additional_menu, err
	}

	// menus = append(menus, questMenu...)

	// Menu Lv2
	for _, menu := range menus {
		var childMenus []models.ChildMenuResponse
		childDatas, err := mu.repository.GetChildMenu(menu.IDMenu, menuList.Menu)
		if err != nil {
			mu.logger.Zap.Error(err)
			return list_menu, additional_menu, err
		}

		// fmt.Println("rows =>", childDatas)

		// menu lv3
		for _, childData := range childDatas {
			var subChildMenus []models.SubChildMenuResponse
			subChildDatas, err := mu.repository.GetSubChildMenu(childData.IDMenu, menuList.Menu)
			if err != nil {
				mu.logger.Zap.Error(err)
				return list_menu, additional_menu, err
			}

			subChildDataQuests, err := mu.repository.GetSubChildMenuQuest(childData.IDMenu, listMenu)
			if err != nil {
				mu.logger.Zap.Error(err)
				return list_menu, additional_menu, err
			}

			subChildDatas = append(subChildDatas, subChildDataQuests...)

			for _, subChildData := range subChildDatas {
				subChildMenus = append(subChildMenus, models.SubChildMenuResponse{
					IDMenu:   subChildData.IDMenu,
					Title:    subChildData.Title,
					Url:      subChildData.Url,
					Icon:     subChildData.Url,
					SvgIcon:  subChildData.SvgIcon,
					FontIcon: subChildData.FontIcon,
				})
			}

			childMenus = append(childMenus, models.ChildMenuResponse{
				IDMenu:   childData.IDMenu,
				Title:    childData.Title,
				Url:      childData.Url,
				Icon:     childData.Icon,
				SvgIcon:  childData.SvgIcon,
				FontIcon: childData.FontIcon,
				SubChild: subChildMenus,
			})
		}

		childDataQuests, err := mu.repository.GetChildMenuQuest(menu.IDMenu, listMenu)
		if err != nil {
			mu.logger.Zap.Error(err)
			return list_menu, additional_menu, err
		}

		for _, childDataQuest := range childDataQuests {
			childMenus = append(childMenus, models.ChildMenuResponse{
				IDMenu:   childDataQuest.IDMenu,
				Title:    childDataQuest.Title,
				Url:      childDataQuest.Url,
				Icon:     childDataQuest.Icon,
				SvgIcon:  childDataQuest.SvgIcon,
				FontIcon: childDataQuest.FontIcon,
			})
		}

		list_menu = append(list_menu, models.MenuResponse{
			IDMenu:    menu.IDMenu,
			Title:     menu.Title,
			Url:       menu.Url,
			Deskripsi: menu.Deskripsi,
			Icon:      menu.Icon,
			SvgIcon:   menu.SvgIcon,
			FontIcon:  menu.FontIcon,
			Child:     childMenus,
		})
	}

	for _, questMenu := range questMenus {
		var childMenus []models.ChildMenuResponse

		list_menu = append(list_menu, models.MenuResponse{
			IDMenu:    questMenu.IDMenu,
			Title:     questMenu.Title,
			Url:       questMenu.Url,
			Deskripsi: questMenu.Deskripsi,
			Icon:      questMenu.Icon,
			SvgIcon:   questMenu.SvgIcon,
			FontIcon:  questMenu.FontIcon,
			Child:     childMenus,
		})
	}

	addMenu, err := mu.repository.GetAdditionalMenuById(menuList.AdditionalMenu)

	for _, value := range addMenu {
		additional_menu = append(additional_menu, models.AdditionalMenuResponse{
			Id:   value.Id,
			Nama: value.Nama,
			Url:  value.Url,
			Icon: value.Icon,
		})
	}

	return list_menu, additional_menu, err
}

func (mu ManagementUserService) GetUkerKelolaan(request models.UkerKelolaanUserRequest) (responses []models.UkerKelolaanUserResponse, err error) {
	data, err := mu.repository.GetUkerKelolaan(&request)
	if err != nil {
		mu.logger.Zap.Error(err)
		return responses, err
	}

	for _, response := range data {
		responses = append(responses, models.UkerKelolaanUserResponse{
			Id:        response.Id.Int64,
			CreatedAt: response.CreatedAt.String,
			UpdatedAt: response.UpdatedAt.String,
			ExpiredAt: response.ExpiredAt.String,
			IsTemp:    response.IsTemp.Int64,
			PN:        response.PN.String,
			REGION:    response.REGION.String,
			RGDESC:    response.RGDESC.String,
			MAINBR:    response.MAINBR.String,
			MBDESC:    response.MBDESC.String,
			BRANCH:    response.BRANCH.String,
			BRDESC:    response.BRDESC.String,
		})
	}

	return responses, err
}

// GetTreeMenu implements ManagementUserDefinition
func (mu ManagementUserService) GetTreeMenu() (responses []models.MenuResponse, err error) {
	menus, err := mu.repository.GetTreeMenu()
	if err != nil {
		mu.logger.Zap.Error(err)
		return responses, err
	}

	fmt.Println("rows =>", menus)

	for _, menu := range menus {
		var childMenus []models.ChildMenuResponse
		childDatas, err := mu.repository.GetChildTreeMenu(menu.IDMenu)
		if err != nil {
			mu.logger.Zap.Error(err)
			return responses, err
		}

		for _, childData := range childDatas {
			var subChildMenus []models.SubChildMenuResponse
			subChildDatas, err := mu.repository.GetSubChildTreeMenu(childData.IDMenu)

			if err != nil {
				mu.logger.Zap.Error(err)
				return responses, err
			}

			for _, subChildData := range subChildDatas {
				subChildMenus = append(subChildMenus, models.SubChildMenuResponse{
					IDMenu:   subChildData.IDMenu,
					Title:    subChildData.Title,
					Url:      subChildData.Url,
					Icon:     subChildData.Icon,
					SvgIcon:  subChildData.SvgIcon,
					FontIcon: subChildData.FontIcon,
				})
			}

			childMenus = append(childMenus, models.ChildMenuResponse{
				IDMenu:   childData.IDMenu,
				Title:    childData.Title,
				Url:      childData.Url,
				Icon:     childData.Icon,
				SvgIcon:  childData.SvgIcon,
				FontIcon: childData.FontIcon,
				SubChild: subChildMenus,
			})
		}

		responses = append(responses, models.MenuResponse{
			IDMenu:    menu.IDMenu,
			Title:     menu.Title,
			Url:       menu.Url,
			Deskripsi: menu.Deskripsi,
			Icon:      menu.Icon,
			SvgIcon:   menu.SvgIcon,
			FontIcon:  menu.FontIcon,
			Child:     childMenus,
		})
	}

	return responses, err
}

// GetLevelUker implements ManagementUserDefinition
func (mu ManagementUserService) GetLevelUker() (responses []models.LevelUkerResponse, err error) {
	return mu.repository.GetLevelUker()
}

// Enhance Management User By Panji 02/02/2024s
// GetJabatanRole implements ManagementUserDefinition.
func (mu ManagementUserService) GetJabatanRole() (responses []models.JabatanRolesResponse, err error) {
	return mu.repository.GetJabatanRole()
}

func (mu ManagementUserService) GetAdditionalMenu() (responses []models.AdditionalMenuResponse, err error) {
	dataMenu, err := mu.repository.GetAdditionalMenu()

	for _, value := range dataMenu {
		responses = append(responses, models.AdditionalMenuResponse{
			Id:   value.Id,
			Nama: value.Nama,
			Url:  value.Url,
			Icon: value.Icon,
		})
	}

	return responses, err
}
