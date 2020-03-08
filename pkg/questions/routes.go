package questions

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/undeadops/enigma"
)

// Routes - Setup HTTP Routes for Qeustions pkg
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", ListQuestions)
	return router
}

// ListQuestions - Return a list of all questions answered
func ListQuestions(w http.ResponseWriter, r *http.Request) {
	q := enigma.Response{
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

	render.JSON(w, r, q)
}
