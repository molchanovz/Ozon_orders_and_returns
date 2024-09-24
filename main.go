package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var Client_Id, Api_Key string
	Client_Id, err := initEnv("variables.env", "Client_Id")
	Api_Key, err = initEnv("variables.env", "Api_Key")
	if err != nil {
		log.Panic(err)
	}

	writeToGoogleSheets(Client_Id, Api_Key)
}

func writeToGoogleSheets(Client_Id string, Api_Key string) {
	spreadsheetId := "1OzA_3t8f-0Un0CTAKaIbF1lD74ZYTuIPpFBfs_a2Hdo"
	date := time.Now().AddDate(0, 0, -1)
	sheetsName := "Заказы OZON-" + strconv.Itoa(date.Day())

	var values [][]interface{}

	values = append(values, []interface{}{"Отчет за " + date.Format("02.01.2006")})

	writeRange := sheetsName + "!A1"
	write(spreadsheetId, writeRange, values)

	//Заказы FBS
	postingsWithCountFBS := getPostingsMapFBS(Client_Id, Api_Key)
	for article, count := range postingsWithCountFBS {
		fmt.Println(article, " ", count)
	}
	values = [][]interface{}{}
	values = append(values, []interface{}{"Заказы FBS"})
	for article, count := range postingsWithCountFBS {
		values = append(values, []interface{}{article, count})
	}
	writeRange = sheetsName + "!A2:B100"
	write(spreadsheetId, writeRange, values)

	//Заказы FBO
	postingsWithCountFBO := getPostingsMapFBO(Client_Id, Api_Key)
	values = [][]interface{}{}
	values = append(values, []interface{}{"Заказы FBO"})
	for article, count := range postingsWithCountFBO {
		values = append(values, []interface{}{article, count})
	}

	writeRange = sheetsName + "!D2:E100"
	write(spreadsheetId, writeRange, values)

	//Возвраты
	returnsWithCount := getReturnsMap(Client_Id, Api_Key)
	for article, count := range returnsWithCount {
		fmt.Println(article, " ", count)
	}
	values = [][]interface{}{}
	values = append(values, []interface{}{"Возвраты"})
	for article, count := range returnsWithCount {
		values = append(values, []interface{}{article, count})
	}
	writeRange = sheetsName + "!G2:H100"
	write(spreadsheetId, writeRange, values)
}

func getPostingsMapFBS(Client_Id string, Api_Key string) map[string]int {
	postingsWithCountFBS := make(map[string]int)
	potings_list_fbs := get_postings_list_fbs(Client_Id, Api_Key)
	for _, posting := range potings_list_fbs.Result.PostingsFBS {
		if posting.Status != "cancelled" {
			for _, product := range posting.Products {
				postingsWithCountFBS[product.OfferId] += product.Quantity
			}
		}
	}
	return postingsWithCountFBS
}
func getPostingsMapFBO(Client_Id string, Api_Key string) map[string]int {
	postingsWithCountFBO := make(map[string]int)
	postings_list_fbo := get_postings_list_fbo(Client_Id, Api_Key)
	for _, posting := range postings_list_fbo.Result {
		if posting.Status != "cancelled" {
			for _, product := range posting.Products {
				postingsWithCountFBO[product.OfferId] += product.Quantity
			}
		}
	}
	return postingsWithCountFBO
}

func getReturnsMap(Client_Id string, Api_Key string) map[string]int {
	LastID := 0
	returnsWithCount := make(map[string]int)
	/*
		Лимит у запроса 1000, но нам нужны все возвраты,
		поэтому делаем цикл с LastID и добавляем в срез returnsFBO
	*/
	returns_fbo, LastID := get_returns_fbo(Client_Id, Api_Key, LastID)
	returnsFBO := make([]ReturnFBO, 0, len(returns_fbo))
	returnsFBO = append(returnsFBO, returns_fbo...)
	for LastID != 0 {
		returns_fbo, LastID = get_returns_fbo(Client_Id, Api_Key, LastID)
		returnsFBO = append(returnsFBO, returns_fbo...)
	}
	for i := range returnsFBO {
		parsedTime := date_parser(returnsFBO[i].ReturnedToOzonMoment)
		// Получаем год, месяц и день
		year := parsedTime.Year()
		month := parsedTime.Month()
		day := parsedTime.Day()

		if year == time.Now().Year() && month == time.Now().Month() && day == time.Now().Day()-1 {
			posting := get_posting_fbo(Client_Id, Api_Key, returnsFBO[i].PostingNumber)
			for _, product := range posting.Result.Products {
				returnsWithCount[product.OfferId] += product.Quantity
			}
		}
	}

	LastID = 0

	returns_fbs, LastID := get_returns_fbs(Client_Id, Api_Key, LastID)
	returnsFBS := make([]ReturnFBS, 0, len(returns_fbs))
	returnsFBS = append(returnsFBS, returns_fbs...)

	for LastID != 0 {
		returns_fbs, LastID = get_returns_fbs(Client_Id, Api_Key, LastID)
		returnsFBS = append(returnsFBS, returns_fbs...)
	}

	for i := range returnsFBS {
		parsedTime := date_parser(returnsFBS[i].ReturnDate)
		year := parsedTime.Year()
		month := parsedTime.Month()
		day := parsedTime.Day()
		if year == time.Now().Year() && month == time.Now().Month() && day == time.Now().Day()-1 {
			posting := get_posting_fbs(Client_Id, Api_Key, returnsFBS[i].PostingNumber)
			for _, product := range posting.Result.Products {
				returnsWithCount[product.OfferId] += product.Quantity
			}

		}
	}
	return returnsWithCount
}

func date_parser(date string) time.Time {

	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("Ошибка разбора даты:", err)
		return time.Time{}
	}
	return parsedTime
}
func initEnv(path, name string) (string, error) {
	err := godotenv.Load(path)
	if err != nil {
		log.Printf("Ошибка загрузки файла %s: %v\n", path, err)
		return "", fmt.Errorf("ошибка загрузки файла " + path)
	}
	// Получаем значения переменных среды
	env := os.Getenv(name)

	if env == "" {
		return "", fmt.Errorf("переменная среды " + name + " не установлена")
	}
	return env, err
}
