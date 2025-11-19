package jenistask

import (
	models "riskmanagement/models/jenistask"
	repo "riskmanagement/repository/jenistask"

	"gitlab.com/golang-package-library/logger"
)

type JenisTaskDefinition interface {
	GetData(request *models.TaskRequest) (response []models.JenisTask, err error)
}

type JenisTaskService struct {
	logger    logger.Logger
	jenisTask repo.JenisTaskDefinition
}

func NewJenisTaskService(
	logger logger.Logger,
	jenisTask repo.JenisTaskDefinition,
) JenisTaskDefinition {
	return JenisTaskService{
		logger:    logger,
		jenisTask: jenisTask,
	}
}

func (s JenisTaskService) GetData(request *models.TaskRequest) (response []models.JenisTask, err error) {
	response, err = s.jenisTask.GetData(request)
	if err != nil {
		s.logger.Zap.Error(err)
	}

	return response, err
}
