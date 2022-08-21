package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"hackaton/src/models"
	"hackaton/src/routes"
	"log"
	"net/http"
)

func main() {
	// create server
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "hack",
		AppName:      "hack",
		BodyLimit:    20 * 1024 * 1024, // Max 20mb body size
		ErrorHandler: func(c *fiber.Ctx, err error) error { // error response handler
			return c.Status(http.StatusOK).JSON(models.Response{
				Status:  false,
				Payload: err.Error(),
			})
		},
	})

	var ConfigDefault = cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(logger.New())

	app.Use(cors.New(ConfigDefault))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	routes.Router().InitTrainerRouter(v1)

	log.Fatal(app.Listen(":8080"))
}
