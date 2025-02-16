package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/joho/godotenv"
	"github.com/jojomak13/pdf-toolbox/core"
	"github.com/jojomak13/pdf-toolbox/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	core.LoadLogger()

}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: os.Getenv("APP_NAME")}))

	app.Use(healthcheck.New())

	app.Use(requestid.New(requestid.Config{
		Generator: utils.UUIDv4,
	}))

	app.Post("/merge", handlers.Merge)

	var port = "3000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
