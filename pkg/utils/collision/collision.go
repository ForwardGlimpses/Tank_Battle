package collision

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	// "github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
)

// Collider represents an collider that can be spread across one or more Cells in a Space. An Collider is essentially an AABB (Axis-Aligned Bounding Box) Rectangle.
type Collider struct {
	Shape         IShape             // A shape for more specific inspector-checking.
	Space         *Space             // Reference to the Space the Collider exists within
	Position      vector2.Vector     // The position of the Collider in the Space
	Size          vector2.Vector     // The size of the Collider in the Space
	TouchingCells []*Cell            // An array of Cells the Collider is touching
	Data          interface{}        // A pointer to a user-definable collider
	ignoreList    map[*Collider]bool // Set of Colliders to ignore when checking for inspectors
	tags          []string           // A list of tags the Collider has
}

// NewCollider returns a new Collider of the specified position and size.
func NewCollider(x, y, w, h float64, tags ...string) *Collider {
	col := &Collider{
		Position:   vector2.NewVector(x, y),
		Size:       vector2.NewVector(w, h),
		tags:       []string{},
		ignoreList: map[*Collider]bool{},
	}

	if len(tags) > 0 {
		col.AddTags(tags...)
	}

	space.Add(col)

	return col
}

// RandNewCollider returns a new Collider of the
func RandNewCollider(w, h float64, tags ...string) *Collider {

	col := &Collider{
		Position:   vector2.NewVector(0, 0),
		Size:       vector2.NewVector(w, h),
		tags:       []string{},
		ignoreList: map[*Collider]bool{},
	}

	var c *Cell
	for c == nil {
		rx := rand.Intn(space.Width())
		ry := rand.Intn(space.Height())
		c = space.Cell(rx, ry)
		fmt.Println(c.Occupied())
		if c.Occupied() {
			c = nil
		} else {
			col.Position.X, col.Position.Y = space.SpaceToWorld(c.X, c.Y)
		}
	}

	if len(tags) > 0 {
		col.AddTags(tags...)
	}

	space.Add(col)

	return col
}

// 销毁
func (col *Collider) Destruction() {
	space.Remove(col)
}

// Clone clones the Collider with its properties into another Collider. It also clones the Collider's Shape (if it has one).
func (col *Collider) Clone() *Collider {
	newObj := NewCollider(col.Position.X, col.Position.Y, col.Size.X, col.Size.Y, col.Tags()...)
	newObj.Data = col.Data
	if col.Shape != nil {
		newObj.SetShape(col.Shape.Clone())
	}
	for k := range col.ignoreList {
		newObj.AddToIgnoreList(k)
	}
	return newObj
}

// Update updates the collider's association to the Cells in the Space. This should be called whenever an Collider is moved.
// This is automatically called once when creating the Collider, so you don't have to call it for static colliders.
func (col *Collider) Update() {

	if col.Space != nil {

		// Collider.Space.Remove() sets the removed collider's Space to nil, indicating it's been removed. Because we're updating
		// the Collider (which is essentially removing it from its previous Cells / position and re-adding it to the new Cells /
		// position), we store the original Space to re-set it.

		space := col.Space

		col.Space.Remove(col)

		col.Space = space

		cx, cy, ex, ey := col.BoundsToSpace(0, 0)

		for y := cy; y <= ey; y++ {

			for x := cx; x <= ex; x++ {

				c := col.Space.Cell(x, y)

				if c != nil {
					c.register(col)
					col.TouchingCells = append(col.TouchingCells, c)
				}

			}

		}

	}

	if col.Shape != nil {
		col.Shape.SetPosition(col.Position.X, col.Position.Y)
	}
}

func (col *Collider) Move(move vector2.Vector) {
	col.Position = col.Position.Add(move)
	col.Update()
}

// AddTags adds tags to the Collider.
func (col *Collider) AddTags(tags ...string) {
	col.tags = append(col.tags, tags...)
}

// RemoveTags removes tags from the Collider.
func (col *Collider) RemoveTags(tags ...string) {

	for _, tag := range tags {

		for i, t := range col.tags {

			if t == tag {
				col.tags = append(col.tags[:i], col.tags[i+1:]...)
				break
			}

		}

	}

}

// HasTags indicates if an Collider has any of the tags indicated.
func (col *Collider) HasTags(tags ...string) bool {

	for _, tag := range tags {

		for _, t := range col.tags {

			if t == tag {
				return true
			}

		}

	}

	return false

}

// Tags returns the tags an Collider has.
func (col *Collider) Tags() []string {
	return append([]string{}, col.tags...)
}

// SetShape sets the Shape on the Collider, in case you need to use precise per-Shape intersection detection. SetShape calls Collider.Update() as well, so that it's able to
// update the Shape's position to match its Collider as necessary. (If you don't use this, the Shape's position might not match the Collider's, depending on if you set the Shape
// after you added the Collider to a Space and if you don't call Collider.Update() yourself afterwards.)
func (col *Collider) SetShape(shape IShape) {
	if col.Shape != shape {
		col.Shape = shape
		col.Update()
	}
}

// BoundsToSpace returns the Space coordinates of the shape (x, y, w, and h), given its world position and size, and a supposed movement of dx and dy.
func (col *Collider) BoundsToSpace(dx, dy float64) (int, int, int, int) {
	cx, cy := col.Space.WorldToSpace(col.Position.X+dx, col.Position.Y+dy)
	ex, ey := col.Space.WorldToSpace(col.Position.X+col.Size.X+dx-1, col.Position.Y+col.Size.Y+dy-1)
	return cx, cy, ex, ey
}

// SharesCells returns whether the Collider occupies a cell shared by the specified other Collider.
func (col *Collider) SharesCells(other *Collider) bool {
	for _, cell := range col.TouchingCells {
		if cell.Contains(other) {
			return true
		}
	}
	return false
}

// SharesCellsTags returns if the Cells the Collider occupies have an collider with the specified tags.
func (col *Collider) SharesCellsTags(tags ...string) bool {
	for _, cell := range col.TouchingCells {
		if cell.ContainsTags(tags...) {
			return true
		}
	}
	return false
}

// Center returns the center position of the Collider.
func (col *Collider) Center() vector2.Vector {
	return vector2.NewVector(col.Position.X+(col.Size.X/2.0), col.Position.Y+(col.Size.Y/2.0))
}

// SetCenter sets the Collider such that its center is at the X and Y position given.
func (col *Collider) SetCenter(x, y float64) {
	col.Position.X = x - (col.Size.X / 2)
	col.Position.Y = y - (col.Size.Y / 2)
}

// SetCenterVec sets the Collider such that its center is at the X and Y position given.
func (col *Collider) SetCenterVec(pos vector2.Vector) {
	col.Position.X = pos.X - (col.Size.X / 2)
	col.Position.Y = pos.Y - (col.Size.Y / 2)
}

// CellPosition returns the cellular position of the Collider's center in the Space.
func (col *Collider) CellPosition() (int, int) {
	return col.Space.WorldToSpaceVec(col.Center())
}

// SetRight sets the X position of the Collider so the right edge is at the X position given.
func (col *Collider) SetRight(x float64) {
	col.Position.X = x - col.Size.X
}

// SetBottom sets the Y position of the Collider so that the bottom edge is at the Y position given.
func (col *Collider) SetBottom(y float64) {
	col.Position.Y = y - col.Size.Y
}

// Bottom returns the bottom Y coordinate of the Collider (i.e. col.Y + col.H).
func (col *Collider) Bottom() float64 {
	return col.Position.Y + col.Size.Y
}

// Right returns the right X coordinate of the Collider (i.e. col.X + col.W).
func (col *Collider) Right() float64 {
	return col.Position.X + col.Size.X
}

func (col *Collider) SetBounds(topLeft, bottomRight vector2.Vector) {
	col.Position.X = topLeft.X
	col.Position.Y = topLeft.Y
	col.Size.X = bottomRight.X - col.Position.X
	col.Size.Y = bottomRight.Y - col.Position.Y
}

// Check checks the space around the collider using the designated delta movement (dx and dy). This is done by querying the containing Space's Cells
// so that it can see if moving it would coincide with a cell that houses another Collider (filtered using the given selection of tag strings). If so,
// Check returns a Inspector. If no colliders are found or the Collider does not exist within a Space, this function returns nil.
func (col *Collider) Check(dx, dy float64, tags ...string) *Inspector {

	if col.Space == nil {
		return nil
	}

	cc := newInspector()
	cc.checkingCollider = col

	if dx < 0 {
		dx = math.Min(dx, -1)
	} else if dx > 0 {
		dx = math.Max(dx, 1)
	}

	if dy < 0 {
		dy = math.Min(dy, -1)
	} else if dy > 0 {
		dy = math.Max(dy, 1)
	}

	cc.dx = dx
	cc.dy = dy

	cx, cy, ex, ey := col.BoundsToSpace(dx, dy)

	collidersAdded := map[*Collider]bool{}
	cellsAdded := map[*Cell]bool{}

	for y := cy; y <= ey; y++ {

		for x := cx; x <= ex; x++ {

			if c := col.Space.Cell(x, y); c != nil {

				for _, o := range c.Colliders {

					// We only want cells that have colliders other than the checking collider, or that aren't on the ignore list.
					if ignored := col.ignoreList[o]; o == col || ignored {
						continue
					}

					if _, added := collidersAdded[o]; (len(tags) == 0 || o.HasTags(tags...)) && !added {

						cc.Colliders = append(cc.Colliders, o)
						collidersAdded[o] = true
						if _, added := cellsAdded[c]; !added {
							cc.Cells = append(cc.Cells, c)
							cellsAdded[c] = true
						}
						continue

					}

				}

			}

		}

	}

	if len(cc.Colliders) == 0 {
		return nil
	}

	// ox := cc.checkingCollider.X + (cc.checkingCollider.W / 2)
	// oy := cc.checkingCollider.Y + (cc.checkingCollider.H / 2)

	oc := cc.checkingCollider.Center()

	sort.Slice(cc.Colliders, func(i, j int) bool {

		return cc.Colliders[i].Center().Sub(oc).Magnitude() < cc.Colliders[j].Center().Sub(oc).Magnitude()

	})

	cw := cc.checkingCollider.Space.CellWidth
	ch := cc.checkingCollider.Space.CellHeight

	sort.Slice(cc.Cells, func(i, j int) bool {

		return vector2.NewVector(float64(cc.Cells[i].X*cw+(cw/2)), float64(cc.Cells[i].Y*ch+(ch/2))).Sub(oc).Magnitude() <
			vector2.NewVector(float64(cc.Cells[j].X*cw+(cw/2)), float64(cc.Cells[j].Y*ch+(ch/2))).Sub(oc).Magnitude()

	})

	return cc

}

// Overlaps returns if an Collider overlaps another Collider.
func (col *Collider) Overlaps(other *Collider) bool {
	return other.Position.X <= col.Position.X+col.Size.X && other.Position.X+other.Size.X >= col.Position.X && other.Position.Y <= col.Position.Y+col.Size.Y && other.Position.Y+other.Size.Y >= col.Position.Y
}

// AddToIgnoreList adds the specified Collider to the Collider's internal inspector ignoral list. Cells that contain the specified Collider will not be counted when calling Check().
func (col *Collider) AddToIgnoreList(ignoreObj *Collider) {
	col.ignoreList[ignoreObj] = true
}

// RemoveFromIgnoreList removes the specified Collider from the Collider's internal inspector ignoral list. Colliders removed from this list will once again be counted for Check().
func (col *Collider) RemoveFromIgnoreList(ignoreObj *Collider) {
	delete(col.ignoreList, ignoreObj)
}
