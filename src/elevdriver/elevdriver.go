package elevdriver
import "fmt"
import "time"

type Direction int
const (
	NONE Direction = iota
	UP
	DOWN
)

type button struct {
	floor int
	dir Direction
}

const MAX_SPEED = 4024
const MIN_SPEED = 2048

func Init() {
	val := IoInit()
	if !val {
		fmt.Printf("Driver initiated\n")
	} else {
		fmt.Printf("Driver not initiated\n")
	}

	ClearDoor()
	ClearStopButton()
	ClearLight(1, UP)
	ClearLight(2, UP)
	ClearLight(3, UP)
	ClearLight(2, DOWN)
	ClearLight(3, DOWN)
	ClearLight(4, DOWN)
	ClearLight(1, NONE)
	ClearLight(2, NONE)
	ClearLight(3, NONE)
	ClearLight(4, NONE)

	buttonChan = make(chan button)
	floorChan = make(chan int)
	motorChan = make(chan Direction)
	stopButtonChan = make(chan bool)
	obsChan = make(chan bool)

	go listen()
	go motorHandler()
}

var buttonChan chan button
var floorChan chan int
var motorChan chan Direction
var stopButtonChan chan bool
var obsChan chan bool

func MotorHandler() {
	currentDir := NONE
	Write_analog(MOTOR, MIN_SPEED)
	for {
		newDir := <- motorChan
		if (newDir == NONE) && (currentDir == UP) {
			Set_bit(MOTORDIR)
			Write_analog(MOTOR, MIN_SPEED)
		} else if (newDir == NONE) && (currentDir == DOWN) {
			Clear_bit(MOTORDIR)
			Write_analog(MOTOR, MIN_SPEED)
		} else if (newDir == UP) {
			Clear_bit(MOTORDIR)
			Write_analog(MOTOR, MAX_SPEED)
		} else if (newDir == DOWN) {
			Set_bit(MOTORDIR)
			Write_analog(MOTOR, MAX_SPEED)
		} else {
			Write_analog(MOTOR, MIN_SPEED)
		}
		currentDir = newDir
	}
}

func Listen() {
	var floorMap = map[int] int {
		SENSOR1 : 1,
		SENSOR2 : 2,
		SENSOR3 : 3,
		SENSOR4 : 4,
	}

	var buttonMap = map[int] button {
		FLOOR_COMMAND1 : { 1, NONE },
		FLOOR_COMMAND2 : { 2, NONE },
		FLOOR_COMMAND3 : { 3, NONE },
		FLOOR_COMMAND4 : { 4, NONE },
		FLOOR_UP1      : { 1,   UP },
		FLOOR_UP2      : { 2,   UP },
		FLOOR_UP3      : { 3,   UP },
		FLOOR_DOWN2    : { 2, DOWN },
		FLOOR_DOWN3    : { 3, DOWN },
		FLOOR_DOWN4    : { 4, DOWN },
	}

	buttonList := make(map[int]bool)
	for key, _ := range buttonMap {
		buttonList[key] = Read_bit(key)
	}

	floorList := make(map[int]bool)
	for key, _ := range floorMap {
		floorList[key] = Read_bit(key)
	}

	//oldStop := false
	//oldObs := false
	
	// Here are the changes we did to the elevdriver!
	atFloor := false
	
	for {	
		atFloor := false
		time.Sleep(1E7)
		for key, floor := range floorMap {
			if Read_bit(key) {
				select {
				case floorChan <- floor:
				default:
				}
				atFloor = true
			}
		}
		
		if !atFloor {
			select {
			case floorChan <- -1:
			default:
			}
		}
		/*
		for key, floor := range floorMap {
			newValue := Read_bit(key)
			if newValue != floorList[key] {
				newFloor := floor
				go func() {
					floorChan <- newFloor
				}()
			}
			floorList[key] = newValue
		}
		*/
		
		for key, btn := range buttonMap {
			newValue := Read_bit(key)
			if newValue && !buttonList[key] {
				newButton := btn
				go func() {
					buttonChan <- newButton
				}()
			}
			buttonList[key] = newValue
		}
		/*
		newStop := Read_bit(STOP)
		if newStop && !oldStop {
			go func() {
				stopButtonChan <- true
			}()
		}
		oldStop = newStop

		newObs := Read_bit(OBSTRUCTION)
		if newObs != oldObs {
			go func() {
				obsChan <- newObs
			}()
		}
		oldObs = newObs
		*/
	}

}


func SetLight (floor int, dir Direction) {
	switch {
	case	floor == 1 && dir == NONE:
			Set_bit(LIGHT_COMMAND1)
	case	floor == 2 && dir == NONE:
			Set_bit(LIGHT_COMMAND2)
	case	floor == 3 && dir == NONE:
			Set_bit(LIGHT_COMMAND3)
	case	floor == 4 && dir == NONE:
			Set_bit(LIGHT_COMMAND4)
	case	floor == 1 && dir == UP:
			Set_bit(LIGHT_UP1)
	case	floor == 2 && dir == UP:
			Set_bit(LIGHT_UP2)
	case	floor == 3 && dir == UP:
			Set_bit(LIGHT_UP3)
	case	floor == 2 && dir == DOWN:
			Set_bit(LIGHT_DOWN2)
	case	floor == 3 && dir == DOWN:
			Set_bit(LIGHT_DOWN3)
	case	floor == 4 && dir == DOWN:
			Set_bit(LIGHT_DOWN4)
	}
}

func ClearLight (floor int, dir Direction) {
	switch {
	case	floor == 1 && dir == NONE:
			Clear_bit(LIGHT_COMMAND1)
	case	floor == 2 && dir == NONE:
			Clear_bit(LIGHT_COMMAND2)
	case	floor == 3 && dir == NONE:
			Clear_bit(LIGHT_COMMAND3)
	case	floor == 4 && dir == NONE:
			Clear_bit(LIGHT_COMMAND4)
	case	floor == 1 && dir == UP:
			Clear_bit(LIGHT_UP1)
	case	floor == 2 && dir == UP:
			Clear_bit(LIGHT_UP2)
	case	floor == 3 && dir == UP:
			Clear_bit(LIGHT_UP3)
	case	floor == 2 && dir == DOWN:
			Clear_bit(LIGHT_DOWN2)
	case	floor == 3 && dir == DOWN:
			Clear_bit(LIGHT_DOWN3)
	case	floor == 4 && dir == DOWN:
			Clear_bit(LIGHT_DOWN4)
	}
}

func MotorUp () {
	motorChan <- UP
}

func MotorDown () {
	motorChan <- DOWN
}

func MotorStop () {
	motorChan <- NONE
}

func GetButton () (int, Direction) {
	btn :=  <- buttonChan
	return btn.floor, btn.dir
}

func GetFloor () (int) {
	floor :=  <- floorChan
	return floor
}

func SetFloor (floor int) {
	switch floor {
	case 1:
		Clear_bit (FLOOR_IND1)
		Clear_bit (FLOOR_IND2)
	case 2:
		Clear_bit (FLOOR_IND1)
		Set_bit   (FLOOR_IND2)
	case 3:
		Set_bit   (FLOOR_IND1)
		Clear_bit (FLOOR_IND2)
	case 4:
		Set_bit   (FLOOR_IND1)
		Set_bit   (FLOOR_IND2)
	}
}

func GetStopButton () {
	// <- stopButtonChan
	return Read_bit(STOP)
}

func SetStopButton() {
	Set_bit(LIGHT_STOP)
}

func ClearStopButton() {
	Clear_bit(LIGHT_STOP)
}

func GetObs() bool {
	// return <- obsChan
	return Read_bit(OBSTRUCTION)
}

func SetDoor() {
	Set_bit(DOOR_OPEN)
}

func ClearDoor() {
	Clear_bit(DOOR_OPEN)
}
