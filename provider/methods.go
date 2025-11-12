package provider

import (
	"encoding/json"
	"fmt"

	"github.com/WorldException/go_tradesoft_api/common"

	"github.com/go-resty/resty/v2"
)

// Service represents the Provider service
type Service struct {
	client     *resty.Client
	baseURL    string
	TsUser     string
	TsPassword string
}

// NewService creates a new Provider service client
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
func NewService(client *resty.Client, baseURL string, user, password string) *Service {
	return &Service{
		client:     client,
		baseURL:    baseURL,
		TsUser:     user,
		TsPassword: password,
	}
}

// GetProducerList gets list of manufacturers by code
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
// - providerID: ID поставщика (может быть получен из списка доступных поставщиков - GetProviderList)
// - code: Номер детали
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
// - options: Опциональные параметры поиска (если доступны)
func (s *Service) GetProducerList(code string, providerID common.ProviderID, login, pwd string, options common.Options) (*ProducerListResponse, error) {
	request := ProducerListRequest{
		User:      s.TsUser,
		Password:  s.TsPassword,
		Service:   "provider",
		Action:    "getProducerList",
		Timelimit: 10,
		Container: []ProducerListItem{
			{
				Provider: providerID,
				Login:    login,
				Password: pwd,
				Code:     code,
				Options:  options,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/get-producer-list/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result ProducerListResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetPriceList gets list of offers by manufacturers
// Параметры:
// - providerID: ID поставщика
// - code: Номер детали
// - producer: Производитель
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
// - options: Опциональные параметры поиска
func (s *Service) GetPriceList(code, producer string, providerID common.ProviderID, login, pwd string, options common.Options) (*PriceListResponse, error) {
	request := PriceListRequest{
		User:      s.TsUser,
		Password:  s.TsPassword,
		Service:   "provider",
		Action:    "getPriceList",
		Timelimit: 10,
		Container: []PriceListItem{
			{
				Provider: providerID,
				Login:    login,
				Password: pwd,
				Code:     code,
				Producer: producer,
				Options:  options,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/get-price-list/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result PriceListResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetAdditionalPartInfo gets additional information about part
// Параметры:
// - providerID: ID поставщика
// - providerCode: Код детали на сайте поставщика (может быть получен методом getPriceList)
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
func (s *Service) GetAdditionalPartInfo(providerCode common.ProviderCode, providerID common.ProviderID, login, pwd string) (*AdditionalPartInfoResponse, error) {
	request := AdditionalPartInfoRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "provider",
		Action:   "getAdditionalPartInfo",
		Container: []AdditionalPartInfoItem{
			{
				Provider:     providerID,
				Login:        login,
				Password:     pwd,
				ProviderCode: providerCode,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/get-additional-part-info/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result AdditionalPartInfoResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetOptionsList gets list of suppliers' options
// Параметры:
// - providerID: ID поставщика
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
func (s *Service) GetOptionsList(providerID common.ProviderID, login, pwd string) (*OptionsListResponse, error) {
	request := OptionsListRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "provider",
		Action:   "getOptionsList",
		Container: []OptionsListItem{
			{
				Provider: providerID,
				Login:    login,
				Password: pwd,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/get-options-list/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result OptionsListResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetProviderList gets list of available suppliers
func (s *Service) GetProviderList() (*ProviderListResponse, error) {
	request := ProviderListRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "provider",
		Action:   "getProviderList",
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/get-provider-list/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result ProviderListResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// PreOrderSearch gets list of items for order
// Параметры:
// - providerID: ID поставщика
// - code: Номер детали
// - producer: Производитель
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
// - itemHash: Хэш позиции (может быть получен методом getPriceList)
func (s *Service) PreOrderSearch(code, producer string, providerID common.ProviderID, login, pwd string, itemHash common.ItemHash) (*PreOrderSearchResponse, error) {
	request := PreOrderSearchRequest{
		User:      s.TsUser,
		Password:  s.TsPassword,
		Service:   "provider",
		Action:    "preOrderSearch",
		Timelimit: 10,
		Container: []PreOrderSearchItem{
			{
				Provider: providerID,
				Login:    login,
				Password: pwd,
				Code:     code,
				Producer: producer,
				ItemHash: itemHash,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/pre-order-search/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result PreOrderSearchResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// MakeOrderOffline places an order to a supplier
// Параметры:
// - providerID: ID поставщика
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
// - items: Список заказываемых позиций
// - comment: Комментарий к заказу
// - clientOrderNumber: Номер заказа клиента
func (s *Service) MakeOrderOffline(providerID common.ProviderID, login, pwd string, items []MakeOrderItem, comment, clientOrderNumber string) (*MakeOrderOfflineResponse, error) {
	request := MakeOrderOfflineRequest{
		User:      s.TsUser,
		Password:  s.TsPassword,
		Service:   "provider",
		Action:    "makeOrderOffline",
		Timelimit: 10,
		Param: []MakeOrderParam{
			{
				Provider:          providerID,
				Login:             login,
				Password:          pwd,
				Comment:           comment,
				ClientOrderNumber: clientOrderNumber,
				Items:             items,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/make-order-offline/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result MakeOrderOfflineResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetStatusList gets list of available statuses
// Параметры:
// - providerID: ID поставщика
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
func (s *Service) GetStatusList(providerID common.ProviderID, login, pwd string) (*StatusListResponse, error) {
	request := StatusListRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "provider",
		Action:   "getStatusList",
		Container: []StatusListItem{
			{
				Provider: providerID,
				Login:    login,
				Password: pwd,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/get-status-list/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result StatusListResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetItemsStatus gets statuses of ordered items
// Параметры:
// - providerID: ID поставщика
// - login: Логин на сайте поставщика
// - password: Пароль на сайте поставщика
// - items: Список ID заказанных позиций
func (s *Service) GetItemsStatus(providerID common.ProviderID, login, pwd string, items []common.OrderItemID) (*ItemsStatusResponse, error) {
	request := ItemsStatusRequest{
		User:     s.TsUser,
		Password: s.TsPassword,
		Service:  "provider",
		Action:   "getItemsStatus",
		Container: []ItemsStatusItem{
			{
				Provider: providerID,
				Login:    login,
				Password: pwd,
				Items:    items,
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/provider/get-items-status/")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result ItemsStatusResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}
