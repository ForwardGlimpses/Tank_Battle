package player

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	Up int = iota
	Down
	Left
	Right
)

type Player struct {
	Tank *tank.Tank
}

func New() *Player {
	return &Player{
		Tank: tank.New(),
	}
}

func (p *Player) Update() {
	direction, pressed := GetDirection()
	if pressed {
		p.Tank.Move(direction)
		//p.Tank.Rotate()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
		p.Tank.Draw(screen)  
}
func GetDirection() (int, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return Up, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return Left, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return Right, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		return Down, true
	}
	return 0, false
}

