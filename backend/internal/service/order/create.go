package order

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/pkg/utils"
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
}

type Order struct {
	CartJSON    []CartItem `json:"cart_json"`
	ClientEmail string     `json:"client_email"`
	ClientPhone string     `json:"client_phone"`
	ClientID    string     `json:"clientid"`
	Expiry      string     `json:"expiry"`
	OrderID     string     `json:"orderid"`
	PayAmount   float64    `json:"pay_amount"`
	ServiceName string     `json:"service_name"`
	Token       string     `json:"token"`
}

func (s *service) Create(ctx context.Context, data *dto.CreateOrder, userID string) (string, error) {
	if err := s.validator.Struct(data); err != nil {
		return "", err
	}

	user := s.config.PayKeeperUser
	password := s.config.PayKeeperPass
	serverPaykeeper := s.config.PayKeeperServer

	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	basket, err := s.basketService.GetByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	articles := make([]dto.GetManyCert, len(basket.Items))
	for _, item := range basket.Items {
		categoryArticle := strings.Split(item.Article, ".")[0]
		articles = append(articles, dto.GetManyCert{CategoryArticle: categoryArticle})
	}

	certs, err := s.certService.GetMany(ctx, &articles)
	if err != nil {
		return "", err
	}

	certMap := make(map[string]string)

	for _, cert := range *certs {
		certMap[cert.CategoryArticle] = cert.TRU
	}

	carts := make([]CartItem, len(basket.Items))
	for _, item := range basket.Items {
		categoryArticle := strings.Split(item.Article, ".")[0]

		cartItem := CartItem{
			ItemType:    "goods",
			PaymentType: "full",
			SKU:         "",
			Name:        item.Name,
			Price:       item.Price,
			Quantity:    item.Quantity,
			ItemCode:    "",
			TruCode:     certMap[categoryArticle],
			Tax:         "none",
			Sum:         basket.TotalPrice,
		}
		carts = append(carts, cartItem)
	}

	jsonData, err := json.Marshal(carts)
	if err != nil {
		s.logger.Errorf("Ошибка при парсинге JSON для создания заказа: %v", err)
		return "", err
	}

	expireDate := time.Now().AddDate(0, 0, 1) // Текущая дата + 1 день
	expire := expireDate.Format("2006-01-02")
	serviceName := fmt.Sprintf(";PKC|%s|", jsonData)
	order := Order{
		CartJSON:    carts,
		ClientEmail: data.Email,
		ClientPhone: data.PhoneNumber,
		ClientID:    data.FIO,
		Expiry:      expire,
		OrderID:     "",
		PayAmount:   basket.TotalPrice,
		ServiceName: serviceName,
		Token:       "",
	}

	tokenURI := "/info/settings/token/"

	options := utils.RequestOptions{
		Method: "GET",
		URL:    serverPaykeeper + tokenURI,
		Headers: map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Basic " + auth,
		},
	}
	tokenBody, err := utils.MakeRequest(options)
	if err != nil {
		return "", err
	}

	var tokenResponse map[string]string
	if err := json.Unmarshal(tokenBody, &tokenResponse); err != nil {
		s.logger.Errorf("Ошибка при парсинге JSON для получения токена: %v", err)
		return "", err
	}

	// В ответе должно быть заполнено поле token, иначе - ошибка
	token, ok := tokenResponse["token"]
	if !ok {
		s.logger.Errorf("Поле 'token' отсутствует в ответе")
		return "", fmt.Errorf("Поле 'token' отсутствует в ответе")
	}

	// Готовим запрос 3.4 JSON API на получение счёта
	invoiceURI := "/change/invoice/preview/"

	// Добавляем токен в структуру заказа
	order.Token = token

	options = utils.RequestOptions{
		Method: "POST",
		URL:    serverPaykeeper + invoiceURI,
		Headers: map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Basic " + auth,
		},
		FormData: map[string]string{
			"cart_json":    fmt.Sprintf("%v", order.CartJSON),
			"client_email": order.ClientEmail,
			"client_phone": order.ClientPhone,
			"clientid":     order.ClientID,
			"expiry":       order.Expiry,
			"orderid":      order.OrderID,
			"pay_amount":   fmt.Sprintf("%.2f", order.PayAmount),
			"service_name": order.ServiceName,
			"token":        order.Token,
		},
	}

	invoiceBody, err := utils.MakeRequest(options)
	if err != nil {
		return "", err
	}
	// Парсим JSON-ответ
	var invoiceResponse map[string]string
	if err := json.Unmarshal(invoiceBody, &invoiceResponse); err != nil {
		s.logger.Errorf("Ошибка при парсинге JSON для создания счёта: %v", err)
		return "", err
	}

	invoiceID, ok := invoiceResponse["invoice_id"]
	if !ok {
		s.logger.Errorf("Поле 'invoice_id' отсутствует в ответе")
		return "", fmt.Errorf("Поле 'invoice_id' отсутствует в ответе")
	}

	link := fmt.Sprintf("%s/bill/%s/", serverPaykeeper, invoiceID)

	return link, nil
}
