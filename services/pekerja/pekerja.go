package pekerja

import (
	"riskmanagement/repository/pekerja"

	models "riskmanagement/models/pekerja"

	"gitlab.com/golang-package-library/logger"
)

type PekerjaDefinition interface {
	GetApproval(request *models.RequestApproval) (responses []models.DataPekerjaResponse, err error)
	GetAllPekerjaBranch(request *models.PekerjaUkerRequest) (responses []models.DataPekerjaResponse, err error)
}

type PekerjaService struct {
	logger logger.Logger
	repo   pekerja.PekerjaDefinition
}

func NewPekerjaService(
	logger logger.Logger,
	repo pekerja.PekerjaDefinition,
) PekerjaDefinition {
	return PekerjaService{
		logger: logger,
		repo:   repo,
	}
}

// GetAllPekerjaBranch implements PekerjaDefinition.
func (p PekerjaService) GetAllPekerjaBranch(request *models.PekerjaUkerRequest) (responses []models.DataPekerjaResponse, err error) {
	dataPekerja, err := p.repo.GetAllPekerjaBranch(request)

	if err != nil {
		p.logger.Zap.Error("Error getting data")
		return responses, err
	}

	for _, value := range dataPekerja {
		responses = append(responses, models.DataPekerjaResponse{
			Pernr:   value.Pernr,
			Sname:   value.Sname,
			Stell:   value.Stell,
			StellTx: value.StellTx,
			Branch:  value.Branch,
		})
	}

	return responses, nil
}

// GetApproval implements PekerjaDefinition.
func (p PekerjaService) GetApproval(request *models.RequestApproval) (responses []models.DataPekerjaResponse, err error) {
	dataPekerja, err := p.repo.GetApproval(request)

	if err != nil {
		p.logger.Zap.Error("Error getting data")
		return responses, err
	}

	for _, value := range dataPekerja {
		responses = append(responses, models.DataPekerjaResponse{
			Pernr:   value.Pernr,
			Sname:   value.Sname,
			StellTx: value.StellTx,
			Branch:  value.Branch,
		})
	}

	return responses, nil
}
