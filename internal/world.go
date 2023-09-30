package internal

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	position   *ecs.Component
	renderable *ecs.Component
)

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()

	//Adding the image for the player
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	//Creating a player entity which has all the structs that we
	//created earlier, also we are asigning it with the corresponging
	//tags
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		})

	//Adding a view for the player
	players := ecs.BuildTag(player, position)
	tags["players"] = players

	//Adding a view for all rendable objects in the game
	//And where each should be (NB!!!)
	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	return manager, tags
}
