package main

import (
	"fmt"
	"log"

	"tradesoft"
)

func main() {
	// Создание клиента
	client := tradesoft.NewClient("https://service.tradesoft.ru/3")

	// Установка аутентификации
	client.SetAuth("username", "password")

	// Пример вызова метода получения списка поставщиков
	providers, err := client.Provider().GetProviderList("username", "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Поставщики: %+v\n", providers)

	// Пример вызова метода получения информации о детали
	partInfo, err := client.Info().GetPartInfo("username", "password", "0356150003", "BOSCH", "ru")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Информация о детали: %+v\n", partInfo)

	// Пример вызова метода получения аналогов
	analogInfo, err := client.Analog().GetAnalogs("username", "password", "0356150003", "BOSCH", "ru", 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Аналоги: %+v\n", analogInfo)

	// Пример вызова метода отправки SMS
	smsResult, err := client.Messenger().SendSMS("username", "password", "+79991114444", "Тестовое сообщение")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Отправка SMS: %+v\n", smsResult)
}
