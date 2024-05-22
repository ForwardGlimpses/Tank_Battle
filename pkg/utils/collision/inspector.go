package collision

import "github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"

// Inspector contains the results of an Collider.Check() call, and represents a inspector between an Collider and cells that contain other Colliders.
// The Colliders array indicate the Colliders collided with.
type Inspector struct {
	checkingCollider *Collider   // The checking collider
	dx, dy           float64     // The delta the checking collider was moving on that caused this inspector
	Colliders        []*Collider // Slice of colliders that were collided with; sorted according to distance to calling Collider.
	Cells            []*Cell     // Slice of cells that were collided with; sorted according to distance to calling Collider.
}

func newInspector() *Inspector {
	return &Inspector{
		Colliders: []*Collider{},
	}
}

// HasTags returns whether any colliders within the Inspector have all of the specified tags. This slice does not contain the Collider that called Check().
func (cc *Inspector) HasTags(tags ...string) bool {

	for _, o := range cc.Colliders {

		if o == cc.checkingCollider {
			continue
		}
		if o.HasTags(tags...) {
			return true
		}

	}

	return false
}

// CollidersByTags returns a slice of Colliders from the cells reported by a Inspector collider by searching for Colliders with a specific set of tags.
// This slice does not contain the Collider that called Check().
func (cc *Inspector) CollidersByTags(tags ...string) []*Collider {

	colliders := []*Collider{}

	for _, o := range cc.Colliders {

		if o == cc.checkingCollider {
			continue
		}
		if o.HasTags(tags...) {
			colliders = append(colliders, o)
		}

	}

	return colliders

}

// ContactWithCollider returns the delta to move to have the checking collider come into contact with the specified Collider.
func (cc *Inspector) ContactWithCollider(collider *Collider) vector2.Vector {

	delta := vector2.NewVector(0, 0)

	if cc.dx < 0 {
		delta.X = collider.Position.X + collider.Size.X - cc.checkingCollider.Position.X
	} else if cc.dx > 0 {
		delta.X = collider.Position.X - cc.checkingCollider.Size.X - cc.checkingCollider.Position.X
	}

	if cc.dy < 0 {
		delta.Y = collider.Position.Y + collider.Size.Y - cc.checkingCollider.Position.Y
	} else if cc.dy > 0 {
		delta.Y = collider.Position.Y - cc.checkingCollider.Size.Y - cc.checkingCollider.Position.Y
	}

	return delta

}

// ContactWithCell returns the delta to move to have the checking collider come into contact with the specified Cell.
func (cc *Inspector) ContactWithCell(cell *Cell) vector2.Vector {

	delta := vector2.NewVector(0, 0)

	cx := float64(cell.X * cc.checkingCollider.Space.CellWidth)
	cy := float64(cell.Y * cc.checkingCollider.Space.CellHeight)

	if cc.dx < 0 {
		delta.X = cx + float64(cc.checkingCollider.Space.CellWidth) - cc.checkingCollider.Position.X
	} else if cc.dx > 0 {
		delta.X = cx - cc.checkingCollider.Size.X - cc.checkingCollider.Position.X
	}

	if cc.dy < 0 {
		delta.Y = cy + float64(cc.checkingCollider.Space.CellHeight) - cc.checkingCollider.Position.Y
	} else if cc.dy > 0 {
		delta.Y = cy - cc.checkingCollider.Size.Y - cc.checkingCollider.Position.Y
	}

	return delta

}

// SlideAgainstCell returns how much distance the calling Collider can slide to avoid a inspector with the targetCollider, and
// a boolean indicating if such a slide was possible.
// This only works on vertical and horizontal axes (x and y directly), primarily for platformers / top-down games.
// avoidTags is a sequence of tags (as strings) to indicate when sliding is valid (i.e. if a Cell contains an
// Collider that has the tag given in the avoidTags slice, then sliding CANNOT happen).
func (cc *Inspector) SlideAgainstCell(cell *Cell, avoidTags ...string) (vector2.Vector, bool) {

	sp := cc.checkingCollider.Space

	collidingCell := cc.Cells[0]
	ccX, ccY := sp.SpaceToWorld(collidingCell.X, collidingCell.Y)
	hX := float64(sp.CellWidth) / 2.0
	hY := float64(sp.CellHeight) / 2.0

	ccX += hX
	ccY += hY

	center := cc.checkingCollider.Center()

	diffX := center.X - ccX
	diffY := center.Y - ccY

	left := sp.Cell(collidingCell.X-1, collidingCell.Y)
	right := sp.Cell(collidingCell.X+1, collidingCell.Y)
	up := sp.Cell(collidingCell.X, collidingCell.Y-1)
	down := sp.Cell(collidingCell.X, collidingCell.Y+1)

	slide := vector2.NewVector(0, 0)

	// Moving vertically
	if cc.dy != 0 {

		if diffX > 0 && (right == nil || !right.ContainsTags(avoidTags...)) {
			// Slide right
			slide.X = ccX + hX - cc.checkingCollider.Position.X
		} else if diffX < 0 && (left == nil || !left.ContainsTags(avoidTags...)) {
			// Slide left
			slide.X = ccX - hX - (cc.checkingCollider.Position.X + cc.checkingCollider.Size.X)
		} else {
			return vector2.NewVector(0, 0), false
		}
	}

	if cc.dx != 0 {
		if diffY > 0 && (down == nil || !down.ContainsTags(avoidTags...)) {
			// Slide down
			slide.Y = ccY + hY - cc.checkingCollider.Position.Y
		} else if diffY < 0 && (up == nil || !up.ContainsTags(avoidTags...)) {
			// Slide up
			slide.Y = ccY - hY - (cc.checkingCollider.Position.Y + cc.checkingCollider.Size.Y)
		} else {
			return vector2.NewVector(0, 0), false
		}
	}

	return slide, true

}
