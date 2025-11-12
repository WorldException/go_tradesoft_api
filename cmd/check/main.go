package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/WorldException/go_tradesoft_api"
)

func main() {
	// Создание флагов
	var username = flag.String("u", "", "Имя пользователя tradesoft")
	var password = flag.String("p", "", "Пароль tradesoft")
	var article = flag.String("a", "", "Артикул")
	var brand = flag.String("b", "", "Бренд")
	var serviceName = flag.String("sname", "", "имя сервиса")
	var serviceUser = flag.String("suser", "", "Пользователь сервиса")
	var servicePassword = flag.String("spass", "", "Пароль сервиса")

	// Парсинг флагов
	flag.Parse()

	// Проверка обязательных параметров
	if *username == "" || *password == "" {
		log.Fatal("Необходимо указать имя пользователя и пароль с помощью флагов -u и -p")
	}

	if *serviceUser == "" || *servicePassword == "" || *serviceName == "" {
		log.Fatal("Необходимо указать параметры учетной записи сервиса")
	}

	// Создание клиента
	client := go_tradesoft_api.NewClientDefault()

	// Установка аутентификации
	client.SetAuth(*username, *password)

	// Пример вызова метода получения списка поставщиков
	providers, err := client.Provider().GetProviderList(*username, *password)
	if err != nil {
		log.Fatal(err)
	}
	for i, item := range providers.Data {
		fmt.Printf("Поставщик %d, Name: %s, Active: %v\n", i, item.Name, item.Active)
	}
	if len(providers.Data) == 0 {
		fmt.Printf("Поставщики: %#v\n", providers)
	}

	// Пример вызова метода получения информации о детали
	partInfo, err := client.Info().GetPartInfo(*username, *password, *article, *brand, "ru")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Информация о детали: %#v\n", partInfo)
	partJson, err := json.MarshalIndent(partInfo, "", "  ")
	println(string(partJson))

	// Пример вызова метода получения аналогов
	analogInfo, err := client.Analog().GetAnalogs(*username, *password, *article, *brand, "ru", 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Аналоги: %+v\n", analogInfo)
	//for i, item := range analogInfo.Data {
	//}
	ai, err := json.MarshalIndent(analogInfo, "", "  ")
	println(string(ai))

	// Пример вызова метода отправки SMS
	// smsResult, err := client.Messenger().SendSMS(*username, *password, "+79991114444", "Тестовое сообщение")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Отправка SMS: %+v\n", smsResult)
}
