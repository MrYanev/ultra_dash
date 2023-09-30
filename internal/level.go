package internal

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Level holds the tile formation for the level
type Level struct {
	Tiles []MapTile
	Rooms []Rect
}

// NewLevel creates a new level
func NewLevel() Level {
	l := Level{}
	tiles := l.CreateTiles()
	l.Tiles = tiles
	return l
}

// A structure to hold all the individual tiles and squares
type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	//Draw the map
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// This function gets the index of the map array from given X and Y
// tile coordinates. These coordinates are logical tiles not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// This function creates the map of tiles
// First if handles the walls around the edges
// Else handles the flooring tiles
func (level *Level) CreateTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)
	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}
			tile := MapTile{
				PixelX:  x * gd.TitleWidth,
				PixelY:  y * gd.TitleHeight,
				Blocked: true,
				Image:   wall,
			}
			tiles[index] = tile
		}
	}
	return tiles
}

// A function for creating the rooms
func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}
