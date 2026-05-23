package routes

import (
	"notes-api-v2/handlers"

	"github.com/go-chi/chi/v5"
)

func Mount(r chi.Router, notes *handlers.NoteHandler) {
	r.Get("/notes", notes.List)
	r.Post("/notes", notes.Create)
	r.Get("/notes/{id}", notes.Get)
	r.Delete("/notes/{id}", notes.Delete)
	r.Put("/notes/{id}", notes.Update)
	r.Post("/notes/bulk", notes.CreateBulk)
}
