package questions

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/undeadops/enigma/pkg/config"
	"github.com/undeadops/enigma/pkg/models"
)

// Handler - Context for Handling of Questions routes
type Handler struct {
	repo models.QuestionsData
	qset models.QuestionSetData
}

// NewHandler - Initialize Handler
func NewHandler(q models.QuestionsData, qs models.QuestionSetData) *Handler {
	return &Handler{
		repo: q,
		qset: qs,
	}
}

// Router - A completely separate router the questions data storage handle
func Router(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(authHandler)

	r.Get("/", config.HoneycombMiddleware(h.GetResponses))
	r.Get("/{id:[0-9a-f]+}", config.HoneycombMiddleware(h.GetResponseID))
	r.Post("/", config.HoneycombMiddleware(h.SaveResponse))
	r.Put("/{id:[0-9a-f]+}", config.HoneycombMiddleware(h.UpdateResponse))
	r.Delete("/{id:[0-9a-f]+}", config.HoneycombMiddleware(h.DeleteResponse))

	r.Get("/sets", config.HoneycombMiddleware(h.GetQuestionSet))
	r.Post("/sets", config.HoneycombMiddleware(h.SaveQuestionSet))
	r.Delete("/{id:[0-9a-f]+", config.HoneycombMiddleware(h.DeleteQuestionSet))

	return r
}
