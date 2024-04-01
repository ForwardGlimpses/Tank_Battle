package main

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
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