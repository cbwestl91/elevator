package network

// In this part, communication regarding orders and availability is done over TCP

import(
	"fmt"
	"net"
	"strings"
	"time"
)

func TCPconnectionHandler(communicator commChannels) {
	// Umbrella function for TCP part. Goroutines started here
	go mapOverseer()
	go listenTCP()
	go sendTCP(communicator)	
	
	for {
		select {
		case newIP := <- newIPchan: // new UDP source detected, which means we need a new TCP connection			
			go connectTCP(newIP)
		case conn := <- startNewReceivechan: //
			conn.receiveTCP(communicator)
		}
	}
}

// OUTPUTS: TCPmap over GETCURRENTMAP
//	    TCPmap[wantedIP] over GETSINGLECONN
// INPUTS:  newMapEntry from UPDATETCPMAP
//	    deadIP from ISDEADCHAN
//	    bool from GIVEMECURRENTMAP
//	    wantedIP from GIVEMECONN
func mapOverseer() {
	TCPmap := make(map[string]net.Conn)
	for {
		select {
		case newMapEntry := <- updateTCPmap: // new entry detected
			_, exists := TCPmap[newMapEntry.IP]
			if !exists {
				TCPmap[newMapEntry.IP] = newMapEntry.socket
				startNewReceivechan <- newMapEntry
			}
		case deadIP := <- isDeadchan: // someone stopped transmitting UDP, and needs to be removed from map
			delete(TCPmap, deadIP)
			// NEED TO STOP RECEIVING ON CONNECTION WITH THIS IP

		case <- giveMeCurrentMap: // send function wants full map
			getCurrentMap <- TCPmap
		case wantedIP := <- giveMeConn: // if only one connection is wanted
			getSingleConn <- TCPmap[wantedIP]
		}
	}
}

// OUTPUTS: newMapEntry over UPDATETCPMAP
func listenTCP(){ // listens for TCP connections 
	service := ":" + TCPport
	addr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println("error resolving TCP adress")
	} else {
		listener, err := net.ListenTCP("tcp4", addr)
		fmt.Println("listening for new TCP connections")
		if err != nil {
			fmt.Println("error listening for TCP connections")
		} else {
			for {
				socket, err := listener.Accept()
				if err != nil {
					fmt.Println("error accepting TCP connection")
				} else {
					remoteElevAddr := socket.RemoteAddr().String()
			
					remoteElevSplitter := strings.Split(remoteElevAddr, ":")
					remoteElevIP := remoteElevSplitter[0]
					newMapEntry := TCPconnection{socket, remoteElevIP}

					// found new peer. will forward info about peer
					updateTCPmap <- newMapEntry
				} // what happens if several identical copies are made? overwrite?
			}
		}
	}
}

// OUTPUTS: newMapEntry over UPDATETCPMAP
func connectTCP(remoteIP string) {
	service := remoteIP + ":" + TCPport
	_, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println("error resolving TCP adress")
	} else {
		conn, err := net.Dial("tcp4", service)
		// As of now, this function might fight with listenTCP()

		if err != nil {
			fmt.Println("error dialing TCP")
		} else {
			newMapEntry := TCPconnection{conn, remoteIP}
			updateTCPmap <- newMapEntry
		}
	}
}

// OUTPUTS: true over GIVEMECURRENTMAP
// 	    input.IP over GIVEMECONN
// INPUTS:  input struct from COMMUNICATOR.SENDTOALL
//	    TCPmap from GETCURRENTMAP
//	    input struct from COMMUNICATOR.SENDTOONE
//	    socket from GETSINGLECONN
func sendTCP(communicator commChannels){ 
	for { // communication is done over channels, so function should never return
		select {
			case message := <- communicator.sendToAll:
				fmt.Println("Sending message to all")
				giveMeCurrentMap <- true
				TCPmap := <- getCurrentMap
				if TCPmap == nil {
					fmt.Println("There are no active connections")
				} else {
					for ip := range TCPmap {
						socket := TCPmap[ip]
						socket.SetWriteDeadline(time.Now().Add(300*time.Millisecond))
						socket.Write(message.content)
						fmt.Println("message successfully sent to %s", ip)
					}
				}

			case message := <- communicator.sendToOne:
				fmt.Println("Sending message to one")
				giveMeConn <- message.IP
				socket := <- getSingleConn
				// NEED ERROR CHECK HERE ASWELL
				socket.Write(message.content)
				fmt.Println("message successfully sent to %s", message.IP)
		}
	}
}

// OUTPUTS: messages received over MESSAGERECEIVEDCHAN
func (conn TCPconnection) receiveTCP(communicator commChannels){
	var msg [512]byte
	for {
		n, err := conn.socket.Read(msg[0:])
		if err != nil {
			fmt.Println("error receiving on TCP connection: %s", conn.IP)
			return
		} else {
			newMessage := message{conn.IP, msg[0:n]}
			fmt.Println("New message has been received")
			communicator.messageReceivedchan <- newMessage
		}
	}
}
