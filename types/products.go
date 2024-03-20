package types

type Product struct {
	Id          int
	Name        string
	Category    string
	Price       float32
	Description string
	Visible     bool
}

type ProductList = map[string][]Product
