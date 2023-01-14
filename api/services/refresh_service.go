package services

import (
	"exams-api/auth"
	"exams-api/models"

	"github.com/gofiber/fiber/v2"
)

func RefreshService(c *fiber.Ctx) error {
	var res models.Response
	var pair models.JWTPair
	var email string
	var role string
	var accessTk *fiber.Cookie
	var refreshTk *fiber.Cookie
	var err error

	res.Status = string(models.ERROR)
	res.Message = "Something went wrong"

	email, err = auth.EmailFromToken(c.Cookies("refresh-token"))

	if err != nil {
		return c.Status(500).JSON(res)
	}

	role, err = auth.RoleFromToken(c.Cookies("refresh-token"))

	if err != nil {
		return c.Status(500).JSON(res)
	}

	pair, err = auth.GenerateJWTPair(email, role)

	if err != nil {
		return c.Status(500).JSON(res)
	}

	accessTk = auth.GenerateAccessCookie(pair.Token)
	refreshTk = auth.GenerateRefreshCookie(pair.RefreshToken)

	c.Cookie(accessTk)
	c.Cookie(refreshTk)

	res.Status = string(models.OK)
	res.Message = "Session Refreshed"

	return c.Status(200).JSON(res)
}
