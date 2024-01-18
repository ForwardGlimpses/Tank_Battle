package config

var (
	defaultConfig = &Config{
		Window: Window{
			Width:  600,
			Height: 600,
		},
	}
)

type Config struct {
	Window Window
}

type Window struct {
	Width  int
	Height int
}

func GetWindowSize() (int, int) {
	return defaultConfig.Window.Width, defaultConfig.Window.Height
}
