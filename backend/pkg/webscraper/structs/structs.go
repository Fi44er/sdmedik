package structs

type Category struct {
	URL  string
	Name string
}

type Item struct {
	Price  float64
	Region string
}

type Items struct {
	CategoryArticle string
	CategoryName    string
	Items           []Item
	Articles        []ParseProductsArticlesType
}

type ParseProductsArticlesType struct {
	Article string
	Name    string
}
