package models

//go:generate moq -out mockQuestionsData_test.go . QuestionsData
//go:generate moq -out mockUserData_test.go . UserData

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// QuestionsData - Storage Abstraction for question objects
type QuestionsData interface {
	SaveResponse(ctx context.Context, reps *Response) error
	DeleteResponse(ctx context.Context, respID string) error
	ListResponses(ctx context.Context) ([]*Response, error)
	GetByID(ctx context.Context, qID string) (*Response, error)
}

// Response - Answers to daily questions
type Response struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Date      time.Time          `json:"date" bson:"date"`
	Questions []Question         `json:"questions" bson:"questions"`
}

// Question - Question And Answers
type Question struct {
	Question string `json:"question" bson:"question"`
	Answer   string `json:"answer" bson:"answer"`
}
