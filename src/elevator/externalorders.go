
package elevator

func (elevinf *Elevatorinfo) ExternalOrderDetector () {
	
	var checker int = 0
	
	for {
		//Checking for "own" external orders
		for i := 0; i < 4; i++ {
			for j := 0; j < 2; j++ {
					if elevinf.external_orders[i][j] == 1 {
						checker++
					}
			}
		}
		
		if checker > 0  {
		//I have an external order! I must share this with everyone and find out who gets to
		// put it into their internal orders!!! I MUST HURRY!
		}
		checker = 0
	}
	
}

func (elevinf *Elevatorinfo) ExternalOrderReceiver () {
	// This function will run as a thread, waiting external orders from
	// other elevators to arrive -> run cost function on that orders -> return cost
	// and then wait for a signal to either run the order or just save it.
}