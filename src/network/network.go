package network

import(
	"fmt"
	"net"
)


func EstablishConnection(multicastAddr, multicastPort) {
	mcAddr, err = net.ResolveUDPAddr("udp", multicastAddr:multicastPort)
	_, _ = net.ListenMulticastUDP("udp", nil, mcAddr)

}

func MulticastImAlive() {

}
