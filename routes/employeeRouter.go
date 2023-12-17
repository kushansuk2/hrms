package routes

import (
	"hrms/handlers"
	"gofr.dev/pkg/gofr"
)

func EmployeeRoutes(app *gofr.Gofr) {
	app.GET("/employee",handlers.GetAllEmployees())
	app.POST("/employee",handlers.CreateEmployee())
	app.PUT("/employee/{id}",handlers.UpdateEmployee())
	app.DELETE("/employee/{id}",handlers.DeleteEmployee())
}