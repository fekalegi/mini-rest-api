package note

import (
	"mini-rest-api/repository"
	"mini-rest-api/usecase"
)

type noteImplementation struct {
	repo repository.NoteRepository
}

func NewNoteImplementation(repo repository.NoteRepository) usecase.NoteService {
	return &noteImplementation{
		repo: repo,
	}
}
