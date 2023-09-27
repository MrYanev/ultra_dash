package internal

import (
	"github.com/bytearena/ecs"
)

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	player := manager.NewComponent()
	position := manager.NewComponent()
	renderable := manager.NewComponent()
	movable := manager.NewComponent()

	//Creating a player entity which has all the structs that we
	//created earlier, also we are asigning it with the corresponging
	//tags
	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, Renderable{}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: 40,
			Y: 25,
		})

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	return manager, tags
}
