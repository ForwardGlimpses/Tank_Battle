package player

import (

	"container/list"
	tankImage "github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	Up int = iota
	Down
	Left
	Right
)

var index = 0

var Tanknumer = 2

type Player struct {
	Tank   *tank.Tank
	Index  int
	action config.Action
}

func init() {
	tankbattle.RegisterInit(Init, 3)
	tankbattle.RegisterUpdate(Update, 2)
}

var globalPlayer = make(map[int]*Player)

func Init() error {
	for _, actions := range config.DefaultPlayers {
		player := &Player{
			Tank:   tank.New("Player", (index+2)*100, (index+2)*100),
			Index:  index,
			action: actions,
		}
		globalPlayer[player.Index] = player
		index++
	}
	return nil
}

type Position struct {
	X int
	Y int
}

func PlayerDetection(dx, dy int) Position {

	// 使用list包实现队列
	queue := list.New()
	queue.PushBack(Position{X: dx, Y: dy}) // 将根节点入队

	SizeX, SizeY := config.GetWindowSize()

	visited := make([][]bool, SizeX)
	for i := range visited {
		visited[i] = make([]bool, SizeY)
	}
	visited[dx][dy] = true

	directions := [][]int{{-10, 0}, {10, 0}, {0, -10}, {0, 10}} // 上下左右四个方向
	for queue.Len() > 0 {
		// 出队一个位置
		e := queue.Front()
		queue.Remove(e)
		pos := e.Value.(Position)

		// 遍历四个方向
		for _, dir := range directions {
			newX, newY := pos.X+dir[0], pos.Y+dir[1]
			// 检查新位置是否合法且未访问过且不是障碍物
			if newX > 0 && newX < SizeX && newY > 0 && newY < SizeY {
				if visited[newX][newY] {
					continue
				}
				visited[newX][newY] = true
				t := collision.NewCollider(float64(newX), float64(newY), float64(tankImage.PlayerImage.Bounds().Dx()), float64(tankImage.PlayerImage.Bounds().Dx()))
				if check := t.Check(float64(dx), float64(dy)); check != nil {
					queue.PushBack(Position{X: newX, Y: newY}) // 将新位置入队
				} else {
					return Position{X: newX, Y: newY}
				}
			}
		}
	}
	return Position{dx, dy}
}

func CreatePlayer(dx, dy ,indexx int) {
	t := PlayerDetection(dx, dy)
	if t.X == dx && t.Y == dy {
		return
	}
	NewX := t.X
	NewY := t.Y
	player := &Player{
		Tank:   tank.New("Player", NewX, NewY),
		Index:  indexx,
		action: config.DefaultPlayers[indexx],
	}
	globalPlayer[player.Index] = player
}

func Update() {

	var Destroyed []Player
	var Create []int

	for _, player := range globalPlayer {
		if player.Tank.Hp <= 0 {
			Destroyed = append(Destroyed, *player)
		} else {
			player.Update()
		}
	}

	for _, player := range Destroyed {
		delete(globalPlayer, player.Index)
		player.Tank.Collider.Destruction()
		delete(tank.GlobalTanks, player.Tank.Index)
		Create = append(Create, player.Index)
	}

	for _, indexx := range Create {
		CreatePlayer((indexx+2)*100, (indexx+2)*100,indexx)
	}
}

func (p *Player) Update() {
	direction, pressed := p.GetDirection()
	if pressed {
		p.Tank.Direction = direction
		p.Tank.Move = true
	} else {
		p.Tank.Move = false
	}
	if p.Attack() {
		p.Tank.Attack = true
	}
	tank.GlobalTanks[p.Tank.Index] = p.Tank
}

func (p *Player) GetDirection() (direction.Direction, bool) {
	if ebiten.IsKeyPressed(config.KeyMap((config.DefaultPlayers[p.Index].Up))) {
		return direction.Up, true
	}
	if ebiten.IsKeyPressed(config.KeyMap((config.DefaultPlayers[p.Index].Down))) {
		return direction.Down, true
	}
	if ebiten.IsKeyPressed(config.KeyMap((config.DefaultPlayers[p.Index].Left))) {
		return direction.Left, true
	}
	if ebiten.IsKeyPressed(config.KeyMap((config.DefaultPlayers[p.Index].Right))) {
		return direction.Right, true
	}
	return 0, false
}

func (p *Player) Attack() bool {
	return inpututil.IsKeyJustPressed(config.KeyMap(config.DefaultPlayers[p.Index].Attack))
}
