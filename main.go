package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

// Update logical state
func (g *Game) Update() error {
	mouseIsClicked := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	if mouseIsClicked {
		fmt.Print("Cursor position is : ")
		fmt.Print(ebiten.CursorPosition())
		fmt.Println()
	}
	return nil
}

// Render screen
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "This is the beginning of Game of Life")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
