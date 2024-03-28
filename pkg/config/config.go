package config

var (
	defaultConfig = &Config{
		Window: Window{
			Width:  600,
			Height: 600,
			MinWidth: 0,
			MinHeight: 0,
		},
	}
)

type Config struct {
	Window Window
}

type Window struct {
	Width  int
	Height int
	MinHeight int
	MinWidth int
}

func GetWindowSize() (int, int) {
	return defaultConfig.Window.Width, defaultConfig.Window.Height
}
func GetWindowLimit() (int, int) {
	return defaultConfig.Window.MinWidth, defaultConfig.Window.MinHeight
}
