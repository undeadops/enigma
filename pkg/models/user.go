package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserData - Storage Abstraction for user objects
type UserData interface {
	CreateUser(user *User) error
	GetUser(user string) (*User, error)
	DeleteUser(user string) error
	ListUsers() ([]*User, error)
}

// User - Authenticated User that is making these responses
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username" bson:"username"`
	PhoneNum string             `json:"phone_num" bson:"phone_num"`
}
