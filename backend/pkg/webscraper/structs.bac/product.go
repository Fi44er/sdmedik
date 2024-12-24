package structs

type Product struct {
	Article string `json:"article"`
	Price   string `json:"price"`
	Region  string `json:"region"`
}

type Products struct {
	Article      string        `json:"article"`
	RegionPrices []RegionPrice `json:"region_prices"`
}

type RegionPrice struct {
	Region string `json:"region"`
	Price  string `json:"price"`
}
