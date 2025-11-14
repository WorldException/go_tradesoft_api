package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/WorldException/go_tradesoft_api"
	"github.com/WorldException/go_tradesoft_api/common"
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

	if serviceUser == nil || servicePassword == nil || serviceName == nil {
		log.Fatal("Необходимо указать параметры учетной записи сервиса")
	}

	// Создание клиента
	client := go_tradesoft_api.NewClientDefault(*username, *password)

	// Пример вызова метода получения списка поставщиков
	providers, err := client.Provider().GetProviderList()
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
	partInfo, err := client.Info().GetPartInfo(*article, *brand, "ru")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Информация о детали: %#v\n", partInfo)
	partJson, err := json.MarshalIndent(partInfo, "", "  ")
	println(string(partJson))

	// Пример вызова метода получения аналогов
	analogInfo, err := client.Analog().GetAnalogs(*article, *brand, "ru", 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Аналоги: %+v\n", analogInfo)
	//for i, item := range analogInfo.Data {
	//}
	ai, err := json.MarshalIndent(analogInfo, "", "  ")
	println(string(ai))

	// find items
	if article != nil && brand != nil {
		opt := make(common.Options)
		opt["analogs"] = "Y"
		prices, err := client.Provider().GetPriceList(*article, *brand, common.ProviderID(*serviceName), *serviceUser, *servicePassword, opt)
		if err != nil {
			log.Fatal(err)
		}
		cnt := 0
		for _, price := range prices.Container {
			fmt.Printf("Price: %s\n", price.Provider)
			for _, item := range price.Data {
				fmt.Printf("item %s: %d, %f\n", item.Code, item.Amount, item.Price)
				cnt = cnt + 1
			}
		}
		fmt.Printf("Find items prices:%d, items:%d", len(prices.Container), cnt)
	}

	// Пример вызова метода отправки SMS
	// smsResult, err := client.Messenger().SendSMS(*username, *password, "+79991114444", "Тестовое сообщение")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Отправка SMS: %+v\n", smsResult)
}
