package tankbattle

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/player"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *player.Player
	tank   *tank.Tank
}

func NewGame() (*Game, error) {
	return &Game{
		Player: player.New(), tank: tank.New(),
	}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.GetWindowSize()
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.Player.Update()
	g.tank.Update()
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(ebiten.NewImage(config.GetWindowSize()), &ebiten.DrawImageOptions{})
}
