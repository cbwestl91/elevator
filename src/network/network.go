package network

import(
	"fmt"
	"net"
	"time"
	"strings"
)

var(
	localIP = "129.241.187.142"
	broadcast = "235.241.187.255" //må se nærmere på adressen
	
	UDPport = "8769"
)

func ListenforMaster(broadcast, UDPport) (masterexists bool) { //returns true if master exists, false if not
	pipe := make(chan bool)

	conn := UDPconnEstablisher(broadcast, UDPport)
	
	go UDPlistener(conn, pipe)

	for{
		select{
			case <- pipe:
				masterexists := true
	
			case <-time.After(500 * time.Millisecond):
				fmt.Println("No master currently exists")
				masterexists := false
		}
		return masterexists
	}
}


func UDPconnEstablisher(broadcast, UDPport) (conn *net.UDPConn){

	destination := broadcast + ":" + UDPport
	mcAddr, err := net.ResolveUDPAddr("udp", destination)
	errorhandler(err)

	fmt.Println("Multicast adress obtained")

	conn, err = net.ListenMulticastUDP("udp", nil, mcAddr)
	errorhandler(err)

	fmt.Println("UDP broadcast found")
	
	return conn
}

func UDPlistener(conn *net.UDPConn, pipe chan int){
	v := 0
	var buf [512]byte
	for{
		n, _, _ := conn.ReadFromUDP(buf[0:])
		json.Unmarshal(buf[0:n], &v)
		pipe <- v
	}
}


func MulticastFromMaster(broadcast, UDPport) {
	destination := broadcast + ":" + UDPport
	mcAddr, err := net.ResolveUDPAddr("udp", destination)
	errorhandler(err)

	conn, err = net.ListenMulticastUDP("udp", nil, mcAddr)
	

}

func findmyIP() (){

}
func errorhandler()
