package network

import(
	"net"

	"fmt"
)

func getIP() string{
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("ERROR FINDING IP ", err)
	}
	
	for _, a := range addr {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			fmt.Println(ipnet.IP.String())
			return ipnet.IP.String()
		}
	}
	return "failure"
}
