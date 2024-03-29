
#### 坦克

坦克至少需要包含字段：
+ 图片[^1]
+ 坐标
+ 血量
+ 武器（可以等实现武器类再补）

至少需要包含方法：
+ 移动
+ 攻击（可以等实现武器类再补）
+ 绘制

[^1]: 图片的引用方式参考 ebitn [blend 示例](https://github.com/hajimehoshi/ebiten/blob/main/examples/blend/main.go#L56)，图片尽量小一些，能看出区别即可。图片放在 assets 文件夹中。

#### 玩家

玩家在游戏初始化时创建

玩家至少包含字段：
+ 坦克，源于坦克类
+ 名称

至少需要包含方法：
+ 更新，获取键盘输入，确定当前帧玩家动作
+ 绘制，可以调用坦克类的绘制函数

#### 敌人

敌人每帧循环时确定是否创建

敌人至少包含字段：
+ 坦克，源于坦克类

至少需要包含方法：
+ 更新，使用随机数决定行为即可
+ 绘制，可以调用坦克类的绘制函数

#### 武器&子弹

武器为接口类型，包含“攻击”函数，至少设计两种武器

子弹至少包含字段：
+ 图片
+ 坐标
+ 伤害

子弹至少包含方法：
+ 更新，子弹的更新方法中需要判断碰撞，然后决定子弹是否生效
+ 绘制，
