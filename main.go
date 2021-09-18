package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	goodbye()
}

func hello() {
	fmt.Println("Hello, world!")
}

func goodbye() {
	fmt.Println("Goodbye, world!")
}
