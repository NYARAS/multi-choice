// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type Answer struct {
	ID         string    `json:"id" db:"id"`
	QuestionID string    `json:"questionID" db:"questionID"`
	OptionID   string    `json:"optionID" db:"optionID"`
	IsCorrect  bool      `json:"isCorrect" db:"isCorrect"`
	CreatedAt  time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updatedAt"`
}

type AnswerResponse struct {
	Message  string    `json:"message" db:"message"`
	Status   int       `json:"status" db:"status"`
	Data     *Answer   `json:"data" db:"data"`
	DataList []*Answer `json:"dataList" db:"dataList"`
}

type Question struct {
	ID             string            `json:"id" db:"id"`
	Title          string            `json:"title" gorm:"unique" db:"title"`
	QuestionOption []*QuestionOption `json:"questionOption" db:"questionOption"`
	CreatedAt      time.Time         `json:"createdAt" db:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt" db:"updatedAt"`
}

type QuestionInput struct {
	Title   string                 `json:"title" db:"title"`
	Options []*QuestionOptionInput `json:"options" db:"options"`
}

type QuestionOption struct {
	ID         string    `json:"id" db:"id"`
	QuestionID string    `json:"questionID" db:"questionID"`
	Title      string    `json:"title" db:"title"`
	Position   int       `json:"position" db:"position"`
	IsCorrect  bool      `json:"isCorrect" db:"isCorrect"`
	CreatedAt  time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updatedAt"`
}

type QuestionOptionInput struct {
	Title     string `json:"title" db:"title"`
	Position  int    `json:"position" db:"position"`
	IsCorrect bool   `json:"isCorrect" db:"isCorrect"`
}

type QuestionResponse struct {
	Message  string      `json:"message" db:"message"`
	Status   int         `json:"status" db:"status"`
	Data     *Question   `json:"data" db:"data"`
	DataList []*Question `json:"dataList" db:"dataList"`
}
