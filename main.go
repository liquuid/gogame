package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/liquuid/gogame/window"
	game "github.com/liquuid/gogame/game"
	_ "image/jpeg"
	"log"
)

var (
	windowDecorated = flag.Bool("windowdecorated", true, "whether the window is decorated")
	terminated = errors.New("terminated")
	zoomScale = 1
	sprites = []*game.Sprite{}
	//sprites = &game.Sprites{make([]*game.Sprite, 10), 10}

)
func init() {
	
}

const (
	initScreenWidth  = 512
	initScreenHeight = 320
	initScreenScale  = 2
)

func update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return terminated
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ToggleZoom()
	}

	window.Update(screen)
	//pl.Tick(screen)

	return nil
}

func ToggleZoom() {
	fmt.Println("Zoomtoggled")
}

func main() {
	flag.Parse()

	fmt.Printf("Device scale factor: %0.2f\n", ebiten.DeviceScaleFactor())

	w, h := ebiten.ScreenSizeInFullscreen()
	fmt.Printf("Screen size in fullscreen: %d, %d\n", w, h)

	ebiten.SetWindowDecorated(*windowDecorated)

	//pl.Init(100,100, 0,77,90,77,4)
	fmt.Println(len(sprites))
	for i := 0; i < 5 ; i++{
		sprites = append(sprites, &game.Player{})
	}

	if err := ebiten.Run(update, initScreenWidth, initScreenHeight, initScreenScale, "Window Size (Ebiten Demo)"); err != nil && err != terminated {
		log.Fatal(err)
	}
}
