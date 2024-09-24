package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
API для учета возвратов
*/
func v2_posting_fbo_get(Client_Id, Api_Key, PostingNumber string) string {

	url := "https://api-seller.ozon.ru/v2/posting/fbo/get"
	body := []byte(`{
  "posting_number": "` + PostingNumber + `",
  "translit": true,
  "with": {
    "analytics_data": false,
    "financial_data": false
  }
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v2_posting_fbo_get: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonString, _ := io.ReadAll(resp.Body)

	// Выводим ответ
	return string(jsonString)
}
func v3_returns_company_fbo(Client_Id, Api_Key string, LastID int) string {

	url := "https://api-seller.ozon.ru/v3/returns/company/fbo"
	body := []byte(`{
  "filter": {
    "status": [
      "ReturnedToOzon"
    ]
  },
  "last_id":` + strconv.Itoa(LastID) + `,
  "limit": 1000
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v3_returns_company_fbo: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonString, _ := io.ReadAll(resp.Body)

	// Выводим ответ
	return string(jsonString)
}

func v3_returns_company_fbs(Client_Id, Api_Key string, LastID int) string {

	url := "https://api-seller.ozon.ru/v3/returns/company/fbs"
	body := []byte(`{
  "filter": {
    "status": "moving_to_resale"
  },
  "limit": 1000,
  "last_id": ` + strconv.Itoa(LastID) + `
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v3_returns_company_fbs: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonString, _ := io.ReadAll(resp.Body)

	// Выводим ответ
	return string(jsonString)
}
func v3_posting_fbs_get(Client_Id, Api_Key, PostingNumber string) string {

	url := "https://api-seller.ozon.ru/v3/posting/fbs/get"
	body := []byte(`{
  	"posting_number": "` + PostingNumber + `",
  	"with": {
    	"analytics_data": false,
    	"barcodes": false,
    	"financial_data": false,
    	"product_exemplars": false,
    	"translit": false
 	 }
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v3_returns_company_fbs: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonString, _ := io.ReadAll(resp.Body)

	// Выводим ответ
	return string(jsonString)
}

/*
API для отчетов
*/
func v1_report_postings_create(Client_Id, Api_Key string) string {

	url := "https://api-seller.ozon.ru/v1/report/postings/create"
	body := []byte(`{
  "filter": {
    "processed_at_from": "` + time.Now().AddDate(0, 0, -1).Format("2006-01-02T15:04:05.000Z") + `",
    "processed_at_to": "` + time.Now().Format("2006-01-02T15:04:05.000Z") + `",
    "delivery_schema": [
      "fbo"
    ],
    "sku": [],
    "cancel_reason_id": [],
    "offer_id": "",
    "status_alias": [],
    "statuses": [],
    "title": ""
  },
  "language": "DEFAULT"
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v2_posting_fbo_get: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonString, _ := io.ReadAll(resp.Body)

	// Выводим ответ
	return string(jsonString)
}
func v1_report_info(Client_Id, Api_Key string) string {
	jsonString := v1_report_postings_create(Client_Id, Api_Key)
	var response ReportResponse

	// Парсим JSON-ответ
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	// Получаем значение поля code
	code := response.Result.Code

	url := "https://api-seller.ozon.ru/v1/report/info"
	body := []byte(`{
  	"code": "` + code + `"
	}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v2_posting_fbo_get: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonBytes, _ := io.ReadAll(resp.Body)
	jsonString = string(jsonBytes)
	// Выводим ответ
	return string(jsonString)
}

/*
API для учета заказов
*/
func v3_posting_fbs_list(Client_Id, Api_Key string) string {

	url := "https://api-seller.ozon.ru/v3/posting/fbs/list"
	daysAgo := 1
	body := []byte(`{
  "dir": "ASC",
  "filter": {
    "since": "` + time.Now().AddDate(0, 0, daysAgo*(-1)-1).Format("2006-01-02") + `T21:00:00.000Z",
    "to": "` + time.Now().AddDate(0, 0, daysAgo*(-1)).Format("2006-01-02") + `T21:00:00.000Z"
  },
  "limit": 1000,
  "offset": 0,
  "with": {
    "analytics_data": false,
    "barcodes": false,
    "financial_data": false,
    "translit": false
  }
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v2_posting_fbo_get: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonString, _ := io.ReadAll(resp.Body)

	// Выводим ответ
	return string(jsonString)
}
func v2_posting_fbo_list(Client_Id, Api_Key string) string {

	url := "https://api-seller.ozon.ru/v2/posting/fbo/list"
	daysAgo := 1
	body := []byte(`{
  "dir": "ASC",
  "filter": {
    "since": "` + time.Now().AddDate(0, 0, daysAgo*(-1)-1).Format("2006-01-02") + `T21:00:00.000Z",
    "to": "` + time.Now().AddDate(0, 0, daysAgo*(-1)).Format("2006-01-02") + `T21:00:00.000Z"
  },
  "limit": 1000,
  "offset": 0,
  "translit": false,
  "with": {
    "analytics_data": false,
    "financial_data": false
  }
}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Устанавливаем необходимые заголовки (если нужны)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", Client_Id)
	req.Header.Set("Api-Key", Api_Key)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка v2_posting_fbo_get: получен статус %s", resp.Status)
	}

	// Читаем тело ответа
	jsonString, _ := io.ReadAll(resp.Body)

	// Выводим ответ
	return string(jsonString)
}
