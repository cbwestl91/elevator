package main

import(
	"network"
	"time"
)

func main() {
	var communicator network.CommChannels
	communicator.CommChanInit()
	network.NetworkInit(communicator)

	time.Sleep(time.Second)
	
	go sendTESTmail(communicator)

	for {
		time.Sleep(time.Millisecond)
	}
}

func sendTESTmail(communicator network.CommChannels) {
	for {
		time.Sleep(time.Second)
	}
}
