package services

import (
	"exams-api/auth"
	"exams-api/database"
	"exams-api/models"
	"exams-api/util"

	"github.com/gofiber/fiber/v2"
)

func SignUpService(c *fiber.Ctx) error {

	var newUser models.User

	res := models.Response{
		Status:  string(models.DENIED),
		Message: "Not enough privileges",
	}

	role := c.Params("role")

	if role != string(models.ADMINISTRATOR) && role != string(models.STUDENT) {
		res.Message = "Bad route"
		return c.Status(400).JSON(res)
	}

	if role == string(models.ADMINISTRATOR) {
		tk := c.Cookies("access-token")
		hasSession, err := auth.ValidateToken(tk)

		if !hasSession || err != nil {
			return c.Status(401).JSON(res)
		}

		tkRole, err := auth.RoleFromToken(tk)

		if err != nil || tkRole != string(models.ADMINISTRATOR) {
			return c.Status(401).JSON(res)
		}
	}

	err := c.BodyParser(&newUser)

	if err != nil {
		res.Message = "Malformed body"
		return c.Status(400).JSON(res)
	}

	newUser.Type = role
	newUser.Password, err = util.HashPassword(newUser.Password)

	res.Message = "Something went wrong"
	res.Status = string(models.ERROR)

	if err != nil {
		return c.Status(500).JSON(res)
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		return c.Status(500).JSON(res)
	}

	res.Message = "Registered Successfully"
	res.Status = string(models.OK)

	return c.Status(201).JSON(res)
}
