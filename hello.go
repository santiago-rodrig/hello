package main

import (
	"fmt"
	"log"

	"github.com/santiago-rodrig/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Santiago")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(message)
}
