// Here the entire statemachine for the elevator will be declared
package elevator

import "elevdriver"
import "fmt"
import "time"

// Creating states
type State int

const (
	IDLE State = iota // iota gives states int from 0 (increment)
	UP 
	DOWN 
	OPEN_DOOR 
	EMERGENCY 
)

var counter, last_floor, last_direction int
var N_FLOORS, N_BUTTONS int = 4, 3

// Order Array
order_slice := make([][]int, N_FLOORS)
for i := range(order_slice){
	order_slice[i] = make([]int, N_BUTTONS)
}

func BootStatemachine(state State, ){
	
	last_floor = 0
	
	Initiate(state, event, order_slice)
	
	go ReceiveOrders(state, event, order_slice)
	go elevdriver.MotorHandler
	go elevdriver.Listen
}

func UpdateStatemachine(state State){
	
	if state == UP || state == DOWN {
		last_direction = state
	} 	
		
	FloorIndicator()
		
	CheckLights()
	
}

func RunStatemachine(state State, event Event){
	
	switch state {
		case IDLE:
			statemachineIdle(state,event)
		case UP:
			statemachineUp(state,event)
		case DOWN:
			statemachineDown(state,event)
		case OPEN_DOOR:
			statemachineOpendoor(state,event)
		case EMERGENCY:
			statemachineEmergency(state,event)
	}
	
}

func statemachineIdle(state State, event Event)() {
	
	switch event {
		case ORDER:
			if DetermineDirection(last_direction, order_slice) != 2 && GetFloor() == -1 {
				Initiate(state, event, order_slice)
				state = IDLE
				break
			}
			StartMotor(DetermineDirection(last_direction,order_slice))
			if DetermineDirection(last_direction, order_slice) == -2 {
				state = OPEN_DOOR
			}
			else if DetermineDirection(last_direction, order_slice) == -1 {
				state = DOWN
			}
			else if DetermineDirection(last_direction, order_slice) == 1 {
				if GetFloor() == -1 {
					Initiate(state, event, order_slice)
					state = IDLE
					break
				}
				state = UP
			}
		case STOP:
			StopButtonPushed(state, event, order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed(state, event, order_slice)
			state = EMERGENCY
		case SENSOR:
		case NO_EVENT:
	}
	
}

func statemachineUp(state State, event Event)() {

	switch event {
		case ORDER:
		case STOP:
			StopButtonPushed(state, event, order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed(state, event, order_slice)
			state = EMERGENCY
		case SENSOR: // Destination reached || someone wants to go UP || no orders above DOWN
			FloorIndicator()
			if(
		case NO_EVENT:
	}
	
}

func statemachineDown(state State, event Event)() {

	switch event {
		case ORDER:
		case STOP:
			StopButtonPushed(state, event, order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed(state, event, order_slice)
			state = EMERGENCY
		case SENSOR:
		case NO_EVENT:
	}
	
}

func statemachineOpendoor(state State, event Event)() {
	
	switch event {
		case ORDER:
		case STOP:
		case OBSTRUCTION:
		case SENSOR:
		case NO_EVENT:
	}
	
}

func statemachineEmergency(state State, event Event)() {

	switch event {
		case ORDER:
		case STOP:
		case OBSTRUCTION:
		case SENSOR:
		case NO_EVENT:
	}
	
}






























