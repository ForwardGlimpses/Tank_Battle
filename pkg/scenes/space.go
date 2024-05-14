package scenes

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/solarlune/resolv"
)

var space *resolv.Space

func init() {
	width, height := config.GetWindowSize()
	space = resolv.NewSpace(width, height, 3, 3)
}

func SpaceAdd(objects ...*resolv.Object) {
	space.Add(objects...)
}

func SpaceRemove(objects ...*resolv.Object) {
	space.Remove(objects...)
}
