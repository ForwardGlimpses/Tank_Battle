package scorer

import (
	"fmt"
	"image/color"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

var (
	GlobalScore = map[int]*Score{}
)

func New(Playerindex int) {
	score := &Score{
		PlayerIndex: Playerindex,
		PlayerScore: 0,
	}
	GlobalScore[score.PlayerIndex] = score
}

func init() {
	tankbattle.RegisterDraw(Draw, 100)
}

type Score struct {
	PlayerIndex int
	PlayerScore int
}

func (s *Score) AddPoints(points int) {
	s.PlayerScore += points
}

func AddPoints(playerindex int, points int) {
	score := GlobalScore[playerindex]
	score.AddPoints(points)
}

func Draw(screen *ebiten.Image) {
	for playerIndex, score := range GlobalScore {
		score.Draw(screen, playerIndex)
	}
}

func (s *Score) Draw(screen *ebiten.Image, playerIndex int) {
	scoreText := fmt.Sprintf("PalyerIndex: %d,PalyerScore: %d", s.PlayerIndex, s.PlayerScore)
	// 使用基础字体绘制白色文字，绘制在屏幕上
	text.Draw(screen, scoreText, basicfont.Face7x13, 200, playerIndex*100, color.White)
}
