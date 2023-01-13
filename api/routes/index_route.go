package routes

import (
	"exams-api/app"
	"exams-api/auth"
	"exams-api/services"
)

func indexRoute() {
	app.AddGet("/", auth.ValidateAccessToken, services.IndexService)
}
