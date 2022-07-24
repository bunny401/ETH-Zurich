package main

import (
	"fmt"
	"net"
	"time"
)

//Created a structure
type Packet struct {
	Payload byte
	Type    int
}

//ability to write data  in the channel
func BestEffort(Payload byte, Type int, ch chan<- Packet) {
	for i := 1; i <= 10; i++ {
		ch <- Packet{Payload, Type} //Sending 10 best effort packets to channel;

	}

}
func Reservation(Payload byte, Type int, ch chan<- Packet) {

	for i := 1; i <= 10; i++ {
		ch <- Packet{Payload, Type} //Sending 10 Reservation Packets to channel;

	}
}

func main() {

}
