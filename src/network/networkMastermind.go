package network

// Here, functions implemented in this package are used and wrapped up for easy use by elevator packages

// well.. atleast they will be in the near future. hopefully

func channelInit() {
	updateTCPmap = make(chan TCPconnection)
	newIPchan = make(chan string)
	isDeadchan = make(chan string)
	isAlicechan = make(chan int)
	giveMeCurrentMap = make(chan bool)
	getCurrentMap = make(chan map[string]net.Conn)
	giveMeConn = make(chan string)
	getSingleConn = make(chan net.Conn)
	startNewReceivechan = make(chan TCPconnection)
}

func (commChan *commChannels) commChanInit() {
	commChan.sendToAll = make(chan message)
	commChan.sendToOne = make(chan message)
	commChan.messageReceivedchan = make(chan message)
}
