
// All events that can happen will be covered in this part

package elevator

import (
	"elevdriver"
	"fmt"
	"time"
)

type Event int

const (
	STOP Event = iota
	WAIT
	GO_UP 
	GO_DOWN
	STOPBUTTON
	OBSTRUCTION 
)


