package domain

type Product struct {
	Id    int     `json:"id"`
	Date  string  `json:"date"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductResponse struct {
	Id    int     `json:"id"`
	Date  string  `json:"date"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
