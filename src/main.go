// Trying to get this main to work!
package main

import "elevator"

type State int

const (
	IDLE State = iota // iota gives states int from 0 (increment)
	UP 
	DOWN 
	OPEN_DOOR 
	EMERGENCY 
)

type Event int

const (
	ORDER Event = iota
	STOP
	OBSTRUCTION
	SENSOR
	NO_EVENT
)

func main () {
	elevinf := elevator.Elevatorinfo{}

	elevinf.HandleElevator()
	return

}

