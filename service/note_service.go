package service

import (
	"errors"
	"strings"

	"notes-api-v2/models"
	"notes-api-v2/store"
)

var ErrNotFound = errors.New("note not found")

type NoteService struct {
	store *store.MemoryStore
}

func NewNoteService(s *store.MemoryStore) *NoteService {
	return &NoteService{store: s}
}
func (s *NoteService) List() []models.Note {
	return s.store.GetAll()
}
func (s *NoteService) Get(id int) (models.Note, error) {
	note, err := s.store.GetByID(id)
	if err != nil && strings.Contains(err.Error(), "not found") {
		return models.Note{}, ErrNotFound
	}
	return note, err
}
func (s *NoteService) Create(input models.CreateNoteInput) (models.Note, error) {
	if err := input.Validate(); err != nil {
		return models.Note{}, err
	}
	return s.store.Create(input), nil
}
func (s *NoteService) Delete(id int) error {
	if err := s.store.Delete(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ErrNotFound
		}
		return err
	}
	return nil
}
func (s *NoteService) Update(id int, input models.CreateNoteInput) (models.Note, error) {
	if err := input.Validate(); err != nil {
		return models.Note{}, err
	}
	note, err := s.store.Update(id, input)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return models.Note{}, ErrNotFound
		}
		return models.Note{}, err
	}
	return note, nil
}
func (s *NoteService) CreateBulk(inputs []models.CreateNoteInput) ([]models.Note, error) {
	for _, input := range inputs {
		if err := input.Validate(); err != nil {
			return nil, err
		}
	}
	notes, err := s.store.CreateBulk(inputs)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
