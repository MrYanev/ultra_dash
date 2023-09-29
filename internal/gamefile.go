package internal

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map       GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

func NewGame() *Game {
	g := &Game{}
	world, tags := InitializeWorld()
	g.Map = NewGameMap()
	g.WorldTags = tags
	g.World = world
	return g
}

// Draw is called each draw cycle
// Here we build the visuals
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the map
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
}

// Update executes each tic.
func (g *Game) Update() error {
	TryMovePlayer(g)
	return nil
}

// Layout returns the screen dimensions
func (g *Game) Layout(ScreenWidth, ScreenHeight int) (int, int) {
	gd := NewGameData()
	return gd.TitleWidth * gd.ScreenWidth, gd.TitleHeight * gd.ScreenHeight
}
