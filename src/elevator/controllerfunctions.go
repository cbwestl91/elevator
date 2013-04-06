
package elevator

import "elevdriver"
import "fmt"
import "time"

func Initiate (int state, int event, [][]int order_slice)(){
	
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


