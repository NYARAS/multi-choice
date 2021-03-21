package interfaces

import (
	"github.com/NYARAS/multi-choice/graph/domain/repository/answer"
	"github.com/NYARAS/multi-choice/graph/domain/repository/question"
	"github.com/NYARAS/multi-choice/graph/domain/repository/question_option"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AnsService            answer.AnswerService
	QuestionService       question.QuestionService
	QuestionOptionService question_option.OptService
}
