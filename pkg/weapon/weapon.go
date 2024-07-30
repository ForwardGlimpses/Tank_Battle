package weapon

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
)

var (
	weapons = make(map[int]Weapon)
)
type Weapon interface {
	Fight(position vector2.Vector, direction direction.Direction,camp string)
}

// 当前武器是一个抽象概念，不需要实际的图片
type DefaultWeapon struct {
	Damage int
}

func (D *DefaultWeapon) Fight(position vector2.Vector, direction direction.Direction,camp string) {
	opt := &bullet.CreateOption{
		Position:  position,
		Direction: direction,
		Camp: camp,
	}
	bullet.Create(opt)
}

func GetWeapon(Type int) Weapon {
    return weapons[Type]
}

func init() {
    weapons[0] = &DefaultWeapon{Damage: 10}
}