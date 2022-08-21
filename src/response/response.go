package response

import (
	"github.com/gofiber/fiber/v2"
	"hackaton/src/models"
	"net/http"
)

type Response struct{}

func (response *Response) Success(data interface{}, c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(models.Response{
		Status:  true,
		Payload: data,
	})
}

func (response *Response) Image(data []byte, c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "image/jpeg")
	return c.Status(http.StatusOK).Send(data)
}

func (response *Response) PDF(data []byte, c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/pdf")
	return c.Status(http.StatusOK).Send(data)
}

func (response *Response) File(data []byte, c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/octet-stream")
	return c.Status(http.StatusOK).Send(data)
}

func (response *Response) Code(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(models.Response{
		Status:  true,
		Payload: "successful operation",
	})
}

func (response *Response) Error(description string, c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(models.Response{
		Status:  false,
		Payload: description,
	})
}
