package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NYARAS/multi-choice/graph/generated"
	"github.com/NYARAS/multi-choice/graph/models"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, question models.QuestionInput) (*models.QuestionResponse, error) {
	// vaidate the title
	if question.Title == "" {
		return &models.QuestionResponse{
			Message: "The title is required",
			Status:  http.StatusBadRequest,
		}, nil
	}

	ques := &models.Question{
		Title: question.Title,
	}
	ques.CreatedAt = time.Now()
	ques.UpdatedAt = time.Now()

	// save the question
	quest, err := r.QuestionService.CreateQuestion(ques)
	if err != nil {
		fmt.Println("the error with this", err)
		return &models.QuestionResponse{
			Message: error.Error(),
			Status:  http.StatusInternalServerError,
		}, nil
	}

	// validate the question options
	for _, v := range question.Options {
		if ok, errorString := helpers.ValidateInputs(*v); !ok {
			return &models.QuestionResponse{
				Message: errorString,
				Status:  http.StatusUnprocessableEntity,
			}, nil
		}
		questOpt := &models.QuestionOption{
			QuestionID: ques.ID,
			Title:      v.Title,
			Position:   v.Position,
			IsCorrect:  v.IsCorrect,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		_, err := r.QuestionOptionService.CreateQuestionOption(questOpt)
		if err != nil {
			return &models.QuestionResponse{
				Message: "Error creating question option",
				Status:  http.StatusInternalServerError,
			}, nil
		}
	}
	return &models.QuestionResponse{
		Message: "Successfully created question",
		Status:  http.StatusCreated,
		Data:    quest,
	}, nil
}

func (r *mutationResolver) UpdateQuestion(ctx context.Context, id string, question models.QuestionInput) (*models.QuestionResponse, error) {
	// validate the title
	if question.Title == "" {
		return &models.QuestionResponse{
			Message: "The title is required",
			Status:  http.StatusBadRequest,
		}, nil
	}

	// get the question
	ques, err := r.QuestionService.GetQuestionByID(id)
	if err != nil {
		return &models.QuestionResponse{
			Message: "Error getting the question",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	ques.Title = question.Title
	ques.UpdatedAt = time.Now()

	// save the question
	quest, err := r.QuestionService.UpdateQuestion(ques)
	if err != nil {
		return &models.QuestionResponse{
			Message: "Error creating the question",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	// For the options, we will discard the previous options and insert
	// the new ones
	err = r.QuestionOptionService.DeleteQuestionOptionByQuestionID(quest.ID)
	if err != nil {
		return &models.QuestionResponse{
			Message: "Error deleting the question options",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	for _, v := range question.Options {
		if ok, errorString := helpers.ValidateInputs(*v); !ok {
			return &models.QuestionResponse{
				Message: errorString,
				Status:  http.StatusUnprocessableEntity,
			}, nil
		}

		quesOpt := &models.QuestionOption{
			QuestionID: quest.ID,
			Title:      v.Title,
			Position:   v.Position,
			IsCorrect:  v.IsCorrect,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		_, err := r.QuestionService.CreateQuestionOption(quesOpt)
		if err != nil {
			return &models.QuestionResponse{
				Message: "Error creating the question option",
				Status:  http.StatusInternalServerError,
			}, nil
		}
	}

	return &models.QuestionResponse{
		Message: "Successfully updated the question",
		Status:  http.StatusOK,
		Data:    quest,
	}, nil
}

func (r *mutationResolver) DeleteQuestion(ctx context.Context, id string) (*models.QuestionResponse, error) {
	err := r.QuestionService.DeleteQuestion(id)
	if err != nil {
		return &models.QuestionResponse{
			Message: "Something went wrong deleting the question",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	//also delete the options created too

	err = r.QuestionOptionService.DeleteQuestionOptionByQuestionID(id)
	if err != nil {
		return &models.QuestionResponse{
			Message: "Error Deleting the question options",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.QuestionResponse{
		Message: "Successfully deleted the question",
		Status:  http.StatusOK,
	}, nil
}

func (r *queryResolver) GetOneQuestion(ctx context.Context, id string) (*models.QuestionResponse, error) {
	question, err := r.QuestionService.GetQuestionByID(id)
	if err != nil {
		log.Println("getting question error:", err)
		return &models.QuestionResponse{
			Message: "Something went wrong getting the question",
			Status:  http.StatusInternalServerError,
		}, nil
	}
	return &models.QuestionResponse{
		Message: "Successfully retrieve the question",
		Status:  http.StatusOK,
		Data:    question,
	}, nil
}

func (r *queryResolver) GetAllQuestions(ctx context.Context) (*models.QuestionResponse, error) {
	questions, err := r.QuestionService.GetAllQuestions()
	if err != nil {
		log.Println("getting a;; questions error:", err)
		return &models.QuestionResponse{
			Message: "Something went wrong getting all the questions",
			Status:  http.StatusInternalServerError,
		}, nil
	}
	return &models.QuestionResponse{
		Message:  "Successfully retrived all the questions",
		Status:   http.StatusOK,
		DataList: questions,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
