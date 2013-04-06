package network

import(
	"fmt"
	"net"
	"time"
	"strings"
)

const sleepduration = 1000 //interval between alivemessages given in milliseconds

var(
	localIP = "129.241.187.142"
	broadcast = "235.241.187.255" //må se nærmere på adressen
	
	UDPport = "8769"
	TCPport = " 8770"
)


func sendImAlive() {
	destination := broadcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp", destination)
	errorhandler(err)

	isaliveconn, err := net.DialUDP("udp", nil, addr)
	errorhandler(err)
	
	isaliveMessage := []byte("1")
	for {
		_, err := isaliveconn.Write(isaliveMessage)
		if err != nil {
			fmt.Println("Error sending Imalive message")
		} 
		else {
			fmt.Println("Imalive message sent")
		}
		time.Sleep(sleepduration * time.Millisecond)
	}
}

func listenImAlive() [
	fmt.Println(localIP)

	destination := broadcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp", destination)
	errorhandler(err)

	isaliveconn, err := net.ListenUDP("udp", addr)
	errorhandler(err)

	var data [512]byte
	for {
		_, _, err := isaliveconn.ReadFromUDP(data[0:])



func aliveCounter() {


}

func findmyIP() string{
	systemIPs, err := net.InterfaceAddrs()
	errorhandler(err)

	IPstring := make([]string, len(systemIPs))
	
	for i := range systemIPs{
		temp := systemIPs[i].String()
		ip := strings.Split(temp, "/")
		tempIPstring[i] = ip[0]
	}
	myIP := tempIPstring[2]
	return myIP
}

func errorhandler(err error){
	if err != nil {
		fmt.Println("Fatal error in network package")
	}
}

/*
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
*/
