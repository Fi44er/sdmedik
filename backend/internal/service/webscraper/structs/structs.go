package structs

type Element struct {
	Order     int         `json:"order"`
	Value     string      `json:"value"`
	Reference interface{} `json:"reference"` // Используем interface{}, так как reference всегда null
}

// BodyItem описывает объект в массиве body
type BodyItem struct {
	Elements       []Element   `json:"elements"`
	ID             int         `json:"id"`
	UID            string      `json:"uid"`
	ClassifierID   int         `json:"classifierId"`
	ClassifierName string      `json:"classifierName"`
	Hierarchical   interface{} `json:"hierarchical"` // Используем interface{}, так как hierarchical всегда null
}

// HeaderItem описывает объект в массиве header
type HeaderItem struct {
	Order            int    `json:"order"`
	Value            string `json:"value"`
	ID               int    `json:"id"`
	Type             string `json:"type"`
	Key              bool   `json:"key"`
	Autokey          bool   `json:"autokey"`
	ParentKey        bool   `json:"parentKey"`
	AutokeyBase      bool   `json:"autokeyBase"`
	AutokeyBaseOrder int    `json:"autokeyBaseOrder"`
}

// ApiResponse описывает структуру всего JSON-ответа
type ApiResponse struct {
	Errors      []interface{} `json:"errors"` // Используем interface{}, так как errors всегда пустой массив
	Header      []HeaderItem  `json:"header"`
	Body        []BodyItem    `json:"body"`
	Total       int           `json:"total"`
	CurrentPage int           `json:"currentPage"`
}
