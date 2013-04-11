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

func (elevinf *Elevatorinfo) BootStatemachine (){
	fmt.Printf("STATEMACHINE BOOTING...\n")
	elevinf.l  ast_floor = 0
	
	elevinf.Initiate()
	
	fmt.Printf("Balle\n")
	go elevinf.CheckLights()
	go elevinf.ReceiveOrders()
	go elevinf.SetEvent()
	
	
	fmt.Printf("STATEMACHINE BOOTED\n")
}

func (elevinf *Elevatorinfo) UpdateLastDirection(){
	
	if elevinf.state == UP{
		elevinf.last_direction = 1
	} else if elevinf.state == DOWN {
		elevinf.last_direction = 2
	}
		
	FloorIndicator()
}

func (elevinf *Elevatorinfo) RunStatemachine(){
	
	switch elevinf.state {
		case IDLE:
			elevinf.statemachineIdle()
		case UP:
			elevinf.statemachineUp()
		case DOWN:
			elevinf.statemachineDown()
		case OPEN_DOOR:
			elevinf.statemachineOpendoor()
		case EMERGENCY:
			elevinf.statemachineEmergency()
	}
	
}

func (elevinf *Elevatorinfo) statemachineIdle() {
	
	switch elevinf.event {
		case ORDER:
			if elevinf.DetermineDirection() != 2 && elevdriver.GetFloor() == -1 {
				elevinf.Initiate()
				elevinf.state = IDLE
				break
			}
			StartMotor(elevinf.DetermineDirection())
			if elevinf.DetermineDirection() == -2 {
				elevinf.state = OPEN_DOOR
			} else if elevinf.DetermineDirection() == -1 {
				elevinf.state = DOWN
			} else if elevinf.DetermineDirection() == 1 {
				if elevdriver.GetFloor() == -1 {
					elevinf.Initiate()
						elevinf.state = IDLE
					break
				}
				elevinf.state = UP
			}
		case STOP:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case OBSTRUCTION:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case SENSOR:
		case NO_EVENT:
	}
	
}

func (elevinf *Elevatorinfo) statemachineUp() {

	switch elevinf.event {
		case ORDER:
		case STOP:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case OBSTRUCTION:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case SENSOR: // Destination reached || someone wants to go UP || no orders above DOWN
			FloorIndicator()
			if elevinf.StopAtCurrentFloor() == 1 {
				elevdriver.MotorStop()
				elevinf.state = OPEN_DOOR
				elevinf.DeleteOrders()
				break
			} else if elevdriver.GetFloor() == 4 {
				elevdriver.MotorStop()
				elevinf.state = IDLE
			}
		case NO_EVENT:
	}
	
}

func (elevinf *Elevatorinfo) statemachineDown() {

	switch elevinf.event {
		case ORDER:
		case STOP:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case OBSTRUCTION:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case SENSOR:
			FloorIndicator()
			if elevinf.StopAtCurrentFloor() == -1 {
				elevdriver.MotorStop()
				elevinf.state = OPEN_DOOR
				elevinf.DeleteOrders()
				break
			} else if elevdriver.GetFloor() == 1 {
				elevdriver.MotorStop()
				elevinf.state = IDLE
			}
		case NO_EVENT:
	}
	
}

func (elevinf *Elevatorinfo) statemachineOpendoor() {
	
	switch elevinf.event {
		case ORDER:
			fmt.Printf("The door is open\n")
			elevdriver.SetDoor()
			for i := 0; i < 300; i++{
				if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == false {
					elevdriver.ClearDoor()
					elevinf.state = IDLE
					break
				} else if elevdriver.GetStopButton() == true {
					elevinf.StopButtonPushed()
					elevinf.state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == true {
					fmt.Printf("Obstruction detected, door staying open\n")
					i = 0
				}
			}
			elevinf.DeleteOrders()
			elevdriver.ClearDoor()
			if elevinf.DetermineDirection() == -2 {
				elevinf.state = OPEN_DOOR
			} else if elevinf.DetermineDirection() == -1 {
				elevinf.state = DOWN
				elevdriver.MotorDown()
			} else if elevinf.DetermineDirection() == 1 {
				elevinf.state = UP
				elevdriver.MotorUp()
			} else if elevinf.DetermineDirection() == 2 {
				elevinf.state = IDLE
			}
		case STOP:
			elevinf.StopButtonPushed()
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
					elevinf.StopButtonPushed()
					elevinf.state = EMERGENCY
					break
				}
				time.Sleep(10*time.Millisecond)
				if elevdriver.GetObs() == true {
					i = 0
				}
			}
			elevinf.DeleteOrders()
			elevdriver.ClearDoor()
			if elevinf.DetermineDirection() == -2 {
				elevinf.state = OPEN_DOOR
			} else if elevinf.DetermineDirection() == -1 {
				elevinf.state = DOWN
				elevdriver.MotorDown()
			} else if elevinf.DetermineDirection() == 1 {
				elevinf.state = UP
				elevdriver.MotorUp()
			} else if elevinf.DetermineDirection() == 2 {
				elevinf.state = IDLE
			}
		case NO_EVENT:
	}
	
}

func (elevinf *Elevatorinfo) statemachineEmergency() {

	switch elevinf.event {
		case ORDER:
			for i := 0; i < 4; i++{
				if elevinf.internal_orders[i][2] == 1 {
					elevdriver.ClearStopButton()
					if elevinf.DetermineDirection() != 2 && elevdriver.GetFloor() == -1 {
						for elevdriver.GetFloor() == -1 {
							elevdriver.MotorDown()
							if elevinf.StopAtCurrentFloor() == 2 {
								elevdriver.MotorStop()
								elevinf.state = OPEN_DOOR
								elevinf.DeleteOrders()
								break
							}
							if elevdriver.GetStopButton() || elevdriver.GetObs() {
								elevinf.StopButtonPushed()
								elevinf.state = EMERGENCY
								break
							}
						}
						break
					}
					if elevinf.DetermineDirection() == -2 {
						elevinf.state = OPEN_DOOR
					} else if elevinf.DetermineDirection() == -1 {
						elevinf.state = DOWN
						elevdriver.MotorDown()
					} else if elevinf.DetermineDirection() == 1 {
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

