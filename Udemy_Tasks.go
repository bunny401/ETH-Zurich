package main

import (
	"fmt"
	"sync"
)

func printsomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

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
	wg.Add(len(words))

	for i, x := range words {
		go printsomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()
	wg.Add(1)

	//time.Sleep(1 * time.Second)

	printsomething("This is the second thing to be printed", &wg)

}
