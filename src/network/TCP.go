package network

// In this part, communication regarding orders and availability is done over TCP

import(
	"fmt"
	"net"
	"strings"
)

func TCPconnectionHandler() {
	for {
		select {
		case newIP := <- newIPchan:
			_, exists := TCPmap[newIP]
			if !exists {
				go connectTCP(newIP)
			} else {
				fmt.Println("This IP already exists in map.. which is weird")
			}
		case deadIP := <- isDeadchan:

}

func mapOverseer() {
	TCPmap := make(map[string]net.Conn)
	for {
		select {
		case newIP := <- newIPchan:
			_, exists := TCPmap[newIP]
			if !exists {
				go connectTCP(newIP)
			} else {
				fmt.Println("This IP already exists in map.. which is weird")
			}
		case deadIP := <- isDeadchan:

}

func listenTCP(){ // listens for TCP connections 
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

func connectTCP(remoteIP string) {
	destination := remoteIP + ":" + TCPport
	addr, err := net.ResolveTCPAddr("tcp", destination)
	if err != nil {
		fmt.Println("error resolving TCP adress")
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	// As of now, this function might fight with listenTCP()

	if err != nil {
		fmt.Println("error dialing TCP")
		
	} else {
		newMapEntry := TCPconnection{conn, remoteIP}
		updateTCPmap <- newMapEntry
	}
}

func sendTCP(communicator commChannels){ // for sending information over one or all TCP connections
	for { // communication is done over channels, so function should never return
		select {
			case: input := <- communicator.sendToAll:
				giveMeCurrentMap <- true
				TCPmap := <- getCurrentMap
				if TCPmap == nil {
					fmt.Println("There are no active connections")
				} else {
					for ip := range TCPmap {
						socket := TCPmap[ip]
						socket.Write(input.content)
						fmt.Println("message successfully sent to %s", ip)
					}
				}

			case: input := <- communicator.sendToOne:
				giveMeConn <- input.IP
				socket := <- getSingleConn
				// NEED ERROR CHECK HERE ASWELL
				socket.Write(input.content)
				fmt.Println("message successfully sent to %s", input.IP)
		}
	}
}

func (conn TCPconnection) receiveTCP(communicator commChannels){ //temptxt: kalles med conn.receiveTCP(communicator)
	var msg [512]byte
	for {
		_, err := conn.socket.Read(msg[0:])
		if err != nil {
			fmt.Println("error receiving on TCP connection: %s", conn.IP)
		} else {
			newMessage := message{conn.IP, msg[0:]}
			fmt.Println("New message has been received")
			messageReceivedchan <- newMessage
		}
	}
}
