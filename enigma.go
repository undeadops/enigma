package enigma

import (
	"time"
)

// Domains data types?... I think thats what this would be called

// Response - Answers to daily questions
type Response struct {
	Date      time.Time  `json:"date" bson:"date"`
	Questions []Question `json:"questions" bson:"questions"`
}

// Question - Question And Answers
type Question struct {
	Date     time.Time `json:"date" bson:"date"`
	Question string    `json:"question" bson:"question"`
	Answer   string    `json:"answer" bson:"answer"`
}

// User - Authenticated User that is making these responses
type User struct {
	Username string `json:"username" bson:"username"`
	PhoneNum string `json:"phone_num" bson:"phone_num"`
}
