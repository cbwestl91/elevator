package main

import(
	"network"
	"time"
)

func main() {
	var communicator network.CommChannels
	communicator.CommChanInit()
	
	time.Sleep(time.Second)
	
	go sendTESTmail(communicator)

	for {
		newMessage := <- communicator.MessageReceivedchan
		time.Sleep(time.Millisecond)
	}
}

func sendTESTmail(communicator network.CommChannels) {
	for {
		message := network.Message{content: []byte("test")}
		communicator.SendToAll <- message
		time.Sleep(time.Second)
	}
}
