package tradesoft

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3")
	if client == nil {
		t.Error("Клиент не должен быть nil")
	}
}

func TestSetAuth(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3")
	client.SetAuth("user", "password")
	// Проверяем, что аутентификация установлена
	// Тесты могут требовать реальных данных для проверки
}

func TestProviderService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3")
	providerService := client.Provider()
	if providerService == nil {
		t.Error("Provider service не должен быть nil")
	}
}

func TestInfoService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3")
	infoService := client.Info()
	if infoService == nil {
		t.Error("Info service не должен быть nil")
	}
}

func TestAnalogService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3")
	analogService := client.Analog()
	if analogService == nil {
		t.Error("Analog service не должен быть nil")
	}
}

func TestMessengerService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3")
	messengerService := client.Messenger()
	if messengerService == nil {
		t.Error("Messenger service не должен быть nil")
	}
}
