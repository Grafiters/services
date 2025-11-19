package questionner

import (
	models "riskmanagement/models/questionner"
	questionRepo "riskmanagement/repository/questionner"

	"gitlab.com/golang-package-library/logger"
)

type QuestionnerDefinition interface {
	GetQuestionner() (responses []models.QuestionerList, err error)
}

type QuestionerService struct {
	logger       logger.Logger
	questionRepo questionRepo.QuestionnerDefinition
}

func NewQuestionnerService(
	logger logger.Logger,
	questionRepo questionRepo.QuestionnerDefinition,
) QuestionnerDefinition {
	return QuestionerService{
		logger:       logger,
		questionRepo: questionRepo,
	}
}

// GetQuestionner implements QuestionnerDefinition
func (q QuestionerService) GetQuestionner() (responses []models.QuestionerList, err error) {
	questionnerList, err := q.questionRepo.GetQuestion()

	for _, res := range questionnerList {
		Answer, err := q.questionRepo.GetAnswer(res.Id)
		if err != nil {
			q.logger.Zap.Error(err)
			return responses, err
		}

		responses = append(responses, models.QuestionerList{
			Id:       res.Id,
			Question: res.Question,
			Answer:   Answer,
		})
	}

	return responses, err
}
