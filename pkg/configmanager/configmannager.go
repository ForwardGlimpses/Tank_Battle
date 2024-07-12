package configmanager

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
)


func init() {
	tankbattle.RegisterInit(Init,1)
}

func Init() error{
	sizeX, sizeY := config.GetWindowSize()
	collision.Init(sizeX, sizeY, 2, 2)
	return nil
}