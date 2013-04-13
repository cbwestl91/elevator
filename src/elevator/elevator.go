//-----------------------------------------------------------------------------------------//
//                                   ELEVATOR	                                           //
//-----------------------------------------------------------------------------------------//
package elevator

import "fmt"
import "network"
// import "time"

type Direction int

const (
	NONE Direction = iota
	UP
	DOWN
)

type Elevatorinfo struct {
	state State
	event Event
	internal_orders [][]int
	external_orders [][]int
	last_floor int
	last_direction Direction
}

var N_FLOORS, N_BUTTONS int = 4, 3

func (elevinf *Elevatorinfo) HandleElevator() {

	sendToAll = make(chan string)
	sendToOne = make(chan network.DecodedMessage)
	receivedCostchan = make(chan receivedCost)
	receivedGochan = make(chan bool)
	receivedNoGochan = make(chan bool)
	receiveDeletion = make(chan string)
	gochan =  make(chan string)
	noGochan = make(chan string)
	
	var communicator network.CommChannels
	communicator.CommChanInit()
	
	elevinf.state = IDLE
	elevinf.event = NO_EVENT
	
	// Initializing order arrays
	elevinf.internal_orders = make([][]int, N_FLOORS)
	for i := range(elevinf.internal_orders){
		elevinf.internal_orders[i] = make([]int, N_BUTTONS)
	}
	elevinf.external_orders = make([][]int, N_FLOORS)
	for i := range(elevinf.internal_orders){
		elevinf.external_orders[i] = make([]int, N_BUTTONS-1)
	}
	
	network.NetworkInit(communicator)
	go elevinf.ExternalOrderMaster(communicator)
	go elevinf.ExternalOrderSlave()
	go elevinf.ExternalOrderTimer()
	go elevinf.ExternalRecvDelete()
	go ExternalOrderSend(communicator)
	go ExternalOrderReceive(communicator)
	
	elevinf.BootStatemachine()
	
	elevinf.RunStatemachine()	

}

func (elevinf *Elevatorinfo) PrintStatus() {
	// for {
		var s1, s2, s3 string
		switch elevinf.state {
			case IDLE:
				s1 = "IDLE"
			case ASCENDING:
				s1 = "ASCENDING"
			case DECENDING:
				s1 = "DECENDING"
			case OPEN_DOOR:
				s1 = "OPEN_DOOR"
			case EMERGENCY:
				s1 = "EMERGENCY"
		}
		switch elevinf.event {
			case ORDER:
				s2 = "ORDER"
			case STOP:
				s2 = "STOP"
			case OBSTRUCTION:
				s2 = "OBSTRUCTION"
			case SENSOR:
				s2 = "SENSOR"
			case NO_EVENT:
				s2 = "NO_EVENT"
		}
		switch elevinf.last_direction {
			case NONE:
				s3 = "NONE"
			case UP:
				s3 = "UP"
			case DOWN:
				s3 = "DOWN"
		}
		fmt.Printf("Elevatorstatus--> State: %s Event: %s LastFloor: %d LastDirection: %s\n",
					s1, s2, elevinf.last_floor, s3)	
		// time.Sleep(100*time.Millisecond)
	// }
}














