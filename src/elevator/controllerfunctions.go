
package elevator

import "elevdriver"
import "fmt"

func (elevinf *Elevatorinfo) Initiate (){
	
	elevdriver.Init()
	
	elevinf.CheckLights()
	
	elevdriver.MotorDown()
	
	for elevdriver.GetFloor() == -1 {}
	
	elevdriver.SetFloor(elevdriver.GetFloor())
	
	elevdriver.MotorUp()
	elevdriver.MotorStop()
	
	fmt.Printf("Elevator initiation complete!\n")
	
}

// This fucntions checks how many orders are under and above, and returns a number telling where it will go
func (elevinf *Elevatorinfo) DetermineDirection ()(int){
	
	current_floor := elevdriver.GetFloor()
	orders_over := 0
	orders_under := 0 
	orders_at_current := 0
	
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if elevinf.order_slice[i][j] == 1 && i < current_floor {
				orders_under++
			} else if elevinf.order_slice[i][j] == 1 && i > current_floor {
				orders_over++
			} else if elevinf.order_slice[i][j] == 1 && i == current_floor {
				orders_at_current++
			}	
		}
	}
	
	if orders_at_current > 0 {
		return -2 //Stay at floor
	} else if (orders_under > 0 && elevinf.last_direction == 2) || (orders_under > 0 && orders_over == 0) {
		return -1 //Keep going down
	} else if (orders_over > 0 && elevinf.last_direction == 1) || (orders_over > 0 && orders_under == 0) {
		return 1 //Keep going up
	} else {
		return 2 //No orders, no direction
	}
	
	return 0

} 

func StartMotor(direction int)() {
	
	if direction == -1 {
		elevdriver.MotorDown()
	} else if direction == 1 {
		elevdriver.MotorUp()
	}
}


func (elevator *Elevatorinfo) StopButtonPushed() {
	
	elevdriver.SetStopButton()
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			elevator.order_slice[i][j] = 0
		}
	}
	elevdriver.MotorStop()

}







