package main

import (
	"fmt"
	"net/http"

	"notes-api-v2/handlers"
	"notes-api-v2/routes"
	"notes-api-v2/service"
	"notes-api-v2/store"

	"github.com/go-chi/chi/v5"
)

func main() {
	noteStore := store.NewMemoryStore()
	noteSvc := service.NewNoteService(noteStore)
	noteH := handlers.NewNoteHandler(noteSvc)

	r := chi.NewRouter()
	routes.Mount(r, noteH)

	fmt.Println("📝 Notes API v2 (routes · service · handlers) on :8080")
	http.ListenAndServe(":8080", r)
}
