package network

import(
	"encoding/json"
	"fmt"
)

/* INCOMING:
	communicator.MessageReceivedchan
	communicator.SendToAll
	communicator.SendToOne

   OUTGOING:
	communicator.decodedMessagechan	
	NEED NEW INTERNAL CHANNEL TO ALL
	NEED NEW INTERNAL CHANNEL TO ONE
*/
func messageHandler(communicator CommChannels) { // makes right format for incoming/outgoing orders and forwards to the right channel
	for {
		select {
		case incoming := <- internal.MessageReceivedchan:
			var decoded string
			err := json.Unmarshal(incoming.Content, &decoded)
			if err != nil {
				fmt.Println("FATAL ERROR: failed decoding message from: ", incoming.IP)
			} else {
				message := DecodedMessage{incoming.IP, decoded}
				communicator.DecodedMessagechan <- message
			}
		case outgoing := <- communicator.SendToAll:
			// local elevator has something for everyone. must encode into Message and forward to sendTCP
			encoded, err := json.Marshal(outgoing.Content)
			if err != nil {
				fmt.Println("FATAL ERROR: encoding before sending to all FAILED")
			} else {
				final := encodedMessage{"null", encoded}
				internal.encodedMessageSendAll <- final
			}
		case outgoing := <- communicator.SendToOne:
			encoded, err := json.Marshal(outgoing.Content)	
			if err != nil {
				fmt.Println("FATAL ERROR: encoding before sending to one FAILED: ", outgoing.IP)
			} else {
				final := encodedMessage{outgoing.IP, encoded}
				internal.encodedMessageSendOne <- final
			}
		}
	}
}
