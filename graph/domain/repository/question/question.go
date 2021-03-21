package question

import "github.com/NYARAS/multi-choice/graph/models"

type QuestionService interface {
	CreateQuestion(question *models.Question) (*models.Question, error)
	UpdateQuestion(question *models.Question) (*models.Question, error)
	DeleteQuestion(id string) error
	GetQuestionByID(id string) (*models.Question, error)
	GetAllQuestions() ([]*models.Question, error)
}
