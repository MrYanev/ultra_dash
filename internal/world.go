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
	userMessage *ecs.Component
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
	userMessage = manager.NewComponent()

	//Adding the image for the player
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	monsterImg, _, err := ebitenutil.NewImageFromFile("assets/skelly.png")
	if err != nil {
		log.Fatal(err)
	}

	orcImg, _, err := ebitenutil.NewImageFromFile("assets/orc.png")
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
			Name:          "Battle Axe",
			MinimumDamage: 5,
			MaximumDamage: 13,
			ToHitBonus:    2,
		}).
		AddComponent(armor, &Armor{
			Name:       "Obsidian Armor",
			Defense:    7,
			ArmorClass: 13,
		}).
		AddComponent(userMessage, &UserMessage{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		}).
		AddComponent(name, &Name{Label: "Player"})

	//Add a Monster in each room except the player's room
	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()

			//Flip a coin to see what to add...
			mobSpawn := GetDiceRoll(2)

			if mobSpawn == 1 {
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{
						Image: orcImg,
					}).
					AddComponent(position, &Position{
						X: mX,
						Y: mY,
					}).
					AddComponent(health, &Health{
						MaxHealth:     30,
						CurrentHealth: 30,
					}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:          "Machete",
						MinimumDamage: 4,
						MaximumDamage: 11,
						ToHitBonus:    1,
					}).
					AddComponent(armor, &Armor{
						Name:       "Leather",
						Defense:    5,
						ArmorClass: 6,
					}).
					AddComponent(name, &Name{Label: "Orc"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			} else {
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
						MinimumDamage: 3,
						MaximumDamage: 10,
						ToHitBonus:    0,
					}).
					AddComponent(armor, &Armor{
						Name:       "Bone",
						Defense:    3,
						ArmorClass: 4,
					}).
					AddComponent(name, &Name{Label: "Skeleton"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			}

		}
	}
	//Adding a view for the player
	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name, userMessage)
	tags["players"] = players

	//Adding a view for all rendable objects in the game
	//And where each should be (NB!!!)
	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	//Adding a view for monsters for faster access
	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name, userMessage)
	tags["monsters"] = monsters

	messengers := ecs.BuildTag(userMessage)
	tags["messengers"] = messengers

	return manager, tags
}
