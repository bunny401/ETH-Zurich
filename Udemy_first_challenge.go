package main

import (
	"fmt"
)

var msg string

func updateMessage(s string) {

	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 27, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	go updateMessage("Hello, universe!")
	printMessage()

	go updateMessage("Hello, cosmos!")
	printMessage()

	go updateMessage("Hello, world!")
	printMessage()
}
