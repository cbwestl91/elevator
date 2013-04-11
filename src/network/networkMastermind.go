package network

import(
	"net"
	"fmt"
)

// Here, functions implemented in this package are used and wrapped up for easy use by elevator and main packages

// well.. atleast they will be in the near future. hopefully

func Channelinit() {
	updateTCPmap = make(chan TCPconnection)
	newIPchan = make(chan string)
	isDeadchan = make(chan string)
	isAlivechan = make(chan int)
	giveMeCurrentMap = make(chan bool)
	getCurrentMap = make(chan map[string]net.Conn)
	giveMeConn = make(chan string)
	getSingleConn = make(chan net.Conn)
	startNewReceivechan = make(chan TCPconnection)
	fmt.Println("channels initialized")
}

func (commChan *commChannels) commChaninit() {
	commChan.sendToAll = make(chan message)
	commChan.sendToOne = make(chan message)
	commChan.messageReceivedchan = make(chan message)
	fmt.Println("communicator initialized")
}

func UDPinit() {
	go sendImAlive()
	go listenImAlive()
}

func TCPinit() {
	communicator := commChannels{}
	communicator.commChaninit()
	go TCPconnectionHandler(communicator)
}

