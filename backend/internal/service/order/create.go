package order

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/Fi44er/sdmedik/backend/internal/dto"
	"github.com/Fi44er/sdmedik/backend/internal/model"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (s *service) Create(ctx context.Context, data *dto.CreateOrder, userID string, sess *session.Session) (string, error) {
	if err := s.validator.Struct(data); err != nil {
		return "", err
	}

	basket, err := s.basketService.GetByUserID(ctx, userID, sess)
	if err != nil {
		return "", err
	}

	articles := make([]dto.GetManyCert, len(basket.Items))

	for _, item := range basket.Items {
		categoryArticle := strings.Split(item.Article, ".")[0]
		articles = append(articles, dto.GetManyCert{CategoryArticle: categoryArticle, RegionIso: item.Iso})
	}

	link, orderModel, err := s.sendToPaykeeper(ctx, data, basket, articles, userID)
	if err != nil {
		return "", err
	}

	chatID := userID
	if chatID == "" {
		chatID = sess.ID()
	}

	fragmentID, err := s.chatService.AddEndMsgID(ctx, chatID)
	if err != nil {
		return "", err
	}

	if userID != "" {
		orderModel.UserID = sql.NullString{String: userID, Valid: true}
	}

	orderModel.FragmentLink = s.config.FrontendURL + "/admin/admin_chat?fragment=" + fragmentID
	if err := s.repo.Create(ctx, orderModel); err != nil {
		return "", err
	}

	s.logger.Infof("Link: %v", link)

	orderItems := []model.OrderItem{}
	templateData := struct {
		Date          string
		ClientName    string
		ClientPhone   string
		ClientEmail   string
		ClientAddress string
		Data          struct {
			Items      []model.OrderItem
			TotalPrice float64
		}
	}{
		Date:          string(time.Now().Format("02.01.2006")),
		ClientName:    data.FIO,
		ClientPhone:   data.PhoneNumber,
		ClientEmail:   data.Email,
		ClientAddress: data.Address,
		Data: struct {
			Items      []model.OrderItem
			TotalPrice float64
		}{
			Items:      make([]model.OrderItem, 0),
			TotalPrice: 0,
		},
	}
	for _, item := range basket.Items {
		orderItem := model.OrderItem{
			OrderID:         orderModel.ID,
			Name:            item.Name,
			Price:           item.Price,
			Quantity:        item.Quantity,
			TotalPrice:      item.TotalPrice,
			ProductID:       item.ProductID,
			SelectedOptions: item.SelectedOptions,
		}
		orderItems = append(orderItems, orderItem)
		templateData.Data.Items = append(templateData.Data.Items, orderItem)
		templateData.Data.TotalPrice += orderItem.TotalPrice

		if err := s.basketService.DeleteItem(ctx, item.ID, userID, sess); err != nil {
			return "", err
		}
	}

	if err := s.repo.AddItems(ctx, &orderItems); err != nil {
		return "", err
	}

	s.mailer.SendMailAsync(
		s.config.MailFrom,
		"Новый заказ",
		templateData,
		[]string{"sales@sdmedik.ru", "amanager@sdmedik.ru", "admin@sdmedik.ru"},
	)

	return link, nil
}
