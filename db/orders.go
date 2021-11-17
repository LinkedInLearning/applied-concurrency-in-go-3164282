package db

import (
	"fmt"
	"sync"

	"github.com/applied-concurrency-in-go/models"
)

type OrderDB struct {
	placedOrders sync.Map
}

// NewOrders creates a new empty order service
func NewOrders() *OrderDB {
	return &OrderDB{}
}

// Find order for a given id, if one exists
func (o *OrderDB) Find(id string) (models.Order, error) {
	po, ok := o.placedOrders.Load(id)
	if !ok {
		return models.Order{}, fmt.Errorf("no order found for %s order id", id)
	}

	return toOrder(po), nil
}

// Upsert creates or updates an order in the orders DB
func (o *OrderDB) Upsert(order models.Order) {
	o.placedOrders.Store(order.ID, order)
}

// toOrder attempts to convert an interface to an order
// panics if if this not possible
func toOrder(po interface{}) models.Order {
	order, ok := po.(models.Order)
	if !ok {
		panic(fmt.Errorf("error casting %v to order", po))
	}
	return order
}