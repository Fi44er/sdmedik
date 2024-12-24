package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Fi44er/sdmedik/backend/pkg/webscraper/utils"
	"github.com/andybalholm/brotli"
	"golang.org/x/net/html"
)

func main() {
	region := "RU-VGG"
	article := "07-01-01"
	articleType := strings.Split(article, "-")[0]
	url := fmt.Sprintf("https://ktsr.sfr.gov.ru/ru-RU/service/compensation/product-header?region=%v&type=%v&code=%v", region, articleType, article)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0")
	req.Header.Set("Accept", "text/html, */*; q=0.01")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	br := brotli.NewReader(bytes.NewReader(body))
	decodedBody, err := io.ReadAll(br)
	if err != nil {
		fmt.Println("Ошибка при декодировании Brotli:", err)
		return
	}

	doc, err := html.Parse(strings.NewReader(string(decodedBody)))
	if err != nil {
		fmt.Println("Ошибка при парсинге HTML:", err)
		return
	}

	certificatePrice := utils.ParcePrice(doc)

	if certificatePrice != 0 {
		fmt.Println("Стоимость сертификата:", certificatePrice)
	} else {
		fmt.Println("Стоимость сертификата не найдена.")
	}
}
