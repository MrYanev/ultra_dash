package main

import (
	_ "image/png"
	"log"

	"github.com/MrYanev/ultra_dash/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := internal.NewGame()
	ebiten.SetWindowTitle("Ultra DASH")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
