package db

import (
	"fmt"
	"sort"

	"github.com/applied-concurrency-in-go/models"
	"github.com/applied-concurrency-in-go/utils"
)

type ProductDB struct {
	products map[string]models.Product
}

// NewProducts creates a new empty product DB
func NewProducts() (*ProductDB, error) {
	p := &ProductDB{
		products: make(map[string]models.Product),
	}
	// load start position
	if err := utils.ImportProducts(p.products); err != nil {
		return nil, err
	}

	return p, nil
}

// Exists checks whether a product with a give id exists
func (p *ProductDB) Exists(id string) error {
	if _, ok := p.products[id]; !ok {
		return fmt.Errorf("no product found for id %s", id)
	}

	return nil
}

// Find returns a given product if one exists
func (p *ProductDB) Find(id string) (models.Product, error) {
	prod, ok := p.products[id]
	if !ok {
		return models.Product{}, fmt.Errorf("no product found for id %s", id)
	}

	return prod, nil
}

// Upsert creates or updates a product in the orders DB
func (p *ProductDB) Upsert(prod models.Product) {
	p.products[prod.ID] = prod
}

// FindAll returns all products in the system
func (p *ProductDB) FindAll() []models.Product {
	var allProducts []models.Product
	for _, product := range p.products {
		allProducts = append(allProducts, product)
	}
	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].ID < allProducts[j].ID
	})
	return allProducts
}
