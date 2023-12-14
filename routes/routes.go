package routes

import (
	"Go-ToDo/controllers"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {

	app.Get("/toDo", controllers.GetTasks)
	app.Post("/toDo", controllers.CreateTask)
	app.Get("/toDo/:taskId", controllers.GetTask)
	app.Put("/toDo/:taskId", controllers.CompleteTask)
	app.Delete("/toDo/:taskId", controllers.DeleteTask)

}
