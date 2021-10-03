package main

import (
	"fmt"
	"time"
)

// greetings in many languages
var hellos = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}
var goodbyes = []string{"Goodbye!", "Arrivederci!", "Adios!", "Hej Hej!", "La revedere!"}

func main() {
	// create a channel
	ch := make(chan string, 1)
	ch2 := make(chan string, 1)
	// start the greeter to provide a greeting
	go greet(hellos, ch)
	go greet(goodbyes, ch2)
	// sleep for a long time
	time.Sleep(1 * time.Second)
	fmt.Println("Main ready!")
	for {
		select {
		case gr, ok := <-ch:
			if !ok {
				ch = nil
				break
			}
			printGreeting(gr)
		case gr2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			printGreeting(gr2)
		default:
			return
		}
	}
}

// greet writes a greet to the given channel and then says goodbye
func greet(greetings []string, ch chan<- string) {
	fmt.Println("Greeter ready!")
	// greet
	for _, g := range greetings {
		ch <- g
	}
	close(ch)
	fmt.Println("Greeter completed!")
}

// printGreeting sleeps and prints the greeting given
func printGreeting(greeting string) {
	// sleep and print
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Greeting received!", greeting)
}
