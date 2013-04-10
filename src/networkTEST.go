package main

import(
	"network"
	"fmt"
)

func main() {
	network.Channelinit()
	//network.UDPinit()
	
	go network.SendImAlive()
	go network.ListenImAlive()
	
	
	network.TCPinit()
	
	fmt.Println("SPAM")
	for {
	}
}
