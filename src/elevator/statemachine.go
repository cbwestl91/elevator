//-----------------------------------------------------------------------------------------//
//                                   STATEMACHINE                                          //
//-----------------------------------------------------------------------------------------//
package elevator

import "elevdriver"
import "fmt"
import "time"

type State int

const ( // Giving the states values with iote -> increments from 0
	IDLE State = iota
	ASCENDING
	DECENDING
	OPEN_DOOR 
	EMERGENCY 
)

func (elevinf *Elevatorinfo) BootStatemachine (){ // Called once, prepares elevator for use...
	fmt.Printf("STATEMACHINE BOOTING...\n")
	elevinf.last_floor = 0
	
	elevinf.Initiate()
	go FloorIndicator()
	go elevinf.CheckLights()
	go elevinf.ReceiveOrders()
	go elevinf.SetEvent()
	go elevinf.PrintStatus()
	go elevinf.UpdateLastDirection()
	
	fmt.Printf("STATEMACHINE BOOTED!\n")
}

func (elevinf *Elevatorinfo) UpdateLastDirection(){
	for{ 
		if elevinf.state == ASCENDING{
			elevinf.last_direction = 1
		} else if elevinf.state == DECENDING {
			elevinf.last_direction = 2
		}
		time.Sleep(1E7)
	}
}

func (elevinf *Elevatorinfo) RunStatemachine(){
	switch elevinf.state {
		case IDLE:
			fmt.Printf("Case Idle engaged...\n")
			elevinf.statemachineIdle()
		case ASCENDING:
			elevinf.statemachineAscending()
		case DECENDING:
			elevinf.statemachineDecending()
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
				elevinf.state = DECENDING
			} else if elevinf.DetermineDirection() == 1 {
				if elevdriver.GetFloor() == -1 {
					elevinf.Initiate()
						elevinf.state = IDLE
					break
				}
				elevinf.state = ASCENDING
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

func (elevinf *Elevatorinfo) statemachineAscending() {
	switch elevinf.event {
		case ORDER:
		case STOP:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case OBSTRUCTION:
			elevinf.StopButtonPushed()
			elevinf.state = EMERGENCY
		case SENSOR: // Destination reached || someone wants to go UP || no orders above DECENDING
			FloorIndicator()
			if elevinf.StopAtCurrentFloor() == 1 {
				elevinf.StopMotor()
				elevinf.state = OPEN_DOOR
				elevinf.DeleteOrders()
				break
			} else if elevdriver.GetFloor() == 4 {
				elevinf.StopMotor()
				elevinf.state = IDLE
			}
		case NO_EVENT:
	}
}

func (elevinf *Elevatorinfo) statemachineDecending() {
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
				elevinf.StopMotor()
				elevinf.state = OPEN_DOOR
				elevinf.DeleteOrders()
				break
			} else if elevdriver.GetFloor() == 1 {
				elevinf.StopMotor()
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
				elevinf.state = DECENDING
				StartMotor(-1)
			} else if elevinf.DetermineDirection() == 1 {
				elevinf.state = ASCENDING
				StartMotor(1)
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
				elevinf.state = DECENDING
				StartMotor(-1)
			} else if elevinf.DetermineDirection() == 1 {
				elevinf.state = ASCENDING
				StartMotor(1)
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
							StartMotor(-1)
							if elevinf.StopAtCurrentFloor() == 2 {
								elevinf.StopMotor()
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
						elevinf.state = DECENDING
						StartMotor(-1)
					} else if elevinf.DetermineDirection() == 1 {
						elevinf.state = ASCENDING
						StartMotor(1)
					}
				}
			}
		case STOP:
		case OBSTRUCTION:
		case SENSOR:
		case NO_EVENT:
	}
}

