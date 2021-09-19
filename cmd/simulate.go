package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/applied-concurrency-in-go/models"
)

const simulationCount int = 50
const ordersEndpoint string = "http://localhost:3000/orders"
const indexEndpoint string = "http://localhost:3000/"
const maxOrderAmount int = 15

var products []string = []string{"MWBLU", "MWLEM", "MWORG", "MWPEA", "MWRAS", "MWSTR", "MWCRA", "MWMAN"}

// Load test the server
func main() {
	log.Println("Welcome to the Orders App simulator!")
	log.Printf("We will now simulate %d orders. Hold onto your hats!\n", simulationCount)
	if err := checkIndex(); err != nil {
		log.Fatalf("Endpoint %s is not up. Please start the server before running simulations.", indexEndpoint)
	}
	var wg sync.WaitGroup
	wg.Add(simulationCount)
	for i := 0; i < simulationCount; i++ {
		go createRandomOrder(i, &wg)
	}
	wg.Wait()
}

func createRandomOrder(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	amount := rand.Intn(maxOrderAmount) + 1
	product := products[rand.Intn(len(products))]
	item := models.Item{
		ProductID: product,
		Amount:    amount,
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
