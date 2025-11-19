package questionner

import (
	"riskmanagement/lib"
	models "riskmanagement/models/questionner"
	"time"

	"gitlab.com/golang-package-library/logger"
	"gorm.io/gorm"
)

type QuestionnerDefinition interface {
	WithTrx(trxHandle *gorm.DB) QuestionnerRepository
	GetQuestion() (responses []models.Question, err error)
	GetAnswer(id int64) (responses []models.Answer, err error)
}

type QuestionnerRepository struct {
	db      lib.Database
	logger  logger.Logger
	timeout time.Duration
}

func NewQuestionnerRepository(
	db lib.Database,
	logger logger.Logger,
) QuestionnerDefinition {
	return QuestionnerRepository{
		db:      db,
		logger:  logger,
		timeout: 0,
	}
}

// GetAnswer implements QuestionnerDefinition
func (q QuestionnerRepository) GetAnswer(id int64) (responses []models.Answer, err error) {
	return responses, q.db.DB.Table("tbl_answers").Where(`id_question = ?`, id).Find(&responses).Error
}

// GetQuestion implements QuestionnerDefinition
func (q QuestionnerRepository) GetQuestion() (responses []models.Question, err error) {
	return responses, q.db.DB.Table("tbl_questions").Find(&responses).Error
}

// WithTrx implements QuestionnerDefinition
func (q QuestionnerRepository) WithTrx(trxHandle *gorm.DB) QuestionnerRepository {
	if trxHandle == nil {
		q.logger.Zap.Error("transaction Database not found in gin context.")
		return q
	}

	q.db.DB = trxHandle
	return q
}
