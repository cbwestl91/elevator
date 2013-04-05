
// Here ElevatorHandler is created,
// ready to be called in a slave or master (hopefully)

package elevator

import (
	"elevdriver"
	"fmt"
	"time"
)

const ( 
	N_FLOORS = 3
	UP
	DOWN
	NONE
)

func ElevatorHandler()(){
	elevatorInit()
	
}

func elevatorInit()(){
	fmt.Printf("Initiating elevator...\n")
	elevdriver.Init()
}







