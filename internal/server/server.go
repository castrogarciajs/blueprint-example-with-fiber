package server

import (
	"github.com/gofiber/fiber/v2"

	"blueprint-example-with-fiber/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "blueprint-example-with-fiber",
			AppName:      "blueprint-example-with-fiber",
		}),

		db: database.New(),
	}

	return server
}
