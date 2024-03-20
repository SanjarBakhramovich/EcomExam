package models

type Product struct {
	ID        int
	Name      string
	MainShelf int
}

type Order struct {
	ID        int
	Number    int
	ProductID int
	Quantity  int
}

type Shelf struct {
	ID   int
	Name string
}
