package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kyzykyky/softwarearch/svcreg/pkg/domain"
	"go.uber.org/zap/zapcore"
)

// Send request to stock service
func (s *Service) GetStock(id int) (domain.Stock, error) {
	stockService, err := s.ConsulConnection.SelectServiceInstance("stock")
	if err != nil {
		s.log.Error("Error selecting stock service", zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return domain.Stock{}, err
	}
	s.log.Info("Selected stock service", zapcore.Field{Key: "stockService", Type: zapcore.StringType, String: stockService})

	// Call stock service
	url := fmt.Sprintf("http://%s/api/stock?id=%d", stockService, id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		s.log.Error("Error creating request", zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return domain.Stock{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		s.log.Error("Error sending request", zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return domain.Stock{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.log.Error("Error reading response", zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return domain.Stock{}, err
	}

	stock := domain.Stock{}
	err = json.Unmarshal(body, &stock)
	if err != nil {
		s.log.Error("Error unmarshalling response", zapcore.Field{Key: "error", Type: zapcore.ErrorType, Interface: err})
		return domain.Stock{}, err
	}

	return stock, nil
}
