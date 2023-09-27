package internal

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
