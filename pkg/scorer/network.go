package scorer

import (

	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/json"
)

func init() {
	network.RegisterClient("scorer", &neteworkClient{})
	network.RegisterServer("scorer", &networkServer{})
}

type scorerMassage struct {
	PlayerIndex int 
	PlayerScore int 
}

type neteworkClient struct{}

func (a *neteworkClient) Send() string {
	return ""
}

func (a *neteworkClient) Receive(m string) {
	massage := []scorerMassage{}
	json.Unmarshal([]byte(m), &massage)
	//fmt.Println("接收数据:", massage)
	for _, scoremassage := range massage {
		score, ok := GlobalScore[scoremassage.PlayerIndex]
		if ok {
			score.PlayerScore = scoremassage.PlayerScore
		} else {
			scorer := &Score{
				PlayerScore: scoremassage.PlayerScore,
				PlayerIndex: scoremassage.PlayerIndex,
			}
			GlobalScore[scorer.PlayerIndex] = scorer
		}
	}
}

type networkServer struct{}

func (a *networkServer) Send() string {
	massage := []scorerMassage{}
	for _, score := range GlobalScore {
		massage = append(massage,scorerMassage{
			PlayerIndex: score.PlayerIndex,
			PlayerScore: score.PlayerScore,
		})
	}
	date := json.MarshalToString(massage)
	return date

}

func (a *networkServer) Receive(m string) {
	//fmt.Println("--------------")
}
