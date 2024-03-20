package main

import (
	"github.com/gofiber/fiber/v2"
	"restApi/user"
)

func Routers(app *fiber.App) {
	app.Get("/user", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()
	var app = fiber.New()
	app.Get("/", hello)
	Routers(app)
	app.Listen(":3000")
}

func hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome")
}
