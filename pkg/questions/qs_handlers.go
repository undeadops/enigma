package questions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/undeadops/enigma/pkg/models"
)

// SaveQuestionSet - Save Set of Questions to database
func (h *Handler) SaveQuestionSet(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	post := &models.QuestionSet{}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ev.AddField("save.questionset", post.ID)

	queryStart := time.Now()
	err = h.qset.SaveQuestionSet(r.Context(), post)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error"+err.Error())
	}
	ev.AddField("timers.db.questionset_insert_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// GetQuestionSet - GET list of all questions
func (h *Handler) GetQuestionSet(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	queryStart := time.Now()
	results, err := h.qset.ListQuestionSet(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}
	ev.AddField("timers.db.questionset_get_ms", time.Since(queryStart)/time.Millisecond)

	ev.AddField("get.questionset_return_count", len(results))

	respondWithJSON(w, http.StatusOK, results)
}

// DeleteQuestionSet - DELETE question set id from question sets
func (h *Handler) DeleteQuestionSet(w http.ResponseWriter, r *http.Request) {
	var err error
	ev := hnyEventFromRequest(r)
	defer addFinalErr(&err, ev)

	id := chi.URLParam(r, "id")

	ev.AddField("response.id", id)

	queryStart := time.Now()
	err = h.qset.DeleteQuestionSet(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error Deleting ID")
	}
	ev.AddField("timers.db.questionset_delete_ms", time.Since(queryStart)/time.Millisecond)

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Deleted"})
}
