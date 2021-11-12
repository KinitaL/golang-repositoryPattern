package repositories

import (
	"repository/pkg/models"
)

type SqlReposiroty struct{}

func (SqlReposiroty) Get() []models.Note {
	db := connectToDB()
	db.AutoMigrate(&models.Note{})
	var notes []models.Note
	db.Find(&notes)
	return notes
}

func (SqlReposiroty) GetOne(id int) models.Note {
	db := connectToDB()
	var note models.Note
	db.Where("id = ?", id).First(&note)
	return note
}

func (SqlReposiroty) Save(note models.Note) string {
	db := connectToDB()
	db.Create(&note)
	return "All right"
}
