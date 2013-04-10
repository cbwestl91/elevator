package network

// In this part, remote elevators are pinged through UDP
// pings are also received, so that we may keep track of who is alive and who isn't

import(
	"fmt"
	"net"
	"time"
	"os"
)


func UDPconnectionHandler(remoteElev string) { //goroutine that keeps track of who is alive and who isn't
	for{
		select{
			case <- isAlivechan:
				// IMPLEMENT DEATH TIMER
				continue
			case <- time.After(toleratedLosses * sleepduration * time.Millisecond):
				// remote elevator death detected. TRANSMIT!
				isDeadchan <- remoteElev
		}
	}
}	

func SendImAlive() {
	service := broadcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp4", service)
	errorhandler(err)

	isaliveconn, err := net.DialUDP("udp4", nil, addr)
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

func ListenImAlive() {
	service := broadcast + ":" + UDPport
	addr, err := net.ResolveUDPAddr("udp4", service)
	errorhandler(err)

	isaliveconn, err := net.ListenUDP("udp4", addr)
	errorhandler(err)

	var data [512]byte
	anotherElev := make(map[string]chan int)

	for {
		_, senderAddr, err := isaliveconn.ReadFromUDP(data[0:])
		errorhandler(err)
		
		if localIP != senderAddr.IP.String(){ // makes sure we don't pick up packets from ourselves
			fmt.Println("ImAlive message received")
			
			remoteElev := senderAddr.IP.String()
			_, inMap := anotherElev[remoteElev]
			
			if inMap{ // inform handler that some IP already in map is still alive, and reset death timer
				anotherElev[remoteElev] <- isAlive
			} else{ //new participant found, must add to map and give designated handler
				
				anotherElev[remoteElev] = isAlivechan
				go UDPconnectionHandler(remoteElev)
				
				anotherElev[remoteElev] <- isAlive
			}
			newIPchan <- remoteElev
		}
	}
}

func errorhandler(err error){ // tidies up code. will be replaced by individualized error handling for each error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
