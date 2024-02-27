package main

import (
	"log"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := tankbattle.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowTitle("坦克大战")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
