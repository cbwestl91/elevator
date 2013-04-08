
package elevator

import "elevdriver"
import "fmt"
import "time"

func Initiate (state State, event Event, order_slice [][]int)(){
	
	elevdriver.Init()
	
	CheckLights(state,event,order_slice)
	
	elevdriver.MotorDown()
	
	for GetFloor() == -1 {}
	
	SetFloor(GetFloor())
	
	elevdriver.MotorUp()
	elevdriver.MotorStop()
	
	int current_floor = GetFloor()
	
	fmt.Printf("Elevator initiation complete!\n")
	
}

// This fucntions checks how many orders are under and above, and returns a number telling where it will go
func DetermineDirection(last_direction int, order_slice [][]int)(int){
	
	int current_floor = GetFloor()
	int orders_over := 0
	int orders_under := 0 
	int orders_at_current := 0
	
	for i = 0; i < 4; i++ {
		for j = 0; j < 3; j++ {
			if order_slice[i][j] == 1 && i < current {
				orders_under++
			}
			else if order_slice[i][j] == 1 && i > current {
				orders_over++
			}
			else if order_slice[i][j] == 1 && i = current {
				orders_at_current++
			}	
		}
	}
	
	if orders_at_current > 0 {
		return -2 //Stay at floor
	}
	else if (orders_under > 0 && last_direction == 2) || (orders_under > 0 && orders_over == 0) {
		return -1 //Keep going down
	}
	else if (orders_over > 0 && last_direction == 1) || (orders_over > 0 && orders_under == 0) {
		return 1 //Keep going up
	}
	else {
		return 2 //No orders, no direction
	}

} 

func StartMotor(direction int)() {
	
	if direction == -1 {
		elevdriver.MotorDown()
	}
	else if direction == 1 {
		elevdriver.MotorUp()
	}
}


func StopButtonPushed(state State, event Event, order_slice [][]int)() {
	
	elevdriver.SetStopButton()
	for i = 0; i < 4; i++ {
		for j = 0; j < 3; j++ {
			order_slice[i][j] = 0
		}
	}
	CheckLights(state, event, order_slice)
	elevdriver.MotorStop()

}







