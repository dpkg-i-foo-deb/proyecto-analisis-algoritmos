package routes

import (
	"exams-api/app"
	"exams-api/services"
)

func signUpRoute() {
	app.AddPost("/sign-up/:role", services.SignUpService)
}
