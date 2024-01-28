package player

import "github.com/ForwardGlimpses/Tank_Battle/pkg/tank"

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
}

func GetDirection() (int, bool) {
	return 0, true
}
