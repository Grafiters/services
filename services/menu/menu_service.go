package menu

import (
	"fmt"
	"riskmanagement/lib"

	menuModels "riskmanagement/models/menu"
	roles "riskmanagement/models/mstrole"
	types "riskmanagement/models/type"

	menuRepo "riskmanagement/repository/menu"
	mstRole "riskmanagement/repository/mstrole"

	jwt "riskmanagement/services/auth"

	"gitlab.com/golang-package-library/logger"
)

var (
	timeNow = lib.GetTimeNow("timestime")
)

type MenuServiceDefinition interface {
	GetMenuTree(request menuModels.MenuRequest) (responses []menuModels.MenuResponse, err error)
	GetKuisioner(request menuModels.RequestKuisioner) (responses menuModels.ResponseData, err error)

	// Modul Qna
	GetAll() (responses []menuModels.MenuQnaResponse, err error)
	SubMenuCheck(request menuModels.MenuQnaRequest) (response menuModels.MenuQnaResponse, err error)
	GetAllMstMenu() (responses []menuModels.MstMenu, err error)
	DeleteMenuRRM(request menuModels.MstMenuRequest) (status bool, err error)
	StoreMstRRM(request menuModels.MstMenu) (id string, err error)
	DeleteRole(request menuModels.MstMenuRequest) (err error)
	StoreRoleRRM(request []types.Role) (err error)
	SetStatus(request menuModels.MstMenu) (response bool, err error)

	GetLastID() (id_menu int64, err error)
}

type MenuService struct {
	logger     logger.Logger
	repo       menuRepo.MenuDefinition
	role       mstRole.MstRoleDefinition
	db         lib.Database
	jwtService jwt.JWTAuthService
}

func NewMenuService(
	db lib.Database,
	logger logger.Logger,
	repo menuRepo.MenuDefinition,
	role mstRole.MstRoleDefinition,
	jwtService jwt.JWTAuthService,
) MenuServiceDefinition {
	return MenuService{
		db:         db,
		logger:     logger,
		repo:       repo,
		role:       role,
		jwtService: jwtService,
	}
}

// GetMenu implements MenuServiceDefinition.
func (m MenuService) GetMenuTree(request menuModels.MenuRequest) (responses []menuModels.MenuResponse, err error) {
	data, err := m.role.GetMenuList(roles.MenuListRequest{
		TIPEUKER: request.TipeUker,
		HILFM:    request.Hilfm,
		ORGEH:    request.Orgeh,
		PERNR:    request.Pernr,
		KOSTL:    request.Kostl,
		StellTx:  request.StellTx,
		Jgpg:     request.Jgpg,
	})

	if err != nil {
		m.logger.Zap.Error(err)
		return responses, err
	}

	menus, err := m.repo.GetMenuTree(data.MfeMenu, 0)
	// fmt.Println("menus", menus)

	if err != nil {
		return responses, err
	}

	for _, res := range menus {
		var childMenus []menuModels.SubMenu
		child, err := m.repo.GetMenuTree(data.MfeMenu, int(res.Id))
		if err != nil {
			m.logger.Zap.Error(err)
			return responses, err
		}

		for _, resChild := range child {
			var subChildMenus []menuModels.Menu
			subChild, err := m.repo.GetMenuTree(data.MfeMenu, int(resChild.Id))
			if err != nil {
				m.logger.Zap.Error(err)
				return responses, err
			}

			for _, resSubChild := range subChild {
				subChildMenus = append(subChildMenus, menuModels.Menu{
					Id:        resSubChild.Id,
					Title:     resSubChild.Title,
					Icon:      resSubChild.Icon,
					Path:      resSubChild.Path,
					IsSection: resSubChild.IsSection,
					ParentID:  resSubChild.ParentID,
				})
			}

			childMenus = append(childMenus, menuModels.SubMenu{
				Id:           resChild.Id,
				Title:        resChild.Title,
				Icon:         resChild.Icon,
				Path:         resChild.Path,
				ParentID:     resChild.ParentID,
				IsSection:    resChild.IsSection,
				SubChildMenu: subChildMenus,
			})
		}

		responses = append(responses, menuModels.MenuResponse{
			Id:        res.Id,
			Title:     res.Title,
			Icon:      res.Icon,
			Path:      res.Path,
			ParentID:  res.ParentID,
			IsSection: res.IsSection,
			Submenu:   childMenus,
		})
	}

	return responses, nil
}

// GetKuisioner implements MenuServiceDefinition.
func (m MenuService) GetKuisioner(request menuModels.RequestKuisioner) (responses menuModels.ResponseData, err error) {
	roles, err := m.repo.LoadRole(menuModels.RequestKuisioner{
		TipeUker: request.TipeUker,
		Hilfm:    request.Hilfm,
		Orgeh:    request.Orgeh,
		Pernr:    request.Pernr,
		Kostl:    request.Kostl,
		StellTx:  request.StellTx,
		Jgpg:     request.Jgpg,
	})

	if err != nil {
		m.logger.Zap.Error(err)
		return responses, err
	}

	fmt.Println("Roles", roles)

	// Send To MQ Service
	baseUrl, err := lib.GetVarEnv("MQUrl")
	if err != nil {
		return responses, err
	}

	token := m.jwtService.CreateRealisasiToken(request.Pernr)

	Url := baseUrl + "menu/getMenuKusioner"

	headers := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}

	type Payload struct {
		Keyword string `json:"keyword"`
		IdRole  string `json:"id_role"`
		Limit   int64  `json:"limit"`
		Offset  int64  `json:"offset"`
		Pernr   string `json:"pernr"`
	}

	requestBody := &Payload{
		Keyword: request.Keyword,
		IdRole:  roles.Id,
		Limit:   request.Limit,
		Offset:  request.Offset,
		Pernr:   request.Pernr,
	}

	fmt.Println("Url =>", Url)
	fmt.Println("Headers =>", headers)
	fmt.Println("Request Body =>", requestBody)
	fmt.Println("Response =>", responses)

	err = lib.MakeRequest("POST", Url, headers, requestBody, &responses)

	return responses, err
}

// Menu Quisioner
func (menu MenuService) GetAll() (responses []menuModels.MenuQnaResponse, err error) {
	dataMenus, err := menu.repo.GetAll()
	if err != nil {
		return responses, err
	}

	for _, menu := range dataMenus {
		responses = append(responses, menuModels.MenuQnaResponse{
			ID:      menu.ID,
			Name:    menu.Name,
			Submenu: menu.Submenu,
		})
	}

	return responses, err
}

func (menu MenuService) SubMenuCheck(request menuModels.MenuQnaRequest) (response menuModels.MenuQnaResponse, err error) {
	dataMenu, err := menu.repo.SubMenuCheck(request)

	response = menuModels.MenuQnaResponse{
		ID:      dataMenu.ID,
		Name:    dataMenu.Name,
		Submenu: dataMenu.Submenu,
	}

	return response, err
}

func (menu MenuService) GetAllMstMenu() (responses []menuModels.MstMenu, err error) {
	dataMenus, err := menu.repo.GetAllMstMenu()
	if err != nil {
		return responses, err
	}

	for _, menu := range dataMenus {
		responses = append(responses, menuModels.MstMenu{
			IDMenu:      menu.IDMenu,
			Title:       menu.Title,
			IDParent:    menu.IDParent,
			ChildStatus: menu.ChildStatus,
			Urutan:      menu.Urutan,
		})
	}

	return responses, err
}

func (menu MenuService) DeleteMenuRRM(request menuModels.MstMenuRequest) (status bool, err error) {
	err = menu.repo.DeleteMenuRRM(request)

	if err != nil {
		return false, err
	}

	return true, err
}

func (menu MenuService) StoreMstRRM(request menuModels.MstMenu) (id string, err error) {
	tx := menu.db.DB.Begin()

	id, err = menu.repo.StoreMstRRM(request, tx)

	if err != nil {
		return id, err
	}

	tx.Commit()
	return id, err
}

func (menu MenuService) DeleteRole(request menuModels.MstMenuRequest) (err error) {
	roleRequest := types.Role{
		MenuID: request.IDMenu,
	}

	err = menu.repo.DeleteRole(&roleRequest)

	if err != nil {
		return err
	}

	return err
}

func (menu MenuService) StoreRoleRRM(request []types.Role) (err error) {
	for _, role := range request {
		err = menu.repo.StoreRoleRRM(&role)
	}

	if err != nil {
		return err
	}

	return err
}

func (menu MenuService) SetStatus(request menuModels.MstMenu) (response bool, err error) {
	response, err = menu.repo.SetStatus(&request)

	if err != nil {
		return response, err
	}

	return response, err
}

func (menu MenuService) GetLastID() (id_menu int64, err error) {
	id_menu, err = menu.repo.GetLastID()

	if err != nil {
		return id_menu, err
	}

	return id_menu, err
}
