package main

import (
	"fmt"
	"net"
	"time"
)

//Created a structure
type Packet struct {
	payload byte
	type1    int
}

//ability to write data  in the channel
func BestEffort(payload byte, type1 int, ch chan<- Packet) {
	for i := 1; i <= 10; i++ {
		ch <- Packet{payload, type1} //Sending 10 best effort packets to channel;

	}

}
func Reservation(payload byte, type1 int, ch chan<- Packet) {

	for i := 1; i <= 10; i++ {
		ch <- Packet{payload, type1} //Sending 10 Reservation Packets to channel;

	}
}

func main() {
	var p1 Packet
	results := make(chan Packet)

	If p1.type1==0{
		go BestEffort(ch)
	} else if p1.type1==1{
		go Reservation(ch)
	}












}
