package model

type Trade struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Id     uint64  `json:"id"`
	Market string
	Move   string
}

type Book struct {
	Bids   [][2]string `json:"bids"`
	Asks   [][2]string `json:"asks"`
	Market string
}
