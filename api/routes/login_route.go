package routes

import (
	"exams-api/app"
	"exams-api/services"
)

func loginRoute() {
	app.AddPost("/login/:role", services.LoginService)
}
