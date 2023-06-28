package routes

import (
	"errors"
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

type UserProfile struct {
	//this is not the model UserProfile, see this as the serializer
	Id          uint   `json:"id"`
	UserProfile string `json:"user_profile"`
	Biography   string `json:"biography"`
	ProfilePic  string `json:"profile_pic"`
}

func CreateResponseUserProfile(userProfileModel models.UserProfile) UserProfile {
	return UserProfile{Id: userProfileModel.Id, UserProfile: userProfileModel.UserProfile, Biography: userProfileModel.Biography, ProfilePic: userProfileModel.UserProfile}
}

func CreateUserProfile(c *fiber.Ctx) error {
	var userProfile models.UserProfile

	if err := c.BodyParser(&userProfile); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&userProfile)
	responseUserProfile := CreateResponseUserProfile(userProfile)

	return c.Status(200).JSON(responseUserProfile)
}

func GetUserProfiles(c *fiber.Ctx) error {
	userProfiles := []models.UserProfile{}

	database.Database.Db.Find(&userProfiles)
	responseUserProfiles := []UserProfile{}
	for _, userProfile := range userProfiles {
		responseUserProfile := CreateResponseUserProfile(userProfile)
		responseUserProfiles = append(responseUserProfiles, responseUserProfile)
	}

	return c.Status(200).JSON(responseUserProfiles)
}

func FindUserProfile(id int, userProfile *models.UserProfile) error {
	database.Database.Db.Find(&userProfile, "id = ?", id)
	if userProfile.Id == 0 {
		return errors.New("userProfile does not exist")
	}
	return nil
}

func GetUserProfile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userProfile models.UserProfile

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindUserProfile(id, &userProfile); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUserProfile := CreateResponseUserProfile(userProfile)

	return c.Status(200).JSON(responseUserProfile)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userProfile models.UserProfile

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindUserProfile(id, &userProfile); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUserProfile struct {
		Biography  string `json:"biography"`
		Profilepic string `json:"profile_pic"`
	}

	var updateData UpdateUserProfile

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	userProfile.Biography = updateData.Biography
	userProfile.Profilepic = updateData.Profilepic

	database.Database.Db.Save(&userProfile)

	responseUserProfile := CreateResponseUserProfile(userProfile)

	return c.Status(200).JSON(responseUserProfile)

}

func DeleteUserProfile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var userProfile models.UserProfile

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindUserProfile(id, &userProfile); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&userProfile).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfuly Deleted User")
}
