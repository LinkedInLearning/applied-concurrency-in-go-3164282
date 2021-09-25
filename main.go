package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel
	ch := make(chan string)
	// start the greeter to provide a greeting
	go greet(ch)
	// sleep for a long time
	time.Sleep(10 * time.Second)
	fmt.Println("Main ready!")
	// receive greeting
	greeting := <-ch
	// sleep and print
	time.Sleep(2 * time.Second)
	fmt.Println(greeting)

}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan string) {
	fmt.Println("Greeter ready!")
	// greet
	ch <- "Hello, world!"
	fmt.Println("Greeter completed!")
}
