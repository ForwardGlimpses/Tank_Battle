package tank

import (
	"image"

	//"log"
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Up int = iota
	Down
	Left
	Right
	step int = 1
)

type Tank struct {
	dx    int
	dy    int
	Hp    int
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
	if direction == Up {
		t.dx -= step
	} else if direction == Down {
		t.dx += step
	} else if direction == Left {
		t.dy -= step
	} else {
		t.dy += step
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

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(t.dx), float64(t.dy))
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
