package internal

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TitleWidth   int
	TitleHeight  int
}

// Constructor to set the tile size
func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TitleWidth:   16,
		TitleHeight:  16,
	}
	return g
}

// A structure to hold all the individual tiles and squares
type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

// This function gets the index of the map array from given X and Y
// tile coordinates. These coordinates are logical tiles not pixels.
func GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// This function creates the map of tiles
// First if handles the walls around the edges
// Else handles the flooring tiles
func CreateTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, 0)

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {
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
				tiles = append(tiles, tile)
			} else {
				floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					PixelX:  x * gd.TitleWidth,
					PixelY:  y * gd.TitleHeight,
					Blocked: false,
					Image:   floor,
				}
				tiles = append(tiles, tile)
			}
		}
	}
	return tiles
}
