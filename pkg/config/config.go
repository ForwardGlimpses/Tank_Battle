package config

import (
	//"strings"

	"github.com/hajimehoshi/ebiten/v2"
)


var C Config

var (
	defaultConfig = &Config{
		Window: Window{
			Width:     600,
			Height:    600,
			MinWidth:  0,
			MinHeight: 0,
		},
	}
	DefaultPlayers = [2]Action{
		{
			Up: "ArrowUp",
			Down: "ArrowDown",
			Left: "ArrowLeft",
			Right: "ArrowRight",
			Attack: "Space",
		},
		{
			Up: "W",
			Down: "S",
			Left: "A",
			Right: "D",
			Attack: "J",
		},
	}
)

type Config struct {
	Window Window
	Players []Action
	//Network Network
	Plat string
}

type Action struct {
	Up string
	Down string
	Left string
	Right string
	Attack string
}

type Window struct {
	Width     int
	Height    int
	MinHeight int
	MinWidth  int
}

func KeyMap(name string)ebiten.Key{
	switch name{
	case "0":
		return ebiten.Key0
	case "1":
		return ebiten.Key1
	case "2":
		return ebiten.Key2
	case "3":
		return ebiten.Key3
	case "4":
		return ebiten.Key4
	case "5":
		return ebiten.Key5
	case "6":
		return ebiten.Key6
	case "7":
		return ebiten.Key7
	case "8":
		return ebiten.Key8
	case "9":
		return ebiten.Key9
	case "Q":
		return ebiten.KeyQ
	case "W":
		return ebiten.KeyW
	case "E":
		return ebiten.KeyE
	case "R":
		return ebiten.KeyR
	case "A":
		return ebiten.KeyA
	case "S":
		return ebiten.KeyS
	case "ArrowUp":
		return ebiten.KeyArrowUp
	case "ArrowDown":
		return ebiten.KeyArrowDown
	case "ArrowLeft":
		return ebiten.KeyArrowLeft
	case "ArrowRight":
		return ebiten.KeyArrowRight
	case "Space":
		return ebiten.KeySpace
	case "F":
		return ebiten.KeyF
	case "G":
		return ebiten.KeyG
	case "H":
		return ebiten.KeyH
	case "J":
		return ebiten.KeyJ
	case "K":
		return ebiten.KeyK
	case "L":
		return ebiten.KeyL
	case "Z":
		return ebiten.KeyZ
	case "D":
		return ebiten.KeyD
	default:
	    return ebiten.KeyF11
	}
}

func GetWindowSize() (int, int) {
	return defaultConfig.Window.Width, defaultConfig.Window.Height
}
func GetWindowLimit() (int, int) {
	return defaultConfig.Window.MinWidth, defaultConfig.Window.MinHeight
}
