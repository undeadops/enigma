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

// QuestionSetData - Storage Abstraction for what questions to ask
type QuestionSetData interface {
	SaveQuestionSet(ctx context.Context, qs *QuestionSet) error
	DeleteQuestionSet(ctx context.Context, qs string) error
	ListQuestionSet(ctx context.Context) ([]*QuestionSet, error)
	GetQuestionSet(ctx context.Context, qsID string) (*QuestionSet, error)
}

// QuestionSet - Set of Questions to ask
type QuestionSet struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Date        time.Time          `bson:"date" json:"date"`
	QuestionSet []string           `bson:"question_set" json:"question_set"`
}
