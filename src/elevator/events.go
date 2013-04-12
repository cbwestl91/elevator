//-----------------------------------------------------------------------------------------//
//                                   EVENTS                                                //
//-----------------------------------------------------------------------------------------//
package elevator

import "elevdriver"
import "time"

type Event int

const (
	ORDER Event = iota
	STOP
	OBSTRUCTION
	SENSOR
	NO_EVENT
)

func (elevinf *Elevatorinfo) SetEvent(){
	for{
		if elevdriver.GetStopButton() && elevinf.state != EMERGENCY {
			elevinf.event = STOP
		} else if elevdriver.GetObs() {
			elevinf.event = OBSTRUCTION
		} else if elevinf.DetermineDirection() != 2 && elevinf.state != ASCENDING && elevinf.state != DECENDING {
			elevinf.event = ORDER
		} else if elevdriver.GetFloor() != -1 {
			elevinf.event = SENSOR
			elevinf.last_floor = elevdriver.GetFloor()
		} else {
			elevinf.event = NO_EVENT
		}
		time.Sleep(1E7)
	}
}
