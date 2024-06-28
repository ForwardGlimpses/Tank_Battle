package player

import (
	//"fmt"

	"container/list"
	"fmt"

	tankImage "github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
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

	SizeX,SizeY:=config.GetWindowSize()

	visited := make([][]bool, SizeX) 
	for i := range visited {  
		visited[i] = make([]bool, SizeY)  
	}  
	visited[dx][dy] = true

	directions := [][]int{{-20, 0}, {20, 0}, {0, -20}, {0, 20}} // 上下左右四个方向
	for queue.Len() > 0 {
		// 出队一个位置
		e := queue.Front()
		queue.Remove(e)
		pos := e.Value.(Position)

		// 遍历四个方向
		for _, dir := range directions {
			newX, newY := pos.X+dir[0], pos.Y+dir[1]
			fmt.Println(newX," ",newY)
			// 检查新位置是否合法且未访问过且不是障碍物
			fmt.Println(config.C.Window.Width,config.C.Window.Height)
			fmt.Println(visited[newX][newY])
			if newX > 0 && newX < SizeX && newY > 0 && newY < SizeY && !visited[newX][newY] {
				visited[newX][newY] = true
				t := collision.NewCollider(float64(newX), float64(newY), float64(tankImage.PlayerImage.Bounds().Dx()), float64(tankImage.PlayerImage.Bounds().Dx()))
				if check := t.Check(0,0); check != nil {
					queue.PushBack(Position{X: newX, Y: newY}) // 将新位置入队
				} else {
					return Position{X: newX, Y: newY}
				}
			}
		}
	}
	return Position{dx, dy}
}

func CreatePlayer(dx,dy int) {
	t := PlayerDetection(dx,dy)
	if t.X == dx && t.Y == dy{
		return 
	}
	NewX := t.X
	NewY := t.Y
	index %= 2
	player := &Player{
		Tank:   tank.New("Player", NewX, NewY),
		Index:  index,
		action: config.DefaultPlayers[index],
	}
	globalPlayer[player.Index] = player
	index++
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
		Create = append(Create, player.Index)
	}
	
	for _,indexx := range Create {
		CreatePlayer((indexx+2)*100,(indexx+2)*100)
	}
}

func (p *Player) Update() {
	direction, pressed := p.GetDirection()
	if pressed {
		p.Tank.Move(direction)
	}
	if p.Attack() {
		p.Tank.Fight()
	}
}

func Draw(screen *ebiten.Image) {
	for _, player := range globalPlayer {
		player.Draw(screen)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Tank.Draw(screen)
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

func GetCreatEnemy() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyQ)
}
