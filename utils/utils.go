package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/applied-concurrency-in-go/models"
)

const productInputPath string = "./input/products.csv"

// importProducts imports the start position of the products DB
func ImportProducts(products map[string]models.Product) error {
	input, err := readCsv(productInputPath)
	if err != nil {
		return err
	}

	// each line in the file is a product
	for _, line := range input {
		// bad csv line continue import
		if len(line) != 5 {
			continue
		}
		id := line[0]
		stock, err := strconv.Atoi(line[2])
		// bad csv line continue import
		if err != nil {
			continue
		}
		price, err := strconv.ParseFloat(line[4], 64)
		// bad csv line continue import
		if err != nil {
			continue
		}
		products[id] = models.Product{
			ID:    id,
			Name:  fmt.Sprintf("%s(%s)", line[1], line[3]),
			Stock: stock,
			Price: price,
		}
	}
	return nil
}

// readCsv reads file and converts it to an array of strings
// the format of the csv file is hardcoded so we can take some
// error handling liberties for the sake of brevity
func readCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
