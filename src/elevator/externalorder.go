//-----------------------------------------------------------------------------------------//
//                                   EXTERNALORDERS                                        //
//-----------------------------------------------------------------------------------------//
package elevator

import "time"
import "network"

var(
	sendToAll chan string
	sendToOne chan string
	receivedCostchan chan int
	receivedGochan chan bool
	receivedNoGochan chan bool
	
	gochan chan string // IP string
	noGochan chan string
	
	incExternal chan network.DecodedMessage
)

var receivedCost struct {
	IP string
	cost int
}
	

func (elevinf *Elevatorinfo) ExternalOrderMaster () {
	for {
		checker, pos1, pos2, order_int, my_cost := 0, 0, 0, 0, 0
		var message string
		for { // Checking for "own" external orders
			for i := 0; i < 4; i++ {
				for j := 0; j < 2; j++ {
						if elevinf.external_orders[i][j] == 1 {
							checker++
							if checker > 0 {
								pos1 = i
								pos2 = j
								break
							}
						}
					if checker > 0 {
						break
					}
				}
			}
		
			if checker == 1  { // External order detected!
				order_int, message = OrderPacker(pos1,pos2)
				
				my_cost = elevinf.MyCost(order_int)
				
				sendToAll <- message
				
				// CREATE MAP::::                        important :       var costMap map[string]int
				costMap = make(map[string]int)
				
				communicator.GiveMeCurrentAlives <- true
				aliveMap := <- communicator.GetCurrentAlives
				
				howManyIPs := len(aliveMap)
				
				for len(costMap) < howManyIPs {
					select {
						case costStruct := <- receivedCostchan
							costMap[costStruct.IP] = costStruct.IP
						case <- time.After(100*time.Millisecond)
						break
					}
				}
				currentBest := my_cost
				currentBestIP := "Handle self"
				
				for ip, _ := range(costMap){
					if costMap[ip] < currentBest {
						currentBest = costMap[ip]
						currentBestIP = ip
					}
				}
				
				if currentBestIP == "Handle self" {
					// handle self and send nogo to everyone else
					elevinf.internal_orders[pos1][pos2] = 1
					for ip, _ := range(costMap) {
						noGochan <- ip
					}
				} else {
					for ip, _ := range(costMap) {
						if costMap[ip] == currentBest {
							gochan <- ip
						} else {
							noGochan <- ip
						}
					}
				}
				elevinf.external_orders[pos1][pos2] = -1
			}
			checker = 0
			pos_one = 0
			pos_two = 0
		}
	}
}

func (elevinf *Elevatorinfo) ExternalOrderSlave () {
	for {
		receiver := <- incExternal
		ip := receiver.IP
		message := receiver.Content
		
		pos1, pos2, order_int := OrderUnpacker(message)
		
		my_cost := elevinf.MyCost(order_int)
		
		decoded := network.DecodedMessage{ip, message}
		
		sendToOne <- decoded
		
		select {
			case <- receivedGochan:
				elevinf.internal_orders[pos1][pos2] = 1
			case <- receivedNoGochan:
				elevinf.external_orders[pos1][pos2] = -1
			default:
				time.Sleep(time.Millisecond)
		}
	}
}

func (elevinf *Elevatorinfo) ExternalOrderTimer () {
	for {	
		for elevinf.external_orders[][]
	}
}

func OrderPacker (floor int, button int)(order_code int, message string){
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

func OrderUnpacker (message string)(floor int, button int, order_code int){
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
