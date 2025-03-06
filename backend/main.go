package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

func main() {
	// Инициализация Fiber приложения
	app := fiber.New()

	storage := redis.New(redis.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "", // если есть пароль, укажите его здесь
		Database: 0,  // используйте 0 для дефолтной БД Redis
		Reset:    false,
	})

	// Инициализация хранилища сессий в памяти
	store := session.New(session.Config{
		Storage: storage,
	})

	// Маршрут для добавления элемента в корзину
	app.Post("/cart/add", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при получении сессии")
		}

		log.Errorf("%+v", sess)

		// Получаем текущую корзину из сессии
		var cart []string
		if sess.Get("cart") != nil {
			cart = sess.Get("cart").([]string)
		}

		// Получаем элемент из тела запроса
		var item struct {
			Item string `json:"item"`
		}
		if err := c.BodyParser(&item); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Неверный формат данных")
		}

		// Добавляем элемент в корзину
		cart = append(cart, item.Item)
		sess.Set("cart", cart)

		// Сохраняем сессию
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при сохранении сессии")
		}

		return c.SendString("Элемент добавлен в корзину")
	})

	// Маршрут для удаления элемента из корзины
	app.Delete("/cart/remove", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при получении сессии")
		}

		// Получаем текущую корзину из сессии
		var cart []string
		if sess.Get("cart") != nil {
			cart = sess.Get("cart").([]string)
		} else {
			return c.Status(fiber.StatusNotFound).SendString("Корзина пуста")
		}

		// Получаем элемент из тела запроса
		var item struct {
			Item string `json:"item"`
		}
		if err := c.BodyParser(&item); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Неверный формат данных")
		}

		// Удаляем элемент из корзины
		var newCart []string
		for _, i := range cart {
			if i != item.Item {
				newCart = append(newCart, i)
			}
		}

		sess.Set("cart", newCart)

		// Сохраняем сессию
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при сохранении сессии")
		}

		return c.SendString("Элемент удален из корзины")
	})

	// Маршрут для получения текущей корзины
	app.Get("/cart", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Ошибка при получении сессии")
		}

		// Получаем текущую корзину из сессии
		var cart []string
		if sess.Get("cart") != nil {
			cart = sess.Get("cart").([]string)
		} else {
			return c.Status(fiber.StatusNotFound).SendString("Корзина пуста")
		}

		return c.JSON(cart)
	})

	// Запуск сервера
	app.Listen(":3000")
}
