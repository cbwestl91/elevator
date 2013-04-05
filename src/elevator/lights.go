package elevator

import "elevdriver"
import "fmt"
import "time"

// checks pressed buttons and set lights accordingly
func CheckLights(int state, int event, [][]int orderarray)(){
	
	if state != 4 || (state == 4 && event == 0) {
		for i := 0; i < 3; i++ {
			if orderarray[i][0] == 1 {
				elevdriver.SetLight(i, UP)
			}
			else if orderarray[i][0] == 0 {
				elevdriver.ClearLight(i, UP)
			}
		}
		for i := 1; i < 4; i++ {
			if orderarray[i][0] == 1 {
				elevdriver.SetLight(i, DOWN)
			}
			else if orderarray[i][0] == 0 {
				elevdriver.ClearLight(i, DOWN)
			}
		}
	}
	
	for i := 0; i < 4; i++ {
		if array[k][2] == 1 {
			elevdriver.SetLight(k, NONE)
		}
		else if array[k][2] == 0 {
			elevdriver.ClearLight(k, NONE)
		}
	}
}



