package network

import(
	"fmt"
	"net"
	"time"
	"strings"
)

const(
	sleepduration = 1000 //interval between alivemessages given in milliseconds
	isAlive = 1
	dead = 0
)

var(
	localIP = "129.241.187.142"
	broadcast = "235.241.187.255" //må se nærmere på adressen
	
	UDPport = "8769"
	TCPport = " 8770"

)

var(
	isDeadchan chan int
	isAlivechan chan int
)

func connectionHandler(remoteElev string) { //goroutine that keeps track of who is alive and who isn't


}


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
	anotherElev := make(map[string]chan int)

	for {
		_, senderAddr, err := isaliveconn.ReadFromUDP(data[0:])
		errorhandler(err)
		
		if localIP != senderAddr.IP.String(){
			fmt.Println("ImAlive message received")
			
			remoteElev := senderAddr.IP.String()
			inMap := anotherElev[remoteElev] //might require an additional input var
			
			if inMap{ // inform handler that some IP already in map is still alive, and reset death timer
				anotherElev[remoteElev] <- isAlive
			}
			else{ //new participant found, must add to map and give designated handler
				anotherElev[remoteElev] = isAlivechan
				go connectionHandler(remoteElev)
				
		





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
	break
	}
}
