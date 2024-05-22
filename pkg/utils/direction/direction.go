package direction

import "github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
	Static
)

func (d *Direction) TurnRight() {
	*d = (*d + 1) % 4
}

func (d *Direction) TurnBack() {
	*d = (*d + 2) % 4
}

func (d *Direction) TurnLeft() {
	*d = (*d + 3) % 4
}

func (d *Direction) Theta() float64 {
	return float64(*d) * 90
}

func (d *Direction) DirectionVector2() vector2.Vector {
	switch *d {
	case Up:
		return vector2.NewVector(0, -1)
	case Right:
		return vector2.NewVector(1, 0)
	case Down:
		return vector2.NewVector(0, 1)
	case Left:
		return vector2.NewVector(-1, 0)
	}
	return vector2.NewVector(0, 0)
}
