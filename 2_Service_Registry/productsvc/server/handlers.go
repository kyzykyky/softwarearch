package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyzykyky/softwarearch/svcreg/productsvc/service"
	"go.uber.org/zap/zapcore"
)

func (s Server) SetRoutes() {
	group := s.app.Group("/api/product")
	group.Get("/", s.GetProduct)
}

func (s Server) GetProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		s.log.Error("invalid id for product",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	product, err := s.Service.GetProduct(intId)
	if err != nil {
		switch err {
		case service.ErrProductNotFound:
			return c.Status(404).JSON(fiber.Map{"error": service.ErrProductNotFound.Error()})
		default:
			s.log.Error("error getting product",
				zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
			return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
		}
	}
	return c.Status(200).JSON(product)
}
