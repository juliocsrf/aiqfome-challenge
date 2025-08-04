package repository

type FakestoreapiProductResponse struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int64   `json:"count"`
	} `json:"rating"`
}
