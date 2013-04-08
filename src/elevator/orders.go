
package elevator

import "elevdriver"

var floor_button int
var direction_button int
		
func ReceiveOrders (state State, event Event, order_slice [][]int)(){
	
	for {
		floorbutton, directionbutton := elevdriver.GetButton()
	
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

func StopAtCurrentFloor(state State, order_slice [][]int)(int){
	
	var current int = elevdriver.GetFloor()
	
	if state == UP {
		for i := 0; i < 3; i = i+2 {
			if current == 0 && order_slice[current][i] == 1 {
				return 1
			} else if current == 1 && order_slice[current][i] == 1 {
				return 1
			} else if current == 2 && order_slice[current][i] == 1 {
				return 1
			} else if current == 3 && order_slice[current][i] == 1 {
				return 1
			}
		}
		
		orders_above_current := 0
		
		for i := current+1; i < 4; i++ {
			for j := 0; j < 3; j++ {
				if order_slice[i][j] == 1 {
					orders_above_current++
				}
			}
		}
		
		if order_slice[current][1] == 1 && orders_above_current == 0 {
			return 1
		}
		
		if current == 3 && order_slice[3][1] == 1 {
			return 1	
		}
	} else if state == DOWN {
		for i := 1; i < 3; i++ {
			if current == 0 && order_slice[0][i] == 1 {
				return -1
			} else if current == 1 && order_slice[1][i] == 1 {
				return -1
			} else if current == 2 && order_slice[2][i] == 1 {
				return -1
			} else if current == 3 && order_slice[3][i] == 1 {
				return -1
			}
		}
		
		if current == 0 && order_slice[0][0] == 1 {
			return -1
		}
		
		orders_below_current := 0
		
		for i := 0; i < current; i++ {
			for j := 0; j < 3; j++ {
				if order_slice[i][j] == 1 {
					orders_below_current++
				}
			}
		}
		
		if order_slice[current][0] == 1 && orders_below_current == 0{
			return -1
		}
	} else if state == EMERGENCY {
		for i := 0; i < 3; i++ {
			if current == 0 && order_slice[0][i] == 1{
				return 2
			} else if current == 1 && order_slice[1][i] == 1{
				return 2
			} else if current == 2 && order_slice[2][i] == 1{
				return 2
			} else if current == 3 && order_slice[3][i] == 1{
				return 2
			}
		}
	}
	
	return 0
}

func DeleteOrders(order_slice [][]int)(){
	if elevdriver.GetFloor() == 1{
		for i := 0; i < 4; i++ {
			order_slice[0][i] = 0
		}
	} else if elevdriver.GetFloor() == 2{
		for i := 0; i < 4; i++ {
			order_slice[1][i] = 0
		}
	} else if elevdriver.GetFloor() == 3{
		for i := 0; i < 4; i++ {
			order_slice[2][i] = 0
		}
	} else if elevdriver.GetFloor() == 4{
		for i := 0; i < 4; i++ {
			order_slice[3][i] = 0
		}
	}
}











































