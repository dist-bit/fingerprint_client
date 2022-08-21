package routes

import (
	"github.com/gofiber/fiber/v2"
	"hackaton/src/container"
)

func (router *router) InitTrainerRouter(r fiber.Router) {
	controller := container.ServiceContainer().InjectTrainerController()
	trainer := r.Group("/services")

	trainer.Post("/fingerprints", func(c *fiber.Ctx) error {
		return controller.GetFingerprintsController(c)
	})
	trainer.Post("/fingerprints/process", func(c *fiber.Ctx) error {
		return controller.ProcessFingerprintController(c)
	})
	trainer.Post("/fingerprints/nfiq", func(c *fiber.Ctx) error {
		return controller.GetFingerprintNfiqScoreController(c)
	})
	trainer.Post("/wsq", func(c *fiber.Ctx) error {
		return controller.GetWSQFingerprintController(c)
	})
	trainer.Get("/iso", func(c *fiber.Ctx) error {
		return controller.GetISOFIle(c)
	})
}
