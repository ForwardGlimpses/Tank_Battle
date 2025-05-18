# 坦克大战

本项目是一个基于 [Ebiten](https://ebiten.org/) 的 Golang 游戏项目，旨在通过开发坦克大战游戏，熟悉 Go 语言语法和项目构建流程。项目重心在于工程实践与架构设计，而非游戏玩法的深度优化。

## 项目结构

```
.
├── assets/         # 游戏资源（图片等，使用 embed 静态导入）
├── cmd/            # 主程序入口
├── configs/        # 配置文件
├── docs/           # 项目文档
├── examples/       # 示例代码
├── pkg/            # 主要业务逻辑
├── test/           # 集成测试
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 主要功能

- 基础战斗单元（坦克、玩家、敌人、武器&子弹）
- 地图与障碍物
- 配置文件支持
- 计分与胜利条件
- 联机功能

## 快速开始

1. **克隆项目**
    ```sh
    git clone https://github.com/yourname/tank-battle.git
    cd tank-battle
    ```

2. **构建并运行**
    ```sh
    make run
    ```
    或直接使用 Go 命令：
    ```sh
    go run ./cmd/tankbattle
    ```

3. **资源说明**
    - 所有图片等资源位于 [assets/](assets/README.md)
    - 配置文件位于 [configs/](configs/README.md)


## 贡献

欢迎提交 issue 和 PR，完善项目功能或修复 bug。

## License

[MIT](LICENSE)
