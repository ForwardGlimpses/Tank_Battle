package config

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
	DefaultPlayers = [2]Player{
		{
			Up:     "ArrowUp",
			Down:   "ArrowDown",
			Left:   "ArrowLeft",
			Right:  "ArrowRight",
			Attack: "Space",
		},
		{
			Up:     "W",
			Down:   "S",
			Left:   "A",
			Right:  "D",
			Attack: "J",
		},
	}
)

type Config struct {
	Window  Window
	Players []Player
	//Network Network
	Plat string
}

type Player struct {
	Up     string
	Down   string
	Left   string
	Right  string
	Attack string
}

type Window struct {
	Width     int
	Height    int
	MinHeight int
	MinWidth  int
}

func GetWindowSize() (int, int) {
	return defaultConfig.Window.Width, defaultConfig.Window.Height
}
func GetWindowLimit() (int, int) {
	return defaultConfig.Window.MinWidth, defaultConfig.Window.MinHeight
}
