
package elevator

import "elevdriver"
import "fmt"
import "time"

type Event int

const (
	ORDER Event = iota
	STOP
	OBSTRUCTION
	SENSOR
	NO_EVENT
)

func SetEvent(last_direction int, state State, order_slice [][]int)(event Event){
	
	if GetStopButton() && state != EMERGENCY {
		event = STOP
	}
	else if GetObs() {
		event = OBSTRUCTION
	}
	else if DetermineDirection(last_direction, order_slice) != 2 && state != UP && state != DOWN {
		event = ORDER
	}
	else if GetFloor() != -1 {
		event = SENSOR
		last_floor = GetFloor()
	}
	else {
		event = NO_EVENT
	}
	
	return event
	
}
