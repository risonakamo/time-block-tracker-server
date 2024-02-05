package main

import "github.com/gofiber/fiber/v2"

func main() {
	var app *fiber.App = fiber.New(fiber.Config{
		CaseSensitive:     true,
		EnablePrintRoutes: false,
	})

	app.Pos
}