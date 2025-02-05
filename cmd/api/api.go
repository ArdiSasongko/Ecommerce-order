package api

import (
	"github.com/ArdiSasongko/Ecommerce-order/internal/handler"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Application struct {
	config  Config
	handler handler.Handler
}

type Config struct {
	addrHTTP string
	log      *logrus.Logger
	db       DBConfig
	auth     AuthConfig
}

type DBConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type AuthConfig struct {
	secret string
	iss    string
	aud    string
}

func (a *Application) Mount() *fiber.App {
	r := fiber.New()
	r.Use(recover.New())

	r.Get("/health", a.handler.Health.Check)

	v1 := r.Group("/v1")
	order := v1.Group("/order")
	order.Patch("/:orderID", a.handler.Order.UpdateStatus)

	order.Use(a.handler.Middleware.AuthMiddleware())
	order.Post("/", a.handler.Order.CreateOrder)
	order.Get("/", a.handler.Order.GetOrders)
	order.Get("/:orderID", a.handler.Order.GetOrder)

	return r
}

func (a *Application) Run(r *fiber.App) error {
	a.config.log.Printf("http server has run, port%v", a.config.addrHTTP)
	return r.Listen(a.config.addrHTTP)
}
