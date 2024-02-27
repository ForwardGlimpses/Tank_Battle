package tank

import (

	"bytes"
	"image"
	_ "image/png"
	//"log"
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/hajimehoshi/ebiten/v2"
)
var (
	m image.Image
)

const (
	Up int = iota
	Down
	Left
	Right
	step int = 1 
)

type Tank struct {
	dx int
	dy int
	Hp int
}

func New() *Tank {
	return &Tank{}
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
	Width , Height := config.GetWindowSize()
	MinWidth , MinHeight := config.GetWindowLimit()
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
func (t *Tank) Update() {

}

func (t *Tank) Layout(w, h int) (int, int) {
	return t.dx, t.dy
}
func (t *Tank) Analysis() {
	m, _, _ = image.Decode(bytes.NewReader(tank.Tank0_png))
	// if err := ebiten.RunGame(&Tank{}); err != nil {
	// 	log.Fatal(err)
	// }
}

func (t *Tank) Draw(screen *ebiten.Image) {

	screen.DrawImage(ebiten.NewImageFromImage(m), &ebiten.DrawImageOptions{})
}

