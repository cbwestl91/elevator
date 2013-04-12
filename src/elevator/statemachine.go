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
	// go elevinf.CheckLights()
	// go elevinf.ReceiveOrders()
	// go elevinf.SetEvent()
	// go elevinf.PrintStatus()
	// go elevinf.UpdateLastDirection()
	
	fmt.Printf("STATEMACHINE BOOTED!\n")
}

func (elevinf *Elevatorinfo) UpdateLastDirection(){
	// for{ 
		if elevinf.state == ASCENDING{
			elevinf.last_direction = 1
		} else if elevinf.state == DECENDING {
			elevinf.last_direction = 2
		}
		fmt.Printf("Updated direction\n")
		time.Sleep(100*time.Millisecond)
	// }
}

func (elevinf *Elevatorinfo) RunStatemachine(){
	for {
		elevinf.UpdateLastDirection()
		elevinf.ReceiveOrders()
		elevinf.CheckLights()
		elevinf.SetEvent()
		elevinf.PrintStatus()
		
		switch elevinf.state {
			case IDLE:
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
		fmt.Printf("ran statemachine\n")
	}
}

func (elevinf *Elevatorinfo) statemachineIdle() {
	// for elevinf.state == IDLE {
		time.Sleep(1*time.Millisecond)
		switch elevinf.event {
			case ORDER:
				if elevinf.DetermineDirection() != 2 && elevdriver.GetFloor() == -1 {
					elevinf.Initiate()
					elevinf.state = IDLE
					break
				}
				if elevinf.DetermineDirection() == -2 {
					elevinf.state = OPEN_DOOR
				} else if elevinf.DetermineDirection() == -1 {
					elevinf.state = DECENDING
					StartMotor(-1)
				} else if elevinf.DetermineDirection() == 1 {
					/*if elevdriver.GetFloor() == -1 {
						elevinf.Initiate()
							elevinf.state = IDLE
						break
					}*/
					elevinf.state = ASCENDING
					StartMotor(1)
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
	// }
}

func (elevinf *Elevatorinfo) statemachineAscending() {
	// for elevinf.state == ASCENDING {
		time.Sleep(1*time.Millisecond)
		switch elevinf.event {
			case ORDER:
			case STOP:
				elevinf.StopButtonPushed()
				elevinf.state = EMERGENCY
			case OBSTRUCTION:
				elevinf.StopButtonPushed()
				elevinf.state = EMERGENCY
			case SENSOR: // Destination reached || someone wants to go UP || no orders above DECENDING
				fmt.Printf("FLOOR DETECTED\n")
				if elevinf.StopAtCurrentFloor() == 1 {
					fmt.Printf("GOING TO STOP HERE\n")
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
	// }
}

func (elevinf *Elevatorinfo) statemachineDecending() {
	// for elevinf.state == DECENDING {
		switch elevinf.event {
			case ORDER:
			case STOP:
				elevinf.StopButtonPushed()
				elevinf.state = EMERGENCY
			case OBSTRUCTION:
				elevinf.StopButtonPushed()
				elevinf.state = EMERGENCY
			case SENSOR:
				fmt.Printf("FLOOR DETECTED\n")
				if elevinf.StopAtCurrentFloor() == -1 {
					fmt.Printf("GOING TO STOP HERE\n")
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
		time.Sleep(1*time.Millisecond)
	// }
}

func (elevinf *Elevatorinfo) statemachineOpendoor() {
	// for elevinf.state == OPEN_DOOR {
		switch elevinf.event {
			case ORDER:
				fmt.Printf("The door is open\n")
				elevdriver.SetDoor()
				elevinf.DeleteOrders()
				for i := 0; i < 30; i++{
					if elevdriver.GetFloor() == -1 && elevdriver.GetObs() == false {
						elevdriver.ClearDoor()
						elevinf.state = IDLE
						break
					} else if elevdriver.GetStopButton() == true {
						elevinf.StopButtonPushed()
						elevinf.state = EMERGENCY
						break
					}
					elevinf.ReceiveOrders()
					elevinf.CheckLights()
					time.Sleep(10*time.Millisecond)
					if elevdriver.GetObs() == true {
						fmt.Printf("Obstruction detected, door staying open\n")
						i = 0
					}
				}
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
					elevinf.ReceiveOrders()
					elevinf.CheckLights()
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
		time.Sleep(1*time.Millisecond)
	// }
}

func (elevinf *Elevatorinfo) statemachineEmergency() {
	// for elevinf.state == EMERGENCY {
		switch elevinf.event {
			case ORDER:
				for i := 0; i < 4; i++{
					if elevinf.internal_orders[i][2] == 1 {
						elevdriver.ClearStopButton()
						if elevinf.DetermineDirection() != 2 && elevdriver.GetFloor() == -1 {
							for elevdriver.GetFloor() == -1 {
								StartMotor(-1)
								elevinf.ReceiveOrders()
								elevinf.CheckLights()
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
		time.Sleep(1*time.Millisecond)
	// }
}

