package player

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
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
	}
	if GetAttack() {
		p.Tank.Fight()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Tank.Draw(screen)
}

func GetDirection() (direction.Direction, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return direction.Up, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return direction.Left, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return direction.Right, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		return direction.Down, true
	}
	return 0, false
}

func GetAttack() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}
