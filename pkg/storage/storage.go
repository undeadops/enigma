package storage

//go:generate moq -out mockQuestionsData_test.go . QuestionsData
//go:generate moq -out mockUserData_test.go . UserData

import (
	"github.com/undeadops/enigma"
)

// QuestionsData - Storage Abstraction for question objects
type QuestionsData interface {
	SaveResponse(*enigma.Response) error
	ListResponses() ([]*enigma.Response, error)
}

// UserData - Storage Abstraction for user objects
type UserData interface {
	CreateUser(*enigma.User) error
	GetUser(string) (*enigma.User, error)
	DeleteUser(string) error
	ListUsers() ([]*enigma.User, error)
}
