package fiber

import (
	"github.com/gofiber/fiber/v2"
	dataErrors "github.com/kyzykyky/softwarearch/bookservice/internal/data/errors"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"go.uber.org/zap/zapcore"
)

type errorMessage struct {
	Error string `json:"error"`
}

// Translates Service errors to HTTP errors
func ErrorHandler(c *fiber.Ctx, err error) error {
	switch err {
	case nil:
		return nil

	// Switching on the service error types
	case dataErrors.ErrConflict:
		return c.Status(fiber.StatusConflict).JSON(errorMessage{err.Error()})
	case dataErrors.ErrNotFound:
		return c.Status(fiber.StatusNotFound).JSON(errorMessage{err.Error()})
	case dataErrors.ErrInvalid:
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorMessage{err.Error()})
	case dataErrors.ErrInternal:
		return c.Status(fiber.StatusInternalServerError).JSON(errorMessage{err.Error()})

	default:
		logger.Logger().Error("fiber: Unhandled error",
			zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return c.Status(fiber.StatusInternalServerError).JSON(errorMessage{err.Error()})
	}
}
