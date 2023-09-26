package internal

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Tiles []MapTile
}

func NewGame() *Game {
	g := &Game{}
	g.Tiles = CreateTiles()
	return g
}

//Draw is called each draw cycle
//Here we build the visuals
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the map
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := g.Tiles[GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

//Update executes each tic.
func (g *Game) Update() error {
	return nil
}

//Layout returns the screen dimensions
func (g *Game) Layout(ScreenWidth, ScreenHeight int) (int, int) {
	return 1280, 800
}
