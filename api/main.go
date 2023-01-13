package main

import (
	"exams-api/app"
	_ "exams-api/database"
	_ "exams-api/routes"
)

func main() {
	app.Start()
}
