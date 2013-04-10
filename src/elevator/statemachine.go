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

func (elevinf *elevatorinfo) BootStatemachine (event Event){
	
	elevinf.last_floor = 0
	
	Initiate(event)
	
	go CheckLights()
	go ReceiveOrders()
	go SetEvent()
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

func (elevinf *elevatorinfo) RunStatemachine(){
	
	switch elevinf.state {
		case IDLE:
			statemachineIdle()
		case UP:
			statemachineUp()
		case DOWN:
			statemachineDown()
		case OPEN_DOOR:
			statemachineOpendoor()
		case EMERGENCY:
			statemachineEmergency()
	}
	
}

func (elevinf *elevatorinfo) statemachineIdle() {
	
	switch elevinf.event {
		case ORDER:
			if DetermineDirection() != 2 && elevdriver.GetFloor() == -1 {
				Initiate()
				elevinf.state = IDLE
				break
			}
			StartMotor(DetermineDirection())
			if DetermineDirection() == -2 {
				elevinf.state = OPEN_DOOR
			} else if DetermineDirection() == -1 {
				elevinf.state = DOWN
			} else if DetermineDirection() == 1 {
				if elevdriver.GetFloor() == -1 {
					Initiate()
						elevinf.state = IDLE
					break
				}
				elevinf.state = UP
			}
		case STOP:
			StopButtonPushed()
			elevinf.state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed()
			elevinf.state = EMERGENCY
		case SENSOR:
		case NO_EVENT:
	}
	
}

func (elevinf *elevatorinfo) statemachineUp() {

	switch elevinf.event {
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

func (elevinf *elevatorinfo) statemachineDown() {

	switch elevinf.event {
		case ORDER:
		case STOP:
			StopButtonPushed()
			elevinf.state = EMERGENCY
		case OBSTRUCTION:
			StopButtonPushed()
			elevinf.state = EMERGENCY
		case SENSOR:
			FloorIndicator()
			if StopAtCurrentFloor() == -1 {
				elevdriver.MotorStop()
				elevinf.state = OPEN_DOOR
				DeleteOrders()
				break
			} else if elevdriver.GetFloor() == 1 {
				elevdriver.MotorStop()
				elevinf.state = IDLE
			}
		case NO_EVENT:
	}
	
}

func (elevinf *elevatorinfo) statemachineOpendoor() {
	
	switch elevinf.event {
		case ORDER:
			elevdriver.SetDoor()
			for i := 0; i < 300; i++{
				if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == false {
					elevdriver.ClearDoor()
					elevinf.state = IDLE
					break
				} else if elevdriver.GetStopButton() == true {
					StopButtonPushed()
					elevinf.state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == true {
					i = 0
				}
			}
			DeleteOrders()
			elevdriver.ClearDoor()
			if DetermineDirection() == -2 {
				elevinf.state = OPEN_DOOR
			} else if DetermineDirection() == -1 {
				elevinf.state = DOWN
				elevdriver.MotorDown()
			} else if DetermineDirection() == 1 {
				elevinf.state = UP
				elevdriver.MotorUp()
			} else if DetermineDirection() == 2 {
				elevinf.state = IDLE
			}
		case STOP:
			StopButtonPushed()
			elevinf.state = EMERGENCY
		case OBSTRUCTION:
		case SENSOR:
			fmt.Printf("Elevator reached floor %d\n", elevdriver.GetFloor())
			elevdriver.SetDoor()
			for i := 0; i < 300; i++{
				if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == false {
					elevdriver.ClearDoor()
					elevinf.state = IDLE
					break
				} else if elevdriver.GetStopButton() == true {
					StopButtonPushed()
					elevinf.state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == true {
					i = 0
				}
			}
			DeleteOrders()
			elevdriver.ClearDoor()
			if DetermineDirection() == -2 {
				elevinf.state = OPEN_DOOR
			} else if DetermineDirection() == -1 {
				state = DOWN
				elevinf.elevdriver.MotorDown()
			} else if DetermineDirection() == 1 {
				state = UP
				elevinf.elevdriver.MotorUp()
			} else if DetermineDirection() == 2 {
				elevinf.state = IDLE
			}
		case NO_EVENT:
	}
	
}

func (elevinf *elevatorinfo) statemachineEmergency() {

	switch elevinf.event {
		case ORDER:
			for i := 0; i < 4; i++{
				if elevinf.order_slice[i][2] == 1 {
					elevdriver.ClearStopButton()
					if DetermineDirection() != 2 && elevdriver.GetFloor() == -1 {
						for elevdriver.GetFloor() == -1 {
							elevdriver.MotorDown()
							if StopAtCurrentFloor() == 2 {
								elevdriver.MotorStop()
								elevinf.state = OPEN_DOOR
								DeleteOrders()
								break
							}
							if elevdriver.GetStopButton() || elevdriver.GetObs() {
								StopButtonPushed()
								elevinf.state = EMERGENCY
								break
							}
						}
						break
					}
					if DetermineDirection() == -2 {
						elevinf.state = OPEN_DOOR
					} else if DetermineDirection() == -1 {
						elevinf.state = DOWN
						elevdriver.MotorDown()
					} else if DetermineDirection() == 1 {
						elevinf.state = UP
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

