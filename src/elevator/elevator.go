
package elevator

type elevatorinfo struct {
	state State
	event Event
	order_slice [][]int
	last_floor int
	last_direction Direction
}

var N_FLOORS, N_BUTTONS int = 4, 3
	
func (elevinf *elevatorinfo) HandleELevator() {
	
	elevinf.state := IDLE
	elevinf.event := NO_EVENT
	
	// Order Array
	elevator.order_slice := make([][]int, N_FLOORS)
	for i := range(elevator.order_slice){
		elevator.order_slice[i] = make([]int, N_BUTTONS)
	}
	BootStatemachine()
	
	for {
		UpdateStatemachine()
		RunStatemachine(event)	
	}
	
}
