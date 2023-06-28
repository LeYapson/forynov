package routes

import (
	"errors"
	"main/database"
	"time"

	"main/models"

	"github.com/gofiber/fiber/v2"
)

type Subject struct {
	Id              uint `json:"id"`
	CreatedAt       time.Time
	Subjectname     string `json:"subject_name"`
	Description     string `json:"description"`
	Messagequantity int    `json:"quantity_of_messages"`
}

func CreateResponseSubject(routeModel models.Subject) Subject {
	return Subject{Id: routeModel.Id, Subjectname: routeModel.Subjectname, Description: routeModel.Description}
}

func CreateSubject(c *fiber.Ctx) error {
	var subject models.Subject

	if err := c.BodyParser(&subject); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&subject)
	responseRoute := CreateResponseSubject(subject)

	return c.Status(200).JSON(responseRoute)
}

func GetSubjects(c *fiber.Ctx) error {
	subjects := []models.Subject{}

	database.Database.Db.Find(&subjects)
	responseSubjects := []Subject{}
	for _, subject := range subjects {
		responseSubject := CreateResponseSubject(subject)
		responseSubjects = append(responseSubjects, responseSubject)
	}

	return c.Status(200).JSON(responseSubjects)
}

func FindSubject(id int, subject *models.Subject) error {
	database.Database.Db.Find(&subject, "id = ?", id)
	if subject.Id == 0 {
		return errors.New("subject does not exist")
	}
	return nil
}

func GetSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var subject models.Subject

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindSubject(id, &subject); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseSubject := CreateResponseSubject(subject)

	return c.Status(200).JSON(responseSubject)
}

func DeleteSubject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var subject models.Subject

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindSubject(id, &subject); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&subject).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfuly Deleted User")
}
