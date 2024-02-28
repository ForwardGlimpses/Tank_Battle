package tank

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
)

var (
	//go:embed tank0.png
	Tank0_png   []byte
	PlayerImage image.Image
)

func init() {
	PlayerImage, _, _ = image.Decode(bytes.NewReader(Tank0_png))
}
