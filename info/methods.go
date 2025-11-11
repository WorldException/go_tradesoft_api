package info

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Service represents the Info service
type Service struct {
	client  *resty.Client
	baseURL string
}

// NewService creates a new Info service client
func NewService(client *resty.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}

// GetPartInfo gets part information
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
// - code: Номер детали
// - brand: Производитель
// - lang: Язык вывода (по умолчанию ru)
func (s *Service) GetPartInfo(user, password string, code, brand, lang string) (*PartInfoResponse, error) {
	request := PartInfoRequest{
		User:     user,
		Password: password,
		Service:  "info",
		Action:   "getPartInfo",
		Param: PartInfoParam{
			Code:  code,
			Brand: brand,
			Lang:  lang,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/info/get-part-info")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result PartInfoResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetBrandsByArticle gets list of brands by part number
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
// - code: Номер детали
// - lang: Язык вывода (по умолчанию ru)
func (s *Service) GetBrandsByArticle(user, password string, code, lang string) (*BrandsByArticleResponse, error) {
	request := BrandsByArticleRequest{
		User:     user,
		Password: password,
		Service:  "info",
		Action:   "getBrandsByArticle",
		Param: BrandsByArticleParam{
			Code: code,
			Lang: lang,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/info/get-brands-by-article")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result BrandsByArticleResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetBrandsByBarcode gets list of brands by barcode
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
// - barcode: Штрих-код детали
func (s *Service) GetBrandsByBarcode(user, password string, barcode int64) (*BrandsByBarcodeResponse, error) {
	request := BrandsByBarcodeRequest{
		User:     user,
		Password: password,
		Service:  "info",
		Action:   "getBrandsByBarcode",
		Params: BrandsByBarcodeParams{
			Barcode: barcode,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/info/get-brands-by-barcode")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result BrandsByBarcodeResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}
