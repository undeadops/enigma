package questions

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/undeadops/enigma/pkg/models"
)

// Handler - Context for Handling of Questions routes
type Handler struct {
	repo models.QuestionsData
}

// NewHandler - Initialize Handler
func NewHandler(q models.QuestionsData) *Handler {
	return &Handler{
		repo: q,
	}
}

// Router - A completely separate router the questions data storage handle
func Router(h *Handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", h.GetResponses)
	r.Get("/{id:[0-9]+}", h.GetResponseID)
	r.Post("/", h.SaveResponse)
	r.Put("/{id:[0-9]+}", h.UpdateResponse)
	r.Delete("/{id:[0-9]+}", h.DeleteResponse)
	return r
}
