package unitkerja

import (
	"fmt"
	models "riskmanagement/models/unitkerja"
	repository "riskmanagement/repository/unitkerja"

	"gitlab.com/golang-package-library/logger"
)

type UnitKerjaDefinition interface {
	GetAll() (responses []models.UnitKerjaResponse, err error)
	GetOne(id int64) (responses models.UnitKerjaResponse, err error)
	Store(request *models.UnitKerjaRequest) (err error)
	Update(request *models.UnitKerjaRequest) (err error)
	Delete(id int64) (err error)
	GetRegionList(request models.RegionRequest) (responses []models.RegionList, err error)
	GetMainbrList(request models.MainbrRequest) (responses []models.MainbrList, err error)
	GetBranchList(request models.BranchRequest) (responses []models.BranchList, err error)

	GetRegionName(REGION string) (name string, err error)
	GetMainbrName(MAINBR string) (name string, err error)
	GetBranchName(BRANCH string) (name string, err error)
	GetMainbrKWList(request models.MainbrKWRequest) (responses []models.MainbrList, err error)
	GetEmployeeRegion(request models.EmployeeRegionRequest) (responses models.EmployeeRegionResponse, err error)

	// DisasterMaps
	GetMapRegionList(request *models.MapLocationRequest) (response []models.MapRegionOffice, err error)
	GetMapBranchList(request *models.MapLocationRequest) (response []models.MapBranchOffice, err error)
	GetMapUnitList(request *models.MapLocationRequest) (response []models.MapUnitOffice, err error)
}

type UnitKerjaService struct {
	logger     logger.Logger
	repository repository.UnitKerjaDefinition
}

func NewUnitKerjaService(
	logger logger.Logger,
	repository repository.UnitKerjaDefinition,
) UnitKerjaDefinition {
	return UnitKerjaService{
		logger:     logger,
		repository: repository,
	}
}

// Delete implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) Delete(id int64) (err error) {
	return unitKerja.repository.Delete(id)
}

// GetAll implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) GetAll() (responses []models.UnitKerjaResponse, err error) {
	return unitKerja.repository.GetAll()
}

// GetOne implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) GetOne(id int64) (responses models.UnitKerjaResponse, err error) {
	return unitKerja.repository.GetOne(id)
}

// Store implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) Store(request *models.UnitKerjaRequest) (err error) {
	fmt.Println("service =", request)
	status, err := unitKerja.repository.Store(request)
	if !status || err != nil {
		return err
	}

	return nil
}

// Update implements UnitKerjaDefinition
func (unitKerja UnitKerjaService) Update(request *models.UnitKerjaRequest) (err error) {
	status, err := unitKerja.repository.Update(request)
	if !status || err != nil {
		return err
	}

	return nil
}

func (unitKerja UnitKerjaService) GetRegionList(request models.RegionRequest) (responses []models.RegionList, err error) {
	responses, err = unitKerja.repository.GetRegionList(&request)

	return responses, err
}
func (unitKerja UnitKerjaService) GetMainbrList(request models.MainbrRequest) (responses []models.MainbrList, err error) {
	responses, err = unitKerja.repository.GetMainbrList(&request)

	return responses, err
}
func (unitKerja UnitKerjaService) GetBranchList(request models.BranchRequest) (responses []models.BranchList, err error) {
	responses, err = unitKerja.repository.GetBranchList(&request)

	return responses, err
}

func (unitKerja UnitKerjaService) GetRegionName(REGION string) (name string, err error) {
	var uker models.UkerName
	uker, err = unitKerja.repository.GetRegionName(REGION)

	return uker.BRDESC, err
}
func (unitKerja UnitKerjaService) GetMainbrName(MAINBR string) (name string, err error) {
	var uker models.UkerName
	uker, err = unitKerja.repository.GetMainbrName(MAINBR)

	return uker.BRDESC, err
}
func (unitKerja UnitKerjaService) GetBranchName(BRANCH string) (name string, err error) {
	var uker models.UkerName
	uker, err = unitKerja.repository.GetBranchName(BRANCH)

	return uker.BRDESC, err
}

func (unitKerja UnitKerjaService) GetMainbrKWList(request models.MainbrKWRequest) (responses []models.MainbrList, err error) {
	responses, err = unitKerja.repository.GetMainbrKWList(&request)

	return responses, err
}

func (unitKerja UnitKerjaService) GetEmployeeRegion(request models.EmployeeRegionRequest) (responses models.EmployeeRegionResponse, err error) {
	return unitKerja.repository.GetEmployeeRegion(&request)
}

// GetMapRegionList implements UnitKerjaDefinition.
func (uk UnitKerjaService) GetMapRegionList(request *models.MapLocationRequest) (response []models.MapRegionOffice, err error) {
	data, err := uk.repository.GetMapRegionList(request)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetMapBranchList implements UnitKerjaDefinition.
func (uk UnitKerjaService) GetMapBranchList(request *models.MapLocationRequest) (response []models.MapBranchOffice, err error) {
	data, err := uk.repository.GetMapBranchList(request)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (uk UnitKerjaService) GetMapUnitList(request *models.MapLocationRequest) (response []models.MapUnitOffice, err error) {
	data, err := uk.repository.GetMapUnitList(request)

	if err != nil {
		return nil, err
	}

	return data, nil
}
