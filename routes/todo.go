package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	controllers "github.com/mohammedalizi/go-fiber-todo/controllers" // replace
)

func TodoRoute(route fiber.Router) {
    route.Get("", controllers.GetTodos)
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}