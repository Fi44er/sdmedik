package ut``

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// RequestOptions содержит параметры для выполнения запроса
type RequestOptions struct {
	Method  string            // Тип запроса (GET, POST, PUT, DELETE и т.д.)
	URL     string            // URL запроса
	Query   map[string]string // Query-параметры
	Headers map[string]string // Заголовки запроса
	Body    interface{}       // Тело запроса (может быть nil)
}

// MakeRequest выполняет HTTP-запрос и возвращает ответ
func MakeRequest(options RequestOptions) ([]byte, error) {
	// Добавляем query-параметры к URL, если они есть
	if len(options.Query) > 0 {
		queryParams := url.Values{}
		for key, value := range options.Query {
			queryParams.Add(key, value)
		}
		options.URL += "?" + queryParams.Encode()
	}

	// Преобразуем тело запроса в JSON, если оно есть
	var bodyReader io.Reader
	if options.Body != nil {
		bodyBytes, err := json.Marshal(options.Body)
		if err != nil {
			return nil, fmt.Errorf("ошибка при сериализации тела запроса: %v", err)
		}
		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	// Создаем новый запрос
	req, err := http.NewRequest(options.Method, options.URL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	// Добавляем заголовки
	for key, value := range options.Headers {
		req.Header.Add(key, value)
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	// Проверяем статус ответа
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("ошибка сервера: %s, тело ответа: %s", resp.Status, string(responseBody))
	}

	return responseBody, nil
}

func main() {
	// Пример использования функции

	// GET-запрос с query-параметрами
	options := RequestOptions{
		Method: "GET",
		URL:    "https://esnsi.gosuslugi.ru/rest/ext/v1/classifiers/10616/data",
		Query: map[string]string{
			"query": "цйвцйв",
		},
	}

	response, err := MakeRequest(options)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Ответ GET-запроса:", string(response))
}
