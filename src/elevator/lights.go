// Functions for controlling the lights on the elevatorpanel
package elevator

import "elevdriver"

// checks pressed buttons and set lights accordingly
func CheckLights(state State, event Event, order_slice [][]int)(){
	
	for {
		if state != EMERGENCY || (state == EMERGENCY && event == ORDER) {
			for i := 0; i < 3; i++ {
				if order_slice[i][0] == 1 {
					elevdriver.SetLight(i, 1)
				} else if order_slice[i][0] == 0 {
					elevdriver.ClearLight(i, 1)
				}
			}
			for i := 1; i < 4; i++ {
				if order_slice[i][0] == 1 {
					elevdriver.SetLight(i, 2)
				} else if order_slice[i][0] == 0 {
					elevdriver.ClearLight(i, 2)
				}
			}
		}
	
		for i := 0; i < 4; i++ {
			if order_slice[i][2] == 1 {
				elevdriver.SetLight(i, 0)
			} else if order_slice[i][2] == 0 {
				elevdriver.ClearLight(i, 0)
			}
		}
	}
	
}

// sets floorindicator light
func FloorIndicator(){

	if elevdriver.GetFloor()  > 0 { 
		elevdriver.SetFloor(elevdriver.GetFloor())
	}
	
}


