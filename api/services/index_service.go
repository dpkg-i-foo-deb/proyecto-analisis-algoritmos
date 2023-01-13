package services

import (
	"exams-api/models"

	"github.com/gofiber/fiber/v2"
)

func IndexService(c *fiber.Ctx) error {

	res := models.Response{
		Status:  string(models.OK),
		Message: "Welcome!",
	}

	return c.Status(200).JSON(res)
}
