
package elevator

import "elevdriver"
import "fmt"
import "time"

type Event int

const (
	ORDER Event = iota
	STOP
	OBSTRUCTION
	SENSOR
	NO_EVENT
)


