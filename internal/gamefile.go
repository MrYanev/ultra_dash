package internal

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Map GameMap
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	return g
}

//Draw is called each draw cycle
//Here we build the visuals
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the map
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)
}

//Update executes each tic.
func (g *Game) Update() error {
	return nil
}

//Layout returns the screen dimensions
func (g *Game) Layout(ScreenWidth, ScreenHeight int) (int, int) {
	gd := NewGameData()
	return gd.TitleWidth * gd.ScreenWidth, gd.TitleHeight * gd.ScreenHeight
}
