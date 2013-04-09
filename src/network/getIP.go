package network

import(
	"strings"
	"net"
	"os"
	"fmt"
)


func getMyIP() string{
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error retrieving IP adresses")
		os.Exit(1)
	}

	for _, a := range addr {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			temp := ipnet.IP.String() // net.IsGlobalUnicast()?
			return temp
			}
		}
	}
}

/*
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
*/
