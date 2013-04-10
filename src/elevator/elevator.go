
package elevator

import "elevdriver"

type Elevatorinfo struct {
	state State
	event Event
	order_slice [][]int
	last_floor int
	last_direction elevdriver.Direction
}

var N_FLOORS, N_BUTTONS int = 4, 3
	
func (elevinf *Elevatorinfo) HandleElevator() {
	
	elevinf.state = IDLE
	elevinf.event = NO_EVENT
	
	// Order Array
	elevinf.order_slice = make([][]int, N_FLOORS)
	for i := range(elevinf.order_slice){
		elevinf.order_slice[i] = make([]int, N_BUTTONS)
	}
	elevinf.BootStatemachine()
	
	for {
		elevinf.UpdateLastDirection()
		elevinf.RunStatemachine()	
	}
	
}
