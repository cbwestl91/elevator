package network

// In this part, communication regarding orders and availability is done over TCP

import(
	"fmt"
	"net"
	"strings"
	"time"
)

func TCPHandler(communicator CommChannels) {
	// Manager function for TCP part. Goroutines for connects are spawned here
	
	for {
		select {
		case newIP := <- internal.newIPchan: // new UDP source detected, will connect
			go connectTCP(newIP)
		case conn := <- internal.startNewReceivechan:
			go conn.receiveTCP(communicator)
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
		case newMapEntry := <- internal.updateTCPmap: // new entry detected
			_, exists := TCPmap[newMapEntry.IP]
			if !exists {
				TCPmap[newMapEntry.IP] = newMapEntry.socket
				internal.startNewReceivechan <- newMapEntry
			} else {
				fmt.Println("IP already exists in TCPmap. Won't do anything")
			}
		case deadIP := <- internal.isDeadchan: // someone stopped transmitting UDP, and needs to be removed from map
			internal.closeConn <- deadIP
			delete(TCPmap, deadIP)
			// NEED TO STOP RECEIVING ON CONNECTION WITH THIS IP

		case <- internal.giveMeCurrentMap: // send function wants full map
			internal.getCurrentMap <- TCPmap
		case wantedIP := <- internal.giveMeConn: // if only one connection is wanted
			internal.getSingleConn <- TCPmap[wantedIP]
		}
	}
}

// OUTPUTS: newMapEntry over UPDATETCPMAP
func listenTCP(){ // listens for TCP connections 
	service := ":" + TCPport
	addr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println("net.ResolveTCPAddr error in listenTCP")
		internal.setupFail <- true
	} else {
		listener, err := net.ListenTCP("tcp4", addr)
		if err != nil {
			fmt.Println("error listening for TCP connections: ", err)
			internal.setupFail <- true
		} else {
			for {
				socket, err := listener.Accept()
				if err != nil {
					fmt.Println("listener.Accept() error in listenTCP, FIX ME PLS!")
				} else {
					remoteElevAddr := socket.RemoteAddr().String()

					remoteElevSplitter := strings.Split(remoteElevAddr, ":")
					remoteElevIP := remoteElevSplitter[0]
					newMapEntry := TCPconnection{socket, remoteElevIP}

					// found new peer. will forward info about peer
					fmt.Println("successfully connected to: ", remoteElevIP)
					internal.updateTCPmap <- newMapEntry
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
		fmt.Println("net.ResolveTCPAddr error in connectTCP for: ", remoteIP)
	
	} else {
		conn, err := net.Dial("tcp4", service)
		// As of now, this function might fight with listenTCP()

		if err != nil {
			fmt.Println("error dialing TCP")
		} else {
			newMapEntry := TCPconnection{conn, remoteIP}
			internal.updateTCPmap <- newMapEntry
			fmt.Println("successfully connected to: ", remoteIP)
		}
	}
}

// OUTPUTS: true over GIVEMECURRENTMAP
// 	    input.IP over GIVEMECONN
// INPUTS:  input struct from COMMUNICATOR.SENDTOALL
//	    TCPmap from GETCURRENTMAP
//	    input struct from COMMUNICATOR.SENDTOONE
//	    socket from GETSINGLECONN
func sendTCP(communicator CommChannels){ 
	for { // communication is done over channels, so function should never return
		select {
			case message := <- communicator.SendToAll:
				fmt.Println("Sending message to all")
				internal.giveMeCurrentMap <- true
				TCPmap := <- internal.getCurrentMap
				if TCPmap == nil {
					fmt.Println("Unable to send to all: there are no active connections")
				} else {
					for ip := range TCPmap {
						socket := TCPmap[ip]
						socket.SetWriteDeadline(time.Now().Add(300*time.Millisecond))
						_, err := socket.Write(message.content)
						if err != nil {
							fmt.Println("error sending on all TCP conns: ", err)
						} else {
							fmt.Println("message successfully sent to %s", ip)
						}
					}
				}

			case message := <- communicator.SendToOne:
				fmt.Println("Sending message to one")
				internal.giveMeConn <- message.IP
				socket := <- internal.getSingleConn
				// NEED ERROR CHECK HERE ASWELL
				socket.Write(message.content)
				fmt.Println("message successfully sent to %s", message.IP)
		}
	}
}

// OUTPUTS: messages received over MESSAGERECEIVEDCHAN
func (conn TCPconnection) receiveTCP(communicator CommChannels){
	var msg [512]byte
	for {
		select {
		case deadIP := <- internal.closeConn:
			if deadIP == conn.IP {
				break
			}
		default:
			n, err := conn.socket.Read(msg[0:])
			if err != nil {
				fmt.Println("Read error in receiveTCP: ", conn.IP, err)
			} else {
				newMessage := Message{conn.IP, msg[0:n]}
				fmt.Println("New message has been received")
				communicator.MessageReceivedchan <- newMessage
			}
		}
	}
}
