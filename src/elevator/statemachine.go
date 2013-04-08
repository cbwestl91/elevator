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
	
	go CheckLights(state, event, order_slice)
	go ReceiveOrders(state, event, order_slice)
	go elevdriver.MotorHandler
	go elevdriver.Listen
}

func UpdateStatemachine(state State){
	
	if state == UP || state == DOWN {
		last_direction = state
	} 	
		
	FloorIndicator()
	
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
				if elevdriver.GetFloor() == -1 {
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
			if StopAtCurrentFloor(state, order_slice) == 1 {
				elevdriver.MotorStop()
				state = OPEN_DOOR
				DeleteOrders(order_slice)
				break
			}
			else if elevdriver.GetFloor() == 4 {
				elevdriver.MotorStop()
				state = IDLE
			}
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
			FloorIndicator()
			if StopAtCurrentFloor(state, order_slice) == -1 {
				elevdriver.MotorStop()
				state = OPEN_DOOR
				DeleteOrders(order_slice)
				break
			}
			else if elevdriver.GetFloor() == 1 {
				elevdriver.MotorStop()
				state = IDLE
			}
		case NO_EVENT:
	}
	
}

func statemachineOpendoor(state State, event Event)() {
	
	switch event {
		case ORDER:
			elevdriver.SetDoor()
			for i = 0; i < 300; i++{
				if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == 0 {
					elevdriver.ClearDoor()
					state = IDLE
					break
				}
				else if elevdriver.GetStopButton() == 1 {
					StopButtonPushed()
					state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == 1 {
					i = 0
				}
			}
			DeleteOrders(order_slice)
			elevdriver.ClearDoor()
			if DetermineDirection(last_direction,order_slice) == -2 {
				state = OPEN_DOOR
			}
			else if DetermineDirection(last_direction,order_slice) == -1 {
				state = DOWN
				elevdriver.MotorDown()
			}
			else if DetermineDirection(last_direction,order_slice) == 1 {
				state = UP
				elevdriver.MotorUp()
			}
			else if DetermineDirection(last_direction,order_slice) == 2 {
				state = IDLE
			}
		case STOP:
			StopButtonPushed(state, event, order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
		case SENSOR:
			fmt.Printf("Elevator reached floor %d\n", GetFloor())
			elevdriver.SetDoor()
			for i = 0; i < 300; i++{
				if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == 0 {
					elevdriver.ClearDoor()
					state = IDLE
					break
				}
				else if elevdriver.GetStopButton() == 1 {
					StopButtonPushed()
					state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == 1 {
					i = 0
				}
			}
			DeleteOrders(order_slice)
			elevdriver.ClearDoor()
			if DetermineDirection(last_direction,order_slice) == -2 {
				state = OPEN_DOOR
			}
			else if DetermineDirection(last_direction,order_slice) == -1 {
				state = DOWN
				elevdriver.MotorDown()
			}
			else if DetermineDirection(last_direction,order_slice) == 1 {
				state = UP
				elevdriver.MotorUp()
			}
			else if DetermineDirection(last_direction,order_slice) == 2 {
				state = IDLE
			}
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































