package controllers

import (
	"repository/pkg/models"
	"repository/pkg/repositories"
	"repository/pkg/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	service := services.NoteService{NoteRepository: repositories.SqlReposiroty{}}
	return c.JSON(service.GetNotes())
}

func GetOne(c *fiber.Ctx) error {
	service := services.NoteService{NoteRepository: repositories.SqlReposiroty{}}
	id, _ := strconv.Atoi(c.Params("id"))
	return c.JSON(service.GetNote(id))
}

func Post(c *fiber.Ctx) error {
	service := services.NoteService{NoteRepository: repositories.SqlReposiroty{}}
	var note models.Note
	c.BodyParser(&note)
	return c.JSON(service.SaveNote(note))
}
