package scenes

// 使用二维数组表示地图
type Map [][]ScenesType

// TODO: 写一个默认的地图，根据 Window 大小编写（也可以修改一下 Window）
var defMap = Map{
	{0, 0, 1},
	{1, 1, 1},
	{1, 0, 0},
}
