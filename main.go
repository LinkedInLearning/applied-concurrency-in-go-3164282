package main

import (
	"fmt"
	"time"
)

// greetings in many languages
var greetings = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}

func main() {
	// create a channel
	ch := make(chan string, 1)
	// start the greeter to provide a greeting
	go greet(ch)
	// sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")
	for {
		// receive greeting
		greeting := <-ch
		// sleep and print
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Greeting received!", greeting)
	}

}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan<- string) {
	fmt.Println("Greeter ready!")
	// greet
	for _, g := range greetings {
		ch <- g
	}
	fmt.Println("Greeter completed!")
}
