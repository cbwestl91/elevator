//-----------------------------------------------------------------------------------------//
//                                   EXTERNALORDERS                                        //
//-----------------------------------------------------------------------------------------//
package elevator

import "network"

func (elevinf *Elevatorinfo) ExternalOrderDetector () {

	var checker int = 0
	
	for { //Checking for "own" external orders
		for i := 0; i < 4; i++ {
			for j := 0; j < 2; j++ {
					if elevinf.external_orders[i][j] == 1 {
						checker++
					}
			}
		}
		
		if checker > 0  { // External order detected!
			
		}
		checker = 0
	}
	
}

func (elevinf *Elevatorinfo) ExternalOrderReceiver () {
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
}

func (elevinf *Elevatorinfo) ExternalOrderTimer () {

}


// external_orders har kun verdien 1 idet den blir oppdaget
// må ha en annen verdi f.eks -1 mens orderen kjøres i en 
// intern heis