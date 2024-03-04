package tank

import (
	//"fmt"
	"image"
	_ "image/png"
	"math"

	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	 "github.com/ForwardGlimpses/ebiten/v2"
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

// func (t *Tank) Rotate() {

// 	sin, cos := math.Sincos(t.theta)

// 	tx := cos*float64(t.dx) - sin*float64(t.dy)
// 	ty := sin*float64(t.dx) + cos*float64(t.dy)
// 	t.dx = int(tx)
// 	t.dy = int(ty)
// 	// radian := angle * math.Pi / 180.0
//     // var angle = this.rotate * Math.PI / 180;
// 	// rx := this.point.x + this.point.width / 2, ry = this.point.y + this.point.height / 2; // the rotation x and y
// 	// px := rx, py := ry; // the objects center x and y
// 	// radius := ry - py; // the difference in y positions or the radius
// 	// dx := rx + radius * math.sin(angle); // the draw x 
// 	// dy := ry - radius * math.cos(angle); // the draw y
// 	// canvas.translate(dx, dy);
// 	// canvas.rotate(angle);
// 	// canvas.translate(-dx, -dx);
// }

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	 opt.GeoM.Rotate(float64(t.theta%360) * 2 * math.Pi / 360)
	opt.GeoM.Translate(float64(t.dx), float64(t.dy))
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
