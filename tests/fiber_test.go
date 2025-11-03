package tests

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/pprof"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/stretchr/testify/require"
)

func init() {
}

func TestFiber_1(t *testing.T) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := fiber.New(fiber.Config{
		AppName: "Test Fiber Server",
	})
	//app.Use(slogfiber.New(log))
	app.Use(logger.New())
	app.Use(pprof.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Hooks().OnRoute(func(r fiber.Route) error {
		log.Info("Route info",
			slog.String("name", r.Name),
			slog.String("method", r.Method),
			slog.String("path", r.Path),
		)
		return nil
	})

	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())

	app.Get("/", func(c fiber.Ctx) error {
		req := c.Request()
		log.Info("req: " + req.String())

		c.WriteString("Hello World")
		return nil
	})

	app.Get("/json", func(c fiber.Ctx) error {
		action := c.Params("action")
		req := c.Request()
		log.Info("req: " + req.String())
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
