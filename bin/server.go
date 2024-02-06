package main

import (
	"fmt"
	"time-block-tracker/lib/timeblocks"

	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var app *fiber.App=fiber.New(fiber.Config {
		CaseSensitive:true,
		EnablePrintRoutes:false,
	})

	var blocks timeblocks.TimeBlocks=make(timeblocks.TimeBlocks)

	// add a time block
	app.Post("/new-time-block",func (c *fiber.Ctx) error {
		fmt.Println("adding time block")
		timeblocks.AddTimeBlock(blocks)
		spew.Dump(blocks)
		return nil
	})

	app.Post("/toggle-time-block/:id",func (c *fiber.Ctx) error {
		var timeblockId string=c.Params("id")

		timeblocks.ToggleTimeBlock(blocks,timeblockId)

		spew.Dump(blocks)
		return nil
	})

	app.Get("/time-blocks",func (c *fiber.Ctx) error {
		return c.JSON(blocks)
	})
}