package main


import (
	"fmt"
	"math/rand"
	"time"
)

type packet struct {
	payload byte
	type1   int
}

func checktype() {
	rangeLower := 0
	rangeUpper := 1
	rand.Seed(time.Now().Unix())
	

	for i := 0; i <= 4; i++ {
		randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
		fmt.Println(randomNum)
		if randomNum == 0 {
			//go to best effort function
			//for interfaces 1 and 3
			go besteffort_packet(randomNum, c)
			go besteffort_packet(randomNum, c)
		} else {
			//go to reservation function
			//for interfaces 2
			go reservation_packet(randomNum, c)
		}
	}

}

func besteffort_packet(randomNum int, c chan string) {
	println("This is the best-effort Package")
	c <- "The type of the packet is 0 "
}

func reservation_packet(randomNum int, c chan string) {
	println("This is the reservation package")
	c <- "The type of the packet is 1"
}

func main() {
	c := make(chan string)
	checktype()
	fmt.Println(<-c)

}
