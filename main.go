package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"go-fiber-todos/Todo"
)

func main() {
	app := fiber.New()
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello world")
	})

	Todo.SetupTodoAPI(app)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}
