package tankbattle

import (
	"sort"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/hajimehoshi/ebiten/v2"
)

var InitList []Initfunc
var UpdateList []Updatefunc
var DrawList []Drawfunc

type Initfunc struct {
	function func() error
	priority int
}

type Updatefunc struct {
	function func()
	priority int
}

type Drawfunc struct {
	function func(screen *ebiten.Image)
	priority int
}

func RegisterInit(a func() error,b int) {
	InitList = append(InitList,Initfunc{
		function: a,
		priority: b,
	})
	sort.Slice(InitList,func(i int, j int) bool{
		return InitList[i].priority < InitList[j].priority
	})
	//fmt.Println("Init")
}

func RegisterUpdate(a func(),b int) {
	UpdateList = append(UpdateList,Updatefunc{
		function: a,
		priority: b,
	})
	sort.Slice(UpdateList,func(i int, j int) bool{
		return UpdateList[i].priority < UpdateList[j].priority
	})
	//fmt.Println("Update")
}

func RegisterDraw(a func(screen *ebiten.Image),b int) {
	DrawList = append(DrawList,Drawfunc{
		function: a,
		priority: b,
	})
	sort.Slice(DrawList,func(i int, j int) bool{
		return DrawList[i].priority < DrawList[j].priority
	})
	//fmt.Println("Draw")
}

type Game struct {

}

func NewGame() (*Game, error) {
	// sizeX, sizeY := config.GetWindowSize()
	// collision.Init(sizeX, sizeY, 2, 2)
	for _,f :=range InitList {
		if err :=f.function(); err !=nil {
			return nil , err
		}
		//f.function()
	}
	return &Game{}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.GetWindowSize()
}

// Update updates the current game state.
func (g *Game) Update() error{
	for _,f :=range UpdateList {
		f.function()
	}
	
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	for _,f :=range DrawList {
		f.function(screen)
	}

	
}
