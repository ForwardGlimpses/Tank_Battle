package tank

import (
	"image"
	_ "image/png"
	"math"

	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/weapon"
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
	Hp        int
	position  *vector2.Vector2
	direction direction.Direction
	weapon    weapon.Weapon
	Image     image.Image
}

func New() *Tank {
	return &Tank{
		position: vector2.New(28, 25),
		Hp:       100,
		weapon:   &weapon.DefaultWeapon{},
		Image:    tank.PlayerImage,
	}
}

func (t *Tank) Move(direction direction.Direction) {
	t.direction = direction
	t.position = t.position.Add(direction.DirectionVector2().MulScalar(step))

	// TODO: 通过碰撞检测限制移动

	// Width, Height := config.GetWindowSize()
	// MinWidth, MinHeight := config.GetWindowLimit()
	// if t.dx < MinHeight {
	// 	t.dx = MinHeight
	// }
	// if t.dy < MinWidth {
	// 	t.dy = MinWidth
	// }
	// if t.dx > Height {
	// 	t.dx = Height
	// }
	// if t.dy > Width {
	// 	t.dy = Width
	// }
}

func (t *Tank) Fight() {
	// TODO: 计算子弹发射位置（坦克正前方）
	t.weapon.Fight(t.position, t.direction)
}

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(55)/2, -float64(49)/2)
	opt.GeoM.Rotate(t.direction.Theta() * 2 * math.Pi / 360)
	opt.GeoM.Translate(t.position.ValueFloat64())
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
