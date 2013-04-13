//-----------------------------------------------------------------------------------------//
//                                    EXTERNALORDERFUNCTIONS                               //
//-----------------------------------------------------------------------------------------//
package elevator

import "network"
import "strconv"
import "time"
// import "fmt"

func ExternalOrderSend (communicator network.CommChannels) {
	for {
		select {
		case message := <- sendToAll:
			send_struct := network.DecodedMessage{"null", message}
			communicator.SendToAll <- send_struct
		case message := <- sendToOne: // sends own cost
			communicator.SendToOne <- message
		case ip := <- gochan:
			goStruct := network.DecodedMessage{ip, "Go"}
			communicator.SendToOne <- goStruct
		case ip := <- noGochan:
			noGoStruct := network.DecodedMessage{ip, "noGo"}
			communicator.SendToOne <- noGoStruct
		}
	}
}

func ExternalOrderReceive (communicator network.CommChannels) {
	for {
		received := <- communicator.DecodedMessagechan
		intver, err := strconv.Atoi(received.Content)
		
		if err == nil { // could convert int to string, so cost was received
			cost := receivedCost{received.IP, intver}
			receivedCostchan <- cost
		} else if received.Content == "Go" {
			receivedGochan <- true
		} else if received.Content == "noGo" {
			receivedNoGochan <- true
		} else if received.Content == "up 1" || received.Content == "up 2" || received.Content == "up 3" || received.Content == "down 2" || received.Content == "down 3" || received.Content == "down 4" {
		// EXTERNAL ORDERS RECEIVED, dvs. up 1, up 2 osv..
			incExternal <- received
		} else { // EXTERNAL DELETION RECEIVED, dvs. del up 1, del up 2 osv..
			content := received.Content
			receiveDeletion <- content
		}
		time.Sleep(10*time.Millisecond)
	}
}


