package internal

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/norendren/go-fov/fov"
)

type TileType int

const (
	WALL TileType = iota
	FLOOR
)

// Level holds the tile formation for the level
type Level struct {
	Tiles         []*MapTile
	Rooms         []Rect
	PlayerVisible *fov.View
}

// NewLevel creates a new level
func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	l.PlayerVisible = fov.New()
	return l
}

// A structure to hold all the individual tiles and squares
type MapTile struct {
	PixelX     int
	PixelY     int
	Blocked    bool
	IsRevealed bool
	TileType   TileType
	Image      *ebiten.Image
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	//Draw the map
	//If statement to check for the visible range of the player
	gd := NewGameData()

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			idx := level.GetIndexFromXY(x, y)
			tile := level.Tiles[idx]
			isVis := level.PlayerVisible.IsVisible(x, y)
			if isVis {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				screen.DrawImage(tile.Image, op)
				level.Tiles[idx].IsRevealed = true
			} else if tile.IsRevealed == true {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
				op.ColorScale.Scale(0.8, 1, 1, 0.25)
				screen.DrawImage(tile.Image, op)
			}
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
func (level *Level) CreateTiles() []*MapTile {
	gd := NewGameData()
	tiles := make([]*MapTile, gd.ScreenHeight*gd.ScreenWidth)
	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}
			tile := MapTile{
				PixelX:     x * gd.TitleWidth,
				PixelY:     y * gd.TitleHeight,
				Blocked:    true,
				Image:      wall,
				IsRevealed: false,
				TileType:   WALL,
			}
			tiles[index] = &tile
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
			level.Tiles[index].TileType = FLOOR
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

// GenerateLevelTiles creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	tiles := level.CreateTiles()
	level.Tiles = tiles
	contains_rooms := false

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(gd.ScreenHeight - h - 1)
		new_room := NewRect(x, y, w, h)

		okToAdd := true

		for _, otherRoom := range level.Rooms {
			if new_room.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(new_room)
			if contains_rooms {
				newX, newY := new_room.Center()
				prevX, prevY := level.Rooms[len(level.Rooms)-1].Center()
				coinflip := GetDiceRoll(2)
				if coinflip == 2 {
					level.createHorizontalTunnel(prevX, newX, prevY)
					level.createVerticalTunnel(prevY, newY, newX)

				} else {
					level.createHorizontalTunnel(prevX, newX, newY)
					level.createVerticalTunnel(prevY, newY, prevX)
				}

			}

			level.Rooms = append(level.Rooms, new_room)
			contains_rooms = true
		}
	}

}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// A function for creating horizontal tunnels
func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	gd := NewGameData()
	for x := min(x1, x2); x < max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

// A function for creating vertical tunnels
func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	gd := NewGameData()
	for y := min(y1, y2); y < max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)

		if index > 0 && index < gd.ScreenWidth*gd.ScreenHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = FLOOR
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

// Checking for walls for the vision radius
func (level Level) InBounds(x, y int) bool {
	gd := NewGameData()
	if x < 0 || x > gd.ScreenWidth || y < 0 || y > gd.ScreenHeight {
		return false
	}
	return true
}

// TODO: Change this to check for WALL, not blocked
func (level Level) IsOpaque(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	return level.Tiles[idx].TileType == WALL
}
