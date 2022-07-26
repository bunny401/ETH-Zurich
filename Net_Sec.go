package main

import (
	"fmt"
	"math/rand"
)

type packet struct {
	payload byte
	type1   int
}

func checktype() {
	rangeLower := 0
	rangeUpper := 1
	randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	fmt.Println(randomNum)

	for i := 0; i <= 5; i++ {
		if randomNum == 0 {
			//go to best effort function
		} else {
			//go to reservation function
		}
	}
}

func main() {

}
