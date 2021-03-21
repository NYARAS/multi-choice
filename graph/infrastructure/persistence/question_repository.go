package persistence

import (
	"errors"
	"strings"

	"github.com/NYARAS/multi-choice/graph/domain/repository/question"
	"github.com/NYARAS/multi-choice/graph/models"

	"gorm.io/gorm"
)

type questionService struct {
	db *gorm.DB
}

func NewQuestion(db *gorm.DB) *questionService {
	return &questionService{
		db,
	}
}

// We implement the interface defined in the domain
var _ question.QuestionService = &questionService{}

func (s *questionService) CreateQuestion(question *models.Question) (*models.Question, error) {
	err := s.db.Create(&question).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("question title already taken")
		}
		return nil, err
	}
	return question, nil
}

func (s *questionService) UpdateQuestion(question *models.Question) (*models.Question, error) {
	err := s.db.Save(&question).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("question title already taken")
		}
		return nil, err
	}
	return question, nil
}

func (s *questionService) DeleteQuestion(id string) error {
	ques := &models.Question{}
	err := s.db.Where("id = ?", id).Delete(&ques).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *questionService) GetQuestionByID(id string) (*models.Question, error) {
	ques := &models.Question{}
	err := s.db.Where("id = ?", id).Preload("QuestionOption").Take(&ques).Error
	if err != nil {
		return nil, err
	}
	return ques, nil
}

func (s *questionService) GetAllQuestions() ([]*models.Question, error) {
	var questions []*models.Question
	err := s.db.Preload("QuestionOption").Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}
