package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/applied-concurrency-in-go/models"
)

const simulationCount int = 19
const ordersEndpoint string = "http://localhost:3000/orders"
const productsEndpoint string = "http://localhost:3000/products"
const indexEndpoint string = "http://localhost:3000/"

const product = "MWBLU"
const expectedStock = float64(1)

type productsResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// Load test the server
func main() {
	log.Println("Welcome to the Orders App integration test!")
	log.Printf("We will now simulate %d orders. Hold onto your hats!\n", simulationCount)
	if err := checkIndex(); err != nil {
		log.Fatalf("Endpoint %s is not up. Please start the server before running simulations.", indexEndpoint)
	}
	var wg sync.WaitGroup
	wg.Add(simulationCount)
	for i := 0; i < simulationCount; i++ {
		go createOrder(i, &wg)
	}
	wg.Wait()
	verifyStock()
}

func createOrder(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	item := models.Item{
		ProductID: product,
		Amount:    1,
	}
	log.Printf("[simulation-%d]: sending order %+v", number, item)

	ibytes, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", ordersEndpoint, bytes.NewBuffer(ibytes))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[simulation-%d]: completed", number)
}

func verifyStock() {
	log.Printf("verify stock")

	req, err := http.NewRequest("GET", productsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dataResponse productsResponse
	if err := json.Unmarshal(bodyBytes, &dataResponse); err != nil {
		log.Fatal(err)
	}
	data := dataResponse.Data.([]interface{})
	for _, d := range data {
		fields := d.(map[string]interface{})
		if fields["id"].(string) == product {
			prodStock := fields["stock"].(float64)
			if prodStock != expectedStock {
				panic(fmt.Sprintf("stock for product %s not as expected: got %v, want %v",
					product, prodStock, expectedStock))
			}
		}
	}

	log.Printf("verify stock completed")
}

func checkIndex() error {
	req, err := http.NewRequest("GET", indexEndpoint, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
