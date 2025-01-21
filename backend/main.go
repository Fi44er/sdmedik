// package main
//
// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )
//
// type CartItem struct {
// 	ItemType    string  `json:"item_type"`
// 	PaymentType string  `json:"payment_type"`
// 	SKU         string  `json:"sku"`
// 	Name        string  `json:"name"`
// 	Price       float64 `json:"price"`
// 	Quantity    int     `json:"quantity"`
// 	ItemCode    string  `json:"item_code"`
// 	TruCode     string  `json:"tru_code"`
// 	Tax         string  `json:"tax"`
// 	Sum         float64 `json:"sum"`
// }
//
// type Order struct {
// 	CartJSON    []CartItem `json:"cart_json"`
// 	ClientEmail string     `json:"client_email"`
// 	ClientPhone string     `json:"client_phone"`
// 	ClientID    string     `json:"clientid"`
// 	Expiry      string     `json:"expiry"`
// 	OrderID     string     `json:"orderid"`
// 	PayAmount   float64    `json:"pay_amount"`
// 	ServiceName string     `json:"service_name"`
// 	Token       string     `json:"token"`
// }
//
// func main() {
// 	// Логин и пароль от личного кабинета PayKeeper
// 	user := "admin"
// 	password := "1$Fgtkmcby2019#"
//
// 	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))
//
// 	serverPaykeeper := "https://sdmedik.server.paykeeper.ru"
//
// 	cart := []CartItem{
// 		{
// 			ItemType:    "goods",
// 			PaymentType: "full",
// 			SKU:         "",
// 			Name:        "Кресло коляска для инвалидов Ortonica Trend 35",
// 			Price:       28910.3,
// 			Quantity:    1,
// 			ItemCode:    "",
// 			TruCode:     "266014120.170000111",
// 			Tax:         "none",
// 			Sum:         28910.3,
// 		},
// 	}
//
// 	jsonData, err := json.Marshal(cart)
// 	if err != nil {
// 		fmt.Println("Ошибка при сериализации в JSON:", err)
// 		return
// 	}
//
// 	// Формируем строку serviceName
// 	serviceName := fmt.Sprintf(";PKC|%s|", jsonData)
// 	// Параметры платежа
// 	order := Order{
// 		CartJSON:    cart,
// 		ClientEmail: "q@mail.ru",
// 		ClientPhone: "+7",
// 		ClientID:    "q",
// 		Expiry:      "2025-02-28 03:00",
// 		OrderID:     "",
// 		PayAmount:   28910.3,
// 		ServiceName: serviceName,
// 		Token:       "",
// 	}
//
// 	log.Println(serviceName == order.ServiceName)
// 	log.Println(serviceName)
//
// 	// Готовим первый запрос на получение токена безопасности
// 	tokenURI := "/info/settings/token/"
//
// 	// Создаем HTTP-запрос для получения токена
// 	req, err := http.NewRequest("GET", serverPaykeeper+tokenURI, nil)
// 	if err != nil {
// 		log.Fatalf("Ошибка при создании запроса для получения токена: %v", err)
// 	}
//
// 	// Устанавливаем заголовки
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Authorization", "Basic "+auth)
//
// 	// Выполняем запрос
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Ошибка при выполнении запроса для получения токена: %v", err)
// 	}
// 	defer resp.Body.Close()
//
// 	// Читаем ответ
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Ошибка при чтении ответа для получения токена: %v", err)
// 	}
//
// 	// Парсим JSON-ответ
// 	var tokenResponse map[string]string
// 	if err := json.Unmarshal(body, &tokenResponse); err != nil {
// 		log.Fatalf("Ошибка при парсинге JSON для получения токена: %v", err)
// 	}
//
// 	// В ответе должно быть заполнено поле token, иначе - ошибка
// 	token, ok := tokenResponse["token"]
// 	if !ok {
// 		log.Fatalf("Поле 'token' отсутствует в ответе")
// 	}
//
// 	// Готовим запрос 3.4 JSON API на получение счёта
// 	invoiceURI := "/change/invoice/preview/"
//
// 	// Добавляем токен в структуру заказа
// 	order.Token = token
//
// 	// Преобразуем структуру заказа в URL-encoded форму
// 	formData := url.Values{}
// 	formData.Set("cart_json", fmt.Sprintf("%v", order.CartJSON))
// 	formData.Set("client_email", order.ClientEmail)
// 	formData.Set("client_phone", order.ClientPhone)
// 	formData.Set("clientid", order.ClientID)
// 	formData.Set("expiry", order.Expiry)
// 	formData.Set("orderid", order.OrderID)
// 	formData.Set("pay_amount", fmt.Sprintf("%.2f", order.PayAmount))
// 	formData.Set("service_name", order.ServiceName)
// 	formData.Set("token", order.Token)
//
// 	req, err = http.NewRequest("POST", serverPaykeeper+invoiceURI, strings.NewReader(formData.Encode()))
// 	if err != nil {
// 		log.Fatalf("Ошибка при создании запроса для создания счёта: %v", err)
// 	}
//
// 	// Устанавливаем заголовки
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Authorization", "Basic "+auth)
//
// 	// Выполняем запрос
// 	resp, err = client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Ошибка при выполнении запроса для создания счёта: %v", err)
// 	}
// 	defer resp.Body.Close()
//
// 	// Читаем ответ
// 	body, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Ошибка при чтении ответа для создания счёта: %v", err)
// 	}
//
// 	// Парсим JSON-ответ
// 	var invoiceResponse map[string]string
// 	if err := json.Unmarshal(body, &invoiceResponse); err != nil {
// 		log.Fatalf("Ошибка при парсинге JSON для создания счёта: %v", err)
// 	}
//
// 	invoiceID, ok := invoiceResponse["invoice_id"]
// 	if !ok {
// 		log.Fatalf("Поле 'invoice_id' отсутствует в ответе")
// 	}
//
// 	// В этой переменной прямая ссылка на оплату с заданными параметрами
// 	link := fmt.Sprintf("%s/bill/%s/", serverPaykeeper, invoiceID)
//
// 	// Теперь её можно использовать как угодно, например, выводим ссылку на оплату
// 	fmt.Println("Ссылка на оплату:", link)
// }
