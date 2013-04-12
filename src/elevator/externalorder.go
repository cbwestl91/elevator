//-----------------------------------------------------------------------------------------//
//                                   EXTERNALORDERS                                        //
//-----------------------------------------------------------------------------------------//
package elevator

import "time"

func (elevinf *Elevatorinfo) ExternalOrderDetector () {

	checker, pos_one, pos_two, order_int, my_cost := 0, 0, 0, 0, 0
	var message string
	for { // Checking for "own" external orders
		for i := 0; i < 4; i++ {
			for j := 0; j < 2; j++ {
					if elevinf.external_orders[i][j] == 1 {
						checker++
						if checker > 0 {
							pos_one = i
							pos_two = j
							break
						}
					}
				if checker > 0 {
					break
				}
			}
		}
		
		if checker == 1  { // External order detected!
			order_int, message = OrderEncoder(pos_one,pos_two)
			my_cost = elevinf.MyCost(order_int)
			
		}
		checker = 0
		pos_one = 0
		pos_two = 0
	}
	
}

func (elevinf *Elevatorinfo) ExternalOrderReceiver () {
	
}

func (elevinf *Elevatorinfo) ExternalOrderTimer () {
	for {	
		for elevinf.external_orders[][]
	}
}

func (elevinf *Elevatorinfo) OrderEncoder (floor int, button int)(order_code int, message string){
	if floor = 0 && b = 0 {
		order_code = 0
		message = "up 1"
	} else if floor = 1 && button = 0{
		order_code = 1
		message = "up 2"
	} else if floor = 2 && button = 0{
		order_code = 2
		message = "up 3"
	} else if floor = 1 && button = 1{
		order_code = 3
		message = "down 2"
	} else if floor = 2 && button = 1{
		order_code = 4
		message = "down 3"
	} else if floor = 3 && button = 1{
		order_code = 5
		message = "down 4"
	}
	return order_code, message
}

func (elevinf *Elevatorinfo) OrderDecoder (message string)(floor int, button int, order_code int){
	if message == "up 1" {
		floor, button, order_code = 0, 0, 0
	} else if  message == "up 2" {
		floor, button, order_code = 1, 0, 1
	} else if  message == "up 3" {
		floor, button, order_code = 2, 0, 2
	} else if  message == "down 2" {
		floor, button, order_code = 1, 1, 3
	} else if  message == "down 3" {
		floor, button, order_code = 2, 1, 4
	} else if  message == "down 4" {
		floor, button, order_code = 3, 1, 5
	}
	return floor, button, order_code
}

// This function will run as a thread, waiting for external orders from
	// other elevators to arrive -> run cost function on that orders -> return cost
	// and then wait for a signal to either run the order or just save it.
	// 1. Ligge her å blokke til det kmr en ExternalOrder inn
	// 2. Oppdatere egen externalorder, Kjøre cost funksjon på gitt ordre, sende tilbake kost
	// 3. Vente på et go signal av noe slag fra sender, hvis kjør: legg external i internal
	//		Hvis ikke gå tilbake til start.
	// SIDENOTE: Når man oppdager en -1 i external_order starter man en timer som setter orderen til 1 igjen
	// hvis det går for lang tid -> da tar kanskje orderen litt lenger tid, men den blir ikke TAPT
	// Denne funskjonen sjekker hele tiden slik at hvis orderen blir 0 av en heis, slutter den timeren.
	// 		Må ha en detektor, som gjør at med en gang en heis setter sin eksterne order til 0
	//		Så gjør alle det! Dette må skje før timeren til -1 signalet oppstår
	// 4. Delete funksjonen må være slik at når en ekstern order blir fullført
	// 		må det sendes et signal til alle om at den er blitt nettopp det, slik at alle
	// 		endrer verdien i external_orders til 0 fra -1
	
	// external_orders har kun verdien 1 idet den blir oppdaget
	// må ha en annen verdi f.eks -1 mens orderen kjøres i en 
	// intern heis
