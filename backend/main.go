package main

import (
	"fmt"
	"log"
	// "math/rand"
	"os"
	// "strconv"
	"strings"
	// "time"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/custom"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/analysis/token/ngram"
	"github.com/blevesearch/bleve/v2/mapping"
)

// Product представляет структуру товара
type Product struct {
	ID    string
	Name  string
	Price float64
}

const indexDir = "product_index" // Директория для хранения индекса

func main() {
	// Создаем или загружаем индекс
	index, err := createOrLoadIndex()
	if err != nil {
		log.Fatalf("Ошибка при создании/загрузке индекса: %v", err)
	}

	// Добавляем товары в индекс (если их нет)
	// if err := addSampleProducts(index); err != nil {
	// 	log.Fatalf("Ошибка при добавлении товаров: %v", err)
	// }

	// Выполняем поиск по индексу
	query := "тел sams тел"
	words := strings.Fields(query) // Разбиваем запрос на слова

	// Создаем BooleanQuery
	booleanQuery := bleve.NewBooleanQuery()
	for _, word := range words {
		matchQuery := bleve.NewMatchQuery(word)
		booleanQuery.AddMust(matchQuery)
	}

	searchRequest := bleve.NewSearchRequest(booleanQuery)
	searchRequest.Fields = []string{"Name", "Price"} // Указываем, какие поля включить в результаты
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		log.Fatalf("Ошибка при поиске: %v", err)
	}

	// Выводим результаты поиска
	fmt.Printf("Найдено %d результатов по запросу '%s':\n", searchResult.Total, query)
	for _, hit := range searchResult.Hits {
		name := hit.Fields["Name"]
		price := hit.Fields["Price"]
		fmt.Printf("ID: %s, Название: %v, Цена: $%.2f\n", hit.ID, name, price)
	}
}

// createOrLoadIndex создает новый индекс или загружает существующий
func createOrLoadIndex() (bleve.Index, error) {
	// Проверяем, существует ли индекс на диске
	if _, err := os.Stat(indexDir); os.IsNotExist(err) {
		// Индекс не существует, создаем новый
		return createIndex()
	}

	// Индекс существует, загружаем его
	return bleve.Open(indexDir)
}

// createIndex создает новый индекс с n-gram токенизацией
func createIndex() (bleve.Index, error) {
	// Создаем маппинг для индекса
	indexMapping := bleve.NewIndexMapping()
	// Регистрируем n-gram токенизатор
	err := indexMapping.AddCustomTokenFilter("ngram3", map[string]interface{}{
		"type": ngram.Name,
		"min":  3,
		"max":  3,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при регистрации токенизатора: %v", err)
	}

	// Регистрируем lowercase фильтр
	err = indexMapping.AddCustomTokenFilter("lowercase", map[string]interface{}{
		"type": lowercase.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при регистрации lowercase фильтра: %v", err)
	}

	// Создаем кастомный анализатор
	err = indexMapping.AddCustomAnalyzer("custom", map[string]interface{}{
		"type":      custom.Name,
		"tokenizer": "unicode", // Используем стандартный токенизатор
		"token_filters": []string{
			"lowercase", // Приводим текст к нижнему регистру
			"ngram3",    // Применяем n-gram фильтр
		},
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при добавлении кастомного анализатора: %v", err)
	}
	// Устанавливаем кастомный анализатор по умолчанию
	indexMapping.DefaultAnalyzer = "custom"

	// Создаем маппинг для товаров
	productMapping := bleve.NewDocumentMapping()
	productMapping.AddFieldMappingsAt("Name", mapping.NewTextFieldMapping())
	// productMapping.AddFieldMappingsAt("Price", mapping.NewNumericFieldMapping())

	// Добавляем маппинг товаров в индекс
	indexMapping.AddDocumentMapping("product", productMapping)

	// Создаем новый индекс на диске
	index, err := bleve.New(indexDir, indexMapping)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании индекса: %v", err)
	}

	return index, nil
}

// // addSampleProducts добавляет тестовые товары в индекс
// func addSampleProducts(index bleve.Index) error {
// 	// Проверяем, есть ли уже товары в индексе
// 	count, err := index.DocCount()
// 	if err != nil {
// 		return fmt.Errorf("ошибка при проверке количества документов: %v", err)
// 	}
//
// 	if count > 0 {
// 		// Товары уже есть, ничего не делаем
// 		return nil
// 	}
//
// 	// Генерируем и добавляем товары
// 	products := generateProducts(10000) // Генерируем 10 000 товаров
// 	for _, product := range products {
// 		if err := index.Index(product.ID, product); err != nil {
// 			return fmt.Errorf("ошибка при индексации товара: %v", err)
// 		}
// 	}
//
// 	return nil
// }
//
// // generateProducts генерирует список товаров
// func generateProducts(count int) []Product {
// 	rand.Seed(time.Now().UnixNano()) // Инициализируем генератор случайных чисел
// 	products := make([]Product, 0, count)
//
// 	brands := []string{"Samsung", "Apple", "Sony", "Xiaomi", "Huawei", "LG", "Asus", "Dell", "HP", "Lenovo"}
// 	categories := []string{"Смартфон", "Ноутбук", "Наушники", "Планшет", "Монитор", "Телевизор", "Фотоаппарат", "Часы", "Клавиатура", "Мышь"}
//
// 	for i := 1; i <= count; i++ {
// 		brand := brands[rand.Intn(len(brands))]
// 		category := categories[rand.Intn(len(categories))]
// 		name := fmt.Sprintf("%s %s %d", category, brand, i)
// 		price := 100.0 + rand.Float64()*1000.0 // Генерируем случайную цену от 100 до 1100
//
// 		products = append(products, Product{
// 			ID:    strconv.Itoa(i),
// 			Name:  name,
// 			Price: price,
// 		})
// 	}
//
// 	return products
// }
