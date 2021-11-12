package services

import (
	"repository/pkg/models"
	"repository/pkg/repositories"
)

type NoteService struct {
	NoteRepository repositories.NoteRepositoryInterface
}

func (service NoteService) GetNotes() []models.Note {
	return service.NoteRepository.Get()
}

func (service NoteService) GetNote(id int) models.Note {
	return service.NoteRepository.GetOne(id)
}

func (service NoteService) SaveNote(note models.Note) string {
	return service.NoteRepository.Save(note)
}
