package network

import(
	"net"
	"fmt"
	"time"
)

// Here, functions implemented in this package are used and wrapped up for easy use by elevator and main packages

func NetworkInit(communicator CommChannels) {	
	internalInit()
	go UDPHandler(communicator)
	go listenImAlive()
	go sendImAlive()

	time.Sleep(time.Millisecond)
	
	go TCPHandler(communicator)
	go mapOverseer()
	go listenTCP()
	go sendTCP(communicator)

	time.Sleep(time.Second)

	for {
		select {
		case <- internal.setupFail:
			internal.quitsendImAlive <- true
			internal.quitlistenImAlive <- true
			
			time.Sleep(time.Millisecond)
			
			go listenImAlive()
			go sendImAlive()
			go listenTCP()
		case <- time.After(400 * time.Millisecond):
			fmt.Println("network properly initialized")
			return
		}
	}	
}


func internalInit() {
	internal.updateTCPmap = make(chan TCPconnection)
	internal.newIPchan = make(chan string)
	internal.isDeadchan = make(chan string)
	internal.isAlivechan = make(chan string)
	internal.giveMeCurrentMap = make(chan bool)
	internal.getCurrentMap = make(chan map[string]net.Conn)
	internal.giveMeConn = make(chan string)
	internal.getSingleConn = make(chan net.Conn)
	internal.startNewReceivechan = make(chan TCPconnection)
	internal.closeConn = make(chan string)
	internal.quitsendImAlive = make(chan bool)
	internal.quitlistenImAlive = make(chan bool)
	internal.setupFail = make(chan bool)
	fmt.Println("internal channels initialized")
}

func (communicator *CommChannels) CommChanInit() {
	communicator.SendToAll = make(chan Message)
	communicator.SendToOne = make(chan Message)
	communicator.MessageReceivedchan = make(chan Message)
	communicator.getDeadIPchan = make(chan string)
	communicator.sendDeadIPchan = make(chan string)
	fmt.Println("communicator initialized")
}
