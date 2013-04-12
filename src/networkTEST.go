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

	for {
		sendTESTmail(communicator)
		time.Sleep(time.Second)
		receiveTESTmail(communicator)
		time.Sleep(time.Second)
	}
}

func sendTESTmail(communicator network.CommChannels) {
	testvar := "Will this arrive?"
	randomstruct := network.DecodedMessage{"null", testvar}
	communicator.SendToAll <- randomstruct
}

func receiveTESTmail(communicator network.CommChannels) {
	received := <- communicator.DecodedMessagechan
	fmt.Println("received message: ", received.Content)
}
