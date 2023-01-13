package routes

import (
	"exams-api/app"
	"exams-api/services"
)

func indexRoute() {
	app.AddGet("/", services.IndexService)
}
