
// Here the entire statemachine for the elevator will be declared

package elevator

import "elevdriver"
import "fmt"
import "time"

// Creating states
type State int
type Event int

const (
	IDLE State = iota // iota gives states int from 0 (increment)
	UP 
	DOWN 
	OPEN_DOOR 
	EMERGENCY 
)

const (
	ORDER Event = iota
	STOP
	OBSTRUCTION
	SENSOR
	NO_EVENT
)

int counter
int last_floor
int last_direction
int N_FLOORS := 4
int N_BUTTONS := 3
// Order Array
order_slice := make([][]int, 

func statemachineInit(){
	
	last_floor 
	
	state := IDLE
	event := NO_EVENT
	
	
	 
	
}


