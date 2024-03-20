package repositories

import (
	"EcomExam/models"
	"database/sql"
	"fmt"
	"strings"
)

type ProductRepository interface {
	GetProductsByOrderIDs(db *sql.DB, orderIDs []int) ([]models.Product, error)
}

type ShelfRepository interface {
	GetShelvesByProductIDs(db *sql.DB, productIDs []int) ([]models.Shelf, error)
}

type OrderRepository interface {
	GetOrdersByOrderIDs(db *sql.DB, orderIDs []int) ([]models.Order, error)
}

type productRepo struct{}

func NewProductRepository() ProductRepository {
	return &productRepo{}
}

func (r *productRepo) GetProductsByOrderIDs(db *sql.DB, orderIDs []int) ([]models.Product, error) {
	query := fmt.Sprintf("SELECT id, название, стеллаж FROM Товары WHERE id IN (%s);", joinInts(orderIDs))

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.MainShelf)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

type shelfRepo struct{}

func NewShelfRepository() ShelfRepository {
	return &shelfRepo{}
}

func (r *shelfRepo) GetShelvesByProductIDs(db *sql.DB, productIDs []int) ([]models.Shelf, error) {
	query := fmt.Sprintf("SELECT id, название FROM Стеллажи WHERE id IN (%s);", joinInts(productIDs))

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var shelves []models.Shelf
	for rows.Next() {
		var shelf models.Shelf
		err := rows.Scan(&shelf.ID, &shelf.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		shelves = append(shelves, shelf)
	}

	return shelves, nil
}

type orderRepo struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepo{}
}

func (r *orderRepo) GetOrdersByOrderIDs(db *sql.DB, orderIDs []int) ([]models.Order, error) {
	query := fmt.Sprintf("SELECT id, номер, товар, количество FROM Заказы WHERE номер IN (%s);", joinInts(orderIDs))

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Number, &order.ProductID, &order.Quantity)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func joinInts(ints []int) string {
	var strInts []string
	for _, num := range ints {
		strInts = append(strInts, fmt.Sprintf("%d", num))
	}
	return strings.Join(strInts, ",")
}
