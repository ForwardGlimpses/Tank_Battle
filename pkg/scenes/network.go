package scenes

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/json"
)

func init() {
	network.RegisterClient("scenes", &neteworkClient{})
	network.RegisterServer("scenes", &networkServer{})
}

var (
	scenesDetect = map[int]bool{}
)

type scenesMassage struct {
	Dx    int
	Dy    int
	Index int
	Type  ScenesType
	Hp    int
}

type neteworkClient struct{}

func (a *neteworkClient) Send() string {
	return ""
}

func (a *neteworkClient) Receive(m string) {
	massage := []scenesMassage{}
	json.Unmarshal([]byte(m), &massage)
	for _, scenesmassage := range massage {
		dx := scenesmassage.Dx
		dy := scenesmassage.Dy
		Types := scenesmassage.Type
		_,ok:= scenesDetect[scenesmassage.Index]
		if ok {
			globalScenes[scenesmassage.Index].Hp = scenesmassage.Hp
		}else{
			scenes := &Scenes{
				Collider: collision.NewCollider(float64(dx), float64(dy), float64(scenesImages[Types].Bounds().Dx()), float64(scenesImages[Types].Bounds().Dy())),
				Image:    scenesImages[Types],
				index:    scenesmassage.Index,
				Type:     Types,
				Hp:       scenesmassage.Hp,
			}
			globalScenes[scenesmassage.Index] = scenes
			scenesDetect[scenesmassage.Index] = true
		}

	}
}

type networkServer struct{}

func (a *networkServer) Send() string {
	massage := []scenesMassage{}
	for _, scenes := range globalScenes {
		massage = append(massage, scenesMassage{
			Dx:    int(scenes.Collider.Position.X),
			Dy:    int(scenes.Collider.Position.Y),
			Index: scenes.index,
			Type:  scenes.Type,
			Hp:    scenes.Hp,
		})
	}
	return json.MarshalToString(massage)
}

func (a *networkServer) Receive(m string) {

}
