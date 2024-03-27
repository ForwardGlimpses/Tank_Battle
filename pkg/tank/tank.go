package tank

import (
	//"fmt"
	"image"
	_ "image/png"
	"math"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/hajimehoshi/ebiten/v2"
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
	theta  int
	Image image.Image
	Buttle *bullet.Bullet
}

func New() *Tank {
	return &Tank{
		dx:    28,
		dy:    25,
		Hp:    100,
		Image: tank.PlayerImage,
		Buttle: bullet.New(),
	}
}

func (t *Tank) Move(direction int) {
	if direction == Up {
		t.dy -= step
		t.theta = 0
	} else if direction == Down {
		t.dy += step
		t.theta = 180
	} else if direction == Left {
		t.dx -= step
		t.theta = -90
	} else {
		t.dx += step
		t.theta = 90
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

func (t *Tank) Fight (IS_pressed bool) {
	if IS_pressed {
		t.Buttle.Move(t.dx,t.dy)
	}

}

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(55)/2, -float64(49)/2)
	opt.GeoM.Rotate(float64(t.theta%360) * 2 * math.Pi / 360)
	opt.GeoM.Translate(float64(t.dx), float64(t.dy))
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
