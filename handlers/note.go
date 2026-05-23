package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"notes-api-v2/models"
	"notes-api-v2/service"

	"github.com/go-chi/chi/v5"
)

type NoteHandler struct {
	svc *service.NoteService
}

func NewNoteHandler(svc *service.NoteService) *NoteHandler {
	return &NoteHandler{svc: svc}
}
func (h *NoteHandler) List(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, h.svc.List())
}
func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.CreateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "bad JSON")
		return
	}
	note, err := h.svc.Create(input)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, note)
}
func (h *NoteHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	note, err := h.svc.Get(id)
	if errors.Is(err, service.ErrNotFound) {
		writeError(w, http.StatusNotFound, "note not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to get note")
		return
	}
	writeJSON(w, http.StatusOK, note)
}
func (h *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.svc.Delete(id); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			writeError(w, http.StatusNotFound, "note not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "delete failed")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var input models.CreateNoteInput
	/*if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "bad JSON")
		return
	}*/
	if _, err := writeJSONDecoder(w, r, &input); err != nil {
		return
	}
	note, err := h.svc.Update(id, input)
	if errors.Is(err, service.ErrNotFound) {
		writeError(w, http.StatusNotFound, "note not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "update failed")
		return
	}

	writeJSON(w, http.StatusOK, note)
}
