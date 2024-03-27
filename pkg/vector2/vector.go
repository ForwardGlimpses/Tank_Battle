package vector2

// TODO: 如果有需要增加的函数，可以参考 godot Vector2 实现
// https://docs.godotengine.org/zh-cn/4.x/classes/class_vector2i.html
// godot 重载了运算符，go 里面需要对每种运算都实现一个函数

// 向量既可以表示坐标，也可以表示速度
type Vector2 struct {
	X, Y int
}

func New(x, y int) *Vector2 {
	return &Vector2{
		X: x,
		Y: y,
	}
}

func (v *Vector2) Add(other *Vector2) *Vector2 {
	return New(v.X+other.X, v.Y+other.Y)
}

func (v *Vector2) Sub(other *Vector2) *Vector2 {
	return New(v.X-other.X, v.Y-other.Y)
}

func (v *Vector2) MulScalar(scalar int) *Vector2 {
	return New(v.X*scalar, v.Y*scalar)
}

func (v *Vector2) Value() (int, int) {
	return v.X, v.Y
}

func (v *Vector2) ValueFloat64() (float64, float64) {
	return float64(v.X), float64(v.Y)
}
