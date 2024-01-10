package tankbattle

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

func NewGame() (*Game, error) {
	return &Game{}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 0, 0
}

// Update updates the current game state.
func (g *Game) Update() error {
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
}
