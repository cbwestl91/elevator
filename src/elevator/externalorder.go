//-----------------------------------------------------------------------------------------//
//                                   EXTERNALORDERS                                        //
//-----------------------------------------------------------------------------------------//
package elevator

import "time"
import "network"
import "strconv"
import "fmt"

var(
	sendToAll chan string
	sendToOne chan network.DecodedMessage
	receivedCostchan chan receivedCost
	receivedGochan chan bool
	receivedNoGochan chan bool
	receiveDeletion chan string
	
	gochan chan string // IP string
	noGochan chan string
	
	incExternal chan network.DecodedMessage
)

type receivedCost struct {
	IP string
	cost int
}
	

func (elevinf *Elevatorinfo) ExternalOrderMaster (communicator network.CommChannels) {
	for {
		checker, pos1, pos2, order_int:= 0, 0, 0, 0
		var message string
		var my_cost int
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
			//	convCost := strconv
				
				sendToAll <- message
				
				// CREATE MAP::::                        important :       
				var costMap map[string]int
				costMap = make(map[string]int)
				
				communicator.GiveMeCurrentAlives <- true
				aliveMap := <- communicator.GetCurrentAlives
				
				howManyIPs := len(aliveMap)
				
				for len(costMap) < howManyIPs {
					select {
						case costStruct := <- receivedCostchan:
							costMap[costStruct.IP] = costStruct.cost
						case <- time.After(100*time.Millisecond):
						break
					}
				}
				currentBest := my_cost
				currentBestIP := "Handle self"
				
				for ip, _ := range costMap {
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
			pos1 = 0
			pos2 = 0
		}
		time.Sleep(10*time.Millisecond)
	}
}

func (elevinf *Elevatorinfo) ExternalOrderSlave () {
	for {
		receiver := <- incExternal
		ip := receiver.IP
		message := receiver.Content
		
		pos1, pos2, _ := OrderUnpacker(message)
		
		my_cost := elevinf.MyCost(pos1)
		
		cost := strconv.Itoa(my_cost)
		
		
		decoded := network.DecodedMessage{ip, cost}
		
		sendToOne <- decoded
		
		select {
			case <- receivedGochan:
				elevinf.internal_orders[pos1][pos2] = 1
			case <- receivedNoGochan:
				elevinf.external_orders[pos1][pos2] = -1
			default:
				time.Sleep(time.Millisecond)
		}
		time.Sleep(10*time.Millisecond)
	}
}

func (elevinf *Elevatorinfo) ExternalOrderTimer () {
	for {
		time.Sleep(100*time.Millisecond)
		for i := 0; i > 4; i++ {
			for j := 0; j > 3; j++ {
				time.Sleep(100*time.Millisecond)
				if elevinf.external_orders[i][j] == -1 {
					pos1,pos2 := i,j
					checker := true
					for k := 0; k < 12; k++ {
						if elevinf.external_orders[pos1][pos2] == 0 {
							i,j = 4,3
							checker = false
						}
						time.Sleep(time.Second)
					}
					if checker {
						elevinf.external_orders[pos1][pos2] = 1
					}
				}
				fmt.Printf("EOT 1\n")
			}
		}	
	}
}

func ExternalSendDelete (pos1 int, pos2 int) {
	message := DeletionPacker(pos1,pos2)
	sendToAll <- message
}

func (elevinf *Elevatorinfo) ExternalRecvDelete (){
	for{
		fmt.Printf("ERD 1\n")
		receiver := <- receiveDeletion
		floor, button := DeletionUnpacker(receiver)
		elevinf.external_orders[floor][button] = 0
	}
}

func OrderPacker (floor int, button int)(order_code int, message string){
	if floor == 0 && button == 0 {
		order_code = 0
		message = "up 1"
	} else if floor == 1 && button == 0{
		order_code = 1
		message = "up 2"
	} else if floor == 2 && button == 0{
		order_code = 2
		message = "up 3"
	} else if floor == 1 && button == 1{
		order_code = 3
		message = "down 2"
	} else if floor == 2 && button == 1{
		order_code = 4
		message = "down 3"
	} else if floor == 3 && button == 1{
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

func DeletionPacker (floor int, button int)(message string){
	if floor == 0 && button == 0 {
		message = "del up 1"
	} else if floor == 1 && button == 0{
		message = "del up 2"
	} else if floor == 2 && button == 0{
		message = "del up 3"
	} else if floor == 1 && button == 1{
		message = "del down 2"
	} else if floor == 2 && button == 1{
		message = "del down 3"
	} else if floor == 3 && button == 1{
		message = "del down 4"
	}
	return message
}

func DeletionUnpacker (message string)(floor int, button int){
	if message == "del up 1" {
		floor, button = 0, 0
	} else if  message == "del up 2" {
		floor, button = 1, 0
	} else if  message == "del up 3" {
		floor, button = 2, 0
	} else if  message == "del down 2" {
		floor, button = 1, 1
	} else if  message == "del down 3" {
		floor, button = 2, 1
	} else if  message == "del down 4" {
		floor, button = 3, 1
	}
	return floor, button
}

