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

type scenesMassage struct {
	dx    int
	dy    int
	index int
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
		dx := scenesmassage.dx
		dy := scenesmassage.dy
		Types := scenesmassage.Type
		scenes := &Scenes{
			Collider: collision.NewCollider(float64(dx), float64(dy), float64(scenesImages[Types].Bounds().Dx()), float64(scenesImages[Types].Bounds().Dy())),
			Image:    scenesImages[Types],
			index:    scenesmassage.index,
			Type:     Types,
			Hp:       scenesmassage.Hp,
		}
		globalScenes[scenesmassage.index] = scenes
	}
}

type networkServer struct{}

func (a *networkServer) Send() string {
	massage := []scenesMassage{}
	for _, scenes := range globalScenes {
		massage = append(massage, scenesMassage{
			dx:    int(scenes.Collider.Position.X),
			dy:    int(scenes.Collider.Position.Y),
			index: scenes.index,
			Type:  scenes.Type,
			Hp:    scenes.Hp,
		})
	}
	return json.MarshalToString(massage)
}

func (a *networkServer) Receive(m string) {

}
