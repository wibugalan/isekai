package controllers

import (
	"isekai/Isekai/db"
	"isekai/Isekai/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	type Log struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Log

	var user models.User

	err := c.BodyParser(&body)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 0)

	user.Password = password
	user.Name = body.Name
	user.Email = body.Email

	err = db.DB.Create(&user).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	type Log struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body Log

	err := c.BodyParser(&body)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	email := body.Email
	password := body.Password

	var user1 models.User
	db.DB.Where("email = ?", email).First(&user1)

	if email != user1.Email {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword(user1.Password, []byte(password))

	if err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	return err

}
