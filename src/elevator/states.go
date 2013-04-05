
// Here the entire statemachine for the elevator will be declared

package elevator

import (
	"elevdriver"
	"fmt"
	"time"
)

// Creating states
type State int

const (
	START State = iota // iota gives states int from 0 (increment)
	IDLE 
	UP 
	DOWN 
	STOPPED 
	EMERGENCY 
)




