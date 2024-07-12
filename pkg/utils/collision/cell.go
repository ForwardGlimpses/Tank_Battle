package collision

// Cell is used to contain and organize Collider information.
type Cell struct {
	X, Y      int         // The X and Y position of the cell in the Space - note that this is in Grid position, not World position.
	Colliders []*Collider // The Colliders that a Cell contains.
}

// newCell creates a new cell at the specified X and Y position. Should not be used directly.
func newCell(x, y int) *Cell {
	return &Cell{
		X:         x,
		Y:         y,
		Colliders: []*Collider{},
	}
}

// register registers an collider with a Cell. Should not be used directly.
func (cell *Cell) register(collider *Collider) {
	if !cell.Contains(collider) {
		cell.Colliders = append(cell.Colliders, collider)
	}
}

// unregister unregisters an collider from a Cell. Should not be used directly.
func (cell *Cell) unregister(collider *Collider) {

	for i, o := range cell.Colliders {

		if o == collider {
			cell.Colliders[i] = cell.Colliders[len(cell.Colliders)-1]
			cell.Colliders = cell.Colliders[:len(cell.Colliders)-1]
			break
		}

	}

}

// Contains returns whether a Cell contains the specified Collider at its position.
func (cell *Cell) Contains(collider *Collider) bool {
	for _, o := range cell.Colliders {
		if o == collider {
			return true
		}
	}
	return false
}

// ContainsTags returns whether a Cell contains an Collider that has the specified tag at its position.
func (cell *Cell) ContainsTags(tags ...string) bool {
	for _, o := range cell.Colliders {
		if o.HasTags(tags...) {
			return true
		}
	}
	return false
}

// Occupied returns whether a Cell contains any Colliders at all.
func (cell *Cell) Occupied() bool {
	return len(cell.Colliders) > 0
}
