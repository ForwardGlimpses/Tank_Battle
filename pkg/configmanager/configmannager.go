package configmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
)

func init() {
	tankbattle.RegisterInit(Init, 1)
}

func Init() error {

	//cfg, err := LoadConfig("./configs/config.json")
	cfg, err := LoadConfig("C:\\Users\\乔书祥\\Desktop\\远程文件库\\Tank_Battle\\configs\\config.json")
	if err != nil {
		return err
	}

	sizeX, sizeY := cfg.Window.Width, cfg.Window.Height
	collision.Init(sizeX, sizeY, 2, 2)

	config.C = *cfg
	return nil
}

// 读取配置文件
func LoadConfig(filename string) (*config.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %v", err)
	}

	var cfg config.Config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse config file: %v", err)
	}

	return &cfg, nil
}
