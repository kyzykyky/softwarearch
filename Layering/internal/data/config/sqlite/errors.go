package sqlite

import (
	"github.com/kyzykyky/softwarearch/bookservice/internal/data/errors"
	"github.com/kyzykyky/softwarearch/bookservice/internal/integration/logger"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

// Convert to repository common errors
func ConvertError(err error) error {
	switch err {
	case nil:
		logger.Logger().Error("convert error: nil error")
		return nil

	case gorm.ErrRecordNotFound:
		logger.Logger().Error("convert error: not found", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return errors.ErrNotFound
	case gorm.ErrInvalidData:
		logger.Logger().Error("convert error: invalid data", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return errors.ErrInvalid

	default:
		logger.Logger().Error("convert error: unknown error", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return errors.ErrUnknown
	}
}
