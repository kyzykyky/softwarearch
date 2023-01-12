package domain

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Isbn   string  `json:"isbn"`
	Price  float32 `json:"price"`
}
