package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kyzykyky/softwarearch/svcreg/stocksvc/service"
	"go.uber.org/zap/zapcore"
)

func (s Server) SetRoutes() {
	group := s.app.Group("/api/stock")
	group.Get("/", s.GetStock)
}

func (s Server) GetStock(c *fiber.Ctx) error {
	id := c.Query("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		s.log.Error("invalid id for stock",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	stock, ok := service.Stock[intId]
	if !ok {
		s.log.Error("stock not found",
			zapcore.Field{Key: "id", Type: zapcore.Int64Type, Integer: int64(intId)})
		return c.Status(404).JSON(fiber.Map{"error": "stock not found"})
	}
	return c.Status(200).JSON(stock)
}
