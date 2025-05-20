package order

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/Fi44er/sdmedik/backend/internal/response"
	"github.com/Fi44er/sdmedik/backend/pkg/errors"
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

func (s *service) sendToPaykeeper(
	ctx context.Context,
	data *dto.CreateOrder,
	basket *response.BasketResponse,
	articles []dto.GetManyCert,
	userID string,
) (string, *model.Order, error) {

	// s.logger.Infof("Basket: %+v", basket)
	// s.logger.Infof("Articles: %+v", articles)
	// s.logger.Infof("User ID: %v", userID)
	// s.logger.Infof("Data: %+v", data)

	user := s.config.PayKeeperUser
	password := s.config.PayKeeperPass
	serverPaykeeper := s.config.PayKeeperServer

	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	certs, err := s.certService.GetMany(ctx, &articles)
	if err != nil {
		return "", nil, err
	}

	certMap := make(map[string]string)

	for _, cert := range *certs {
		certMap[cert.CategoryArticle] = cert.TRU
	}

	carts := []CartItem{}
	for _, item := range basket.Items {
		categoryArticle := strings.Split(item.Article, ".")[0]
		truCode := certMap[categoryArticle] + "00000000643"
		cartItem := CartItem{
			ItemType:    "goods",
			PaymentType: "full",
			SKU:         "",
			Name:        item.Name,
			Price:       item.Price,
			Quantity:    item.Quantity,
			ItemCode:    "",
			TruCode:     truCode,
			Tax:         "none",
			Sum:         item.TotalPrice,
		}
		carts = append(carts, cartItem)
	}

	expireDate := time.Now().AddDate(0, 0, 3) // Текущая дата + 1 день
	expire := expireDate.Format("2006-01-02")
	cartsJson, err := json.MarshalIndent(carts, "", "  ")
	if err != nil {
		return "", nil, err
	}
	serviceName := ";PKC|" + string(cartsJson) + "|"

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
		return "", nil, err
	}

	var tokenResponse map[string]string
	if err := json.Unmarshal(tokenBody, &tokenResponse); err != nil {
		s.logger.Errorf("Ошибка при парсинге JSON для получения токена: %v", err)
		return "", nil, err
	}

	token, ok := tokenResponse["token"]
	if !ok {
		s.logger.Errorf("Поле 'token' отсутствует в ответе")
		return "", nil, fmt.Errorf("Поле 'token' отсутствует в ответе")
	}

	invoiceURI := "/change/invoice/preview/"

	order.Token = token

	test := map[string]string{
		"cart_json":    string(cartsJson),
		"client_email": order.ClientEmail,
		"client_phone": order.ClientPhone,
		"clientid":     order.ClientID,
		"expiry":       order.Expiry,
		"orderid":      order.OrderID,
		"pay_amount":   fmt.Sprintf("%.2f", order.PayAmount),
		"service_name": order.ServiceName,
		"token":        order.Token,
	}

	// Преобразуем в JSON с отступами
	jsonData, _ := json.MarshalIndent(test, "", "    ")

	// Выводим результат
	fmt.Println(string(jsonData))

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
		return "", nil, err
	}

	var invoiceResponse map[string]string
	if err := json.Unmarshal(invoiceBody, &invoiceResponse); err != nil {
		s.logger.Errorf("Ошибка при парсинге JSON для создания счёта: %v", err)
		return "", nil, err
	}

	invoiceID, ok := invoiceResponse["invoice_id"]
	if !ok {
		s.logger.Errorf("Поле 'invoice_id' отсутствует в ответе")
		s.logger.Errorf("Invoice res: %+v", invoiceResponse)
		return "", nil, errors.New(400, invoiceResponse["msg"])
	}

	link := fmt.Sprintf("%s/bill/%s/", serverPaykeeper, invoiceID)

	orderModel := model.Order{
		TotalPrice: basket.TotalPrice,
		Email:      data.Email,
		Phone:      data.PhoneNumber,
		FIO:        data.FIO,
		Address:    data.Address,
		Status:     "pending",
	}

	if userID != "" {
		orderModel.UserID = sql.NullString{String: userID, Valid: true}
	}
	if err := s.repo.Create(ctx, &orderModel); err != nil {
		return "", nil, err
	}

	return link, &orderModel, nil
}
