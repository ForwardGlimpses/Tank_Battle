package main

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	_"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	_"github.com/ForwardGlimpses/Tank_Battle/pkg/bullet"
	_"github.com/ForwardGlimpses/Tank_Battle/pkg/player"
	_"github.com/ForwardGlimpses/Tank_Battle/pkg/enemy"
	_"github.com/ForwardGlimpses/Tank_Battle/pkg/configmanager"
	_"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	_"github.com/ForwardGlimpses/Tank_Battle/pkg/network/protocol/tcp"
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