package Todo

import (
	"fmt"
	"github.com/gofiber/fiber"
)

func deleteTodo(ctx *fiber.Ctx) {
	var id string
	id = ctx.Params("id")
	err := deleteTodoFromDB(id)
	if err != nil {
		ctx.Status(fiber.StatusNotFound).Send(err.Error())
		return
	} else {
		ctx.Status(fiber.StatusOK)
	}
}

func getTodo(ctx *fiber.Ctx) {
	var id string
	id = ctx.Params("id")
	todo, err := getTodoFromDB(id)
	if err != nil {
		ctx.Status(fiber.StatusNotFound).Send(err.Error())
		return
	}
	_ = ctx.Status(fiber.StatusOK).JSON(&todo)
}

func createTodo(ctx *fiber.Ctx) {
	var todo Todo
	err := ctx.BodyParser(&todo)
	if err == nil {
		insertErr := insertTodoIntoDB(todo)
		if insertErr != nil {
			ctx.Status(fiber.StatusBadRequest).Send(insertErr.Error())
		} else {
			ctx.Status(fiber.StatusCreated)
		}
	} else {
		fmt.Println(err)
		ctx.Status(fiber.StatusBadRequest).Send(err.Error())
	}
}

func getTodos(ctx *fiber.Ctx) {
	_ = ctx.Status(fiber.StatusOK).JSON(getTodosFromDB())
}

func updateTodo(ctx *fiber.Ctx) {
	var todo Todo
	id := ctx.Params("id")
	err := ctx.BodyParser(&todo)
	if err == nil {
		updateErr := updateTodoInDB(id, todo)
		if updateErr == nil {
			ctx.Status(fiber.StatusOK)
		} else {
			ctx.Status(fiber.StatusNotFound).Send(updateErr.Error())
		}
	} else {
		ctx.Status(fiber.StatusBadRequest).Send(err.Error())
	}
}
