//-----------------------------------------------------------------------------------------//
//                                   LIGHTS                                                //
//-----------------------------------------------------------------------------------------//
package elevator

import "elevdriver"
import "fmt"
import "time"

// checks pressed buttons and set lights accordingly
func (elevinf *Elevatorinfo) CheckLights (){
	for {
		if elevinf.state != EMERGENCY || (elevinf.state == EMERGENCY && elevinf.event == ORDER) {
			for i := 0; i < 3; i++ {
				if elevinf.external_orders[i][0] == 1 {
					elevdriver.SetLight(i, 1)
				} else if elevinf.external_orders[i][0] == 0 {
					elevdriver.ClearLight(i, 1)
				}
			}
			for i := 1; i < 4; i++ {
				if elevinf.external_orders[i][0] == 1 {
					elevdriver.SetLight(i, 2)
				} else if elevinf.external_orders[i][0] == 0 {
					elevdriver.ClearLight(i, 2)
				}
			}
		}
		for i := 0; i < 4; i++ {
			if elevinf.internal_orders[i][2] == 1 {
				elevdriver.SetLight(i, 0)
			} else if elevinf.internal_orders[i][2] == 0 {
				elevdriver.ClearLight(i, 0)
			}
		}
		time.Sleep(1E7)
	}
}

// sets floorindicator light
func FloorIndicator(){
	if elevdriver.GetFloor()  > 0 { 
		elevdriver.SetFloor(elevdriver.GetFloor())
		fmt.Printf("Floor detected\n")
	}
}


