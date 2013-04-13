//-----------------------------------------------------------------------------------------//
//                                    EXTERNALORDERFUNCTIONS                               //
//-----------------------------------------------------------------------------------------//
package elevator

import "network"
import "strconv"

func ExternalOrderSend () {
	for {
		select {
		case message := <- sendToAll:
			send_struct := network.DecodedMessage{"null", message}
			communicator.SendToAll <- send_struct
		case message := <- sendToOne: // sends own cost
			communicator.SendToOne <- message
		case ip := <- gochan:
			goStruct := network.DecodeMessage{ip, "Go"}
			communicator.SendToOne <- goStruct
		case ip := <- noGochan:
			noGoStruct := network.DecodeMessage{ip, "noGo"}
			communicator.SendToOne <- noGoStruct
		}
	}
}

func ExternalOrderReceive () {
	for {
		received := <- communicator.DecodedMessagechan
		
		intver, err := strconv.AtoI(received.Content)
		
		if err == nil { // could convert int to string, so cost was received
			cost := receivedCost{received.IP, intver}
			receivedCostchan <- cost
		} elsif received.Content == "Go" {
				
		} elsif received.Content == "noGo" {
		
		} else {
		// EXTERNAL ORDERS RECEIVED, dvs. up 1, up 2 osv..
		}
		
		
		incExternal <- received
	}
}


