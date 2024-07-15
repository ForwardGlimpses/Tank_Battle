package player

import (

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/ebitenextend"
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


func CreatePlayer(dx, dy ,indexx int) {
	t := tank.TankBorn(dx, dy)
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
}

func (p *Player) GetDirection() (direction.Direction, bool) {
	KeyUp , _ := ebitenextend.KeyNameToKeyCode(config.DefaultPlayers[p.Index].Up)
	KeyDown , _ := ebitenextend.KeyNameToKeyCode(config.DefaultPlayers[p.Index].Down)
	KeyLeft , _ := ebitenextend.KeyNameToKeyCode(config.DefaultPlayers[p.Index].Left)
	KeyRight , _ := ebitenextend.KeyNameToKeyCode(config.DefaultPlayers[p.Index].Right)
	if ebiten.IsKeyPressed(ebiten.Key(KeyUp)) {
		return direction.Up, true
	}
	if ebiten.IsKeyPressed(ebiten.Key(KeyDown)){
		return direction.Down, true
	}
	if ebiten.IsKeyPressed(ebiten.Key(KeyLeft)) {
		return direction.Left, true
	}
	if ebiten.IsKeyPressed(ebiten.Key(KeyRight)) {
		return direction.Right, true
	}
	return 0, false
}

func (p *Player) Attack() bool {
	attack , _ := ebitenextend.KeyNameToKeyCode(config.DefaultPlayers[p.Index].Attack)
	return inpututil.IsKeyJustPressed(ebiten.Key(attack))
}
