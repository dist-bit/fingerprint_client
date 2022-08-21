package routes

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type router struct{}

type IRouter interface {
	InitTrainerRouter(fiber.Router)
}

var (
	m          *router
	routerOnce sync.Once
)

func Router() IRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
