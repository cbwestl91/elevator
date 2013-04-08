// Function for deciding which event is taking place
package elevator

import "elevdriver"

type Event int

const (
	ORDER Event = iota
	STOP
	OBSTRUCTION
	SENSOR
	NO_EVENT
)

func SetEvent(last_direction int, state State, order_slice [][]int)(event Event){
	
	if elevdriver.GetStopButton() && state != EMERGENCY {
		event = STOP
	} else if elevdriver.GetObs() {
		event = OBSTRUCTION
	} else if DetermineDirection(last_direction, order_slice) != 2 && state != UP && state != DOWN {
		event = ORDER
	} else if elevdriver.GetFloor() != -1 {
		event = SENSOR
		last_floor = elevdriver.GetFloor()
	} else {
		event = NO_EVENT
	}
	
	return event
	
}
