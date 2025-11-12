package go_tradesoft_api

import (
	"os"
	"testing"
)

var (
	user     = os.Getenv("TS_USER")
	password = os.Getenv("TS_PASSWORD")
)

func TestNewClient(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3", user, password)
	if client == nil {
		t.Error("Клиент не должен быть nil")
	}
}

func TestSetAuth(t *testing.T) {
	NewClient("https://test.tradesoft.ru/3", user, password)
}

func TestProviderService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3", user, password)
	providerService := client.Provider()
	if providerService == nil {
		t.Error("Provider service не должен быть nil")
	}
}

func TestInfoService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3", user, password)
	infoService := client.Info()
	if infoService == nil {
		t.Error("Info service не должен быть nil")
	}
}

func TestAnalogService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3", user, password)
	analogService := client.Analog()
	if analogService == nil {
		t.Error("Analog service не должен быть nil")
	}
}

func TestMessengerService(t *testing.T) {
	client := NewClient("https://test.tradesoft.ru/3", user, password)
	messengerService := client.Messenger()
	if messengerService == nil {
		t.Error("Messenger service не должен быть nil")
	}
}
