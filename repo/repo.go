package repo

import (
	"fmt"
	"math"

	"github.com/applied-concurrency-in-go/db"
	"github.com/applied-concurrency-in-go/models"
)

// repo holds all the dependencies required for repo operations
type repo struct {
	products *db.ProductDB
	orders   *db.OrderDB
}

// Repo is the interface we expose to outside packages
type Repo interface {
	CreateOrder(item models.Item) (*models.Order, error)
	GetAllProducts() []models.Product
	GetOrder(id string) (models.Order, error)
}

// New creates a new Order repo with the correct database dependencies
func New() (Repo, error) {
	p, err := db.NewProducts()
	if err != nil {
		return nil, err
	}
	o := repo{
		products: p,
		orders:   db.NewOrders(),
	}
	return &o, nil
}

// GetAllProducts returns all products in the system
func (r repo) GetAllProducts() []models.Product {
	return r.products.FindAll()
}

// GetProduct returns the given order if one exists
func (r repo) GetOrder(id string) (models.Order, error) {
	return r.orders.Find(id)
}

// CreateOrder creates a new order for the given item
func (r repo) CreateOrder(item models.Item) (*models.Order, error) {
	if err := r.validateItem(item); err != nil {
		return nil, err
	}
	order := models.NewOrder(item)
	r.orders.Upsert(order)
	r.processOrders(&order)
	return &order, nil
}

// validateItem runs validations on a given order
func (r repo) validateItem(item models.Item) error {
	if item.Amount < 1 {
		return fmt.Errorf("order amount must be at least 1:got %d", item.Amount)
	}
	if err := r.products.Exists(item.ProductID); err != nil {
		return fmt.Errorf("product %s does not exist", item.ProductID)
	}
	return nil
}

func (r repo) processOrders(order *models.Order) {
	r.processOrder(order)
	r.orders.Upsert(*order)
	fmt.Printf("Processing order %s completed\n", order.ID)
}

// processOrder is an internal method which completes or rejects an order
func (r repo) processOrder(order *models.Order) {
	item := order.Item
	product, err := r.products.Find(item.ProductID)
	if err != nil {
		order.Status = models.OrderStatus_Rejected
		order.Error = err.Error()
		return
	}
	if product.Stock < item.Amount {
		order.Status = models.OrderStatus_Rejected
		order.Error = fmt.Sprintf("not enough stock for product %s:got %d, want %d", item.ProductID, product.Stock, item.Amount)
		return
	}
	remainingStock := product.Stock - item.Amount
	product.Stock = remainingStock
	r.products.Upsert(product)

	total := math.Round(float64(order.Item.Amount)*product.Price*100) / 100
	order.Total = &total
	order.Complete()
}
