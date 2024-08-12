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
}

type Operate struct {
	Up     ebiten.Key
	Down   ebiten.Key
	Left   ebiten.Key
	Right  ebiten.Key
	Attack ebiten.Key
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

	direction, pressed := p.GetDirection()
	if pressed {
		p.Tank.Direction = direction
		p.Tank.Move = true
	} else {
		p.Tank.Move = false
	}
	if p.Attack() {
		p.Tank.Attack = true
		fmt.Println("攻击-----攻击-----")
	}
	//fmt.Println("已更新，已更新，已更新---")
}

func (p *Player) Reset() {
	dx, _ := strconv.Atoi(p.Index)
	dy, _ := strconv.Atoi(p.Index)
	p.Tank = tank.New("Player", (dx+2)*10, (dy+2)*100)
}

func (p *Player) GetDirection() (direction.Direction, bool) {
	op := p.Operate

	if p.PlayerUuid != Uuid {
		return 0, false
	}
	if ebiten.IsKeyPressed(op.Up) {
		//fmt.Println("上上上上上")
		return direction.Up, true
	}
	if ebiten.IsKeyPressed(op.Down) {
		return direction.Down, true
	}
	if ebiten.IsKeyPressed(op.Left) {
		return direction.Left, true
	}
	if ebiten.IsKeyPressed(op.Right) {
		return direction.Right, true
	}
	//fmt.Println("移动移动移动---")
	return 0, false
}

func (p *Player) Attack() bool {
	return inpututil.IsKeyJustPressed(p.Operate.Attack)
}
