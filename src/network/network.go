package network

import(
	"fmt"
	"net"
)

var(
	localIP = "129.241.187.142"
	broadcast = "129.241.187.255" //må se nærmere på adressen
	
	UDPport = "8769"
)

func ListenforMaster(broadcast, UDPport) {
	destination := broadcast + ":" + UDPport
	mcAddr, err := net.ResolveUDPAddr("udp", destination)
	_, _ = net.ListenMulticastUDP("udp", nil, mcAddr)
	

}

func MulticastImAlive() {

}
