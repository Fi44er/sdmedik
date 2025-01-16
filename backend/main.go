package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Логин и пароль от личного кабинета PayKeeper
	user := "admin"
	password := "1$Fgtkmcby2019#"

	// Basic-авторизация передаётся как base64
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	// Укажите адрес ВАШЕГО сервера PayKeeper, адрес demo.rsb-processing.ru - пример!
	serverPaykeeper := "https://sdmedik.server.paykeeper.ru"

	// Параметры платежа, сумма - обязательный параметр
	// Остальные параметры можно не задавать
	paymentData := map[string]string{
		"pay_amount":   "42.50",
		"clientid":     "Иванов Иван Иванович",
		"orderid":      "Заказ № 10",
		"pstype":       "cert",
		"client_email": "test@example.com",
		"service_name": "Услуга",
		"client_phone": "8 (910) 123-45-67",
	}

	// Готовим первый запрос на получение токена безопасности
	tokenURI := "/info/settings/token/"

	// Создаем HTTP-запрос для получения токена
	req, err := http.NewRequest("GET", serverPaykeeper+tokenURI, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса для получения токена:", err)
		return
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+auth)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса для получения токена:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа для получения токена:", err)
		return
	}

	// Парсим JSON-ответ
	var tokenResponse map[string]string
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		fmt.Println("Ошибка при парсинге JSON для получения токена:", err)
		return
	}

	// В ответе должно быть заполнено поле token, иначе - ошибка
	token, ok := tokenResponse["token"]
	if !ok {
		fmt.Println("Поле 'token' отсутствует в ответе")
		return
	}

	// Готовим запрос 3.4 JSON API на получение счёта
	invoiceURI := "/change/invoice/preview/"

	// Формируем список POST параметров
	paymentData["token"] = token
	formData := url.Values{}
	for key, value := range paymentData {
		formData.Set(key, value)
	}

	// Создаем HTTP-запрос для создания счёта
	req, err = http.NewRequest("POST", serverPaykeeper+invoiceURI, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Ошибка при создании запроса для создания счёта:", err)
		return
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+auth)

	// Выполняем запрос
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса для создания счёта:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа для создания счёта:", err)
		return
	}

	// Парсим JSON-ответ
	var invoiceResponse map[string]string
	if err := json.Unmarshal(body, &invoiceResponse); err != nil {
		fmt.Println("Ошибка при парсинге JSON для создания счёта:", err)
		return
	}

	// В ответе должно быть поле invoice_id, иначе - ошибка
	invoiceID, ok := invoiceResponse["invoice_id"]
	if !ok {
		fmt.Println("Поле 'invoice_id' отсутствует в ответе")
		return
	}

	// В этой переменной прямая ссылка на оплату с заданными параметрами
	link := fmt.Sprintf("%s/bill/%s/", serverPaykeeper, invoiceID)

	// Теперь её можно использовать как угодно, например, выводим ссылку на оплату
	fmt.Println("Ссылка на оплату:", link)
}

// package main
//
// import (
// 	"crypto/tls"
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"net/url"
// 	"time"
// )
//
// func main() {
// 	// Инициализация генератора случайных чисел
// 	rand.Seed(time.Now().UnixNano())
//
// 	// URL для отправки запроса
// 	apiUrl := "https://ask.oksei.ru/vote.php"
//
// 	// Создаем кастомный HTTP-клиент с отключенной проверкой SSL
// 	client := &http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Отключаем проверку SSL
// 		},
// 	}
//
// 	// Цикл для отправки 1000 запросов
// 	for i := 0; i < 1000; i++ {
// 		// Генерация случайного IP-адреса
// 		ip := generateRandomIP()
//
// 		// Подготовка данных для отправки
// 		data := url.Values{}
// 		data.Set("id", "9307")
// 		data.Set("type", "question")
// 		data.Set("operation", "like")
// 		data.Set("ip", ip)
//
// 		// Отправка POST-запроса
// 		resp, err := client.PostForm(apiUrl, data)
// 		if err != nil {
// 			fmt.Printf("Ошибка при отправке запроса: %v\n", err)
// 			continue
// 		}
//
// 		// Закрытие тела ответа
// 		resp.Body.Close()
//
// 		// Вывод статуса запроса
// 		fmt.Printf("Запрос %d: IP=%s, Статус=%s\n", i+1, ip, resp.Status)
// 	}
// }
//
// // Функция для генерации случайного IP-адреса
// func generateRandomIP() string {
// 	// Генерация четырех случайных чисел от 0 до 255
// 	part1 := rand.Intn(256)
// 	part2 := rand.Intn(256)
// 	part3 := rand.Intn(256)
// 	part4 := rand.Intn(256)
//
// 	// Форматирование IP-адреса
// 	return fmt.Sprintf("%d.%d.%d.%d", part1, part2, part3, part4)
// }
