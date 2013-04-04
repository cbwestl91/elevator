package main
import "elevdriver"
import "fmt"
import "time"

const FLOORS = 4

func main () {
	elevdriver.Init()
	fmt.Printf("Started!\n")
	go readButtons()
	go turnAround()
	go watchObs()
	time.Sleep(1*time.Second)
	elevdriver.MotorUp()
	for {
		select {
		case <-time.After(1*time.Second):
		}
	}
}

func readButtons () {
	var current [FLOORS][3] bool
	for {
		floor, dir := elevdriver.GetButton()
		current[floor-1][dir] = !current[floor-1][dir]
		if current[floor-1][dir] {
			elevdriver.SetLight(floor, dir)
		} else {
			elevdriver.ClearLight(floor, dir)
		}
	}
}

func turnAround () {
	for {
		floor := elevdriver.GetFloor()
		go elevdriver.SetFloor(floor)
		switch floor {
		case 1:
			go elevdriver.MotorUp()
		case 4:
			go elevdriver.MotorDown()
		}
	}
}

func watchObs () {
	for {
		
		if elevdriver.GetObs() {
			go elevdriver.SetDoor()
		} else {
			go elevdriver.ClearDoor()
		}
	}
}
