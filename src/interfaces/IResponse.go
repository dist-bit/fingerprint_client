package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type IResponse interface {
	Success(data interface{}, c *fiber.Ctx) error
	Image(data []byte, c *fiber.Ctx) error
	PDF(data []byte, c *fiber.Ctx) error
	Code(c *fiber.Ctx) error
	Error(description string, c *fiber.Ctx) error
	File(data []byte, c *fiber.Ctx) error
}
