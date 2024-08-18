package tests

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/gofiber/fiber/v3"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/stretchr/testify/require"
)

func init() {
}

func TestFiber_1(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := fiber.New(fiber.Config{
		AppName: "Test Fiber Server",
	})
	app.Use(slogfiber.New(logger))

	app.Hooks().OnRoute(func(r fiber.Route) error {
		logger.Info("Route info",
			slog.String("name", r.Name),
			slog.String("method", r.Method),
			slog.String("path", r.Path),
		)
		return nil
	})
	app.Get("/", func(c fiber.Ctx) error {
		c.WriteString("Hello World")
		return nil
	})

	app.Get("/json", func(c fiber.Ctx) error {
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
