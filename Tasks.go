package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	var out = x + y
	return out

}

func main() {
	var num1 int = 3
	var num2 int = 5

	var num3 float64 = 9
	var result = math.Sqrt(num3)

	fmt.Println(add(num1, num2))
	fmt.Println(result)

	//Added If-Else
	var ass int = 5
	if ass <= 4 {
		fmt.Println("Hi", ass)

	} else {
		fmt.Println("Bye", ass)

	}

	var number int = 3
	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Error 404")

	}

}
