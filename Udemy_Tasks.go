package main

import (
	"fmt"
	"time"
)

func printsomething(s string) {
	fmt.Println(s)
}

func main() {
	//declared a slice
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	for i, x := range words {
		go printsomething(fmt.Sprintf("%d: %s", i, x))
	}

	time.Sleep(1 * time.Second)

	printsomething("This is the second thing to be printed")

}
