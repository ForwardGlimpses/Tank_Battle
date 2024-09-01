package player

import (
	"fmt"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/ebitenextend"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	TankIndex    int
	Index        string
	Operate      Operate
	Action       Action
	NetworkCount int
	Local        bool
}

type Operate struct {
	Up     ebiten.Key
	Down   ebiten.Key
	Left   ebiten.Key
	Right  ebiten.Key
	Attack ebiten.Key
}

type Action struct {
	Direction direction.Direction
	Attack    bool
	Move      bool
}

func init() {
	tankbattle.RegisterInit(Init, 30)
	tankbattle.RegisterUpdate(Update, 20)
}

var globalPlayer = make(map[string]*Player)

func Init() (err error) {
	for _, cfg := range config.DefaultPlayers {
		player := New(cfg)
		globalPlayer[player.Index] = player
	}
	return
}

var (
	index = 0
)

func New(cfg config.Player) *Player {
	index++
	return &Player{
		TankIndex: tank.New("Player", 100, 100).Index,
		Index:     fmt.Sprintf("%s%d", Uuid, index),
		Local:     true,
		Operate: Operate{
			Up:     ebitenextend.KeyNameToKeyCode(cfg.Up),
			Down:   ebitenextend.KeyNameToKeyCode(cfg.Down),
			Left:   ebitenextend.KeyNameToKeyCode(cfg.Left),
			Right:  ebitenextend.KeyNameToKeyCode(cfg.Right),
			Attack: ebitenextend.KeyNameToKeyCode(cfg.Attack),
		},
	}
}

func Update() {
	for _, player := range globalPlayer {
		player.Update()
	}
}

func (p *Player) Update() {
	if p.Local {
		p.GetDirection()
		p.Attack()
	}
	t := tank.Get(p.TankIndex)
	t.Move = p.Action.Move
	t.Attack = p.Action.Attack
	t.Direction = p.Action.Direction
	//fmt.Println("-----------")
}

func (p *Player) GetDirection() {
	op := p.Operate

	p.Action.Move = false
	if ebiten.IsKeyPressed(op.Up) {
		p.Action.Direction = direction.Up
		p.Action.Move = true
	}
	if ebiten.IsKeyPressed(op.Down) {
		p.Action.Direction = direction.Down
		p.Action.Move = true
	}
	if ebiten.IsKeyPressed(op.Left) {
		p.Action.Direction = direction.Left
		p.Action.Move = true
	}
	if ebiten.IsKeyPressed(op.Right) {
		p.Action.Direction = direction.Right
		p.Action.Move = true
	}
}

func (p *Player) Attack() {
	p.Action.Attack = ebiten.IsKeyPressed(p.Operate.Attack)
}
