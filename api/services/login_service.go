package services

import (
	"exams-api/auth"
	"exams-api/database"
	"exams-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(c *fiber.Ctx) error {

	var user models.User
	var pass string
	var tk models.JWTPair
	var role string

	res := models.Response{
		Status:  string(models.DENIED),
		Message: "Invalid credentials",
	}

	role = c.Params("role")

	if role != string(models.STUDENT) && role != string(models.ADMINISTRATOR) {
		res.Message = "Bad route"
		return c.Status(400).JSON(res)
	}

	err := c.BodyParser(&user)

	if err != nil {
		res.Message = "Malformed Entity"
		return c.Status(400).JSON(res)
	}

	pass = user.Password

	err = database.DB.Where("email = ?", user.Email).First(&user).Error

	if err != nil {
		return c.Status(401).JSON(res)
	}

	if user.Type != c.Params("role") {
		res.Message = "Not enough privilieges"
		return c.Status(401).JSON(res)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if err != nil {
		return c.Status(401).JSON(res)
	}

	tk, err = auth.GenerateJWTPair(user.Email, role)

	if err != nil {
		res.Status = string(models.ERROR)
		res.Message = "Something failed"
		return c.Status(500).JSON(res)
	}

	refresh := auth.GenerateRefreshCookie(tk.RefreshToken)
	auth := auth.GenerateAccessCookie(tk.Token)

	c.Cookie(refresh)
	c.Cookie(auth)

	res.Status = string(models.OK)
	res.Message = "Welcome, " + user.Email

	return c.Status(200).JSON(res)
}
