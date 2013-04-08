
package elevator

import "elevdriver"
import "fmt"
import "time"

func HandleELevator() {
	
	var state State := IDLE
	var event Event := NO_EVENT
	
	BootStatemachine()
	
	for {
		UpdateStatemachine(state)
		
		// Setting event
		
		Event event := SetEvent(last_direction, state, order_slice)
		
		RunStatemachine(state, event)	
		
	}
	
}
