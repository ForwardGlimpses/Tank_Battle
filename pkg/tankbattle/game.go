package tankbattle

import (
	//"sort"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/hajimehoshi/ebiten/v2"
)

var initList []func() error 
var initPriorities []int

var updateList []func()
var updatePriorities []int

var drawList []func(screen *ebiten.Image)
var drawPriorities []int

func RegisterInit(f func() error, priority int) {
	initList = append(initList, f)
	initPriorities = append(initPriorities, priority)
	sortFuncsByPriority(initList,initPriorities)
}

func RegisterUpdate(f func(), priority int) {
	updateList = append(updateList, f)
	updatePriorities = append(updatePriorities, priority)
	sortFuncsByPriority(updateList,updatePriorities)
}

func RegisterDraw(f func(screen *ebiten.Image), priority int) {
	drawList = append(drawList, f)
	drawPriorities = append(drawPriorities, priority)
	sortFuncsByPriority(drawList,drawPriorities)
}

func sortFuncsByPriority[T any](funcs []T, priorities []int) {
	for i := 1; i < len(priorities); i++ {
		keyFunc := funcs[i]
		keyPriority := priorities[i]
		j := i - 1

		for j >= 0 && priorities[j] >= keyPriority {
			funcs[j+1] = funcs[j]
			priorities[j+1] = priorities[j]
			j = j - 1
		}
		funcs[j+1] = keyFunc
		priorities[j+1] = keyPriority
	}
}


type Game struct {
}

func NewGame() (*Game, error) {
	for _, f := range initList {
		if err := f(); err != nil {
			return nil, err
		}
	}
	return &Game{}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.GetWindowSize()
}

// Update updates the current game state.
func (g *Game) Update() error {
	for _, f := range updateList {
		f()
	}

	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, f := range drawList {
		f(screen)
	}
}
