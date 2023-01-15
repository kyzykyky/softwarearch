package server

import "github.com/gofiber/fiber/v2"

func (s Server) SetRoutes() {
	group := s.app.Group("/api/product")
	group.Get("/product", s.GetProduct)
}

func (s Server) GetProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	return c.Status(200).JSON(fiber.Map{"id": id, "name": "bacon"})
}
