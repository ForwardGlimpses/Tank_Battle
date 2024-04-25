package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	Width  = 640
	Height = 360
)

// 此例子源于 github.com/solarlune/resolv/examples/worldBouncer.go
func main() {
	ebiten.RunGame(NewGame())
}
