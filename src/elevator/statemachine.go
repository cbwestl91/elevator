// Here the entire statemachine for the elevator will be declared
package elevator

import "elevdriver"
import "fmt"
import "time"

// Creating states
type State int

const (
	IDLE State = iota // iota gives states int from 0 (increment)
	UP 
	DOWN 
	OPEN_DOOR 
	EMERGENCY 
)

var counter, last_floor, last_direction int
var N_FLOORS, N_BUTTONS int = 4, 3

// Order Array
order_slice := make([][]int, N_FLOORS)
for i := range(order_slice){
	order_slice[i] = make([]int, N_BUTTONS)
}

func BootStatemachine(state State, ){
	
	last_floor = 0
	
	Initiate(state, event, order_slice)
	
	go ReceiveOrders(state, event, order_slice)
	
}

func UpdateStatemachine(state State){
	
	if state == UP || state == DOWN {
		last_direction = state
	} 	
		
	FloorIndicator()
		
	CheckLights()
	
}

func RunStatemachine(state State, event Event){
	
	switch state {
		case IDLE:
			statemachineIdle(event)
		case UP:
			statemachineUp(event)
		case DOWN:
			statemachineDown(event)
		case OPEN_DOOR:
			statemachineOpendoor(event)
		case EMERGENCY:
			statemachineEmergency(event)
	}
	
}

func statemachineIdle(event Event)() {

}

func statemachineUp(event Event)() {

}

func statemachineDown(event Event)() {

}

func statemachineOpendoor(event Event)() {

}

func statemachineEmergency(event Event)() {

}






























