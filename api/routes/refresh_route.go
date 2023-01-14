package routes

import (
	"exams-api/app"
	"exams-api/auth"
	"exams-api/services"
)

func refreshRoute() {
	app.AddPost("/refresh", auth.ValidateRefreshToken,
		auth.ValidateRefreshTokenDate,
		services.RefreshService)
}
