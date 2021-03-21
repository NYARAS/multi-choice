package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/NYARAS/multi-choice/graph/models"
)

func (r *mutationResolver) CreateAnswer(ctx context.Context, questionID string, optionID string) (*models.AnswerResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAnswer(ctx context.Context, id string, questionID string, optionID string) (*models.AnswerResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAnswer(ctx context.Context, id string) (*models.AnswerResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOneAnswer(ctx context.Context, id string) (*models.AnswerResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllQuestionAnswer(ctx context.Context, questionID string) (*models.AnswerResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
