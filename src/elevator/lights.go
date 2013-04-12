//-----------------------------------------------------------------------------------------//
//                                   LIGHTS                                                //
//-----------------------------------------------------------------------------------------//
package elevator

import "elevdriver"
import "fmt"
import "time"

func (elevinf *Elevatorinfo) CheckLights (){  // Updates lights according to the order_arrays!
	// for {
		for i := 1; i < 4; i++ {
			if elevinf.internal_orders[i-1][0] == 1 {
				elevdriver.SetLight(i, 1)
			} else if elevinf.internal_orders[i-1][0] == 0 {					
				elevdriver.ClearLight(i, 1)
			}
		}
		for i := 2; i < 5; i++ {
			if elevinf.internal_orders[i-1][1] == 1 {
				elevdriver.SetLight(i, 2)
			} else if elevinf.internal_orders[i-1][1] == 0 {
				elevdriver.ClearLight(i, 2)
			}
		}
		for i := 0; i < 4; i++ {
			if elevinf.internal_orders[i][2] == 1 {
				elevdriver.SetLight(i+1, 0)
			} else if elevinf.internal_orders[i][2] == 0 {
				elevdriver.ClearLight(i+1, 0)
			}
		}
		fmt.Printf("checkedlights\n")
		// time.Sleep(100*time.Millisecond)
	// }
}

func FloorIndicator(){
	for {
		if elevdriver.GetFloor()  > 0 { 
			elevdriver.SetFloor(elevdriver.GetFloor())
		}
		time.Sleep(100*time.Millisecond)
	}
}


