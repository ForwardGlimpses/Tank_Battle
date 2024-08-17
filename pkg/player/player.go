package player

import (
	"fmt"
	"strconv"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/configmanager"
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

func Init() (err error) {
	for _, cfg := range config.DefaultPlayers {
		player := New(cfg)
		combinedKey := fmt.Sprintf("%s%s", player.PlayerUuid, player.Index)
		globalPlayer[combinedKey] = player
	}
	C, err := configmanager.LoadConfig("C:\\Users\\乔书祥\\Desktop\\远程文件库\\Tank_Battle\\config.json")
	Cfg := C.Network
	if Cfg.Type == "client" {
		IsCreatTank = false
	} else {
		IsCreatTank = true
	}
	return
}

var (
	index       = 0
	IsCreatTank = true
)

func New(cfg config.Player) *Player {
	index++
	indexStr := fmt.Sprintf("%d", index)
	if IsCreatTank {
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
	} else {
		return &Player{
			//Tank:       &tank.Tank{},//tank.New("Player", (index+2)*100, (index+2)*100),
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
}

func Update() {
	for _, player := range globalPlayer {
		if player.Tank!=nil {
			player.Update()
		}
	}
}

func (p *Player) Update() {

	if p.Tank.Hp <= 0 {
		p.Reset()
	}

	if p.PlayerUuid == Uuid {
		p.GetDirection()
		p.Attack()
	}
	p.Tank.Move = p.Action.Move
	p.Tank.Attack = p.Action.Attack
	p.Tank.Direction = p.Action.Direction
	if p.Action.Attack {
		fmt.Println("攻击----------")
	}
	//fmt.Println("下标：",p.Tank.Index,"方向：",p.Tank.Direction)
}
func (p *Player) Reset() {
	dx, _ := strconv.Atoi(p.Index)
	dy, _ := strconv.Atoi(p.Index)
	p.Tank = tank.New("Player", (dx+2)*100, (dy+2)*100)
}

func (p *Player) GetDirection() {
	op := p.Operate

	//fmt.Println("")
	if p.PlayerUuid != Uuid {
		return
	}
	p.Action.Move = false
	if ebiten.IsKeyPressed(op.Up) {
		//fmt.Println("上----")
		p.Action.Direction = direction.Up
		p.Action.Move = true
	}
	if ebiten.IsKeyPressed(op.Down) {
		//fmt.Println("下----")
		p.Action.Direction = direction.Down
		p.Action.Move = true
	}
	if ebiten.IsKeyPressed(op.Left) {
		//fmt.Println("左----")
		p.Action.Direction = direction.Left
		p.Action.Move = true
	}
	if ebiten.IsKeyPressed(op.Right) {
		//fmt.Println("右----")
		p.Action.Direction = direction.Right
		p.Action.Move = true
	}
}

func (p *Player) Attack() {
	p.Action.Attack = inpututil.IsKeyJustPressed(p.Operate.Attack)
}
