
package elevator

import "elevdriver"

var floor_button int
var direction_button int
		
func (elevinf *Elevatorinfo) ReceiveOrders (){
	
	for {
		floorbutton, directionbutton := elevdriver.GetButton()
	
		if elevinf.state != EMERGENCY || (elevinf.state == EMERGENCY || elevinf.event == ORDER) {
			// First column of the order slice refers to UP buttons
			for i := 1; i <= N_FLOORS - 1; i++ {
				if i == floorbutton && directionbutton == 1 {
					elevinf.external_orders[i-1][0] = 1
				}
			}
			// Second column of the order slice refers to DOWN buttons
			for i := 2; i <= N_FLOORS; i++ {
				if i == floorbutton && directionbutton == 2 {
					elevinf.external_orders[i-1][1] = 1
				}
			}
		}
	
		// Third column of the order slice refers to COMMAND buttons
		for i := 1; i <= N_FLOORS; i++ {
			if i == floorbutton && directionbutton == 0 {
				elevinf.internal_orders[i-1][2] = 1
			}
		}
		
		// Clearing the unused slice spaces
		elevinf.internal_orders[3][0] = 0
		elevinf.internal_orders[0][1] = 0
	
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

func (elevinf *Elevatorinfo) StopAtCurrentFloor()(int){
	
	var current int = elevdriver.GetFloor()
	
	if elevinf.state == UP {
		for i := 0; i < 3; i = i+2 {
			if current == 0 && elevinf.internal_orders[current][i] == 1 {
				return 1
			} else if current == 1 && elevinf.internal_orders[current][i] == 1 {
				return 1
			} else if current == 2 && elevinf.internal_orders[current][i] == 1 {
				return 1
			} else if current == 3 && elevinf.internal_orders[current][i] == 1 {
				return 1
			}
		}
		
		orders_above_current := 0
		
		for i := current+1; i < 4; i++ {
			for j := 0; j < 3; j++ {
				if elevinf.internal_orders[i][j] == 1 {
					orders_above_current++
				}
			}
		}
		
		if elevinf.internal_orders[current][1] == 1 && orders_above_current == 0 {
			return 1
		}
		
		if current == 3 && elevinf.internal_orders[3][1] == 1 {
			return 1	
		}
	} else if elevinf.state == DOWN {
		for i := 1; i < 3; i++ {
			if current == 0 && elevinf.internal_orders[0][i] == 1 {
				return -1
			} else if current == 1 && elevinf.internal_orders[1][i] == 1 {
				return -1
			} else if current == 2 && elevinf.internal_orders[2][i] == 1 {
				return -1
			} else if current == 3 && elevinf.internal_orders[3][i] == 1 {
				return -1
			}
		}
		
		if current == 0 && elevinf.internal_orders[0][0] == 1 {
			return -1
		}
		
		orders_below_current := 0
		
		for i := 0; i < current; i++ {
			for j := 0; j < 3; j++ {
				if elevinf.internal_orders[i][j] == 1 {
					orders_below_current++
				}
			}
		}
		
		if elevinf.internal_orders[current][0] == 1 && orders_below_current == 0{
			return -1
		}
	} else if elevinf.state == EMERGENCY {
		for i := 0; i < 3; i++ {
			if current == 0 && elevinf.internal_orders[0][i] == 1{
				return 2
			} else if current == 1 && elevinf.internal_orders[1][i] == 1{
				return 2
			} else if current == 2 && elevinf.internal_orders[2][i] == 1{
				return 2
			} else if current == 3 && elevinf.internal_orders[3][i] == 1{
				return 2
			}
		}
	}
	
	return 0
}

func (elevinf *Elevatorinfo) DeleteOrders(){
	if elevdriver.GetFloor() == 1{
		for i := 0; i < 3; i++ {
			elevinf.internal_orders[0][i] = 0
		}
	} else if elevdriver.GetFloor() == 2{
		for i := 0; i < 3; i++ {
			elevinf.internal_orders[1][i] = 0
		}
	} else if elevdriver.GetFloor() == 3{
		for i := 0; i < 3; i++ {
			elevinf.internal_orders[2][i] = 0
		}
	} else if elevdriver.GetFloor() == 4{
		for i := 0; i < 3; i++ {
			elevinf.internal_orders[3][i] = 0
		}
	}
}











































