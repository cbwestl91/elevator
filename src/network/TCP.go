package network

// In this part, communication regarding orders and availability is done over TCP

import(
	"fmt"
	"net"
	"time"
	"strings"
)

var(
	commChan chan string
	

func sendTCP(){


}

func receiveTCP(){


}

func listenTCP(){
	destination := ":" + TCPport
	addr, err := net.ResolveTCPAddr("tcp", destination)
	if err != nil {
		fmt.Println("error resolving TCP adress")
	} else {
		listener, err := net.ListenTCP("tcp", addr)
		fmt.Println("listening for new TCP connections")
		if err != nil {
			fmt.Println("error listening for TCP connections")
		} else {
			socket, err := listener.Accept()
			if err != nil {
				fmt.Println("error accepting TCP connection")
			} else {
				remoteElev := socket.RemoteAddr().String()
				
				
			
			
	
	
	

}

func connectTCP(){
	

}
