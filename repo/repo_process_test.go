package repo

import (
	"fmt"
	"sync"
	"testing"

	"github.com/applied-concurrency-in-go/db"
	"github.com/applied-concurrency-in-go/models"
	"github.com/stretchr/testify/assert"
)

const productCode = "TEST"
const productStock = 11

// how many goroutines we will place orders on
const concurrentOrders = 10

// THIS TEST IS FLAKY. FOR DEMO PURPOSES ONLY
func Test_ProcessOrder(t *testing.T) {
	// Uncomment out line below to skip it
	// t.Skip("Skipping process Order test")

	prod := &db.ProductDB{}
	prod.Upsert(models.Product{
		ID:    productCode,
		Stock: productStock,
	})
	r := &repo{
		orders:   db.NewOrders(),
		products: prod,
	}
	item := models.Item{
		ProductID: productCode,
		Amount:    1,
	}

	t.Run(fmt.Sprintf("%d concurrent orders", concurrentOrders), func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrentOrders)
		for j := 0; j < concurrentOrders; j++ {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				order := models.NewOrder(item)
				r.processOrders(&order)
			}(&wg)
		}
		wg.Wait()
		expected := productStock - concurrentOrders
		assertStock(t, r, expected)
	})

}

func assertStock(t *testing.T, r *repo, expectedStock int) {
	prod, err := r.products.Find(productCode)
	assert.Nil(t, err)
	assert.Equal(t, expectedStock, prod.Stock)
}
