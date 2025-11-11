package main

import (
	"flag"
	"fmt"
	"log"

	"tradesoft"
)

func main() {
	// Создание флагов
	var username = flag.String("u", "", "Имя пользователя")
	var password = flag.String("p", "", "Пароль")
	var article = flag.String("a", "", "Артикул")
	var brand = flag.String("b", "", "Бренд")

	// Парсинг флагов
	flag.Parse()

	// Проверка обязательных параметров
	if *username == "" || *password == "" {
		log.Fatal("Необходимо указать имя пользователя и пароль с помощью флагов -u и -p")
	}

	// Создание клиента
	client := tradesoft.NewClientDefault()

	// Установка аутентификации
	client.SetAuth(*username, *password)

	// Пример вызова метода получения списка поставщиков
	providers, err := client.Provider().GetProviderList(*username, *password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Поставщики: %+v\n", providers)

	// Пример вызова метода получения информации о детали
	partInfo, err := client.Info().GetPartInfo(*username, *password, *article, *brand, "ru")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Информация о детали: %+v\n", partInfo)

	// Пример вызова метода получения аналогов
	analogInfo, err := client.Analog().GetAnalogs(*username, *password, *article, *brand, "ru", 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Аналоги: %+v\n", analogInfo)

	// Пример вызова метода отправки SMS
	// smsResult, err := client.Messenger().SendSMS(*username, *password, "+79991114444", "Тестовое сообщение")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Отправка SMS: %+v\n", smsResult)
}
