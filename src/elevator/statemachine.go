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

func BootStatemachine(state State, event Event, order_slice [][]int){
	
	last_floor = 0
	
	Initiate(state, event, order_slice)
	
	go CheckLights(state, event, order_slice)
	go ReceiveOrders(state, event, order_slice)
	go elevdriver.MotorHandler()
	go elevdriver.Listen()
}

func UpdateStatemachine(state State){
	
	if state == UP{
		last_direction = 1
	} else if state == DOWN {
		last_direction = 2
	}
		
	FloorIndicator()
	
}

func RunStatemachine(state State, event Event, order_slice [][]int){
	
	switch state {
		case IDLE:
			statemachineIdle(state,event,order_slice)
		case UP:
			statemachineUp(state,event,order_slice)
		case DOWN:
			statemachineDown(state,event,order_slice)
		case OPEN_DOOR:
			statemachineOpendoor(state,event,order_slice)
		case EMERGENCY:
			statemachineEmergency(state,event,order_slice)
	}
	
}

func statemachineIdle(state State, event Event, order_slice [][]int)() {
	
	switch event {
		case ORDER:
			if DetermineDirection(last_direction, order_slice) != 2 && elevdriver.GetFloor() == -1 {
				Initiate(state, event, order_slice)
				state = IDLE
				break
			}
			StartMotor(DetermineDirection(last_direction,order_slice))
			if DetermineDirection(last_direction, order_slice) == -2 {
				state = OPEN_DOOR
			} else if DetermineDirection(last_direction, order_slice) == -1 {
				state = DOWN
			} else if DetermineDirection(last_direction, order_slice) == 1 {
				if elevdriver.GetFloor() == -1 {
					Initiate(state, event, order_slice)
					state = IDLE
					break
				}
				state = UP
			}
		case STOP:
			StopButtonPushed(order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed(order_slice)
			state = EMERGENCY
		case SENSOR:
		case NO_EVENT:
	}
	
}

func statemachineUp(state State, event Event, order_slice [][]int)() {

	switch event {
		case ORDER:
		case STOP:
			StopButtonPushed(order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed(order_slice)
			state = EMERGENCY
		case SENSOR: // Destination reached || someone wants to go UP || no orders above DOWN
			FloorIndicator()
			if StopAtCurrentFloor(state, order_slice) == 1 {
				elevdriver.MotorStop()
				state = OPEN_DOOR
				DeleteOrders(order_slice)
				break
			} else if elevdriver.GetFloor() == 4 {
				elevdriver.MotorStop()
				state = IDLE
			}
		case NO_EVENT:
	}
	
}

func statemachineDown(state State, event Event, order_slice [][]int)() {

	switch event {
		case ORDER:
		case STOP:
			StopButtonPushed(order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed(order_slice)
			state = EMERGENCY
		case SENSOR:
			FloorIndicator()
			if StopAtCurrentFloor(state, order_slice) == -1 {
				elevdriver.MotorStop()
				state = OPEN_DOOR
				DeleteOrders(order_slice)
				break
			} else if elevdriver.GetFloor() == 1 {
				elevdriver.MotorStop()
				state = IDLE
			}
		case NO_EVENT:
	}
	
}

func statemachineOpendoor(state State, event Event, order_slice [][]int)() {
	
	switch event {
		case ORDER:
			elevdriver.SetDoor()
			for i := 0; i < 300; i++{
				if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == false {
					elevdriver.ClearDoor()
					state = IDLE
					break
				} else if elevdriver.GetStopButton() == true {
					StopButtonPushed(order_slice)
					state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == true {
					i = 0
				}
			}
			DeleteOrders(order_slice)
			elevdriver.ClearDoor()
			if DetermineDirection(last_direction,order_slice) == -2 {
				state = OPEN_DOOR
			} else if DetermineDirection(last_direction,order_slice) == -1 {
				state = DOWN
				elevdriver.MotorDown()
			} else if DetermineDirection(last_direction,order_slice) == 1 {
				state = UP
				elevdriver.MotorUp()
			} else if DetermineDirection(last_direction,order_slice) == 2 {
				state = IDLE
			}
		case STOP:
			StopButtonPushed(order_slice)
			state = EMERGENCY
		case OBSTRUCTION:
		case SENSOR:
			fmt.Printf("Elevator reached floor %d\n", elevdriver.GetFloor())
			elevdriver.SetDoor()
			for i := 0; i < 300; i++{
				if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == false {
					elevdriver.ClearDoor()
					state = IDLE
					break
				} else if elevdriver.GetStopButton() == true {
					StopButtonPushed(order_slice)
					state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == true {
					i = 0
				}
			}
			DeleteOrders(order_slice)
			elevdriver.ClearDoor()
			if DetermineDirection(last_direction,order_slice) == -2 {
				state = OPEN_DOOR
			} else if DetermineDirection(last_direction,order_slice) == -1 {
				state = DOWN
				elevdriver.MotorDown()
			} else if DetermineDirection(last_direction,order_slice) == 1 {
				state = UP
				elevdriver.MotorUp()
			} else if DetermineDirection(last_direction,order_slice) == 2 {
				state = IDLE
			}
		case NO_EVENT:
	}
	
}

func statemachineEmergency(state State, event Event, order_slice [][]int)() {

	switch event {
		case ORDER:
			for i := 0; i < 4; i++{
				if order_slice[i][2] == 1 {
					elevdriver.ClearStopButton()
					if DetermineDirection(last_direction, order_slice) != 2 && elevdriver.GetFloor() == -1 {
						for elevdriver.GetFloor() == -1 {
							elevdriver.MotorDown()
							if StopAtCurrentFloor(state,order_slice) == 2 {
								elevdriver.MotorStop()
								state = OPEN_DOOR
								DeleteOrders(order_slice)
								break
							}
							if elevdriver.GetStopButton() || elevdriver.GetObs() {
								StopButtonPushed(order_slice)
								state = EMERGENCY
								break
							}
						}
						break
					}
					if DetermineDirection(last_direction, order_slice) == -2 {
						state = OPEN_DOOR
					} else if DetermineDirection(last_direction, order_slice) == -1 {
						state = DOWN
						elevdriver.MotorDown()
					} else if DetermineDirection(last_direction, order_slice) == 1 {
						state = UP
						elevdriver.MotorDown()
					}
				}
			}
		case STOP:
		case OBSTRUCTION:
		case SENSOR:
		case NO_EVENT:
	}
	
}































