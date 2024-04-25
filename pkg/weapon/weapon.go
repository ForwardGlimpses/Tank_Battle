package weapon

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
//	"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/solarlune/resolv"
)

type Weapon interface {
	Fight(position resolv.Vector, direction direction.Direction)
}

// 当前武器是一个抽象概念，不需要实际的图片
type DefaultWeapon struct {
	Damage int
}

func (D *DefaultWeapon) Fight(position resolv.Vector, direction direction.Direction) {
	opt := &bullet.CreateOption{
		Position:  position,
		Direction: direction,
	}
	bullet.Create(opt)
}
