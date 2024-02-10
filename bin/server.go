package main

import (
	"fmt"
	"time-block-tracker/lib/timeblocks"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// --- states ---
	var blocks timeblocks.TimeBlocks=make(timeblocks.TimeBlocks)



	// --- app setup ---
	var app *fiber.App=fiber.New(fiber.Config {
		CaseSensitive:true,
		EnablePrintRoutes:false,
	})

	app.Use(cors.New())



	// --- routes ---
	// add a time block. returns the new timeblocks
	app.Post("/new-time-block",func (c *fiber.Ctx) error {
		fmt.Println("adding time block")
		timeblocks.AddTimeBlock(blocks)

		return nil
	})

	// toggle a time block given id of a block
	app.Post("/toggle-time-block/:id",func (c *fiber.Ctx) error {
		var timeblockId string=c.Params("id")

		timeblocks.ToggleTimeBlock(blocks,timeblockId)

		return nil
	})

	// get all current time blocks
	app.Get("/time-blocks",func (c *fiber.Ctx) error {
		return c.JSON(blocks)
	})

	app.Listen(":4201")
}