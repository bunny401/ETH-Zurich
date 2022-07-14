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

}
