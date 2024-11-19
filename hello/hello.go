package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// import "rsc.io/quote"
//
//	func main() {
//		fmt.Println("Hello, World!")
//	}
//
//	func main() {
//		fmt.Println(quote.Go())
//	}
func main() {
	// Get a greeting message and print it.
	//message := greetings.Hello("Gladys")
	// __________________________________
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request greeting message.
	message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
