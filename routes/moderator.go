package routes

import (
	"errors"
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

type Moderator struct {
	//this is not the model Moderator, see this as the serializer
	Id       uint   `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func CreateResponseModerator(moderatorModel models.Moderator) Moderator {
	return Moderator{Id: moderatorModel.Id, Uuid: moderatorModel.Uuid, Username: moderatorModel.Username, Password: moderatorModel.Password, Email: moderatorModel.Email}
}

func CreateModerator(c *fiber.Ctx) error {
	var moderator models.Moderator

	if err := c.BodyParser(&moderator); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&moderator)
	responseModerator := CreateResponseModerator(moderator)

	return c.Status(200).JSON(responseModerator)
}

func GetModerators(c *fiber.Ctx) error {
	moderators := []models.User{}

	database.Database.Db.Find(&moderators)
	responseModerators := []User{}
	for _, moderator := range moderators {
		responseModerator := CreateResponseUser(moderator)
		responseModerators = append(responseModerators, responseModerator)
	}

	return c.Status(200).JSON(responseModerators)
}

func FindModerator(id int, moderator *models.Moderator) error {
	database.Database.Db.Find(&moderator, "id = ?", id)
	if moderator.Id == 0 {
		return errors.New("moderator does not exist")
	}
	return nil
}

func GetModerator(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var moderator models.Moderator

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindModerator(id, &moderator); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseModerator := CreateResponseModerator(moderator)

	return c.Status(200).JSON(responseModerator)
}

func UpdateModerator(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var moderator models.Moderator

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindModerator(id, &moderator); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateModerator struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var updateData UpdateModerator

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	moderator.Username = updateData.Username
	moderator.Password = updateData.Password

	database.Database.Db.Save(&moderator)

	responseModerator := CreateResponseModerator(moderator)

	return c.Status(200).JSON(responseModerator)

}

func DeleteModerator(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var moderator models.Moderator

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindModerator(id, &moderator); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&moderator).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfuly Deleted User")
}
