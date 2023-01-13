package routes

import (
	"exams-api/app"
	"exams-api/services"
)

func logOutRoute() {
	app.AddPost("/log-out", services.LogOutService)
}
