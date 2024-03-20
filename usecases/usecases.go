package usecases

import (
	"EcomExam/db"
	"EcomExam/models"
	"EcomExam/repositories"
	"log"
	"strconv"
	"strings"
)

type AssemblyPageUsecase interface {
	AssemblePage(orderNumbers []string) string
}

type assemblyPageUsecase struct {
	productRepo repositories.ProductRepository
	shelfRepo   repositories.ShelfRepository
	orderRepo   repositories.OrderRepository
}

func NewAssemblyPageUsecase() AssemblyPageUsecase {
	return &assemblyPageUsecase{
		productRepo: repositories.NewProductRepository(),
		shelfRepo:   repositories.NewShelfRepository(),
		orderRepo:   repositories.NewOrderRepository(),
	}
}

func (u *assemblyPageUsecase) AssemblePage(orderNumbers []string) string {
	db, err := db.NewDB()
	var orders []models.Order
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}
	defer db.Close()

	orderIDs := convertOrderNumbersToInt(orderNumbers)
	orders, err = u.orderRepo.GetOrdersByOrderIDs(db, orderIDs)
	if err != nil {
		log.Fatalf("Error getting orders from the database: %v", err)
	}

	productIDs := extractProductIDsFromOrders(orders)
	products, err := u.productRepo.GetProductsByOrderIDs(db, productIDs)
	if err != nil {
		log.Fatalf("Error getting products from the database: %v", err)
	}

	shelves, err := u.shelfRepo.GetShelvesByProductIDs(db, productIDs)
	if err != nil {
		log.Fatalf("Error getting shelves from the database: %v", err)
	}

	return formatOutput(orders, products, shelves)
}

func convertOrderNumbersToInt(orderNumbers []string) []int {
	var orderIDs []int
	for _, num := range orderNumbers {
		id, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("Error converting order number %s to integer: %v", num, err)
		}
		orderIDs = append(orderIDs, id)
	}
	return orderIDs
}

func extractProductIDsFromOrders(orders []models.Order) []int {
	var productIDs []int
	for _, order := range orders {
		productIDs = append(productIDs, order.ProductID)
	}
	return productIDs
}

func formatOutput(orders []models.Order, products []models.Product, shelves []models.Shelf) string {
	var output strings.Builder

	output.WriteString("=+=+=+=\n")
	output.WriteString("Страница сборки заказов\n\n")

	// Добавьте обработку orders, products, shelves для формирования вывода

	return output.String()
}
