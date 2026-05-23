package store

import (
	"fmt"
	"sync"
	"time"

	"notes-api-v2/models"
)

//has three fields: notes, nextID, mu(a lock)

type MemoryStore struct {
	mu     sync.RWMutex
	notes  map[int]models.Note
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{notes: make(map[int]models.Note), nextID: 1}
}

func (s *MemoryStore) GetAll() []models.Note {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []models.Note{}
	for _, n := range s.notes {
		out = append(out, n)
	}
	return out
}

func (s *MemoryStore) GetByID(id int) (models.Note, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n, ok := s.notes[id]
	if !ok {
		return models.Note{}, fmt.Errorf("note %d not found", id)
	}
	return n, nil
}
func (s *MemoryStore) Create(input models.CreateNoteInput) models.Note {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := models.Note{
		ID: s.nextID, Title: input.Title, Body: input.Body, CreatedAt: time.Now(),
	}
	s.notes[s.nextID] = n
	s.nextID++
	return n
}
func (s *MemoryStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.notes[id]; !ok {
		return fmt.Errorf("note %d not found", id)
	}
	delete(s.notes, id)
	return nil
}
