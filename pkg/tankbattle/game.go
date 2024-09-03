package tankbattle

import (
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
	insertSort(&initList, &initPriorities, f, priority)
}

func RegisterUpdate(f func(), priority int) {
	insertSort(&updateList, &updatePriorities, f, priority)
}

func RegisterDraw(f func(screen *ebiten.Image), priority int) {
	insertSort(&drawList, &drawPriorities, f, priority)
}

func insertSort[T any](funcs *[]T, priorities *[]int, f T, priority int) {
	i := len(*priorities)
	for i > 0 && (*priorities)[i-1] > priority {
		i--
	}
	*funcs = append(*funcs, f)
	copy((*funcs)[i+1:], (*funcs)[i:])
	(*funcs)[i] = f

	*priorities = append(*priorities, priority)
	copy((*priorities)[i+1:], (*priorities)[i:])
	(*priorities)[i] = priority
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
