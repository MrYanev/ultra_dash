package internal

//Holds all the levels
type GameMap struct {
	Dungeons []Dungeon
}

//This function creates set of maps for the entire game
//			(Game Map (Dungeons (Levels)))
func NewGameMap() GameMap {
	//Return a game map for a single level
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)
	//Add the level to the dungeon
	d := Dungeon{Name: "Dungeon 1", Levels: levels}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)
	//Add the dungeon to the Game Map
	gm := GameMap{Dungeons: dungeons}
	return gm
}
