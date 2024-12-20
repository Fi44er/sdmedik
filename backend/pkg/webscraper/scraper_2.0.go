package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"golang.org/x/net/html"
)

func main() {
	url := "https://ktsr.sfr.gov.ru/ru-RU/service/compensation/product-header?region=RU-VGG&type=07&code=07-01-01"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0")
	req.Header.Set("Accept", "text/html, */*; q=0.01")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	// req.Header.Set("X-CSRF-Token", "rubRkPiNqWYbCX6QI0I-rL_-OyhhinBDZR5OZCSEsu3IrYGnoOrfA3NEMeFic23Z_awMfRjIOzYyKX4maMefjA==")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Cookie", "favorite=2aee8fa1fa5cdcc0cc132cba95f4ac1ec3fc4446c74bfcffe14bbfad7230a48fa%3A2%3A%7Bi%3A0%3Bs%3A8%3A%22favorite%22%3Bi%3A1%3Bs%3A128%3A%22_NCImZPk8xwHtPCeQY8MDN9ZHVulw-4w2pn_bK7vqK3H49mKs6JnGYcKdoqHKtCLTHFTfWRm1aE7sjB46EqtP4lku4BgWSgx-M44O29pS-adGb5EzCjvrzAePCoJmU5u%22%3B%7D; _ym_uid=1733215528322384010; _ym_d=1733215528; KTSRSESSID=WlxnoCMDTVrrwzWsQk-rA17M7jUYXipw5RYPhD-CGdMM%2CiruAoAiapXlBHELI2FdJImr8Hd1eUqYlejSApodJWTkF-kiIkCBTMj9jSNx8%2C2E%2CONAECpm07oLkH6r1N2Q; _csrf=4436a251adf1267fc7b6f44e5537c3b8a37d95e2819428371a311c3d61cacda3a%3A2%3A%7Bi%3A0%3Bs%3A5%3A%22_csrf%22%3Bi%3A1%3Bs%3A32%3A%22fKP7XgvehMOqA1SuBR7UyBKuW70BLC-a%22%3B%7D; _ym_isad=2")
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

	var certificatePrice string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "catalog-products__info-price catalog-products__info-price_space" {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode && c.Data == "span" {
							for spanChild := c.FirstChild; spanChild != nil; spanChild = spanChild.NextSibling {
								if spanChild.Type == html.TextNode {
									certificatePrice += spanChild.Data
								}
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if certificatePrice != "" {
		fmt.Println("Стоимость сертификата:", strings.TrimSpace(certificatePrice))
	} else {
		fmt.Println("Стоимость сертификата не найдена.")
	}
}
