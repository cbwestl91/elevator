
package elevator

import "elevdriver"
import "fmt"
import "time"

func Initiate (int state, int event, [][]int orderarray)(){
	
	elevdriver.Init()
	
	CheckLights(state,event,orderarray)
	
	elevdriver.MotorDown()
	
	for GetFloor() == -1 {}
	
	SetFloor(GetFloor())
	
	elevdriver.MotorUp()
	elevdriver.MotorStop()
	
	int current_floor = GetFloor()
	
	fmt.Printf("Elevator initiating complete!\n")
	
}



