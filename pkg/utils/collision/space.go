package collision

import (
	"math"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
)

var space *Space

// Space represents a inspector space. Internally, each Space contains a 2D array of Cells, with each Cell being the same size. Cells contain information on which
// Colliders occupy those spaces.
type Space struct {
	Cells                 [][]*Cell
	CellWidth, CellHeight int // Width and Height of each Cell in "world-space" / pixels / whatever
}

// spaceWidth and spaceHeight is the width and height of the Space (usually in pixels), which is then populated with cells of size
// cellWidth by cellHeight. Generally, you want cells to be the size of the smallest collide-able colliders in your game, and you want to move Colliders at a maximum
// speed of one cell size per inspector check to avoid missing any possible inspectors.
func Init(spaceWidth, spaceHeight, cellWidth, cellHeight int) {

	space = &Space{
		CellWidth:  cellWidth,
		CellHeight: cellHeight,
	}

	space.Resize(spaceWidth/cellWidth, spaceHeight/cellHeight)

}

// Add adds the specified Colliders to the Space, updating the Space's cells to refer to the Collider.
func (sp *Space) Add(colliders ...*Collider) {

	if sp == nil {
		panic("ERROR: space is nil")
	}

	for _, collider := range colliders {

		collider.Space = sp

		// We call Update() once to make sure the collider gets its cells added.
		collider.Update()

	}

}

// Remove removes the specified Colliders from being associated with the Space. This should be done whenever an Collider is removed from the
// game.
func (sp *Space) Remove(colliders ...*Collider) {

	if sp == nil {
		panic("ERROR: space is nil")
	}

	for _, collider := range colliders {

		for _, cell := range collider.TouchingCells {
			cell.unregister(collider)
		}

		collider.TouchingCells = []*Cell{}

		collider.Space = nil

	}

}

// Colliders loops through all Cells in the Space (from top to bottom, and from left to right) to return all Colliders
// that exist in the Space. Of course, each Collider is counted only once.
func (sp *Space) Colliders() []*Collider {

	collidersAdded := map[*Collider]bool{}
	colliders := []*Collider{}

	for cy := range sp.Cells {

		for cx := range sp.Cells[cy] {

			for _, o := range sp.Cells[cy][cx].Colliders {

				if _, added := collidersAdded[o]; !added {
					colliders = append(colliders, o)
					collidersAdded[o] = true
				}

			}

		}

	}

	return colliders

}

// Resize resizes the internal Cells array.
func (sp *Space) Resize(width, height int) {

	sp.Cells = [][]*Cell{}

	for y := 0; y < height; y++ {

		sp.Cells = append(sp.Cells, []*Cell{})

		for x := 0; x < width; x++ {
			sp.Cells[y] = append(sp.Cells[y], newCell(x, y))
		}

	}

}

// Cell returns the Cell at the given cellular / spatial (not world) X and Y position in the Space. If the X and Y position are
// out of bounds, Cell() will return nil.
func (sp *Space) Cell(x, y int) *Cell {

	if y >= 0 && y < len(sp.Cells) && x >= 0 && x < len(sp.Cells[y]) {
		return sp.Cells[y][x]
	}
	return nil

}

// CheckCells checks a set of cells (from x,y to x + w, y + h in cellular coordinates) and returns
// a slice of the colliders found within those Cells.
// The colliders must have any of the tags provided (if any are provided).
func (sp *Space) CheckCells(x, y, w, h int, tags ...string) []*Collider {

	res := []*Collider{}

	for ix := x; ix < x+w; ix++ {

		for iy := y; iy < y+h; iy++ {

			cell := sp.Cell(ix, iy)

			if cell != nil {

				if len(tags) > 0 {

					if cell.ContainsTags(tags...) {
						for _, collider := range cell.Colliders {
							if collider.HasTags(tags...) {
								res = append(res, collider)
							}
						}
					}

				} else if cell.Occupied() {
					res = append(res, cell.Colliders...)
				}

			}

		}

	}

	return res

}

// CheckWorld checks the cells of the Grid with the given world coordinates.
// Internally, this is just syntactic sugar for calling Space.WorldToSpace() on the
// position and size given.
func (sp *Space) CheckWorld(x, y, w, h float64, tags ...string) []*Collider {

	sx, sy := sp.WorldToSpace(x, y)
	cw, ch := sp.WorldToSpace(w, h)

	return sp.CheckCells(sx, sy, cw, ch, tags...)

}

// CheckWorldVec checks the cells of the Grid with the given world coordinates.
// This function takes vectors for the position and size of the checked area.
// Internally, this is just syntactic sugar for calling Space.WorldToSpace() on the
// position and size given.
func (sp *Space) CheckWorldVec(pos, size vector2.Vector, tags ...string) []*Collider {

	sx, sy := sp.WorldToSpace(pos.X, pos.Y)
	cw, ch := sp.WorldToSpace(size.X, size.Y)

	return sp.CheckCells(sx, sy, cw, ch, tags...)

}

// UnregisterAllColliders unregisters all Colliders registered to Cells in the Space.
func (sp *Space) UnregisterAllColliders() {

	for y := 0; y < len(sp.Cells); y++ {

		for x := 0; x < len(sp.Cells[y]); x++ {
			cell := sp.Cells[y][x]
			sp.Remove(cell.Colliders...)
		}

	}

}

// WorldToSpace converts from a world position (x, y) to a position in the Space (a grid-based position).
func (sp *Space) WorldToSpace(x, y float64) (int, int) {
	fx := int(math.Floor(x / float64(sp.CellWidth)))
	fy := int(math.Floor(y / float64(sp.CellHeight)))
	return fx, fy
}

// WorldToSpaceVec converts from a world position vector2.Vector to a position in the Space (a grid-based position).
func (sp *Space) WorldToSpaceVec(position vector2.Vector) (int, int) {
	return sp.WorldToSpace(position.X, position.Y)
}

// SpaceToWorld converts from a position in the Space (on a grid) to a world-based position, given the size of the Space when first created.
func (sp *Space) SpaceToWorld(x, y int) (float64, float64) {
	fx := float64(x * sp.CellWidth)
	fy := float64(y * sp.CellHeight)
	return fx, fy
}

func (sp *Space) SpaceToWorldVec(x, y int) vector2.Vector {
	outX, outY := sp.SpaceToWorld(x, y)
	return vector2.NewVector(outX, outY)
}

// Height returns the height of the Space grid in Cells (so a 320x240 Space with 16x16 cells would have a height of 15).
func (sp *Space) Height() int {
	return len(sp.Cells)
}

// Width returns the width of the Space grid in Cells (so a 320x240 Space with 16x16 cells would have a width of 20).
func (sp *Space) Width() int {
	if len(sp.Cells) > 0 {
		return len(sp.Cells[0])
	}
	return 0
}

func (sp *Space) CellsInLine(startX, startY, endX, endY int) []*Cell {

	cells := []*Cell{}
	cell := sp.Cell(startX, startY)
	endCell := sp.Cell(endX, endY)

	if cell != nil && endCell != nil {

		dv := vector2.NewVector(float64(endX-startX), float64(endY-startY)).Unit()
		dv.X *= float64(sp.CellWidth / 2)
		dv.Y *= float64(sp.CellHeight / 2)

		pX, pY := sp.SpaceToWorld(startX, startY)
		p := vector2.NewVector(pX+float64(sp.CellWidth/2), pY+float64(sp.CellHeight/2))

		alternate := false

		for cell != nil {

			if cell == endCell {
				cells = append(cells, cell)
				break
			}

			cells = append(cells, cell)

			if alternate {
				p.Y += dv.Y
			} else {
				p.X += dv.X
			}

			cx, cy := sp.WorldToSpace(p.X, p.Y)
			c := sp.Cell(cx, cy)
			if c != cell {
				cell = c
			}
			alternate = !alternate

		}

	}

	return cells

}
