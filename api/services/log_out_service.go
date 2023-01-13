package services

import (
	"exams-api/auth"
	"exams-api/models"

	"github.com/gofiber/fiber/v2"
)

func LogOutService(c *fiber.Ctx) error {

	refresh := auth.GenerateFakeRefreshCookie()
	auth := auth.GenerateFakeAccessCookie()

	c.Cookie(auth)
	c.Cookie(refresh)

	res := models.Response{
		Status:  string(models.OK),
		Message: "Logged Out",
	}

	return c.Status(200).JSON(res)
}
