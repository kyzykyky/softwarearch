package domain

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Stock struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
