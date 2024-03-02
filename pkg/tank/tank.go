package tank

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"math"
)

const (
	Up int = iota
	Down
	Left
	Right
	step int = 10
)

type Tank struct {
	dx    int
	dy    int
	Hp    int
	count int
	Image image.Image
}

func New() *Tank {
	return &Tank{
		dx:    0,
		dy:    0,
		Hp:    100,
		Image: tank.PlayerImage,
	}
}

func (t *Tank) Move(direction int) {
	t.count += 90
	if direction == Up {
		t.dy -= step
	} else if direction == Down {
		t.dy += step
	} else if direction == Left {
		t.dx -= step
	} else {
		t.dx += step
	}
	Width, Height := config.GetWindowSize()
	MinWidth, MinHeight := config.GetWindowLimit()
	if t.dx < MinHeight {
		t.dx = MinHeight
	}
	if t.dy < MinWidth {
		t.dy = MinWidth
	}
	if t.dx > Height {
		t.dx = Height
	}
	if t.dy > Width {
		t.dy = Width
	}
}

// func (t *Tank) Rotate(direction int) {
// 	theta := 0.0
// 	if direction == t.direction {
// 		return ;
// 	} else if t.direction - direction == 2||direction-t.direction == 2{
// 		theta = 1.0
// 	} else if direction - t.direction == 1||t.direction - direction == 3{
// 		theta = 0.45
// 	} else {
// 		theta = -0.45
// 	}

// 	sin, cos := math.Sincos(theta)

// 	tx := cos*float64(t.dx) - sin*float64(t.dy)
// 	ty := sin*float64(t.dx) + cos*float64(t.dy)

// 	t.dx = int(tx)
// 	t.dy = int(ty)
// 	t.direction = direction
// }

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Rotate(float64(t.count%360) * 2 * math.Pi / 360)
	opt.GeoM.Translate(float64(t.dx), float64(t.dy))
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
