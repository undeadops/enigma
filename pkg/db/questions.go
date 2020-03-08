package db

import (
	"time"

	"github.com/undeadops/enigma"
)

// SaveResponse - Save Reponse to database
func (qr *QuestionsRepo) SaveResponse(r *enigma.Response) error {

	return nil
}

// ListResponses - Return a List of Responses
func (qr *QuestionsRepo) ListResponses() ([]*enigma.Response, error) {
	q := &enigma.Response{
		Date: time.Now(),
		Questions: []enigma.Question{
			enigma.Question{
				Question: "What have you eaten thus far?",
				Answer:   "Half a costco muffin, some pistacios",
			},
			enigma.Question{
				Question: "How are you feeling?",
				Answer:   "Fine, not working so I'm not feeling stress from that",
			},
			enigma.Question{
				Question: "Read anything lately?",
				Answer:   "Golang articles",
			},
		},
	}

	resp := []*enigma.Response{
		q,
	}
	return resp, nil
}
