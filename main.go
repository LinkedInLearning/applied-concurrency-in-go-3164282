package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/applied-concurrency-in-go/counter"
)

func main() {
	fmt.Println("Welcome to the Word Counter!")
	// read input file
	lines, err := readInput("./counter/landon_hotel.txt")
	if err != nil {
		fmt.Println("Could not read file:", err)
		return
	}

	// count the from the file
	stats := countWords(lines)

	// print the stats to the console
	stats.PrintStats()
	fmt.Println("Goodbye from the Word Counter!")
}

// countWords counts all the word occurrences slice of lines
// returns a map of words and their counts
func countWords(lines []string) counter.Statistics {
	stats := counter.New()
	for _, l := range lines {
		countLine(stats, l)
	}
	return *stats
}

// countLine counts all the word occurrences in a single line
// returns a map of of words and their counts
func countLine(stats *counter.Statistics, line string) {
	// split the line into words
	lineWords := strings.Split(line, " ")
	for _, w := range lineWords {
		stats.AddWord(w)
	}
}

// readInput reads a file at the given path
// returns a slice of lines from the file or an error
func readInput(filePath string) ([]string, error) {
	// open the file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// create a lines slice to save each line into
	var lines []string
	// use a scanner to read the file in chunks
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
