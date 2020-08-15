package Todo

import "github.com/gofiber/fiber"

func SetupTodoAPI(app *fiber.App) {
	todoRouter := app.Group("/api/v1")
	setupTodosRoutes(todoRouter)
}

func setupTodosRoutes(grp fiber.Router) {
	todosRoutes := grp.Group("/todo")
	todosRoutes.Get("/", getTodos)
	todosRoutes.Post("/", createTodo)
	todosRoutes.Get("/:id", getTodo)
	todosRoutes.Delete("/:id", deleteTodo)
	todosRoutes.Patch("/:id", updateTodo)
}
