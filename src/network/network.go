package network

import(
	"fmt"
	"net"
)

var(
	localIP = "129.241.187.142"
	broadcast = "235.241.187.255" //må se nærmere på adressen
	
	UDPport = "8769"
)

func ListenforMaster(broadcast, UDPport) (masterexists bool) { //returns true if master exists
	destination := broadcast + ":" + UDPport
	mcAddr, err := net.ResolveUDPAddr("udp", destination)
	errorhandler(err)

	conn, err = net.ListenMulticastUDP("udp", nil, mcAddr)
	errorhandler(err)

	var buf [512]byte
	_, _, err = conn.ReadFromUDP(buf[0:])
	errorhandler(err)

	if buf

}

func MulticastFromMaster(broadcast, UDPport) {
	

}

func errorhandler()
