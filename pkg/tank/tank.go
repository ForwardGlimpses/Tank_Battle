package tank

import (
	//"fmt"
	"image"
	_ "image/png"
	"math"

	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"

	//	"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/weapon"
	"github.com/hajimehoshi/ebiten/v2"

	//	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

const (
	Up int = iota
	Down
	Left
	Right
	step float64 = 5
)

type Tank struct {
	Hp        int
	Position  resolv.Vector
	direction direction.Direction
	weapon    weapon.Weapon
	Image     image.Image
}

func New() *Tank {
	return &Tank{
		Position: resolv.NewVector(28, 25),
		Hp:       100,
		weapon:   &weapon.DefaultWeapon{},
		Image:    tank.PlayerImage,
	}
}

func (t *Tank) Move(direction direction.Direction) {
	t.direction = direction
	t.Position = t.Position.Add(direction.DirectionVector2().Scale(step))

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
	t.weapon.Fight(t.Position, t.direction)
}

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(55)/2, -float64(49)/2)
	opt.GeoM.Rotate(t.direction.Theta() * 2 * math.Pi / 360)
	opt.GeoM.Translate(t.Position.X,t.Position.Y)
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
