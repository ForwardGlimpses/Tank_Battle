package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Bouncer struct {
	Object *resolv.Object
	Speed  resolv.Vector
}

type Game struct {
	Width, Height int
	Bouncers      []*Bouncer
	Geometry      []*resolv.Object
	Space         *resolv.Space
}

func NewGame() *Game {

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("resolv test")

	g := &Game{
		Width:  640,
		Height: 360,
	}

	gw := float64(g.Width)
	gh := float64(g.Height)
	cellSize := 8

	g.Space = resolv.NewSpace(g.Width, g.Height, cellSize, cellSize)

	g.Geometry = []*resolv.Object{
		resolv.NewObject(0, 0, 16, gh),
		resolv.NewObject(gw-16, 0, 16, gh),
		resolv.NewObject(0, 0, gw, 16),
		resolv.NewObject(0, gh-24, gw, 32),
	}

	g.Space.Add(g.Geometry...)

	for i := 0; i < 4; i++ {
		g.SpawnObject()
	}

	return g
}

func (g *Game) SpawnObject() {

	bouncer := &Bouncer{
		Object: resolv.NewObject(0, 0, 2, 2),
		Speed: resolv.NewVector(
			(rand.Float64()*8)-4,
			(rand.Float64()*8)-4,
		),
	}

	g.Space.Add(bouncer.Object)

	var c *resolv.Cell
	for c == nil {
		rx := rand.Intn(g.Space.Width())
		ry := rand.Intn(g.Space.Height())
		c = g.Space.Cell(rx, ry)
		if c.Occupied() {
			c = nil
		} else {
			bouncer.Object.Position.X, bouncer.Object.Position.Y = g.Space.SpaceToWorld(c.X, c.Y)
		}
	}

	g.Bouncers = append(g.Bouncers, bouncer)

}

func (g *Game) Update() error {

	for _, b := range g.Bouncers {

		b.Speed.Y += 0.1

		dx := b.Speed.X
		dy := b.Speed.Y

		if check := b.Object.Check(dx, 0); check != nil {
			contact := check.ContactWithCell(check.Cells[0])
			dx = contact.X
			b.Speed.X *= -1
		}

		b.Object.Position.X += dx

		if check := b.Object.Check(0, dy); check != nil {
			contact := check.ContactWithCell(check.Cells[0])
			dy = contact.Y
			b.Speed.Y *= -1
		}

		b.Object.Position.Y += dy

		b.Object.Update()

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, o := range g.Geometry {
		ebitenutil.DrawRect(screen, o.Position.X, o.Position.Y, o.Size.X, o.Size.Y, color.RGBA{60, 60, 60, 255})
	}

	for _, b := range g.Bouncers {
		o := b.Object
		ebitenutil.DrawRect(screen, o.Position.X, o.Position.Y, o.Size.X, o.Size.Y, color.RGBA{0, 80, 255, 255})
	}

}

func (g *Game) Layout(w, h int) (int, int) {
	return g.Width, g.Height
}

func main() {
	ebiten.RunGame(NewGame())
}
