
package elevator

import "elevdriver"
import "fmt"
import "time"

var floor_button int
var direction_button int
		
func ReceiveOrders (state State, event Event, order_slice [][]int)(){
	
	for {
		floorbutton, directionbutton := GetButton()
	
		if state != EMERGENCY || (state == EMERGENCY || event == ORDER) {
			// First column of the order slice refers to UP buttons
			for i := 1; i <= N_FLOORS - 1; i++ {
				if i == floorbutton && directionbutton == 1 {
					order_slice[i-1][0] = 1
				}
			}
			// Second column of the order slice refers to DOWN buttons
			for i := 2; i <= N_FLOORS; i++ {
				if i == floorbutton && directionbutton == 2 {
					order_slice[i-1][1] = 1
				}
			}
		}
	
		// Third column of the order slice refers to COMMAND buttons
		for i := 1; i <= N_FLOORS; i++ {
			if i == floorbutton && directionbutton == 3 {
				order_slice[i-1][2] = 1
			}
		}
		
		// Clearing the unused slice spaces
		order_slice[3][0] = 0
		order_slice[0][1] = 0
	
		floorbutton = 0
		directionbutton = 0
	}
	
}
/*
00 up1 01 emp 02 co1
10 up2 11 do2 12 co2
20 up3 21 do3 22 co3
30 emp 31 do4 32 co4
*/

func StopCurrentFloor

	current int = GetFloor()
	
