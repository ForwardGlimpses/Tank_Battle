package types

// 可被伤害的 struct 需要实现此接口
// GetCamp() string 获取阵营信息，相同阵营的不受影响。创建 tank 和 bullet 时需要携带阵营信息并存储
// TakeDamage(int) 处理受伤操作，实例各自创建自己的受伤函数，由子弹调用
type TakeDamage interface {
	GetCamp() string
	TakeDamage(int)
}

// 当前只有一个空函数，不做任何处理，实现此函数即认为是障碍物
type Obstacle interface {
	Obstacle()
	TankIsPassable() bool
	BulletIsPassable() bool
}
