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

func (elevinf *Elevatorinfo) SetEvent(){
	
	if elevdriver.GetStopButton() && elevinf.state != EMERGENCY {
		elevinf.event = STOP
	} else if elevdriver.GetObs() {
		elevinf.event = OBSTRUCTION
	} else if elevinf.DetermineDirection() != 2 && elevinf.state != UP && elevinf.state != DOWN {
		elevinf.event = ORDER
	} else if elevdriver.GetFloor() != -1 {
		elevinf.event = SENSOR
		elevinf.last_floor = elevdriver.GetFloor()
	} else {
		elevinf.event = NO_EVENT
	}
	
}
