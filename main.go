package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/liquuid/gogame/window"
	_ "image/jpeg"
	"log"
)

var (
	windowDecorated = flag.Bool("windowdecorated", true, "whether the window is decorated")
)

func init() {
	
}

const (
	initScreenWidth  = 512
	initScreenHeight = 320
	initScreenScale  = 2
)

var (

	terminated = errors.New("terminated")

)

func update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return terminated
	}

	window.Update(screen)


	return nil
}

func main() {
	flag.Parse()

	fmt.Printf("Device scale factor: %0.2f\n", ebiten.DeviceScaleFactor())

	w, h := ebiten.ScreenSizeInFullscreen()
	fmt.Printf("Screen size in fullscreen: %d, %d\n", w, h)

	ebiten.SetWindowDecorated(*windowDecorated)


	if err := ebiten.Run(update, initScreenWidth, initScreenHeight, initScreenScale, "Window Size (Ebiten Demo)"); err != nil && err != terminated {
		log.Fatal(err)
	}
}
