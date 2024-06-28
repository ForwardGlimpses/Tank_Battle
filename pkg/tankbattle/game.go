package tankbattle

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/enemy"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/player"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *player.Player
	Enemy  *enemy.Enemy
}

func NewGame() (*Game, error) {
	sizeX, sizeY := config.GetWindowSize()
	collision.Init(sizeX, sizeY, 2, 2)
	scenes.Init()
	player.Init()
	return &Game{
		// Player: player.New(),
		//Scense: scenes.New(),
	}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.GetWindowSize()
}

// Update updates the current game state.
func (g *Game) Update() error {
	player.Update()
	enemy.Update()
	bullet.Update()
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	player.Draw(screen)
	enemy.Draw(screen)
	bullet.Draw(screen)
	scenes.Draw(screen)
}
