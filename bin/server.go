package main

import (
	"fmt"
	timeblocks "time-block-tracker/lib"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var app *fiber.App=fiber.New(fiber.Config {
		CaseSensitive:true,
		EnablePrintRoutes:false,
	})

	var blocks timeblocks.TimeBlocks=make(timeblocks.TimeBlocks)
	blocks["asdasd"]=&timeblocks.TimeBlock{
		Title:"hello",
	}

	app.Get("/get-timeblocks",func (c *fiber.Ctx) error {
		fmt.Println("huh")
		return c.JSON(timeblocks.TimeBlock {
			Title:"hello",
		})
	})

	app.Use(cors.New(cors.Config {
		AllowOrigins:"http://localhost:4200",
		AllowHeaders:"Origin, Content-Type, Accept",
	}))

	app.Listen(":4201")
}