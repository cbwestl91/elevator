
package elevator

func HandleELevator() {
	
	state := IDLE
	event := NO_EVENT
	
	// Order Array
	order_slice := make([][]int, N_FLOORS)
	for i := range(order_slice){
		order_slice[i] = make([]int, N_BUTTONS)
	}
	BootStatemachine(state, order_slice)
	
	for {
		UpdateStatemachine(state)
		event = SetEvent(last_direction, state, order_slice)
		RunStatemachine(state, event, order_slice)	
	}
	
}
