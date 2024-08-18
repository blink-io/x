package tests

import (
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestFiber_1(t *testing.T) {
	app := fiber.New(fiber.Config{
		AppName: "Test Fiber Server",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.WriteString("Hello World")
		return nil
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		action := c.Params("action")
		fmt.Println(action)
		c.JSON(fiber.Map{
			"hello": "world",
			"type":  "info",
		})
		return nil
	})

	err := app.Listen(":4004")
	require.NoError(t, err)
}
