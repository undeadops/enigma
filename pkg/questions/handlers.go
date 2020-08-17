package questions

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/undeadops/enigma/pkg/models"
)

// SaveResponse - write out question to database
func (h *Handler) SaveResponse(w http.ResponseWriter, r *http.Request) {
	post := &models.Response{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.SaveResponse(r.Context(), post)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error"+err.Error())
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// UpdateResponse - PUT question updated with additional answers
func (h *Handler) UpdateResponse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	post := models.Response{}
	json.NewDecoder(r.Body).Decode(&post)

	_, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not update question")
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Question Updated"})
}

// GetResponseID - Fetch a question from givin ID in url
func (h *Handler) GetResponseID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not Fetch Response")
	}

	respondWithJSON(w, http.StatusOK, result)
}

// GetResponses - GET list of all questions/responses
func (h *Handler) GetResponses(w http.ResponseWriter, r *http.Request) {

	results, err := h.repo.ListResponses(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondWithJSON(w, http.StatusOK, results)
}

// DeleteResponse - DELETE question id from questions
func (h *Handler) DeleteResponse(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.repo.DeleteResponse(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error Deleting ID")
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Deleted"})
}

// respondwithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
