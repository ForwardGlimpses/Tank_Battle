package player

import (
	"fmt"
	"strconv"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/ebitenextend"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Tank       *tank.Tank
	PlayerUuid string
	Index      string
	Operate    Operate
	Action     Action
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
	tankbattle.RegisterInit(Init, 3)
	tankbattle.RegisterUpdate(Update, 2)
}

var globalPlayer = make(map[string]*Player)

func Init() error {
	for _, cfg := range config.DefaultPlayers {
		player := New(cfg)
		combinedKey := fmt.Sprintf("%s%s", player.PlayerUuid, player.Index)
		globalPlayer[combinedKey] = player
	}
	return nil
}

var index = 0

func New(cfg config.Player) *Player {
	index++
	indexStr := fmt.Sprintf("%d", index)
	return &Player{
		Tank:       tank.New("Player", (index+2)*100, (index+2)*100),
		PlayerUuid: Uuid,
		Index:      indexStr,
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
	if p.Tank.Hp <= 0 {
		p.Reset()
	}

	p.GetDirection()
	p.Attack()
	p.Tank.Move = p.Action.Move
	p.Tank.Attack = p.Action.Attack
	p.Tank.Direction = p.Action.Direction
}

func (p *Player) Reset() {
	dx, _ := strconv.Atoi(p.Index)
	dy, _ := strconv.Atoi(p.Index)
	p.Tank = tank.New("Player", (dx+2)*10, (dy+2)*100)
}

func (p *Player) GetDirection() {
	op := p.Operate

	if p.PlayerUuid != Uuid {
		return
	}

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
	p.Action.Move = false
}

func (p *Player) Attack() {
	p.Action.Attack = inpututil.IsKeyJustPressed(p.Operate.Attack)
}
