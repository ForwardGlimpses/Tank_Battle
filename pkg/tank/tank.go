package tank

import (
	"container/list"
	//"fmt"
	//"fmt"
	//"fmt"
	//"fmt"
	"image"
	_ "image/png"
	"math"

	tankassets "github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/weapon"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Up int = iota
	Down
	Left
	Right
	step        float64 = 3
	PlayerImage         = 0
	EnemyImage          = 1
)

var GlobalTanks = make(map[int]*Tank)

var (
	TankIndex = 0
)

type Tank struct {
	Hp          int
	Collider    *collision.Collider
	Direction   direction.Direction
	weapon      weapon.Weapon
	Image       image.Image
	Attack      bool
	Move        bool
	Camp        string
	Index       int
	PlayerIndex int
}

type Position struct {
	X int
	Y int
}

func New(camp string, tankx int, tanky int, playerindex int) *Tank {
	position := TankBorn(tankx, tanky)
	tank := &Tank{
		Collider:    collision.NewCollider(float64(position.X), float64(position.Y), float64(tankassets.PlayerImage.Bounds().Dx()), float64(tankassets.PlayerImage.Bounds().Dy())),
		Hp:          100,
		weapon:      weapon.Weapons[0],
		Image:       tankassets.TankImage[camp],
		Camp:        camp,
		Index:       TankIndex,
		PlayerIndex: playerindex,
	}
	tank.Collider.Data = tank
	GlobalTanks[tank.Index] = tank
	TankIndex++
	return tank
}

func Get(index int) *Tank {
	return GlobalTanks[index]
}

func init() {
	tankbattle.RegisterDraw(Draw, 10)
	tankbattle.RegisterUpdate(Update, 30)
}

func (t *Tank) Update(direction direction.Direction) {
	t.Direction = direction
	increment := direction.DirectionVector2().MulScale(step)
	dx := increment.X
	dy := increment.Y
	stop := false
	if check := t.Collider.Check(dx, dy); check != nil {
		for _, obj := range check.Colliders {
			if tt, ok := obj.Data.(types.Obstacle); ok {
				if !tt.TankIsPassable() {
					stop = true
				}
			}
		}
	}
	if !stop {
		t.Collider.Position = t.Collider.Position.Add(direction.DirectionVector2().MulScale(step))
	}
	t.Collider.Update()
}

func (t *Tank) SetPosition(position Position) {
	t.Collider.Position.X = float64(position.X)
	t.Collider.Position.Y = float64(position.Y)
}

func Update() {
	var Destroyed []Tank
	cfg := config.C.Network
	for _, tank := range GlobalTanks {
		if tank.Hp <= 0 {
			Destroyed = append(Destroyed, *tank)
		} else if tank.Move {
			tank.Update(tank.Direction)
		}
	}

	for _, tank := range Destroyed {
		if tank.Camp == "Player" {
			if cfg.Type == "server" {
				d := TankBorn((tank.Index+2)*100, (tank.Index+1)*100)
				GlobalTanks[tank.Index].Collider.Position.X = float64(d.X)
				GlobalTanks[tank.Index].Collider.Position.Y = float64(d.Y)
				GlobalTanks[tank.Index].Collider.Update()
				GlobalTanks[tank.Index].Hp = 100
			}
		} else {
			tank.Collider.Destruction()
			delete(GlobalTanks, tank.Index)
		}
	}

	for _, tank := range GlobalTanks {
		tank.weapon.Cooling()
		if tank.Attack {
			tank.Fight()
			//fmt.Println("攻击----")
			tank.Attack = false
		}
	}
}

func (t *Tank) Fight() {
	t.weapon.Fight(t.Collider.Position, t.Direction, t.Camp, t.PlayerIndex)
}

func TankBorn(dx, dy int) Position {

	queue := list.New()
	queue.PushBack(Position{X: dx, Y: dy})
	SizeX, SizeY := config.GetWindowSize()
	visited := make([][]bool, SizeX)
	for i := range visited {
		visited[i] = make([]bool, SizeY)
	}

	visited[dx][dy] = true
	directions := [][]int{{-20, 0}, {20, 0}, {0, -20}, {0, 20}}
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		pos := e.Value.(Position)
		for _, dir := range directions {
			newX, newY := pos.X+dir[0], pos.Y+dir[1]
			if newX > 0 && newX < SizeX && newY > 0 && newY < SizeY {
				if visited[newX][newY] {
					continue
				}
				visited[newX][newY] = true
				t := collision.NewCollider(float64(dx), float64(dy), float64(tankassets.PlayerImage.Bounds().Dx()), float64(tankassets.PlayerImage.Bounds().Dy()))
				if check := t.Check(float64(newX-dx), float64(newY-dy)); check != nil {
					queue.PushBack(Position{X: newX, Y: newY})
				} else {
					return Position{X: newX, Y: newY}
				}
			}
		}
	}
	return Position{dx, dy}
}

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	tranX := float64(t.Image.Bounds().Dx()) / 2
	tranY := float64(t.Image.Bounds().Dy()) / 2
	opt.GeoM.Translate(-tranX, -tranY)
	opt.GeoM.Rotate(t.Direction.Theta() * 2 * math.Pi / 360)
	opt.GeoM.Translate(t.Collider.Position.X+tranX, t.Collider.Position.Y+tranY)
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}

func Draw(screen *ebiten.Image) {
	for _, tank := range GlobalTanks {
		tank.Draw(screen)
	}
}

func (t *Tank) TankIsPassable() bool {
	return false
}

func (t *Tank) BulletIsPassable() bool {
	return false
}

func (t *Tank) GetCamp() string {
	return t.Camp
}

func (t *Tank) TakeDamage(damage int) {
	t.Hp -= damage
}
