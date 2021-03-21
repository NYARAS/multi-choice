package question_option

import (
	"github.com/NYARAS/multi-choice/graph/models"
)

type OptService interface {
	CreateQuestionOption(question *models.QuestionOption) (*models.QuestionOption, error)
	UpdateQuestionOption(question *models.QuestionOption) (*models.QuestionOption, error)
	DeleteQuestionOption(id string) error
	DeleteQuestionOptionByQuestionID(questionID string) error
	GetQuestionOptionByID(id string) (*models.QuestionOption, error)
	GetQuestionOptionByQuestionID(questionID string) ([]*models.QuestionOption, error)
}
