package questions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/undeadops/enigma/pkg/models"
)

// SaveResponse - write out question to database
func (h *Handler) SaveResponse(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	post := &models.Response{}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ev.AddField("save.post", post.ID)

	queryStart := time.Now()
	err = h.repo.SaveResponse(r.Context(), post)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error"+err.Error())
	}
	ev.AddField("timers.db.response_insert_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// UpdateResponse - PUT question updated with additional answers
func (h *Handler) UpdateResponse(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	id := chi.URLParam(r, "id")

	post := models.Response{}
	json.NewDecoder(r.Body).Decode(&post)

	queryStart := time.Now()
	_, err = h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not update question")
	}
	ev.AddField("timers.db.response_update_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Question Updated"})
}

// GetResponseID - Fetch a question from givin ID in url
func (h *Handler) GetResponseID(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	id := chi.URLParam(r, "id")

	ev.AddField("response.id", id)

	queryStart := time.Now()
	result, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not Fetch Response")
	}
	ev.AddField("timers.db.response_get_id_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, result)
}

// GetResponses - GET list of all questions/responses
func (h *Handler) GetResponses(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	queryStart := time.Now()
	results, err := h.repo.ListResponses(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}
	ev.AddField("timers.db.response_get_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, results)
}

// DeleteResponse - DELETE question id from questions
func (h *Handler) DeleteResponse(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	id := chi.URLParam(r, "id")

	ev.AddField("response.id", id)

	queryStart := time.Now()
	err = h.repo.DeleteResponse(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error Deleting ID")
	}
	ev.AddField("timers.db.response_delete_ms", time.Since(queryStart)/time.Millisecond)

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
