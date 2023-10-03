package internal

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	position    *ecs.Component
	renderable  *ecs.Component
	monster     *ecs.Component
	health      *ecs.Component
	meleeWeapon *ecs.Component
	armor       *ecs.Component
	name        *ecs.Component
)

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	player := manager.NewComponent()
	monster := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	health = manager.NewComponent()
	meleeWeapon = manager.NewComponent()
	armor = manager.NewComponent()
	name = manager.NewComponent()

	//Adding the image for the player
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	monsterImg, _, err := ebitenutil.NewImageFromFile("assets/skelly.png")
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
		}).
		AddComponent(health, &Health{
			MaxHealth:     30,
			CurrentHealth: 30,
		}).
		AddComponent(meleeWeapon, &MeleeWeapon{
			Name:          "Fist",
			MinimumDamage: 1,
			MaximumDamage: 3,
			ToHitBonus:    2,
		}).
		AddComponent(armor, &Armor{
			Name:       "Burlap Sack",
			Defense:    1,
			ArmorClass: 1,
		}).
		AddComponent(name, &Name{Label: "Player"})

	//Add a Monster in each room except the player's room
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()
			manager.NewEntity().
				AddComponent(monster, &Monster{}).
				AddComponent(renderable, &Renderable{
					Image: monsterImg,
				}).
				AddComponent(position, &Position{
					X: mX,
					Y: mY,
				}).
				AddComponent(health, &Health{
					MaxHealth:     10,
					CurrentHealth: 10,
				}).
				AddComponent(meleeWeapon, &MeleeWeapon{
					Name:          "Short Sword",
					MinimumDamage: 2,
					MaximumDamage: 6,
					ToHitBonus:    0,
				}).
				AddComponent(armor, &Armor{
					Name:       "Bone",
					Defense:    3,
					ArmorClass: 4,
				}).
				AddComponent(name, &Name{Label: "Skeleton"})

		}
	}

	//Adding a view for the player
	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name)
	tags["players"] = players

	//Adding a view for all rendable objects in the game
	//And where each should be (NB!!!)
	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	//Adding a view for monsters for faster access
	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name)
	tags["monsters"] = monsters

	return manager, tags
}
