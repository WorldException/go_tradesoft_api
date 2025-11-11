package messenger

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Service represents the Messenger service
type Service struct {
	client  *resty.Client
	baseURL string
}

// NewService creates a new Messenger service client
func NewService(client *resty.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}

// SendSMS sends an SMS to a mobile number
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
// - phoneNumber: Номер телефона в международном формате
// - text: Текст SMS
func (s *Service) SendSMS(user, password string, phoneNumber, text string) (*SendSMSResponse, error) {
	request := SendSMSRequest{
		User:     user,
		Password: password,
		Service:  "messenger",
		Action:   "sendSms",
		Param: SMSParam{
			PhoneNumber: phoneNumber,
			Text:        text,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/messenger/send-sms")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result SendSMSResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetSmsStatus gets information about sent SMS
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
// - smsIDs: Список ID SMS для проверки статуса
func (s *Service) GetSmsStatus(user, password string, smsIDs []string) (*SMSStatusResponse, error) {
	request := SMSStatusRequest{
		User:     user,
		Password: password,
		Service:  "messenger",
		Action:   "getSmsStatus",
		Param:    smsIDs,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/messenger/get-sms-status")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result SMSStatusResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetSmsBalance checks SMS balance (deprecated)
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
func (s *Service) GetSmsBalance(user, password string) (*SMSBalanceResponse, error) {
	request := SMSBalanceRequest{
		User:     user,
		Password: password,
		Service:  "messenger",
		Action:   "getSmsBalance",
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/messenger/get-sms-balance")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result SMSBalanceResponse
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}

// GetSmsBalance2 checks SMS balance (new version)
// Параметры:
// - user: Логин на сайте TradeSoft
// - password: Пароль на сайте TradeSoft
func (s *Service) GetSmsBalance2(user, password string) (*SMSBalance2Response, error) {
	request := SMSBalance2Request{
		User:     user,
		Password: password,
		Service:  "messenger",
		Action:   "getSmsBalance2",
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка при сериализации запроса: %w", err)
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.baseURL + "/messenger/get-sms-balance2")

	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	var result SMSBalance2Response
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при десериализации ответа: %w", err)
	}

	return &result, nil
}
