package network

// In this part, communication regarding orders and availability is done over TCP

import(
	"fmt"
	"net"
	"strings"
)

func mapOverseer() {


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

func connectTCP(){
	

}








/*
func sendTCP(){


}

func receiveTCP(){


}
*/
