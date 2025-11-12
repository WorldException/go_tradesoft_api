package analog

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Service represents the Analog service
type Service struct {
	client     *resty.Client
	baseURL    string
	TsUser     string
	TsPassword string
}

// NewService creates a new Analog service client
func NewService(client *resty.Client, baseURL string, user, password string) *Service {
	return &Service{
		client:     client,
		baseURL:    baseURL,
		TsUser:     user,
		TsPassword: password,
	}
}

// GetAnalogs gets list of interchanges by part number and brand
// Параметры:
// - code: Номер детали
// - brand: Производитель
// - rating: Минимальный рейтинг для отображения замены (от 1 до 5)
// - lang: Язык вывода (по умолчанию ru)
func (s *Service) GetAnalogs(code, brand, lang string, rating int) (*AnalogsResponse, error) {
	request := AnalogsRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "analog",
		Action:   "getAnalogs",
		Rating:   rating,
		Lang:     lang,
		Data: [][]string{
			{code, brand},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/analog/get-analogs")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result AnalogsResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetAnalogsAdv gets list of interchanges by part number and brand (advanced version)
// Параметры:
// - data: Список строк в формате "код:бренд"
// - rating: Минимальный рейтинг для отображения замены (от 1 до 5)
// - lang: Язык вывода (по умолчанию ru)
func (s *Service) GetAnalogsAdv(data []string, lang string, rating int) (*AnalogsAdvResponse, error) {
	request := AnalogsAdvRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "analog",
		Action:   "getAnalogsAdv",
		Rating:   rating,
		Lang:     lang,
		Data:     data,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/analog/get-analogs-adv")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result AnalogsAdvResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetProducersAdv gets list of available manufacturers by part number (advanced version)
// Параметры:
// - data: Список номеров деталей
// - rating: Минимальный рейтинг для отображения замены (от 1 до 5)
func (s *Service) GetProducersAdv(data []string, rating int) (*ProducersAdvResponse, error) {
	request := ProducersAdvRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "analog",
		Action:   "getProducersAdv",
		Rating:   rating,
		Data:     data,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/analog/get-producers-adv")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result ProducersAdvResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}
