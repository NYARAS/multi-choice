package answer

import (
	"github.com/NYARAS/multi-choice/graph/models"
)

type AnswerService interface {
	CreateAnswer(answer *models.Answer) (*models.Answer, error)
	UpdateAnswer(answer *models.Answer) (*models.Answer, error)
	DeleteAnswer(id string) error
	GetAnswerByID(id string) (*models.Answer, error)
	GetAllQuestionAnswers(questionID string) (*models.Answer, error)
}
