package network

// In this part, remote elevators are pinged through UDP
// pings are also received, so that we may keep track of who is alive and who isn't

import(
	"fmt"
	"net"
	"time"
	"strings"
	"os"
)


func UDPconnectionHandler(remoteElev string) { //goroutine that keeps track of who is alive and who isn't
	for{
		select{
			case <- isAlivechan:
				// IMPLEMENT DEATH TIMER
			case <- time.After(toleratedLosses * sleepduration * time.Millisecond):
				// remote elevator death detected. TRANSMIT!
				isDeadchan <- remoteElev
		}
	}
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
		} else {
			fmt.Println("Imalive message sent")
		}
		time.Sleep(sleepduration * time.Millisecond)
	}
}

func listenImAlive() {
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
		
		if localIP != senderAddr.IP.String(){ // makes sure we don't pick up packets from ourselves
			fmt.Println("ImAlive message received")
			
			remoteElev := senderAddr.IP.String()
			inMap := anotherElev[remoteElev] //might require an additional input var
			
			if inMap{ // inform handler that some IP already in map is still alive, and reset death timer
				anotherElev[remoteElev] <- isAlive
			} else{ //new participant found, must add to map and give designated handler
				
				anotherElev[remoteElev] = isAlivechan
				go connectionHandler(remoteElev)
				
				anotherElev[remoteElev] <- isAlive
			}
			remoteIPshare <- remoteElev
		}
	}
}

func findmyIP() string{ // this function is weird, and should be looked at. returns ip6 -.- working on better option
	systemIPs, err := net.InterfaceAddrs()
	errorhandler(err)

	tempIPstring := make([]string, len(systemIPs))
	
	for i := range systemIPs{
		temp := systemIPs[i].String()
		ip := strings.Split(temp, "/")
		tempIPstring[i] = ip[0]
	}
	myIP := tempIPstring[2]
	return myIP
}

func errorhandler(err error){ // tidies up code. will be replaced by individualized error handling for each error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
