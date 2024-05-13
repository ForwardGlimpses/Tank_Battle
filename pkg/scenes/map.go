package scenes

import (
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	// "github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	// "github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
	// "github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

// 使用二维数组表示地图

type World struct {
	Space       *resolv.Space
	Geometry    []*resolv.Object // 四周围墙
	Bullets     []*Bullet     // 所有小球
	Map         [][]ScenesType
}

func (world *World) Init() {
	world.Map = [][]ScenesType{
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{2, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		{2, 0, 3, 0, 0, 0, 0, 0, 3, 2},
		{2, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		{2, 0, 0, 0, 0, 1, 0, 0, 0, 2},
		{2, 0, 0, 2, 0, 0, 0, 3, 0, 2},
		{2, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		{2, 1, 0, 0, 0, 0, 0, 0, 2, 2},
		{2, 0, 0, 0, 0, 0, 0, 0, 0, 2},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	}
}
func (world *World) Init(){

	gw ,gh := 100, 100

	cellSize := 8

	// gw, gh 为碰撞网格的长和高，cellSize 为碰撞检测的精度（cellSize*cellSize 的方格内就算碰撞）
	World.Space = resolv.NewSpace(gw, gh, cellSize, cellSize)

	world.Geometry = []*resolv.Object{
		resolv.NewObject(0, 0, 16, float64(gh)),
		resolv.NewObject(float64(gw-16), 0, 16, float64(gh)),
		resolv.NewObject(0, 0, float64(gw), 16),
		resolv.NewObject(0, float64(gh-24), float64(gw), 32),
	}

	world.Space.Add(world.Geometry...)
	
	world.Bullets = []*Bullet{}

}