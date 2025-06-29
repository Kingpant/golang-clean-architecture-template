package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Option func(*fiberRouter)

type HealthCheckFunc func(c *fiber.Ctx) error

type Pinger interface {
	Ping() error
}

type PingerWithContext interface {
	Ping(ctx *fiber.Ctx) error
}

type PingerWithContextMethod interface {
	PingWithContext(ctx *fiber.Ctx) error
}

func WithPinger(p Pinger) Option {
	return func(fr *fiberRouter) {
		if p != nil {
			fr.healthChecks = append(fr.healthChecks, func(c *fiber.Ctx) error {
				return p.Ping()
			})
		}
	}
}

func WithPingerWithContext(p PingerWithContext) Option {
	return func(fr *fiberRouter) {
		if p != nil {
			fr.healthChecks = append(fr.healthChecks, func(c *fiber.Ctx) error {
				return p.Ping(c)
			})
		}
	}
}

func WithPingerWithContextMethod(p PingerWithContextMethod) Option {
	return func(fr *fiberRouter) {
		if p != nil {
			fr.healthChecks = append(fr.healthChecks, func(c *fiber.Ctx) error {
				return p.PingWithContext(c)
			})
		}
	}
}

type fiberRouter struct {
	app *fiber.App

	healthChecks []HealthCheckFunc
}

func NewFiberRouter(opt ...Option) *fiberRouter {
	app := fiber.New()

	fr := &fiberRouter{
		app: app,
	}

	for _, o := range opt {
		o(fr)
	}

	fr.setupHealthCheck()

	return fr
}

func (fr *fiberRouter) Shutdown() error {
	return fr.app.Shutdown()
}

func (fr *fiberRouter) Listen(port string) error {
	return fr.app.Listen(fmt.Sprintf(":%s", port))
}

func (fr *fiberRouter) App() *fiber.App {
	return fr.app
}

func (fr *fiberRouter) setupHealthCheck() {
	fr.app.Get("/healthz", func(c *fiber.Ctx) error {
		for _, check := range fr.healthChecks {
			if err := check(c); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Health check failed: %v", err))
			}
		}

		return c.SendString("OK")
	})
}
