package tankbattle

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/player"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *player.Player
	//Scense *scenes.Scenes
}

func NewGame() (*Game, error) {
	return &Game{
		Player: player.New(),
		//Scense: scenes.New(),
	}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.GetWindowSize()
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.Player.Update()
	bullet.Update()
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
	bullet.Draw(screen)
	scenes.New().Draw(screen)
}
