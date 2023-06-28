package routes

import (
	"errors"
	"main/database"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	//this is not the model Message, see this as the serializer
	Id                 uint `json:"id"`
	CreatedAt          time.Time
	MessageContent     string `json:"message_content"`
	Authorofthemessage string `json:"author_of_the_message"`
	SubjectId          uint   `json:"subject_id"`
}

func CreateResponseMessage(messageModel models.Message) Message {
	return Message{Id: messageModel.Id, MessageContent: messageModel.MessageContent, Authorofthemessage: messageModel.Authorofthemessage, SubjectId: messageModel.SubjectRefer}
}

func CreateMessage(c *fiber.Ctx) error {
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&message)
	responseMessage := CreateResponseMessage(message)

	return c.Status(200).JSON(responseMessage)
}

func GetMessages(c *fiber.Ctx) error {
	messages := []models.Message{}

	database.Database.Db.Find(&messages)
	responseMessages := []Message{}
	for _, message := range messages {
		responseMessage := CreateResponseMessage(message)
		responseMessages = append(responseMessages, responseMessage)
	}

	return c.Status(200).JSON(responseMessages)
}

func FindMessage(id int, message *models.Message) error {
	database.Database.Db.Find(&message, "id = ?", id)
	if message.Id == 0 {
		return errors.New("message does not exist")
	}
	return nil
}

func GetMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var message models.Message

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindMessage(id, &message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseMessage := CreateResponseMessage(message)

	return c.Status(200).JSON(responseMessage)
}

func DeleteMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var message models.Message

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindMessage(id, &message); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Delete(&message)

	return c.SendString("Message successfully deleted")
}
