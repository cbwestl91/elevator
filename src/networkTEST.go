package main

import(
	"network"
	"time"
	"fmt"
)

func main() {
	var communicator network.CommChannels
	communicator.CommChanInit()
	network.NetworkInit(communicator)

	time.Sleep(time.Second)
	
	go receiveTESTmail(communicator)

	for {
		sendTESTmail(communicator)
		time.Sleep(400*time.Millisecond)
	}
}

func sendTESTmail(communicator network.CommChannels) {
	testvar := "Will this arrive?"
	randomstruct := network.DecodedMessage{"null", testvar}
	communicator.SendToAll <- randomstruct
}

func receiveTESTmail(communicator network.CommChannels) {
	for {
		received := <- communicator.DecodedMessagechan
		fmt.Println("received message: ", received.Content)
	}
}
