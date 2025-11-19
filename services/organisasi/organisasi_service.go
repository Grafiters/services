package organisasi

import (
	models "riskmanagement/models/organisasi"
	repo "riskmanagement/repository/organisasi"
)

type OrganisasiServiceDefinition interface {
	GetCostCenter(request models.CostCenterRequest) (responses []models.CostCenterResponse, err error)
	GetOrgUnit(request models.DepartmentRequest) (responses []models.DepartmentResponse, err error)
	GetHilfm(request models.JabatanRequest) (responses []models.JabatanResponse, err error)
}

type OrganisasiService struct {
	repo repo.OrganisasiDefinition
}

func NewOrganisasiService(
	repo repo.OrganisasiDefinition,
) OrganisasiServiceDefinition {
	return OrganisasiService{
		repo: repo,
	}
}

// GetCostCenter implements OrganisasiServiceDefinition.
func (o OrganisasiService) GetCostCenter(request models.CostCenterRequest) (responses []models.CostCenterResponse, err error) {
	data, err := o.repo.GetCostCenter(request)

	return data, err
}

// GetOrgUnit implements OrganisasiServiceDefinition.
func (o OrganisasiService) GetOrgUnit(request models.DepartmentRequest) (responses []models.DepartmentResponse, err error) {
	data, err := o.repo.GetOrgUnit(request)

	return data, err
}

// GetJabatan implements OrganisasiServiceDefinition.
func (o OrganisasiService) GetHilfm(request models.JabatanRequest) (responses []models.JabatanResponse, err error) {
	data, err := o.repo.GetHilfm(request)

	return data, err
}
