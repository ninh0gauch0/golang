package composition


// Board represents a surface we can work on.
type Board struct {
	NailsNeeded int
    NailsDriven int
}

// NailDriver represents behavior to drive nails into a board.
type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

// NailPuller represents behavior to remove nails into a board.
type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

// NailDrivePuller represents behavior to drive/remove nails into a board.
type NailDrivePuller interface {
	NailDriver
	NailPuller
}

// Mallet is a tool that pounds in nails.
type Mallet struct{}
// DriveNail pounds a nail into the specified board.
	func (Mallet) DriveNail(nailSupply *int, b *Board) {
		// Take a nail out of the supply.
	   *nailSupply–

	   // Pound a nail into the board.
	     b.NailsDriven++
	     fmt.Println("Mallet: pounded nail into the board.")
}

// Crowbar is a tool that removes nails.
type Crowbar struct{}
// PullNail yanks a nail out of the specified board.
func (Crowbar) PullNail(nailSupply *int, b *Board) {
	// Yank a nail out of the board.
	b.NailsDriven–

    // Put that nail back into the supply.
    *nailSupply++

    fmt.Println("Crowbar: yanked nail out of the board.")
}


// Contractor carries out the task of securing boards.
type Contractor struct{}

// Fasten will drive nails into a board. Notice we are using the NailDriver interface and not the NailDrivePuller interface for the parameter’s type
func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded { 
		d.DriveNail(nailSupply, b)
	}
}


// Unfasten will remove nails from a board.
func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}


// ProcessBoards works against boards.
func (c Contractor) ProcessBoards(dp NailDrivePuller, nailSupply *int, boards []Board) {
	for i := range boards {
    b := &boards[i]
    fmt.Printf("contractor: examining board #%d: %+v\n", i+1, b)

    switch {
	    case b.NailsDriven < b.NailsNeeded:
	        c.Fasten(dp, nailSupply, b)

         case b.NailsDriven > b.NailsNeeded:
            c.Unfasten(dp, nailSupply, b)
        }
    }
}