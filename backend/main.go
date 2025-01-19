package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type CartItem struct {
	ItemType    string  `json:"item_type"`
	PaymentType string  `json:"payment_type"`
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	ItemCode    string  `json:"item_code"`
	TruCode     string  `json:"tru_code"`
	Tax         string  `json:"tax"`
	Sum         float64 `json:"sum"`
	ID          int     `json:"id"`
}

type Order struct {
	CartJSON    []CartItem `json:"cart_json"`
	ClientEmail string     `json:"client_email"`
	ClientPhone string     `json:"client_phone"`
	ClientID    string     `json:"clientid"`
	Expiry      string     `json:"expiry"`
	OrderID     string     `json:"orderid"`
	PayAmount   float64    `json:"pay_amount"`
	// ServiceName string     `json:"service_name"`
	Token string `json:"token"`
}

func main() {
	// Логин и пароль от личного кабинета PayKeeper
	user := "admin"
	password := "1$Fgtkmcby2019#"

	// Basic-авторизация передаётся как base64
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	log.Println("Basic Auth:", auth)

	// Укажите адрес ВАШЕГО сервера PayKeeper
	serverPaykeeper := "https://sdmedik.server.paykeeper.ru"

	// Параметры платежа
	order := Order{
		CartJSON: []CartItem{
			{
				ItemType:    "goods",
				PaymentType: "full",
				SKU:         "07-01-03.0001",
				Name:        "Кресло коляска для инвалидов Ortonica Trend 35",
				Price:       28910.3,
				Quantity:    1,
				ItemCode:    "",
				TruCode:     "266014120.170000111",
				Tax:         "none",
				Sum:         28910.3,
				ID:          0,
			},
		},
		ClientEmail: "test@mail.ru",
		ClientPhone: "+7 888 888 88 88",
		ClientID:    "test",
		Expiry:      "2025-01-26 23:59",
		OrderID:     "123123",
		PayAmount:   28910.3,
		// ServiceName: ";PKC|[{\"item_type\":\"goods\",\"payment_type\":\"full\",\"sku\":\"07-01-03.0001\",\"name\":\"Кресло коляска для инвалидов Ortonica Trend 35\",\"price\":28910.3,\"quantity\":1,\"item_code\":\"\",\"tru_code\":\"266014120.170000111\",\"tax\":\"none\",\"sum\":28910.3,\"id\":0}]|",
		Token: "", // Токен будет добавлен позже
	}

	// Готовим первый запрос на получение токена безопасности
	tokenURI := "/info/settings/token/"

	// Создаем HTTP-запрос для получения токена
	req, err := http.NewRequest("GET", serverPaykeeper+tokenURI, nil)
	if err != nil {
		log.Fatalf("Ошибка при создании запроса для получения токена: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+auth)

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса для получения токена: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа для получения токена: %v", err)
	}

	// Парсим JSON-ответ
	var tokenResponse map[string]string
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		log.Fatalf("Ошибка при парсинге JSON для получения токена: %v", err)
	}

	// В ответе должно быть заполнено поле token, иначе - ошибка
	token, ok := tokenResponse["token"]
	if !ok {
		log.Fatalf("Поле 'token' отсутствует в ответе")
	}

	log.Println("Полученный токен:", token)

	// Готовим запрос 3.4 JSON API на получение счёта
	invoiceURI := "/change/invoice/preview/"

	// Добавляем токен в структуру заказа
	order.Token = token

	// Преобразуем структуру заказа в URL-encoded форму
	formData := url.Values{}
	formData.Set("cart_json", fmt.Sprintf("%v", order.CartJSON))
	formData.Set("client_email", order.ClientEmail)
	formData.Set("client_phone", order.ClientPhone)
	formData.Set("clientid", order.ClientID)
	formData.Set("expiry", order.Expiry)
	formData.Set("orderid", order.OrderID)
	formData.Set("pay_amount", fmt.Sprintf("%.2f", order.PayAmount))
	// formData.Set("service_name", order.ServiceName)
	formData.Set("token", order.Token)

	log.Println("Отправляемые данные:", formData.Encode())

	// Создаем HTTP-запрос для создания счёта
	req, err = http.NewRequest("POST", serverPaykeeper+invoiceURI, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatalf("Ошибка при создании запроса для создания счёта: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+auth)

	// Выполняем запрос
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса для создания счёта: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа для создания счёта: %v", err)
	}

	// Парсим JSON-ответ
	var invoiceResponse map[string]string
	if err := json.Unmarshal(body, &invoiceResponse); err != nil {
		log.Fatalf("Ошибка при парсинге JSON для создания счёта: %v", err)
	}

	log.Println("Ответ от сервера:", invoiceResponse)

	// В ответе должно быть поле invoice_id, иначе - ошибка
	invoiceID, ok := invoiceResponse["invoice_id"]
	if !ok {
		log.Fatalf("Поле 'invoice_id' отсутствует в ответе")
	}

	// В этой переменной прямая ссылка на оплату с заданными параметрами
	link := fmt.Sprintf("%s/bill/%s/", serverPaykeeper, invoiceID)

	// Теперь её можно использовать как угодно, например, выводим ссылку на оплату
	fmt.Println("Ссылка на оплату:", link)
}
