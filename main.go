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
	zoomScale = 1.0
	sprites = make([]*game.Sprite, 10)
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

	for i := 0; i < 10 ; i++ {
		sprites[i].ScaleTo(zoomScale)
		sprites[i].Tick(screen)  // .Tick(screen)
	}

	return nil
}

func ToggleZoom() {
	if zoomScale == 1.0{
		zoomScale = 0.5
	} else {
		zoomScale = 1
	}
	fmt.Println("Zoomtoggled")
}

func main() {
	flag.Parse()

	fmt.Printf("Device scale factor: %0.2f\n", ebiten.DeviceScaleFactor())

	w, h := ebiten.ScreenSizeInFullscreen()
	fmt.Printf("Screen size in fullscreen: %d, %d\n", w, h)

	ebiten.SetWindowDecorated(*windowDecorated)
	fmt.Println(len(sprites))

	for i := 0; i < 10 ; i++{

		x := float64(100.0+(10*i))
		y := float64(100.0+(10*i))
		sprites[i] = new(game.Sprite)
		sprites[i].Init(x, y, 1, 0,77,90,77,4)
	}

	if err := ebiten.Run(update, initScreenWidth, initScreenHeight, initScreenScale, "Window Size (Ebiten Demo)"); err != nil && err != terminated {
		log.Fatal(err)
	}

}
